[TOC]

讲述一个知识点的时候并非完全讲透，有时是浅尝辄止，因为后续涉及到后续章节，会在后面进行更深入的讲解

## 一、综述

### 1.1 Hello Word

介绍包管理，编译依赖，运行代码的流程；无需分号结尾以及严格的自动格式化

### 1.2 命令行参数

参数初始化，获取命令行参数的方式，给出了一个低效的循环获取命令行参数的代码，在此基础上进行优化

关于字符串常量的累加（是否是不断创建新值，变量创建后如何存储，结合Java堆｜栈）

### 1.3 查找重复行

strings.join底层发生了什么

map乱序的原因

os.stdin Scan的终止条件

输出错误内容到标准错误

何时可以跳过error检查

### 1.4 GIF 动画

可以生成gif格式的图片

### 1.5 获取一个URL

resp.Body.Close()可以avoid leaking resources，具体发生了什么

io.Copy(dst, src)与ioutil.ReadAll的工作模式区别

### 1.6 并发获取多个URL

当多个goroutine同时对一个channel进行输入输出的时候，会发生阻塞

### 1.7 实现一个 Web 服务器

fmt.Fprintf(dir, src)可以将内容输出到指定输出（web的response、标准错误），因为dir实现了接口（io.Writer）

启动服务程序的时候mac&linux为什么末尾要加&

服务端handler路由匹配到前缀则可以触发，并且开启不同goroutine处理request（那么上限是多少，高访问量会发生什么）

### 1.8 杂项

switch在满足case之后不会继续下沉，且default可以放置在任何位置

switch也可以以tarless的模式书写

goto语法不常用，但是go也提供了

func也可以作为一种类型

结构、指针、方法、接口、包、注释

## 二、程序结构

### 2.1 名字

包名通常小写字母命名

通常来说，对于作用域较短的变量名，Go推荐短命名，如i而不是index，而对于全局变量则倾向于更长，更凸显意义的命名

驼峰而非下划线命名

### 2.2 声明

注意全局变量的作用域最小也是整个包的所有文件，大写则可以跨包

### 2.3 变量

引用类型：slice、pointer、map、channel、function

可以同时初始化多种类型的变量，并且Go没有未初始化的变量

```go
var x float64 = 100 // 此时不使用短变量命名
```

:= 是声明，而 = 是赋值

**巧妙：如果:=左侧部分变量已经声明过（作用域相同），则只会对其进行赋值，而只声明+赋值未声明过的变量，且左侧必须至少有一个未声明才能用:=，且declarations outer block are ignored**

```go
x := 1
p := &x
*p = 2 // 则 x == 1

var x, y int
&x == &x, &x == &y, &x == nil // true false false
```

Go的flag包可以实现获取命令行参数的功能：-help的来源

```go
p := new(int) // p是int类型的指针（或者某个类型的引用），此时*p == 0
*p = 2 // new 并不常用
```

垃圾回收：一个变量如果不可达（unreachable），则会被回收

关于变量的生命周期：全局变量在程序运行周期内一直存在，而局部变量则会在unreachable时会被回收，其生命周期从变量的声明开始，到unreachable时结束

**栈内存：栈内存由编译器自动分配和释放，开发者无法控制。栈内存一般存储函数中的局部变量、参数等，函数创建的时候，这些内存会被自动创建；函数返回的时候，这些内存会被自动释放，栈可用于内存分配，栈的分配和回收速度非常快**

**堆内存：只要有对变量的引用，变量就会存在，而它存储的位置与语言的语义无关。如果可能，变量会被分配到其函数的栈，但如果编译器无法证明函数返回之后变量是否仍然被引用，就必须在堆上分配该变量，采用垃圾回收机制进行管理，从而避免指针悬空。此外，局部变量如果非常大，也会存在堆上。**

**在编译器中，如果变量具有地址，就作为堆分配的候选，但如果逃逸分析可以确定其生存周期不会超过函数返回，就会分配在栈上。**

**总之，分配在堆还是栈完全由编译器确定。而原本看起来应该分配在栈上的变量，如果其生命周期获得了延长，被分配在了堆上，就说它发生了逃逸。编译器会自动地去判断变量的生命周期是否获得了延长，整个判断的过程就叫逃逸分析。**

