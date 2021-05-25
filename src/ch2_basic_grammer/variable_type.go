package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
)

func main()  {

	var a int = 1
	var b int32 = 2
	b = int32(a)
	fmt.Printf("%T %T\n", a, b)

	// 运行时包
	goarch := runtime.GOARCH
	goos := runtime.GOOS
	intSize := strconv.IntSize

	fmt.Println(goarch, goos, intSize)
	// float
	var f1 float32
	var f2 float64

	fmt.Printf("%T %T\n", f1, f2)
	fmt.Printf("%f %f\n",f1, f2)

	// byte rune
	var bt byte = 2
	var ru rune = '中'

	fmt.Printf("%T %T\n" , bt, ru)

	// math.sqrt
	var c, d = 3, 4
	var e int
	temp := c*c + d*d
	e = int(math.Sqrt(float64(temp)))
	fmt.Printf("%T %d\n",e, e)

	// ptr
	var a1 = 1
	ptr := &a1
	ptrptr := &ptr
	fmt.Printf("%T %T %T\n", a1, ptr, ptrptr)
}
