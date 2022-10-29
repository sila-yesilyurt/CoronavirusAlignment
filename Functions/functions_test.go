package Functions

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

type DistanceMatrix [][]float64

/********************************************
 Levenshtein Matrix Tests (Edit-Distance Tests)
*********************************************/

type levenshteinMatrixTestpair struct {
	patterns   []string
	matrixPath string
}

// levenshteinMatrixTests is a list of paired string arrays and file paths
// where the associated files have correct levenshtein matricies for those
// strings. If you want to add more tests, just add another pair and put the
// file in the Tests folder to keep everything organized.
var levenshteinMatrixTests = []levenshteinMatrixTestpair{
	{[]string{"A", "T"},
		"Tests/matrix1.txt"},
	{[]string{"-GCTAG--TACCTATCGGA---", "TAGCTAG---TCGAT", "AGC---TAGGGATCGAAAT----"},
		"Tests/matrix2.txt"},
	{[]string{"ATCGCTAGCTCTTCGATC", "TTTCGATCTTAAGATC"},
		"Tests/matrix3.txt"},
	{[]string{"ATCGCTAGCTCTTCGATC", "TTTCGATCTTAAGATC", "--TTTCGATC---TTAAGATC--", "ATC--GCTAG-CTCTTCG---ATC"},
		"Tests/matrix4.txt"},
	{[]string{"A", "A", "A"},
		"Tests/matrix5.txt"}}

func TestEditDistanceMatrix(t *testing.T) {
	for _, pair := range levenshteinMatrixTests {
		mtx64, _ := ReadMatrixFromFile(pair.matrixPath)
		mtx := TwoDMatrixFloat64ToInt(mtx64)
		v := EditDistanceMatrix(pair.patterns)
		if !reflect.DeepEqual(v, mtx) {
			t.Error(
				"For", pair.patterns,
				"expected", mtx,
				"got", v,
			)
		}
	}
}

//ReadMatrixFromFile takes a file name and reads the information in this file to produce
//a distance matrix and a slice of strings holding the species names.  The first line of the
//file should contain the number of species.  Each other line contains a species name
//and its distance to each other species.
func ReadMatrixFromFile(fileName string) (DistanceMatrix, []string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error: couldn't open the file")
		os.Exit(1)
	}
	var lines []string = make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		fmt.Println("Sorry: there was some kind of error during the file reading")
		os.Exit(1)
	}
	file.Close()

	mtx := make(DistanceMatrix, 0)
	speciesNames := make([]string, 0)

	for idx, _ := range lines {
		if idx >= 1 {
			row := make([]float64, 0)
			nums := strings.Split(lines[idx], "\t")
			for i, num := range nums {
				if i == 0 {
					speciesNames = append(speciesNames, num)
				} else {
					n, err := strconv.ParseFloat(num, 64)
					if err != nil {
						fmt.Println("Error: Wrong format of matrix!")
						os.Exit(1)
					}
					row = append(row, n)
				}
			}
			mtx = append(mtx, row)
		}
	}
	return mtx, speciesNames
}

// The TwoDMatrixFloat64ToIn function converts a two dimentional matrix from
// float64 values to int values. This is nesissary for TestEditDistanceMatrix
// due to some weirdness with Io.ReadMatrixFromFile only producing float64
// because of some of the math it uses for its calculations.
func TwoDMatrixFloat64ToInt(m [][]float64) [][]int {
	duplicate := make([][]int, len(m))
	for i := range m {
		duplicate[i] = make([]int, len(m[i]))
		for j := range m[i] {
			duplicate[i][j] = int(m[i][j])
		}
	}
	return duplicate
}

/********************************************
 Change Tests
*********************************************/

type changetestpair struct {
	money       int
	coins       []int
	minNumCoins int
}

var changeTests = []changetestpair{
	{0, []int{0}, 0},
	{0, []int{1, 2, 7}, 0},
	{10, []int{1, 2, 7}, 3},
	{21, []int{1, 2, 7}, 3},
	{15, []int{5, 8, 2}, 3},
	{8, []int{2}, 4},
	{10, []int{3, 1}, 4}}

func TestChange(t *testing.T) {
	for _, pair := range changeTests {
		v := Change(pair.money, pair.coins)
		if v != pair.minNumCoins {
			t.Error(
				"For", pair.money,
				"and", pair.coins,
				"expected", strconv.Itoa(pair.minNumCoins),
				"got", strconv.Itoa(v),
			)
		}
	}
}

