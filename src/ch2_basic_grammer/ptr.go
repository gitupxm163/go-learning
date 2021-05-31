package main

import "fmt"

func valPass(a int)  {
	a = 100
}

func refPass(a *int)  {
		*a = 200
}

func main()  {

	// 普通变量
	var a int = 10
	// ptr变量
	var b *int
	// 给ptr变量赋值
	b = &a
	fmt.Println(*b)
	// 通过ptr变量给实际变量赋值
	*b = 20
	fmt.Println(a)

	fmt.Println("====")
	valPass(a)
	fmt.Println(a)

	fmt.Println("====")
	refPass(&a)
	fmt.Println(a)
}
