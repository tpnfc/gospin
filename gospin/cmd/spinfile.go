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
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/m1/gospin"
	"github.com/spf13/cobra"
)

// spinfileCmd represents the spinfile command
var spinfileCmd = &cobra.Command{
	Use:   "spinfile",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := gospin.Config{
			StartChar:     startChar,
			EndChar:       endChar,
			DelimiterChar: delimiterChar,
			EscapeChar:    escapeChar,
		}
		spinner := gospin.New(&cfg)
		str, e := openFile(filename)
		if e != nil {
			log.Println(e)
			return
		}
		if times == 1 {
			spun, err := spinner.Spin(str)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			fmt.Println(spun)
			return
		}

		spun, err := spinner.SpinN(str, times)
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
	rootCmd.AddCommand(spinfileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// spinfileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// spinfileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	spinfileCmd.Flags().StringVarP(&filename, "filename", "f", "seed.txt", "path to the file")

}

var filename string

//openFile
func openFile(name string) (data string, err error) {
	path, _ := os.Getwd()
	x := filepath.Join(path, filename)
	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile(x)
	if err != nil {
		log.Fatal(err)
	}
	// Convert []byte to string and print to screen
	text := string(content)
	return text, err
}
