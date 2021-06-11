package main

import "fmt"

func main()  {

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Printf("arr=%v\n",arr)

	s2 := arr[2:6]
	fmt.Printf("s2=%v, len=%d, cap=%d\n",s2,len(s2), cap(s2))

	s3 := arr[2:]
	fmt.Println(s3)

	// create slice based a slice

	s7 := s2[2:5]
	fmt.Println(s7)

	// create slice by make()
	arr1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	s8 := arr1[2:6]

	fmt.Printf("s8=%v, len=%d, cap=%d\n", s8, len(s8), cap(s8))

	s9 := s8[2:6]
	fmt.Println(s9)

}
