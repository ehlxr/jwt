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

package token

import (
	"fmt"
	"os"
	"regexp"

	"github.com/dgrijalva/jwt-go"

	"github.com/spf13/cobra"
)

// showToken pretty-prints the token on the command line.
func ShowToken(cmd *cobra.Command) error {
	flagToken := cmd.LocalFlags().Lookup("token").Value.String()
	flagDebug := promptDebug()
	flagCompact := promptCompact()

	// get the token
	tokData, err := loadData(flagToken)
	if err != nil {
		return fmt.Errorf("couldn't read token: %v", err)
	}

	// trim possible whitespace from token
	tokData = regexp.MustCompile(`\s*$`).ReplaceAll(tokData, []byte{})
	if flagDebug {
		_, _ = fmt.Fprintf(os.Stderr, "Token len: %v bytes\n", len(tokData))
	}

	token, err := jwt.Parse(string(tokData), nil)
	if token == nil {
		return fmt.Errorf("malformed token: %v", err)
	}

	// Print the token details
	fmt.Println("Header:")
	if err := printJSON(token.Header, flagCompact); err != nil {
		return fmt.Errorf("failed to output header: %v", err)
	}

	fmt.Println("Claims:")
	if err := printJSON(token.Claims, flagCompact); err != nil {
		return fmt.Errorf("failed to output claims: %v", err)
	}

	return nil
}
