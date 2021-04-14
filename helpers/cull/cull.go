package cull

import (
	"bufio"
	"fmt"
	"os"

	"github.com/audioo/goseek/helpers/cli"
	"github.com/jung-kurt/gofpdf"
)

// Run ...
func Run() {
	fmt.Println()
	fmt.Println("   Leave field empty to omit.")
	fname, err := generatePdf()
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println("   Output to " + fname + ".pdf")
}

func input(question string, scanner *bufio.Scanner, pdf *gofpdf.Fpdf) {
	fmt.Print("   " + question)
	res := cli.ScanIt(scanner)
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
		"./images/icon.png",
		100.5, 4.5,
		10, 10,
		false,
		gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true},
		0,
		"",
	)

	// CellFormat(width, height, text, border, position after, align, fill, link, linkStr)
	pdf.CellFormat(190, 0, "GOSEEK           RESULT", "0", 0, "CM", false, 0, "")
	pdf.Line(0, 18, 380, 19)
	pdf.CellFormat(190, 200, "", "0", 0, "TL", false, 0, "")
	pdf.SetFontSize(10)

	// Spacers

	pdf.Write(5, "\n")
	pdf.Write(5, "\n")
	pdf.Write(5, "\n")

	// Spacers end
	fmt.Println()

	fmt.Print("   File Name: ")
	fname := cli.ScanIt(scanner)

	fmt.Println()

	input("Legal Name: ", scanner, pdf)
	input("Date of Birth: ", scanner, pdf)
	input("Preferred Pronoun(s): ", scanner, pdf)
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
	input("Vehicle Information: ", scanner, pdf)
	input("Health Issues: ", scanner, pdf)
	input("IP Addresses: ", scanner, pdf)
	input("Embarrassing Information: ", scanner, pdf)
	input("Bank Information: ", scanner, pdf)
	input("Potential Liabilities: ", scanner, pdf)
	input("Description: ", scanner, pdf)

	return fname, pdf.OutputFileAndClose(fname + ".pdf")

}
