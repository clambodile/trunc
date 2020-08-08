package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var length = flag.Int("length", 80, "Length to truncate lines to.")
	var filename = flag.String("file", "", "Name of file to truncate.")
	flag.Parse()

	var dat []byte
	var err error
	if *filename == "" {
		dat, err = ioutil.ReadAll(bufio.NewReader(os.Stdin))
	} else {
		dat, err = ioutil.ReadFile(*filename)
	}

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")
	for _, line := range lines {
		chars := []rune(line)
		oldLength := len(chars)
		var newLines = make([]string, 0)
		for i := 0; i < oldLength; i += *length {
			var stopIndex int
			if i + *length < oldLength {
				stopIndex = i + *length
			} else {
				stopIndex = oldLength
			}
			newLines = append(newLines, string(chars[i:stopIndex]))
		}
		newLineStr := strings.Join(newLines, "\n")
		fmt.Println(newLineStr)
	}
}
