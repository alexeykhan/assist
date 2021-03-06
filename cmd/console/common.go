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
	"log"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	appVersion   = "0.1.0"
	appCopyright = "alexeykhan"
	appLicense   = "MIT"
	appViewWidth = 64

	commandOptionEmpty           = ""
	commandOptionDetailedYearly  = "Y"
	commandOptionDetailedMonthly = "M"

	commandOptions       = " Параметры и опции команды:"
	commandUsageExamples = "\n Примеры использования:\n"

	tableColumnYear     = "Год"
	tableColumnMonth    = "Месяц"
	tableColumnExpenses = "Расходы"

	tableColumnInflationInitial    = "Исходная ценность"
	tableColumnInflationEquivalent = "Эквивалент ценности"

	tableColumnInvestments    = "Вложения"
	tableColumnInterestIncome = "Проценты"
	tableColumnTotalSavings   = "Накопления"

	tableFooterTotal = "ИТОГО"
)

var logo = [...]string{
	`  ______    ______    ______   ______   ______   ________ `,
	` /      \  /      \  /      \ /      | /      \ /        |`,
	`/$$$$$$  |/$$$$$$  |/$$$$$$  |$$$$$$/ /$$$$$$  |$$$$$$$$/ `,
	`$$ |__$$ |$$ \__$$/ $$ \__$$/   $$ |  $$ \__$$/    $$ |   `,
	`$$    $$ |$$      \ $$      \   $$ |  $$      \    $$ |   `,
	`$$$$$$$$ | $$$$$$  | $$$$$$  |  $$ |   $$$$$$  |   $$ |   `,
	`$$ |  $$ |/  \__$$ |/  \__$$ | _$$ |_ /  \__$$ |   $$ |   `,
	`$$ |  $$ |$$    $$/ $$    $$/ / $$   |$$    $$/    $$ |   `,
	`$$/   $$/  $$$$$$/   $$$$$$/  $$$$$$/  $$$$$$/     $$/    `,
}

func yearsDuration(years uint8) string {
	info := "лет"
	if years != 11 && years%10 == 1 {
		info = "года"
	}
	return fmt.Sprintf("%d %s", years, info)
}

func savingsTable(payment, interest float64, years uint8, detailed bool) (rendered string, personalInvestments, interestIncome, totalSavings float64) {
	firstColumn := tableColumnYear
	if detailed {
		firstColumn = tableColumnMonth
	}

	t := getTableWriter(
		firstColumn,
		tableColumnInvestments,
		tableColumnInterestIncome,
		tableColumnTotalSavings)

	var index interface{}
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

		index = i + 1
		if !detailed {
			index = (i + 1) / 12
		}

		if i == periods {
			index = tableFooterTotal
		}

		if detailed || (i+1 >= 12 && (i+1)%12 == 0 || i == periods) {
			t.AppendRow(table.Row{
				index,
				fmt.Sprintf("%.2f", personalInvestments),
				fmt.Sprintf("%.2f", interestIncome),
				fmt.Sprintf("%.2f", totalSavings),
			})
		}
	}

	return t.Render(), personalInvestments, interestIncome, totalSavings
}

func commandOverview(title, about string, examples []string) string {
	boldWhiteText := text.Colors{text.Bold, text.FgHiWhite}
	normalWhiteText := text.Colors{text.FgHiWhite}

	var overview string
	upperCaseTitle := text.FormatUpper.Apply(title)
	formattedTitle := boldWhiteText.Sprintf(" %s", upperCaseTitle)
	overview += formattedTitle + "\n\n"

	wrappedAbout := text.WrapSoft(about, appViewWidth-2)
	for _, line := range strings.Split(wrappedAbout, "\n") {
		trimmedLine := strings.TrimSpace(line)
		overview += normalWhiteText.Sprintf(" %s\n", trimmedLine)
	}

	overview += boldWhiteText.Sprint(commandUsageExamples)
	for _, line := range examples {
		trimmedLine := strings.TrimSpace(line)
		overview += normalWhiteText.Sprintf("  $ %s\n", trimmedLine)
	}

	return overview
}

