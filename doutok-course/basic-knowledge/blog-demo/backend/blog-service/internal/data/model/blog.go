package model

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"
	"time"
)

// Blog 博客模型
type Blog struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BlogStore 博客存储接口
type BlogStore interface {
	List() ([]*Blog, error)
	Get(id string) (*Blog, error)
	Save(blog *Blog) error
	Delete(id string) error
}

// CSVBlogStore CSV实现的博客存储
type CSVBlogStore struct {
	filePath string
	mu       sync.RWMutex
}

// NewCSVBlogStore 创建CSV博客存储
func NewCSVBlogStore(filePath string) (*CSVBlogStore, error) {
	// 检查文件是否存在，不存在则创建
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}
		writer := csv.NewWriter(file)
		// 写入CSV表头
		if err := writer.Write([]string{"id", "title", "content", "updated_at"}); err != nil {
			return nil, err
		}
		writer.Flush()
		file.Close()
	}

	return &CSVBlogStore{
		filePath: filePath,
	}, nil
}

// 内部函数，不加锁，用于读取CSV文件内容
func (s *CSVBlogStore) readBlogsFromFile() ([]*Blog, error) {
	file, err := os.Open(s.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	blogs := make([]*Blog, 0, len(records)-1)
	// 跳过表头
	for i := 1; i < len(records); i++ {
		record := records[i]
		if len(record) < 4 {
			continue
		}

		updatedAt, err := time.Parse(time.RFC3339, record[3])
		if err != nil {
			updatedAt = time.Now()
		}

		blogs = append(blogs, &Blog{
			ID:        record[0],
			Title:     record[1],
			Content:   record[2],
			UpdatedAt: updatedAt,
		})
	}

	return blogs, nil
}

// List 列出所有博客
func (s *CSVBlogStore) List() ([]*Blog, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.readBlogsFromFile()
}

// Get 获取指定ID的博客
func (s *CSVBlogStore) Get(id string) (*Blog, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	blogs, err := s.readBlogsFromFile()
	if err != nil {
		return nil, err
	}

	for _, blog := range blogs {
		if blog.ID == id {
			return blog, nil
		}
	}

	return nil, fmt.Errorf("blog with ID %s not found", id)
}

// Save 保存或更新博客
func (s *CSVBlogStore) Save(blog *Blog) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	blogs, err := s.readBlogsFromFile()
	if err != nil {
		return err
	}

	// 更新现有博客或添加新博客
	found := false
	for i, b := range blogs {
		if b.ID == blog.ID {
			blogs[i] = blog
			found = true
			break
		}
	}

	if !found {
		blogs = append(blogs, blog)
	}

	return s.writeBlogs(blogs)
}

// Delete 删除指定ID的博客
func (s *CSVBlogStore) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	blogs, err := s.readBlogsFromFile()
	if err != nil {
		return err
	}

	newBlogs := make([]*Blog, 0, len(blogs))
	for _, blog := range blogs {
		if blog.ID != id {
			newBlogs = append(newBlogs, blog)
		}
	}

	return s.writeBlogs(newBlogs)
}

// writeBlogs 将博客列表写入CSV文件
func (s *CSVBlogStore) writeBlogs(blogs []*Blog) error {
	file, err := os.Create(s.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 写入表头
	if err := writer.Write([]string{"id", "title", "content", "updated_at"}); err != nil {
		return err
	}

	// 写入数据
	for _, blog := range blogs {
		record := []string{
			blog.ID,
			blog.Title,
			blog.Content,
			blog.UpdatedAt.Format(time.RFC3339),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}
