package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Parse()
	fileName := flag.Arg(0)

	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("error reading file:", err)
	}

	fmt.Printf("Are contents of %s, valid json: %t\n", fileName, isValidJson(data))
}

func isValidJson(data []byte) bool {
	if len(data) == 0 {
		return false
	}

	if !bytes.HasPrefix(data, []byte("{")) {
		return false
	}

	if !bytes.HasSuffix(data, []byte("}")) {
		return false
	}

	return true
}
