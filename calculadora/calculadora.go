package main

type Calculadora struct {
	opcion int
	valor1 int
	valor2 int
}

func (o *Calculadora) HacerOperacion() int {
	switch o.opcion {
	case 1:
		return o.Sumar()
	case 2:
		return o.Restar()
	case 3:
		return o.Multiplicar()
	case 4:
		return o.Dividir()
	default:
		return 0
	}
}

func (o *Calculadora) Sumar() int {
	return o.valor1 + o.valor2
}

func (o *Calculadora) Restar() int {
	return o.valor1 - o.valor2
}

func (o *Calculadora) Multiplicar() int {
	return o.valor1 * o.valor2
}

func (o *Calculadora) Dividir() int {
	return o.valor1 / o.valor2
}
