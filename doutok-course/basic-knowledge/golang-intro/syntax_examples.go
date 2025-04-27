package main

import (
	"fmt"
	"sync"
	"time"
)

// RunAllExamples 运行所有语法示例
func RunAllExamples() {
	fmt.Println("Go语言常用语法示例")
	fmt.Println("==================")

	// ===== 切片(Slice)示例 =====
	fmt.Println("\n1. 切片(Slice)示例:")
	sliceExamples()

	// ===== 映射(Map)示例 =====
	fmt.Println("\n2. 映射(Map)示例:")
	mapExamples()

	// ===== 通道(Channel)示例 =====
	fmt.Println("\n3. 通道(Channel)示例:")
	channelExamples()

	// ===== 遍历方法示例 =====
	fmt.Println("\n4. 遍历方法示例:")
	iterationExamples()

	// ===== 错误处理示例 =====
	fmt.Println("\n5. 错误处理示例:")
	errorHandlingExamples()

	// ===== 结构体和方法示例 =====
	fmt.Println("\n6. 结构体和方法示例:")
	structExamples()
}

// RunSliceExamples 运行切片示例
func RunSliceExamples() {
	fmt.Println("切片(Slice)示例:")
	sliceExamples()
}

// RunMapExamples 运行映射示例
func RunMapExamples() {
	fmt.Println("映射(Map)示例:")
	mapExamples()
}

// RunChannelExamples 运行通道示例
func RunChannelExamples() {
	fmt.Println("通道(Channel)示例:")
	channelExamples()
}

// RunIterationExamples 运行遍历方法示例
func RunIterationExamples() {
	fmt.Println("遍历方法示例:")
	iterationExamples()
}

// RunErrorHandlingExamples 运行错误处理示例
func RunErrorHandlingExamples() {
	fmt.Println("错误处理示例:")
	errorHandlingExamples()
}

// RunStructExamples 运行结构体和方法示例
func RunStructExamples() {
	fmt.Println("结构体和方法示例:")
	structExamples()
}

// =================== 切片示例 ===================
func sliceExamples() {
	// 1. 创建切片的不同方式
	fmt.Println("1.1 创建切片:")

	// 使用字面量创建切片
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("字面量创建: %v (长度:%d, 容量:%d)\n", s1, len(s1), cap(s1))

	// 使用make创建切片
	s2 := make([]int, 5) // 长度为5，元素初始值为0
	fmt.Printf("make创建(指定长度): %v (长度:%d, 容量:%d)\n", s2, len(s2), cap(s2))

	s3 := make([]int, 3, 10) // 长度为3，容量为10
	fmt.Printf("make创建(指定长度和容量): %v (长度:%d, 容量:%d)\n", s3, len(s3), cap(s3))

	// 从数组创建切片
	arr := [5]int{10, 20, 30, 40, 50}
	s4 := arr[1:4] // 包含索引1到3的元素
	fmt.Printf("从数组创建: %v (长度:%d, 容量:%d)\n", s4, len(s4), cap(s4))

	// 2. 切片操作
	fmt.Println("\n1.2 切片操作:")

	// 添加元素
	s5 := []int{1, 2, 3}
	s5 = append(s5, 4)       // 添加单个元素
	s5 = append(s5, 5, 6, 7) // 添加多个元素
	fmt.Printf("添加元素: %v\n", s5)

	// 合并切片
	s6 := []int{8, 9, 10}
	s5 = append(s5, s6...) // 使用...展开切片
	fmt.Printf("合并切片: %v\n", s5)

	// 复制切片
	s7 := make([]int, len(s5))
	copied := copy(s7, s5) // 返回实际复制的元素个数
	fmt.Printf("复制切片: %v (复制了%d个元素)\n", s7, copied)

	// 切片的切片
	s8 := s5[2:5] // 索引2到4的元素
	fmt.Printf("切片的切片: %v\n", s8)

	// 3. 切片的内部工作原理示例
	fmt.Println("\n1.3 切片内部原理:")

	original := []int{1, 2, 3, 4, 5}
	fmt.Printf("原始切片: %v\n", original)

	// 创建一个共享底层数组的切片
	slice1 := original[1:3]
	fmt.Printf("派生切片: %v\n", slice1)

	// 修改派生切片，会影响原始切片
	slice1[0] = 99
	fmt.Printf("修改派生切片后，原始切片: %v\n", original)

	// 当append导致容量不足时会创建新的底层数组
	slice2 := original[1:2]
	fmt.Printf("派生切片2: %v (容量:%d)\n", slice2, cap(slice2))

	// 添加足够多的元素使其扩容
	slice2 = append(slice2, 100, 200, 300, 400)
	fmt.Printf("扩容后的派生切片2: %v (容量:%d)\n", slice2, cap(slice2))

	// 此时修改slice2不会影响original
	slice2[0] = 888
	fmt.Printf("修改扩容后的切片不影响原始切片: %v\n", original)
}

