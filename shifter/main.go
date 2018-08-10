package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func dieUsage(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error(), "\n")
	}
	fmt.Fprintf(os.Stderr, "Usage: %s [shiftnum]\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "Reads stdin and shifts each byte up by shiftnum, writing the result to stdout.")
	fmt.Fprintln(os.Stderr, "")
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		dieUsage(nil)
	}

	shiftnum, err := strconv.ParseInt(os.Args[1], 0, 0)
	if err != nil {
		dieUsage(err)
	}

	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	for {
		c, err := in.ReadByte()
		if err == io.EOF {
			break
		}
		c += byte(shiftnum)
		out.WriteByte(c)
	}
    out.Flush()
}
