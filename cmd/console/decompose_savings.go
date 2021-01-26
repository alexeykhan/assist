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

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var decomposeSavingsConfig = struct {
	title    string
	about    string
	overview string
	results  string
	detailed string
	examples []string
}{
	title: "Декомпозиция накопления суммы",
	about: "Узнайте, какую сумму необходимо инвестировать каждый месяц, чтобы при " +
		"заданных доходности портфеля R% годовых, горизонте инвестирования N лет и " +
		"ежемесячной капитализации процентов накопить к концу срока нужную сумму X.",
	overview: "Задача: рассчитать сумму, которую необходимо инвестировать каждый месяц " +
		"на протяжении %s, чтобы при средней доходности портфеля %.2f%% " +
		"годовых и ежемесячной капитализации процентов накопить %.2f руб.",
	results: " > Сумма ежемесячных инвестиций составит: %.2f\n\n",
	detailed: "\n > Сумма ежемесячных инвестиций составит: %.2f\n" +
		" > Сумма собственных вложений за период: %.2f\n" +
		" > Сумма начисленных процентов за период: %.2f\n\n",
	examples: []string{
		"./bin/assist decompose savings --goal=1234567.89 --years=10 --interest=6.5 --detailed=M",
		"./bin/assist decompose savings -g=1234567.89 -y=10 -i=6.5 -d=M",
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
		detailed := getString(cmd, detailedFlag.Name)

		if err = validateDetailedOption(detailed); err != nil {
			return err
		}

		overview := fmt.Sprintf(decomposeSavingsConfig.overview, yearsDuration(years), interest, goal)
		taskOverview := getTaskOverview(decomposeSavingsConfig.title, overview)

		var payment float64
		if payment, err = core.DecomposeSavings(goal, interest, years); err != nil {
			return err
		}

		if detailed == commandOptionEmpty {
			fmt.Println(taskOverview)
			fmt.Printf(decomposeSavingsConfig.results, payment)
			return
		}

		rendered, personalInvestments, interestIncome, _ := savingsTable(
			payment, interest, years, detailed == commandOptionDetailedMonthly)

		fmt.Println(taskOverview)
		fmt.Println(rendered)
		fmt.Printf(decomposeSavingsConfig.detailed, payment, personalInvestments, interestIncome)

		return nil
	},
}
