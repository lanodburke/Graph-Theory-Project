package res

// IntoPost function
func IntoPost(infix string) string {
	// Map of speacial characters mapped to numbers, higher the number higher the presidence of the operator
	specials := map[rune]int{'*':10, '.': 9, '|': 8, '?': 7, '+': 6}

	// Array of rune's, rune's are any unicode character
	// Postfix will be the expression re-written after the loop has finished in post fix notation
	// Stack is where operators are stored in the infix expression
	pofix, stack := []rune{}, []rune{} 

	// Loop over infix expression character by character
	for _, r := range infix {
		switch {
		case r == '(':
			// append onto the end of the stack
			stack = append(stack, r)
		case r == ')':
			// pop characters off the stack until we see an open bracket and append them onto postfix expression
			for stack[len(stack)-1] != '(' {
				pofix, stack = append(pofix, stack[len(stack)-1]), stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		case specials[r] > 0:
			// if there are elements on the stack and the presidence of the current character 
			// is less than the presidence of the character on the top of the stack, pop elements 
			// off the stack and append them to the postfix expression 
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