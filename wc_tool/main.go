// Challenge: https://codingchallenges.fyi/challenges/challenge-wc/
// References:
// https://blog.ippon.tech/no-framework-approach-to-building-a-cli-with-go/
// https://pkg.go.dev/flag
// https://www.scaler.com/topics/golang/golang-read-file/

package main

import (
	"flag"
	"fmt"
	"os"
)

var bytesCount string
var lineCount string
var wordCount string

var usage = `Usage:
NAME
	./wc_tool

DESCRIPTION
	-c <filepath>: outputs the number of bytes in a file
`

// fileBytes returns the count of bytes in a file.
func fileBytes() (int, error) {
	content, err := os.ReadFile(bytesCount) // returns a byte array of content
	if err != nil {
		return 0, err
	}

	return len(content), nil
}

// fileLines returns the count of lines in a file.
func fileLines() (int, error) {
	content, err := os.ReadFile(lineCount)

	if err != nil {
		return 0, err
	}

	var lc int
	for _, c := range content {
		if c == byte('\n') {
			lc++
		}
	}

	return lc, nil
}

// fileWords returns the count of words in a file.
func fileWords() (int, error) {
	content, err := os.ReadFile(wordCount)

	if err != nil {
		return 0, err
	}

	var wc int
	bc := len(content)
	for i := 0; i < bc-1; i++ {
		c := content[i]
		c1 := content[i+1]
		if (c == byte(' ') || c == byte('\t') || c == byte('\r') || c == byte('\n')) &&
			(c1 != byte(' ') && c1 != byte('\t') && c1 != byte('\r') && c1 != byte('\n')) {
			wc++
		}
	}
	if content[bc-1] != byte(' ') || content[bc-1] == byte('\t') || content[bc-1] == byte('\n') {
		wc++
	}
	return wc, nil
}

func main() {
	fmt.Println(usage)

	// TODO: The file is assumed to be in same directory.
	flag.StringVar(&bytesCount, "c", "", "File name to be parsed for bytes count")
	flag.StringVar(&lineCount, "l", "", "File name to be parsed for line count")
	flag.StringVar(&wordCount, "w", "", "File name to be parsed for word count")
	flag.Parse()

	// If number of bytes are requested.
	if bytesCount != "" {
		// bytes in the file
		contentLen, err := fileBytes()
		if err != nil {
			fmt.Printf("Error while processing file %s - %s\n", bytesCount, err)
			return
		}
		fmt.Printf("%d %s\n", contentLen, bytesCount)
	}

	// If number of lines are requested.
	if lineCount != "" {
		// lines in the file
		contentLen, err := fileLines()
		if err != nil {
			fmt.Printf("Error while processing file %s - %s\n", lineCount, err)
			return
		}
		fmt.Printf("%d %s\n", contentLen, lineCount)
	}

	// If number of words are requested.
	if wordCount != "" {
		// words in the file
		contentLen, err := fileWords()
		if err != nil {
			fmt.Printf("Error while processing file %s - %s\n", wordCount, err)
			return
		}
		fmt.Printf("%d %s\n", contentLen, wordCount)
	}
}
