package main

import "fmt"

func main()  {

	var a = 120

	switch a {
	case 19:
		fmt.Println("a=19")
	case 20:
		fmt.Println("a=20")
		fallthrough
	case 21,22,23,24,25,26,120:
		fmt.Println("a in [21,22,23,24,25,26,120]")
	default:
		fmt.Println("UnKnown")
	}

}
