package main

import (
	"fmt"	
)



func main(){

	switch 4 {
		case 1 :
			fmt.Println("one is printed")
		case 2:
			fmt.Println("two is printed")
		case 3:
			fmt.Println("three is printed") 
		case 4:
			fmt.Println("four is printed")
		default:
			fmt.Println("not in range ! ")
	}
	increment := 10 
	switch {
	case increment<10 :
		fmt.Println("increment is less than 10")
	case increment==10:
		fmt.Println("increment is equal to 10")
	case increment>10:
		fmt.Println("increment is grater than 10")
	default:
		fmt.Println("error!")
	}

   var mySuperHero interface{} = 52.33
   fmt.Println(mySuperHero)

   var myObject interface{}  = "dhaka" //  this is anytype of data 
   fmt.Println(myObject)

   switch myObject.(type) {
	case int:
		fmt.Println("this is integer")
	case float32:
		fmt.Println("this is float data type")
	case string:
		fmt.Println("this data type is string")
   }

  


}