/********************************************
 Global Alignment Score Tests
*********************************************/

type GlobalAlignmentInput struct {
	str1     string
	str2     string
	match    float64
	mismatch float64
	gap      float64
}

type globaltestpair struct {
	input    GlobalAlignmentInput
	solScore float64
}

type globalscoretestpair struct {
	input       GlobalAlignmentInput
	scoreMatrix [][]float64
}

var globalScoreTests = []globalscoretestpair{
	{GlobalAlignmentInput{"A",
		"A",
		1.0, 0.5, 1.0},
		[][]float64{[]float64{0, -1}, []float64{-1, 1}}},

	{GlobalAlignmentInput{"GAAC",
		"CAAG",
		1.0, 1.0, 0.5},
		[][]float64{[]float64{0, -0.5, -1, -1.5, -2},
			[]float64{-0.5, -1, -1.5, -2, -0.5},
			[]float64{-1, -1.5, 0, -0.5, -1},
			[]float64{-1.5, -2, -0.5, 1, 0.5},
			[]float64{-2, -0.5, -1, 0.5, 0}}},

	{GlobalAlignmentInput{"GAAC",
		"CAAG",
		1.0, 0.5, 1.0},
		[][]float64{[]float64{0, -1, -2, -3, -4},
			[]float64{-1, -0.5, -1.5, -2.5, -2},
			[]float64{-2, -1.5, 0.5, -0.5, -1.5},
			[]float64{-3, -2.5, -0.5, 1.5, 0.5},
			[]float64{-4, -2, -1.5, 0.5, 1}}},

	{GlobalAlignmentInput{"TACG",
		"CACG",
		1.0, 1.0, 0.5},
		[][]float64{[]float64{0, -0.5, -1, -1.5, -2},
			[]float64{-0.5, -1, -1.5, -2, -2.5},
			[]float64{-1, -1.5, 0, -0.5, -1},
			[]float64{-1.5, 0, -0.5, 1, 0.5},
			[]float64{-2, -0.5, -1, 0.5, 2}}},

	{GlobalAlignmentInput{"TACGG",
		"CACG",
		1.0, 0.5, 1.0},
		[][]float64{[]float64{0, -1, -2, -3, -4},
			[]float64{-1, -0.5, -1.5, -2.5, -3.5},
			[]float64{-2, -1.5, 0.5, -0.5, -1.5},
			[]float64{-3, -1, -0.5, 1.5, 0.5},
			[]float64{-4, -2, -1.5, 0.5, 2.5},
			[]float64{-5, -3, -2.5, -0.5, 1.5}}},

	{GlobalAlignmentInput{"ATCGATCGT",
		"ATCGGCTAC",
		1.0, 1.0, 0.5},
		[][]float64{[]float64{0, -0.5, -1, -1.5, -2, -2.5, -3, -3.5, -4, -4.5},
			[]float64{-0.5, 1, 0.5, 0, -0.5, -1, -1.5, -2, -2.5, -3},
			[]float64{-1, 0.5, 2, 1.5, 1, 0.5, 0, -0.5, -1, -1.5},
			[]float64{-1.5, 0, 1.5, 3, 2.5, 2, 1.5, 1, 0.5, 0},
			[]float64{-2, -0.5, 1, 2.5, 4, 3.5, 3, 2.5, 2, 1.5},
			[]float64{-2.5, -1, 0.5, 2, 3.5, 3, 2.5, 2, 3.5, 3},
			[]float64{-3, -1.5, 0, 1.5, 3, 2.5, 2, 3.5, 3, 2.5},
			[]float64{-3.5, -2, -0.5, 1, 2.5, 2, 3.5, 3, 2.5, 4},
			[]float64{-4, -2.5, -1, 0.5, 2, 3.5, 3, 2.5, 2, 3.5},
			[]float64{-4.5, -3, -1.5, 0, 1.5, 3, 2.5, 4, 3.5, 3}}},

	{GlobalAlignmentInput{"ATCGATCGT",
		"AAC",
		1.0, 1.0, 0.5},
		[][]float64{[]float64{0, -0.5, -1, -1.5},
			[]float64{-0.5, 1, 0.5, 0},
			[]float64{-1, 0.5, 0, -0.5},
			[]float64{-1.5, 0, -0.5, 1},
			[]float64{-2, -0.5, -1, 0.5},
			[]float64{-2.5, -1, 0.5, 0},
			[]float64{-3, -1.5, 0, -0.5},
			[]float64{-3.5, -2, -0.5, 1},
			[]float64{-4, -2.5, -1, 0.5},
			[]float64{-4.5, -3, -1.5, 0}}}}

