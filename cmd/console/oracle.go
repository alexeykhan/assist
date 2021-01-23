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
)

const (
	appVersion   = "0.1.0"
	appUpdated   = "2021-01-23"
	appCopyright = "alexeykhan"
	appLicense   = "MIT"
	appViewWidth = 78
)

var logo = []string{
	`  ooooooo   oooooooooo       o       oooooooo8 ooooo       ooooooooooo`,
	`o888   888o  888    888     888    o888     88  888         888    88 `,
	`888     888  888oooo88     8  88   888          888         888ooo8   `,
	`888o   o888  888  88o     8oooo88  888o     oo  888      o  888    oo `,
	`  88ooo88   o888o  88o8 o88o  o888o 888oooo88  o888ooooo88 o888ooo8888`,
}

var oracle = &cobra.Command{
	Use:   "oracle",
	SilenceUsage: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		printHeader()
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from oracle")
	},
}
