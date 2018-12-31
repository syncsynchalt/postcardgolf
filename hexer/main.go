package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dieUsage(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error(), "\n")
	}
	fmt.Fprintf(os.Stderr, "Usage: %s [filename]\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Reads a text file of hex codes and outputs the binary version.")
	fmt.Fprintln(os.Stderr, "")
	os.Exit(1)
}

func die(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", a...)
	os.Exit(1)
}

func fileExists(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	}
	return false
}

func main() {
	if len(os.Args) != 2 {
		dieUsage(nil)
	}

	infile, err := os.Open(os.Args[1])
	if err != nil {
		dieUsage(err)
	}
	defer infile.Close()

	fscanner := bufio.NewScanner(infile)
	for fscanner.Scan() {
		line := fscanner.Text()
		if comment := strings.Index(line, ";"); comment != -1 {
			line = line[0:comment]
		}

		words := strings.Fields(line)
		output := make([]byte, 0, 100)
		for _, word := range words {
			var num int64
			if len(word) > 1 && word[0] == '\'' {
				num, err = strconv.ParseInt(word[1:], 10, 8)
			} else {
				num, err = strconv.ParseInt(word, 16, 8)
			}
			if err != nil || num < 0 || num > 255 {
				die("Bad number %s", word)
			}
			output = append(output, byte(num))
		}
		fmt.Print(string(output))
	}
}
