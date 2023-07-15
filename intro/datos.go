package intro

import (
	"fmt"
	"reflect"
)

func Datos() {
	var string1 string = "soy un string"
	var textGrande string = `
	Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.
	
	`
	var booleano bool = true
	var flotante32 float32 = 32.32
	var flotante64 float64 = 64.64

	numero := 8238023482834234234

	fmt.Println("====================================")
	fmt.Println(string1)
	fmt.Println(reflect.TypeOf(numero))
	fmt.Println("====================================")
	fmt.Println(textGrande)
	fmt.Println("====================================")
	fmt.Println(booleano)
	fmt.Println("====================================")
	fmt.Println(flotante32)
	fmt.Println("====================================")
	fmt.Println(flotante64)
}
