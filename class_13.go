package main 

import (
	"fmt"
)

func myName (){
	fmt.Println("Md Ruhul Amin")
}

func main(){

	// assignming a function to a variable and the variable will act like a function also 
	x := myName 
	x()

	// another way to assign a function inside variable and variable act like a function 
	test :=func(){

		fmt.Println("assigning a function inside a variable and variable act like a function ! ")
	}
	test()


	for x:=0;x<10;x++{
		func(){
			fmt.Println("this is an annonimus function call inside main function!!")
		}()
	}

	// function parameter handeling 

	age_info :=func(age int)int{
		return age*2 -1
	}(26)
	fmt.Println(age_info)


	calculator := func(num1 int, num2 int)(summation int, substruction int , multiplication int , division int){
		summation =num1+num2 
		substruction = num1- num2 
		multiplication = num1*num2 
		division =num1/num2
		return 
	}
	s,sb,m,d := calculator(12,22)
	fmt.Println("summation : " ,s)
	fmt.Println("substruction : ",sb)
	fmt.Println("multiplication : ",m)
	fmt.Println("division : ",d)




}