package main

import "fmt"

type Person struct {
	FirstName, SurName, Email string
	Age                       int
}

func describePerson(p Person) {

	fmt.Printf("This person's name is: %s, %s, and they are %d years old. If you want to contact them, their email is: %s",
		p.FirstName, p.SurName, p.Age, p.Email)

}

func main() {
	u := Person{FirstName: "Joe", SurName: "O'Reilly", Email: "joe@test.com", Age: 28}
	describePerson(u)
	// fmt.Println("hello world")
	// fmt.Println(u)
}
