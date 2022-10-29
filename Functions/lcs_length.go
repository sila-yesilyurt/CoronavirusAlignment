package Functions

//LCSLength takes two strings as input. It returns the length of a longest common
//subsequence of the two strings.
func LCSLength(str1, str2 string) int {
	if len(str1) == 0 || len(str2) == 0 {
		panic("boo")
	}

	//if we're here, we know that both strings are nonempty

	scoringMatrix := LCSScoreMatrix(str1, str2)

	//return the lower right corner
	return scoringMatrix[len(str1)][len(str2)]
}

//LCSScoreMatrix takes two strings as input.
//It returns the scoring matrix for longest common subsequence using dynamic programming.
func LCSScoreMatrix(str1, str2 string) [][]int {
	// let's make our matrix

	numRows := len(str1) + 1
	numCols := len(str2) + 1

	//make our matrix

	scoringMatrix := make([][]int, numRows)
	for i := range scoringMatrix {
		//make cols
		scoringMatrix[i] = make([]int, numCols)
	}

	//the 0-th row and column have default zero values which is good :)

	//range over the rest of the matrix, starting at (1, 1)
	for i := 1; i < numRows; i++ {
		for j := 1; j < numCols; j++ {
			//set the value at (i,j)
			up := scoringMatrix[i-1][j]
			left := scoringMatrix[i][j-1]
			diag := scoringMatrix[i-1][j-1]

			//if there was a match, increment diag
			if str1[i-1] == str2[j-1] {
				diag++
			}

			// now I have set my three values, so take their max!
			scoringMatrix[i][j] = Max(up, left, diag)
		}
	}

	return scoringMatrix
}

// we need a function to take the max of integers.
// variadic functions take arbitrary number of inputs
// e.g., fmt.Println(blah, bleeh, foo, bar )

func Max(nums ...int) int {
	//regardless of how many inputs we give, these inputs are converted into
	//an array called "nums"
	if len(nums) == 0 {
		panic("no")
	}
	m := nums[0]

	//let's range over array nums and see if something is larger
	for i := 1; i < len(nums); i++ {
		if nums[i] > m {
			m = nums[i] // update new max
		}
	}

	return m
}
