package main

import "strings"

func main(){
    alphabet = "0abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

    rucksacks, err := os.ReadFile("/home/rsegt/personal/adventofcode/rucksacks.txt")
    lines := strings.Split(rucksacks, "\n")

    for _, rucksack := range lines {
        amountItems := len(string(rucksack))

        compartmentA := rucksack[:amountItems/2]
        compartmentB := rucksack[amountItems/2:]




    }
}
