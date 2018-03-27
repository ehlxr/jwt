// Copyright © 2018 ehlxr <ehlxr.me@gmail.com>
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

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/cobra"
)

// Verify a token and output the claims.  This is a great example
// of how to verify and view a token.
func VerifyToken(cmd *cobra.Command) error {
	flagToken := cmd.LocalFlags().Lookup("token").Value.String()
	flagKey := cmd.LocalFlags().Lookup("key").Value.String()
	flagCompact := promptCompact()
	flagDebug := promptDebug()

	flagAlg, err := promptAlg()
	if err != nil {
		return fmt.Errorf("Prompt flagAlg failed %v\n", err)
	}

	// get the token
	tokData, err := loadData(flagToken)
	if err != nil {
		return fmt.Errorf("Couldn't read token: %v", err)
	}

	// trim possible whitespace from token
	tokData = regexp.MustCompile(`\s*$`).ReplaceAll(tokData, []byte{})
	if flagDebug {
		fmt.Fprintf(os.Stderr, "Token len: %v bytes\n", len(tokData))
	}

	// Parse the token.  Load the key from command line option
	token, err := jwt.Parse(string(tokData), func(t *jwt.Token) (interface{}, error) {
		data, err := loadData(flagKey)
		if err != nil {
			return nil, err
		}
		if isEs(flagAlg) {
			return jwt.ParseECPublicKeyFromPEM(data)
		} else if isRs(flagAlg) {
			return jwt.ParseRSAPublicKeyFromPEM(data)
		}
		return data, nil
	})

	// Print some debug data
	if flagDebug && token != nil {
		fmt.Fprintf(os.Stderr, "Header:\n%v\n", token.Header)
		fmt.Fprintf(os.Stderr, "Claims:\n%v\n", token.Claims)
	}

	// Print an error if we can't parse for some reason
	if err != nil {
		return fmt.Errorf("Couldn't parse token: %v", err)
	}

	// Is token invalid?
	if !token.Valid {
		return fmt.Errorf("Token is invalid")
	}

	// Print the token details
	if err := printJSON(token.Claims, flagCompact); err != nil {
		return fmt.Errorf("Failed to output claims: %v", err)
	}

	return nil
}
