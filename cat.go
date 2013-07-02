package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) <= 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s FILES\n", os.Args[0])
		os.Exit(-1)
	}

	for _, file := range args {
		f, e := os.Open(file)
		if e != nil {
			fmt.Fprintf(os.Stderr, "%s doesn't exist\n", file)
			continue
		}
		st, e := f.Stat()
		if e != nil {
			fmt.Fprintf(os.Stderr, "%s\n", e.Error())
			continue
		}
		if st.IsDir() {
			fmt.Fprintf(os.Stderr, "%s is a directory\n", file)
			continue
		}
		io.Copy(os.Stdout, f)
		f.Close()
	}
}
