package main

import (
	"fmt"

	"github.com/Dev-barbosajv/go-a-linguagem-do-google/pkg"
)

func main() {
	hello := pkg.Hello()
	variables := pkg.Types_variables()
	fmt.Println(hello)
	fmt.Println(variables)
}
