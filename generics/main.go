// 一个简单的泛型示例
package main

import "fmt"

type iface1 interface {
	Foo()
}

type iface2[T iface1] interface {
	Get() T
}

type s1 struct{}

func (s *s1) Foo() {
	fmt.Println("this is s1.Foo")
}

type S2 struct {
	s1 *s1
}

func (s *S2) Get() *s1 {
	return s.s1
}

func NewS2() *S2 {
	return &S2{
		s1: &s1{},
	}
}

func test[T iface1](obj iface2[T]) {
	obj.Get().Foo()
}

func main() {
	obj := NewS2()
	test[*s1](obj)
}
