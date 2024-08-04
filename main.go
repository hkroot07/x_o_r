package main

import (
	"flag"
	"fmt"
	"os"
)

var cipherMode = flag.Bool("cipher", true, "Enable cipher mode. This is the default mode.")
var decipherMode = flag.Bool("decipher", false, "Enable decipher mode.")
var secretKey = flag.String("secret", "", "Your secret key. Must contain at least 1 character")

func main() {
	flag.Parse()

	if *cipherMode && *decipherMode {
		fmt.Println("Please specify only one mode at a time.")
		os.Exit(1)
	}

	if len(*secretKey) == 0 {
		fmt.Println("No secret is provided! Exiting now ...")
		os.Exit(1)
	}
}
