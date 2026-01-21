package main

import "fmt"

// Polymorphism:
// Different things respond to the same thing in different ways

/*
Go achieves polymorphism through:
- Interfaces
- Implicit interface implementation
- Method sets
- Composition (embedding)
- Function-based polymorphism
- Generics
*/

/*
Interface:
It defines behavior not data.
*/

type Speaker interface{
	Speak() string
}
// Any type that has method Speak implements Speaker interface automatically


type human struct {
	name string
}

func (h human) Speak() string {
	return "I am " + h.name
}

type cat struct {
	name string
}

func (c cat) Speak() string {
	return "meowwww"
}

func makeSpeak(s Speaker){
	fmt.Println(s.Speak())
}


func interfacesAndPolymorphism() {
	fmt.Println("-------------Interfaces-------------")

	h := human{"Abdul Ghaffar"}
	// h.Speak()
	c := cat{"kinky"}
	// c.Speak()


	/// Now using interface -> Same function, same interface but different behavior 
	makeSpeak(h)
	makeSpeak(c)

	// Implicit interface implementation
	// A type implements an interface implicitly by having the required methods.

}
