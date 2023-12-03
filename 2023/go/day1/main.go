package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)
func readInput(filePath string) ([] string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        log.Fatalf("error opening the file")
        return nil, err
    }
    defer file.Close()
    var out [] string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        out = append(out, scanner.Text())
    }
    return out, nil
}

func extractDigits(line string) [] int {
    splitLine := strings.Split(line, "")
    out := make([]int, 0)
    for _, char := range splitLine {
        value, err := strconv.Atoi(char)
        if err != nil {
            continue
        }
        out = append(out, value)
    }
    return out
}

func main() {
    args := os.Args[1:]
    if len(args) != 1 {
        log.Fatalf("too many arguments given")
        return
    }

    input, err := readInput(args[0])
    if err != nil {
        log.Fatalf("failed parsing the input")
        return
    }
    result := 0
    for _, line := range input {
        digits := extractDigits(line)
        if len(digits) == 0 {
            continue
        }
        first, last := digits[0], digits[len(digits) - 1]
        lineValue := first * 10 + last
        fmt.Println("line: ", digits, " => ", lineValue)
        result += lineValue
    }
    fmt.Println(result)
}
