package service

import (
	"time"

	"github.com/seanyang20/banking/domain"
	"github.com/seanyang20/banking/dto"
	"github.com/seanyang20/banking/errs"
)

const dbTSLayout = "2006-01-02 15:04:05"

type AccountService interface {
	// NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

// hold reference to secondary port
type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	// err := req.Validate()
	// if err != nil {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	// a := domain.Account{
	// 	CustomerId: req.CustomerId,
	// 	// OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
	// 	OpeningDate: dbTSLayout,
	// 	AccountType: req.AccountType,
	// 	Amount:      req.Amount,
	// 	Status:      "1",
	// }
	// newAccount, err := s.repo.Save(a)
	// if err != nil {
	account := domain.NewAccount(req.CustomerId, req.AccountType, req.Amount)
	if newAccount, err := s.repo.Save(account); err != nil {
		return nil, err
		// }
	} else {
		return newAccount.ToNewAccountResponseDto(), nil
	}
	// response := newAccount.ToNewAccountResponseDto()

	// return &response, nil
}

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	// incoming request validation
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	// server side validation for checking the available balance in the account
	if req.IsTransactionTypeWithdrawal() {
		account, err := s.repo.FindBy(req.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Insufficient balance in the account")
		}
	}
	// if all is well, build the domain object & save the transaction
	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format(dbTSLayout),
	}
	transaction, appError := s.repo.SaveTransaction(t)
	if appError != nil {
		return nil, appError
	}
	response := transaction.ToDto()
	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
