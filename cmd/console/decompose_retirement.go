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

var decomposeRetirementConfig = struct {
	title    string
	about    string
	overview string
	results  string
	detailed string
	examples []string
}{
	title: "Декомпозиция пенсии",
	about: "Узнайте, какую минимальную сумму нужно накопить, чтобы при доходности " +
		"портфеля R% годовых можно было на протяжении N лет тратить X рублей в месяц " +
		"без дополнительного дохода, потратив к концу срока все сбережения.",
	overview: "Задача: рассчитать сумму, которую необходимо накопить, чтобы при доходности " +
		"портфеля %.2f%% годовых можно было на протяжении %s тратить %2.f рублей в месяц " +
		"без дополнительного дохода, потратив к концу срока все сбережения.",
	results: " > Минимальная сумма накоплений составит: %.2f\n\n",
	detailed: "\n > Минимальная сумма накоплений составит: %.2f\n" +
		" > Сумма начисленных процентов за период: %.2f\n\n",
	examples: []string{
		"./bin/assist decompose retirement --expenses=150000 --years=25 --interest=6.5 --detailed=M",
		"./bin/assist decompose retirement -e=150000 -y=25 -i=6.5 -d=M",
		"./bin/assist decompose retirement --help",
	},
}

var decomposeRetirementFlags = struct {
	Years    pflag.Flag
	Interest pflag.Flag
	Expenses pflag.Flag
}{
	Years: pflag.Flag{
		Name: "years", Shorthand: "y",
		Usage: "Количество лет, на протяжении которых будут ежемесячные траты",
	},
	Interest: pflag.Flag{
		Name: "interest", Shorthand: "i",
		Usage: "Доходность вашего инвестиционного портфеля в процентах годовых",
	},
	Expenses: pflag.Flag{
		Name: "expenses", Shorthand: "e",
		Usage: "Сумма ежемесячных расходов в течение пенсионного периода",
	},
}

var decomposeRetirement = &cobra.Command{
	Use: "retirement",
	Example: commandOverview(
		decomposeRetirementConfig.title,
		decomposeRetirementConfig.about,
		decomposeRetirementConfig.examples,
	),
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		printHeader()

		years := getUint8(cmd, decomposeRetirementFlags.Years.Name)
		interest := getFloat64(cmd, decomposeRetirementFlags.Interest.Name)
		expenses := getFloat64(cmd, decomposeRetirementFlags.Expenses.Name)
		detailed := getString(cmd, detailedFlag.Name)

		if err = validateDetailedOption(detailed); err != nil {
			return err
		}

		view := core.View()

		boldWhiteText := text.Colors{text.Bold, text.FgHiWhite}
		normalWhiteText := text.Colors{text.FgHiWhite}

		upperCaseTitle := text.FormatUpper.Apply(decomposeRetirementConfig.title)
		formattedTitle := boldWhiteText.Sprintf(" %s", upperCaseTitle)

		yearsInfo := view.YearsDuration(years)
		filledTask := fmt.Sprintf(decomposeRetirementConfig.overview, interest, yearsInfo, expenses)
		wrappedTask := text.WrapSoft(filledTask, appViewWidth-2)

		taskOverview := formattedTitle + "\n\n"
		for _, line := range strings.Split(wrappedTask, "\n") {
			trimmedLine := strings.TrimSpace(line)
			taskOverview += normalWhiteText.Sprintf(" %s\n", trimmedLine)
		}

		var retirement float64
		if retirement, err = core.DecomposeRetirement(expenses, interest, years); err != nil {
			return err
		}

		if detailed == commandOptionEmpty {
			fmt.Println(taskOverview)
			fmt.Printf(decomposeRetirementConfig.results, retirement)
			return
		}

		t := getTableWriter(
			tableColumnYear,
			tableColumnInterestIncome,
			tableColumnExpenses,
			tableColumnTotalSavings)

		var (
			next           int
			index          interface{}
			interestIncome float64
			totalExpenses  float64
		)

		savingsLeft := retirement
		periods := 12 * int(years)
		periodRate := interest * 0.01 / 12
		detailedMonthly := detailed == commandOptionDetailedMonthly
		for i := 0; i < periods; i++ {
			interest := (savingsLeft - expenses) * periodRate
			if interest < 0 {
				interest = 0
			}
			interestIncome += interest
			totalExpenses += expenses
			savingsLeft += interest - expenses

			next = i + 1
			index = next

			if !detailedMonthly {
				index = next / 12
			}

			if savingsLeft < 0 {
				savingsLeft = 0
			}

			if detailedMonthly || (next >= 12 && next%12 == 0 || i == periods) {
				t.AppendRow(table.Row{
					index,
					fmt.Sprintf("%.2f", interestIncome),
					fmt.Sprintf("%.2f", totalExpenses),
					fmt.Sprintf("%.2f", savingsLeft),
				})
			}

			if i == periods-1 {
				t.AppendSeparator()
				t.AppendRow(table.Row{
					tableFooterTotal,
					fmt.Sprintf("%.2f", interestIncome),
					fmt.Sprintf("%.2f", totalExpenses),
					fmt.Sprintf("%.2f", savingsLeft),
				})
			}
		}

		fmt.Println(taskOverview)
		fmt.Println(t.Render())
		fmt.Printf(decomposeRetirementConfig.detailed, retirement, interestIncome)

		return nil
	},
}