func computeScore(alignment [2]string, match float64, mismatch float64, gap float64) float64 {
	score := 0.0
	str1 := alignment[0]
	str2 := alignment[1]
	gapC := "-"

	for i, c := range str1 {
		if string(c) == string(str2[i]) {
			score += match
		} else if string(c) == gapC || string(str2[i]) == gapC {
			score += gap
		} else {
			score += mismatch
		}
	}
	return score
}

func TestGlobalScoreTable(t *testing.T) {
	for _, pair := range globalScoreTests {
		v := GlobalScoreTable(pair.input.str1, pair.input.str2, pair.input.match, pair.input.mismatch, pair.input.gap)
		if !reflect.DeepEqual(v, pair.scoreMatrix) {
			t.Error(
				"For", pair.input,
				"expected scoring matrix", pair.scoreMatrix,
				"got", v,
			)
		}
	}
}

/********************************************
 Global Alignment Tests
*********************************************/

var globalTests = []globaltestpair{
	{GlobalAlignmentInput{"A",
		"A",
		1.0, 0.5, 1.0},
		1.000},

	{GlobalAlignmentInput{"GAAC",
		"CAAG",
		1.0, 1.0, 0.5},
		4.000},

	{GlobalAlignmentInput{"GAAC",
		"CAAG",
		1.0, 0.5, 1.0},
		3.000},

	{GlobalAlignmentInput{"TACG",
		"CACG",
		1.0, 1.0, 0.5},
		4.000},

	{GlobalAlignmentInput{"TACGG",
		"CACG",
		1.0, 0.5, 1.0},
		4.500},

	{GlobalAlignmentInput{"ATCGATCGT",
		"ATCGGCTAC",
		1.0, 1.0, 0.5},
		9.000},

	{GlobalAlignmentInput{"ATCGATCGT",
		"AAC",
		1.0, 1.0, 0.5},
		6.000}}

func TestGlobalAlignment(t *testing.T) {
	for _, pair := range globalTests {
		v := GlobalAlignment(pair.input.str1, pair.input.str2, pair.input.match, pair.input.mismatch, pair.input.gap)
		score := computeScore(v, pair.input.match, pair.input.mismatch, pair.input.gap)
		if score != pair.solScore {
			t.Error(
				"For", pair.input,
				"expected alignment with score", strconv.FormatFloat(pair.solScore, 'f', 3, 64),
				"got", v, "with a score of", strconv.FormatFloat(score, 'f', 3, 64),
			)
		}
	}
}

/********************************************
 Hamming Distance Tests
*********************************************/

type hammingTestPair struct {
	str1  string
	str2  string
	hDist int
}

/********************************************
 LCS length Tests
*********************************************/

type lcsLengthTestpair struct {
	str1      string
	str2      string
	lcsLength int
}

var lcsLengthTests = []lcsLengthTestpair{
	{"ATGCGGCTAGCTTAGCCTAGATCGATCGGCTAGCTAGCTAGCCGAGGCTCTCGATCGATCGCGCTAGG",
		"ATGCGGCTGGCTTAGCCTAGTTCGATCGCGTTCGTAGCTATAGAGCTAGCTAGATCGATCGCGCTAGG", 59},
	{"TTGCGGAGCTAGGGATCCGATCGAATATCGATATTCGATCGGGAACACAGATCGAT",
		"GGTACATCGATTCTAGATTCTATAGCGCGCTTCGATCGATTCGATCGATCGAAAAG", 37},
	{"TTGCGGAGCTAGGGATCCGATCGAATATCGATATTCGATCGGGAACACAGATCGAT",
		"TTGCGGAGCTAGGGATCCGATCGAATATCGATATTCGATCGGGAACACAGATCGAT", 56},
	{"TTGCGGAGCTAGGGATCCGATCGAATATCGATATTCGATCGGGAACACAGATCGAT",
		"TTGCGGAGCTAGGGATCCGATCGAATATCGATATAGCATCGGGAACACAGATCGAT", 54},
	{"AGCT",
		"TCGA", 1},
	{"AATT",
		"CCGG", 0},
	{"AGCT",
		"TGAC", 2},
	{"---ATCGATC--ATCGATT-GGACAT",
		"--ATGATGCATC---ATCGGTAG---GCTTAGCTTTTAG", 20}}