// =================== 映射示例 ===================
func mapExamples() {
	// 1. 创建映射的不同方式
	fmt.Println("2.1 创建映射:")

	// 使用字面量创建映射
	m1 := map[string]int{
		"apple":  5,
		"banana": 8,
		"orange": 7,
	}
	fmt.Printf("字面量创建: %v\n", m1)

	// 使用make创建映射
	m2 := make(map[string]int)
	fmt.Printf("make创建(空): %v\n", m2)

	m3 := make(map[string]int, 10) // 指定初始容量
	fmt.Printf("make创建(指定容量): %v (初始容量:10)\n", m3)

	// 2. 映射操作
	fmt.Println("\n2.2 映射操作:")

	// 添加或更新元素
	m2["apple"] = 5
	m2["banana"] = 8
	fmt.Printf("添加元素: %v\n", m2)

	// 获取元素
	value := m2["apple"]
	fmt.Printf("获取元素 'apple': %d\n", value)

	// 检查键是否存在
	value, exists := m2["grape"]
	if exists {
		fmt.Printf("'grape'存在，值为: %d\n", value)
	} else {
		fmt.Printf("'grape'不存在\n")
	}

	// 删除元素
	delete(m2, "banana")
	fmt.Printf("删除'banana'后: %v\n", m2)

	// 3. 映射的特性
	fmt.Println("\n2.3 映射的特性:")

	// 映射是引用类型
	m4 := m1
	m4["apple"] = 100
	fmt.Printf("原始映射m1: %v\n", m1) // m1也被修改
	fmt.Printf("引用映射m4: %v\n", m4)

	// 映射的大小
	fmt.Printf("映射的大小: %d\n", len(m1))
}

// =================== 通道示例 ===================
func channelExamples() {
	// 1. 创建通道的不同方式
	fmt.Println("3.1 创建通道:")

	// 创建无缓冲通道
	ch1 := make(chan int)
	fmt.Printf("无缓冲通道: %T\n", ch1)

	// 创建有缓冲通道
	ch2 := make(chan string, 3)
	fmt.Printf("有缓冲通道(容量:%d): %T\n", ch2, cap(ch2))

	// 2. 通道的基本操作
	fmt.Println("\n3.2 通道操作:")

	// 使用有缓冲通道示例
	ch := make(chan string, 2)

	// 发送数据到通道
	ch <- "hello"
	ch <- "world"
	fmt.Println("发送了两条消息到通道")

	// 从通道接收数据
	msg1 := <-ch
	msg2 := <-ch
	fmt.Printf("从通道接收: %s, %s\n", msg1, msg2)

	// 3. 通道的常见模式
	fmt.Println("\n3.3 通道常见模式:")

	// 示例1: 使用通道进行同步
	done := make(chan bool)
	go func() {
		fmt.Println("goroutine 执行中...")
		time.Sleep(time.Millisecond * 500)
		done <- true
	}()
	<-done // 等待goroutine完成
	fmt.Println("goroutine 已完成")

	// 示例2: 通道方向限制
	channelDirection()

	// 示例3: select语句
	selectExample()
}

// 演示通道方向限制
func channelDirection() {
	fmt.Println("\n通道方向限制示例:")

	ch := make(chan int)

	// 启动生产者和消费者
	go producer(ch)
	go consumer(ch)

	// 等待一段时间让goroutines运行
	time.Sleep(time.Second)
}

// 只能向通道发送数据
func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("发送: %d\n", i)
		time.Sleep(time.Millisecond * 100)
	}
	close(ch) // 关闭通道
}

// 只能从通道接收数据
func consumer(ch <-chan int) {
	// 使用for range从通道接收数据，直到通道关闭
	for num := range ch {
		fmt.Printf("接收: %d\n", num)
	}
}

