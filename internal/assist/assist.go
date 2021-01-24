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
)

type (
	Assist interface {
		DecomposeSavings(goal, interest float32, years uint8) (float32, error)
		Validator() Validator
		View() View
	}
	assist struct {
		validator Validator
		view      View
	}
)

var _ Assist = assist{}

func New() Assist {
	return assist{
		view:      view{},
		validator: validator{},
	}
}

func (a assist) DecomposeSavings(goal, interest float32, years uint8) (payment float32, err error) {
	if err = a.validator.HumanLifeYears(years); err != nil {
		return payment, err
	}
	if err = a.validator.PositiveFloat32(interest); err != nil {
		return payment, fmt.Errorf("invalid interest rate: %w", err)
	}
	if err = a.validator.PositiveFloat32(goal); err != nil {
		return payment, fmt.Errorf("invalid financial goal: %w", err)
	}

	periodRate := interest * 0.01 / 12
	coefficient := 1 + periodRate

	finalCoefficient := coefficient
	for i := 1; i < 12*int(years); i++ {
		finalCoefficient *= coefficient
	}

	// Формула сложных процентов, начисляемых несколько раз в течение года,
	// выходит из суммы геометрической прогрессии, в которой первый член
	// равен payment*(1+periodRate), а знаменатель прогрессии - (1+periodRate).
	payment = (goal * periodRate) / (coefficient*finalCoefficient - coefficient)

	return
}

func (a assist) Validator() Validator {
	return a.validator
}

func (a assist) View() View {
	return a.view
}
