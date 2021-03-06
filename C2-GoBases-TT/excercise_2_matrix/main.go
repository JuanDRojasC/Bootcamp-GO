package main

import (
	"errors"
	. "fmt"
	"log"
	"strings"
)

// MATRIX STRUCTURE

/* Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que represente una matriz de datos.
Para ello requieren una estructura Matrix que tenga los métodos:
Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho, si es cuadrática y cuál es el valor máximo. */

type Matrix struct {
	Values   [][]float64
	Rows     int
	Columns  int
	Square   bool
	MaxValue float64
}

func (m *Matrix) Set(values ...[]float64) error {
	m.Values = values
	m.MaxNumber()
	consistentRows := values[0]
	for _, row := range values {
		if len(row) != len(consistentRows) {
			m.Square = false
			return errors.New(Sprintf("Matrix must be have %d rows and %d columns for each rows", m.Rows, m.Columns))
		}
	}
	m.Square = true
	return nil
}

func (m Matrix) Print() {
	maxLenCell := ""
	for _, r := range m.Values {
		for _, v := range r {
			currentStrNumber := Sprintf("%.1f", v)
			if len(currentStrNumber) > len(maxLenCell) {
				maxLenCell = currentStrNumber
			}
		}
	}
	lenCell := (len(maxLenCell) + 3) * len(m.Values[0])
	rowDivision := "\n" + strings.Repeat("-", lenCell)
	Printf("MATRIX(%d, %d)", m.Rows, m.Columns)
	Println(rowDivision)
	for _, r := range m.Values {
		for _, v := range r {
			whiteSpaces := (len(maxLenCell) + 2) - len(Sprintf("%.1f", v))
			if whiteSpaces%2 == 0 {
				Printf("%s%.1f%s|", strings.Repeat(" ", whiteSpaces/2), v, strings.Repeat(" ", whiteSpaces/2))
			} else {
				Printf("%s%.1f%s|", strings.Repeat(" ", whiteSpaces/2), v, strings.Repeat(" ", (whiteSpaces/2)+1))
			}
		}
		Println(rowDivision)
	}
}

func (m *Matrix) MaxNumber() {
	maxNumber := m.Values[0][0]
	for _, r := range m.Values {
		for _, v := range r {
			if v > float64(maxNumber) {
				maxNumber = v
			}
		}
	}
	m.MaxValue = maxNumber
}

func newMatrix(rows int, columns int) Matrix {
	return Matrix{Rows: rows, Columns: columns}
}

func main() {
	values := [][]float64{{1, 2, 3, 4, 5}, {5, 4, 3, 2, 1}}
	m1 := newMatrix(len(values), len(values[0]))
	ok := m1.Set(values...)
	if ok != nil {
		log.Fatal(ok)
	}
	m1.Print()

	values2 := [][]float64{{1, 2, 3, 4, 5, 6}, {7, 8, 9, 10, 11, 12}, {5, 7, 3, 8, 6, 3}, {45, 76, 23, 58, 336, 76}}
	m2 := newMatrix(len(values2), len(values2[0]))
	ok2 := m2.Set(values2...)
	if ok != nil {
		log.Fatal(ok2)
	}
	m2.Print()
}
