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
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/atotto/clipboard"

	"github.com/manifoldco/promptui"
)

// Helper func:  Read input from specified file or string
func loadData(p string) ([]byte, error) {
	if p == "" {
		return nil, fmt.Errorf("No path or arg specified")
	}

	var rdr io.Reader
	if p == "-" {
		t, _ := clipboard.ReadAll()
		fmt.Printf("read data from clipboard: %s\n", t)
		rdr = strings.NewReader(t)
	} else if p == "+" {
		return []byte("{}"), nil
	} else {
		is, path := isPath(p)
		if is {
			if f, err := os.Open(path); err == nil {
				rdr = f
				defer f.Close()
			} else {
				return nil, err
			}
		} else {
			rdr = strings.NewReader(p)
		}
	}
	return ioutil.ReadAll(rdr)
}

func isPath(path string) (bool, string) {
	absPath, err := filepath.Abs(path)
	_, err = os.Stat(absPath)
	if err == nil {
		return true, absPath
	}
	return false, ""
}

// Print a json object in accordance with the prophecy (or the command line options)
func printJSON(j interface{}, flagCompact bool) error {
	var out []byte
	var err error

	if flagCompact == false {
		out, err = json.MarshalIndent(j, "", "    ")
	} else {
		out, err = json.Marshal(j)
	}

	if err == nil {
		fmt.Println(string(out))
	}

	return err
}

func isEs(flagAlg string) bool {
	return strings.HasPrefix(flagAlg, "ES")
}

func isRs(flagAlg string) bool {
	return strings.HasPrefix(flagAlg, "RS")
}

func promptAlg() (string, error) {
	prompt := promptui.SelectWithAdd{
		Label:    "select a signing algorithm identifier",
		Items:    []string{"HS256", "RS256"},
		AddLabel: "Other",
	}
	_, flagAlg, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("Prompt flagAlg failed %v\n", err)
	}

	return flagAlg, nil
}

func promptDebug() bool {
	prompt := promptui.Prompt{
		Label:     "print out all kinds of debug data",
		IsConfirm: true,
	}

	if _, err := prompt.Run(); err != nil {
		return false
	}
	return true
}

func promptCompact() bool {
	prompt := promptui.Prompt{
		Label:     "output compact JSON",
		IsConfirm: true,
	}

	if _, err := prompt.Run(); err != nil {
		return false
	}
	return true
}

type ArgList map[string]string

func (l ArgList) String() string {
	data, _ := json.Marshal(l)
	return string(data)
}

func (l ArgList) Set(arg string) error {
	parts := strings.SplitN(arg, "=", 2)
	if len(parts) != 2 {
		return fmt.Errorf("Invalid argument '%v'.  Must use format 'key=value'. %v", arg, parts)
	}
	l[parts[0]] = parts[1]
	return nil
}

func (l ArgList) Type() string {
	return "argList"
}
