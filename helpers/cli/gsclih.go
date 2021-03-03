package cli

import (
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

	fmt.Println("")

}

// Dispban ...
func Dispban(text string) {
	var reset string = "\033[0m"
	var red string = "\033[31m"
	fmt.Println("   [ " + string(red) + text + string(reset) + " ]")
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
