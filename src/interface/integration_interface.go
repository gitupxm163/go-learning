package main

import "fmt"

//integration interface

type Phone interface {
	Call(num string)
}

type Camera interface {
	Take() string
}

type SmartPhone interface {
	Phone
	Camera
	Download(url string) string
}

func GetAnyType(any interface{})  {
	switch t := any.(type) {
	case int:
		fmt.Println("this is int type")
	case string:
		fmt.Println("this is string type")
	case *MiPhone:
		fmt.Println("this is MiPhone Pointer type")
	case *IPhone:
		fmt.Println("this is IPhone Pointer type")
	default:
		fmt.Println("unexpect type: ", t)



	}
}



func ListSmartPhoneFunction(mp SmartPhone)  {
	phone, ok:= mp.(*MiPhone)
	if ok {
		phone.Call("10086")
		fmt.Println(phone.Take())
		fmt.Println(phone.Download("bilibili"))
	}else {
		fmt.Println("it`s not a SmartPhone type")
	}

}

// struct
type MiPhone struct {
	version string
}

// implement interface

func (mp *MiPhone) Call(num string) {
	fmt.Println("miphone call: ", num)
}

func (mp *MiPhone) Take() string {
	return "picture url"
}

func (mp *MiPhone) Download(url string) string {
	return "https://bilibili.com"
}

type IPhone struct {
	version string
}

func (mp *IPhone) Call(num string) {
	fmt.Println("Iphone call: ", num)
}

func (mp *IPhone) Take() string {
	return "Iphone take picture url"
}

func (mp *IPhone) Download(url string) string {
	return "Iphone download url https://bilibili.com"
}


func main()  {
	var mp *MiPhone
	mp = new(MiPhone)
	mp.version = "Mi 10"

	ListSmartPhoneFunction(mp)

	fmt.Println("=======")
	iphone := new(IPhone)
	iphone.version = "iphone 12"

	ListSmartPhoneFunction(iphone)
	GetAnyType(iphone)
	fmt.Println("========")
	map1 := new(map[string]string)
	GetAnyType(map1)

}
