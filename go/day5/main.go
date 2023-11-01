package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Stack[T any] struct {
    elems []T
}

func (s *Stack[T]) Push(elem T) {
    s.elems = append(s.elems, elem)
}

func (s *Stack[T]) Pop() T {
    length := len(s.elems)
    if length < 0 {
        log.Fatalf("the stack is empty")
    }
    out := s.elems[length - 1]
    s.elems = s.elems[:length - 1]
    return out
}

func readFromFile(path string) []string {
    file, err := os.Open(path)
    if err != nil {
        log.Fatalf("encountered error %v when reading from file", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var out []string
    for scanner.Scan() {
        line := scanner.Text()
        out = append(out, line)
    }

    return out
}

func parseInputData(data []string) ([]string, []string) {
    var stacks []string
    var operations []string
    delimiterIndex := slices.Index(data, "")
    stacks = data[:delimiterIndex]
    operations = data[delimiterIndex + 1:]
    return stacks, operations
}

func parseStacks(inputStacks []string) []Stack[string] {
    length := len(inputStacks)
    buckets := strings.Split(inputStacks[length-1], "")
    numberOfStacks := 0
    for _, char := range buckets {
        _, err := strconv.Atoi(char)
        if err == nil {
           numberOfStacks += 1
       }
    }
    outputStacks := make([]Stack[string], numberOfStacks)
    for i := length - 2; i >= 0; i-- {
        for index, elem := range inputStacks[i] {
            if string(elem) == "[" {
                stackIndex := index / 4
                outputStacks[stackIndex].Push(string(inputStacks[i][index + 1]))
            }
        }

    }
    fmt.Println("parsed stacks")
    return outputStacks
}

func parseOperations(inputOperations []string) [][]int {
    outputOperations := make([][]int, 0)
    for _, operation := range inputOperations {
        tmp := strings.Split(operation, " ")
        quantity, err := strconv.Atoi(tmp[1])
        if err != nil {
            log.Fatalf("failed to parse quantity for %v", operation)
        }
        source, err := strconv.Atoi(tmp[3])
        if err != nil {
            log.Fatalf("failed to parse source for %v", operation)
        }
        destination, err := strconv.Atoi(tmp[5])
        if err != nil {
            log.Fatalf("failed to parse destination for %v", operation)
        }
        op := []int{quantity, source - 1, destination - 1}
        outputOperations = append(outputOperations, op)
    }
    return outputOperations
}

func computeOperations(stacks []Stack[string], operations [][]int) string {
    for _, operation := range operations {
        quantity, source, destination := operation[0], operation[1], operation[2]
        // the crates need to be moved in order for part 2
        crates := make([]string, 0)
        for quantity > 0 {
            crate := stacks[source].Pop()
            crates = append(crates, crate)

            quantity -= 1
        }
        length := len(crates) - 1
        for index := range(crates) {
            stacks[destination].Push(crates[length - index])
        }
    }
    out := ""
    for _, stack := range stacks {
        out += stack.Pop()
    }
    log.Println("finished the computation")
    return out
}

func main()  {
    args := os.Args[1:]
    if len(args) != 1 {
        log.Fatalf("too many args")
    }
    filePath := args[0]
    rawData := readFromFile(filePath)
    inputStacks, inputOperations := parseInputData(rawData)
    outputStacks := parseStacks(inputStacks)
    for _, stack := range outputStacks {
        fmt.Println(stack)
    }
    outputOperations := parseOperations(inputOperations)
    fmt.Println(computeOperations(outputStacks, outputOperations))
}
