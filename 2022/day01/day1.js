import * as fs from 'fs';

let data;

try {
    data = fs.readFileSync('input.txt', 'utf8').split('\n');
} catch (err) {
    console.error(err);
}

let currentCalories = 0;

const elfCalories = [];

for (const line of data) {
    if (line === '') {
        elfCalories.push(currentCalories);
        currentCalories = 0;
    }
    currentCalories += Number(line);
}

elfCalories.sort((a, b) => a-b).reverse();

console.log('Highest Calories: ' + elfCalories[0]);
console.log('Top 3 Calories: ' + (elfCalories[0] + elfCalories[1] + elfCalories[2]));