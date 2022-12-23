import * as fs from 'fs';

let data;

try {
    data = fs.readFileSync('input.txt', 'utf8').split('\n');
} catch (err) {
    console.error(err);
}


function checkIfElfHasNeigbours(x, y, field) {
    return [
        field[y - 1][x - 1] + field[y - 1][x] + field[y - 1][x + 1] > 0,
        field[y + 1][x + 1] + field[y + 1][x] + field[y + 1][x - 1] > 0,
        field[y - 1][x - 1] + field[y][x - 1] + field[y + 1][x - 1] > 0,
        field[y - 1][x + 1] + field[y][x + 1] + field[y + 1][x + 1] > 0
    ];
}

const parseInput = (data) => {
    const elfs = [];

    // predict that elves do not move more than 200 steps
    // not the best solution
    const size = 200;
    const field = [];
    for (let i = 0; i < size; i++) {
        field.push(new Array(size).fill(0));
    }

    const currenField = data.map(row => row.split(''));

    const mid = Math.floor((size - currenField.length) / 2);

    for (let y = 0; y < currenField.length; y++) {
        for (let x = 0; x < currenField[y].length; x++) {
            if (currenField[y][x] === '#') {
                field[y + mid][x + mid] = 1;
                elfs.push([x + mid, y + mid]);
            }
        }
    }

    return [elfs, field];
}

function moveElves(field, elfs, rounds) {

    for (let round = 0; round < rounds; round++) {

        const currentFirstDirection = round % 4;
        const newSpacesUsed = new Map();
        
        const newElfs = elfs.map(([x, y]) => {
            const neighbours = checkIfElfHasNeigbours(x, y, field);

            const occupiedSides = neighbours.filter(x => x > 0).length;

            if (occupiedSides === 0 || occupiedSides === 4) {
                // add to freq table
                newSpacesUsed.set([x, y].toString(), (newSpacesUsed.get([x, y].toString()) + 1) || 1)

                // cant move so return original value
                return [x, y];
            } else {
                // potentially can move
                for (let i = currentFirstDirection; i < currentFirstDirection + 4; i++) {
                    if (!neighbours[(i % 4)]) {

                        const moves = [
                            [0, -1],
                            [0, 1],
                            [-1, 0],
                            [1, 0]
                        ];

                        let currentMove = moves[(i % 4)];
                        const newPos = [(x + currentMove[0]), (y + currentMove[1])];

                        if (newPos == null) throw new Error('That is wrong!');
                        newSpacesUsed.set(newPos.toString(), (newSpacesUsed.get(newPos.toString()) + 1) || 1);

                        return newPos;
                    }
                }
            }
        });


        for (let i = 0; i < elfs.length; i++) {

            const [x, y] = elfs[i];
            const [newX, newY] = newElfs[i];

            // console.log('Elf moves from ' + x +' '+ y + ' to ' + newX + ' ' + newY);

            // check if he moves and only one elf wants the new space
            if ((newY != y || newX != x) && newSpacesUsed.get(newElfs[i].toString()) === 1) {

                elfs[i] = [newX, newY];

                field[y][x] = 0;
                field[newY][newX] = 1;
            }

        }
    }
}

function moveElvesPart2(field, elfs) {

    let currentElfCount = -1;
    let rounds = 0;

    while (currentElfCount !== 0) {
        currentElfCount = 0;

        const currentFirstDirection = rounds % 4;
        const newSpacesUsed = new Map();
        
        const newElfs = elfs.map(([x, y]) => {
            const neighbours = checkIfElfHasNeigbours(x, y, field);

            const occupiedSides = neighbours.filter(x => x > 0).length;

            if (occupiedSides === 0 || occupiedSides === 4) {
                // add to freq table
                newSpacesUsed.set([x, y].toString(), (newSpacesUsed.get([x, y].toString()) + 1) || 1)

                // cant move so return original value
                return [x, y];
            } else {
                // potentially can move
                for (let i = currentFirstDirection; i < currentFirstDirection + 4; i++) {
                    if (!neighbours[(i % 4)]) {

                        const moves = [
                            [0, -1],
                            [0, 1],
                            [-1, 0],
                            [1, 0]
                        ];

                        let currentMove = moves[(i % 4)];
                        const newPos = [(x + currentMove[0]), (y + currentMove[1])];

                        if (newPos == null) throw new Error('That is wrong!');
                        newSpacesUsed.set(newPos.toString(), (newSpacesUsed.get(newPos.toString()) + 1) || 1);

                        return newPos;
                    }
                }
            }
        });


        for (let i = 0; i < elfs.length; i++) {

            const [x, y] = elfs[i];
            const [newX, newY] = newElfs[i];

            // console.log('Elf moves from ' + x +' '+ y + ' to ' + newX + ' ' + newY);

            // check if he moves and only one elf wants the new space
            if ((newY != y || newX != x) && newSpacesUsed.get(newElfs[i].toString()) === 1) {

                elfs[i] = [newX, newY];
                currentElfCount++;

                field[y][x] = 0;
                field[newY][newX] = 1;
            }

        }
        
        rounds++;
    }
    return rounds;
}

function countEmpty(elves, field) {
    
    const eRows = elves.map(([y, _]) => y);
    const eCols = elves.map(([_, x]) => x);

    const [minRow, maxRow] = [Math.min(...eRows), Math.max(...eRows)];
    const [minCol, maxCol] = [Math.min(...eCols), Math.max(...eCols)];

    return (maxRow - minRow + 1) * (maxCol - minCol + 1) - field.flat().flat().reduce((partialSum, a) => partialSum + a, 0);
  }

const part1 = (data) => {

    const [elfs, field] = parseInput(data);
    moveElves(field, elfs, 10);

    return countEmpty(elfs, field);
}

const part2 = (data) => {

    const [elfs, field] = parseInput(data);
    return moveElvesPart2(field, elfs);
}



console.log('Part 1: ' + part1(data));
console.log('Part 2: ' + part2(data));