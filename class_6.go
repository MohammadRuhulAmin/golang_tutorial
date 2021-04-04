package main 

import "fmt" 

func main(){

	firstNum :=33
	secondNum :=55 
	if firstNum == secondNum {
		fmt.Println(firstNum , "  & ", secondNum , "  Are Same " )
	}
	if firstNum >secondNum {
		fmt.Println(firstNum ," number is greater than ",secondNum)
	}
	if firstNum < secondNum{
		fmt.Println(firstNum , " is less than  ",secondNum)
	}

	myEmail :="alfaBinomialbeta23@gmail.com"
	tryEmail :="alfaBinomialbeta23@gmail.com"
	if myEmail == tryEmail{
		fmt.Println("| mail is valid |") 
	}

	highestScore := -100 
	currentScore := -5
	msg := "Please Set your score again " 
	if currentScore<0 || currentScore>100 && highestScore>100 && highestScore<0{
		fmt.Println(msg)
	}


	if currentScore>=highestScore {

		highestScore = currentScore

	}


	guess :=100 
	if guess >100 {
		fmt.Println("grater than 100 ")
	}else if guess < 100 {
		fmt.Println("less than 100  ")
	}else  {
		fmt.Println("Equals to 100 ")
	}



}