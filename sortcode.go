package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"time"
)

func readLines(file string) (linesArray []string, err error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		const delim = '\n'
		line, err := r.ReadString(delim)
		if err == nil || len(line) > 0 {
			if err != nil {
				line += string(delim)
			}
			linesArray = append(linesArray, line)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return linesArray, nil
}

func writeLines(lines []string) (err error) {
	f, err := os.Create("output.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()
	for _, line := range lines {
		_, err := w.WriteString(line)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	now := time.Now()
	linesArray, err := readLines(`../../sort-challenge/sample_large.txt`)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	sort.Strings(linesArray)
	err = writeLines(linesArray)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	now2 := time.Now()
	fmt.Println(now2.Second()-now.Second(), "Secs")
}
