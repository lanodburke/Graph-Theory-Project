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
func Poregtonfa(pofix string) *Nfa {
	// array of pointers to nfa structs
	nfastack := []*Nfa{}

	for _, r := range pofix {
		switch r {
			case '.':
				frag2 := nfastack[len(nfastack)-1]
				nfastack = nfastack[:len(nfastack)-1]
				frag1 := nfastack[len(nfastack)-1]
				nfastack = nfastack[:len(nfastack)-1]

				frag1.accept.edge1 = frag2.initial

				nfastack = append(nfastack, &Nfa{initial: frag1.initial, accept: frag2.accept})

			case '|':
				frag2 := nfastack[len(nfastack)-1]
				nfastack = nfastack[:len(nfastack)-1]
				frag1 := nfastack[len(nfastack)-1]
				nfastack = nfastack[:len(nfastack)-1]

				initial := State{edge1: frag1.initial, edge2: frag2.initial}
				accept := State{}

				frag1.accept.edge1 = &accept
				frag2.accept.edge1 = &accept

				nfastack = append(nfastack, &Nfa{initial: &initial, accept: &accept})
			
			case '*':
				frag := nfastack[len(nfastack)-1]
				nfastack = nfastack[:len(nfastack)-1]

				accept := State{}
				initial := State{edge1: frag.initial, edge2: &accept}
				
				frag.accept.edge1 = frag.initial
				frag.accept.edge2 = &accept

				nfastack = append(nfastack, &Nfa{initial: &initial, accept: &accept})
				case '+':
					frag := nfastack[len(nfastack)-1]

					accept := state{}
					initial := state{edge1: frag.initial, edge2: &accept}

					frag.accept.edge1 = &initial

					nfastack = append(nfastack, &nfa{initial: frag.initial, accept: &accept})
				case '?':
					frag := nfastack[len(nfastack)-1]
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

	current := []*State{}
	next := []*State{}

	current = addState(current[:], ponfa.initial, ponfa.accept)

	for _, r := range s {
		for _, c := range current {
			if c.symbol == r {
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}
		current, next = next, []*State{}
	}

	for _, c := range current {
		if c == ponfa.accept {
			ismatch = true
			break
		}
	}

	return ismatch
}