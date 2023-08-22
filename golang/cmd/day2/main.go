package main

import (
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
    score := 0

    gameResults := map[string]int{
        "A X": 4,
        "A Y": 8,
        "A Z": 3,
        "B X": 1,
        "B Y": 5,
        "B Z": 9,
        "C X": 7,
        "C Y": 2,
        "C Z": 6,
    }

    strategy, err := os.ReadFile("/home/rsegt/personal/adventofcode/strategy.txt")
   	check(err)

    lines := strings.Split(string(strategy), "\n")

   for _, line := range lines {
        println(line)

        score += gameResults[line]
    }

    println(score)
}
