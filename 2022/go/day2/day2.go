package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readFromFile(path string) ([]string,[]string,  error) {
    file, err := os.Open(path)
    if err != nil {
        fmt.Println("failed to read from file")
    }
    defer file.Close()

    var elfChoices []string
    var myChoices []string

    scanner := bufio.NewScanner(file)
    for scanner.Scan(){
        line := scanner.Text()
        lineSplit := strings.Split(line, " ")
        elfChoice, myChoice := lineSplit[0], lineSplit[1]
        elfChoices = append(elfChoices, elfChoice)
        myChoices = append(myChoices, myChoice)
    }

    return elfChoices, myChoices, scanner.Err()
}

func makeChoice(elfChoice, myChoice string) string {
    winMap := map[string]string {
        "A": "Y",
        "B": "Z",
        "C": "X",
    }
    loseMap := map[string]string {
        "A": "Z",
        "B": "X",
        "C": "Y",
    }
    drawMap := map[string]string {
        "A": "X",
        "B": "Y",
        "C": "Z",
    }
    switch myChoice {
    case "X":
        return loseMap[elfChoice]
    case "Y":
        return drawMap[elfChoice]
    case "Z":
        return winMap[elfChoice]
    }
    return ""
    
}

func calculateScore(elfChoice, myChoice string) int {
    choiceMap := map[string]string{
        "X": "C",
        "Y": "A",
        "Z": "B",
    }
    choiceWrights := map[string]int {
        "X": 1,
        "Y": 2,
        "Z": 3,
        "A": 1,
        "B": 2,
        "C": 3,
    }
    if choiceWrights[elfChoice] == choiceWrights[myChoice] {
        return 3 + choiceWrights[myChoice]
    } else if choiceMap[myChoice] == elfChoice {
        return 6 + choiceWrights[myChoice]
    } else {
        return choiceWrights[myChoice]
    }
}

func main() {
    args := os.Args[1:]
    if len(args) != 1 {
        log.Fatalf("please provide only one argument i.e path to the input")
    }

    elfChoices, myChoices, err := readFromFile(args[0])
    if err != nil {
        log.Fatal("failed to read from file")
    }


    var sum int
    for index := range elfChoices {
        elfChoice := elfChoices[index]
        matchResult := myChoices[index]
        myChoice := makeChoice(elfChoice, matchResult)
        fmt.Printf("elf choice: %s my choice: %s\n", elfChoice, myChoice)
        sum += calculateScore(elfChoice, myChoice)
        fmt.Println("sum is ", sum)
    }
    
    fmt.Printf("total sum is: %d\n", sum)
}
