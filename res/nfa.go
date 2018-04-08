package res

import "fmt"

// State struct
type State struct {
	symbol 	rune
	edge1 	*State
	edge2 	*State
}

// Nfa struct
type Nfa struct {
	initial *State
	accept 	*State
}

// Poregtonfa function returns *nfa
// Postfix regular expression to non-deterministic finite automaton
func Poregtonfa(pofix string) *Nfa {
	// array of pointers to nfa structs
	nfastack := []*Nfa{}

	// Looping over postfix expression character by character
	for _, r := range pofix {
		switch r {
			// match any element except newline
			case '.':
				// Pop last element off the stack
				frag2 := nfastack[len(nfastack)-1]
				// Remove last element off the stack
				nfastack = nfastack[:len(nfastack)-1]
				// Pop last element off the stack
				frag1 := nfastack[len(nfastack)-1]
				// Then remove the element from the stack
				nfastack = nfastack[:len(nfastack)-1]
				
				// Join accept state of first fragment to accept state of the second fragment
				frag1.accept.edge1 = frag2.initial
				
				// Then append the new Nfa fragment to the stack
				nfastack = append(nfastack, &Nfa{initial: frag1.initial, accept: frag2.accept})

			// seperate any alternative characters
			case '|':
				frag2 := nfastack[len(nfastack)-1]
				nfastack = nfastack[:len(nfastack)-1]
				frag1 := nfastack[len(nfastack)-1]
				nfastack = nfastack[:len(nfastack)-1]

				// Points two edges to inital states of the fragments
				initial := State{edge1: frag1.initial, edge2: frag2.initial}
				accept := State{}

				// Set frag1 and frag2's edge to the accept state
				frag1.accept.edge1 = &accept
				frag2.accept.edge1 = &accept

				// Push new nfa onto the stack
				nfastack = append(nfastack, &Nfa{initial: &initial, accept: &accept})
			
			// zero or more characters
			case '*':
				// Pop one fragment off the stack
				frag := nfastack[len(nfastack)-1]
				// Remove fragment from the stack
				nfastack = nfastack[:len(nfastack)-1]

				// New accept state
				accept := State{}
				// Set inital state to frag1's inital state and the accept stage
				initial := State{edge1: frag.initial, edge2: &accept}
				
				frag.accept.edge1 = frag.initial
				frag.accept.edge2 = &accept

				// push nfa onto the stack
				nfastack = append(nfastack, &Nfa{initial: &initial, accept: &accept})
			
			// match one or more than one element 
			case '+':
				// Pop fragment off the stack
				frag := nfastack[len(nfastack)-1]
				// Remove elemt from the stack
				nfastack = nfastack[:len(nfastack)-1]
				// Set accept state
				accept := state{}
				initial := state{edge1: frag.initial, edge2: &accept}

				// set fragments accept state to inital state
				frag.accept.edge2 = &initial

				// push new nfa onto the stack
				nfastack = append(nfastack, &nfa{initial: frag.initial, accept: &accept})
			
			// match one or none elements
			case '?':
				// Pop fragment off of the stack
				frag := nfastack[len(nfastack)-1]
				// Then remove the element from the stack
				nfastack = nfastack[:len(nfastack)-1]

				initial := state{edge1: frag.initial, edge2: frag.accept}

				nfastack = append(nfastack, &nfa{initial: &initial, accept: frag.accept})
						
			default:
				accept := State{}
				initial := State{symbol: r, edge1: &accept}

				nfastack = append(nfastack, &Nfa{initial: &initial, accept: &accept})

		}
	}

	if len(nfastack) != 1 {
		fmt.Println("Sorry more than 1 nfa found", len(nfastack), nfastack)
	}

	return nfastack[0]
}

func addState(l []*State, s *State, a *State) []*State {
	l = append(l,s)

	if s != a && s.symbol == 0 {
		l = addState(l, s.edge1, a)
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}

	return l
}

// Pomatch function
func Pomatch(po string, s string) bool {
	ismatch := false
	ponfa := Poregtonfa(po)

	// current states
	current := []*State{}
	// next states
	next := []*State{}

	current = addState(current[:], ponfa.initial, ponfa.accept)

	// for every character in s loop over every character in c and compare them one by one
	for _, r := range s {
		for _, c := range current {
			// if they are equal
			if c.symbol == r {
				// if current state is equal to rune character
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}
		current, next = next, []*State{}
	}

	// if they are in the accept state then the regular expression matches the nfa
	for _, c := range current {
		if c == ponfa.accept {
			ismatch = true
			break
		}
	}

	return ismatch
}