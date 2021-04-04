package main

import "fmt"

func main(){
	statePopulation :=map[string]int{ // here all the key type is string and al
											//l the value type is integer
		"dhaka": 9555522,
		"chittagong":885552,
		"rajshahi":52222,
		"borishal":88552,

	}

	fmt.Println(statePopulation)

// decleration of a map 
  myContactNumber := make(map[string]string)
// assigning the value inside map
	myContactNumber = map[string]string{
		"grameen":"01221222144",
		"banglalink":"01403288711",
	}

	fmt.Println(myContactNumber) 

	countryList := make(map[string]string)

	countryList= map[string]string{
		"dhaka":"bangladesh",
		"noyadilli":"india",
	}
	fmt.Println("the capitals of the courtrys are : ",countryList)

	// getting the value of each key 
	fmt.Println(countryList["dhaka"])
	
	// adding a value in map 
	countryList["srilanka"] = "kualalaaampur"
	fmt.Println(countryList)

	// deletiing aa value from the map 

	delete(countryList,"srilanka")
	fmt.Println(countryList)
	// getting the length of a map 

	fmt.Println("The lenght of the map of countryList",len(countryList))
	


}