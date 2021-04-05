package main

import "fmt"
 
func main() {
    g:= func(){
		fmt.Println("function call G")
	}
	e:=func(){
		fmt.Println("function call E")
	}
	t:=func(){
		fmt.Println("function T")
	}

	for x:=0;x<10;x++{
		func(){
			g()
			e()
			t()
		}()
	}



	f := func(){
		fmt.Println("anonimus function another way to call ")
	}
	f()
	fmt.Println("anonimus function / return values / Functions as types / Methods ")
	var myMsg string ="this is Message : "
	sayMessage("hi world")
	myEmailInfo("alfaBinomialbeta23@gmail.com")
	for i:=0;i<10;i++{
		validationList(myMsg,i) 
	}
	summation(1,2,3,4,5)
	fmt.Println(multiplication(1,2,3,4,5))
	fmt.Println(calculation(1,2,3,4,5,6,7,8,9,10))

	func (){
		fmt.Println("this annonimus function call")
	}()

	fmt.Println("---------------------------------------")
	for x:=0;x<10;x++{
		func(){
			fmt.Println("hi this is annonimus function call ", x)
		}()
	}
	fmt.Println("-------------------------------------")
	for l:=0;l<10;l++{
		func(l int){
			fmt.Println(l)
		}(l)
	}
 
}



func sayMessage (msg string){
	fmt.Println(msg)
}

func myEmailInfo(email string){
	fmt.Println(email)
}

func validationList(msg string,index int){
	fmt.Println(msg, " " , index )
}

func summation (values ...int){
	result:=0
	for _,v:=range values{
		result +=v
	}
	fmt.Println("The summation is :: ",result)
}

// function return statements 

func multiplication (values ...int) int{

	result :=1
	for _, v :=range values{
		result *=v 
	}
	return result
	
}

func calculation(values ...int) int{
	result :=0  
	for _,v:=range values{
		result +=v 
	}
	return result
}
