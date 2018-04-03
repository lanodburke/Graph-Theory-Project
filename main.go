package main

import (
	"fmt"
)

func main() {
	fmt.Println("Infix:		", "a.b.c*")
	fmt.Println("Postfix:		", intoPost("a.b.c*"))

	fmt.Println("Infix:		", "(a.(b|d))*")
	fmt.Println("Postfix:		", intoPost("(a.(b|d))*"))

	fmt.Println("Infix:		", "a.(b|d).c*")
	fmt.Println("Postfix:		", intoPost("a.(b|d).c*"))

	fmt.Println("Infix:		", "a.(b.b)+.c")
	fmt.Println("Postfix:		", intoPost("a.(b.b)+.c"))


}