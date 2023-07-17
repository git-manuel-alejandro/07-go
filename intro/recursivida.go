package intro

import "fmt"

func Recursion(valor int) {

	contador := valor + 1
	fmt.Println(contador)
	if contador < 10 {
		Recursion(contador)
	}
}
