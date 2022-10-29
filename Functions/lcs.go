package Functions

//LongestCommonSubsequence takes two strings as input.
//It returns a longest common subsequence of the two strings.
func LongestCommonSubsequence(str0, str1 string) string {
	backtrack := LCSBacktrack(str0, str1) // this will be a matrix storing my "pointers"
	return OutputLCS(str0, str1, backtrack)
}

//OutputLCS takes two strings and a matrix of (LCS) backtracking pointers.
//It returns an LCS of the two strings.
func OutputLCS(str0, str1 string, backtrack [][]string) string {
	s := ""

	//idea: start at sink, backtrack, and add any match symbols we encounter
	//to the start of s

	for len(str0) > 0 || len(str1) > 0 { // while the strings have symbols
		row := len(str0)
		col := len(str1)

		if backtrack[row][col] == "UP" {
			//chop off the last symbol of str0
			str0 = str0[:len(str0)-1]
		} else if backtrack[row][col] == "LEFT" {
			//chop off the last symbol of str1
			str1 = str1[:len(str1)-1]
		} else if backtrack[row][col] == "DIAG" {
			//was this a match? If so, add to our string :)
			if str0[len(str0)-1] == str1[len(str1)-1] {
				// we have a match!
				s = string(str0[len(str0)-1]) + s
			}
			//chop off two symbols
			str0 = str0[:len(str0)-1]
			str1 = str1[:len(str1)-1]
		} else {
			panic("Error: non-standard backtracking pointer.")
		}
	}

	return s
}

//LCSBacktrack takes two strings as input.
//It returns a 2-D slice of strings storing backtrack pointers ("up", "left", "diag")
//for finding the longest common subsequence of the two strings.
func LCSBacktrack(str0, str1 string) [][]string {
	if len(str0) == 0 || len(str1) == 0 {
		panic("Blah")
	}

	numRows := len(str0) + 1
	numCols := len(str1) + 1

	backtrack := make([][]string, numRows)

	for i := range backtrack {
		backtrack[i] = make([]string, numCols)
	}

	//grab the scores from the LCS scoring table
	scoreTable := LCSScoreMatrix(str0, str1)

	//now, use the scoring table to set backtracking pointers

	//first, easy cases: 0-th row and column
	for j := 1; j < numCols; j++ {
		backtrack[0][j] = "LEFT"
	}
	for i := 1; i < numRows; i++ {
		backtrack[i][0] = "UP"
	}

	//traverse rest of scoring table
	for i := 1; i < numRows; i++ {
		for j := 1; j < numCols; j++ {
			// which value was used to produce scoring matrix at (i, j)?
			if scoreTable[i][j] == scoreTable[i-1][j] {
				backtrack[i][j] = "UP"
			} else if scoreTable[i][j] == scoreTable[i][j-1] {
				backtrack[i][j] = "LEFT"
			} else {
				// I should be more precise, making sure that this value is indeed met, and throwing an
				// error if not
				backtrack[i][j] = "DIAG"
			}
		}
	}

	return backtrack
}
