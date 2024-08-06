package biz

import (
	"context"
	"helloworld/pkg/db"

	v1 "helloworld/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter is a Greeter model.
type Greeter struct {
	Hello string `gorm:"column:hello;type:varchar(20)"`
}

func (*Greeter) TableName() string {
	return "greater"
}

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	Save(context.Context, *Greeter) (*Greeter, error)
	Update(context.Context, *Greeter) (*Greeter, error)
	FindByID(context.Context, int64) (*Greeter, error)
	ListByHello(context.Context, string) ([]*Greeter, error)
	ListAll(context.Context) ([]*Greeter, error)
}

// GreeterUsecase is a Greeter usecase.
type GreeterUsecase struct {
	db   *db.DBClient
	repo GreeterRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo GreeterRepo, db *db.DBClient, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{db: db, repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *Greeter) (*Greeter, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	var (
		greater *Greeter
		err     error
	)
	//err = uc.db.ExecTx(ctx, func(ctx context.Context) error {
	//	// 更新所有 hello 为 hello + "updated"，且插入新的 hello
	//	greater, err = uc.repo.Save(ctx, g)
	//	_, err = uc.repo.Update(ctx, g)
	//	return err
	//})
	greater, err = uc.repo.Save(ctx, g)
	_, err = uc.repo.Update(ctx, g)
	if err != nil {
		return nil, err
	}
	return greater, nil
}
