package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		clearWindowTitle()
		os.Exit(0)
	}

	switch os.Args[1] {
	case "-v", "--version":
		printVersion()
	case "-?", "-h", "--help":
		printUsage()
	default:
		setTerminalTitle(os.Args[1:])
	}
}

func printUsage() {
	fmt.Println("Usage: " + os.Args[0] + " [title]")
	fmt.Println("Set the terminal title.\n")
	fmt.Println("Options:")
	fmt.Println("  -v, --version\tDisplay script version.")
	fmt.Println("  -?, -h, --help\tShow this help message.")
	os.Exit(0)
}

func printVersion() {
	fmt.Println(os.Args[0] + " v1.1")
	fmt.Println("by PhateValleyman")
	fmt.Println("Jonas.Ned@outlook.com")
	os.Exit(0)
}

func setTerminalTitle(args []string) {
	title := strings.Join(args, " ")
	fmt.Print("\033]0;" + title + "\007")
	fmt.Println("Title set to:\n'\033[01;33m" + title + "\033[0m'")
}

func clearWindowTitle() {
	fmt.Print("\033]0;\007")
	fmt.Println("Title cleared")
}
