package main

import "fmt"


type User struct {
	name string
	age int
	userName string
	email string
	isActive bool
}

func structsAndMethods(){

	/*
		Struct is a composite data type that groups related data together.
	*/

	// Named field initialization	-> Recommended way of initialization
	u := User{
		name: "Abdul Ghaffar",
		age: 22,
		userName: "ags",
		email: "agscontact777@gmail.com",
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
	p.name = "Azlan"  // automatically dereferenced
	fmt.Printf("%+v\n", u1)

	// ------------ Methods --------------
	// A method is a function with a receiver
	
	
}