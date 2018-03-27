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
	"encoding/json"
	"fmt"
	"os"

	"github.com/atotto/clipboard"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/spf13/cobra"
)

// Create, sign, and output a token.  This is a great, simple example of
// how to use this library to create and sign a token.
func SignToken(cmd *cobra.Command) error {
	flagData := cmd.LocalFlags().Lookup("data").Value.String()
	flagClaims := cmd.LocalFlags().Lookup("claims").Value.(ArgList)
	flagHead := cmd.LocalFlags().Lookup("header").Value.(ArgList)
	flagKey := cmd.LocalFlags().Lookup("key").Value.String()
	flagDebug := promptDebug()

	flagAlg, err := promptAlg()
	if err != nil {
		return fmt.Errorf("Prompt flagAlg failed %v\n", err)
	}

	// get the token data from command line arguments
	tokData, err := loadData(flagData)
	if err != nil {
		return fmt.Errorf("Couldn't read data: %v", err)
	} else if flagDebug {
		fmt.Fprintf(os.Stderr, "Token len: %v bytes\n", len(tokData))
		fmt.Fprintf(os.Stderr, "Token data: %v \n", string(tokData))
	}

	// parse the JSON of the claims
	var claims jwt.MapClaims
	if err := json.Unmarshal(tokData, &claims); err != nil {
		return fmt.Errorf("Couldn't parse claims JSON: %v", err)
	}

	// add command line claims
	if len(flagClaims) > 0 {
		for k, v := range flagClaims {
			claims[k] = v
		}
	}

	// get the key
	var key interface{}
	key, err = loadData(flagKey)
	if err != nil {
		return fmt.Errorf("Couldn't read key: %v", err)
	}

	// get the signing alg
	alg := jwt.GetSigningMethod(flagAlg)
	if alg == nil {
		return fmt.Errorf("Couldn't find signing method alg: %v", flagAlg)
	}

	// create a new token
	token := jwt.NewWithClaims(alg, claims)

	// add command line headers
	if len(flagHead) > 0 {
		for k, v := range flagHead {
			token.Header[k] = v
		}
	}

	if isEs(flagAlg) {
		if k, ok := key.([]byte); !ok {
			return fmt.Errorf("Couldn't convert key data to key")
		} else {
			key, err = jwt.ParseECPrivateKeyFromPEM(k)
			if err != nil {
				return err
			}
		}
	} else if isRs(flagAlg) {
		if k, ok := key.([]byte); !ok {
			return fmt.Errorf("Couldn't convert key data to key")
		} else {
			key, err = jwt.ParseRSAPrivateKeyFromPEM(k)
			if err != nil {
				return err
			}
		}
	}

	if out, err := token.SignedString(key); err == nil {
		fmt.Println(out)
		clipboard.WriteAll(out)
	} else {
		return fmt.Errorf("Error signing token: %v", err)
	}

	return nil
}
