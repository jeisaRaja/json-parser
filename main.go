package main

import (
	"flag"
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	filename := flag.String("f", "", "file location")

	flag.Parse()

	if *filename == "" {
		panic("File not found")
	}

	json, err := os.ReadFile(*filename)
	checkError(err)

  fmt.Println(string(json))
}