```go
/* 
此时x虽然是局部变量，但是被分配在堆内存，在f()调用结束后依旧可以通过global获取x的内容，我们称x从f当中escape了

逃逸并非是一件不好的事情，但是需要注意，对于那些需要被回收的短生命周期的变量，不要在编程当中被长生命周期的变量（全局变量）引用，否则会很大程度上影响Go的垃圾回收能力，造成内存分配压力
*/
var global *int
func f() {
	var x int
	x = 1
	global = &x
}
// 此时*y没有从g()当中escape，因此是分配在栈内存当中，调用结束变成unreachable，需要被回收
fun g() {
  y := new(int)
  *y = 1
}
```

### 2.4 赋值

```go
x, y = y, x
a[i], a[j] = a[j], a[i]
// 计算斐波那契数列，=赋值右侧的表达式会按照旧值先计算后赋值给左侧变量
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
}
```

### 2.5 类型声明

```go
type IntA int
type IntB int

var (
  x IntA = 1 // 此时x和y是不同类型，因此无法比较与一起运算
  y IntB = 2
)
```

T(x)将x转成T类型，转换操作可以执行的前提是x和T在底层是相同的类型，或者二者是未命名的指针类型，底层指向相同的类型

这样的转换虽然转化了值的类型，但是并没有改变其代表的值

当然，数值类型的变量之间也允许这种转换（损失精度），或者将string转换成[]byte的切片等，当然这些转化方式将改变值的内容

### 2.6 包和文件

包中.go文件的初始化流程：

1. 如果package p内部import了q，则会先初始化package q
2. main package最后初始化，可以确保main func在执行时所有的package已经完成初始化

### 2.7 作用域

变量的scope（作用域）是处于compile-time（编译时）的特征

变量的lifetime（生命周期）是处于run-time（运行时）的特征

```go
if x := f(); x == 0 {
  fmt.Println(x, y)
} else if y := g(x); x == y {
  fmt.Println(x, y)
} else {
  fmt.Println(x, y)
}
fmt.Println(x, y) // compile error: x and y are not visible here
```

变量作用域的测试如下：

```go
func test() (int, error) {
	return 1, nil
}

func main() {
	x := 0

	for i := 1; i <= 5; i++ {
		x := i
		fmt.Println(x, &x)
	}
	fmt.Println(x, &x) // 此时x依旧是0，说明for内部的x是重新声明的
  x, err := test() // 此时x和err通过:=声明+赋值，但是结合2.3节的内容，此时x已经声明，所以只对其进行赋值为1，但是地址不变
	fmt.Println(x, &x, err) // 此处打印的x == 1时的地址与赋值前x == 0地址相同
}
// 结果
1 0x1400012a010
2 0x1400012a030
3 0x1400012a038
4 0x1400012a040
5 0x1400012a048
0 0x1400012a008
1 0x1400012a008 <nil>
```

## 三、基本数据类型

### 3.1 整数

负数的%运算

&^（位运算符：and not），x &^ y = z，y中1的位，则z中对应为0，否则z中对应为x中的位

```go
00100010 &^ 00000110 = 00100000
```

无符号整数通常不会用于只为了存放非负整数变量，只有当涉及到位运算、特殊的算数运算、hash等需要利用无符号特性的场景下才会去选择使用

比如数组下标i用int存放，而不是uint，因为i--使得i == -1时作为判断遍历结束的标志，如果是uint，则0减1则等于2^64-1，而不是-1，无法结束遍历

注意：int的范围随着当前机器决定是32位还是64位

```go
var x int32 = 1
var y int16 = 2
var z int = x + y // complie error
var z int = int(x) + int(y) // ok
// 大多数数值型的类型转换不会改变值的内容，只会改变其类型（编译器解释这个变量的方式），但是当整数和浮点数以及大范围类型与小范围类型转换时，可能会丢失精度，或者出现意外的结果
```

### 3.2 浮点数

