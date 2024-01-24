//   sqlc v1.16.0

package db

import (
	"context"
)

type Querier interface {
	DeleteAccount(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (Account, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) error
	createAccount(ctx context.Context, arg createAccountParams) (Account, error)
}

var _ Querier = (*Queries)(nil)