// select示例
func selectExample() {
	fmt.Println("\nselect示例:")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Millisecond * 100)
		ch1 <- "来自通道1的消息"
	}()

	go func() {
		time.Sleep(time.Millisecond * 200)
		ch2 <- "来自通道2的消息"
	}()

	// 使用select处理多个通道
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		case <-time.After(time.Millisecond * 300):
			fmt.Println("超时")
		}
	}
}

// =================== 遍历方法示例 ===================
func iterationExamples() {
	// 1. 数组遍历
	fmt.Println("4.1 数组遍历:")

	numbers := [5]int{10, 20, 30, 40, 50}

	// 使用索引遍历
	fmt.Println("索引遍历:")
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("numbers[%d] = %d\n", i, numbers[i])
	}

	// 使用range遍历
	fmt.Println("\nrange遍历:")
	for index, value := range numbers {
		fmt.Printf("索引:%d 值:%d\n", index, value)
	}

	// 只关心值的遍历
	fmt.Println("\n只关心值:")
	for _, value := range numbers {
		fmt.Printf("值:%d\n", value)
	}

	// 只关心索引的遍历
	fmt.Println("\n只关心索引:")
	for index := range numbers {
		fmt.Printf("索引:%d\n", index)
	}

	// 2. 切片遍历
	fmt.Println("\n4.2 切片遍历:")

	fruits := []string{"apple", "banana", "orange", "grape"}

	// 使用range遍历切片
	fmt.Println("range遍历切片:")
	for index, fruit := range fruits {
		fmt.Printf("fruits[%d] = %s\n", index, fruit)
	}

	// 3. 映射遍历
	fmt.Println("\n4.3 映射遍历:")

	ages := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 22,
		"David":   28,
	}

	// 遍历键和值
	fmt.Println("遍历键和值:")
	for name, age := range ages {
		fmt.Printf("%s: %d岁\n", name, age)
	}

	// 只遍历键
	fmt.Println("\n只遍历键:")
	for name := range ages {
		fmt.Printf("姓名: %s\n", name)
	}

	// 4. 字符串遍历
	fmt.Println("\n4.4 字符串遍历:")

	text := "你好，Go!"

	// 按字节遍历
	fmt.Println("按字节遍历(不适合中文):")
	for i := 0; i < len(text); i++ {
		fmt.Printf("%d: %c\n", i, text[i])
	}

	// 按Unicode字符(rune)遍历
	fmt.Println("\n按Unicode字符遍历(支持中文):")
	for index, runeValue := range text {
		fmt.Printf("%d: %c\n", index, runeValue)
	}

	// 5. 通道遍历
	fmt.Println("\n4.5 通道遍历:")

	ch := make(chan int, 5)

	// 发送数据到通道
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch) // 关闭通道，表示没有更多数据了

	// 使用range遍历通道
	fmt.Println("遍历通道:")
	for value := range ch {
		fmt.Printf("接收: %d\n", value)
	}
}

// =================== 错误处理示例 ===================
func errorHandlingExamples() {
	// 1. 基本错误处理
	fmt.Println("5.1 基本错误处理:")

	// 调用可能返回错误的函数
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("结果: %d\n", result)
	}

	// 除以零会返回错误
	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("结果: %d\n", result)
	}

	// 2. panic和recover
	fmt.Println("\n5.2 panic和recover:")

	// 使用匿名函数包装带有recover的代码
	func() {
		// 设置recover来捕获panic
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("已恢复的panic: %v\n", r)
			}
		}()

		fmt.Println("调用可能会panic的函数")
		dangerousFunction(true)
		fmt.Println("此行不会执行") // 如果panic，这行不会执行
	}()

	fmt.Println("继续执行") // recover后程序继续执行

	// 3. defer使用
	fmt.Println("\n5.3 defer使用:")
	deferExample()
}

// 演示可能返回错误的函数
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为零")
	}
	return a / b, nil
}

// 可能会引发panic的函数
func dangerousFunction(shouldPanic bool) {
	if shouldPanic {
		panic("故意引发的panic")
	}
	fmt.Println("安全执行完毕")
}

// defer示例
func deferExample() {
	fmt.Println("函数开始")

	// defer语句按LIFO(后进先出)顺序执行
	defer fmt.Println("defer 1 - 最后执行")
	defer fmt.Println("defer 2 - 倒数第二个执行")
	defer fmt.Println("defer 3 - 第一个执行")

	fmt.Println("函数中间部分")

	// defer常用于资源清理
	// 示例：模拟文件操作
	fmt.Println("\n模拟文件操作:")
	openFile("test.txt")

	fmt.Println("\n函数结束")
}

