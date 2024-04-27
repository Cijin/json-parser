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

	os.Stdout.Write(data)

	fmt.Printf("Are contents of %s, valid json: %t\n", fileName, isValidJson(data))
}

func isValidJson(data []byte) bool {
	data = bytes.TrimSpace(data)

	if len(data) == 0 {
		return false
	}

	if !bytes.HasPrefix(data, []byte("{")) {
		return false
	}

	if !bytes.HasSuffix(data, []byte("}")) {
		return false
	}

	l := NewLexer(data)

	for {
		err := l.NextChar()
		if err != nil {
			fmt.Println(err)
			return false
		}

		if l.ch == 0 {
			break
		}
	}
	// read till end of input

	return true
}
