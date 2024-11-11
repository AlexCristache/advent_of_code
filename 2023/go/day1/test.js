const fs = require('fs');

const mappings = {
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9",
};

const file = fs.readFileSync('input.txt');

function searchAndReplace(line) {
    const keys = Object.keys(mappings);
    keys.forEach(key => {
        const first = key[0]
        const last = key[key.length - 1]
        line = line.replaceAll(key, first.concat(mappings[key], last));
    });
    return line;
}


const res = file.toString().trim().split("\n").map((line) => {
    return searchAndReplace(line)
}).filter((ch) => {
    if (typeof (ch) === 'number') {
        return ch.parseInt();
    }
});

console.log(res)
