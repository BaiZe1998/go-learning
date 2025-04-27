package service

import (
	"context"

	v1 "blog-service/api/blog/v1"
	"blog-service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BlogService struct {
	v1.UnimplementedBlogServer

	uc  *biz.BlogUsecase
	log *log.Helper
}

// NewBlogService 创建博客服务
func NewBlogService(uc *biz.BlogUsecase, logger log.Logger) *BlogService {
	return &BlogService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *BlogService) ListBlogs(ctx context.Context, req *v1.ListBlogsRequest) (*v1.ListBlogsReply, error) {
	blogs, err := s.uc.List(ctx)
	if err != nil {
		return nil, err
	}

	reply := &v1.ListBlogsReply{
		Blogs: make([]*v1.BlogInfo, 0, len(blogs)),
	}

	for _, blog := range blogs {
		reply.Blogs = append(reply.Blogs, &v1.BlogInfo{
			Id:        blog.ID,
			Title:     blog.Title,
			Content:   blog.Content,
			UpdatedAt: timestamppb.New(blog.UpdatedAt),
		})
	}

	return reply, nil
}

func (s *BlogService) GetBlog(ctx context.Context, req *v1.GetBlogRequest) (*v1.GetBlogReply, error) {
	blog, err := s.uc.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	s.log.Infof("blog = %v\n", blog)

	return &v1.GetBlogReply{
		Blog: &v1.BlogInfo{
			Id:        blog.ID,
			Title:     blog.Title,
			Content:   blog.Content,
			UpdatedAt: timestamppb.New(blog.UpdatedAt),
		},
	}, nil
}

func (s *BlogService) CreateBlog(ctx context.Context, req *v1.CreateBlogRequest) (*v1.CreateBlogReply, error) {
	s.log.Infof("CreateBlog req ====== %v\n", req)
	blog, err := s.uc.Create(ctx, &biz.Blog{
		Title:   req.Blog.Title,
		Content: req.Blog.Content,
	})
	if err != nil {
		return nil, err
	}

	return &v1.CreateBlogReply{
		Blog: &v1.BlogInfo{
			Id:        blog.ID,
			Title:     blog.Title,
			Content:   blog.Content,
			UpdatedAt: timestamppb.New(blog.UpdatedAt),
		},
	}, nil
}

func (s *BlogService) UpdateBlog(ctx context.Context, req *v1.UpdateBlogRequest) (*v1.UpdateBlogReply, error) {
	blog, err := s.uc.Update(ctx, req.Id, &biz.Blog{
		Title:   req.Blog.Title,
		Content: req.Blog.Content,
	})
	if err != nil {
		return nil, err
	}

	return &v1.UpdateBlogReply{
		Blog: &v1.BlogInfo{
			Id:        blog.ID,
			Title:     blog.Title,
			Content:   blog.Content,
			UpdatedAt: timestamppb.New(blog.UpdatedAt),
		},
	}, nil
}

func (s *BlogService) DeleteBlog(ctx context.Context, req *v1.DeleteBlogRequest) (*v1.DeleteBlogReply, error) {
	err := s.uc.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.DeleteBlogReply{}, nil
}
