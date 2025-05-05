package knowledgeindexing

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/cloudwego/eino-ext/components/indexer/redis"
	"github.com/cloudwego/eino/components/indexer"
	"github.com/cloudwego/eino/schema"
	"github.com/google/uuid"
	redisCli "github.com/redis/go-redis/v9"

	redispkg "eino_assistant/pkg/redis"
)

func init() {
	err := redispkg.Init()
	if err != nil {
		log.Fatalf("failed to init redis index: %v", err)
	}
}

// newIndexer component initialization function of node 'RedisIndexer' in graph 'KnowledgeIndexing'
func newIndexer(ctx context.Context) (idr indexer.Indexer, err error) {
	// TODO Modify component configuration here.
	redisAddr := os.Getenv("REDIS_ADDR")
	redisClient := redisCli.NewClient(&redisCli.Options{
		Addr:     redisAddr,
		Protocol: 2,
	})

	config := &redis.IndexerConfig{
		Client:    redisClient,
		KeyPrefix: redispkg.RedisPrefix,
		BatchSize: 1,
		DocumentToHashes: func(ctx context.Context, doc *schema.Document) (*redis.Hashes, error) {
			if doc.ID == "" {
				doc.ID = uuid.New().String()
			}
			key := doc.ID

			metadataBytes, err := json.Marshal(doc.MetaData)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal metadata: %w", err)
			}

			return &redis.Hashes{
				Key: key,
				Field2Value: map[string]redis.FieldValue{
					redispkg.ContentField:  {Value: doc.Content, EmbedKey: redispkg.VectorField},
					redispkg.MetadataField: {Value: metadataBytes},
				},
			}, nil
		},
	}
	embeddingIns11, err := newEmbedding(ctx)
	if err != nil {
		return nil, err
	}
	config.Embedding = embeddingIns11
	idr, err = redis.NewIndexer(ctx, config)
	if err != nil {
		return nil, err
	}
	return idr, nil
}
