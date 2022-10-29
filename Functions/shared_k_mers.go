package Functions

//CountSharedKmers takes two strings and an integer k. It returns the number of
//k-mers that are shared by the two strings.
func CountSharedKmers(str1, str2 string, k int) int {
	count := 0

	freqMap1 := FrequencyMap(str1, k)
	freqMap2 := FrequencyMap(str2, k)

	for pattern := range freqMap1 {
		// just take the minimum
		count += Min2(freqMap1[pattern], freqMap2[pattern])
	}
	return count
}

//Min2 takes two integers as input and returns their minimum.
func Min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//FrequencyMap takes a string text and an integer k. It produces a map
//of all k-mers in the string to their number of occurrences.
func FrequencyMap(text string, k int) map[string]int {
	// map declaration is analogous to slices
	// (we don't need to give an initial length)
	freq := make(map[string]int)
	n := len(text)
	for i := 0; i < n-k+1; i++ {
		pattern := text[i : i+k]
		freq[pattern]++
	}
	return freq
}
