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

import "github.com/spf13/cobra"

var decompose = &cobra.Command{
	Use: "decompose",
	Run: func(cmd *cobra.Command, args []string) { _ = cmd.Help() },
	Example: example(
		"Декомпозиция финансовой цели",
		"Используйте команду `savings`, чтобы определить сумму ежемесячных инвестиций, необходимых "+
			"для достижения финансовой цели за заданный период в годах с доходностью портфеля P% годовых "+
			"и капитализацией процентов.\n\nПример: узнайте минимальную сумму ежемесячных инвестиций, чтобы "+
			"при доходности потфеля 5% годовых накопить 5,000,000 руб за ближайшие 10 лет.\n\n"+
			"Используйте команду `expenses`, чтобы определить минимальную сумму, которая при доходности "+
			"P% годовых позволит на протяжении N лет тратить каждый месяц X рублей без дополнительного "+
			"дохода.\n\nПример: узнайте минимальный объем портфеля с доходностью 5% годовых, чтобы к выходу "+
			"на пенсию через 35 лет можно было тратить 100,000 руб в месяц на протяжении следующих 25 лет.",
		[]string{
			"./oracle decompose savings [options]",
			"./oracle decompose expenses [options]",
			"./oracle decompose --help",
			"./oracle decompose -h",
		},
	),
}
