package menu

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/audioo/goseek/helpers/cli"
	"github.com/audioo/goseek/helpers/cull"
	"github.com/audioo/goseek/helpers/ent"
	"github.com/audioo/goseek/helpers/http"
	"github.com/audioo/goseek/helpers/userlookup/run"
	"github.com/audioo/goseek/helpers/vehicle"
)

// MainMenu ...
func MainMenu() {
	cli.Clear()
	cli.Banner()
	cli.Dispban("Select an option")
	fmt.Println()
	fmt.Println(cli.Dispopw("1", "Username Lookup"))
	fmt.Println(cli.Dispopw("2", "Cull"))
	fmt.Println(cli.Dispopw("3", "IP Lookup"))
	fmt.Println(cli.Dispopw("4", "License Plate Lookup (Coming soon...)"))
	fmt.Println(cli.Dispopw("5", "VIN Lookup"))
	fmt.Println(cli.Dispopw("X", "Exit"))
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println()
	fmt.Print("   >> ")
	res := cli.ScanIt(scanner)
	if res == "1" {
		if http.Connected() {
			usernameLookup()
		} else {
			fmt.Println()
			fmt.Println("   You need to have an internet connection in order to use this function.")
			fmt.Println("   Press enter to continue...")
			cli.ScanIt(scanner)
			MainMenu()
		}
	} else if res == "2" {
		cli.Clear()
		cli.Banner()
		cli.Dispban("Cull")
		cull.Run()
		fmt.Println()
		cli.Dispban("Press enter to return to the main menu...")
		cli.ScanIt(scanner)
		MainMenu()
	} else if res == "3" {
		if http.Connected() {
			iplookup()
		} else {
			fmt.Println()
			fmt.Println("   You need to have an internet connection in order to use this function.")
			fmt.Println("   Press enter to continue...")
			cli.ScanIt(scanner)
			MainMenu()
		}
	} else if res == "5" {
		if http.Connected() {
			vin()
		} else {
			fmt.Println()
			fmt.Println("   You need to have an internet connection in order to use this function.")
			fmt.Println("   Press enter to continue...")
			cli.ScanIt(scanner)
			MainMenu()
		}
	} else if res == "X" || res == "x" {
		os.Exit(0)
	} else {
		MainMenu()
	}
}

func usernameLookup() {
	cli.Clear()
	cli.Banner()
	cli.Dispban("Username Lookup")
	fmt.Println()
	scanner := bufio.NewScanner(os.Stdin)
	cli.Dispban("Input Username")
	fmt.Print("   >> ")
	userres := cli.ScanIt(scanner)
	fmt.Println()

	cli.Dispban("Write to file? [y/n]")
	fmt.Print("   >> ")
	write := cli.ScanIt(scanner)

	if write == "y" {
		cli.Banner()
		fmt.Println(cli.Dispopg("GREEN", "ACCOUNT EXISTS"))
		fmt.Println(cli.Dispop("RED", "ACCOUNT DOES NOT EXIST"))
		fmt.Println("")
		run.SendSeeker(userres, true)
		fmt.Println("")
		cli.Dispban("Go-Seek Process Complete :: Results written to " + userres + ".txt")
	} else {
		cli.Banner()
		fmt.Println(cli.Dispopg("GREEN", "ACCOUNT EXISTS"))
		fmt.Println(cli.Dispop("RED", "ACCOUNT DOES NOT EXIST"))
		fmt.Println("")
		run.SendSeeker(userres, false)
		fmt.Println("")
		cli.Dispban("Go-Seek Process Complete | Results have not been written")
	}
	fmt.Println("")
	cli.Dispban("Press enter to return to the main menu...")
	cli.ScanIt(scanner)
	MainMenu()
}

func iplookup() {
	scanner := bufio.NewScanner(os.Stdin)
	cli.Clear()
	cli.Banner()
	cli.Dispban("IP Lookup")
	fmt.Println("")
	fmt.Println("   Enter IP Address")
	fmt.Print("   >> ")
	ipres := cli.ScanIt(scanner)
	fmt.Println("")
	var resp = http.GetReq("http://ip-api.com/json/" + ipres + "?fields=31162361")
	var s = new(ent.IPAddress)
	json.Unmarshal([]byte(resp), &s)
	fmt.Println("   Status: " + s.Status)
	fmt.Println("   Continent: " + s.Continent)
	fmt.Println("   Country: " + s.Country)
	fmt.Println("   Region Name: " + s.RegionName)
	fmt.Println("   City: " + s.City)
	fmt.Println("   District: " + s.District)
	fmt.Println("   Zip: " + s.Zip)
	fmt.Println("   Lat: " + strconv.FormatFloat(s.Lat, 'f', -1, 64))
	fmt.Println("   Lon: " + strconv.FormatFloat(s.Lon, 'f', -1, 64))
	fmt.Println("   Time Zone: " + s.Timezone)
	fmt.Println("   Currency: " + s.Currency)
	fmt.Println("   ISP: " + s.Isp)
	fmt.Println("   Org: " + s.Org)
	fmt.Println("   As: " + s.As)
	fmt.Println("   Asname: " + s.Asname)
	fmt.Println("   Reverse: " + s.Reverse)
	fmt.Println("   Mobile: " + strconv.FormatBool(s.Mobile))
	fmt.Println("   Proxy: " + strconv.FormatBool(s.Proxy))
	fmt.Println("   Hosting: " + strconv.FormatBool(s.Hosting))
	fmt.Println("")
	cli.Dispban("Press enter to return to the main menu...")
	cli.ScanIt(scanner)
	MainMenu()
}

/*
func plate() {
	scanner := bufio.NewScanner(os.Stdin)
	cli.Clear()
	cli.Banner()
	cli.Dispban("License Plate Lookup")
	fmt.Println("")
	fmt.Println("   Enter License Plate")
	fmt.Print("   >> ")
	fmt.Println("")
	fmt.Println("   Enter State Abbrevation (CA, NY, etc...)")
	fmt.Print("   >> ")
	plate := cli.ScanIt(scanner)
}
*/
func vin() {
	scanner := bufio.NewScanner(os.Stdin)
	cli.Clear()
	cli.Banner()
	cli.Dispban("VIN Lookup")
	fmt.Println("")
	fmt.Println("   Enter VIN")
	fmt.Print("   >> ")
	vin := cli.ScanIt(scanner)
	fmt.Println()
	vehicle.VinLookup(vin, 1)
	fmt.Println()
	cli.Dispban("Press enter to return to the main menu...")
	cli.ScanIt(scanner)
	MainMenu()
}
