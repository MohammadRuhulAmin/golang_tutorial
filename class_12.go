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



// multiple return parameter and lable return value example .... 


func advancedCalculator(num1 int , num2 int )(summation int,substruction int , multiplication int , division int){
	summation  = num1 + num2 
	substruction = num1 - num2 
	multiplication = num1*num2 
	division = num1/num2
	return  
}

func main(){
	
	myParameter("hi this is ruhul amin",25)
	test()

	fmt.Println("Basic Calculator ")
	fmt.Println(adding(66,251)) 
	summation , substruction , multiplication , division := basicCalculator(12,3)
	fmt.Println("\n summation  = ",summation , "\n substruction = ",substruction,"\n multiplication = ", multiplication," \n division = " ,division)

	fmt.Println("Advanced Calculator ... ")
	av_sum,ad_sub,ad_mul,ad_div := advancedCalculator(12,4)
	fmt.Println("\nsummation ",av_sum ,"\n substruction ",ad_sub,"\n multiplication ",ad_mul,"\n division ",ad_div)


}