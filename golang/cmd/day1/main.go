package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func totalCaloriesByElf(currentElf []int) int {
	total := 0
	for _, itemCalories := range currentElf {
		total += itemCalories
	}

	return total
}

func main() {
	var caloriesSlice [][]int
	var currentElf []int
	var caloriesByElf []int

	calories, err := os.ReadFile("/home/rsegt/personal/adventofcode/calories.txt")
	check(err)

	lines := strings.Split(string(calories), "\n")

	for _, line := range lines {
		if line == "" {
			if len(currentElf) > 0 {
				caloriesSlice = append(caloriesSlice, currentElf)
				caloriesByElf = append(caloriesByElf, totalCaloriesByElf(currentElf))
				currentElf = nil
			}
		} else {
			num, err := strconv.Atoi(line)
			check(err)

			currentElf = append(currentElf, num)
		}
	}

	if len(currentElf) > 0 {
		caloriesSlice = append(caloriesSlice, currentElf)
		caloriesByElf = append(caloriesByElf, totalCaloriesByElf(currentElf))
	}

	highestAmountCaloriesCarried := 0
	caloriesSum := 0

	for _, innerSlice := range caloriesSlice {
		for _, value := range innerSlice {
			caloriesSum += value
		}

		if caloriesSum > highestAmountCaloriesCarried {
			highestAmountCaloriesCarried = caloriesSum
		}

		caloriesSum = 0
	}

	sort.Slice(caloriesByElf, func(i, j int) bool {
		return caloriesByElf[i] > caloriesByElf[j]
	})

	caloriesTop3 := 0

	for _, amount := range caloriesByElf[:3] {
		caloriesTop3 += amount
	}

	// fmt.Println(highestAmountCaloriesCarried)
	fmt.Println(caloriesTop3)
}
