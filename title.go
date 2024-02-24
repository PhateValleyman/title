package main

import (
	// Importing the fmt package for printing formatted output
	"fmt"
	// Importing the os package for accessing command-line arguments
	"os"
	// Importing the strings package for string manipulation
	"strings"
)

// Define blue color escape sequence
const B = "\033[01;34m"
// Define green color escape sequence
const G = "\033[32m"
// Define red color escape sequence
const R = "\033[31m"
// Define yellow color escape sequence
const Y = "\033[01;33m"
// Define reset color escape sequence
const N = "\033[0m"

func main() {
	// Getting command-line arguments, excluding the program name
	args := os.Args[1:]
	// Checking if arguments are provided
	if len(args) > 0 {
		// Checking if the argument is for version
		if args[0] == "-v" || args[0] == "--version" {
			// Printing program name with version
			fmt.Printf("%s%s%s v1.1\n", Y, os.Args[0], N)
			// Printing author's name
			fmt.Printf("by %sPhateValleyman%s\n", Y, N)
			// Printing author's email
			fmt.Printf("%sJonas.Ned@outlook.com%s\n", G, N)
			// Exiting the program
			os.Exit(0)
		// Checking if the argument is for help
		} else if args[0] == "-?" || args[0] == "-h" || args[0] == "--help" {
			// Printing usage instructions
			fmt.Printf("Usage: %s%s%s '%sTITLE%s'\n", Y, os.Args[0], N, G, N)
			// Printing instructions for setting terminal title
			fmt.Printf("Set terminal title to '%sTITLE%s'\n", G, N)
			// Printing options
			fmt.Println("Options:")
			// Printing help option
			fmt.Println("  -h, --help\t\tShow this help message.")
			// Printing version option
			fmt.Println("  -v, --version\t\tDisplay script version.")
			// Exiting the program
			os.Exit(0)
		}
	}

	// Checking if no arguments are provided
	if len(args) == 0 {
		// Setting terminal title to empty string
		fmt.Print("\033]0;\007")
		// Printing message indicating title unset
		fmt.Printf("%sTitle unset%s\n", Y, N)
	}

	// Checking if one argument is provided
	if len(args) == 1 {
		// Setting terminal title to the provided argument
		fmt.Printf("\033]0;%s\007", args[0])
		// Printing message indicating title set
		fmt.Printf("Title set to:\n'%s%s%s'\n", G, args[0], N)
	}

	// Checking if multiple arguments are provided
	if len(args) > 1 {
		// Joining multiple arguments into a single string
		title := strings.Join(args, " ")
		// Setting terminal title to the joined string
		fmt.Printf("\033]0;%s\007", title)
		// Printing message indicating title set
		fmt.Printf("Title set to:\n'%s%s%s'\n", G, title, N)
	}
}
