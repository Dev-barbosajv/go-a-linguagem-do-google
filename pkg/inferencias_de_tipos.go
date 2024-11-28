package pkg

import (
	"fmt"
	"reflect"
)

func Inferencias_de_tipos() string {
	var name = "admin"
	var idade = 20

	message := fmt.Sprintf("Olá sr %s, sua idade é %d", name, idade)
	fmt.Println("Tipo da variavel:", reflect.TypeOf(name))
	return message
}
