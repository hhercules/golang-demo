package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) <= 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s FILE\n", os.Args[0])
		os.Exit(-1)
	}

	for _, file := range args {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			continue
		}
		st, err := f.Stat()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			continue
		}
		if st.IsDir() {
			fmt.Fprintf(os.Stderr, "%s is a directory\n", file)
			continue
		}
		md5 := md5.New()
		io.Copy(md5, f)
		fmt.Printf("%x\t%s\n", md5.Sum(nil), file)
		f.Close()
	}
}
