/*
Copyright © 2023 <yoshiki.shino.tech@gmail.com>

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
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/yoshikipom/json-masking-go/masking"
)

// maskCmd represents the mask command
var maskCmd = &cobra.Command{
	Use:   "mask",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// masking information
		denyList, _ := cmd.Flags().GetStringArray("deny")
		useRegex, _ := cmd.Flags().GetBool("regex")
		format, _ := cmd.Flags().GetBool("format")
		m := masking.New(denyList, useRegex, format)

		// input json
		var json []byte
		if len(args) > 0 {
			json = []byte(args[0])
		} else {
			bytes, err := ioutil.ReadAll(os.Stdin)
			json = bytes
			if err != nil {
				fmt.Println("Error reading from stdin:", err)
				os.Exit(1)
			}
		}
		fmt.Printf("Original JSON: %v\n", string(json))

		replaced := m.Replace(json)
		fmt.Printf("%s\n", string(replaced))
	},
}

func init() {
	rootCmd.AddCommand(maskCmd)

	maskCmd.Flags().StringArrayP("deny", "d", []string{}, "deny key list")
	maskCmd.Flags().Bool("regex", false, "flag to use regex mode")
	maskCmd.Flags().Bool("format", false, "flag for formatting of output")
}
