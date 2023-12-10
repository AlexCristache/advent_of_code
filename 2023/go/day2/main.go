package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func setLogger() log.Logger {
	return *log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
}

var logger = setLogger()

func fetchInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		logger.Fatal("failed to open the input file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func parseGameRegEx(line string) int {
	gameRegEx := regexp.MustCompile(`\d+`)
	game := gameRegEx.FindString(line)
	gameId, err := strconv.Atoi(string(game))
	if err != nil {
		log.Fatalf("failed to parse game %s", game)
	}
	return gameId
}

func parseCube(line string) []string {
	cubeRegEx := regexp.MustCompile(`\d+ \w+`)
	cubes := cubeRegEx.FindAllString(line, -1)
	return cubes
}

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		log.Fatal("please provide only one argument")
	}
	data, err := fetchInput(args[0])
	if err != nil {
		logger.Fatal("failed reading from file", err)
	}
    sum := 0
    for _, line := range data {
        availableCubes := map[string]int{
            "green": 0,
            "blue":  0,
            "red":   0,
        }
        lineSplit := strings.Split(line, ":")
        game := lineSplit[0]
        setsRaw := lineSplit[1]
        gameId := parseGameRegEx(game)
        if err != nil {
            log.Fatalf("failed converting game id to int")
        }
        sets := strings.Split(setsRaw, ";")
        for _, set := range sets {
            cubes := parseCube(string(set))
            for _, cube := range cubes {
                cubeSplit := strings.Split(cube, " ")
                cubeColor := cubeSplit[1]
                cubeNo, err := strconv.Atoi(cubeSplit[0])
                if err != nil {
                    log.Fatalf("failed parsing cube no to int")
                }
                if availableCubes[cubeColor] < cubeNo {
                    availableCubes[cubeColor] = cubeNo
                }
            }
        }
        power := 1
        for _, no := range availableCubes {
           power *= no
        }
        logger.Println("game", gameId, ":", power)
        sum += power
    }
    log.Println("valid games sum: ", sum)
}
