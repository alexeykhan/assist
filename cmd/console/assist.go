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
	"github.com/spf13/cobra"
)

var assistCmd = &cobra.Command{
	Use:          "assistCmd",
	SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
	Example: commandOverview(
		"Персональный ассистент для планирования личных финансов",
		"Используйте команду `decomposeCmd`, чтобы декомпозировать финансовую цель, например, "+
			"узнать минимальные необходимые условия для достижения вашей цели к конкретному сроку. "+
			"Подход: от желаемого результата.\n\n"+
			"Используйте команду `calculate`, чтобы посмотреть, каких результатов можно достигнуть "+
			"за указанный период, если соблюдать конкретные условия. Подход: от текущей ситуации.",
		[]string{
			"./bin/assist decompose --help",
			"./bin/assist calculate --help",
			"./bin/assist --help",
		},
	),
}
