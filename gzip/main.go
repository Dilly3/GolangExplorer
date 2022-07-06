package main

import (
	"compress/gzip"
	"io"
	"os"
)

func main() {
	Compress("names.txt")
}

func Compress(filename string) error {
	input, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer input.Close()
	output, err := os.Create(filename + ".gz")
	if err != nil {
		return err
	}
	defer output.Close()
	gzoutput := gzip.NewWriter(output)

	_, err = io.Copy(gzoutput, input)

	if err != nil {
		return err
	}
	gzoutput.Close()

	return nil
}
