package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// GenerateSlice 创建一个包含从 1 到 n 的整数的切片，并随机打乱顺序。
func GenerateSlice(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = i + 1
	}
	// 打乱切片
	rand.Shuffle(n, func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	return slice
}

func main() {
	size := []int{128, 512, 1024, 2048, 100000, 1000000, 10000000}
	sortVersion := []struct {
		name string
		sort func([]int)
	}{
		{"Mergesort V1", SequentialMergesortV1},
		{"Mergesort V2", SequentialMergesortV2},
		{"Mergesort V3", SequentialMergesortV3},
	}
	for _, s := range size {
		fmt.Printf("Testing Size: %d\n", s)
		o := GenerateSlice(s)
		for _, v := range sortVersion {
			s := make([]int, len(o))
			copy(s, o)
			start := time.Now()
			v.sort(s)
			elapsed := time.Since(start)
			fmt.Printf("%s: %s\n", v.name, elapsed)
		}
		fmt.Println()
	}
}

func SequentialMergesortV1(s []int) {
	if len(s) <= 1 {
		return
	}
	middle := len(s) / 2
	SequentialMergesortV1(s[:middle])
	SequentialMergesortV1(s[middle:])
	Merge(s, middle)
}

func Merge(s []int, middle int) {
	left := s[:middle]
	right := s[middle:]

	// 初始指针位置
	i := 0 // 左子切片的指针
	j := 0 // 右子切片的指针
	k := 0 // 原切片的指针

	// 合并两个子切片到原切片
	tmp := make([]int, len(s))
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			tmp[k] = left[i]
			i++
		} else {
			tmp[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		tmp[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		tmp[k] = right[j]
		j++
		k++
	}
}

func SequentialMergesortV2(s []int) {
	if len(s) <= 1 {
		return
	}
	middle := len(s) / 2

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		SequentialMergesortV2(s[:middle])
	}()
	go func() {
		defer wg.Done()
		SequentialMergesortV2(s[middle:])
	}()
	wg.Wait()
	Merge(s, middle)
}

const mx = 2048

func SequentialMergesortV3(s []int) {
	if len(s) <= 1 {
		return
	}
	if len(s) < mx {
		SequentialMergesortV1(s)
	} else {
		middle := len(s) / 2

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			SequentialMergesortV3(s[:middle])
		}()
		go func() {
			defer wg.Done()
			SequentialMergesortV3(s[middle:])
		}()

		wg.Wait()
		Merge(s, middle)
	}
}
