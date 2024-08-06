package data

import (
	"context"
	"fmt"

	"helloworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	result := r.data.db.DB(ctx).Create(g)
	return g, result.Error
}

func (r *greeterRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	result := r.data.db.DB(ctx).Model(&biz.Greeter{}).Where("hello = ?", g.Hello).Update("hello", g.Hello+"updated")
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("greeter %s not found", g.Hello)
	}
	return nil, fmt.Errorf("custom error")
	//return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*biz.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(ctx context.Context, hello string) ([]*biz.Greeter, error) {
	var greeters []*biz.Greeter
	result := r.data.db.DB(ctx).Where("hello = ?", hello).Find(&greeters)
	return greeters, result.Error
}

func (r *greeterRepo) ListAll(context.Context) ([]*biz.Greeter, error) {
	return nil, nil
}
