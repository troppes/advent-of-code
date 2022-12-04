const fs = require('fs');

let data;

try {
    data = fs.readFileSync('input.txt', 'utf8').split('\n');
} catch (err) {
    console.error(err);
}

let overlaps = 0;
let partialOverlaps = 0;

for (const line of data) {
    let assignments = line.split(',');

    assignments = assignments.map(e => e.split('-').map(Number));

    if (assignments[0][1] >= assignments[1][1] && assignments[0][0] <= assignments[1][0] ||
        assignments[0][1] <= assignments[1][1] && assignments[0][0] >= assignments[1][0]
    ) {
        overlaps++;
    }

    if (assignments[0][0] <= assignments[1][1] && (assignments[1][0] <= assignments[0][1])) partialOverlaps++;

}
console.log('Overlaps: ' + overlaps);
console.log('Partial Overlaps: ' + partialOverlaps);

