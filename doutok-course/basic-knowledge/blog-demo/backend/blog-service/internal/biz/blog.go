package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// Blog 博客业务对象
type Blog struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BlogRepo 博客存储接口
type BlogRepo interface {
	List(ctx context.Context) ([]*Blog, error)
	Get(ctx context.Context, id string) (*Blog, error)
	Create(ctx context.Context, blog *Blog) (*Blog, error)
	Update(ctx context.Context, id string, blog *Blog) (*Blog, error)
	Delete(ctx context.Context, id string) error
}

// BlogUsecase 博客用例
type BlogUsecase struct {
	repo BlogRepo
	log  *log.Helper
}

// NewBlogUsecase 创建博客用例
func NewBlogUsecase(repo BlogRepo, logger log.Logger) *BlogUsecase {
	return &BlogUsecase{repo: repo, log: log.NewHelper(logger)}
}

// List 获取所有博客
func (uc *BlogUsecase) List(ctx context.Context) ([]*Blog, error) {
	return uc.repo.List(ctx)
}

// Get 获取单个博客
func (uc *BlogUsecase) Get(ctx context.Context, id string) (*Blog, error) {
	return uc.repo.Get(ctx, id)
}

// Create 创建博客
func (uc *BlogUsecase) Create(ctx context.Context, blog *Blog) (*Blog, error) {
	return uc.repo.Create(ctx, blog)
}

// Update 更新博客
func (uc *BlogUsecase) Update(ctx context.Context, id string, blog *Blog) (*Blog, error) {
	return uc.repo.Update(ctx, id, blog)
}

// Delete 删除博客
func (uc *BlogUsecase) Delete(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}
