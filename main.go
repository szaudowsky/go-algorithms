package main

import (
	"fmt"
	"go-algorithms/lis"
)

func main() {
	longest := lis.Lis([]int{10, 22, 9, 33, 21, 50, 41, 60})
	fmt.Print(longest)
}
