package modulos

import (
	"fmt"
	"strings"
)

func Strings() {
	var texto string = "esto es un texto de prueba"

	fmt.Println(strings.ToUpper(texto))
	fmt.Println(strings.ToLower(texto))
	fmt.Println(len(texto))
	fmt.Println(strings.Split(texto, ""))
	fmt.Println(len(strings.Split(texto, "")))
}
