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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var helpFlag = pflag.Flag{
	Name:      "help",
	Shorthand: "h",
	Usage:     "Документация по команде",
	DefValue:  "",
}

func init() {
	oracle.Flags().BoolP(helpFlag.Name, helpFlag.Shorthand, false, helpFlag.Usage)
	decompose.Flags().BoolP(helpFlag.Name, helpFlag.Shorthand, false, helpFlag.Usage)

	decomposeSavings.Flags().Float32P(decomposeSavingsFlags.FinancialGoal.Name, decomposeSavingsFlags.FinancialGoal.Shorthand, 0, decomposeSavingsFlags.FinancialGoal.Usage)
	decomposeSavings.Flags().Uint8P(decomposeSavingsFlags.YearsLeft.Name, decomposeSavingsFlags.YearsLeft.Shorthand, 0, decomposeSavingsFlags.YearsLeft.Usage)
	decomposeSavings.Flags().Float32P(decomposeSavingsFlags.InterestRate.Name, decomposeSavingsFlags.InterestRate.Shorthand, 0, decomposeSavingsFlags.InterestRate.Usage)
	decomposeSavings.Flags().BoolP(helpFlag.Name, helpFlag.Shorthand, false, helpFlag.Usage)
	_ = decomposeSavings.MarkFlagRequired(decomposeSavingsFlags.YearsLeft.Name)
	_ = decomposeSavings.MarkFlagRequired(decomposeSavingsFlags.InterestRate.Name)
	_ = decomposeSavings.MarkFlagRequired(decomposeSavingsFlags.FinancialGoal.Name)

	oracle.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		printHeader()
		printDescriptor(cmd)
	})
}

func main() {
	oracle.AddCommand(decompose)
	decompose.AddCommand(decomposeSavings)

	if err := oracle.Execute(); err != nil {
		os.Exit(1)
	}
}
