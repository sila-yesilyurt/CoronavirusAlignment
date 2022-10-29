package main

import (
	"Alignment/Functions"
	"fmt"
)

func main() {
	fmt.Println("Sequence alignment!")

	//read in data

	/*
		//hemoglobin


		zebrafish := ReadFASTAFile("Data/Hemoglobin/Danio_rerio_hemoglobin.fasta")
		cow := ReadFASTAFile("Data/Hemoglobin/Bos_taurus_hemoglobin.fasta")
		gorilla := ReadFASTAFile("Data/Hemoglobin/Gorilla_gorilla_hemoglobin.fasta")
		human := ReadFASTAFile("Data/Hemoglobin/Homo_sapiens_hemoglobin.fasta")

		//now perform alignments

		match := 1.0
		mismatch := 1.0
		gap := 3.0


		alignment_0 := Functions.GlobalAlignment(zebrafish, human, match, mismatch, gap)

		alignment_1 := Functions.GlobalAlignment(cow, human, match, mismatch, gap)

		alignment_2 := Functions.GlobalAlignment(gorilla, human, match, mismatch, gap)

		fmt.Println("Human and zebrafish")
		PrintAlignment(alignment_0)
		fmt.Println()

		fmt.Println("Human and cow")
		PrintAlignment(alignment_1)
		fmt.Println()

		fmt.Println("Human and gorilla")
		PrintAlignment(alignment_2)
	*/

	sars := ReadFASTAFile("Data/Coronaviruses/SARS-CoV_genome.fasta")

	sars2 := ReadFASTAFile("Data/Coronaviruses/SARS-CoV-2_genome.fasta")

	//let's global align them

	match := 1.0
	mismatch := 10.0
	gap := 1.0

	fmt.Println("Aligning coronavirus genomes.")

	SARS_alignment := Functions.GlobalAlignment(sars, sars2, match, mismatch, gap)

	fmt.Println("Alignment complete! Writing to file.")

	outfile := "Output/coronavirus_alignment.fasta"
	WriteAlignmentToFASTA(SARS_alignment, outfile)

	fmt.Println("Alignment written to file. Program exiting.")

}
