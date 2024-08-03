package main

import "flag"

var cipherMode = flag.Bool("cipher", true, "Enable cipher mode. This is the default mode.")
var decipherMode = flag.Bool("decipher", false, "Enable decipher mode.")
var secretKey = flag.String("secret", "", "Your secret key. Must contain at least 1 character")

func main() {

}
