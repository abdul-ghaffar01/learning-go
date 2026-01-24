package main

import (
	"encoding/json"
	"fmt"
)


type UserData struct {
	Name string		`json:"name"`
	Age int			`json:"age"`
	City string		`json: "city"`
}

func encoding(u UserData){
	jsonData, err := json.Marshal(u)

	if err != nil {
		fmt.Println("Couldn't create json, error: ", err)
	}

	fmt.Println("json data:", string(jsonData))

	// Decoding the same data
	decoding(string(jsonData))
}

func decoding(jsonInString string){
	fmt.Println("Decoding the json-------------")

	/*
		Steps:
			1) Convert string of json to slice of byte
			2) 
	*/

	data := []byte(jsonInString)
	var userDataInstruct UserData
	err := json.Unmarshal(data, &userDataInstruct)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(userDataInstruct)

}

func JsonEncodingDecoding(){
	fmt.Println("Json Encoding Decoding-------------------")
	var u UserData
	u = UserData{"Abdul Ghaffar", 22, "Daharki"}
	fmt.Println(u)

	fmt.Println("Encoding of user object:")
	encoding(u)

}