package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFromFile(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
       log.Fatalf("failed to open file %e", err) 
    }
    defer file.Close()

    text := make([]string, 0)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        text = append(text, scanner.Text())
    }
    return text, scanner.Err()
}

func commonChar(comp1, comp2 string) string {
    for _, c1 := range comp1 {
        for _, c2 := range comp2 {
            if c1 == c2 {
               return string(c1) 
            }
        }
    }
    return ""
}

func generatePriorityMap() map[string]int {
    priorities := make(map[string]int)
    index := 1 
    for ch := 'a'; ch <= 'z'; ch++ {
        priorities[string(ch)] = index
        index += 1
    }

    for ch := 'A'; ch <= 'Z'; ch++ {
        priorities[string(ch)] = index
        index += 1
    }
    return priorities
}

func main()  {
    args := os.Args[1:]
    if len(args) > 1 {
        log.Fatalf("wrong number of args given")
    }

    in, err := readFromFile(args[0])
    if err != nil {
        log.Fatalf("failed to read from file: %e", err)
    }

    priorities := generatePriorityMap()
   
    sum := 0
    for _, bag := range in {
        length := len(bag)
        comp1 := bag[:(length/2)]
        comp2 := bag[(length/2):]
        common := commonChar(comp1, comp2)
        sum += priorities[common]
    }
    fmt.Println(sum)
}
