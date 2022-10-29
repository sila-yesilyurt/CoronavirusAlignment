package Functions

//GlobalScoreTable takes two strings and alignment penalties. It returns a 2-D array
//holding dynamic programming scores for global alignment with these penalties.
func GlobalScoreTable(str0, str1 string, match, mismatch, gap float64) [][]float64 {
	if len(str0) == 0 || len(str1) == 0 {
		panic("Blah")
	}

	numRows := len(str0) + 1
	numCols := len(str1) + 1

	// initialize scoring matrix -- would be great for subroutine

	scoreTable := make([][]float64, numRows)
	for i := range scoreTable {
		scoreTable[i] = make([]float64, numCols)
	}

	//first, penalize the 0-th row and column as all gaps
	for j := 1; j < numCols; j++ {
		//0-th row
		scoreTable[0][j] = float64(j) * (-gap)
	}
	for i := 1; i < numRows; i++ {
		//0th column
		scoreTable[i][0] = float64(i) * (-gap)
	}

	// now I am ready to range row by row and apply the GA recurrence relation
	for i := 1; i < numRows; i++ {
		for j := 1; j < numCols; j++ {
			//apply the recurrence relation
			upValue := scoreTable[i-1][j] - gap   //indel
			leftValue := scoreTable[i][j-1] - gap //indel
			var diagonalWeight float64
			if str0[i-1] == str1[j-1] { //match!
				diagonalWeight = match
			} else { // mismatch!
				diagonalWeight = -mismatch
			}
			diagValue := scoreTable[i-1][j-1] + diagonalWeight
			scoreTable[i][j] = MaxFloat(upValue, leftValue, diagValue)
		}
	}

	return scoreTable
}

func MaxFloat(nums ...float64) float64 {
	m := 0.0
	// nums gets converted to an array
	for i, val := range nums {
		if val > m || i == 0 {
			// update m
			m = val
		}
	}
	return m
}
