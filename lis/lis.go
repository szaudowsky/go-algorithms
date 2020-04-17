package lis

// Lis is a function to find longest icreasing subsequence.
func Lis(array []int) []int {
	if len(array) < 1 {
		return array
	}

	results := make([][]int, len(array))
	for i := 0; i < len(array); i++ {
		results[i] = []int{array[i]}
	}

	for i := range array {
		helperSlice := make([]int, i)
		for j := 0; j < i; j++ {
			if array[j] < array[i] && len(results[j]) > len(results[i])-1 {
				helperSlice = append(results[j], results[i][len(results[i])-1])
				results[i] = helperSlice
			}
		}
	}

	longest := max(results)
	return longest
}

func max(list [][]int) []int {
	var l = list[0]
	var max = len(list[0])

	for i := range list {
		if max < len(list[i]) {
			l = list[i]
		}
	}
	return l
}
