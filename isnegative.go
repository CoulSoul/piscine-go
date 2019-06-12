package main


import "fmt"

func IsNegative() {
	var test int = 4

	if test < 0 {
		fmt.Println("T")
	}else{
		fmt.Println("F")
	}
}