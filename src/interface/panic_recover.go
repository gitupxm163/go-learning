package main

import "fmt"

//panic and recover test
func ComputeValue(devidee, devideer int) int {

	result :=  devidee / devideer
	return result
}

// define common recover protect func
type ComputeFunc func(a,b int) int

func Protect(f ComputeFunc,a,b int) int  {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("catch panic in recover, error information is: %v\n", e)
		}

	}()

	return f(a,b)

}

func main()  {
	fmt.Println("START")
	//ComputeValue(100,0)
	fmt.Println("result is: ", Protect(ComputeValue,100,10))
	fmt.Println("END")
}
