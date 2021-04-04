package main

import (
	"fmt"
	"strconv"
)

var myGlobalInfo int = 662

var (
	contact int    = 0140
	email   string = "alfaBinomialbeta23@gmail.com"
)

var actorName string = "Md Ruhul Amin"
var actorAge int = 32
var actorInformation string = "He is a good man"

func main() {
	var i int
	var j int
	var k int
	i = 232
	j = 2252
	k = 323
	fmt.Println("hi everyone! ")
	fmt.Println("what are you doing here@")
	fmt.Println(i + j + k)
	var ruhul_amin int = 25522
	fmt.Println(ruhul_amin)
	sakib_hasan := 1996
	fmt.Println(sakib_hasan)

	// getting the type of the data of value and type

	var floatData float32 = 63.222
	fmt.Println(floatData)

	
	myData := 855.22 // showing float64 of the data type!
	fmt.Printf("%v,%T", myData, myData)
	fmt.Println()
	fmt.Println(myGlobalInfo)
	fmt.Println(actorName)
	fmt.Println(actorInformation)
	fmt.Println(actorAge)

	fmt.Println(email)
	fmt.Println(contact)

	// type convertion

	var x float32 = 522.02214
	fmt.Println(int(x))

	var myNumber int = 100121
	fmt.Println(strconv.Itoa(myNumber))

}
