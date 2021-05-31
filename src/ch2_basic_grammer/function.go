package main

import (
	"fmt"
	"os"
)

func swap(a,b int) (int, int)  {
	return b, a
}

func openFile()  {
	file,err := os.Open("for.go")
	if err != nil {
		fmt.Println("打开文件有误")
	}else {
		fmt.Println(file)
	}
}

func operate(x,y int, op string) (int, error) {
	switch op {
	case "+":
		return x + y, nil
	case "-":
		return x - y, nil
	case "*":
		return x * y, nil
	case "/":
		return x / y, nil
	default:
		return 0, fmt.Errorf("do not support operate")
	}
}

func say(n Name, world string)  {
	fmt.Println(n(world))
}

type Name func(name string) string

func world(world string) string {
	return "hello " + world
}

func golang(golang string) string {
	return "hi "+ golang
}
func main()  {

	x, y := swap(1, 2)
	fmt.Println(x,y)

	fmt.Println("=====")
	openFile()

	fmt.Println("=====")
	result, err := operate(4, 2, "x")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println(result)
	}

	fmt.Println("======")
	say(world,"world")

	say(golang,"golang")
}
