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
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type template struct {
	title, about string
	examples     []string
}

func (t template) normalize() string {
	boldWhiteText := text.Colors{text.Bold, text.FgHiWhite}
	normalWhiteText := text.Colors{text.FgHiWhite}

	var examples string
	upperCaseTitle := text.FormatUpper.Apply(t.title)
	formattedTitle := boldWhiteText.Sprintf(" %s", upperCaseTitle)
	examples += formattedTitle + "\n\n"

	wrappedAbout := text.WrapSoft(t.about, appViewWidth-2)
	for _, line := range strings.Split(wrappedAbout, "\n") {
		trimmedLine := strings.TrimSpace(line)
		examples += normalWhiteText.Sprintf(" %s\n", trimmedLine)
	}

	examples += boldWhiteText.Sprint("\n Примеры использования:\n")
	for _, line := range t.examples {
		trimmedLine := strings.TrimSpace(line)
		examples += normalWhiteText.Sprintf("  $ %s\n", trimmedLine)
	}

	return examples
}

func example(title, about string, examples []string) string {
	return template{title: title, about: about, examples: examples}.normalize()
}

func wrapUsage(usage string, max, indent int) string {
	var final string
	wrapped := text.WrapSoft(usage, max-indent)
	lines := strings.Split(wrapped, "\n")
	if len(lines) > 1 {
		for i, line := range lines {
			if i == 0 {
				final += line + "\n"
				continue
			}
			final += strings.Repeat(" ", indent) + line
		}
	} else {
		final += lines[0]
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
		text.Colors{text.FgHiWhite, text.BgGreen}.Sprint(" license ") +
			text.Colors{text.Bold, text.FgHiGreen}.Sprintf(" %s ", appLicense),
		text.Colors{text.FgHiWhite, text.BgGreen}.Sprint(" copyright ") +
			text.Colors{text.Bold, text.FgHiGreen}.Sprintf(" %s ", appCopyright),
		text.Colors{text.FgHiWhite, text.BgGreen}.Sprint(" version ") +
			text.Colors{text.Bold, text.FgHiGreen}.Sprintf(" %s ", appVersion),
		text.Colors{text.FgHiWhite, text.BgGreen}.Sprint(" updated ") +
			text.Colors{text.Bold, text.FgHiGreen}.Sprintf(" %s ", appUpdated),
	}

	shieldsText := strings.Join(shields, " ")
	centeredShields := text.AlignCenter.Apply(shieldsText, appViewWidth)

	var logoText string
	for _, logoLine := range logo {
		formattedLine := boldGreenFormat.Sprint(logoLine)
		logoText += formattedLine + "\n"
	}
	var centeredLogo string
	for _, logoLine := range logo {
		formattedLine := boldGreenFormat.Sprint(logoLine)
		centeredLogo += text.AlignCenter.Apply(formattedLine, appViewWidth) + "\n"
	}

	fmt.Printf("\n%s\n%s\n\n", centeredLogo, centeredShields)
}

func printDescriptor(cmd *cobra.Command) {
	var maxLen int
	var flagLines []string
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

		varname, usage := unquoteUsage(flag)
		line += " " + varname

		line += "\x00"
		if len(line) > maxLen {
			maxLen = len(line)
		}

		line += usage
		flagLines = append(flagLines, line)
	})

	if len(flagLines) > 0 {
		fmt.Println(text.Colors{text.Bold, text.FgHiWhite}.Sprint(" Параметры и опции команды:"))
		for _, line := range flagLines {
			sidx := strings.Index(line, "\x00")
			spacing := strings.Repeat(" ", maxLen-sidx)
			concatenated := line[:sidx] + spacing + " " + wrapUsage(line[sidx+1:], appViewWidth, maxLen+1)
			fmt.Println(text.Colors{text.FgHiWhite}.Sprint(concatenated))
		}
	}
	fmt.Println()
}

func getTableWriter() table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	return t
}

func getBool(cmd *cobra.Command, name string) bool {
	value, err := cmd.Flags().GetBool(name)
	if err != nil {
		_ = cmd.Help()
		log.Fatal(err)
	}
	return value
}

func getFloat32(cmd *cobra.Command, name string) float32 {
	value, err := cmd.Flags().GetFloat32(name)
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
