package main

import (
	"fmt"
	"strings"
)

func Codifica(texto string) (codificado string) {
	str := strings.Split(texto, " ")
	fmt.Printf("%v\n", str[0])
	return
}

func Decodifica(texto string) (traduzido string) {
	str := strings.Split(texto, " ")
	fmt.Printf("%v\n", str[0])
	return
}

func main() {
	fmt.Println(Codifica("A B C"))
}
