package main

import (
	"os"
)

func main() {
	exit(0)
}

func exit(code int) {
	os.Exit(code)
}
