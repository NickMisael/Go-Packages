package morse

import (
	"strconv"
	"strings"
)

// alfabeto morse
var alfam = [27]string{"._", "_...", "_._.", "_..", ".", ".._.", "_..", "....", "..", ".___", "_._", "._..", "__", "_.", "___", ".__.", "__._", "._.", "...", "_", ".._", "..._", ".__", "_.._", "_.__", "__..", " "}

var alfa = [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

// Função que codifica
func Codificar(texto string) (codificado string) {
	texto = strings.ToUpper(texto)
	str := strings.Split(texto, "")
	for i := 0; i < len(str); i++ {
		for y := 0; y < len(alfa); y++ {
			if str[i] == "A" && y == 0 {
				codificado += alfam[y]
			}
			if str[i] == "B" && y == 1 {
				codificado += alfam[y]
			}
			if str[i] == "C" && y == 2 {
				codificado += alfam[y]
			}
			if str[i] == "D" && y == 3 {
				codificado += alfam[y]
			}
			if str[i] == "E" && y == 4 {
				codificado += alfam[y]
			}
			if str[i] == "F" && y == 5 {
				codificado += alfam[y]
			}
			if str[i] == "G" && y == 6 {
				codificado += alfam[y]
			}
			if str[i] == "H" && y == 7 {
				codificado += alfam[y]
			}
			if str[i] == "I" && y == 8 {
				codificado += alfam[y]
			}
			if str[i] == "J" && y == 9 {
				codificado += alfam[y]
			}
			if str[i] == "K" && y == 10 {
				codificado += alfam[y]
			}
			if str[i] == "L" && y == 11 {
				codificado += alfam[y]
			}
			if str[i] == "M" && y == 12 {
				codificado += alfam[y]
			}
			if str[i] == "N" && y == 13 {
				codificado += alfam[y]
			}
			if str[i] == "O" && y == 14 {
				codificado += alfam[y]
			}
			if str[i] == "P" && y == 15 {
				codificado += alfam[y]
			}
			if str[i] == "Q" && y == 16 {
				codificado += alfam[y]
			}
			if str[i] == "R" && y == 17 {
				codificado += alfam[y]
			}
			if str[i] == "S" && y == 18 {
				codificado += alfam[y]
			}
			if str[i] == "T" && y == 19 {
				codificado += alfam[y]
			}
			if str[i] == "U" && y == 20 {
				codificado += alfam[y]
			}
			if str[i] == "V" && y == 21 {
				codificado += alfam[y]
			}
			if str[i] == "W" && y == 22 {
				codificado += alfam[y]
			}
			if str[i] == "X" && y == 23 {
				codificado += alfam[y]
			}
			if str[i] == "Y" && y == 24 {
				codificado += alfam[y]
			}
			if str[i] == "Z" && y == 25 {
				codificado += alfam[y]
			}
		}
		codificado += " "
	}
	return
}

// Função que decodifica
func Decodificar(texto string) (traduzido string) {
	if _, err := strconv.Atoi(texto); err == nil {
		traduzido = "Não foi possivel traduzir"
		return
	}
	str := strings.Split(texto, " ")
	for i := 0; i < len(str); i++ {
		for j := 0; j < 27; j++ {
			if j == 26 && alfam[j] == str[i] {
				traduzido += " "
			}
			if j == 0 && alfam[j] == str[i] {
				traduzido += "A"
			}
			if j == 1 && alfam[j] == str[i] {
				traduzido += "B"
			}
			if j == 2 && alfam[j] == str[i] {
				traduzido += "C"
			}
			if j == 3 && alfam[j] == str[i] {
				traduzido += "D"
			}
			if j == 4 && alfam[j] == str[i] {
				traduzido += "E"
			}
			if j == 5 && alfam[j] == str[i] {
				traduzido += "F"
			}
			if j == 6 && alfam[j] == str[i] {
				traduzido += "G"
			}
			if j == 7 && alfam[j] == str[i] {
				traduzido += "H"
			}
			if j == 8 && alfam[j] == str[i] {
				traduzido += "I"
			}
			if j == 9 && alfam[j] == str[i] {
				traduzido += "J"
			}
			if j == 10 && alfam[j] == str[i] {
				traduzido += "K"
			}
			if j == 11 && alfam[j] == str[i] {
				traduzido += "L"
			}
			if j == 12 && alfam[j] == str[i] {
				traduzido += "M"
			}
			if j == 13 && alfam[j] == str[i] {
				traduzido += "N"
			}
			if j == 14 && alfam[j] == str[i] {
				traduzido += "O"
			}
			if j == 15 && alfam[j] == str[i] {
				traduzido += "P"
			}
			if j == 16 && alfam[j] == str[i] {
				traduzido += "Q"
			}
			if j == 17 && alfam[j] == str[i] {
				traduzido += "R"
			}
			if j == 18 && alfam[j] == str[i] {
				traduzido += "S"
			}
			if j == 19 && alfam[j] == str[i] {
				traduzido += "T"
			}
			if j == 20 && alfam[j] == str[i] {
				traduzido += "U"
			}
			if j == 21 && alfam[j] == str[i] {
				traduzido += "V"
			}
			if j == 22 && alfam[j] == str[i] {
				traduzido += "W"
			}
			if j == 23 && alfam[j] == str[i] {
				traduzido += "X"
			}
			if j == 24 && alfam[j] == str[i] {
				traduzido += "Y"
			}
			if j == 25 && alfam[j] == str[i] {
				traduzido += "Z"
			}
		}
	}
	return
}