// 模拟文件打开和关闭
func openFile(filename string) {
	fmt.Printf("打开文件: %s\n", filename)

	// 确保文件关闭
	defer func() {
		fmt.Printf("关闭文件: %s\n", filename)
	}()

	// 处理文件...
	fmt.Println("处理文件内容")

	// 即使发生错误，defer也会执行
	if filename == "test.txt" {
		fmt.Println("文件处理完成")
	} else {
		fmt.Println("文件处理错误")
	}
}

// =================== 结构体和方法示例 ===================
func structExamples() {
	// 1. 定义和创建结构体
	fmt.Println("6.1 结构体定义和创建:")

	// 使用字面量创建结构体
	p1 := Person{
		Name: "Alice",
		Age:  30,
	}
	fmt.Printf("字面量创建: %+v\n", p1)

	// 使用new创建结构体指针
	p2 := new(Person)
	p2.Name = "Bob" // 自动解引用
	p2.Age = 25
	fmt.Printf("使用new创建: %+v\n", p2)

	// 创建结构体并省略字段名
	p3 := Person{"Charlie", 35}
	fmt.Printf("省略字段名创建: %+v\n", p3)

	// 2. 结构体嵌套
	fmt.Println("\n6.2 结构体嵌套:")

	e := Employee{
		Person: Person{
			Name: "David",
			Age:  28,
		},
		Title:  "软件工程师",
		Salary: 12000,
	}
	fmt.Printf("嵌套结构体: %+v\n", e)
	fmt.Printf("访问嵌套字段: %s, %d岁, %s\n", e.Name, e.Age, e.Title)

	// 3. 结构体方法
	fmt.Println("\n6.3 结构体方法:")

	// 调用值接收者方法
	fmt.Printf("%s的个人介绍: %s\n", p1.Name, p1.Description())

	// 修改年龄
	p1.SetAge(31)
	fmt.Printf("修改后的年龄: %d\n", p1.Age)

	// 4. 接口示例
	fmt.Println("\n6.4 接口示例:")

	// 创建不同类型的实体
	alice := Person{Name: "Alice", Age: 30}
	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 2.5}

	// 使用共同接口
	printInfo(alice)
	printInfo(&rect)
	printInfo(circle)

	// 计算形状的面积
	printArea(rect)
	printArea(circle)

	// 5. 并发安全的结构体
	fmt.Println("\n6.5 并发安全的结构体:")

	// 创建并发安全的计数器
	counter := SafeCounter{value: 0}
	var wg sync.WaitGroup

	// 启动多个goroutine递增计数器
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
			fmt.Printf("计数器值: %d\n", counter.Value())
		}()
	}

	wg.Wait()
	fmt.Printf("最终计数器值: %d\n", counter.Value())
}

// 定义Person结构体
type Person struct {
	Name string
	Age  int
}

// 值接收者方法
func (p Person) Description() string {
	return fmt.Sprintf("%s, %d岁", p.Name, p.Age)
}

// 指针接收者方法
func (p *Person) SetAge(age int) {
	p.Age = age
}

// 定义Employee结构体，嵌套Person
type Employee struct {
	Person // 匿名嵌套
	Title  string
	Salary float64
}

// 定义形状接口
type Shape interface {
	Area() float64
}

// 定义矩形结构体
type Rectangle struct {
	Width, Height float64
}

// 实现Area方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 为Rectangle添加Description方法
func (r Rectangle) Description() string {
	return fmt.Sprintf("矩形 宽:%.1f 高:%.1f", r.Width, r.Height)
}

// 定义圆形结构体
type Circle struct {
	Radius float64
}

// 实现Area方法
func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

// 为Circle添加Description方法
func (c Circle) Description() string {
	return fmt.Sprintf("圆形 半径:%.1f", c.Radius)
}

// 定义Describer接口
type Describer interface {
	Description() string
}

// 打印任何实现了Describer接口的类型信息
func printInfo(d Describer) {
	fmt.Printf("描述: %s\n", d.Description())
}

// 打印形状的面积
func printArea(s Shape) {
	fmt.Printf("面积: %.2f\n", s.Area())
}

// 并发安全的计数器结构体
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

// 递增计数器
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// 获取计数器值
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}
