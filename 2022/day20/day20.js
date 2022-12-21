import * as fs from 'fs';

let input;
try {
    input = fs.readFileSync('input.txt', 'utf8').split('\n').map(Number);
} catch (err) {
    console.error(err);
}


const part1 = (input) => {

    let original = [];

    input.forEach(element => { original.push([original.length, element]); });

    let list = JSON.parse(JSON.stringify(original));

    for (let [index, value] of original) {
        let pos = list.findIndex(elem => elem[0] === index && elem[1] === value);

        let newPos = ((pos + value) % (list.length - 1));
        if (newPos < 0) newPos += (list.length - 1);

        let currentItem = list.splice(pos, 1)[0];
        list.splice(newPos, 0, currentItem);
    }

    // start with zero
    let zeroPos = list.findIndex(elem => elem[1] === 0);
    return list[(zeroPos + 1000) % list.length][1] + list[(zeroPos + 2000) % list.length][1] + list[(zeroPos + 3000) % list.length][1];

}

const part2 = (input) => {

    let original = [];

    input.forEach(element => { original.push([original.length, (element * 811589153)]); });

    let list = JSON.parse(JSON.stringify(original));

    for (let i = 0; i < 10; i++) {
        for (let [index, value] of original) {
            let pos = list.findIndex(elem => elem[0] === index && elem[1] === value);

            let newPos = ((pos + value) % (list.length - 1));
            if (newPos < 0) newPos += (list.length - 1);

            let currentItem = list.splice(pos, 1)[0];
            list.splice(newPos, 0, currentItem);
        }
    }

    // start with zero
    let zeroPos = list.findIndex(elem => elem[1] === 0);
    return list[(zeroPos + 1000) % list.length][1] + list[(zeroPos + 2000) % list.length][1] + list[(zeroPos + 3000) % list.length][1];
}

console.log('Part 1: ' + part1(input));
console.log('Part 1: ' + part2(input));