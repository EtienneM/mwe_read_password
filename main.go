package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func main() {
	_, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println(err)
}
