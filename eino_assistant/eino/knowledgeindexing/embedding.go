package knowledgeindexing

import (
	"context"
	"os"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/cloudwego/eino/components/embedding"
)

func newEmbedding(ctx context.Context) (eb embedding.Embedder, err error) {
	// TODO Modify component configuration here.
	config := &ark.EmbeddingConfig{
		BaseURL: "https://ark.cn-beijing.volces.com/api/v3",
		APIKey:  os.Getenv("ARK_API_KEY"),
		Model:   os.Getenv("ARK_EMBEDDING_MODEL"),
	}
	eb, err = ark.NewEmbedder(ctx, config)
	if err != nil {
		return nil, err
	}
	return eb, nil
}
