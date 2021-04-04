package main 

import "fmt"

func main(){

	//myArray := [10]int{1,2,3,4,5,6,7,8,9,10}
	myArray := [...]int{1,2,3,4,5,6,7,8,9,10}

	fmt.Println(len(myArray)) // count length function !

	a:= myArray[:]
	b:= myArray[1:5]
	c:= myArray[:7]
	d:= myArray[6:]
	myArray[2]=552
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	slycingElem :=[]int{}
	slycingElem = append(slycingElem,2,34,55)

	slycingElem = append(slycingElem,-552,8855)

	fmt.Println(slycingElem)

	

	






}