package day1

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	math "aoc/internal/math"
)

func Run() error {
	leftSlice, rightSlice, err := readList("assets/day1_list")
	if err != nil {
		return fmt.Errorf("failed to read list: %w", err)
	}

	fmt.Println("sortedDiff:", calcListDiff(leftSlice, rightSlice))
	fmt.Println("similarity:", calcSimilarity(leftSlice, rightSlice))

	return nil
}

func calcSimilarity(left, right []int) int {
	rmap := make(map[int]int)

	for _, r := range right {
		_, ok := rmap[r]
		if ok {
			rmap[r] += 1
		} else {
			rmap[r] = 1
		}
	}

	sum := 0
	for _, l := range left {
		value, ok := rmap[l]
		if ok {
			sum += l * value
		}
	}

	return sum
}

func calcListDiff(left, right []int) int {
	slices.Sort(left)
	slices.Sort(right)

	sum := 0
	for i := range left {
		sum += math.AbsInt(left[i] - right[i])
	}

	return sum
}

func readList(path string) ([]int, []int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open: %w", err)
	}
	defer file.Close()

	var leftSlice, rightSlice []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		elements := strings.Split(scanner.Text(), "   ")

		left, err := strconv.Atoi(elements[0])
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse left: %w", err)
		}
		leftSlice = append(leftSlice, left)

		right, err := strconv.Atoi(elements[1])
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse right: %w", err)
		}
		rightSlice = append(rightSlice, right)
	}

	return leftSlice, rightSlice, nil
}
