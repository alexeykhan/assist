// Copyright (c) 2021 Alexey Khan
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package check

import (
	"fmt"
)

type (
	// Checker осуществляет проверку входных данных.
	Checker interface {
		HumanLifeYears(years uint8) error
		PositiveFloat64(value float64) error
	}
	check struct{}
)

// New создает новый экземпляр Checker.
func New() Checker {
	return check{}
}

// HumanLifeYears проверяет входное число на соразмерность длительности человеческой жизни.
func (c check) HumanLifeYears(years uint8) error {
	if years == 0 || years > 100 {
		return fmt.Errorf("meaningless number of years: got %d, expected from 1 to 100", years)
	}
	return nil
}

// PositiveFloat64 проверяет, что входное число является неотрицательным float64.
func (c check) PositiveFloat64(value float64) error {
	if value <= 0 {
		return fmt.Errorf("expected positive value: %.2f", value)
	}
	return nil
}
