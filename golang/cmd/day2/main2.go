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
        "A X": 3,
        "A Y": 4,
        "A Z": 8,
        "B X": 1,
        "B Y": 5,
        "B Z": 9,
        "C X": 2,
        "C Y": 6,
        "C Z": 7,
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
