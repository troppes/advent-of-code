import * as fs from 'fs';
import { maxHeaderSize } from 'http';

let data;
let startSymbol = 'S';
let endSymbol = 'E';
let start = null;
let end = null;

try {
    data = fs.readFileSync('input.txt', 'utf8').split('\n').map((row, y) => row.split('').map((c, x) => {
        if (c == startSymbol) {
            start = [x, y];
            c = 'a';
        } else if (c == endSymbol) {
            end = [x, y];
            c = 'z';
        }
        return c.charCodeAt(0) - 'a'.charCodeAt(0);
    }));
} catch (err) {
    console.error(err);
}





const getShortestPath = (start) => {
    let finalPath = [];
    let queue = [[start]];
    let visited = [start.toString()];

    while (queue.length > 0 && finalPath.length == 0) {
        const path = queue.shift();
        const [x, y] = path[path.length - 1];

        let directions = [
            [x + 1, y],
            [x - 1, y],
            [x, y + 1],
            [x, y - 1]
        ];

        for (const direction of directions) {

            // check oob
            if (direction[0] < 0 || direction[0] >= data[0].length || direction[1] < 0 || direction[1] >= data.length) continue;
            // check if already visited
            if (visited.includes(direction.toString())) continue;
            // check if distance too big
            if (data[direction[1]][direction[0]] - data[y][x] > 1) continue;

            if (direction.toString() === end.toString()) {
                return path.concat([end]).length - 1;
            }

            visited.push(direction.toString());
            queue.push(path.concat([direction]));
        }
    }
    // no path available
    return Number.MAX_SAFE_INTEGER;
}

console.log('Part 1: ' + getShortestPath(start));

let shortest = Number.MAX_SAFE_INTEGER;
for (let y = 0; y < data.length; y++) {
    for (let x = 0; x < data[y].length; x++) {
        if (data[y][x] === 0) {
            let pathLength = getShortestPath([x, y])
            if (shortest > pathLength) shortest = pathLength;
        }
    }
}

console.log('Part 2: ' + shortest);