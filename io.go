package main

import (
	"Alignment/Functions"
	"bufio"
	"fmt"
	"os"
)

func PrintAlignment(a Functions.Alignment) {
	fmt.Println(a[0])
	fmt.Println(a[1])
}

//ReadFASTAFile takes a file name with a single FASTA header and reads out all elements
//as a string other than header elements.
func ReadFASTAFile(filename string) string {
	file, err := os.Open(filename)

	if err != nil {
		// error in opening file
		panic("Error: something went wrong with file open (probably you gave wrong filename).")
	}

	scanner := bufio.NewScanner(file)
	genome := ""
	for scanner.Scan() {
		currentLine := scanner.Text() // grabs one line of text and returns a string
		if len(currentLine) > 0 && currentLine[0] != '>' {
			genome += currentLine
		}
	}
	return genome
}

//WriteAlignmentToFASTA takes an alignment and a file name and writes the alignment
//to the file as a FASTA. It uses "string_1" and "string_2" as the headers.
func WriteAlignmentToFASTA(a Functions.Alignment, filename string) {
	file, err := os.Create(filename)
	if err != nil { // panic if anything went wrong
		panic(err)
	}

	writer := bufio.NewWriter(file)
	// first header
	fmt.Fprintln(writer, ">string_1")

	//first string
	fmt.Fprintln(writer, a[0])

	//second header
	fmt.Fprintln(writer, ">string_2")

	//second string
	fmt.Fprintln(writer, a[1])

	writer.Flush()

	file.Close()
}
