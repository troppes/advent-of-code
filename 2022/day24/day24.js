import * as fs from 'fs';

let data;

try {
    data = fs.readFileSync('input.txt', 'utf8').split('\n');
} catch (err) {
    console.error(err);
}

const blizzardDirections = {
    "^": [0, -1],
    "v": [0, 1],
    ">": [1, 0],
    "<": [-1, 0],
}

const parseInput = (data) => {

    const currentField = data.map(row => row.split(''));
    const blizzards = [];

    for (let y = 0; y < currentField.length; y++) {
        for (let x = 0; x < currentField[y].length; x++) {
            if (Object.keys(blizzardDirections).includes(currentField[y][x])) {
                blizzards.push({ x, y, direction: currentField[y][x] });
                currentField[y][x] = '.';
            }
        }
    }

    return [currentField, blizzards];
}

const moveBlizzards = (field, blizzards) => {
    let newBlizzards = [];
    blizzards.forEach(blizzard => {
        let nextStep = blizzardDirections[blizzard.direction];
        let newX = blizzard.x + nextStep[0];
        let newY = blizzard.y + nextStep[1];

        // if newpos wand => reset
        if (field[newY][newX] === '#') {
            switch (blizzard.direction) {
                case '>':
                    newX = 1;
                    break;
                case '<':
                    newX = field[newY].length - 2;
                    break;
                case 'v':
                    newY = 1;
                    break;
                case '^':
                    newY = field.length - 2;
                    break;
            }
        }
        newBlizzards.push({ x: newX, y: newY, direction: blizzard.direction });
    });

    return newBlizzards;

}

const getPossibleMoves = (pos, blizzards, field) => {
    // down / up / left / right / wait
    const noBlizzard = [true, true, true, true, true];
    const moves = [];
    // down
    blizzards.forEach(b => {
        // down up
        if (b.x === pos[0] && b.y === (pos[1] + 1)) noBlizzard[0] = false;
        if (b.x === pos[0] && b.y === (pos[1] - 1)) noBlizzard[1] = false;

        // left right
        if (b.x === (pos[0] - 1) && b.y === pos[1]) noBlizzard[2] = false;
        if (b.x === (pos[0] + 1) && b.y === pos[1]) noBlizzard[3] = false;

        if (b.x === pos[0] && b.y === pos[1]) noBlizzard[4] = false;

    });

    if (noBlizzard[0] &&
        field[(pos[1] + 1)] &&
        field[(pos[1] + 1)][pos[0]] === '.') moves.push([pos[0], (pos[1] + 1)]);

    if (noBlizzard[1] &&
        field[(pos[1] - 1)] &&
        field[(pos[1] - 1)][pos[0]] === '.') moves.push([pos[0], (pos[1] - 1)]);

    if (noBlizzard[2] &&
        field[pos[1]][(pos[0] - 1)] &&
        field[pos[1]][(pos[0] - 1)] === '.') moves.push([(pos[0] - 1), pos[1]]);

    if (noBlizzard[3] &&
        field[pos[1]][(pos[0] + 1)] &&
        field[pos[1]][(pos[0] + 1)] === '.') moves.push([(pos[0] + 1), pos[1]]);

    if (noBlizzard[4]) moves.push([pos[0], pos[1]]);

    return moves;

}

const findBestPath = (field, blizzards, start, goal) => {

    let seen = new Set();
    let bfsQueue = [];

    let startState = { currentTime: 0, x: start[0], y: start[1], currentBlizzards: blizzards };

    let finishedRuns = [];
    let currentShortest = Number.MAX_SAFE_INTEGER;

    bfsQueue.push(startState);

    while (bfsQueue.length > 0) {
        let currentState = bfsQueue.shift();

        // Eliminate Idles
        if (currentState.currentTime > currentShortest) continue;

        // if we hit the goal, write into victory array
        if (currentState.x === goal[0] && currentState.y === goal[1]) {
            console.log('We made it!');
            currentShortest = currentState.currentTime;
            finishedRuns.push(currentState);
            continue;
        }

        let currentTime = currentState.currentTime + 1;

        // move blizzards
        let currentBlizzards = moveBlizzards(field, currentState.currentBlizzards);

        const possibleMoves = getPossibleMoves([currentState.x, currentState.y], currentBlizzards, field);

        for (const possibleMove of possibleMoves) {

            let newState = { x: possibleMove[0], y: possibleMove[1], currentTime, currentBlizzards };

            // if we are currently at the same position with the same time, remove it
            let hash = `t:${currentTime} pos:${newState.x},${newState.y}`;

            // optimize run withs hashes to skip double calculation
            if (seen.has(hash)) {
                continue;
            }
            seen.add(hash);

            bfsQueue.push(newState);
        }
    }
    return finishedRuns.sort((a, b) => a.currentTime > b.currentTime)[0];

}

const part1 = (data) => {

    const [field, blizzards] = parseInput(data);
    const goal = [(field[(field.length - 1)].length - 2), (field.length - 1)];

    return findBestPath(field, blizzards, [0,1], goal).currentTime;
}

const part2 = (data) => {

    const [field, blizzards] = parseInput(data);
    const goal = [(field[(field.length - 1)].length - 2), (field.length - 1)];

    let r1 = findBestPath(field, blizzards, [1,0], goal);
    let r2 = findBestPath(field, r1.currentBlizzards, goal, [1, 0]);
    let r3 = findBestPath(field, r2.currentBlizzards, [1, 0], goal);

    return r1.currentTime + r2.currentTime + r3.currentTime;
}

console.log('Part 1: ' + part1(data));
console.log('Part 2: ' + part2(data));

