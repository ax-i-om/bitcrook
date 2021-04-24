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

	"github.com/audioo/goseek/internal/cli"
	"github.com/audioo/goseek/pkg/opsec"
	"github.com/spf13/cobra"
)

// genidenCmd represents the geniden command
var genidenCmd = &cobra.Command{
	Use:   "geniden [gender(m/f/n)]",
	Short: "Generate a fake identity",
	Long:  `Generate a fake identity including full name, location, date of birth, username, and password.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Usage()
		} else {
			iden := opsec.Identity(args[0])
			cli.Banner()
			cli.Dispban("Identity Generator")
			fmt.Println("")
			fmt.Println("   Name: " + iden.Name)
			fmt.Println("   Location: " + iden.Location)
			fmt.Println("   Date of Birth: " + iden.Dob)
			fmt.Println("   Username: " + iden.Username)
			fmt.Println("   Password: " + iden.Password)
		}
	},
}

func init() {
	rootCmd.AddCommand(genidenCmd)
}
