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

package oracle

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"retire/internal/console"
)

var logo = `
  ooooooo  oooooooooo       o       oooooooo8 ooooo       ooooooooooo
o888   888o 888    888     888    o888     88  888         888    88
888     888 888oooo88     8  88   888          888         888ooo8
888o   o888 888  88o     8oooo88  888o     oo  888      o  888    oo
  88ooo88  o888o  88o8 o88o  o888o 888oooo88  o888ooooo88 o888ooo8888`

var oracle = &cobra.Command{
	Use:   "oracle",
	Short: "Oracle — Программа для декомпозиции финансовых целей",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cmd.SilenceUsage = true
	},
}

var accumulateFlags = struct {
	YearsLeft     pflag.Flag
	AnnualRate    pflag.Flag
	FinancialGoal pflag.Flag
	Capitalize    pflag.Flag
}{
	YearsLeft: pflag.Flag{
		Name:      "years",
		Shorthand: "y",
		Usage:     "Number of years it takes to save up or spend money",
		DefValue:  "",
	},
	AnnualRate: pflag.Flag{
		Name:      "rate",
		Shorthand: "r",
		Usage:     "Annual interest or inflation rate",
		DefValue:  "",
	},
	FinancialGoal: pflag.Flag{
		Name:      "goal",
		Shorthand: "g",
		Usage:     "Your financial goal to achieve during given period",
		DefValue:  "",
	},
	// Fixed transaction amount to replenish or withdraw money
	Capitalize: pflag.Flag{
		Name:      "capitalize",
		Shorthand: "c",
		Usage:     "Whether to take into account interest capitalization or not",
		DefValue:  "",
	},
}

var predict = &cobra.Command{
	Use:   "predict",
	Short: "Построить модель решения задачи",
	Run:   func(cmd *cobra.Command, args []string) { _ = cmd.Help() },
}

var savings = &cobra.Command{
	Use:   "savings",
	Short: "Рассчитать, сколько денег будет по итогу периоду",
	Example: console.Examples(
		"# Узнайте, сколько денег сможете накопить за период:\n" +
			"./bin/oracle predict savings --goal 10234567.89 --rate 6.5 --years 10\n" +
			"./bin/oracle predict savings -g 10234567.89 -r 6.5 -y 10",
	),
	Run: func(cmd *cobra.Command, args []string) {
		yearsLeft := console.GetUint8(cmd, accumulateFlags.YearsLeft.Name)
		annualRate := console.GetFloat32(cmd, accumulateFlags.AnnualRate.Name)
		financialGoal := console.GetFloat32(cmd, accumulateFlags.FinancialGoal.Name)
		capitalize := console.GetBool(cmd, accumulateFlags.Capitalize.Name)

		capitalizeInfo := "выключена"
		if capitalize {
			capitalizeInfo = "включена"
		}

		var yearsInfo string
		if yearsLeft > 10 && yearsLeft < 15 {
			yearsInfo = "лет"
		} else {
			remainder := yearsLeft / 10
			lastDigit := yearsLeft - remainder*10
			switch lastDigit {
			case 1:
				yearsInfo = "год"
			case 2, 3, 4:
				yearsInfo = "года"
			default:
				yearsInfo = "лет"
			}
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

		t := console.GetTableWriter()
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

func init() {
	savings.Flags().BoolP(accumulateFlags.Capitalize.Name, accumulateFlags.Capitalize.Shorthand, true, accumulateFlags.Capitalize.Usage)
	savings.Flags().Uint8P(accumulateFlags.YearsLeft.Name, accumulateFlags.YearsLeft.Shorthand, 0, accumulateFlags.YearsLeft.Usage)
	savings.Flags().Float32P(accumulateFlags.AnnualRate.Name, accumulateFlags.AnnualRate.Shorthand, 0, accumulateFlags.AnnualRate.Usage)
	savings.Flags().Float32P(accumulateFlags.FinancialGoal.Name, accumulateFlags.FinancialGoal.Shorthand, 0, accumulateFlags.FinancialGoal.Usage)
	_ = savings.MarkFlagRequired(accumulateFlags.YearsLeft.Name)
	_ = savings.MarkFlagRequired(accumulateFlags.AnnualRate.Name)
	_ = savings.MarkFlagRequired(accumulateFlags.FinancialGoal.Name)

	oracle.AddCommand(predict)
	predict.AddCommand(savings)

	txt := text.Colors{text.Bold, text.FgHiGreen}

	fmt.Printf(
		"%s\n\n%s\n\n",
		txt.Sprint(logo),
		txt.Sprint("[version|0.1.0] [updated|2021-01-22] [copyright|Alexey Khan]"))
}

// Execute ...
func Execute() {
	if err := oracle.Execute(); err != nil {
		os.Exit(1)
	}
}
