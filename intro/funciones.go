package intro

func Funciones(n1, n2 int) int {

	return n1 + n2

}

var suma = func(n1, n2 int) int {

	return n1 + n2
}

func FuncionAnonima() int {

	return suma(10, 20)

}
