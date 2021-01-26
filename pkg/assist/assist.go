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

package assist

import (
	"fmt"
	"math"

	"github.com/alexeykhan/assist/pkg/check"
)

type (
	// Assist описывает функционал для проведения финансовых расчетов.
	Assist interface {
		DecomposeSavings(goal, interest float64, years uint8) (float64, error)
		DecomposeRetirement(expenses, interest float64, years uint8) (float64, error)
		CalculateSavings(payment, interest float64, years uint8) (float64, error)
		CalculateInflation(current, inflation float64, years uint8) (float64, error)
	}
	assist struct {
		check check.Checker
	}
)

var _ Assist = assist{}

// New создает новый экземпрляр Assist.
func New() Assist {
	return assist{
		check: check.New(),
	}
}

// DecomposeSavings рассчитывает ежемесячный платеж, необходимый для накопления суммы goal,
// при условии поддержания средней доходности портфеля не ниже interest% на протяжении years лет.
func (a assist) DecomposeSavings(goal, interest float64, years uint8) (payment float64, err error) {
	if err = a.check.HumanLifeYears(years); err != nil {
		return payment, err
	}
	if err = a.check.PositiveFloat64(interest); err != nil {
		return payment, fmt.Errorf("invalid interest rate: %w", err)
	}
	if err = a.check.PositiveFloat64(goal); err != nil {
		return payment, fmt.Errorf("invalid financial goal: %w", err)
	}

	periodRate := interest * 0.01 / 12
	coefficient := 1 + periodRate

	finalCoefficient := coefficient
	for i := 1; i < 12*int(years); i++ {
		finalCoefficient *= coefficient
	}

	payment = (goal * periodRate) / (coefficient*finalCoefficient - coefficient)
	return
}

// DecomposeRetirement рассчитывает минимальную суммы накоплений, которая позволит выйти на
// пенсию и на протяжении years лет ежемесячно тратить сумму expenses, при условии, что накопления
// будут приносить доходность не менее interest% годовых.
func (a assist) DecomposeRetirement(expenses, interest float64, years uint8) (retirement float64, err error) {
	if err = a.check.HumanLifeYears(years); err != nil {
		return retirement, err
	}
	if err = a.check.PositiveFloat64(interest); err != nil {
		return retirement, fmt.Errorf("invalid interest rate: %w", err)
	}
	if err = a.check.PositiveFloat64(expenses); err != nil {
		return retirement, fmt.Errorf("invalid expenses: %w", err)
	}

	R := interest * 0.01 / 12
	RPlusOnePow := math.Pow(R+1, 12*float64(years)-1)
	retirement = expenses * (RPlusOnePow*(R+1) - 1) / (RPlusOnePow * R)
	return
}

// CalculateSavings рассчитывает будущие накопления при условии ежемесячных инвестиций на сумму
// не менее payment, средней доходностью портфеля interest% и ежемесячной капитализацией процентов
// на протяжении years лет.
func (a assist) CalculateSavings(payment, interest float64, years uint8) (savings float64, err error) {
	if err = a.check.HumanLifeYears(years); err != nil {
		return savings, err
	}
	if err = a.check.PositiveFloat64(interest); err != nil {
		return savings, fmt.Errorf("invalid interest rate: %w", err)
	}
	if err = a.check.PositiveFloat64(payment); err != nil {
		return savings, fmt.Errorf("invalid payment: %w", err)
	}

	periodRate := interest * 0.01 / 12
	coefficient := 1 + periodRate

	finalCoefficient := coefficient
	for i := 1; i < 12*int(years); i++ {
		finalCoefficient *= coefficient
	}

	savings = payment * (coefficient*finalCoefficient - coefficient) / periodRate
	return
}

// CalculateInflation рассчитывает изменение покупательской способности относительно заданной
// суммы current по прошествии years лет и средним показателем инфляции inflation% в год,
// возвращая эквивалент исходной суммы current по прошествии срока years.
func (a assist) CalculateInflation(current, inflation float64, years uint8) (inflated float64, err error) {
	if err = a.check.HumanLifeYears(years); err != nil {
		return inflated, err
	}
	if err = a.check.PositiveFloat64(current); err != nil {
		return inflated, fmt.Errorf("invalid current value: %w", err)
	}
	if err = a.check.PositiveFloat64(inflation); err != nil {
		return inflated, fmt.Errorf("invalid inflation: %w", err)
	}

	inflated = current * math.Pow(1+inflation*0.01, float64(years))
	return
}
