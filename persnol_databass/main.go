package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var choice int

type User struct {
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	PhoneNo string  `json:"phone_no"`
	Address Address `json:"address"`
}
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	PinCode string `json:"pin_code"`
}

var userDetails = map[string][]User{
	"8930008915": {User{Name: "Vikash", Email: "vikash@123", PhoneNo: "893008915", Address: Address{Street: "Street 1", City: "City1", State: "Haryana", PinCode: "123456"}}},
	"7988323110": {User{Name: "Pawan", Email: "pawan@123", PhoneNo: "7988323110", Address: Address{Street: "Street 3", City: "City1", State: "Haryana", PinCode: "123456"}}},
	"8295068010": {User{Name: "Ashu", Email: "ashu@123", PhoneNo: "8295068010", Address: Address{Street: "Street 2", City: "City1", State: "Haryana", PinCode: "123"}}},
}

func greet() {
	fmt.Println("Welcome to our application enter your choice")
	fmt.Println("1: Print existing data")
	fmt.Println("2: Search by phone number")
	fmt.Println("3: Add new member")

}
func PrintVal() {
	data, err := json.MarshalIndent(userDetails, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	var printData = fmt.Sprint(string(data))
	fmt.Println(printData)
}
func searchByPhone(PhoneNo string) {

	fmt.Println(userDetails[PhoneNo])

}

func addMember(firstName, email, phoneNo string) {
	newUser := User{
		Name:    firstName,
		Email:   email,
		PhoneNo: phoneNo,
		// Address: Address{}, // Initialize Address with empty values
	}
	userDetails[phoneNo] = append(userDetails[phoneNo], newUser)

}

func main() {

	greet()
	fmt.Println("Enter Here!")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		PrintVal()
	case 2:
		fmt.Println("Enter phone number to search")
		var searchPhoneNo string
		fmt.Scan(&searchPhoneNo)
		searchByPhone(searchPhoneNo)
	case 3:
		var firstName string
		var email string
		var phoneNo string
		// var address Address
		fmt.Println("Enter first name:")
		fmt.Scan(&firstName)
		fmt.Println("Enter email id:")
		fmt.Scan(&email)
		fmt.Println("Enter Your Phone number:")
		fmt.Scan(&phoneNo)
		addMember(firstName, email, phoneNo)
		PrintVal()
	case 4:
		var phoneNo string
		fmt.Println("Enter User key which is user phone number :")
		fmt.Scan(&phoneNo)
		
	default:
		fmt.Println("Invalid choice")
	}
}
