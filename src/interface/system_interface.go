package main

import (
	"fmt"
	"sort"
)

// 解引用

type Appender interface {
	Append(int)
}

type Lener interface {
	Len() int
}

type List []int

func (l *List) Append(val int) {
	*l = append(*l, val)
}

func (l List) Len() int {
	return len(l)
}

// 定义调用方法
func Counter(a Appender, start, end int)  {

	for i := start; i < end ; i++ {
		a.Append(i)
	}
}

func IsLarge(l Lener)  bool {
	return l.Len() > 50
}

// 系统接口 Stringer

type Course struct {
	name string
}

func (c *Course) String() string {
	return fmt.Sprintf("name=%v", c.name)
}

// 系统接口 sort.Interface
func Sort(data sort.Interface)  {
	for pass := 0; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-1; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}




type MyList []int

func (l *MyList) Len()	int{
	return len(*l)
}

func (l MyList) Less(i, j int) bool {
	return l[i] < l[j]
}

func (l MyList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}


func main()  {
	l := new(List)
	Counter(l,0,10)
	fmt.Println(l)

	fmt.Println("====")
	fmt.Println("list is large than 50: ", IsLarge(l))

	fmt.Println("====")
	c := new(Course)
	c.name = "golang"

	fmt.Println(c)

	myList := new(MyList)

	*myList = []int{20, 10 ,12, 98}

	Sort(myList)

	fmt.Println(myList)

}

















