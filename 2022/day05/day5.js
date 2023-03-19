import * as fs from 'fs';
const ESACPE_LETTER = 'X';

let data;

try {
    data = fs.readFileSync('input.txt', 'utf8').split('\n');
} catch (err) {
    console.error(err);
}

const stacks = [];


while (data.length > 0) {
    // replace 4 whitespace with X, then remove brackets and lastly remove whitespaces
    let cleaned = data[0].replace(/\s{4}/g, ESACPE_LETTER).replace(/[\[\]']+/g, '').replace(/\s/g, '');

    data.splice(0, 1); // remove item
    if (cleaned.startsWith('1')) {
        // remove blank line
        data.splice(0, 1);
        break;
    }

    for (let i = 0; i < cleaned.length; i++) {
        if (cleaned[i] !== ESACPE_LETTER) {
            if (Array.isArray(stacks[i])) {
                stacks[i].push(cleaned[i]);
            } else {
                stacks[i] = new Array(cleaned[i]);
            }
        }
    }
}


for (const line of data) {
    let [qty, from, to] = line.match(/\d+/g).map(Number);

    // Remove reverse() to finish part 2
    stacks[to - 1].unshift(...stacks[from - 1].splice(0, qty).reverse());
}


let answer = '';

for (let stack of stacks) {
    answer += stack[0];
}

console.log(answer);


