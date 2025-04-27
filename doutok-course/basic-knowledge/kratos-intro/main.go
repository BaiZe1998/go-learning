package main

import (
	"fmt"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// 这是一个简化的 Kratos 应用示例
// 在真实环境中，需要安装 Kratos 并使用 kratos 命令行工具生成完整项目

func main() {
	// 初始化 logger
	logger := log.NewStdLogger(os.Stdout)
	log := log.NewHelper(logger)

	// 创建 HTTP 服务器
	httpSrv := http.NewServer(
		http.Address(":8000"),
		http.Middleware(
			recovery.Recovery(), // 添加异常恢复中间件
		),
	)
	// 使用路由组
	r := httpSrv.Route("/")
	r.GET("", func(ctx http.Context) error {
		return ctx.String(200, "Hello Kratos!")
	})

	r.GET("/hello", func(ctx http.Context) error {
		return ctx.String(200, "Hello Kratos API!")
	})

	// 创建 Kratos 应用实例
	app := kratos.New(
		kratos.Name("kratos-demo"),
		kratos.Server(
			httpSrv,
		),
		kratos.Logger(logger),
	)

	// 启动应用
	if err := app.Run(); err != nil {
		log.Errorf("启动应用失败: %v", err)
		return
	}

	fmt.Println("服务启动成功，监听端口: 8000")
}

// 说明：这是示例代码，需要先安装依赖才能运行
// 安装 Kratos:
// go get github.com/go-kratos/kratos/v2
//
// 完整项目建议使用 Kratos 工具创建:
// go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
// kratos new server
