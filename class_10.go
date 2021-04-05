package main 

import "fmt"

func main(){
	var mydata interface{} = "object --> "
	fmt.Println(mydata)
	var a int = 42 
	var b *int = &a  
	fmt.Println(&a,"  ",b)
	
	if &a == b {
		fmt.Println("both address are same")
	}else{
		fmt.Println("address are not same ")
	}

	var a1 int = 22 
	fmt.Println(&a1)

}