func wrapUsage(usage string, max, indent int) string {
	var final string
	wrapped := text.WrapSoft(usage, max-indent)
	lines := strings.Split(wrapped, "\n")
	if len(lines) > 1 {
		for i, line := range lines {
			if i > 0 {
				final += strings.Repeat(" ", indent)
			}
			final += line + "\n"
		}
	} else {
		final += lines[0] + "\n"
	}

	return final
}

func unquoteUsage(flag *pflag.Flag) (name string, usage string) {
	usage = flag.Usage
	for i := 0; i < len(usage); i++ {
		if usage[i] == '`' {
			for j := i + 1; j < len(usage); j++ {
				if usage[j] == '`' {
					name = usage[i+1 : j]
					usage = usage[:i] + name + usage[j+1:]
					return name, usage
				}
			}
			break // Only one back quote; use type name.
		}
	}

	name = flag.Value.Type()
	switch name {
	case "float64":
		name = "float"
	case "int64":
		name = "int"
	case "uint64":
		name = "uint"
	case "stringSlice":
		name = "[]string"
	case "intSlice":
		name = "[]int"
	case "uintSlice":
		name = "[]uint"
	case "boolSlice":
		name = "[]bool"
	}

	return
}

func printHeader() {
	boldGreenFormat := text.Colors{text.Bold, text.FgHiGreen}

	shields := []string{
		text.Colors{text.FgHiWhite, text.BgGreen}.Sprint(" version ") +
			text.Colors{text.Bold, text.FgHiGreen}.Sprintf(" %s ", appVersion),
		text.Colors{text.FgHiWhite, text.BgGreen}.Sprint(" license ") +
			text.Colors{text.Bold, text.FgHiGreen}.Sprintf(" %s ", appLicense),
		text.Colors{text.FgHiWhite, text.BgGreen}.Sprint(" copyright ") +
			text.Colors{text.Bold, text.FgHiGreen}.Sprintf(" %s ", appCopyright),
	}

	var centeredLogo string
	for _, logoLine := range logo {
		formattedLine := boldGreenFormat.Sprint(logoLine)
		centeredLogo += text.AlignCenter.Apply(formattedLine, appViewWidth) + "\n"
	}

	fmt.Printf("\n%s\n %s\n\n", centeredLogo, strings.Join(shields, " "))
}

func printDescriptor(cmd *cobra.Command) {
	var (
		maxLen    int
		flagLines []string
	)
	fmt.Println(cmd.Example)
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

		varType, usage := unquoteUsage(flag)
		line += " " + varType

		line += "\x00"
		if len(line) > maxLen {
			maxLen = len(line)
		}

		line += usage
		flagLines = append(flagLines, line)
	})

	if len(flagLines) > 0 {
		fmt.Println(text.Colors{text.Bold, text.FgHiWhite}.Sprint(commandOptions))
		for _, line := range flagLines {
			sIdx := strings.Index(line, "\x00")
			spacing := strings.Repeat(" ", maxLen-sIdx)
			concatenated := line[:sIdx] + spacing + " " + wrapUsage(line[sIdx+1:], appViewWidth, maxLen+1)
			fmt.Print(text.Colors{text.FgHiWhite}.Sprint(concatenated))
		}
	}
	fmt.Println()
}

