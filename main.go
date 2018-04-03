package main

import (
	"fmt"
	res "./res"
)

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