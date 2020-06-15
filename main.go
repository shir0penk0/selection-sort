package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := getArgs()
	fmt.Println("Befor: ", s)
	fmt.Println("After: ", sort(s))
}

func getArgs() []int {
	var s []int
	for _, v := range os.Args {
		j, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		s = append(s, j)
	}
	return s
}

func sort(s []int) []int {
	if len(s) == 0 {
		return []int{}
	}
	var min int
	var minIndex int
	for i, v := range s {
		if i == 0 || v < min {
			minIndex = i
			min = v
		}
	}
	rest := append(s[:minIndex], s[(minIndex+1):]...)
	return append([]int{min}, sort(rest)...)
}
