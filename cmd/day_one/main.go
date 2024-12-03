package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var left, right []int

func main() {
	left, right = readInput()
	fmt.Println(len(left))
	fmt.Println(len(right))

	sort.Ints(left)
	sort.Ints(right)

	distance()
}

func distance() {
	totalDistance := 0
	// both 1k elements, range doesn't matter
	for i := 0; i < len(left); i++ {
		if left[i] > right[i] {
			totalDistance = totalDistance + left[i] - right[i]
		} else {
			totalDistance = totalDistance + right[i] - left[i]
		}
	}

	fmt.Println(totalDistance)
}

func readInput() ([]int, []int) {
	file, err := os.Open("/Users/kieran.barty/Documents/repos/personaldev/advent-of-code-24/cmd/day_one/input.txt")

	left, right := []int{}, []int{}

	if err != nil {
		fmt.Println("err", err)
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//var l, r int
		split := strings.Split(scanner.Text(), "   ")
		l, _ := strconv.Atoi(split[0])
		r, _ := strconv.Atoi(split[1])

		left = append(left, l)
		right = append(right, r)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return left, right
}
