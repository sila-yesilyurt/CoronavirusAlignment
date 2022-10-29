package Functions

//EditDistance takes two strings as input. It returns the Levenshtein distance
//between the two strings; that is, the minimum number of substitutions, insertions, and deletions
//needed to transform one string into the other.
func EditDistance(str1, str2 string) int {
	if len(str1) == 0 || len(str2) == 0 {
		panic("boo")
	}

	//if we're here, we know that both strings are nonempty

	scoringMatrix := EditMatrix(str1, str2)

	//return the lower right corner
	return scoringMatrix[len(str1)][len(str2)]
}

//EditMatrix takes two strings as input. It returns a matrix of values
//corresponding to edit distance, where (i, j) in the matrix is the edit distance
//between the substring of v up to the i-th symbol and the substring of w up to the j-th symbol.
func EditMatrix(str1, str2 string) [][]int {
	if len(str1) == 0 || len(str2) == 0 {
		panic("boo")
	}

	numRows := len(str1) + 1
	numCols := len(str2) + 1

	//make our matrix

	scoringMatrix := make([][]int, numRows)
	for i := range scoringMatrix {
		//make cols
		scoringMatrix[i] = make([]int, numCols)
	}

	// at this point in time, everyone has a default zero value

	// we should first set the 0-th row and column. each of these consists of only
	// insertions and deletions, so the i-th element's value should be i.
	// let's range over 0-th row
	for j := range scoringMatrix[0] {
		scoringMatrix[0][j] = j
	}

	// now, range over the 0-th column (i.e., range over rows)
	for i := range scoringMatrix {
		scoringMatrix[i][0] = i
	}

	//now, starting at (1, 1), range over the interior of matrix and fill the rest of the values
	//using dynamic programming.
	for row := 1; row < numRows; row++ {
		for col := 1; col < numCols; col++ {
			up := scoringMatrix[row-1][col] + 1
			left := scoringMatrix[row][col-1] + 1
			diag := scoringMatrix[row-1][col-1] // as of now, this assumes a match
			// if the diagonal edge corresponds to a mismatch, add 1 to it
			if str1[row-1] != str2[col-1] {
				diag++
			}
			//only thing that remains is to set current matrix value equal to the minimum
			//of these three values
			scoringMatrix[row][col] = Min(up, left, diag)
		}

	}
	return scoringMatrix
}

//Min is a variadic function that takes an arbitrary number of integers
//as input and returns their minimum.
func Min(nums ...int) int {
	//regardless of how many inputs we give, these inputs are converted into
	//an array called "nums"
	if len(nums) == 0 {
		panic("no")
	}
	m := nums[0]

	//let's range over array nums and see if something is larger
	for i := 1; i < len(nums); i++ {
		if nums[i] < m {
			m = nums[i] // update new max
		}
	}
	return m
}
