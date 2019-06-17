package piscine 

import "fmt"

func UltimateDivMod(a *int, b *int)  {

	y := *a

	*a = *a / *b 
	*b = y % *b 

	fmt.Println(*a)

	fmt.Println(*b)

	
}