// Copyright © 2020 xrv <xrg@live.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"

	"github.com/ehlxr/jwt/token"
	"github.com/spf13/cobra"
)

// signCmd represents the sign command
var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "JWT 签名",
	Long: `
签名 JWT token 并复制到剪切板
标记 * 号的 flag 为必须项`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := token.SignToken(cmd); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(signCmd)

	signCmd.Flags().StringP("data", "d", "", "* path or json to claims object to sign, '-' to read from clipboard, or '+' to use only -claim args")
	signCmd.Flags().StringP("key", "k", "", "* path of keyfile or key argument")
	signCmd.Flags().VarP(make(token.ArgList), "claims", "c", "add additional claims. may be used more than once")
	signCmd.Flags().VarP(make(token.ArgList), "header", "H", "add additional header params. may be used more than once")
}
