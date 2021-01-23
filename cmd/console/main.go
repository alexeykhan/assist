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
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func init() {
	decomposeSavings.Flags().BoolP(decomposeSavingsFlags.Capitalize.Name, decomposeSavingsFlags.Capitalize.Shorthand, true, decomposeSavingsFlags.Capitalize.Usage)
	decomposeSavings.Flags().Uint8P(decomposeSavingsFlags.YearsLeft.Name, decomposeSavingsFlags.YearsLeft.Shorthand, 0, decomposeSavingsFlags.YearsLeft.Usage)
	decomposeSavings.Flags().Float32P(decomposeSavingsFlags.InterestRate.Name, decomposeSavingsFlags.InterestRate.Shorthand, 0, decomposeSavingsFlags.InterestRate.Usage)
	decomposeSavings.Flags().Float32P(decomposeSavingsFlags.FinancialGoal.Name, decomposeSavingsFlags.FinancialGoal.Shorthand, 0, decomposeSavingsFlags.FinancialGoal.Usage)
	decomposeSavings.Flags().BoolP(decomposeSavingsFlags.Help.Name, decomposeSavingsFlags.Help.Shorthand, false, decomposeSavingsFlags.Help.Usage)
	_ = decomposeSavings.MarkFlagRequired(decomposeSavingsFlags.YearsLeft.Name)
	_ = decomposeSavings.MarkFlagRequired(decomposeSavingsFlags.InterestRate.Name)
	_ = decomposeSavings.MarkFlagRequired(decomposeSavingsFlags.FinancialGoal.Name)

	oracle.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		printHeader()

		fmt.Println(cmd.Example)
		fmt.Println(text.Colors{text.Bold, text.FgHiWhite}.Sprint(" Параметры и опции команды:"))

		var maxlen int
		var flagLines []string
		cmd.LocalFlags().VisitAll(func(flag *pflag.Flag) {
			if flag.Hidden {
				return
			}

			line := ""
			if flag.Shorthand != "" && flag.ShorthandDeprecated == "" {
				line = fmt.Sprintf("  -%s, --%s", flag.Shorthand, flag.Name)
			} else {
				line = fmt.Sprintf("      --%s", flag.Name)
			}

			varname, usage := unquoteUsage(flag)
			line += " " + varname

			line += "\x00"
			if len(line) > maxlen {
				maxlen = len(line)
			}

			line += usage
			flagLines = append(flagLines, line)
		})

		for _, line := range flagLines {
			sidx := strings.Index(line, "\x00")
			spacing := strings.Repeat(" ", maxlen-sidx)
			concatenated := line[:sidx] + spacing + " " + wrapUsage(line[sidx+1:], appViewWidth, maxlen+1)
			fmt.Println(text.Colors{text.FgHiWhite}.Sprint(concatenated))
		}
	})
}

func main() {
	oracle.AddCommand(decompose)
	decompose.AddCommand(decomposeSavings)

	if err := oracle.Execute(); err != nil {
		// _ = cmd.Help()
		os.Exit(1)
	}
}
