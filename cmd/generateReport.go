/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"bufio"
	"fmt"
	"os"

	"github.com/jung-kurt/gofpdf"
	"github.com/spf13/cobra"
)

// generateReportCmd represents the generateReport command
var generateReportCmd = &cobra.Command{
	Use:   "generateReport",
	Short: "Generate a report via terminal",
	Long:  `Generate a report via terminal`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		fmt.Println("Leave field empty to omit.")
		fname, err := generatePdf()
		if err != nil {
			panic(err)
		}
		fmt.Println()
		fmt.Println("Output to " + fname + ".pdf")
	},
}

func init() {
	rootCmd.AddCommand(generateReportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateReportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateReportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func scanIt(s *bufio.Scanner) string {
	s.Scan()
	return s.Text()
}

func input(question string, scanner *bufio.Scanner, pdf *gofpdf.Fpdf) {
	fmt.Print(question)
	res := scanIt(scanner)
	if len(res) != 0 {
		pdf.Write(5, question+res+"\n")
	}
}

func generatePdf() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	// ImageOptions(src, x, y, width, height, flow, options, link, linkStr)
	pdf.ImageOptions(
		"./images/quasmall.png",
		100.8, 3,
		13, 13,
		false,
		gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true},
		0,
		"",
	)

	// CellFormat(width, height, text, border, position after, align, fill, link, linkStr)
	pdf.CellFormat(190, 0, "GOSEEK          RESULT", "0", 0, "CM", false, 0, "")
	pdf.Line(0, 18, 380, 19)
	pdf.CellFormat(190, 200, "", "0", 0, "TL", false, 0, "")
	pdf.SetFontSize(10)

	// Spacers

	pdf.Write(5, "\n")
	pdf.Write(5, "\n")
	pdf.Write(5, "\n")

	// Spacers end
	fmt.Println()

	fmt.Print("File Name: ")
	fname := scanIt(scanner)

	fmt.Println()

	input("Legal Name: ", scanner, pdf)
	input("Date of Birth: ", scanner, pdf)
	input("Preferred Pronoun(s)", scanner, pdf)
	input("Place of Birth: ", scanner, pdf)
	input("Usernames: ", scanner, pdf)
	input("Passwords: ", scanner, pdf)
	input("Aliases: ", scanner, pdf)
	input("Email Addresses: ", scanner, pdf)
	input("Phone Numbers: ", scanner, pdf)
	input("Addresses: ", scanner, pdf)
	input("Contacts: ", scanner, pdf)
	input("Social Media Links: ", scanner, pdf)
	input("Family Information: ", scanner, pdf)
	input("Group Affiliations: ", scanner, pdf)
	input("Personal Tragedies: ", scanner, pdf)
	input("Relationship Status: ", scanner, pdf)
	input("Job Information: ", scanner, pdf)
	input("Health Issues: ", scanner, pdf)
	input("IP Addresses: ", scanner, pdf)
	input("Embarrassing Information: ", scanner, pdf)
	input("Bank Information: ", scanner, pdf)
	input("Potential Liabilities: ", scanner, pdf)
	input("Description: ", scanner, pdf)

	return fname, pdf.OutputFileAndClose(fname + ".pdf")

}
