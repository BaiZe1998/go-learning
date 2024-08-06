package db

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"helloworld/internal/conf"
)

type DBClient struct {
	db *gorm.DB
}

func NewDBClient(c *conf.Data) (*DBClient, error) {
	dsn := c.Database.Source
	if c.Database.Driver == "mysql" {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			SkipDefaultTransaction: true,
		})
		if err != nil {
			return nil, err
		}
		return &DBClient{db: db}, nil
	}
	return nil, nil
}

func (c *DBClient) GetDB() *gorm.DB {
	return c.db
}

type Transaction interface {
	ExecTx(context.Context, func(ctx context.Context) error) error
}

type contextTxKey struct{}

// ExecTx gorm Transaction
func (c *DBClient) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return c.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

func (c *DBClient) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return c.db
}
