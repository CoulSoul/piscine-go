package piscine

import "fmt"

func PrintComb()  {
	for i:= {
		for j := 1; j:= 9; j++{
			fmt.Printf(j)
			for k := 3; k :=9; k++{ 
				fmt.Printf(k)
			}
		}
	}
}