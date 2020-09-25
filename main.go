package goarea

import "math"

/*
Pi e uma proporcao numerica definida pela relacao entre
o perimetro de uma circuferencia e seu diametro
*/
const Pi = 3.1416

/*
Circ e responsavel por calcular a area de uma circuferencia
*/
func Circ(radio float64) float64 {
	return math.Pow(radio, 2) * Pi
}

/*
Rect e responsavel por calcular a area de um retangulo
*/
func Rect(base, altura float64) float64 {
	return base * altura
}

// Nao e visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
