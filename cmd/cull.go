/*
Copyright © 2021 Audio <hyperaudio@protonmail.com>

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

// Package cmd ...
package cmd

import (
	"fmt"

	"github.com/audioo/goseek/helpers/cli"
	"github.com/audioo/goseek/helpers/cull"
	"github.com/spf13/cobra"
)

// cullCmd represents the cull command
var cullCmd = &cobra.Command{
	Use:   "cull",
	Short: "Information Narrowing",
	Long:  `Compile information on an individual and output the results to a PDF file`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.Clear()
		cli.Banner()
		cli.Dispban("Cull")
		cull.Run()
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(cullCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cullCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cullCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}