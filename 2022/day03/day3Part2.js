import * as fs from 'fs';

let data;

try {
    data = fs.readFileSync('input.txt', 'utf8').split('\n');
} catch (err) {
    console.error(err);
}

let score = 0;

for (let i = 0; i < data.length; i += 3) {

    const backpackElfOne = data[i].split('');
    const backpackElfTwo = data[i + 1].split('');
    const backpackElfThree = data[i + 2].split('');

    let intersection = backpackElfOne.filter(element => {
        return backpackElfTwo.includes(element) && backpackElfThree.includes(element);
    })

    // remove duplicates
    intersection = [...new Set(intersection)];

    for (let elem of intersection) {
        if (elem === elem.toUpperCase()) {
            // chatcode 65 to get to zero + 27 to get to value
            score += elem.charCodeAt(0) - 38;
        } else {
            score += elem.charCodeAt(0) - 96;
        }
    }
}

console.log('Score: ' + score);

