package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Result struct {
	Type   string
	Length int
	Value  string
}

type InterfaceInput interface {
	GetValue() string
}

type Input struct {
	Value string
}

func (inp *Input) GetValue() string {
	return inp.Value
}

func SplitInputType(inp []string) (string, error) {
	aux := ""
	for i := 0; i < 2; i++ {
		aux += inp[i]
	}
	return aux, nil
}

func SplitInputLength(inp []string) (int, error) {
	aux1 := "0"
	aux2 := "0"
	for i := 2; i < 4; i++ {
		if i == 2 {
			aux1 = inp[i]
		}
		if i == 3 {
			aux2 = inp[i]
		}

	}
	intAux1, _ := strconv.ParseInt(aux1, 10, 0)
	intAux2, _ := strconv.ParseInt(aux2, 10, 0)
	return (int(intAux1*10) + int(intAux2)), nil
}

func SplitInputValue(inp []string, length int) (string, error) {
	aux := ""
	rango := 0
	for j := range inp {
		rango = j
	}
	if rango+1 == 4+length {
		for i := 4; i < 4+length; i++ {
			aux += inp[i]
		}
	}
	return aux, nil
}

func main() {
	input := Input{"T502A3"}
	split := strings.Split(input.GetValue(), "")
	resultType, errType := SplitInputType(split)
	if (errType != nil) || (resultType == "") {
		panic("se produjo un error en el tipo del argumento insertado")
	}
	resultLength, errLength := SplitInputLength(split)
	if (errLength != nil) || (resultLength == 0) {
		panic("se produjo un error en la longitud del argumento insertado")
	}
	resultValue, errValue := SplitInputValue(split, resultLength)
	if (errValue != nil) || (resultValue == "") {
		panic("se produjo un error en el valor del argumento insertado")
	}
	finalResult := Result{resultType, resultLength, resultValue}
	/*fmt.Println(split)
	fmt.Println(resultType)
	fmt.Println(resultLength)
	fmt.Println(resultValue)*/
	fmt.Println(finalResult)
}
