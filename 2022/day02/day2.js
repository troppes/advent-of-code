import * as fs from 'fs';

let data;

try {
    data = fs.readFileSync('input.txt', 'utf8').split('\n');
} catch (err) {
    console.error(err);
}

// A || X = Rock   B || Y = Paper   C || Z = Scissors
const getScore = (a, b) => {
    // normalize Input
    // -23 to convert X to A and -64 to create A=1 B=2 C=3
    a = a.charCodeAt(0) - 64;
    b = b.charCodeAt(0) - 23 - 64;


    if (a === (b - 1) || ((b - 1) === 0 && a === 3)) {
        return b + 6;
    } else if (a === b) {
        return b + 3;
    } else {
        return b;
    }
}

//X = Loose   Y = Draw   Z = Win
const getScoreStarTwo = (a, b) => {
    // normalize Input
    // -23 to convert X to A and -64 to create A=1 B=2 C=3
    a = a.charCodeAt(0) - 64;

    switch (b) {
        case 'X':
            if ((a - 1) === 0) return 3;
            return a - 1;
        case 'Y':
            return a + 3;
        case 'Z':
            if ((a + 1) === 4) return 7;
            return (a + 1) + 6;
        default:
            console.error('INVALID INPUT!');
    }
}

let currentScore = 0;
let currentScoreStarTwo = 0;

for (const line of data) {
    let game = line.split(' ');

    currentScore += getScore(game[0], game[1]);
    currentScoreStarTwo += getScoreStarTwo(game[0], game[1]);
}

console.log('Part 1:' + currentScore);
console.log('Part 2:' + currentScoreStarTwo);

