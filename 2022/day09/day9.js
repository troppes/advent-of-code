import * as fs from 'fs';
import {Field} from './field.js';

let data;

try {
    data = fs.readFileSync('input.txt', 'utf8').split('\n');
} catch (err) {
    console.error(err);
}

const part1 = new Field(2);
const part2 = new Field(10);

for (let line of data) {
    const [direction, distance] = line.split(' ');
    part1.handleMovement(direction, +distance);
    part2.handleMovement(direction, +distance);
}

console.log('Part 2 : ' + part1.visited.size);
console.log('Part 2 : ' + part2.visited.size);
