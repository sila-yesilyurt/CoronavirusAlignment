package Functions

//EditDistanceMatrix takes a slice of strings as input. It returns a matrix of integers
//whose (i,j)-th value is the edit distance between the i-th and j-th strings in
//the input slice.
func EditDistanceMatrix(patterns []string) [][]int {
	numRows := len(patterns)

	//create an n x n matrix of values
	mtx := make([][]int, numRows)
	for i := range mtx {
		mtx[i] = make([]int, numRows)
	}

	// let's set the values now
	for i := range patterns {
		for j := i + 1; j < len(patterns); j++ {
			d := EditDistance(patterns[i], patterns[j])
			//set two values of the matrix
			mtx[i][j] = d
			mtx[j][i] = d
		}
	}

	return mtx
}
