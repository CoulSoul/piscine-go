package main


import "github.com/01-edu/z01"

func IsNegative(nb int) {
	var nb int = 4

	if nb < 0 {
		z01.PrintRune("T")
	}else{
		z01.PrintRune("F")
	}
}