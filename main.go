package main

import (
	"fmt"
)

type App struct {
}

func main() {
	//if os.Args[1] == "help" {
	//		fmt.Println("MGC - Mint Graphics calc commands : \n 1. help - show help info \n 2. version - show version of program \n 3. save - saving your last graphic")
	//	} else if os.Args[1] == "version" {
	//		fmt.Println("MGC version : 0.1.0")
	//	}

	mp := MathParser{}

	form := "10*x + 11"
	tokens := MakeTokens(form)

	y := mp.Eval(tokens, float64(5))
	fmt.Printf("%.2f %.2f\n", float64(5), y)
}
