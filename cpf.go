package cpf

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func Valida(cf string) (ver bool) {
	var soma int
	/*if _, err := strconv.Atoi(cf); err != nil {
		return
	}*/
	cp := strings.Split(cf, "")
	if len(cf) < 11 || len(cf) > 11 {
		return
	}
	if cf == "" || cf == "12345678901" || cf == "00000000000" || cf == "11111111111" || cf == "22222222222" || cf == "33333333333" || cf == "44444444444" || cf == "55555555555" || cf == "66666666666" || cf == "77777777777" || cf == "88888888888" || cf == "99999999999" {
		return
	}
	dig1 := [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	for i := 0; i < 9; i++ {
		num, _ := strconv.Atoi(cp[i])
		soma += num * dig1[i]
	}
	cpf := (soma * 10) % 11
	if cpf == 10 {
		cpf = 0
	}
	cpp, _ := strconv.Atoi(cp[len(cp)-2])
	if cpp != cpf {
		return
	}
	soma = 0
	for i := 0; i < 10; i++ {
		num, _ := strconv.Atoi(cp[i])
		soma += num * (dig1[i] + 1)
	}
	cpf = (soma * 10) % 11
	if cpf == 10 {
		cpf = 0
	}
	cpp, _ = strconv.Atoi(cp[len(cp)-1])
	if cpp != cpf {
		return
	}
	ver = true
	return
}

func Gera1() (cpf string) {
	rand.Seed(time.Now().UTC().UnixNano())
	for {
		num := rand.Intn(9 - 0)
		cpf += strconv.Itoa(num)
		if len(cpf) == 11 {
			if Valida(cpf) == false {
				cpf = ""
			} else {
				break
			}
		}
	}
	return
}

func GeraS(qtd int) []string {
	rand.Seed(time.Now().UTC().UnixNano())
	cpf := make([]string, qtd)
	for i := 0; i < qtd; i++ {
		for {
			num := rand.Intn(9 - 0)
			cpf[i] += strconv.Itoa(num)
			if len(cpf[i]) == 11 {
				if Valida(cpf[i]) == false {
					cpf[i] = ""
				} else {
					break
				}
			}
		}
	}
	return cpf
}

/*
func main() {
	/*var cpf string
	fmt.Printf("CPF -> ")
	fmt.Scanf("%s", &cpf)
	if Valida(cpf) == false {
		fmt.Println("CPF Inválido!")
	} else {
		fmt.Println("CPF Válido")
	}
	fmt.Println(GeraS(3))
}*/
