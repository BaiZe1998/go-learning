package data

import (
	"blog-service/internal/conf"
	"blog-service/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewBlogRepo)

// Data .
type Data struct {
	blogStore model.BlogStore
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	blogStore, err := model.NewCSVBlogStore("blogs.csv")
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	return &Data{
		blogStore: blogStore,
	}, cleanup, nil
}
