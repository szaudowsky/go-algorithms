package lis

import (
	"fmt"
	"testing"
)

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestLis(t *testing.T) {
	var tests = []struct {
		array []int
		want  []int
	}{
		{
			[]int{},
			[]int{},
		},
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{3, 1},
			[]int{3},
		},
		{
			[]int{1, 9, 5, 13, 3, 11, 7, 15, 2, 10, 6, 14, 4, 12, 8, 16},
			[]int{1, 5, 7, 10, 14, 16},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("array : %d want: %d", tt.array, tt.want)
		t.Run(testname, func(t *testing.T) {
			longest := Lis(tt.array)
			fmt.Print(longest)
			if !Equal(longest, tt.want) {
				t.Errorf("got: %d, want: %d", longest, tt.want)
			}
		})
	}

}
