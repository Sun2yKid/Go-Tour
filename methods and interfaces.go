package main

import (
	"fmt"
	"io"
	"math"
	"strings"
	"time"
)

type Vertex struct {
	X, Y float64
}

//方法
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//函数
func Abs_func(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 指针接收者的方法，接收者既能为值也能为指针
func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

//
func (v Vertex) Scale_nop(f float64) {
	v.X *= f
	v.Y *= f
}

// 接受一个值为参数的函数必须接收一个指定类型的值
func Scale(v Vertex, f float64) {
	v.X *= f
	v.Y *= f
}

// 带指针参数的函数必须接受一个指针
func Scale_p(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}

// 为非结构体类型声明方法
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}

// 接口类型是由一组方法签名定义的集合
type Abser interface {
	Abs() float64
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice of %v is %v\n", v, v*2)
	case string:
		fmt.Printf("lenth of %q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("don't know the type %T\n", v)
	}
}

// Stringer

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v old)", p.Name, p.Age)
}

type IPAddr [4]byte

func (p IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", p[0], p[1], p[2], p[3])
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%v time happen %v\n", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"shit",
	}
}

func main() {
	defer fmt.Println("Finished!")

	fmt.Println("Hello")
	v := Vertex{3, 4}
	fmt.Println(v, v.Abs())
	fmt.Println(v, Abs_func(v))
	fmt.Println("math.Sqrt2", math.Sqrt2)
	f := MyFloat(-math.Sqrt2)
	fmt.Println("f, f.abs()", f, f.Abs())
	v.Scale(2)
	fmt.Println("v.Scale(2)", v)
	Scale(v, 0.5)
	fmt.Println("Scale(v, 0.5)", v)
	Scale_p(&v, 0.5)
	fmt.Println("Scale_p(&v, 0.5)", v)
	p := &v
	p.Scale(2)

	fmt.Println("p.Scale(2)", v)
	v.Scale_nop(3)
	fmt.Println("v.Scale_nop", v)

	var a Abser
	//a = f
	//a = &f
	//a = v
	a = &v

	fmt.Println(a.Abs())
	describe(a)

	var v2 Vertex
	a = &v2
	fmt.Println(a.Abs())
	describe(a)
	var f2 MyFloat
	a = f2
	fmt.Println(a.Abs())
	describe(a)

	var i interface{} = "Hello"
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	ff, ok := i.(float64)
	fmt.Println(ff, ok)

	//fff := i.(float64) // panic: interface conversion: interface {} is string, not float64
	//fmt.Println(fff)

	do(43)
	do("night")
	do(true)

	c := Person{"Yao", 18}
	b := Person{"Kid", 2}
	fmt.Println(c, b)

	hosts := map[string]IPAddr{
		"loopback": {127, 0, 0, 1},
		"google":   {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Println(name, ip)
		fmt.Printf("%v: %v\n", name, ip)
	}

	if err := run(); err != nil {
		fmt.Println(err)
	}

	r := strings.NewReader("Hello, Reader!")
	d := make([]byte, 8)
	for {
		n, err := r.Read(d)
		fmt.Printf("n:%v, err:%v, b:%v", n, err, d)
		fmt.Printf("d[:n]=%q\n", d[:n])
		if err == io.EOF {
			break
		}
	}

}

func describe(i Abser) {
	fmt.Printf("(%v, %T)\n", i, i)
}
