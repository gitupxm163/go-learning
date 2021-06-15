package main

import "fmt"

func main()  {

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := arr[2:6]
	fmt.Println(s2,len(s2),cap(s2),arr)

	s2 = append(s2, 10)
	fmt.Println(s2,len(s2),cap(s2),arr)

	s2 = append(s2, 20, 30, 40, 50)
	fmt.Println(s2,len(s2),cap(s2),arr)

	//copy
	s3 := make([]int, 3)
	copy(s3,s2)
	fmt.Println(s3,s2)

	s2 = append(s2[:4], s2[5:]...)
	fmt.Println(s2)

}
