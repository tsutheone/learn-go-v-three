package main

import "fmt"

type Person struct {
	first string
	age   string
}

func main() {
	p1 := Person{
		first: "James Bond",
		age:   "42",
	}
	p1.Speak()
}

func (p Person) Speak() {
	fmt.Println("I am", p.first, "and my age is", p.age)
}

/*
● Create a user defined struct with
	○ the identifier “person”
	○ the fields:
		■ first
		■ age
● attach a method to type person with
	○ the identifier “speak”
	○ the method should have the person say their name and age
● create a value of type person
● call the method from the value of type person
*/
