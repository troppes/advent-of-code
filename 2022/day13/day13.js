import * as fs from 'fs';

let data;
try {
    data = fs.readFileSync('input.txt', 'utf8');
} catch (err) {
    console.error(err);
}

let result1 = 0;
let dataPart1 = data.split('\n\n');
for (let i = 0; i < dataPart1.length; i++) {
    let [arrayOne, arrayTwo] = dataPart1[i].split('\n').map(elem => JSON.parse(elem));

    if (hasCorrectOrder(arrayOne, arrayTwo)) result1 += (i + 1);
}
console.log('Part 1: ' + result1);

let dataPart2 = data.replace(/\n\n/g, "\n").split("\n").map((line) => JSON.parse(line));

// add divider packets
let divider2 = [[2]];
let divider6 = [[6]];
dataPart2 = dataPart2.concat([divider2, divider6]);

dataPart2.sort((a, b) => {
    const aCopy = JSON.parse(JSON.stringify(a));
    const bCopy = JSON.parse(JSON.stringify(b));

    return hasCorrectOrder(aCopy, bCopy) ? -1 : 1;
})

// find positions
const position2 = dataPart2.indexOf(divider2) + 1;
const position6 = dataPart2.indexOf(divider6) + 1;

console.log('Part 2: ' + position2 * position6);



function hasCorrectOrder(array1, array2) {
    while (array1.length > 0 && array2.length > 0) {

        let elem1 = array1.shift();
        let elem2 = array2.shift();

        // both numbers
        if (typeof (elem1) === 'number' && typeof (elem2) === 'number') {
            if (elem1 === elem2) continue;
            return (elem1 < elem2);
        } else {

            // one array one number or both arrays , but do the same
            elem1 = typeof (elem1) === 'object' ? elem1 : [elem1];
            elem2 = typeof (elem2) === 'object' ? elem2 : [elem2];

            const result = hasCorrectOrder(elem1, elem2);

            if (result === undefined) continue;
            return result;
        }
    }

    if (array1.length === array2.length) return undefined;
    return array1.length < array2.length;
}