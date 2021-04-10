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
	fname, err := generatePdf()
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println("   Output to " + fname + ".pdf")
}

func generatePdf() (string, error) {
	var res string
	scanner := bufio.NewScanner(os.Stdin)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	// ImageOptions(src, x, y, width, height, flow, options, link, linkStr)
	pdf.ImageOptions(
		"./images/icon.png",
		104, 4.5,
		10, 10,
		false,
		gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true},
		0,
		"",
	)

	// CellFormat(width, height, text, border, position after, align, fill, link, linkStr)
	pdf.CellFormat(190, 0, "CREDCULL           RESULT", "0", 0, "CM", false, 0, "")
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

	fmt.Print("   Legal Name: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Legal Name: "+res+"\n")

	fmt.Print("   Date of Birth: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Date of Birth: "+res+"\n")

	fmt.Print("   Preferred Pronoun(s): ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Preferred Pronoun(s):"+res+"\n")

	fmt.Print("   Place of Birth: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Place of Birth: "+res+"\n")

	fmt.Print("   Known Usernames: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Known Usernames: "+res+"\n")

	fmt.Print("   Known Passwords: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Known Passwords: "+res+"\n")

	fmt.Print("   Known Nicknames: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Known Nicknames: "+res+"\n")

	fmt.Print("   Email Addresses: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Email Addresses: "+res+"\n")

	fmt.Print("   Phone Numbers: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Phone Numbers: "+res+"\n")

	fmt.Print("   Known Addresses: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Known Addresses: "+res+"\n")

	fmt.Print("   Known Contacts: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Known Contacts: "+res+"\n")

	fmt.Print("   Social Media Links: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Social Media Links: "+res+"\n")

	fmt.Print("   Family Tree Info: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Family Tree Info: "+res+"\n")

	fmt.Print("   Group Affiliations: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Group Affiliations: "+res+"\n")

	fmt.Print("   Personal Tragedies: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Personal Tragedies: "+res+"\n")

	fmt.Print("   Relationship Status: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Relationship Status: "+res+"\n")

	fmt.Print("   Children's Information: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Children's Information: "+res+"\n")

	fmt.Print("   Job Info: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Job Info: "+res+"\n")

	fmt.Print("   Health Issues: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Health Issues: "+res+"\n")

	fmt.Print("   IP Addresses: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "IP Addresses: "+res+"\n")

	fmt.Print("   Bluetooth IDs: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Bluetooth IDs: "+res+"\n")

	fmt.Print("   Known SSIDs: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Known SSIDs: "+res+"\n")

	fmt.Print("   Known IMEIs: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Known IMEIs: "+res+"\n")

	fmt.Print("   Known Social Footprint: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Known Social Footprint: "+res+"\n")

	fmt.Print("   Embarassing Data: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Embarassing Data: "+res+"\n")

	fmt.Print("   Estimated Yearly Income: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Estimated Yearly Income: "+res+"\n")

	fmt.Print("   Estimated Yearly Expenses: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Estimated Yearly Expenses: "+res+"\n")

	fmt.Print("   Standout Transactions: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Standout Transactions: "+res+"\n")

	fmt.Print("   Bank Information: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Bank Information: "+res+"\n")

	fmt.Print("   Potential Liabilities: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Potential Liabilities: "+res+"\n")

	fmt.Print("   Tax Information: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Tax Information: "+res+"\n")

	fmt.Print("   Credit Rating: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Credit Rating: "+res+"\n")

	fmt.Print("   Description: ")
	res = cli.ScanIt(scanner)
	pdf.Write(5, "Description: "+res+"\n")

	return fname, pdf.OutputFileAndClose(fname + ".pdf")

}