func TestLCSLength(t *testing.T) {
	for _, pair := range lcsLengthTests {
		v := LCSLength(pair.str1, pair.str2)
		if v != pair.lcsLength {
			t.Error(
				"For", pair.str1,
				"and", pair.str2,
				"expected", strconv.Itoa(pair.lcsLength),
				"got", strconv.Itoa(v),
			)
		}
	}
}

/********************************************
 Levenshtein Distance Tests
*********************************************/

type levenshteinDistanceTestpair struct {
	str1    string
	str2    string
	levDist int
}

var levenshteinDistanceTests = []levenshteinDistanceTestpair{
	{"A--T",
		"A--T", 0},
	{"ATGCGGCTAGCTTAGCCTAGATCGATCGGCTAGCTAGCTAGCCGAGGCTCTCGATCGATCGCGCTAGG",
		"ATGCGGCTGGCTTAGCCTAGTTCGATCGCGTTCGTAGCTATAGAGCTAGCTAGATCGATCGCGCTAGG", 13},
	{"TTGCGGAGCTAGGGATCCGATCGAATATCGATATTCGATCGGGAACACAGATCGAT",
		"GGTACATCGATTCTAGATTCTATAGCGCGCTTCGATCGATTCGATCGATCGAAAAG", 29},
	{"TTGCGGAGCTAGGGA---TCCGATCGAATATCGATAT---TCGATCGGGAACAC---AGATCGAT",
		"TTGCG-GAGCTAGGGATCCGATCGAATATCGATATTCGATCGGGAA---CACAGATCGAT", 13},
	{"TTGCGGAGCTAGGGATCCGATCGAATATCGATATTCGATCGGGAACACAGATCGAT",
		"TTGCGGAGCTAGGGATCCGATCGAATATCGATATAGCATCGGGAACACAGATCGAT", 3},
	{"TTGCGGAG-CTAGGGATCCGATCGAATA---TCGATATTCGATCGGGAACACAGATCGAT",
		"TT--GCGGAGCTAGGGATCCGATCGAATATCGATATAGCAT----CGGGAAC--ACAGATCGAT", 15}}

func TestEditDistance(t *testing.T) {
	for _, pair := range levenshteinDistanceTests {
		v := EditDistance(pair.str1, pair.str2)
		if v != pair.levDist {
			t.Error(
				"For", pair.str1,
				"and", pair.str2,
				"expected", strconv.Itoa(pair.levDist),
				"got", strconv.Itoa(v),
			)
		}
	}
}

/********************************************
 Local Alignment Score Tests
*********************************************/

type LocalAlignmentInput struct {
	str1     string
	str2     string
	match    float64
	mismatch float64
	gap      float64
}

type localscoretestpair struct {
	input       LocalAlignmentInput
	scoreMatrix [][]float64
}

