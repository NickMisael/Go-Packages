package ascii

import (
	"errors"
	"fmt"
)

func Code(texto string) (code string) {
	for _, item := range texto {
		code += fmt.Sprintf("%X", item)
	}
	return
}

func CodeS(textos []string) (code []string) {
	for _, txt := range textos {
		var str string
		for _, char := range txt {
			str += fmt.Sprintf("%X", char)
		}
		code = append(code, str)
	}
	return
}

func Decode(texto string) string {
	var str string
	for _, txt := range texto {
		str += fmt.Sprintf("%c", txt)
	}
	return str
}

func DecodeS(text []byte) (string, error) {
	if len(text) == 0 {
		return "", errors.New("Nenhum byte encontrado! :(")
	}
	var tr string
	for i := 0; i < len(text); i++ {
		tr += fmt.Sprintf("%c", text[i])
	}
	return tr, nil
}
