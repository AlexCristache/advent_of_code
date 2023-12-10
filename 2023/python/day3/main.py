from typing import List
import re
import sys


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
            if char != "." and not char.isdigit():
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


if __name__ == "__main__":
    args = sys.argv[1:]
    main(args[0])