var localScoreTests = []localscoretestpair{
	{LocalAlignmentInput{"A",
		"A",
		1.0, 0.5, 1.0},
		[][]float64{[]float64{0, 0}, []float64{0, 1}}},

	{LocalAlignmentInput{"GAAC",
		"CAAG",
		1.0, 1.0, 0.5},
		[][]float64{[]float64{0, 0, 0, 0, 0},
			[]float64{0, 0, 0, 0, 1},
			[]float64{0, 0, 1, 1, 0.5},
			[]float64{0, 0, 1, 2, 1.5},
			[]float64{0, 1, 0.5, 1.5, 1}}},

	{LocalAlignmentInput{"GAAC",
		"CAAG",
		1.0, 0.5, 1.0},
		[][]float64{[]float64{0, 0, 0, 0, 0},
			[]float64{0, 0, 0, 0, 1},
			[]float64{0, 0, 1, 1, 0},
			[]float64{0, 0, 1, 2, 1},
			[]float64{0, 1, 0, 1, 1.5}}},

	{LocalAlignmentInput{"TACG",
		"CACG",
		1.0, 1.0, 0.5},
		[][]float64{[]float64{0, 0, 0, 0, 0},
			[]float64{0, 0, 0, 0, 0},
			[]float64{0, 0, 1, 0.5, 0},
			[]float64{0, 1, 0.5, 2, 1.5},
			[]float64{0, 0.5, 0, 1.5, 3}}},

	{LocalAlignmentInput{"TACGG",
		"CACG",
		1.0, 0.5, 1.0},
		[][]float64{[]float64{0, 0, 0, 0, 0},
			[]float64{0, 0, 0, 0, 0},
			[]float64{0, 0, 1, 0, 0},
			[]float64{0, 1, 0, 2, 1},
			[]float64{0, 0, 0.5, 1, 3},
			[]float64{0, 0, 0, 0, 2}}},

	{LocalAlignmentInput{"ATCGATCGT",
		"ATCGGCTAC",
		1.0, 1.0, 0.5},
		[][]float64{[]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			[]float64{0, 1, 0.5, 0, 0, 0, 0, 0, 1, 0.5},
			[]float64{0, 0.5, 2, 1.5, 1, 0.5, 0, 1, 0.5, 0},
			[]float64{0, 0, 1.5, 3, 2.5, 2, 1.5, 1, 0.5, 1.5},
			[]float64{0, 0, 1, 2.5, 4, 3.5, 3, 2.5, 2, 1.5},
			[]float64{0, 1, 0.5, 2, 3.5, 3, 2.5, 2, 3.5, 3},
			[]float64{0, 0.5, 2, 1.5, 3, 2.5, 2, 3.5, 3, 2.5},
			[]float64{0, 0, 1.5, 3, 2.5, 2, 3.5, 3, 2.5, 4},
			[]float64{0, 0, 1, 2.5, 4, 3.5, 3, 2.5, 2, 3.5},
			[]float64{0, 0, 1, 2, 3.5, 3, 2.5, 4, 3.5, 3}}},

	{LocalAlignmentInput{"ATCGATCGT",
		"AAC",
		1.0, 1.0, 0.5},
		[][]float64{[]float64{0, 0, 0, 0},
			[]float64{0, 1, 1, 0.5},
			[]float64{0, 0.5, 0.5, 0},
			[]float64{0, 0, 0, 1.5},
			[]float64{0, 0, 0, 1},
			[]float64{0, 1, 1, 0.5},
			[]float64{0, 0.5, 0.5, 0},
			[]float64{0, 0, 0, 1.5},
			[]float64{0, 0, 0, 1},
			[]float64{0, 0, 0, 0.5}}}}

func TestLocalScoreTable(t *testing.T) {
	for _, pair := range localScoreTests {
		v := LocalScoreTable(pair.input.str1, pair.input.str2, pair.input.match, pair.input.mismatch, pair.input.gap)
		if !reflect.DeepEqual(v, pair.scoreMatrix) {
			t.Error(
				"For", pair.input,
				"expected scoring matrix", pair.scoreMatrix,
				"got", v,
			)
		}
	}
}

/********************************************
 Local Alignment Tests
*********************************************/

type Solution struct {
	score  float64
	start0 int
	end0   int
	start1 int
	end1   int
}

type testpair struct {
	input LocalAlignmentInput
	sol   Solution
}

var tests = []testpair{
	{LocalAlignmentInput{"GAAC",
		"CAAG",
		1.0, 1.0, 0.5},
		Solution{2.000, 1, 3, 1, 3}},

	{LocalAlignmentInput{"TAAC",
		"TAAG",
		1.0, 1.0, 0.5},
		Solution{3.000, 0, 3, 0, 3}},

	{LocalAlignmentInput{"A",
		"A",
		1.0, 0.5, 1.0},
		Solution{1.000, 0, 1, 0, 1}},

	{LocalAlignmentInput{"TACGG",
		"CACGTG",
		1.0, 1.0, 0.5},
		Solution{4.500, 1, 5, 1, 6}},

	{LocalAlignmentInput{"AGATCGG",
		"AGCGGATTTGCC",
		1.0, 1.0, 0.5},
		Solution{6.000, 0, 7, 0, 5}},

	{LocalAlignmentInput{"AGATCGG",
		"AGTTGATTTGCCG",
		1.0, 0.5, 1.0},
		Solution{3.000, 1, 4, 4, 7}},

	{LocalAlignmentInput{"AGATCGG",
		"AGTCAGCGG",
		1.0, 1.0, 0.5},
		Solution{7.500, 0, 7, 0, 8}},

	{LocalAlignmentInput{"GCGATCGAT",
		"GCGCGATCTTGAT",
		1.0, 1.0, 0.5},
		Solution{10.000, 0, 9, 2, 13}}}

