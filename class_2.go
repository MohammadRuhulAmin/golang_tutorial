package main

import "fmt"

func main() {

	//  grades := [3]int  // array declerations
	//  grades[0]=1
	//  grades[1] = 2
	myNumbers := [5]int{12, 22, 33, 44, 2}
	fmt.Println("myNumbers ", myNumbers)

	myString := [5]string{"ruhul amin", "sakib", "sajid", "saifullah", "ashraful"}
	fmt.Println("my friends  list are : ", myString)

	myLongList := [...]string{"first", "second", "third"}
	fmt.Println("My Friends long list : ", myLongList)

	var students [3]string
	students[0] = "ruhul"
	students[1] = "amin"
	students[2] = "sakib"

	fmt.Println(students)

	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)

	var myPersonalInfo [3][3]string
	myPersonalInfo[0] = [3]string{"ruhul", "amin", "md"}
	myPersonalInfo[1] = [3]string{"MPBHS", "NDC,DHAKA", "AIUB"}
	myPersonalInfo[2] = [3]string{"1223", "323", "#23323"}

	fmt.Println(myPersonalInfo)

	fmt.Println("--------------")
	var myKnowledge [3][3]string
	myKnowledge[0] = [3]string{"ranu", "ruhul", "kamal"}
	myKnowledge[1] = [3]string{"sakib", "jeba", "rahul"}
	myKnowledge[2] = [3]string{"kamal", "tomal", "masum"}

	fmt.Println(myKnowledge)

	// pointing to an array

	a := [3]int{1, 2, 3}
	b := a
	fmt.Println(a)
	fmt.Println(b)

	c := b
	fmt.Println(c)
	d := &a
	d[2] = 1220

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d) 

}
