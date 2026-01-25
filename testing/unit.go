package main

import "fmt"

type User struct {
	Name string
	Age  int
	City string
}

func addUserController(users *[]User) error {
	fmt.Println("Adding a new user")
	newUser := User{"Abdul Ghaffar", 22, "Daharki"}
	*users = append(*users, newUser)
	fmt.Printf("User %v added as a new user\n", newUser.Name)
	return nil
}

func main() {
	fmt.Println("Testing--------------------")
	users := []User{}
	addUserController(&users)

	for _, user := range users {
		fmt.Println(user)
	}

}
