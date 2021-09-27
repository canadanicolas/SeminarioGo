package model

type Result struct {
	Tipo   string
	Length int
	Value  string
}

type InterfaceInput interface {
	GetValue() string
}

type Input struct {
	Value string
}

func NewInput(value string) Input {
	return Input{value}
}

func NewResult(tipo string, length int, value string) Result {
	return Result{tipo, length, value}
}
