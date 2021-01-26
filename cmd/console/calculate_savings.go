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

var calculateSavingsConfig = struct {
	title    string
	about    string
	overview string
	results  string
	detailed string
	examples []string
}{
	title: "Расчет будущих накоплений",
	about: "Узнайте, какую сумму сможете накопить с учетом сложного процента, если на " +
		"протяжении следующих N лет будете ежемесячно инвестировать X рублей под R% годовых " +
		"с ежемесячной капитализацией процентов.",
	overview: "Задача: рассчитать сумму, которую можно накопить с учетом сложного процента, " +
		"если на протяжении следующих %s ежемесячно инвестировать %.2f рублей под %.2f%% " +
		"годовых с ежемесячной капитализацией процентов.",
	results: " > Накопленная сумма составит: %.2f\n\n",
	detailed: "\n > Накопленная сумма составит: %.2f\n" +
		" > Сумма собственных вложений за период: %.2f\n" +
		" > Сумма начисленных процентов за период: %.2f\n\n",
	examples: []string{
		"./bin/assist calculate savings --payment=10000 --years=10 --interest=6.5 --detailed=M",
		"./bin/assist calculate savings -p=10000 -y=10 -i=6.5 -d=M",
		"./bin/assist calculate savings --help",
	},
}

var calculateSavingsFlags = struct {
	Years    pflag.Flag
	Payment  pflag.Flag
	Interest pflag.Flag
}{
	Years: pflag.Flag{
		Name: "years", Shorthand: "y",
		Usage: "Количество лет, на протяжении которых будут производиться накопления",
	},
	Payment: pflag.Flag{
		Name: "payment", Shorthand: "p",
		Usage: "Размер ежемесячного пополнения инвестиционного портфеля",
	},
	Interest: pflag.Flag{
		Name: "interest", Shorthand: "i",
		Usage: "Доходность вашего инвестиционного портфеля в процентах годовых",
	},
}

var calculateSavings = &cobra.Command{
	Use: "savings",
	Example: commandOverview(
		calculateSavingsConfig.title,
		calculateSavingsConfig.about,
		calculateSavingsConfig.examples,
	),
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		printHeader()

		years := getUint8(cmd, calculateSavingsFlags.Years.Name)
		interest := getFloat64(cmd, calculateSavingsFlags.Interest.Name)
		payment := getFloat64(cmd, calculateSavingsFlags.Payment.Name)
		detailed := getString(cmd, detailedFlag.Name)

		if err = validateDetailedOption(detailed); err != nil {
			return err
		}

		overview := fmt.Sprintf(calculateSavingsConfig.overview, yearsDuration(years), payment, interest)
		taskOverview := getTaskOverview(calculateSavingsConfig.title, overview)

		if detailed == commandOptionEmpty {
			var savings float64
			if savings, err = core.CalculateSavings(payment, interest, years); err != nil {
				return err
			}

			fmt.Println(taskOverview)
			fmt.Printf(calculateSavingsConfig.results, savings)
			return
		}

		rendered, personalInvestments, interestIncome, totalSavings := savingsTable(
			payment, interest, years, detailed == commandOptionDetailedMonthly)

		fmt.Println(taskOverview)
		fmt.Println(rendered)
		fmt.Printf(calculateSavingsConfig.detailed, totalSavings, personalInvestments, interestIncome)
		return
	},
}
