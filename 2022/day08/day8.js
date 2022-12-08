import * as fs from 'fs';
import {Field} from './field.js';

let data;

try {
    data = fs.readFileSync('input.txt', 'utf8').split('\n');
} catch (err) {
    console.error(err);
}

const trees = new Field();

for (let line of data) {
    trees.addRow(line.split('').map(Number));

}


console.log(trees.findVisibleTrees());
console.log(trees.findBestSpot());
