const fs = require('fs');

let data;

try {
    data = fs.readFileSync('input.txt', 'utf8').split('\n');
} catch (err) {
    console.error(err);
}

let score = 0;

for (const line of data) {

    let backpack = line.match(new RegExp('.{1,' + line.length / 2 + '}', 'g'));
    let intersection = backpack[0].split('').filter(element => backpack[1].split('').includes(element));

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

