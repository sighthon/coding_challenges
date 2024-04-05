// Challenge: https://codingchallenges.fyi/challenges/challenge-wc/
// References:
// https://blog.ippon.tech/no-framework-approach-to-building-a-cli-with-go/
// https://pkg.go.dev/flag
// https://www.scaler.com/topics/golang/golang-read-file/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

var bytesCount string
var lineCount string
var wordCount string
var characterCount string

var usage = `Usage:
NAME
	./wc_tool

DESCRIPTION
	-c <filepath>: outputs the number of bytes in a file
	-l <filepath>: outputs the number of lines in a file
	-w <filepath>: outputs the number of words in a file
	-m <filepath>: outputs the number of characters in a file
`

// fileBytes returns the count of bytes in a file.
func fileBytes(content []byte) int {
	return len(content)
}

// fileLines returns the count of lines in a file.
func fileLines(content []byte) int {
	var lc int
	for _, c := range content {
		if c == byte('\n') {
			lc++
		}
	}

	return lc
}

// fileWords returns the count of words in a file.
func fileWords(content []byte) int {
	var wc int
	bc := len(content)
	for i := 0; i < bc-1; i++ {
		c := content[i]
		c1 := content[i+1]
		// The next byte has to be not a character
		if (c != byte(' ') && c != byte('\t') && c != byte('\r') && c != byte('\n')) &&
			(c1 == byte(' ') || c1 == byte('\t') || c1 == byte('\r') || c1 == byte('\n')) {
			wc++
		}
	}

	// handling for last character. The last word maynot be followed by a space or new line and wouldn't be
	// counted in previous loop.
	c := content[bc-1]
	if c != byte(' ') && c != byte('\t') && c != byte('\r') && c != byte('\n') {
		wc++
	}

	return wc
}

// fileCharacters returns the count of characters in a file.
func fileCharacters(content []byte) int {
	// utf8.RuneCountInString(characterCount) is used to count runes (characters) in a string.
	runeSlice := make([]rune, 0)
	for len(content) > 0 {
		r, size := utf8.DecodeRune(content)
		runeSlice = append(runeSlice, r)
		content = content[size:]
	}

	return len(runeSlice)
}

func main() {
	// TODO: The file is assumed to be in same directory.
	flag.StringVar(&bytesCount, "c", "", "File name to be parsed for bytes count")
	flag.StringVar(&lineCount, "l", "", "File name to be parsed for line count")
	flag.StringVar(&wordCount, "w", "", "File name to be parsed for word count")
	flag.StringVar(&characterCount, "m", "", "File name to be parsed for character count")

	// Set flag error handling mode to ContinueOnError
	// TODO(sighthon): Not print the error on invalid input
	flag.CommandLine.Init(os.Args[0], flag.ContinueOnError)
	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		// If input is available from the pipe, read it
		stat, _ := os.Stdin.Stat()
		var content string
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			scanner := bufio.NewScanner(os.Stdin)
			// Scan each line and concatenate them into a single string
			for scanner.Scan() {
				line := scanner.Text()
				content += line + "\n" // Add newline to preserve line breaks
			}
		}

		// If flag was provided but flag value wasn't, use the content from string
		// TODO(sighthon): Fix the check. The nil check uses the default value.
		var contentLen int
		if strings.Contains(err.Error(), "-c") {
			contentLen = fileBytes([]byte(content))
		} else if strings.Contains(err.Error(), "-l") {
			contentLen = fileLines([]byte(content))
		} else if strings.Contains(err.Error(), "-w") {
			contentLen = fileWords([]byte(content))
		} else if strings.Contains(err.Error(), "-m") {
			contentLen = fileCharacters([]byte(content))
		}

		fmt.Printf("%d \n", contentLen)
		return
	}

	///////////
	// File parsing beyond this point
	//////////

	// If number of bytes are requested.
	if bytesCount != "" {
		// bytes in the file
		content, err := os.ReadFile(bytesCount) // returns a byte array of content
		if err != nil {
			fmt.Printf("Error while processing file %s - %s\n", bytesCount, err)
			return
		}
		contentLen := fileBytes(content)
		fmt.Printf("%d %s\n", contentLen, bytesCount)
		return
	}

	// If number of lines are requested.
	if lineCount != "" {
		// lines in the file
		content, err := os.ReadFile(lineCount) // returns a byte array of content
		if err != nil {
			fmt.Printf("Error while processing file %s - %s\n", lineCount, err)
			return
		}
		contentLen := fileLines(content)
		fmt.Printf("%d %s\n", contentLen, lineCount)
		return
	}

	// If number of words are requested.
	if wordCount != "" {
		// words in the file
		content, err := os.ReadFile(wordCount) // returns a byte array of content
		if err != nil {
			fmt.Printf("Error while processing file %s - %s\n", wordCount, err)
			return
		}
		contentLen := fileWords(content)
		fmt.Printf("%d %s\n", contentLen, wordCount)
		return
	}

	// If number of characters are requested.
	if characterCount != "" {
		// characters in the file
		content, err := os.ReadFile(characterCount) // returns a byte array of content
		if err != nil {
			fmt.Printf("Error while processing file %s - %s\n", characterCount, err)
			return
		}
		contentLen := fileCharacters(content)
		fmt.Printf("%d %s\n", contentLen, characterCount)
		return
	}

	// If no flag value was provided, use the second argument as filename
	args := os.Args[1:]
	if len(args) > 0 {
		lastArg := args[len(args)-1]
		content, err := os.ReadFile(lastArg) // returns a byte array of content
		if err != nil {
			fmt.Printf("Error while processing file %s - %s\n", lastArg, err)
			return
		}

		a1 := fileBytes(content)
		a2 := fileLines(content)
		a3 := fileWords(content)

		fmt.Printf("%d %d %d %s\n", a2, a3, a1, lastArg)
	} else {
		fmt.Println(usage)
	}
}
