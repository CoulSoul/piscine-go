package piscine


import "fmt"

func DivMod(a,b int, div *int, mod *int)  {
	*div = a / b
	*mod = a % b 

	fmt.Println(*div)
	fmt.Println(*mod)
}