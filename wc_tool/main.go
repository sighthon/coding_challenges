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
	"unicode/utf8"
)

var help = flag.Bool("help", false, "Show help")
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
func fileBytes() (int, error) {
	content, err := os.ReadFile(bytesCount) // returns a byte array of content
	if err != nil {
		return 0, err
	}

	return len(content), nil
}

// fileLines returns the count of lines in a file.
func fileLines() (int, error) {
	content, err := os.ReadFile(lineCount) // returns a byte array of content

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
	content, err := os.ReadFile(wordCount) // returns a byte array of content

	if err != nil {
		return 0, err
	}

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

	return wc, nil
}

// fileCharacters returns the count of characters in a file.
func fileCharacters() (int, error) {
	// utf8.RuneCountInString(characterCount) is used to count runes (characters) in a string.
	content, err := os.ReadFile(characterCount) // returns a byte array of content

	if err != nil {
		return 0, err
	}

	runeSlice := make([]rune, 0)
	for len(content) > 0 {
		r, size := utf8.DecodeRune(content)
		runeSlice = append(runeSlice, r)
		content = content[size:]
	}

	return len(runeSlice), nil
}

// isFlagPassed determines if a particular flag was passed to the program.
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func main() {
	// TODO: The file is assumed to be in same directory.
	flag.StringVar(&bytesCount, "c", "", "File name to be parsed for bytes count")
	flag.StringVar(&lineCount, "l", "ddd", "File name to be parsed for line count")
	flag.StringVar(&wordCount, "w", "", "File name to be parsed for word count")
	flag.StringVar(&characterCount, "m", "", "File name to be parsed for character count")
	flag.Parse()

	// if !isFlagPassed("l") {
	// 	fmt.Println("Yess")
	// }

	// // Usage demo
	// if *help {
	// 	flag.Usage()
	// 	os.Exit(0)
	// }

	// If number of bytes are requested.
	if bytesCount != "" {
		// bytes in the file
		contentLen, err := fileBytes()
		if err != nil {
			fmt.Printf("Error while processing file %s - %s\n", bytesCount, err)
			return
		}
		fmt.Printf("%d %s\n", contentLen, bytesCount)
		return
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
		return
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
		return
	}

	// If number of characters are requested.
	if characterCount != "" {
		// characters in the file
		contentLen, err := fileCharacters()
		if err != nil {
			fmt.Printf("Error while processing file %s - %s\n", characterCount, err)
			return
		}
		fmt.Printf("%d %s\n", contentLen, characterCount)
		return
	}

	// If no flag was provided, use the second argument as filename
	args := os.Args[1:]
	if len(args) > 0 {
		lastArg := args[len(args)-1]
		bytesCount = lastArg
		lineCount = lastArg
		wordCount = lastArg

		a1, err := fileBytes()
		if err != nil {
			fmt.Printf("Error while processing file %s - %s\n", bytesCount, err)
			return
		}
		a2, err := fileLines()
		if err != nil {
			fmt.Printf("Error while processing file %s - %s\n", lineCount, err)
			return
		}
		a3, err := fileWords()
		if err != nil {
			fmt.Printf("Error while processing file %s - %s\n", wordCount, err)
			return
		}
		fmt.Printf("%d %d %d %s\n", a2, a3, a1, lastArg)
	} else {
		fmt.Println(usage)
	}
}
