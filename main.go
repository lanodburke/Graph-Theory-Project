package main

import (
	"fmt"
	res "./res"
)

func menu() {
	fmt.Print("Enter option: \n\t1): Enter infix string to NFA conversion. \n\t2): Enter postfix string to NFA conversion. \n\t3): To exit.\n\n\tInput here: ")
    var input int
	fmt.Scanln(&input)
	
	switch input {
	case 1: 
		var infix string
		fmt.Print("Enter infix string: ")
		fmt.Scanln(&infix)

		infixToPostfix := res.IntoPost(infix)

		var userInput string
		fmt.Print("Enter string to test against NFA: ")
		fmt.Scanln(&userInput)

		fmt.Println("Match: ", res.Pomatch(infixToPostfix, userInput))

	case 2: 
		var postfix string
		fmt.Print("Enter postfix string: ")
		fmt.Scanln(&postfix)

		var userInput string
		fmt.Print("Enter string to test against NFA: ")
		fmt.Scanln(&userInput)

		fmt.Println("Match: ", res.Pomatch(postfix, userInput))
	}
	
}

func main() {
	menu()
}