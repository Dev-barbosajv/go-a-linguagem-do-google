package pkg

import "fmt"

func Types_variables() string {
	var name string = "Joao Vitor"
	var idade int = 23
	var version float32 = 1.1

	message := fmt.Sprintf("Olá Sr %s, que legal saber que tem %d anos.\nVersion system: %.2f", name, idade, version)
	return message

}