func getTableWriter(columns ...string) table.Writer {
	var tableRow []interface{}
	for _, col := range columns {
		tableRow = append(tableRow, col)
	}

	yearColumnWidth := 6
	moneyColumnMaxWidth := (appViewWidth - yearColumnWidth - 17) / 3
	inflationColumnMaxWidth := (appViewWidth - yearColumnWidth - 14) / 2

	t := table.NewWriter()
	t.SetAllowedRowLength(appViewWidth)
	t.AppendHeader(tableRow)
	t.SetStyle(table.Style{
		Name: "Assist",
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
			Name:        tableColumnYear,
			Align:       text.AlignCenter,
			AlignFooter: text.AlignLeft,
			AlignHeader: text.AlignCenter,
			WidthMin:    yearColumnWidth,
			WidthMax:    yearColumnWidth,
		},
		{
			Name:        tableColumnMonth,
			Align:       text.AlignCenter,
			AlignFooter: text.AlignLeft,
			AlignHeader: text.AlignCenter,
			WidthMin:    yearColumnWidth,
			WidthMax:    yearColumnWidth,
		},
		{
			Name:        tableColumnExpenses,
			Align:       text.AlignCenter,
			AlignFooter: text.AlignLeft,
			AlignHeader: text.AlignCenter,
			WidthMin:    moneyColumnMaxWidth,
			WidthMax:    moneyColumnMaxWidth,
		},
		{
			Name:        tableColumnInvestments,
			Align:       text.AlignCenter,
			AlignFooter: text.AlignLeft,
			AlignHeader: text.AlignCenter,
			WidthMin:    moneyColumnMaxWidth,
			WidthMax:    moneyColumnMaxWidth,
		},
		{
			Name:        tableColumnInterestIncome,
			Align:       text.AlignCenter,
			AlignFooter: text.AlignLeft,
			AlignHeader: text.AlignCenter,
			WidthMin:    moneyColumnMaxWidth,
			WidthMax:    moneyColumnMaxWidth,
		},
		{
			Name:        tableColumnTotalSavings,
			Align:       text.AlignCenter,
			AlignFooter: text.AlignLeft,
			AlignHeader: text.AlignCenter,
			WidthMin:    moneyColumnMaxWidth,
			WidthMax:    moneyColumnMaxWidth,
		},
		{
			Name:        tableColumnInflationInitial,
			Align:       text.AlignCenter,
			AlignFooter: text.AlignLeft,
			AlignHeader: text.AlignCenter,
			WidthMin:    inflationColumnMaxWidth,
			WidthMax:    inflationColumnMaxWidth,
		},
		{
			Name:        tableColumnInflationEquivalent,
			Align:       text.AlignCenter,
			AlignFooter: text.AlignLeft,
			AlignHeader: text.AlignCenter,
			WidthMin:    inflationColumnMaxWidth,
			WidthMax:    inflationColumnMaxWidth,
		},
	})

	return t
}

func validateDetailedOption(value string) error {
	if value != commandOptionEmpty &&
		value != commandOptionDetailedYearly &&
		value != commandOptionDetailedMonthly {
		return fmt.Errorf("invalid argument value: detailed = %q; available options: %q, %q",
			value, commandOptionDetailedYearly, commandOptionDetailedMonthly)
	}
	return nil
}

func getTaskOverview(title, overview string) string {
	boldWhiteText := text.Colors{text.Bold, text.FgHiWhite}
	normalWhiteText := text.Colors{text.FgHiWhite}

	upperCaseTitle := text.FormatUpper.Apply(title)
	formattedTitle := boldWhiteText.Sprintf(" %s", upperCaseTitle)
	wrappedTask := text.WrapSoft(overview, appViewWidth-2)
	taskOverview := formattedTitle + "\n\n"

	for _, line := range strings.Split(wrappedTask, "\n") {
		trimmedLine := strings.TrimSpace(line)
		taskOverview += normalWhiteText.Sprintf(" %s\n", trimmedLine)
	}

	return taskOverview
}

func getFloat64(cmd *cobra.Command, name string) float64 {
	value, err := cmd.Flags().GetFloat64(name)
	if err != nil {
		_ = cmd.Help()
		log.Fatal(err)
	}
	return value
}

func getUint8(cmd *cobra.Command, name string) uint8 {
	value, err := cmd.Flags().GetUint8(name)
	if err != nil {
		_ = cmd.Help()
		log.Fatal(err)
	}
	return value
}

func getString(cmd *cobra.Command, name string) string {
	value, err := cmd.Flags().GetString(name)
	if err != nil {
		_ = cmd.Help()
		log.Fatal(err)
	}
	return value
}

func getBool(cmd *cobra.Command, name string) bool {
	value, err := cmd.Flags().GetBool(name)
	if err != nil {
		_ = cmd.Help()
		log.Fatal(err)
	}
	return value
}
