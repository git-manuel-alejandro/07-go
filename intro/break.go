package intro

import "fmt"

func Break() {

	var pais string = "chileasd"

	switch {
	case pais == "chile":
		fmt.Println("chileno")
	case pais == "argentina":
		fmt.Println("argentino")
	case pais == "colombia":
		fmt.Println("chcolombiano")
	case pais == "usa":
		fmt.Println("american")
	default:
		fmt.Println("another thing")
	}

}