```go
math.MaxFloat32
math.MinFloat32
const x = 6.2222334e30 // 科学计数法
// math包中有很多的使用浮点数的函数，并且fmt包有很多适用于浮点数的格式化输出，包括保留小数点的具体精度等
```

float32精度大概6位

float64精度大概15位（更常用，因为单精度计算损失太快）

```go
// 直接用浮点数为返回值结果，再二次用于其他的比较判断返回结果是否有效，有时会有误差导致错误，推荐额外增加一个bool参数
func compute() (value float64, ok bool) {
  if failed {
    return 0, false
  }
  return result, true
}
```

### 3.3 复数

```go
var x complex128 = complex(1, 2) // 1+2i
var y complex128 = complex(3, 4) // 3+4i
fmt.Println(x*y) // -5+10i
fmt.Println(real(x*y)) // -5
fmt.Println(imag(x*y)) // 10
// 两个复数相等当且仅当实部和虚部相当
fmt.Println(cmplx.Sqrt(-1)) // 0+1i
```

### 3.4 布尔量

bool是if或者for的判断条件

```go
s != "" && s[0] == 'x' // 当逻辑运算符左侧表达式可以决定操作结果则将放弃执行右侧表达式
// &&的优先级高于||
```

### 3.5 字符串

string在GO语言中是不可变的量

len获取的是字符串的字节数目，而不是码点（UTF-8 Unicode code point）

字符串第i个字节，并不一定是字符串的第i个字符，因为UTF-8编码对于非ASCII的code point需要2个或更多字节

```golang
str := "hello, world"
fmt.Println(s[:5]) // hello
fmt.Println(s[7:]) // world
fmt.Println(s[:]) // hello world

s := "left"
t := s
s += " right" // 此时s指向新创建的string ”left right“，而t指向之前的s表示的“left”

a := "x"
b := "x"
a == b // true，且string可以按照字典序比较大小
```

string是不可变的，意味着同一个string的拷贝可以共享底层的内存，使得拷贝变得很轻量，比如s和s[:7]可以安全的共享相同的数据，因此substring操作也很轻量。没有新的内存被分配。

![image-20220823211950255](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20220823211950255.png)

反引号中的字符串表示其原生的意思，内容可以多行定义，不支持转义字符

```go
func main() {
	a := `hello 
	world
	lalala`
	fmt.Println(a)
}
```

#### Unicode

UTF-8使用码点描述字符（Unicode code point），在Go中对应术语：rune（GO中使用int32存储）

可以使用一个int32的序列，来代表rune序列，固定长度带来了额外的开销（因为大多数常用字符可以使用16bits描述）

#### UTF-8

可变长编码，使用1-4 bytes来代表一个rune，1byte存储 ASCII ，2 or 3 bytes 存储大多数常用字符 rune，并且采用高位固定的方式来区分范围（前缀编码，无二义性，编码更紧凑）

![image-20220823214937576](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20220823214937576.png)

```go
s := "Hello, 世界"
fmt.Println(len(s)) // 13
fmt.Println(utf8.RuneCountInString(s)) // 9
```

![image-20220823220227181](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20220823220227181.png)

#### 字符串和数组切片

字符串包含了一连串的字节（byte），创建后不可变。而[]byte内容是可变的

```go
s := "abc"
b := []byte(s) // 分配新的字节数组内存
s2 : string(b) // 发生内存拷贝
```

为了避免没有必要的转换和内存分配，bytes包中提供了很多与string包中相同功能的方法，更推荐使用（共享内存）

bytes.Buffer用于字符（字符串）的累加构造字符串操作很方便，高效

```go
// 一些操作Buffer的api
var buf bytes.Buffer
fmt.Fprintf(&buf, "%d", x)
buf.WriteString("abc")
buf.WriteByte('x')
buf.WriteRune(码点值)
```

#### 字符串和数值类型的转换

```go
// 整数转string
x := 123
y := fmt.Sprintf("%d", x)
fmt.Println(y, strconv,Itoa(x)) // "123" "123"
// %b %d %u %x 用于进制转换
s := fmt.Sprintf("x=%b", x) // "x=1111011"
// string转整数
x, err := strconv.Atoi("123")
y, err := strconv.ParseInt("123", 10, 64)// base 10, up to 64 bits，第三个参数表示转换的整型的范围为int64
// fmt.Scanf()可以用于读取混合数据（整型、字符等）
```

