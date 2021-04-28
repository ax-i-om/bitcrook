package cli

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

// Clear ...
func Clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Banner ...
func Banner() {
	Clear()
	fmt.Println("")

	fmt.Println("     ▄██████▄   ▄██████▄     ▄████████    ▄████████    ▄████████    ▄█   ▄█▄ ")
	fmt.Println("    ███    ███ ███    ███   ███    ███   ███    ███   ███    ███   ███ ▄███▀ ")
	fmt.Println("    ███    █▀  ███    ███   ███    █▀    ███    █▀    ███    █▀    ███▐██▀   ")
	fmt.Println("   ▄███        ███    ███   ███         ▄███▄▄▄      ▄███▄▄▄      ▄█████▀    ")
	fmt.Println("  ▀▀███ ████▄  ███    ███ ▀███████████ ▀▀███▀▀▀     ▀▀███▀▀▀     ▀▀█████▄    ")
	fmt.Println("    ███    ███ ███    ███          ███   ███    █▄    ███    █▄    ███▐██▄   ")
	fmt.Println("    ███    ███ ███    ███    ▄█    ███   ███    ███   ███    ███   ███ ▀███▄ ")
	fmt.Println("    ████████▀   ▀██████▀   ▄████████▀    ██████████   ██████████   ███   ▀█▀ ")
	fmt.Println("                                                                   ▀         ")
	fmt.Println("                           OSINT & OPSEC Multi-Tool                          ")
	fmt.Println()
}

// Dispban ...
func Dispban(text string) {
	fmt.Println("   [ " + text + " ]")
}

// Dispop ...
func Dispop(key string, value string) string {
	var reset string = "\033[0m"
	var red string = "\033[31m"
	return fmt.Sprintf("   [ " + string(red) + key + string(reset) + " ] - " + value)
}

// Dispopg ...
func Dispopg(key string, value string) string {
	var reset string = "\033[0m"
	var green string = "\033[32m"
	return fmt.Sprintf("   [ " + string(green) + key + string(reset) + " ] - " + value)
}

// Dispopw ...
func Dispopw(key string, value string) string {
	return fmt.Sprintf("   [ " + key + " ] - " + value)

}

// ScanIt ...
func ScanIt(s *bufio.Scanner) string {
	s.Scan()
	return s.Text()
}

// TreeIt ...
func TreeIt(r string, label string, add string, succeeded bool) (string, bool) {
	if len(add) > 0 || add != "" {
		if !succeeded {
			r += label + add
			succeeded = true
		} else {
			r += "\n" + label + add
		}
	}
	return r, succeeded
}
