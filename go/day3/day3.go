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

func findCommonChar(group []string) string {
    var commonChar string
    backpack1 := group[0]
    backpack2 := group[1]
    backpack3 := group[2]
    for _, c1 := range backpack1 {
        for _, c2 := range backpack2 {
            for _, c3 := range backpack3 {
                if c1 == c2  && c1 == c3 {
                    commonChar = string(c1)
                }
            }
        }
    }
    return commonChar
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
/*    var group []string
    for index, bag := range in {
        group = append(group, bag)
        if index % 3 == 2 && index > 0 {
            start := int(index / 3)
            end := start + 3
            commonChar := findCommonChar(group[start:end])
            sum += priorities[commonChar]
        }

    }
*/
    for len(in) > 0 {
        commonChar := findCommonChar(in[:3])
        in = in[3:]
        fmt.Println(commonChar)
        sum += priorities[commonChar]
    }
    fmt.Println(sum)
}
