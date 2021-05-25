package main

import "fmt"


var(
	aa = 11
	bb = 22
	cc = 33
)
func initVariable()  {
	var a, b, c = 1, "b", true
	var t1 bool
	h, d, e := 2, "d", false
	var str1 = "string"
	fmt.Println(a, b, c)
	fmt.Println(t1)
	fmt.Println(h, d, e)
	fmt.Println(str1)
}


func main()  {
	var a int
	var b string = "hello"
	fmt.Printf("%d %s\n", a, b)
	fmt.Println(aa, bb, cc)
	initVariable()
}