func TestLocalAlignment(t *testing.T) {
	for _, pair := range tests {
		optAlignment, start0, end0, start1, end1 := LocalAlignment(pair.input.str1, pair.input.str2, pair.input.match, pair.input.mismatch, pair.input.gap)
		score := computeScore(optAlignment, pair.input.match, pair.input.mismatch, pair.input.gap)
		v := Solution{score, start0, end0, start1, end1}
		if !reflect.DeepEqual(v, pair.sol) {
			t.Error(
				"For", pair.input,
				"expected alignment with score", pair.sol.score,
				"got", optAlignment, "with score", score,
			)
		}
	}
}

/********************************************
 Shared k-mers Tests
*********************************************/

type sharedKMersTestpair struct {
	str1        string
	str2        string
	k           int
	sharedKmers int
}

var sharedKMersTests = []sharedKMersTestpair{
	{"ATGCGGCTAGCTTAGCCTAGATCGATCGGCTAGCTAGCTAGCCGAGGCTCTCGATCGATCGCGCTAGG",
		"ATGCGGCTGGCTTAGCCTAGTTCGATCGCGTTCGTAGCTATAGAGCTAGCTAGATCGATCGCGCTAGG", 3, 51},
	{"TTGCGGAGCTAGGGATCCGATCGAATATCGATATTCGATCGGGAACACAGATCGAT",
		"GGTACATCGATTCTAGATTCTATAGCGCGCTTCGATCGATTCGATCGATCGAAAAG", 1, 52},
	{"TTGCGGAGCTAGGGATCCGATCGAATATCGATATTCGATCGGGAACACAGATCGAT",
		"TTGCGGAGCTAGGGATCCGATCGAATATCGATATTCGATCGGGAACACAGATCGAT", 2, 55},
	{"TTGCGGAGCTAGGGATCCGATCGAATATCGATATTCGATCGGGAACACAGATCGAT",
		"TTGCGGAGCTAGGGATCCGATCGAATATCGATATTCGATCGGGAACACAGATCGAT", 50, 7},
	{"TTGCGGAGCTAGGGATCCGATCGAATATCGATATTCGATCGGGAACACAGATCGAT",
		"TTGCGGAGCTAGGGATCCGATCGAATATCGATATAGCATCGGGAACACAGATCGAT", 10, 35},
	{"AGCT",
		"TCGA", 1, 4},
	{"AATT",
		"CCGG", 5, 0},
	{"AGCT",
		"TGAC", 2, 0}}

func TestCountSharedKmers(t *testing.T) {
	for _, pair := range sharedKMersTests {
		v := CountSharedKmers(pair.str1, pair.str2, pair.k)
		if v != pair.sharedKmers {
			t.Error(
				"For", pair.str1,
				"and", pair.str2,
				"with k-mers of length", pair.k,
				"expected", strconv.Itoa(pair.sharedKmers),
				"got", strconv.Itoa(v),
			)
		}
	}
}

/********************************************
 LCS Tests
*********************************************/
type lcsTest struct {
	string1 string
	string2 string
	answer  string
}

var lcsTests = []lcsTest{
	{"GACT", "ATG", "AT"},
	{"ACTGAG", "GACTGG", "ACTGG"},
	{"AC", "AC", "AC"},
	{"GGGGT", "CCCCT", "T"},
	{"TCCCC", "TGGGG", "T"},
	{"AA", "CGTGGAT", "A"},
	{"GGTGACGT", "CT", "CT"},
}

func TestLongestCommonSubsequence(t *testing.T) {
	for _, test := range lcsTests {
		output := LongestCommonSubsequence(test.string1, test.string2)
		if output != test.answer {
			t.Error("For", test.string1, "and", test.string2, "expect LCS", test.answer, "but got", output)
		}
	}
}
