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

	"github.com/alexeykhan/assist/internal/assist"
)

var (
	core         assist.Assist
	helpFlag     pflag.Flag
	detailedFlag pflag.Flag
)

func init() {
	core = assist.New()

	helpFlag = pflag.Flag{
		Name: "help", Shorthand: "h",
		Usage: "Документация по команде",
	}

	detailedFlag = pflag.Flag{
		Name: "detailed", Shorthand: "d",
		Usage: "Выводить детализированный ответ. M — по месяцам, Y — по годам.",
	}

	assistCmd.Flags().BoolP(helpFlag.Name, helpFlag.Shorthand, false, helpFlag.Usage)
	decomposeCmd.Flags().BoolP(helpFlag.Name, helpFlag.Shorthand, false, helpFlag.Usage)
	calculateCmd.Flags().BoolP(helpFlag.Name, helpFlag.Shorthand, false, helpFlag.Usage)

	decomposeSavings.Flags().Float64P(decomposeSavingsFlags.Goal.Name, decomposeSavingsFlags.Goal.Shorthand, 0, decomposeSavingsFlags.Goal.Usage)
	decomposeSavings.Flags().Uint8P(decomposeSavingsFlags.Years.Name, decomposeSavingsFlags.Years.Shorthand, 0, decomposeSavingsFlags.Years.Usage)
	decomposeSavings.Flags().Float64P(decomposeSavingsFlags.Interest.Name, decomposeSavingsFlags.Interest.Shorthand, 0, decomposeSavingsFlags.Interest.Usage)
	decomposeSavings.Flags().StringP(detailedFlag.Name, detailedFlag.Shorthand, detailedFlag.DefValue, detailedFlag.Usage)
	decomposeSavings.Flags().BoolP(helpFlag.Name, helpFlag.Shorthand, false, helpFlag.Usage)
	_ = decomposeSavings.MarkFlagRequired(decomposeSavingsFlags.Years.Name)
	_ = decomposeSavings.MarkFlagRequired(decomposeSavingsFlags.Interest.Name)
	_ = decomposeSavings.MarkFlagRequired(decomposeSavingsFlags.Goal.Name)

	decomposeRetirement.Flags().Float64P(decomposeRetirementFlags.Expenses.Name, decomposeRetirementFlags.Expenses.Shorthand, 0, decomposeRetirementFlags.Expenses.Usage)
	decomposeRetirement.Flags().Uint8P(decomposeRetirementFlags.Years.Name, decomposeRetirementFlags.Years.Shorthand, 0, decomposeRetirementFlags.Years.Usage)
	decomposeRetirement.Flags().Float64P(decomposeRetirementFlags.Interest.Name, decomposeRetirementFlags.Interest.Shorthand, 0, decomposeRetirementFlags.Interest.Usage)
	decomposeRetirement.Flags().StringP(detailedFlag.Name, detailedFlag.Shorthand, detailedFlag.DefValue, detailedFlag.Usage)
	decomposeRetirement.Flags().BoolP(helpFlag.Name, helpFlag.Shorthand, false, helpFlag.Usage)
	_ = decomposeRetirement.MarkFlagRequired(decomposeRetirementFlags.Years.Name)
	_ = decomposeRetirement.MarkFlagRequired(decomposeRetirementFlags.Interest.Name)
	_ = decomposeRetirement.MarkFlagRequired(decomposeRetirementFlags.Expenses.Name)

	calculateSavings.Flags().Float64P(calculateSavingsFlags.Payment.Name, calculateSavingsFlags.Payment.Shorthand, 0, calculateSavingsFlags.Payment.Usage)
	calculateSavings.Flags().Uint8P(calculateSavingsFlags.Years.Name, calculateSavingsFlags.Years.Shorthand, 0, calculateSavingsFlags.Years.Usage)
	calculateSavings.Flags().Float64P(calculateSavingsFlags.Interest.Name, calculateSavingsFlags.Interest.Shorthand, 0, calculateSavingsFlags.Interest.Usage)
	calculateSavings.Flags().StringP(detailedFlag.Name, detailedFlag.Shorthand, detailedFlag.DefValue, detailedFlag.Usage)
	calculateSavings.Flags().BoolP(helpFlag.Name, helpFlag.Shorthand, false, helpFlag.Usage)
	_ = calculateSavings.MarkFlagRequired(calculateSavingsFlags.Years.Name)
	_ = calculateSavings.MarkFlagRequired(calculateSavingsFlags.Interest.Name)
	_ = calculateSavings.MarkFlagRequired(calculateSavingsFlags.Payment.Name)

	calculateInflation.Flags().Float64P(calculateInflationFlags.Current.Name, calculateInflationFlags.Current.Shorthand, 0, calculateInflationFlags.Current.Usage)
	calculateInflation.Flags().Uint8P(calculateInflationFlags.Years.Name, calculateInflationFlags.Years.Shorthand, 0, calculateInflationFlags.Years.Usage)
	calculateInflation.Flags().Float64P(calculateInflationFlags.Inflation.Name, calculateInflationFlags.Inflation.Shorthand, 0, calculateInflationFlags.Inflation.Usage)
	calculateInflation.Flags().BoolP(calculateInflationFlags.Detailed.Name, calculateInflationFlags.Detailed.Shorthand, false, calculateInflationFlags.Detailed.Usage)
	calculateInflation.Flags().BoolP(helpFlag.Name, helpFlag.Shorthand, false, helpFlag.Usage)
	_ = calculateInflation.MarkFlagRequired(calculateInflationFlags.Current.Name)
	_ = calculateInflation.MarkFlagRequired(calculateInflationFlags.Years.Name)
	_ = calculateInflation.MarkFlagRequired(calculateInflationFlags.Inflation.Name)

	assistCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		printHeader()
		printDescriptor(cmd)
	})
}

func main() {
	assistCmd.AddCommand(calculateCmd)
	calculateCmd.AddCommand(calculateSavings)
	calculateCmd.AddCommand(calculateInflation)

	assistCmd.AddCommand(decomposeCmd)
	decomposeCmd.AddCommand(decomposeSavings)
	decomposeCmd.AddCommand(decomposeRetirement)

	if err := assistCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
