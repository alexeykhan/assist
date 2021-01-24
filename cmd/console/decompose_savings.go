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
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var decomposeSavingsConfig = struct {
	title    string
	about    string
	template string
	examples []string
}{
	title: "Декомпозиция накопления суммы",
	about: "Узнайте, какую сумму необходимо инвестировать каждый месяц, чтобы при " +
		"заданных доходности портфеля P% годовых, горизонте инвестирования N лет и " +
		"ежемесячной капитализации процентов накопить к концу срока нужную сумму X.",
	template: "Задача: рассчитать сумму, которую необходимо инвестировать каждый месяц " +
		"на протяжении %d %s, чтобы при средней доходности портфеля %.2f%% " +
		"годовых и ежемесячной капитализации процентов накопить %.2f руб.",
	examples: []string{
		"./assist decompose savings --goal=1234567.89 --years=10 --interest=6.5",
		"./assist decompose savings -g=1234567.89 -y=10 -i=6.5",
		"./assist decompose savings --help",
	},
}

var decomposeSavingsFlags = struct {
	YearsLeft     pflag.Flag
	InterestRate  pflag.Flag
	FinancialGoal pflag.Flag
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
}

var decomposeSavings = &cobra.Command{
	Use: "savings",
	Example: example(
		decomposeSavingsConfig.title,
		decomposeSavingsConfig.about,
		decomposeSavingsConfig.examples,
	),
	Run: func(cmd *cobra.Command, args []string) {
		yearsLeft := getUint8(cmd, decomposeSavingsFlags.YearsLeft.Name)
		annualRate := getFloat32(cmd, decomposeSavingsFlags.InterestRate.Name)
		financialGoal := getFloat32(cmd, decomposeSavingsFlags.FinancialGoal.Name)

		yearsInfo := "лет"
		if yearsLeft != 11 && yearsLeft%10 == 1 {
			yearsInfo = "года"
		}

		boldWhiteText := text.Colors{text.Bold, text.FgHiWhite}
		normalWhiteText := text.Colors{text.FgHiWhite}

		var task string
		upperCaseTitle := text.FormatUpper.Apply(decomposeSavingsConfig.title)
		formattedTitle := boldWhiteText.Sprintf(" %s", upperCaseTitle)
		task += formattedTitle + "\n\n"

		filledTask := fmt.Sprintf(decomposeSavingsConfig.template, yearsLeft, yearsInfo, annualRate, financialGoal)
		wrappedTask := text.WrapSoft(filledTask, appViewWidth-2)
		for _, line := range strings.Split(wrappedTask, "\n") {
			trimmedLine := strings.TrimSpace(line)
			task += normalWhiteText.Sprintf(" %s\n", trimmedLine)
		}

		fmt.Println(task)

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

		yearColumnWidth := 6
		moneyColumnMaxWidth := (appViewWidth - yearColumnWidth - 16) / 3

		t := getTableWriter()
		t.SetAllowedRowLength(appViewWidth)
		t.AppendHeader(table.Row{"Год", "Вложения", "Проценты", "Накопления"})
		t.SetStyle(table.Style{
			Name: "myNewStyle",
			Box: table.BoxStyle{
				BottomLeft:       " ┗",
				BottomRight:      "┛",
				BottomSeparator:  "━┻",
				Left:             " ┃",
				LeftSeparator:    " ┣",
				MiddleHorizontal: "━",
				MiddleSeparator:  "━╋",
				MiddleVertical:   " ┃",
				PaddingLeft:      " ",
				PaddingRight:     " ",
				Right:            "┃",
				RightSeparator:   "┫",
				TopLeft:          " ┏",
				TopRight:         "┓",
				TopSeparator:     "━┳",
				UnfinishedRow:    " ~~~",
			},
			Color: table.ColorOptions{
				Footer:       text.Colors{text.FgHiWhite},
				Header:       text.Colors{text.FgHiWhite},
				Row:          text.Colors{text.FgHiWhite},
				RowAlternate: text.Colors{text.FgHiWhite},
			},
			Format: table.FormatOptions{
				Footer: text.FormatUpper,
				Header: text.FormatUpper,
				Row:    text.FormatDefault,
			},
			Options: table.Options{
				DrawBorder:      true,
				SeparateColumns: true,
				SeparateFooter:  true,
				SeparateHeader:  true,
				SeparateRows:    false,
			},
		})
		t.SetColumnConfigs([]table.ColumnConfig{
			{
				Name:        "Год",
				Align:       text.AlignCenter,
				AlignFooter: text.AlignLeft,
				AlignHeader: text.AlignCenter,
				WidthMin:    yearColumnWidth,
				WidthMax:    yearColumnWidth,
			},
			{
				Name:        "Вложения",
				Align:       text.AlignCenter,
				AlignFooter: text.AlignLeft,
				AlignHeader: text.AlignCenter,
				WidthMin:    moneyColumnMaxWidth,
				WidthMax:    moneyColumnMaxWidth,
			},
			{
				Name:        "Проценты",
				Align:       text.AlignCenter,
				AlignFooter: text.AlignLeft,
				AlignHeader: text.AlignCenter,
				WidthMin:    moneyColumnMaxWidth,
				WidthMax:    moneyColumnMaxWidth,
			},
			{
				Name:        "Накопления",
				Align:       text.AlignCenter,
				AlignFooter: text.AlignLeft,
				AlignHeader: text.AlignCenter,
				WidthMin:    moneyColumnMaxWidth,
				WidthMax:    moneyColumnMaxWidth,
			},
		})

		// Проценты с последнего месяца всего горизонта инвестирования начисляются
		// в следующем месяце, поэтому итераций в цикле на 1 больше и в этой последней
		// итерации мы прибавляем только проценты с прошлого месяца.

		var index interface{}
		var next int
		periods := 12 * int(yearsLeft)
		for i := 0; i <= periods; i++ {
			interest := checkTotal * periodRate
			checkInterest += interest
			if i == periods {
				checkTotal += interest
				t.AppendSeparator()
			} else {
				checkTotal += interest + monthlyPayment
				checkPersonal += monthlyPayment
			}

			next = i + 1
			if next >= 12 && next%12 == 0 || i == periods {
				if i == periods {
					index = "ИТОГО"
				} else {
					index = next / 12
				}

				t.AppendRow(table.Row{
					index,
					fmt.Sprintf("%.2f", checkPersonal),
					fmt.Sprintf("%.2f", checkInterest),
					fmt.Sprintf("%.2f", checkTotal),
				})
			}
		}
		t.Render()

		fmt.Printf(
			"\n> Ежемесячный взнос составит: %.2f\n"+
				"> Сумма собственных вложений за период: %.2f\n"+
				"> Сумма начисленных процентов за период: %.2f\n\n",
			monthlyPayment, checkPersonal, checkInterest)
	},
}
