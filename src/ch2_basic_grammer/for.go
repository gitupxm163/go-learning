package main

import "fmt"

func main()  {

	a := 20

	for a < 25 {
		a++
		if a == 24 {
			break
		}
		fmt.Println(a)
	}

	fmt.Println("=======")

	a = 20
	for a < 25 {
		a++
		if a == 24 {
			continue
		}
		fmt.Println(a)

	}

	fmt.Println("======")

	a = 20

LOOP:
	for a < 25 {
		a++
		if a == 24 {
			goto LOOP
		}
		fmt.Println(a)
	}


	fmt.Println("========")

   sum := 0

	for i:=0; i<=100; i++{
		sum += i
	}

	fmt.Println(sum)

}
