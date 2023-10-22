package database

import "context"

type DatabaseInterface interface {
	Connect(ctx context.Context) error
}
