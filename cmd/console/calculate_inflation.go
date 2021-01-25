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

var calculateInflationConfig = struct {
	title    string
	about    string
	overview string
	results  string
	detailed string
	examples []string
}{
	title: "Расчет покупательской способности",
	about: "Узнайте, как изменится конкретное денежное значение через N лет при среднем " +
		"уровне инфляции R% годовых.",
	overview: "Задача: рассчитать изменение покупательской способности по прошествии %s " +
		"при заданном уровне инфляции %.2f%% в год на примере изменения денежной суммы, " +
		"по сегодняшним меркам равной %.2f рублям.",
	results:  " > Эквивалент исходной суммы в будущем: %.2f\n\n",
	detailed: "\n > Эквивалент исходной суммы в будущем: %.2f\n" +
		" > Ценность исходной суммы в будущем: %.2f\n\n",
	examples: []string{
		"./bin/assist calculate inflation --current=150000 --years=10 --inflation=6.5 --detailed",
		"./bin/assist calculate inflation -c=150000 -y=10 -i=6.5 -d",
		"./bin/assist calculate inflation --help",
	},
}

var calculateInflationFlags = struct {
	Years     pflag.Flag
	Current   pflag.Flag
	Inflation pflag.Flag
	Detailed  pflag.Flag
}{
	Years: pflag.Flag{
		Name: "years", Shorthand: "y",
		Usage: "Количество лет, на протяжении которых будет меняться покупательская способность",
	},
	Current: pflag.Flag{
		Name: "current", Shorthand: "c",
		Usage: "Сумма, ценность которой будет изменяться с течением времени в силу инфляции",
	},
	Inflation: pflag.Flag{
		Name: "inflation", Shorthand: "i",
		Usage: "Среднее значение инфляции, при котором будут проводиться расчеты",
	},
	Detailed: pflag.Flag{
		Name: "detailed", Shorthand: "d",
		Usage: "Вывести детализацию изменения ценности по годам",
	},
}

var calculateInflation = &cobra.Command{
	Use: "inflation",
	Example: commandOverview(
		calculateInflationConfig.title,
		calculateInflationConfig.about,
		calculateInflationConfig.examples,
	),
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		printHeader()

		years := getUint8(cmd, calculateInflationFlags.Years.Name)
		inflation := getFloat64(cmd, calculateInflationFlags.Inflation.Name)
		current := getFloat64(cmd, calculateInflationFlags.Current.Name)
		detailed := getBool(cmd, calculateInflationFlags.Detailed.Name)

		overview := fmt.Sprintf(
			calculateInflationConfig.overview,
			core.View().YearsDuration(years),
			inflation, current)

		taskOverview := getTaskOverview(calculateInflationConfig.title, overview)

		if !detailed {
			var inflated float64
			if inflated, err = core.CalculateInflation(current, inflation, years); err != nil {
				return err
			}

			fmt.Println(taskOverview)
			fmt.Printf(calculateInflationConfig.results, inflated)
			return
		}

		t := getTableWriter(tableColumnYear, tableColumnInflationInitial, tableColumnInflationEquivalent)

		square := current * current
		initialValue := current
		equivalentValue := current
		for i := 1; i <= int(years); i++ {
			equivalentValue *= inflation*0.01 + 1
			initialValue = square / equivalentValue
			t.AppendRow(table.Row{i,
				fmt.Sprintf("%.2f (%.2f)", initialValue, initialValue-current),
				fmt.Sprintf("%.2f (+%.2f)", equivalentValue, equivalentValue-current)})
		}

		fmt.Println(taskOverview)
		fmt.Println(t.Render())
		fmt.Printf(calculateInflationConfig.detailed, equivalentValue, initialValue)
		return
	},
}
