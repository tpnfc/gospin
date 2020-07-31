/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/m1/gospin"
	"github.com/spf13/cobra"
)

var (
	startChar     string
	endChar       string
	escapeChar    string
	delimiterChar string
	times         int
)

// spintextCmd represents the spintext command
var spintextCmd = &cobra.Command{
	Use:   "spintext [text]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := gospin.Config{
			StartChar:     startChar,
			EndChar:       endChar,
			DelimiterChar: delimiterChar,
			EscapeChar:    escapeChar,
		}
		spinner := gospin.New(&cfg)
		if times == 1 {
			spun, err := spinner.Spin(args[0])
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			fmt.Println(spun)
			return
		}

		spun, err := spinner.SpinN(args[0], times)
		if err != nil {
			log.Panic(err)
			return
		}
		for key, value := range spun {
			fmt.Println(fmt.Sprintf("%v: %v", key, value))
		}
		js, err := json.Marshal(&spun)
		if err != nil {
			log.Panic(err)
			return
		}

		fmt.Println(string(js))
		return

	},
}

func init() {
	rootCmd.AddCommand(spintextCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// spintextCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// spintextCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	spintextCmd.PersistentFlags().StringVar(&startChar, "start", "{", "Start char for the spinning engine")
	spintextCmd.PersistentFlags().StringVar(&endChar, "end", "}", "End char for the spinning engine")
	spintextCmd.PersistentFlags().StringVar(&delimiterChar, "delimiter", "|", "Delimiter char")
	spintextCmd.PersistentFlags().StringVar(&escapeChar, "escape", "\\", "Escape char")
	spintextCmd.PersistentFlags().IntVar(&times, "times", 1, "How many articles to generate")
	// spintextCmd.Execute()
}
