package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Printf("x")
		}
		println()
	}
	// another syntax of for loop

	println()

	x := 0
	for x <= 10 {
		println("x is on fire! ")
		x++
	}

	// another way to write for loop

	inc := 1
	for {
		println("Print _ of _  increment ", inc)
		inc++
		if inc == 10 {
			break
		}
	}

	// continue statement with for loop

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			continue
		} else {
			fmt.Printf("%v ", i)
		}
	}

	// array value fetching both !
	solution := []int{1, 2, 3, 4}
	for k, v := range solution {
		fmt.Println(k, " -- ", v)
	}

	myArray := []string {"ruhul","sakib", "sajid","affan","ammar"}
	for k,v:=range myArray{
		fmt.Println(k," ",v)
	}
	// loop over maps 

	 myStates := map[string]string {
		"dhaka":"2200112",
		"rajshahi":"225222",
		"khulna" :"77445512",
		"barisal":"885521221",
	}

	for k,v:=range myStates{
		fmt.Println(k, " | " ,v)
	} 
	// what if i just want to get value instead of key
	for _,v:=range myStates{
		fmt.Println( v)
	}

	// what if i just want to get key instead of key
	for k,_:=range myStates{
		fmt.Println(k)
	}





}
