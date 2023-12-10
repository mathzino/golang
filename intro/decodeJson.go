package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
  FullName string `json:"Name"`
  Age int
}
func main() {

	// decode object User
	var jsonString = `{"Name": "john doe", "Age": 27}`
  
	var jsonData = []byte(jsonString)
  
	var data User
	
	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
	  fmt.Println(err.Error())
	  return
	}
  
	fmt.Println("user :", data.FullName)
	fmt.Println("age  :", data.Age)



	// decode map[string]interface{}
	var data2 interface{}
	json.Unmarshal(jsonData, &data2)

	var decodedData = data2.(map[string]interface{})
	fmt.Println("user :", decodedData["Name"])
	fmt.Println("age  :", decodedData["Age"])


	// decode array object
	var arrayJsonString = `[
    {"Name": "john doe", "Age": 27},
    {"Name": "doe john", "Age": 32}]`

	var data3 []User

	var err2 = json.Unmarshal([]byte(arrayJsonString), &data3)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}

	fmt.Println("user 1:", data3[0].FullName)
	fmt.Println("user 2:", data3[1].FullName)
  }
