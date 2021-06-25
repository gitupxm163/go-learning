package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// read file

func ReadFile(fileName string) (string, error)  {

	file, err := os.Open(fileName)
	defer fmt.Println("==END==")
	defer file.Close()
	defer fmt.Println("==START==")
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// error test

// error interface : Error() string

type Devide struct {
	devidee int
	devider int
}

func (d *Devide) Error() string  {
	errContent :=`
	cannot proceede, the devider is zero.
	devidee: %d
	devider: 0
`
	return fmt.Sprintf(errContent, d.devidee)
}

func Compute(d *Devide) (result int, err error){
	fmt.Println(d.devider, d.devidee)
	if d.devider == 0 {
		err = d
	}else {
		result = d.devidee/d.devider
	}
	return
}




func main()  {

	file, err := ReadFile("abc.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(file)

	fmt.Println("============")

	err = errors.New("this a new error.")
	err = fmt.Errorf("error: %v\n", err.Error())
	fmt.Println(err)

	fmt.Println("=============")
	d := new(Devide)
	d.devidee = 100
	d.devider = 1
	result, err := Compute(d)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println("result: ", result)
	}
}