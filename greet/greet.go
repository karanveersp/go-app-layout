package greet

import "fmt"

func SayHello(lang, name string) {
	if lang == "spanish" {
		fmt.Println("Hola", name)
	} else {
		fmt.Println("Hello", name)
	}
}
