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

package main

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var decomposeSavingsFlags = struct {
	YearsLeft     pflag.Flag
	InterestRate  pflag.Flag
	FinancialGoal pflag.Flag
	Capitalize    pflag.Flag
	Help          pflag.Flag
}{
	YearsLeft: pflag.Flag{
		Name:      "years",
		Shorthand: "y",
		Usage:     "Количество лет, за которое необходимо накопить нужную сумму",
		DefValue:  "",
	},
	InterestRate: pflag.Flag{
		Name:      "interest",
		Shorthand: "i",
		Usage:     "Доходность вашего инвестиционного портфеля в процентах годовых",
		DefValue:  "",
	},
	FinancialGoal: pflag.Flag{
		Name:      "goal",
		Shorthand: "g",
		Usage:     "Ваша финансовая цель, которую нужно достгнуть за заданный период",
		DefValue:  "",
	},
	Capitalize: pflag.Flag{
		Name:      "capitalize",
		Shorthand: "c",
		Usage:     "Включать капитализацию процентов или нет",
		DefValue:  "",
	},
	Help: pflag.Flag{
		Name:      "help",
		Shorthand: "h",
		Usage:     "Документация по команде",
		DefValue:  "",
	},
}

var decomposeSavings = &cobra.Command{
	Use:   "savings",
	Short: "Рассчитать, сколько денег будет по итогу периоду",
	Example: example(
		"Декомпозиция финансовой цели",
		"Узнайте, сколько денег необходимо инвестировать каждый месяц, квартал, "+
			"полгода или год, чтобы успешно накопить необходимую сумму за обозначенный "+
			"срок в годах с учетом капиталлизации процентов или без.",
		[]string{
			"./oracle decompose savings --goal=1234567.89 --years=10 --interest=6.5",
			"./oracle decompose savings -g=1234567.89 -y=10 -i=6.5 -c=false",
			"./oracle decompose savings --help",
			"./oracle decompose savings -h",
		},
	),
	Run: func(cmd *cobra.Command, args []string) {
		yearsLeft := getUint8(cmd, decomposeSavingsFlags.YearsLeft.Name)
		annualRate := getFloat32(cmd, decomposeSavingsFlags.InterestRate.Name)
		financialGoal := getFloat32(cmd, decomposeSavingsFlags.FinancialGoal.Name)
		capitalize := getBool(cmd, decomposeSavingsFlags.Capitalize.Name)

		capitalizeInfo := "выключена"
		if capitalize {
			capitalizeInfo = "включена"
		}

		var yearsInfo string
		remainder := yearsLeft / 10
		lastDigit := yearsLeft - remainder*10
		if yearsLeft > 10 && yearsLeft < 15 || lastDigit > 4 {
			yearsInfo = "лет"
		} else if lastDigit == 1 {
			yearsInfo = "год"
		} else {
			yearsInfo = "года"
		}

		fmt.Printf(
			"Входные данные:\n"+
				"> Финансовая цель: %.2f;\n"+
				"> Горизонт инвестирования: %d %s;\n"+
				"> Номинальная процентная ставка: %.2f%%;\n"+
				"> Капитализация процентов: %s;\n\n",
			financialGoal, yearsLeft, yearsInfo, annualRate, capitalizeInfo,
		)

		periodRate := annualRate * 0.01 / 12
		coefficient := 1 + periodRate

		finalCoefficient := coefficient
		for i := 1; i < 12*int(yearsLeft); i++ {
			finalCoefficient *= coefficient
		}

		monthlyPayment := (financialGoal * periodRate) / (coefficient*finalCoefficient - coefficient)

		var checkTotal float32
		var checkInterest float32
		var checkPersonal float32

		t := getTableWriter()
		t.SetTitle("План по достижению цели")
		t.AppendHeader(table.Row{"Месяц", "Вложения", "Проценты", "Накопления"})

		// Проценты с последнего месяца всего горизонта инвестирования начисляются
		// в следующем месяце, поэтому итераций в цикле на 1 больше и в этой последней
		// итерации мы прибавляем только проценты с прошлого месяца.

		var index interface{}
		periods := 12 * int(yearsLeft)
		for i := 0; i <= periods; i++ {
			interest := checkTotal * periodRate
			checkInterest += interest
			if i == periods {
				checkTotal += interest
				index = "Итого"
				t.AppendSeparator()
			} else {
				checkTotal += interest + monthlyPayment
				checkPersonal += monthlyPayment
				index = i + 1
			}

			t.AppendRow(table.Row{
				index,
				fmt.Sprintf("%.2f", checkPersonal),
				fmt.Sprintf("%.2f", checkInterest),
				fmt.Sprintf("%.2f", checkTotal),
			})
		}
		t.Render()

		fmt.Printf(
			"\n> Ежемесячный взнос составит: %.2f\n"+
				"> Сумма собственных вложений за период: %.2f\n"+
				"> Сумма начисленных процентов за период: %.2f\n\n",
			monthlyPayment, checkPersonal, checkInterest)
	},
}
