package inputs

import (
	"fmt"
	"strconv"
)

func GetInt() (int, int) {
	var a, b string
	fmt.Print("Enter the first number pls:")
	fmt.Scanln(&a)
	fmt.Print("Enter the second number pls:")
	fmt.Scanln(&b)
	x, _ := strconv.Atoi(a)
	y, _ := strconv.Atoi(b)
	return x, y
}

func GetFloats() (x float64, y float64) {
	var a, b string
	fmt.Print("Enter the first number pls:")
	fmt.Scanln(&a)
	fmt.Print("Enter the second number pls:")
	fmt.Scanln(&b)
	x, _ = strconv.ParseFloat(a, 64)
	y, _ = strconv.ParseFloat(b, 64)
	return
}

func GetMenu() string {
	var a string
	fmt.Println("| -------- WELCOME TO LEX CLASS CALC ------|")
	fmt.Println("| [A] - Add                                |")
	fmt.Println("| [S] - Subtarct                           |")
	fmt.Println("| [M] - Multiply                           |")
	fmt.Println("| [D] - Divide                             |")
	fmt.Println("| [Q] - Quit                               |")
	fmt.Println("| -----------------------------------------|")
	fmt.Scanln(&a)
	return a
}
