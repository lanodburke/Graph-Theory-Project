# Graph-Theory-Project

Go program to build a non-deterministic finite automaton (NFA) from a regular expression and check a input string given by the user to check if it matches the NFA. This project is a college project for the Graph Theory module in GMIT.

## Research and Development

For the project I followed a set of steps set out in the project specification that are set as follows:

* Parse the regular expression from infix to postfix notation.
   * To parse the regular expression from infix to postfix I used this resource [here](http://jacobappleton.io/2015/07/02/regex-ii-the-shunting-yard-algorithm/). This article outlines how the shunting-yard algorithm works, and also how to implement it in code. This algorithm uses a stack to hold operators rather than numbers. The purpose of the stack is to reverse the order of the operators in the expression. It also serves as a storage structure, since no operator can be printed until both of its operands have appeared.

      Here is a code sample of the shunting-yard algorithm in Go.
      ```go
      func IntoPost(infix string) string {
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
      ```
  
* Build a series of small NFA’s for parts of the regular expression.
  * I used Thompson's construction algorithm to build a series of small Nfa's. Thompson's construction algorithm can be defined as           transforming a regular expression into an equivalent nondeterministic finite automaton. 
    
## Getting Started

To run the program first you will need to clone this repository onto your local machine.
```
git clone https://github.com/lanodburke/Graph-Theory-Project.git
```
Change directory into the Graph-Theory-Project folder.
```
cd Graph-Theory-Project
```
Then you will need to build the project by using the following command.
```
go build main.go
```
To run the program on Windows type this command into the command line.
```
main.exe
```
On macOS run this command.
```
./main
```

## Prerequisites

To run this program you will need to have Go installed on your machine, if you do not have it installed there is a detailed run through on how to install and setup go on your machine found [here](https://golang.org/doc/install)

## References

* [Russ Cox Article](https://swtch.com/~rsc/regexp/regexp1.html)  - Article written by Russ Cox explaing how to make Regular expression matching easy and fast
* [Thompsons Algorithm](https://www.cs.york.ac.uk/fp/lsa/lectures/REToC.pdf) - PDF describing how Thomphson's Algorithm works
* [Shunting-yard Algorithm](http://www.oxfordmathcenter.com/drupal7/node/628) - Article wrriten to describe how the shunting-yard alogirhtm works

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

