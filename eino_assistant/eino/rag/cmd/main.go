package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"eino_assistant/eino/rag"
	"eino_assistant/pkg/env"
)

func main() {
	// 定义命令行参数
	useRedis := flag.Bool("redis", true, "是否使用Redis进行检索增强")
	topK := flag.Int("topk", 3, "检索的文档数量")

	flag.Parse()

	// 检查环境变量
	env.MustHasEnvs("ARK_API_KEY")

	// 构建RAG系统
	ctx := context.Background()
	ragSystem, err := rag.BuildRAG(ctx, *useRedis, *topK)
	if err != nil {
		fmt.Fprintf(os.Stderr, "构建RAG系统失败: %v\n", err)
		os.Exit(1)
	}

	// 显示启动信息
	if *useRedis {
		fmt.Println("启动RAG系统 (使用Redis检索)")
	} else {
		fmt.Println("启动RAG系统 (不使用检索)")
	}
	fmt.Println("输入问题或输入'exit'退出")

	// 创建输入扫描器
	scanner := bufio.NewScanner(os.Stdin)

	// 主循环
	for {
		fmt.Print("\n问题> ")

		// 读取用户输入
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		// 检查退出命令
		if strings.ToLower(input) == "exit" {
			break
		}

		// 处理问题
		answer, err := ragSystem.Answer(ctx, input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "处理问题时出错: %v\n", err)
			continue
		}

		// 显示回答
		fmt.Println("\n回答:")
		fmt.Println(answer)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "读取输入时出错: %v\n", err)
	}

	fmt.Println("再见!")
}