### 3.6 常量

所有常量的底层都是由：boolean、string、number组成（在编译时确定，不可变，常量的运算结果依旧是常量）

```go
const a = 2
const b = 2*a // b 在编译时完成
```

大多数常量的声明没有指定类型，但是也可以指定，没有类型的常量Go中称为无类型常量（untyped constant），具体的类型到使用到的时候确定

#### untyped constant

```go
const a = 10
fmt.Printf("%T\n", a) // int（隐式类型）
var b float64 = 4*a // 在需要的时候，a转变成了float64
fmt.Printf("%T\n", b) // float64
```

在默认情况下，untyped constant 不是没有具体类型，而是隐式转换成了如下类型，因此上述a的类型可以打印为int

并且untyped constant拥有更高的精度，可以认为至少有 256bit 的运算精度

- untyped boolean
- untyped integer （隐式转换成 int）
- untyped rune （隐式转换成 int32）
- untyped floaing-point （隐式转换成 float64）
- untyped complex (隐匿转换成 complex128)
- untyped string

#### 常量生成器

可以参与计算且拥有增长属性

```go
type Flags uint
const (
	a = 1 << iota // 1
	b							// 2
	c  						// 4
	d							// 8
)
const (
  _ = 1 << (10 * iota)
  KiB // 2^10
  MiB // 2^20
  GiB // 2^30
  TiB // 
  PiB // 
  EiB // 
  ZiB // 2^70
  ...
)
```

## 四、复合类型

### 4.1 数组

长度不可变，如果两个数组类型是相同的则可以进行比较，且只有完全相等才会为true

```go
a := [...]int{1, 2} // 数组的长度由内容长度确定
b := [2]int{1, 2}
c := [3]int{1, 2}
```

#### 4.2 切片

切片由三部分组成：指针、长度（len）、容量（cap）

切片可以通过数组创建

```go
// 创建月份数组
months := [...]string{1:"January", 省略部分内容, 12: "December"}
```

基于月份数组创建切片，且不同切片底层可能共享一片数组空间

![image-20220829102017677](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20220829102017677.png)

```go
fmt.Println(summer[:20]) // panic: out of range
endlessSummer := summer[:5] // 如果未超过summer的cap，则会扩展slice的len
fmt.Println(endlessSummer) // "[June July August September October]"
```

[]byte切片可以通过对字符串使用类似上述操作的方式获取

切片之间不可以使用==进行比较，只有当其判断是否为nil才可以使用

切片的zero value是nil，nil切片底层没有分配数组，nil切片的len和cap都为0，但是非nil切片的len和cap也可以为0（Go中len == 0的切片处理方式基本相同）

```go
var s []int // len(s) == 0, s == nil
s = nil // len(s) == 0, s == nil
s = []int(nil) // len(s) == 0, s == nil
s = []int{} // len(s) == 0, s != nil
```

#### The append Function

使用append为slice追加内容，如果cap == len，则会触发slice扩容，下面是一个帮助理解的例子（使用了2倍扩容，并非是Go内置的append处理流程，那将会更加精细，api也更加丰富）：

![image-20220829112405398](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20220829112405398.png)

### 4.3 映射

map(hash table) — **无序集合**，key必须是可以比较的（除了浮点数，这不是一个好的选择）

```go
x := make(map[string]int)
y := map[string]int{
  "alice": 12,
  "tom": 34
}
z := map[string]int{}
// 内置函数
delete(y, "alice")
```

对map的元素进行取地址并不是一个好的注意，因为map的扩容过程中可能伴随着rehash，导致地址发生变化（那么map的扩容规则？）

```go
ages["carol"] = 21 // panic if assignment to entry in nil map
// 判断key-value是否存在的方式
age, ok := ages["alice"]
if age, ok := ages["bob"]; !ok {
  ...
}
```

### 4.4 结构体

