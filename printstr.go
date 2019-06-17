package piscine 

import "fmt"

func PrintStr(str string)  {
	
	for i := 0; i < len(str); i++{
		fmt.Println(str[i])
	}

}