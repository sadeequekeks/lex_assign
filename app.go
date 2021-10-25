package main

import (
	"fmt"

	"github.com/sadeequekeks/lex_assign/functions"
	"github.com/sadeequekeks/lex_assign/inputs"
)

func main() {
	for {
		choice := inputs.GetMenu()
		switch choice {
		case "A":
			a, b := inputs.GetInt()
			ans := functions.Addit(&a, &b)
			fmt.Printf("the sum of %v and %v is %v", a, b, ans)
		case "D":
			p, q := inputs.GetFloats()
			ans2 := functions.Divis(&p, &q)
			fmt.Printf("the divison of %v and %v is %v", p, q, ans2)
		case "Q":
			fmt.Println("Thanks")
			panic("oppppssss")
		}
	}

}
