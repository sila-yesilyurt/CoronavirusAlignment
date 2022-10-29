package Functions

//Alignment is an array of two strings corresponding to the top and bottom
//rows of an alignment of two strings.
type Alignment [2]string

//GlobalAlignment takes two strings, along with match, mismatch, and gap scores.
//It returns a maximum score global alignment of the strings corresponding to these penalties.
func GlobalAlignment(str0, str1 string, match, mismatch, gap float64) Alignment {
	//replace this with your code
	backtrack := GlobalBacktrack(str0, str1, match, mismatch, gap)
	optAlignment := OutputGlobalAlignment(str0, str1, backtrack)
	return optAlignment
}

func OutputGlobalAlignment(str0, str1 string, backtrack [][]string) Alignment {
	var a Alignment // array of two strings
	//a[0] = top row (str0), a[1] = bottom row (str1)

	//start at bottom right and work our way backward

	for len(str0) > 0 || len(str1) > 0 { // while there are symbols in either string
		//which case is satisfied?
		row := len(str0)
		col := len(str1)

		if backtrack[row][col] == "UP" {
			//use a symbol from str0 and align against gap
			a[0] = string(str0[row-1]) + a[0]
			a[1] = "-" + a[1]
			str0 = str0[:len(str0)-1]
		} else if backtrack[row][col] == "LEFT" {
			//use a symbol from str1 and align against gap
			a[0] = "-" + a[0]
			a[1] = string(str1[col-1]) + a[1]
			str1 = str1[:len(str1)-1]

		} else if backtrack[row][col] == "DIAG" {
			//use a symbol from both strings
			a[0] = string(str0[row-1]) + a[0]
			a[1] = string(str1[col-1]) + a[1]
			//shorten strings
			str0 = str0[:len(str0)-1]
			str1 = str1[:len(str1)-1]
		} else {
			panic("Illegal backtracking pointer.")
		}
	}

	return a
}

func GlobalBacktrack(str0, str1 string, match, mismatch, gap float64) [][]string {
	if len(str0) == 0 || len(str1) == 0 {
		panic("Zero length strings.")
	}

	numRows := len(str0) + 1
	numCols := len(str1) + 1

	backtrack := make([][]string, numRows)
	for i := range backtrack {
		backtrack[i] = make([]string, numCols)
	}

	// let's get the scoring matrix values
	scoreTable := GlobalScoreTable(str0, str1, match, mismatch, gap)

	//first, set backtracking pointers of the 0-th row and column
	for j := 1; j < numCols; j++ {
		backtrack[0][j] = "LEFT"
	}
	for i := 1; i < numRows; i++ {
		backtrack[i][0] = "UP"
	}

	//traverse rest of table, checking values in the score table and setting the pointers
	//appropriately.
	for i := 1; i < numRows; i++ {
		for j := 1; j < numCols; j++ {
			//which value was used?
			if scoreTable[i][j] == scoreTable[i][j-1]-gap {
				//looking left
				backtrack[i][j] = "LEFT"
			} else if scoreTable[i][j] == scoreTable[i-1][j]-gap {
				//looking up
				backtrack[i][j] = "UP"
			} else {
				//I should be more precise :)
				backtrack[i][j] = "DIAG"
			}
		}
	}

	return backtrack
}
