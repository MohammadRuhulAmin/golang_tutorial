package main

import "fmt"


func main() {
	myNumb_1 :=1 ==1 
	// myNumb_2 :=2 ==1 

	fmt.Println("-----------------")

	fmt.Println("%v,%T",myNumb_1,myNumb_1)


	fmt.Println("this is something to learn!")
	var myString string = "this is my first string"
	fmt.Println(myString)
	var n bool = true 
	fmt.Println("%v,%T\n",n,n)
	s := "this is a string"
	fmt.Println(string(s[2])); 
	s1 :="string 1"
	s2 :="string 2"
	var s3 string = s1+"  " +s2 
	fmt.Println(s3)

	var str_1 string = "this is going to be a string"
	byte_1 := []byte(str_1)
	fmt.Println(byte_1)


	 

}
