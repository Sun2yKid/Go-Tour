package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

var c, python, java bool

const pi = 3.1415

func sqrt(x float64) (z float64) {
	z = 1.0
	for i := 1; i <= 10; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return
}

func main() {
	defer fmt.Println("Finished!")

	for i := 0; i < 4; i++ {
		defer fmt.Println("counting:", i)
	}
	defer fmt.Println("Shutting down:")

	today := time.Now()
	today_day := time.Now().Weekday()
	fmt.Println(today, today_day)
	switch {
	case today.Hour() < 12:
		fmt.Println("morning!")
	case today.Hour() < 17:
		fmt.Println("afternoon")
	default:
		fmt.Println("Good evening")
	}

	fmt.Println("Go go on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("os:", os)
	}
	os := runtime.GOOS
	fmt.Println("runtime.GOOS:", os)

	fmt.Println("sqrt:", sqrt(2))
	fmt.Println("math.sqrt:", math.Sqrt(2))

	fmt.Println("Good nums:", rand.Intn(100))
	fmt.Println("math.pi", math.Pi)
	fmt.Println("x+y:", add(1, 3))
	a, b := swap("moring", "good")
	fmt.Println("x swap y:", a, b)
	x, y := split(10)
	fmt.Println("split 10:", x, y)

	var i int = 333
	k := 444
	var f float64 = 22.22
	d := float64(f)
	fmt.Println(i, c, python, java, k, d)

	var c, python, java = true, 1, "no"
	fmt.Println(c, python, java)

	fmt.Println(pi)

	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
		fmt.Println("hello,", i, sum)
	}

	for sum < 1000 {
		sum += sum
		fmt.Println("wo~", sum)
	}

}

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum / 2
	y = sum * 2
	return
}
