package main

import (
	"fmt"
	"strconv"
)

type Integer int

func (it Integer) string() {
	itString := strconv.Itoa(int(it))
	fmt.Println(itString)
}

// interface
type Speaker interface {
	Say(message string)
}

type Person struct {
	name string
}

func (p *Person) Say(message string) {
	fmt.Println(p.name," say : ", message)
}

func SpeakerSay(speaker Speaker, message string)  {
	speaker.Say(message)
}

func main()  {
	it := Integer(100)
	it.string()

	fmt.Println("========")
	person := new(Person)
	person.name = "James"

	//person.Say("hello Interface!")
	fmt.Println("========")

	SpeakerSay(person,"this is James")

}
