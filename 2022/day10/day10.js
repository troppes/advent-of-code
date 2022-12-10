import * as fs from 'fs';

let data;
let cycles = 0;
let currentRegister = 1;
let crt = [];
let crtSize = 40;
let queue = [];
let currentNumberToAwait = 0;
let result = 0;
let currentCrtPosition = 0;

try {
    data = fs.readFileSync('input.txt', 'utf8').split('\n');
} catch (err) {
    console.error(err);
}


for (let line of data) {

    const [command, numberToAdd] = line.split(' ');
    queue.push([command, +numberToAdd]);
}

for (let i = 0; i < 8; i++) {
    crt.push(new Array(crtSize).fill('.'));
}

while (queue.length > 0) {
    cycles++;

    if (currentNumberToAwait === null) {
        const [command, numberToAdd] = queue.shift();

        if (command === 'addx') {
            currentNumberToAwait = numberToAdd;
        }

    } else {
        currentRegister += currentNumberToAwait;
        currentNumberToAwait = null;
    }

    if (cycles === 20 || (cycles - 20) % 40 === 0) {
        result += (cycles * currentRegister);
        // console.log('Cycle: ' + cycles);
        // console.log('current register:' + currentRegister);
        // console.log((cycles * currentRegister));

    }

    // handle the crt in here
    let row = Math.floor((cycles + 1) / crtSize);

    if (Math.abs(currentCrtPosition - currentRegister) < 2) {
        crt[row][currentCrtPosition] = '#';
    }

    if (currentCrtPosition < crtSize - 1) {
        currentCrtPosition += 1;
    } else {
        currentCrtPosition = 0;
    }

}

console.log('Part 1:' + result);

console.log('Part 2:')
for (let i = 0; i < 8; i++) {
    let row = '';
    for (let j = 0; j < 40; j++) {
        row += crt[i][j];
    }
    console.log(row);
}