```go
type Point struct {
  x, y int
}
type Circle struct {
	center Point
  radius int
}
type Wheel struct {
  circle Circle
  spokes int
}
w := Wheel{Circle{Point{8, 8}, 5}, 20}
w := Wheel{
		circle: Circle{
			center: Point{x: 8, y: 8},
			radius: 5,
		},
		spokes: 20,
	}
```

### 4.5 JSON

```go
// 将结构体转成存放json编码的byte切片
type Movie struct {
	Title string
	Year  int  `json:"released"` // 重定义json属性名称
	Color bool `json:"color,omitempty"` // 如果是空值则转成json时忽略
}
data, err := json.Marshal(movie)
data2, err := json.MarshalIndent(movie, "", " ")
// 输出结果
{"Title":"s","released":1,"color":true}
{
 "Title": "s",
 "released": 1,
 "color": true
}
// json解码
content := Movie{}
json.Unmarshal(data, &content)
fmt.Println(content)
```

### 4.6 文本和HTML模板

略

## 五、方法

### 5.1 方法声明

```go
// 可以提前声明返回值z
func add(x, y int) (z int) {
	z = x-y
	return
}
```

如果两个方法的参数列表和返回值列表相同，则称之为拥有相同类型（same type）

参数是值拷贝，但是如果传入的参数是：slice、pointer、map、function，channel虽然是值拷贝，但是也是引用类型的值，会对其指向的值做出相应变更

你可能会遇到查看某些go的内置func源码的时候它没有声明func的body部分，例如append方法

```go
// The append built-in function appends elements to the end of a slice. If
// it has sufficient capacity, the destination is resliced to accommodate the
// new elements. If it does not, a new underlying array will be allocated.
// Append returns the updated slice. It is therefore necessary to store the
// result of append, often in the variable holding the slice itself:
//	slice = append(slice, elem1, elem2)
//	slice = append(slice, anotherSlice...)
// As a special case, it is legal to append a string to a byte slice, like this:
//	slice = append([]byte("hello "), "world"...)
func append(slice []Type, elems ...Type) []Type
```

事实上append在代码编译的时候，被替换成runtime.growslice以及相关汇编指令了（可以输出汇编代码查看细节），你可以在go的runtime包中找到相关实现，如下：

```go
// growslice handles slice growth during append.
// It is passed the slice element type, the old slice, and the desired new minimum capacity,
// and it returns a new slice with at least that capacity, with the old data
// copied into it.
// The new slice's length is set to the old slice's length,
// NOT to the new requested capacity.
// This is for codegen convenience. The old slice's length is used immediately
// to calculate where to write new values during an append.
// TODO: When the old backend is gone, reconsider this decision.
// The SSA backend might prefer the new length or to return only ptr/cap and save stack space.
func growslice(et *_type, old slice, cap int) slice {
   if raceenabled {
      callerpc := getcallerpc()
      racereadrangepc(old.array, uintptr(old.len*int(et.size)), callerpc, abi.FuncPCABIInternal(growslice))
   }
   if msanenabled {
      msanread(old.array, uintptr(old.len*int(et.size)))
   }
   if asanenabled {
      asanread(old.array, uintptr(old.len*int(et.size)))
   }
 	// 省略...
}
```

声明函数时指定返回值的名称，可以在return时省略

```go
func add(x, y int) (z int, err error) {
  data, err := deal(x, y)
  if err != nil {
    return // 此时等价于return 0, nil
  }
  // 这里是赋值而不是声明，因为在返回值列表中声明过了
  z = x+y
  return // 此时等价于return z, nil
}
```

### 5.2 错误

error是一个接口，因此可以自定义实现error

```go
type error interface {
   Error() string
}
```

如果一个函数执行失败时需要返回的行为很单一可以通过bool来控制

```go
func test(a int) (y int, ok bool) {
  x, ok := test1(a)
  if !ok {
    return 
	}
  y = x*x
  return 
}
```

更多情况下，函数处理时可能遇到多种类型的错误，则使用error，可以通过判断err是否为nil判断是否发生错误

```go
func test(a int) (y int, err error) {
  x, err := test1(a)
  if err != nil {
    return 
	}
  y = x*x
  return 
}
// 打印错误的值
fmt.Println(err)
fmt.Printf("%v", err)
```

