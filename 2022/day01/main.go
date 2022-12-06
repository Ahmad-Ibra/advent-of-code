package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello")

	groups, err := inputSplitByGroup("day01/input.txt")
	if err != nil {
		fmt.Println(err.Error())
	}

	max := math.MinInt

	for _, group := range groups {
		line := strings.Split(group, "\n")
		count := 0
		for _, l := range line {
			if l == "" {
				continue
			}
			iNum, err := strconv.Atoi(l)
			if err != nil {
				fmt.Println(err.Error())
			}
			count += iNum
		}
		if count > max {
			max = count
		}
	}

	fmt.Println("part1 answer is " + strconv.Itoa(max))
}

func inputSplitByGroup(fileLoc string) ([]string, error) {

	dat, err := os.ReadFile(fileLoc)
	if err != nil {
		return nil, err
	}

	groups := strings.Split(string(dat), "\n\n")

	for i, group := range groups {
		fmt.Println("\ngroup " + strconv.Itoa(i) + ":\n" + group)
	}

	return groups, nil
}
