package tiempo

import (
	"fmt"
	"time"
)

func Tiempo() {

	var now = time.Now()
	fmt.Println("hora actual ", now)
	fmt.Println("año actual ", now.Year())
	fmt.Println("mes actual ", now.Month())
	fmt.Println("dia actual ", now.Day()-10)

	fmt.Printf("%v-0%v-%v", now.Day(), int(now.Month()), now.Year())
}
