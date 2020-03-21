package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

type Vertex struct {
	X int
	Y int
}

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

type Location struct {
	Lat, Long float64
}

var m map[string]Location

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(x int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	defer fmt.Println("Bye")
	today := time.Now().Weekday()
	fmt.Println("hello, number:", rand.Intn(100), "It's ", today)
	// 闭包
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}

	// 函数值
	fmt.Println("math.Pow", compute(math.Pow))
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println("hypot", compute(hypot))

	//映射
	fmt.Println("m", m)
	m = make(map[string]Location)
	fmt.Println("m", m)
	m["Bell Labs"] = Location{0.11, 0.22}
	fmt.Println("m", m, m["Bell Labs"])

	m2 := make(map[string]int)
	m2["answer"] = 42
	fmt.Println(m2, m2["answer"])
	m2["answer"] = 44
	fmt.Println(m2, m2["answer"])
	delete(m2, "answer")
	fmt.Println(m2, m2["answer"])
	value, ok := m2["answer"]
	fmt.Println(value, ok)

	// range
	//for 循环的 range 形式可遍历切片或映射。
	for i, v := range pow {
		fmt.Printf("2**%d=%d\n", i, v)
	}

	pow1 := make([]int, 10)
	printSlice("pow1", pow1)
	for i := range pow1 {
		pow1[i] = 1 << i
	}
	printSlice("pow1", pow1)

	//切片的切片,创建一个井字板
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	fmt.Println("board", board)
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], "|"))
	}

	//切片
	//make 函数会分配一个元素为零值的数组并返回一个引用了它的切片：
	f := make([]int, 5)
	printSlice("f", f)

	//数组
	// 类型[n]T表示拥有n个T类型的值的数组
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{1, 2, 3, 4, 5}
	fmt.Println(primes) // last defalut: 0

	var d []int = primes[1:4]
	fmt.Println(d)

	e := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
		{3, false},
	}
	fmt.Println(e)

	// 结构体
	fmt.Println(Vertex{1, 2})
	v := Vertex{3, 4}
	fmt.Println(v.X, v.Y)
	v.X = 5
	fmt.Println(v.X, v.Y)
	s := &v
	fmt.Println("s, *s", s, *s)
	s.X = 6
	fmt.Println(v.X, v.Y)
	v2 := Vertex{X: 1}
	fmt.Println(v2)

	/*
		类型*T是指向T类型值的指针。其零值为nil
		& 操作符会生成一个指向其操作数的指针。
		* 操作符表示指针指向的底层值。
	*/
	i, j := 43, 2701
	fmt.Println("i, j", i, j)
	p := &i
	fmt.Println("p, *p", p, *p)
	*p = 520
	fmt.Println("i, j", i, j)
	p = &j
	fmt.Println("p, *p", p, *p)
	*p = *p / 37
	fmt.Println("i, j", i, j)

}
