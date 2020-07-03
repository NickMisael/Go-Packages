package morse

import (
	"strconv"
	"strings"
)

var alfa = [27]string{"._", "_...", "_._.", "_..", ".", ".._.", "_..", "....", "..", ".___", "_._", "._..", "__", "_.", "___", ".__.", "__._", "._.", "...", "_", ".._", "..._", ".__", "_.._", "_.__", "__..", " "}

/*m["A"] = "._"
m["B"] = "_..."
m["C"] = "_._."
m["D"] = "_.."
m["E"] = "."
m["F"] = ".._."
m["G"] = "_.."
m["H"] = "...."
m["I"] = ".."
m["J"] = ".___"
m["K"] = "_._"
m["L"] = "._.."
m["M"] = "__"
m["N"] = "_."
m["O"] = "___"
m["P"] = ".__."
m["Q"] = "__._"
m["R"] = "._."
m["S"] = "..."
m["T"] = "_"
m["U"] = ".._"
m["V"] = "..._"
m["W"] = ".__"
m["X"] = "_.._"
m["Y"] = "_.__"
m["Z"] = "__.."
m[" "] = " "*/
func Codifica(texto string) (codificado string) {
	if _, err := strconv.Atoi(texto); err == nil {
		codificado = "Não foi possivel traduzir"
		return
	}
	for i := 0; i < len(str); i++ {
	}
}

func Decodifica(texto string) (traduzido string) {
	if _, err := strconv.Atoi(texto); err == nil {
		traduzido = "Não foi possivel traduzir"
		return
	}
	str := strings.Split(texto, " ")
	for i := 0; i < len(str); i++ {
		for j := 0; j < 27; j++ {
			if j == 26 && alfa[j] == str[i] {
				traduzido += " "
			}
			if j == 0 && alfa[j] == str[i] {
				traduzido += "A"
			}
			if j == 1 && alfa[j] == str[i] {
				traduzido += "B"
			}
			if j == 2 && alfa[j] == str[i] {
				traduzido += "C"
			}
			if j == 3 && alfa[j] == str[i] {
				traduzido += "D"
			}
			if j == 4 && alfa[j] == str[i] {
				traduzido += "E"
			}
			if j == 5 && alfa[j] == str[i] {
				traduzido += "F"
			}
			if j == 6 && alfa[j] == str[i] {
				traduzido += "G"
			}
			if j == 7 && alfa[j] == str[i] {
				traduzido += "H"
			}
			if j == 8 && alfa[j] == str[i] {
				traduzido += "I"
			}
			if j == 9 && alfa[j] == str[i] {
				traduzido += "J"
			}
			if j == 10 && alfa[j] == str[i] {
				traduzido += "K"
			}
			if j == 11 && alfa[j] == str[i] {
				traduzido += "L"
			}
			if j == 12 && alfa[j] == str[i] {
				traduzido += "M"
			}
			if j == 13 && alfa[j] == str[i] {
				traduzido += "N"
			}
			if j == 14 && alfa[j] == str[i] {
				traduzido += "O"
			}
			if j == 15 && alfa[j] == str[i] {
				traduzido += "P"
			}
			if j == 16 && alfa[j] == str[i] {
				traduzido += "Q"
			}
			if j == 17 && alfa[j] == str[i] {
				traduzido += "R"
			}
			if j == 18 && alfa[j] == str[i] {
				traduzido += "S"
			}
			if j == 19 && alfa[j] == str[i] {
				traduzido += "T"
			}
			if j == 20 && alfa[j] == str[i] {
				traduzido += "U"
			}
			if j == 21 && alfa[j] == str[i] {
				traduzido += "V"
			}
			if j == 22 && alfa[j] == str[i] {
				traduzido += "W"
			}
			if j == 23 && alfa[j] == str[i] {
				traduzido += "X"
			}
			if j == 24 && alfa[j] == str[i] {
				traduzido += "Y"
			}
			if j == 25 && alfa[j] == str[i] {
				traduzido += "Z"
			}
		}
	}
	return
}

/*
func main() {
	fmt.Println(Decodifica("_.__ .. _."))
}*/
