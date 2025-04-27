package data

import (
	"context"
	"strconv"
	"time"

	"blog-service/internal/biz"
	"blog-service/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

type blogRepo struct {
	data *Data
	log  *log.Helper
}

// NewBlogRepo .
func NewBlogRepo(data *Data, logger log.Logger) biz.BlogRepo {
	return &blogRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *blogRepo) List(ctx context.Context) ([]*biz.Blog, error) {
	blogs, err := r.data.blogStore.List()
	if err != nil {
		return nil, err
	}

	result := make([]*biz.Blog, 0, len(blogs))
	for _, b := range blogs {
		result = append(result, &biz.Blog{
			ID:        b.ID,
			Title:     b.Title,
			Content:   b.Content,
			UpdatedAt: b.UpdatedAt,
		})
	}
	return result, nil
}

func (r *blogRepo) Get(ctx context.Context, id string) (*biz.Blog, error) {
	blog, err := r.data.blogStore.Get(id)
	if err != nil {
		return nil, err
	}
	return &biz.Blog{
		ID:        blog.ID,
		Title:     blog.Title,
		Content:   blog.Content,
		UpdatedAt: blog.UpdatedAt,
	}, nil
}

func (r *blogRepo) Create(ctx context.Context, blog *biz.Blog) (*biz.Blog, error) {
	blogModel := &model.Blog{
		ID:        strconv.FormatInt(time.Now().UnixNano(), 10),
		Title:     blog.Title,
		Content:   blog.Content,
		UpdatedAt: time.Now(),
	}

	r.log.Infof("blogRepo Create blogModel %v\n", blogModel)
	if err := r.data.blogStore.Save(blogModel); err != nil {
		return nil, err
	}

	return &biz.Blog{
		ID:        blogModel.ID,
		Title:     blogModel.Title,
		Content:   blogModel.Content,
		UpdatedAt: blogModel.UpdatedAt,
	}, nil
}

func (r *blogRepo) Update(ctx context.Context, id string, blog *biz.Blog) (*biz.Blog, error) {
	// 首先检查博客是否存在
	_, err := r.data.blogStore.Get(id)
	if err != nil {
		return nil, err
	}

	blogModel := &model.Blog{
		ID:        id,
		Title:     blog.Title,
		Content:   blog.Content,
		UpdatedAt: time.Now(),
	}

	if err := r.data.blogStore.Save(blogModel); err != nil {
		return nil, err
	}

	return &biz.Blog{
		ID:        blogModel.ID,
		Title:     blogModel.Title,
		Content:   blogModel.Content,
		UpdatedAt: blogModel.UpdatedAt,
	}, nil
}

func (r *blogRepo) Delete(ctx context.Context, id string) error {
	return r.data.blogStore.Delete(id)
}
