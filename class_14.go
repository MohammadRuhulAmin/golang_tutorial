package main

import (
	"fmt"
)

func myInformation(myName func(string) string, myEmail func(string) string, myContact func(string) string) {
	fmt.Println(myName("Md Ruhul Amin"))
	fmt.Println(myEmail("alfaBinomialbeta23@gmail.com"))
	fmt.Println(myContact("0140-3288711"))
}

func main() {

	myName := func(name string) string {
		return name
	}

	myEmail := func(email string) string {
		return email
	}
	myContact := func(contact string) string {
		return contact
	}

	myInformation(myName, myEmail, myContact)

}
