package rag

import (
	"context"
	"encoding/binary"
	"fmt"
	"math"
	"os"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/cloudwego/eino/components/embedding"
)

// RAG 表示一个检索增强生成系统
type RAG struct {
	retriever Retriever
	generator Generator
	topK      int
}

// NewRAG 创建一个新的RAG系统
func NewRAG(retriever Retriever, generator Generator, topK int) *RAG {
	return &RAG{
		retriever: retriever,
		generator: generator,
		topK:      topK,
	}
}

// Answer 处理问题并生成回答
func (r *RAG) Answer(ctx context.Context, query string) (string, error) {
	// 检索相关文档
	docs, err := r.retriever.Retrieve(ctx, query, r.topK)
	if err != nil {
		return "", fmt.Errorf("检索失败: %w", err)
	}

	// 打印检索到的文档
	fmt.Printf("\n===== 检索到 %d 个相关文档 =====\n", len(docs))
	for i, doc := range docs {
		score := 0.0
		if scoreVal, ok := doc.MetaData["score"]; ok {
			if s, ok := scoreVal.(float32); ok {
				score = float64(s)
			}
		}

		// 提取标题（如果有）
		title := "无标题"
		if titleVal, ok := doc.MetaData["title"]; ok {
			if t, ok := titleVal.(string); ok && t != "" {
				title = t
			}
		}

		fmt.Printf("\n文档[%d] 相似度: %.4f  标题: %s\n", i+1, score, title)
		fmt.Printf("----------------------------------------\n")
		// 打印内容摘要（最多显示300个字符）
		content := doc.Content
		if len(content) > 300 {
			content = content[:300] + "..."
		}
		fmt.Printf("%s\n", content)
	}
	fmt.Printf("\n==============================\n\n")

	// 使用检索到的文档生成回答
	answer, err := r.generator.Generate(ctx, query, docs)
	if err != nil {
		return "", fmt.Errorf("生成回答失败: %w", err)
	}

	return answer, nil
}

// BuildRAG 构建一个完整的RAG系统
func BuildRAG(ctx context.Context, useRedis bool, topK int) (*RAG, error) {
	// 创建嵌入模型
	embedder, err := newEmbedding(ctx)
	if err != nil {
		return nil, fmt.Errorf("创建嵌入模型失败: %w", err)
	}

	// 创建检索器
	var retriever Retriever
	if useRedis {
		redisAddr := os.Getenv("REDIS_ADDR")
		if redisAddr == "" {
			redisAddr = "localhost:6379"
		}
		retriever, err = NewRedisRetriever(ctx, redisAddr, embedder)
		if err != nil {
			return nil, fmt.Errorf("创建Redis检索器失败: %w", err)
		}
	} else {
		retriever = NewNoRetriever()
	}

	// 创建生成器
	apiKey := os.Getenv("ARK_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("环境变量ARK_API_KEY未设置")
	}

	modelName := os.Getenv("ARK_CHAT_MODEL")
	if modelName == "" {
		modelName = "doubao" // 默认使用doubao模型
	}

	baseURL := os.Getenv("ARK_API_BASE_URL")
	if baseURL == "" {
		baseURL = "https://ark.cn-beijing.volces.com/api/v3"
	}

	generator := NewArkGenerator(baseURL, apiKey, modelName)

	// 创建RAG系统
	rag := NewRAG(retriever, generator, topK)
	return rag, nil
}

// 重用嵌入模型创建函数
func newEmbedding(ctx context.Context) (eb embedding.Embedder, err error) {
	config := &ark.EmbeddingConfig{
		BaseURL: os.Getenv("ARK_API_BASE_URL"),
		APIKey:  os.Getenv("ARK_API_KEY"),
		Model:   os.Getenv("ARK_EMBEDDING_MODEL"),
	}

	// 使用默认值
	if config.BaseURL == "" {
		config.BaseURL = "https://ark.cn-beijing.volces.com/api/v3"
	}

	if config.Model == "" {
		// 使用默认嵌入模型
		config.Model = "embedding-v2"
	}

	eb, err = ark.NewEmbedder(ctx, config)
	if err != nil {
		return nil, err
	}
	return eb, nil
}

func vectorToBytes(vector []float64) []byte {
	// 转换为float32数组
	float32Vector := make([]float32, len(vector))
	for i, v := range vector {
		float32Vector[i] = float32(v)
	}

	// 创建二进制缓冲区
	buf := make([]byte, len(float32Vector)*4)
	for i, v := range float32Vector {
		binary.LittleEndian.PutUint32(buf[i*4:], math.Float32bits(v))
	}
	return buf
}
