package poo

import "fmt"

type Persona struct {
	Id       int
	Nombre   string
	Correo   string
	Edad     int
	IsWorker bool
}

func Estructura() {
	manuel := Persona{
		Id:     1,
		Nombre: "Manuel",
		Correo: "manuel@gmail.com",

		Edad:     32,
		IsWorker: true,
	}

	fmt.Println(manuel)

	catherin := new(Persona)

	catherin.Id = 1
	catherin.Nombre = "catherin"
	catherin.Correo = "catherin@gmail.com"
	catherin.Edad = 36
	catherin.IsWorker = true

	fmt.Println(catherin)

}
