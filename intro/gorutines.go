package intro

import (
	"fmt"
	"time"
)

func GoRutines() {
	fmt.Println(Saluda("manuel"))
	time.Sleep(time.Second * 5)
	fmt.Println(Saluda("pedro"))
	fmt.Println("despues de la siesta")
	miCanal := make(chan string)
	go func() {
		miCanal <- Saluda("gorutine")
	}()
	fmt.Println(<-miCanal)
	fmt.Println("continuar con la ejecucion")
}

func Saluda(nombre string) string {
	return ("hola " + nombre)

}
