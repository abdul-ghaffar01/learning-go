package main

import "testing"

func TestAddUserController(t *testing.T) {
	users := []User{}
	tests := []struct {
		name string
		fullname string
		age int
		city string
		expected bool
	}{
		{"valid age", "Ahmed", 13, "isb", false},
		{"invalid age", "Ahmed", 12, "isb", true},
	}

	for _, tt := range tests {
		err := addUserController(&users, tt.fullname, tt.age, tt.city)
		if tt.expected == false && err != nil {
			t.Errorf("Expected no error but got one %v", err)
		}

		// Should check all the return values as well they should be same in users
	}
}
