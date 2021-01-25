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
	overview string
	results  string
	examples []string
}{
	title: "Декомпозиция накопления суммы",
	about: "Узнайте, какую сумму необходимо инвестировать каждый месяц, чтобы при " +
		"заданных доходности портфеля R% годовых, горизонте инвестирования N лет и " +
		"ежемесячной капитализации процентов накопить к концу срока нужную сумму X.",
	overview: "Задача: рассчитать сумму, которую необходимо инвестировать каждый месяц " +
		"на протяжении %s, чтобы при средней доходности портфеля %.2f%% " +
		"годовых и ежемесячной капитализации процентов накопить %.2f руб.",
	results: "\n > Ежемесячный взнос составит: %.2f\n" +
		" > Сумма собственных вложений за период: %.2f\n" +
		" > Сумма начисленных процентов за период: %.2f\n\n",
	examples: []string{
		"./bin/assist decompose savings --goal=1234567.89 --years=10 --interest=6.5",
		"./bin/assist decompose savings -g=1234567.89 -y=10 -i=6.5",
		"./bin/assist decompose savings --help",
	},
}

var decomposeSavingsFlags = struct {
	Goal     pflag.Flag
	Years    pflag.Flag
	Interest pflag.Flag
}{
	Goal: pflag.Flag{
		Name: "goal", Shorthand: "g",
		Usage: "Ваша финансовая цель, которую нужно достгнуть за заданный период",
	},
	Years: pflag.Flag{
		Name: "years", Shorthand: "y",
		Usage: "Количество лет, за которое необходимо накопить нужную сумму",
	},
	Interest: pflag.Flag{
		Name: "interest", Shorthand: "i",
		Usage: "Доходность вашего инвестиционного портфеля в процентах годовых",
	},
}

var decomposeSavings = &cobra.Command{
	Use: "savings",
	Example: commandOverview(
		decomposeSavingsConfig.title,
		decomposeSavingsConfig.about,
		decomposeSavingsConfig.examples,
	),
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		printHeader()

		goal := getFloat64(cmd, decomposeSavingsFlags.Goal.Name)
		years := getUint8(cmd, decomposeSavingsFlags.Years.Name)
		interest := getFloat64(cmd, decomposeSavingsFlags.Interest.Name)
		detailed := getBool(cmd, detailedFlag.Name)

		var payment float64
		if payment, err = core.DecomposeSavings(goal, interest, years); err != nil {
			return err
		}

		view := core.View()

		boldWhiteText := text.Colors{text.Bold, text.FgHiWhite}
		normalWhiteText := text.Colors{text.FgHiWhite}

		upperCaseTitle := text.FormatUpper.Apply(decomposeSavingsConfig.title)
		formattedTitle := boldWhiteText.Sprintf(" %s", upperCaseTitle)

		yearsInfo := view.YearsDuration(years)
		filledTask := fmt.Sprintf(decomposeSavingsConfig.overview, yearsInfo, interest, goal)
		wrappedTask := text.WrapSoft(filledTask, appViewWidth-2)

		taskOverview := formattedTitle + "\n\n"
		for _, line := range strings.Split(wrappedTask, "\n") {
			trimmedLine := strings.TrimSpace(line)
			taskOverview += normalWhiteText.Sprintf(" %s\n", trimmedLine)
		}

		t := getTableWriter(
			tableColumnYear,
			tableColumnInvestments,
			tableColumnInterestIncome,
			tableColumnTotalSavings)

		var (
			next  int
			index interface{}

			totalSavings, interestIncome, personalInvestments float64
		)

		periods := 12 * int(years)
		periodRate := interest * 0.01 / 12
		for i := 0; i <= periods; i++ {
			interest := totalSavings * periodRate
			interestIncome += interest
			if i == periods {
				totalSavings += interest
				t.AppendSeparator()
			} else {
				totalSavings += interest + payment
				personalInvestments += payment
			}

			next = i + 1
			index = next

			if !detailed {
				index = next / 12
			}
			if i == periods {
				index = tableFooterTotal
			}
			if detailed || (next >= 12 && next%12 == 0 || i == periods) {
				t.AppendRow(table.Row{
					index,
					fmt.Sprintf("%.2f", personalInvestments),
					fmt.Sprintf("%.2f", interestIncome),
					fmt.Sprintf("%.2f", totalSavings),
				})
			}
		}

		fmt.Println(taskOverview)
		fmt.Println(t.Render())
		fmt.Printf(decomposeSavingsConfig.results, payment, personalInvestments, interestIncome)

		return nil
	},
}
