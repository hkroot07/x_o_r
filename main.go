package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var mode = flag.String("mode", "cipher", "Set to 'cipher' or 'decipher'. Default is 'cipher'.")
var secretKey = flag.String("secret", "", "Your secret key. Must contain at least 1 character")

func main() {
	flag.Parse()

	switch *mode {
	case "cipher":
		plaintext := getUserInput("Enter your text to cipher: ")
		fmt.Println(plaintext)
	case "decipher":
		cipheredText := getUserInput("Enter your ciphered data to decipher: ")
	default:
		fmt.Println("Invalid mode. Use 'cipher' or 'decipher'.")
		os.Exit(1)
	}

	if len(*secretKey) == 0 {
		fmt.Println("No secret is provided! Exiting now ...")
		os.Exit(1)
	}
}

func getUserInput(msg string) string {
	fmt.Print(msg)

	reader := bufio.NewReader(os.Stdin)

	for {
		result, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading the entered text! Please try again")
			continue
		}

		return strings.TrimRight(result, "\n")
	}
}
