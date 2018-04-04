package main

import (
	"fmt"
	res "./res"
)

func menu() {
	fmt.Print("Enter option: \n\t1): Enter infix string to NFA conversion. \n\t2): Enter postfix string to NFA conversion. \n\t3) To exit.")
    var input int
	fmt.Scanln(&input)
	

	switch input {
	case 1: 
		var infix string
		fmt.Print("Enter infix string: ")
		fmt.Scanln(&infix)

		infixToPostfix := res.IntoPost(infix)

		var userInput string
		fmt.Print("Enter infix string: ")
		fmt.Scanln(&userInput)

		fmt.Println("Match: ", res.Pomatch(infixToPostfix, userInput))
	}
}

func main() {
	fmt.Println("Infix:		", "a.b.c*")
	fmt.Println("Postfix:		", res.IntoPost("a.b.c*"))

	fmt.Println("Infix:		", "(a.(b|d))*")
	fmt.Println("Postfix:		", res.IntoPost("(a.(b|d))*"))

	fmt.Println("Infix:		", "a.(b|d).c*")
	fmt.Println("Postfix:		", res.IntoPost("a.(b|d).c*"))

	fmt.Println("Infix:		", "a.(b.b)+.c")
	fmt.Println("Postfix:		", res.IntoPost("a.(b.b)+.c"))




}