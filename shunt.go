package main

import (
	"fmt"
)

func intoPost(infix string) string {
	specials := map[rune]int{'*':10, '.': 9, '|': 8}

	pofix, stack := []rune{}, []rune{}

	for _, r := range infix {
		switch {
		case r == '(':
			stack = append(stack, r)
		case r == ')':
			for stack[len(stack)-1] != '(' {
				pofix, stack = append(pofix, stack[len(stack)-1]), stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		case specials[r] > 0:
			for len(stack) > 0 && specials[r] <= specials[stack[len(stack)-1]] {
				pofix, stack = append(pofix, stack[len(stack)-1]), stack[:len(stack)-1]
			}
			stack = append(stack, r)
		default: 
			pofix = append(pofix, r)
 		}
	}

	for len(stack) > 0 {
		pofix, stack = append(pofix, stack[len(stack)-1]), stack[:len(stack)-1]
	}

	return string(pofix)
}

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