package main 

import "fmt"

// list of field names of a struct 

type Doctor struct{
	number int 
	actorName string 
	company []string
}

type Engineer struct{
	go_developer int 
	developerNameList []string 
	developerContact map[string]string
}

func main(){
	aDoctor :=Doctor{
		number:3,
		actorName:"sakib hasan",
		company:[]string{
			"sakib","sajid","amin",
		},
	}
	fmt.Println(aDoctor)

	
myEngineers := Engineer{
	go_developer : 5,
	developerNameList:[]string{"ruhul","sakib","sajid"},
	developerContact: map[string]string{
		"sakib":"01521-433840",
		"ruhul":"0140-3288711",
		"sajid":"012255222112",
	},
}

	fmt.Println(myEngineers)

	// now getting the value of each items of the struct 
	fmt.Println(myEngineers.developerContact)
	fmt.Println(myEngineers.developerNameList)
	fmt.Println(myEngineers.go_developer)
	fmt.Println(len(myEngineers.developerContact))
	fmt.Println(len(myEngineers.developerNameList))
	
}


