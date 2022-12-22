import * as fs from 'fs';

let data;

try {
    data = fs.readFileSync('input.txt', 'utf8').split('\n\n');
} catch (err) {
    console.error(err);
}

const doMaze = (maze, instructions) => {

    const directions = ['right', 'down', 'left', 'up'];

    let y = 0;
    // find first tile that is not empty
    let x = maze[y].findIndex(tile => tile !== ' ');
    let direction = 0;

    for (let step of instructions) {
        console.log('---' + step + '---');
        if (isNaN(step)) {
            // Turn
            if (step === 'L') {
                direction--;
                if (direction < 0) direction = 3;
            } else {
                direction++;
                if (direction > 3) direction = 0;
            }
        } else {
            // Move
            for (let i = 0; i < step; i++) {

                let newX = x;
                let newY = y;

                let dir = directions[direction];

                switch (dir) {
                    case 'right':
                        newX += 1;
                        break;
                    case 'left':
                        newX -= 1;
                        break;
                    case 'up':
                        newY -= 1;
                        break;
                    case 'down':
                        newY += 1;
                        break;
                }

                let nextStep = checkInBoud(maze, newX, newY);

                // hit wall, so read next instruction
                if (nextStep === '#') {
                    break;
                } else if (nextStep === ' ') {

                    switch (dir) {
                        case 'right':
                            newX = maze[y].findIndex(tile => tile !== ' ');
                            nextStep = maze[y][newX];
                            break;
                        case 'left':
                            // findLastIndex is not working in Node for some reason
                            // reverse is inplace so a new array is needed
                            newX = (maze[y].length - 1) - [...maze[y]].reverse().findIndex(tile => tile !== ' ');
                            nextStep = maze[y][newX];
                            break;
                        case 'up':
                            for (let j = maze.length - 1; j >= 0; j--) {
                                if (maze[j][x] !== ' ' && maze[j][x] !== undefined) {
                                    nextStep = maze[j][x];
                                    newY = j;
                                    break;
                                }
                            }
                            break;
                        case 'down':
                            for (let j = 0; j < maze.length; j++) {
                                if (maze[j][x] !== ' ' && maze[j][x] !== undefined) {
                                    nextStep = maze[j][x];
                                    newY = j;
                                    break;
                                }
                            }
                            break;
                    }
                }

                if (nextStep === '.') {
                    x = newX;
                    y = newY;
                }
            }
        }
    }
    return 1000 * (y + 1) + 4 * (x + 1) + direction;
}

function checkInBoud(maze, x, y) {
    if (y >= 0 && y < maze.length && x >= 0 && x < maze[y].length) return maze[y][x];
    return ' ';
}

const parseInput = (data) => {

    const instructions = data[1].match(/(\d+)|([A-Z]+)/gi).map(step => isNaN(+step) ? step : +step);
    const maze = data[0].split('\n').map(row => row.split(''));

    return [maze, instructions];
}


const part1 = (data) => {
    let [maze, instructions] = parseInput(data);

    return doMaze(maze, instructions);
}

console.log('Part 1: ' + part1(data));