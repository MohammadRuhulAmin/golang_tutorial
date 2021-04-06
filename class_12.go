package main 

import (
	"fmt"
)

func test(){
	fmt.Println("testing a function ")
}

func myParameter(msg string,age int){
	fmt.Println(msg, " ", age)
}

func adding (x int , y int) int{
	return x+y
}

func multipleReturn(x int ,y int) int {
	return x+y 
}

func basicCalculator(num1 int, num2 int)(int,int,int,float64){
	summation :=num1+num2 
	substruction :=num1-num2 
	multiplication :=num1*num2 
	division :=num1/num2 
	return summation , substruction , multiplication , float64(division)
}

func main(){
	
	myParameter("hi this is ruhul amin",25)
	test()

	fmt.Println(adding(66,251)) 
	summation , substruction , multiplication , division := basicCalculator(12,3)
	fmt.Println("\n summation  = ",summation , "\n substruction = ",substruction,"\n multiplication = ", multiplication," \n division = " ,division)

}