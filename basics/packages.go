package main

/*
Rules for package main:
	Executable programs must use package main
	The function main() is the entry point
	libraries never use package main


Rules for package:
	Must be the first non-comment line
	All files in the same directory must declare the same package
	Package name is usually the directory name
*/

/*
Import syntax:

import "single"

import (
	"first"
	"second"
)

Common Packages:
1. fmt
2. os
3. net/http
4. io
5.time
6. strings
7. sync
8. strconv
*/

// Functions that start with capital character are exported while others aren't
// Same applies to Variables, Constants, Types, Struct fields
// Examples:
/*
	func Sum(a, b int) int {	 Exported
		return a+b;
	}

	func sum(a, b int) int {	 Private
		return a+b;
	}

	type Person struct{
		Name string -> Exported
		age int -> private
	}
*/
