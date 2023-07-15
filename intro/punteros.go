package intro

import "fmt"

func Punteros() {

	var color string = "red"
	var puntero = &color

	fmt.Println(&color)
	fmt.Println(*puntero)

}
