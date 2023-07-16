package intro

import "fmt"

func Arrays() {
	var arreglo [10]int

	for i := 0; i < len(arreglo); i++ {
		arreglo[i] = i + 1

		fmt.Println(arreglo)
	}
}

func Slice() {

	var slice = make([]int, 0)

	for i := 1; i <= 10; i++ {

		slice = append(slice, i)

	}

	fmt.Println(slice)

	slice = append(slice[:5], slice[5+1])

	fmt.Println(slice)
}
