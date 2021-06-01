package main

import "fmt"

func changeVal1(arr [3]int)  {
	arr[1] = 100
}

func changeVal2(arr *[3]int){
	arr[0] = 100
}


func main()  {

	// arr
	var arr1 [3]int
	fmt.Println(arr1)

	// arr ...
	 var arr2 = [...]int{1,2,3}
	 fmt.Println(arr2)

    // 2 d arr
    var arr3 = [3][3]int{
    	{1,2,3},
    	{4,5,6},
    	{7,8,9},
	}

	fmt.Println(arr3)

    // change val
    var arr4 = [3]int{1,2,3}
    changeVal1(arr4)

    fmt.Println(arr4)

    changeVal2(&arr4)

    fmt.Println(arr4)


}
