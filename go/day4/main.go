package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(path string) ([]string, error) {
    var out []string
    file, err := os.Open(path)
    if err != nil {
        log.Fatalf("failed to open file")
    }

    scanner := bufio.NewScanner(file)
    defer file.Close()

    for scanner.Scan() {
        out = append(out, scanner.Text())
    }

    return out, scanner.Err()
}

func parseLine(line string) []int {

    elf1, elf2 := strings.Split(line, ",")[0], strings.Split(line, ",")[1]

    start1, end1 := strings.Split(elf1, "-")[0], strings.Split(elf1, "-")[1]
    start2, end2 := strings.Split(elf2, "-")[0], strings.Split(elf2, "-")[1]

    start1Int, err := strconv.Atoi(start1) 
    if err != nil {
        log.Fatal("failed to parse start1")
    }
    start2Int, err := strconv.Atoi(start2) 
    if err != nil {
        log.Fatal("failed to parse start2")
    }
    end1Int, err := strconv.Atoi(end1) 
    if err != nil {
        log.Fatal("failed to parse end1")
    }
    end2Int, err := strconv.Atoi(end2) 
    if err != nil {
        log.Fatal("failed to parse end2")
    }

    out := []int {
        start1Int,
        end1Int,
        start2Int,
        end2Int,
    }

    return out
}

func isOverlapping(sections []int) bool  {
    section1 := sections[:2]
    section2 := sections[2:]


    if section2[0] >= section1[0] && section2[0] <= section1[1] {
        return true
    }
    if section1[0] >= section2[0] && section1[0] <= section2[1] {
        return true
    }

    if section2[1] >= section1[0] && section2[1] <= section1[1] {
        return true
    }
    if section1[1] >= section2[0] && section1[1] <= section2[1] {
        return true
    }

    return false
}

func main()  {
    args := os.Args
    if len(args) < 2 || len(args) > 2 {
        log.Fatal("please provide only one argument as input file")
    }

    lines, err := readInput(args[1])
    if err != nil {
        log.Fatal("encountered error when reading from file", err)
    }

    overlappingSections := 0

    for _, line := range lines {
        sections := parseLine(line)
        if isOverlapping(sections) {
            overlappingSections += 1
        }
    }
    fmt.Println(overlappingSections)
}
