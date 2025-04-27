package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 检查命令行参数
	if len(os.Args) > 1 && os.Args[1] == "syntax" {
		// 如果提供了syntax参数，则运行语法示例
		runSyntaxExamples()
		return
	}

	// 原来的Hello World程序
	// 打印 Hello World
	fmt.Println("Hello, Golang!")

	// 变量声明与使用
	var name string = "学习者"
	age := 25
	fmt.Printf("你好，%s！你今年 %d 岁。\n", name, age)

	// 简单的条件语句
	if age >= 18 {
		fmt.Println("你已经成年了")
	} else {
		fmt.Println("你还未成年")
	}

	// 创建一个切片并遍历
	languages := []string{"Go", "Python", "Java", "JavaScript"}
	fmt.Println("流行的编程语言:")
	for i, lang := range languages {
		fmt.Printf("%d: %s\n", i+1, lang)
	}

	// 调用函数
	result := add(10, 20)
	fmt.Printf("10 + 20 = %d\n", result)

	// 输出提示信息
	fmt.Println("\n提示: 运行 'go run hello.go syntax' 以查看更多语法示例")
	fmt.Println("或者运行 'go run hello.go syntax <数字>' 以查看特定类别的示例:")
	fmt.Println("1: 切片示例")
	fmt.Println("2: 映射示例")
	fmt.Println("3: 通道示例")
	fmt.Println("4: 遍历方法示例")
	fmt.Println("5: 错误处理示例")
	fmt.Println("6: 结构体和方法示例")
}

// 定义一个简单的函数
func add(a, b int) int {
	return a + b
}

// 运行语法示例的函数
func runSyntaxExamples() {
	if len(os.Args) > 2 {
		// 如果提供了第二个参数，根据参数运行特定示例
		category, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("无效的参数，请提供1-6之间的数字")
			return
		}

		switch category {
		case 1:
			RunSliceExamples()
		case 2:
			RunMapExamples()
		case 3:
			RunChannelExamples()
		case 4:
			RunIterationExamples()
		case 5:
			RunErrorHandlingExamples()
		case 6:
			RunStructExamples()
		default:
			fmt.Println("无效的类别，请提供1-6之间的数字")
		}
	} else {
		// 如果没有提供第二个参数，运行所有示例
		RunAllExamples()
	}
}
