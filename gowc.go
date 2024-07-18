package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/akamensky/argparse"
	"io"
	"io/fs"
	"os"
	"strconv"
	"strings"
)

func main() {
	parser := argparse.NewParser("wcgo", "Count line, words, bytes and characters in the given input")

	countBytes := parser.Flag("c", "bytes", &argparse.Options{Default: false, Help: "Prints the number of bytes in the file"})
	countLines := parser.Flag("l", "lines", &argparse.Options{Default: false, Help: "Prints the number of lines in the file"})
	countWords := parser.Flag("w", "words", &argparse.Options{Default: false, Help: "Prints the number of words in the file"})
	countChars := parser.Flag("m", "characters", &argparse.Options{Default: false, Help: "Prints the number of characters in the file"})

	inputFile := parser.File("f", "file", 0600, fs.FileMode(os.O_RDONLY), &argparse.Options{Default: nil, Help: "File to process"})

	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	}

	var inputBytes []byte

	// Handle input source (either stdin or file)
	if argparse.IsNilFile(inputFile) {
		inputBytes, _ = io.ReadAll(os.Stdin)
	} else {
		// Check if input is valid path and is file
		stat, inputFileErr := inputFile.Stat()
		if inputFileErr != nil {
			if errors.Is(err, os.ErrNotExist) {
				fmt.Println("File doesn't exist")
				return
			}
		} else if stat.IsDir() {
			fmt.Println("Path is directory")
			inputFile.Close()
			return
		}

		inputBytes, _ = io.ReadAll(inputFile)
	}

	// Path is existing file
	var message strings.Builder

	if !*countBytes && !*countLines && !*countWords && !*countChars {
		*countBytes = true
		*countLines = true
		*countWords = true
	}

	if countLines != nil && *countLines {
		input := bufio.NewScanner(strings.NewReader(string(inputBytes)))
		num := countIn(input, Lines)
		message.WriteString(fmt.Sprintf(strconv.FormatInt(num, 10)))
	}

	if countWords != nil && *countWords {
		input := bufio.NewScanner(strings.NewReader(string(inputBytes)))
		num := countIn(input, Words)
		message.WriteString("\t")
		message.WriteString(fmt.Sprintf(strconv.FormatInt(num, 10)))
	}

	if countBytes != nil && *countBytes {
		input := bufio.NewScanner(strings.NewReader(string(inputBytes)))
		num := countIn(input, Bytes)
		message.WriteString("\t")
		message.WriteString(fmt.Sprintf(strconv.FormatInt(num, 10)))
	}

	if countChars != nil && *countChars {
		input := bufio.NewScanner(strings.NewReader(string(inputBytes)))
		num := countIn(input, Chars)
		message.WriteString("\t")
		message.WriteString(fmt.Sprintf(strconv.FormatInt(num, 10)))
	}

	message.WriteString(" ")
	if !argparse.IsNilFile(inputFile) {
		message.WriteString(inputFile.Name())
		inputFile.Close()
	}

	fmt.Println("\t", message.String())
}

func countIn(input *bufio.Scanner, operation WhatToCount) int64 {
	var splitFunc bufio.SplitFunc = nil

	switch operation {
	case Lines:
		splitFunc = bufio.ScanLines
	case Words:
		splitFunc = bufio.ScanWords
	case Chars:
		splitFunc = bufio.ScanRunes
	case Bytes:
		splitFunc = bufio.ScanBytes
	}

	input.Split(splitFunc)

	var count int64 = 0
	for input.Scan() {
		count++
	}

	return count
}

type WhatToCount int64

const (
	Bytes WhatToCount = 0
	Words WhatToCount = 1
	Lines WhatToCount = 2
	Chars WhatToCount = 3
)
