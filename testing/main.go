package main

import "fmt"

type User struct {
	Name string `json: "name"`
	Age  int    `json: "age"`
	City string `json: "city"`
}

func addUserController(users *[]User, name string, age int, city string) error {
	fmt.Println("Adding a new user")
	if age < 13 {
		return fmt.Errorf("Minimum age should be 13")
	}
	newUser := User{name, age, city}
	*users = append(*users, newUser)
	fmt.Printf("User %v added as a new user\n", newUser.Name)
	return nil
}

func main() {
	fmt.Println("Testing--------------------")
	users := []User{}
	addUserController(&users, "Abdul Ghaffar", 22, "Islamabad")

	for _, user := range users {
		fmt.Println(user)
	}

}
