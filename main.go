package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	s := []int{}
	if terminal.IsTerminal(0) {
		s = getArgs()
	} else {
		stdin, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal("Invalid numer")
		}
		for _, v := range strings.Fields(string(stdin)) {
			i, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal("Invalid number")
			}
			s = append(s, i)
		}
	}
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
