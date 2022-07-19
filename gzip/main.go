package main

import (
	"compress/gzip"
	"io"
	"os"
	"sync"
)

var Wg = &sync.WaitGroup{}

func main() {

	for _, filename := range os.Args[1:] {
		Wg.Add(1)
		go func(filename string) {
			Compress(filename)
			Wg.Done()
		}(filename)
	}
	Wg.Wait()
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
	defer gzoutput.Close()

	_, err = io.Copy(gzoutput, input)

	if err != nil {
		return err
	}

	return nil
}
