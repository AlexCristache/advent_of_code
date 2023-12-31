package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
    "slices"

)

func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines [] string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

func getMaxCalories(elves map[int]int) int {
    max_calories := 0 

    for _, elf := range elves {
        if elf > max_calories {
            max_calories = elf        
        }
    }

    return max_calories
}

func getTopCaloriesSum(elves map[int]int) int  {
    calories := make([]int, len(elves))

    for index, value := range elves {
        calories[index] =  value
    }

    slices.Sort(calories)
    slices.Reverse(calories)
    return calories[0] + calories[1] + calories[2]
 
}

func main()  {
    args := os.Args[1:]
    
    lines, err := readLines(args[0])
    if err != nil {
        log.Fatalf("readLines: %s", err)
    }

    elves := make(map[int]int)
    index := 0


    for _, line := range lines {
        if line == "" {
            index += 1
            continue
        }
        val, err := strconv.Atoi(line)
        if err != nil {
            log.Fatalf("failed to convert to int value %s", line)
        }
        elves[index] += val 
        //fmt.Println(i, line)
    }
    
    max_calories := getMaxCalories(elves)

    fmt.Printf("the max calories are %d\n", max_calories)

    calories := getTopCaloriesSum(elves)
    fmt.Println(calories)


}
