package main

import (
	"fmt"
	"reflect"
)

// list of field names of a struct
type Information struct {
	userName  string `required max:"100"`
	userEmail string `required max:"100"`
}

type myProfileInformation struct {
	profileName    string            `required max:"200"`
	profileEmail   string            `required max:"400"`
	profileActive  bool              `required max:"2"`
	profileFriends map[string]string `required max:"400"`
}

type Doctor struct {
	number    int
	actorName string
	company   []string
}

type Engineer struct {
	go_developer      int
	developerNameList []string
	developerContact  map[string]string
}

type birdInfo struct {
	birdName  string
	birdSpeed float64
	birdColor string
	birdSound string
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "sakib hasan",
		company: []string{
			"sakib", "sajid", "amin",
		},
	}
	fmt.Println(aDoctor)

	myEngineers := Engineer{
		go_developer:      5,
		developerNameList: []string{"ruhul", "sakib", "sajid"},
		developerContact: map[string]string{
			"sakib": "01521-433840",
			"ruhul": "0140-3288711",
			"sajid": "012255222112",
		},
	}

	fmt.Println(myEngineers)

	// now getting the value of each items of the struct
	fmt.Println(myEngineers.developerContact)
	fmt.Println(myEngineers.developerNameList)
	fmt.Println(myEngineers.go_developer)
	fmt.Println(len(myEngineers.developerContact))
	fmt.Println(len(myEngineers.developerNameList))
	fmt.Println(myEngineers.developerNameList[1])
	fmt.Println(myEngineers.developerContact["sakib"])

	newBird := birdInfo{}
	newBird.birdName = "wyka wyka"
	newBird.birdSound = "uwwikkka "
	newBird.birdSpeed = 56.32
	newBird.birdColor = "white"
	fmt.Println("the New Bird Informatin is \n", newBird)

	u := reflect.TypeOf(Information{})
	userNameField, _ := u.FieldByName("userName")
	fmt.Println(userNameField.Tag)

	userEmailField, _ := u.FieldByName("userEmail")
	fmt.Println(userEmailField.Tag)

	profileInf := reflect.TypeOf(myProfileInformation{})
	check_profileName, _ := profileInf.FieldByName("profileName")
	fmt.Println("this is for check Profile Name : ", check_profileName.Tag)

	check_profileEmail, _ := profileInf.FieldByName("profileEmail")
	fmt.Println("Profile Email", check_profileEmail.Tag)

}
