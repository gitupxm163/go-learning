package main

import (
	"calculate/calculate"
	"fmt"
)

func main()  {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d is even: %v\n", i, calculate.Even(i))
	}
}
