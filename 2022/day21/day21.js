import * as fs from 'fs';

let monkeys = {};

try {
    fs.readFileSync('input.txt', 'utf8')
        .split('\n')
        .map(line => line.split(': '))
        .forEach(([key, value]) => monkeys[key] = isNaN(+value) ? value.split(" ") : +value);
} catch (err) {
    console.error(err);
}

const calcMonkey = (name, monkeys) => {
    const value = monkeys[name];

    // recurive anker is a monkey with only a number
    if (typeof value == 'number') return value;
    // call it eval art
    return eval(`${calcMonkey(value[0], monkeys)} ${value[1]} ${calcMonkey(value[2], monkeys)}`)
}

const part1 = (monkeys) => {
    return calcMonkey("root", monkeys);
}


console.log('Part 1: ' + part1(monkeys));