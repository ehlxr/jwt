/*
Copyright © 2019 ehlxr <ehlxr.me@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

var bannerBase64 = "DQogICBfXyAgICAgX18gICAgIF9fICAgICBfX19fX18gIA0KICAvXCBcICAgL1wgXCAgXyBcIFwgICAvXF9fICBfXCANCiBfXF9cIFwgIFwgXCBcLyAiLlwgXCAgXC9fL1wgXC8gDQovXF9fX19fXCAgXCBcX18vIi5+XF9cICAgIFwgXF9cIA0KXC9fX19fXy8gICBcL18vICAgXC9fLyAgICAgXC9fLyANCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgDQo="
var versionTpl = `%s

Name: jwt
Version: %s
BuildTime: %s
GitCommit: %s
GoVersion: %s
`

var (
	Version   string
	BuildTime string
	GitCommit string
	GoVersion string
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Long: `
Print version of jwt`,
	Run: func(cmd *cobra.Command, args []string) {
		banner, _ := base64.StdEncoding.DecodeString(bannerBase64)
		fmt.Printf(versionTpl, banner, Version, BuildTime, GitCommit, GoVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