Go通过if和return的机制手动返回错误，使得错误的定位更加精确，并且促使你更早的去处理这些错误（而不是像其他语言一样选择抛出异常，可能使得异常由于调用栈的深入，导致最终处理不便）

#### 错误处理策略

一个func的调用返回了err，则调用方有责任正确处理它，下面介绍五种常见处理方式：

1. 传递：

```go
// 某func部分节选
resp, err := http,Get(url)
if err != nil {
  // 将对Get返回的err处理交给当前func的调用方
  return nil, err
}
```

fmt.Errorf()格式化，添加更多描述信息，并创建一个了新的error（参考fmt.Sprintf的格式化）

![image-20220901132429146](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20220901132429146.png)

当error最终被处理的时候，需要反映出其错误的调用链式关系

并且error的内容组织在一个项目中需要统一，以便于后期借助工具统一分析

2. 错误重试

![image-20220902104510073](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20220902104510073.png)

3. 优雅关闭

如果无法处理，可以选择优雅关闭程序，但是推荐将这步工作交给main包的程序，而库函数则选择将error传递给其调用方。

![image-20220902110047619](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20220902110047619.png)

使用log.Fatalf更加方便

![image-20220902110156627](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20220902110156627.png)

会默认输出error的打印时间

![image-20220902110305377](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20220902110305377.png)

4. 选择将错误打印

![image-20220902111135436](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20220902111135436.png)

或者输出到标准错误流

![image-20220902111207116](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20220902111207116.png)

5. 少数情况下，可以选择忽略错误，并且如果错误选择返回，则正确情况下省略else，保持代码整洁

![image-20220902111443372](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20220902111443372.png)

#### EOF（End of File）

输入的时候没有更多内容则触发io.EOF，并且这个error是提前定义好的

![image-20220902113648721](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20220902113648721.png)

![image-20220902113508658](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20220902113508658.png)

### 5.3 函数值

函数是一种类型类型，可以作为参数，并且对应变量是“引用类型”，其零值为nil，相同类型可以赋值

![image-20220902114434886](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20220902114434886.png)

![image-20220902115114210](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20220902115114210.png)

函数作为参数的例子，将一个引用类型的参数传递给多个func，可以为这个参数多次赋值（Hertz框架中使用了这种扩展性的思想）

![image-20220902115541517](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20220902115541517.png)

### 5.4 匿名函数

函数的显式声明需要在package层面，但是在函数的内部也可以创建匿名函数

![image-20220902131528698](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20220902131528698.png)

从上可以看出f存放着匿名函数的引用，并且它是有状态的，维护了一个递增的x

#### 捕获迭代变量引发的问题

正确版本

![image-20220902140757639](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20220902140757639.png)

错误版本

![image-20220902140812965](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20220902140812965.png)

所有循环内创建的func捕获并共享了dir变量（相当于引用类型），所以创建后rmdirs切片内所有元素都有同一个dir，而不是每个元素获得dir遍历时的中间状态

因此正确版本中dir := d的操作为遍历的dir申请了新的内存存放

```go
func main() {
   arr := []int{1, 2, 3, 4, 5}
   temp := make([]func(), 0)
   for _, value := range arr {
      temp = append(temp, func() {
         fmt.Println(value)
      })
   }
   for i := range temp {
      temp[i]()
   }
}
// 结果
5
5
5
5
5
```

另一种错误版本（i最终达到数组长度上界后结束循环，并且导致dirs[i]发生越界）

![image-20220902143256855](https://baize-blog-images.oss-cn-shanghai.aliyuncs.com/img/image-20220902143256855.png)

```go
// 同样是越界的测试函数
func main() {
   arr := []int{1, 2, 3, 4, 5}
   temp := make([]func(), 0)
   for i := 0; i < 5; i++ {
      temp = append(temp, func() {
         fmt.Println(arr[i])
      })
   }
   for i := range temp {
      temp[i]()
   }
}
// 结果
panic: runtime error: index out of range [5] with length 5
```

以上捕获迭代变量引发的问题容易出现在延迟了func执行的情况下（先完成循环创建func、后执行func）

































