package rag

import (
	"context"
	"encoding/json"
	"fmt"
	"math"

	"github.com/cloudwego/eino/components/embedding"
	"github.com/cloudwego/eino/schema"
	"github.com/redis/go-redis/v9"

	redispkg "eino_assistant/pkg/redis"
)

// Retriever 定义检索接口
type Retriever interface {
	Retrieve(ctx context.Context, query string, topK int) ([]*schema.Document, error)
}

// RedisRetriever 实现基于Redis的检索
type RedisRetriever struct {
	client    *redis.Client
	embedder  embedding.Embedder
	indexName string // 在执行FT.SEARCH命令时作为第一个参数使用，指定在哪个索引上执行向量搜索
	prefix    string // 在执行FT.SEARCH命令时作为第二个参数使用，指定在哪个键空间中执行向量搜索
}

// NewRedisRetriever 创建Redis检索器
func NewRedisRetriever(ctx context.Context, redisAddr string, embedder embedding.Embedder) (*RedisRetriever, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Protocol: 2,
	})

	// 测试Redis连接
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("无法连接到Redis: %w", err)
	}

	return &RedisRetriever{
		client:    client,
		embedder:  embedder,
		indexName: fmt.Sprintf("%s%s", redispkg.RedisPrefix, redispkg.IndexName),
		prefix:    redispkg.RedisPrefix,
	}, nil
}

// Retrieve 检索与查询最相关的文档
func (r *RedisRetriever) Retrieve(ctx context.Context, query string, topK int) ([]*schema.Document, error) {
	// 生成查询向量
	queryVectors, err := r.embedder.EmbedStrings(ctx, []string{query})
	if err != nil {
		return nil, fmt.Errorf("生成查询向量失败: %w", err)
	}

	if len(queryVectors) == 0 || len(queryVectors[0]) == 0 {
		return nil, fmt.Errorf("嵌入模型返回空向量")
	}

	queryVector := queryVectors[0]

	// 构建向量搜索查询
	searchQuery := fmt.Sprintf("(*)=>[KNN %d @%s $query_vector AS %s]",
		topK,
		redispkg.VectorField,
		redispkg.DistanceField)

	// 执行向量搜索
	res, err := r.client.Do(ctx,
		"FT.SEARCH", r.indexName, // 执行搜索的索引名称
		searchQuery,   // 向量搜索查询语句
		"PARAMS", "2", // 参数声明，后面有2个参数
		"query_vector", vectorToBytes(queryVector), // 查询向量的二进制表示
		"DIALECT", "2", // 查询方言版本
		"SORTBY", redispkg.DistanceField, // 结果排序字段
		"RETURN", "3", redispkg.ContentField, redispkg.MetadataField, redispkg.DistanceField, // 返回字段
	).Result()

	if err != nil {
		return nil, fmt.Errorf("执行向量搜索失败: %w", err)
	}

	// 将Redis结果转换为Document对象
	return r.parseSearchResults(res)
}

// NoRetriever 实现一个不进行检索的空实现
type NoRetriever struct{}

// NewNoRetriever 创建一个不使用检索的实现
func NewNoRetriever() *NoRetriever {
	return &NoRetriever{}
}

// Retrieve 空实现，返回空文档列表
func (r *NoRetriever) Retrieve(ctx context.Context, query string, topK int) ([]*schema.Document, error) {
	return []*schema.Document{}, nil
}

// 辅助函数：解析Redis搜索结果
func (r *RedisRetriever) parseSearchResults(redisResult interface{}) ([]*schema.Document, error) {
	results, ok := redisResult.([]interface{})
	if !ok || len(results) == 0 {
		return nil, fmt.Errorf("无效的Redis搜索结果")
	}

	// 第一个元素是匹配的文档数量
	count, ok := results[0].(int64)
	if !ok || count == 0 {
		return []*schema.Document{}, nil
	}

	docs := make([]*schema.Document, 0, count)

	// 结果格式为 [总数, 键1, 值数组1, 键2, 值数组2, ...]
	for i := 1; i < len(results); i += 2 {
		if i+1 >= len(results) {
			break
		}

		docID, ok := results[i].(string)
		if !ok {
			continue
		}

		values, ok := results[i+1].([]interface{})
		if !ok || len(values) < 6 {
			continue
		}

		// 提取内容、元数据和相似度
		var content string
		var metadataStr string
		var distance float64

		for j := 0; j < len(values); j += 2 {
			fieldName, ok := values[j].(string)
			if !ok {
				continue
			}

			switch fieldName {
			case redispkg.ContentField:
				if contentVal, ok := values[j+1].(string); ok {
					content = contentVal
				}
			case redispkg.MetadataField:
				if metadataVal, ok := values[j+1].(string); ok {
					metadataStr = metadataVal
				}
			case redispkg.DistanceField:
				if distanceVal, ok := values[j+1].(string); ok {
					if f, err := fmt.Sscanf(distanceVal, "%f", &distance); err != nil || f != 1 {
						// 默认值
						distance = 1.0
					}
				}
			}
		}

		// 解析元数据
		metadata := make(map[string]interface{})
		if metadataStr != "" {
			if err := json.Unmarshal([]byte(metadataStr), &metadata); err != nil {
				// 忽略解析错误，使用空元数据
				metadata = make(map[string]interface{})
			}
		}

		// 计算相似度得分 (余弦距离转换为相似度)
		score := 1 - distance
		// 防止浮点数精度问题
		score = math.Max(0, math.Min(1, score))

		// 创建文档对象
		doc := &schema.Document{
			ID:      docID,
			Content: content,
			MetaData: map[string]interface{}{
				"score": float32(score),
			},
		}

		// 如果需要保留原有元数据
		for k, v := range metadata {
			doc.MetaData[k] = v
		}

		docs = append(docs, doc)
	}

	return docs, nil
}
