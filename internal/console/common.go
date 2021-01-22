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

package console

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// GetTableWriter get standard table
func GetTableWriter() table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	return t
}

// GetStringParameter - get string parameter from command line
func GetStringParameter(cmd *cobra.Command, parameterName string) string {
	parameterValue, err := cmd.Flags().GetString(parameterName)
	if err != nil {
		log.Fatal(err.Error())
	}
	return parameterValue
}

// GetBool - get boolean parameter from command line
func GetBool(cmd *cobra.Command, name string) bool {
	value, err := cmd.Flags().GetBool(name)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

// GetInt32 - get float32 parameter from command line
func GetFloat32(cmd *cobra.Command, name string) float32 {
	value, err := cmd.Flags().GetFloat32(name)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

// GetUint8 - get integer parameter from command line
func GetUint8(cmd *cobra.Command, name string) uint8 {
	value, err := cmd.Flags().GetUint8(name)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

// GetUint64 - get integer parameter from command line
func GetUint64(cmd *cobra.Command, name string) uint64 {
	value, err := cmd.Flags().GetUint64(name)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

// GetIntParameter - get integer parameter from command line
func GetIntParameter(cmd *cobra.Command, parameterName string) int {
	parameterValue, err := cmd.Flags().GetInt(parameterName)
	if err != nil {
		log.Fatal(err.Error())
	}
	return parameterValue
}

// GetStringSliceParameter - get slice of string parameter from command line
func GetStringSliceParameter(cmd *cobra.Command, parameterName string) []string {
	parameterValue, err := cmd.Flags().GetStringSlice(parameterName)
	if err != nil {
		log.Fatal(err.Error())
	}
	return parameterValue
}

// GetStringMapParameter - get map of strings parameter from command line
func GetStringMapParameter(cmd *cobra.Command, parameterName string) map[string]string {
	parameterValue, err := cmd.Flags().GetStringToString(parameterName)
	if err != nil {
		log.Fatal(err.Error())
	}
	return parameterValue
}

// PrintJSON prints JSON to console
func PrintJSON(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "    ")
	fmt.Println(string(b))
}

// AskForConfirmation Спрашивает подтверждение
func AskForConfirmation(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))
		if response == "y" || response == "yes" {
			return true
		}

		if response == "n" || response == "no" {
			fmt.Println("Aborted")
			return false
		}

		fmt.Printf("Cannot recognize %q. Try again\n", response)
	}
}

// SanitizeString убирает всё перевод строк
func SanitizeString(s string) string {
	if s == "" {
		return s
	}

	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "\r", "")

	return strings.TrimSpace(s)
}
