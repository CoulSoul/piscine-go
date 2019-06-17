package piscine 

import "fmt"

func UltimateDivMod(a *int, b *int)  {
	*a = *a / *b 
	*b = *a % *b 

	fmt.Println(*a)
	fmt.Println(*b)
}