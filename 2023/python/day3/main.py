import re
import sys
from typing import List


def readInput(path: str) -> List[str]:
    with open(path) as f:
        data = f.readlines()
    return data


def checkSymbols(data: List[str]):
    validPositions = set()
    neighbours = [(-1, -1), (-1, 0), (-1, 1), (0, -1), (0, 1), (1, 1), (1, 0), (1, -1)]
    for i, line in enumerate(data):
        line = line.strip()
        for j, char in enumerate(line):
            if char == "*":
                for neighbour in neighbours:
                    x, y = neighbour
                    validPositions.add((i + x, j + y))
    return validPositions


def main(input_path: str):
    data = readInput(input_path)
    numberPattern = re.compile(r"\d+")
    sum = 0
    validPositions = checkSymbols(data)
    for idx, x in enumerate(data):
        numbers = re.findall(numberPattern, x)
        totalMatches = 0
        index = 0
        for number in numbers:
            start, length, value = x.find(number, index), len(number), int(number)
            index = start + length - 1
            for a in range(length):
                if (idx, start + a) in validPositions:
                    totalMatches += value
                    break
        sum += totalMatches
    print(sum)

def findSecondGear(gearX, gearY):
    pass


def part2(input_path: str):
    """
    - iterate over each line until you find a number
    - record start position of said number
    """
    data = readInput(input_path)
    numberPattern = re.compile(r"\d+")
    validPositions = checkSymbols(data)
    for i, line in enumerate(data):
        for j, char in enumerate(line):
            if (i, j) in validPositions and char.isdigit():
                print(char)


if __name__ == "__main__":
    args = sys.argv[1:]
    main(args[0])
    part2(args[0])
