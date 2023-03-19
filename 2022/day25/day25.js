import * as fs from 'fs';

let data;

try {
    data = fs.readFileSync('input.txt', 'utf8').split('\n');
} catch (err) {
    console.error(err);
}
const snafuToBase10 = (number) => {

    let result = 0;

    // reverse for easier index
    number = number.split('').reverse();
    for (let i = 0; i < number.length; i++) {
        let currentChar = number[i];

        if (currentChar === '-') currentChar = -1;
        if (currentChar === '=') currentChar = -2;

        let currentNumber = parseInt(currentChar) * Math.pow(5, i);
        result += currentNumber;
    }

    return result;
}

const base10ToSnafu = (number) => {

    let result = "";

    do {
        const remains = number % 5;

        if(remains === 3) {
            result = '=' + result;
            number += 5;
        } else if(remains === 4) {
            result = '-' + result;
            number += 5;
        } else {
            result = remains + result;
        }


        number = Math.floor(number/5);
        
    } while (number !== 0);

    return result;
}

const part1 = (data) => {

    let numberBase10 = 0;

    for (let number of data) {
        numberBase10 += snafuToBase10(number);
    }
    console.log(numberBase10);
    return base10ToSnafu(numberBase10);
}

console.log('Part 1: ' + part1(data));

