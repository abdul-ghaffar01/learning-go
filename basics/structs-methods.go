package main

import (
	"fmt"
)

type User struct {
	name     string
	age      int
	userName string
	email    string
	isActive bool
}

// Method
func (u User) greet() { // value receiver does not modify the original struct.
	fmt.Println("hello", u.name)
}

func (u *User) updateUsername() { // pointer receiver does modify original struct.
	u.userName = "newUsername"
}

// Constructor to create user -> also called factory function
func NewUser(name, userName, email string, age int, isActive bool) (*User, error) {
	if age < 13 {
		return nil, fmt.Errorf("Minimum age should be 13")
	}

	// u := User{name, age, userName, email, isActive}
	return &User{name, age, userName, email, isActive}, nil

}

// Composition/embedding in structs 
type livingThing struct {
	heartBeat int
	isAlive bool
}

type Animal struct {
	name string
	livingThing  // now Animal contains all the properties of livingThing
}

type Human struct {
	name string
	livingThing
}

type Person struct {
	age int
	height float32
	Human	// contains everything from human and livingThing.
}

func structsAndMethods() {

	/*
		Struct is a composite data type that groups related data together.
	*/

	// Named field initialization	-> Recommended way of initialization
	u := User{
		name:     "Abdul Ghaffar",
		age:      22,
		userName: "ags",
		email:    "agscontact777@gmail.com",
		isActive: true,
	}

	fmt.Printf("%+v\n", u)

	// Postitional initialization	-> Avoid to use
	u1 := User{"Abdul Ghaffar", 22, "ags", "ags@gmail.com", true}
	fmt.Printf("%+v\n", u1)

	// Accessing and modifying the fields
	fmt.Println(u.email)
	u.email = "check@gmail.com"
	fmt.Println(u.email)

	// Copying
	u2 := u1
	u2.email = u.email
	fmt.Printf("%+v\n", u2)

	// Pointers with structs
	p := &u1
	p.name = "Azlan" // automatically dereferenced
	fmt.Printf("%+v\n", u1)

	// ------------ Methods --------------
	// A method is a function with a receiver
	u1.greet()          // u1 is receiver.
	u1.updateUsername() // <- go automatically takes reference
	fmt.Printf("%+v\n", u1)

	// Another way to initailize
	u3 := &User{} // u3 is a pointer
	fmt.Printf("%+v\n", u3)
	u3.name = "Ahmed"
	u3.greet()

	// Constructors
	u4, err := NewUser("Aizal", "aizlo", "aizal@gmail.com", 14, true)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("User Created\n%+v\n", u4)

	u5, err := NewUser("Aizal", "aizlo", "aizal@gmail.com", 12, true)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%+v\n", u5)
	}

	// Created an object of embedded structs
	person := Person{
		age: 22,
		height: 5.5,
		Human: Human{
			name: "Abdul Ghaffar",
			livingThing: livingThing{
				heartBeat: 60,
				isAlive: true,
			},
		},
	}

	fmt.Printf("%v\n", person)
}
