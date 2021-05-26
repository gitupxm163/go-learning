package main

import (
	"fmt"
	"math"
)

//type define and type alias

func typeDefineAndTypeAlias()  {

	type MyInt1 int
	type MyInt2 = int
	var i1 MyInt1 = 1
	var i2 MyInt2 = 2
	var i3 int = 2
	i3 = int(i1)

	i3 = i2

	fmt.Printf("%T %T %T\n", i1, i2, i3)

}

// enum

func enumDemo()  {
	const(
		Monday = 1 + iota
		Tuesday
	)

	fmt.Println(Monday,Tuesday)
}

// const
func constDemo()  {
	const s1 string = "hello golang"
	const a, b = 3, 4
	fmt.Println(s1, a, b)

	var c int
	c = int(math.Sqrt(a*a + b*b))

	fmt.Println(c)
}



func main()  {
	typeDefineAndTypeAlias()
	enumDemo()
	constDemo()
}

