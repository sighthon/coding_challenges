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

var fileName string
var usage = `Usage:
NAME
	ccwc

DESCRIPTION
	-c <filepath>: outputs the number of bytes in a file
`

func fileBytes() (int, error) {
	content, err := os.ReadFile(fileName) // returns a byte array of content
	if err != nil {
		return 0, err
	}

	return len(content), nil
}

func main() {
	// TODO: The file is assumed to be in same directory.
	flag.StringVar(&fileName, "c", "", "File name to be parsed")
	flag.Parse()

	fmt.Println(usage)

	// bytes in the file
	if fileName != "" {
		contentLen, err := fileBytes()
		if err != nil {
			fmt.Printf("Error while processing file %s", fileName)
			return
		}
		// TODO: Extra "%" symbol in output ex:"335039 fileName %"
		fmt.Printf("%d %s", contentLen, fileName)
	}

}
