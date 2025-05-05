package rag

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/cloudwego/eino/schema"
)

// Generator 定义生成接口
type Generator interface {
	Generate(ctx context.Context, query string, documents []*schema.Document) (string, error)
}

// ArkGenerator 实现基于ARK的生成器
type ArkGenerator struct {
	baseURL   string
	apiKey    string
	modelName string
}

// NewArkGenerator 创建ARK生成器
func NewArkGenerator(baseURL, apiKey, modelName string) *ArkGenerator {
	return &ArkGenerator{
		baseURL:   baseURL,
		apiKey:    apiKey,
		modelName: modelName,
	}
}

// 请求结构体
type chatRequest struct {
	Model    string        `json:"model"`
	Messages []chatMessage `json:"messages"`
}

type chatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// 响应结构体
type chatResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Choices []struct {
		Index        int         `json:"index"`
		Message      chatMessage `json:"message"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
}

// Generate 生成回答
func (g *ArkGenerator) Generate(ctx context.Context, query string, documents []*schema.Document) (string, error) {
	// 组合上下文信息
	context := ""
	if len(documents) > 0 {
		contextParts := make([]string, len(documents))
		for i, doc := range documents {
			// 如果元数据中有标题，添加标题信息
			titleInfo := ""
			if title, ok := doc.MetaData["title"].(string); ok && title != "" {
				titleInfo = fmt.Sprintf("标题: %s\n", title)
			}
			contextParts[i] = fmt.Sprintf("文档片段[%d]:\n%s%s\n", i+1, titleInfo, doc.Content)
		}
		context = strings.Join(contextParts, "\n---\n")
	}

	// 构建提示
	systemPrompt := "你是一个知识助手。基于提供的文档回答用户问题。如果文档中没有相关信息，请诚实地表明你不知道，不要编造答案。"
	userPrompt := query

	if context != "" {
		userPrompt = fmt.Sprintf("基于以下信息回答我的问题：\n\n%s\n\n问题：%s", context, query)
	}

	// 构建请求
	messages := []chatMessage{
		{Role: "system", Content: systemPrompt},
		{Role: "user", Content: userPrompt},
	}

	reqBody := chatRequest{
		Model:    g.modelName,
		Messages: messages,
	}

	// 序列化请求体
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("序列化请求失败: %w", err)
	}

	// 创建HTTP请求
	endpoint := fmt.Sprintf("%s/chat/completions", g.baseURL)
	req, err := http.NewRequestWithContext(ctx, "POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("创建HTTP请求失败: %w", err)
	}

	// 添加头信息
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", g.apiKey))

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("发送请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败: %w", err)
	}

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API返回错误: %s, 状态码: %d", string(body), resp.StatusCode)
	}

	// 解析响应
	var chatResp chatResponse
	if err := json.Unmarshal(body, &chatResp); err != nil {
		return "", fmt.Errorf("解析响应失败: %w", err)
	}

	// 提取回答
	if len(chatResp.Choices) > 0 {
		return chatResp.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("API没有返回有效回答")
}
