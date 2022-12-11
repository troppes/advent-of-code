import * as fs from 'fs';

let part = 2;

let data;
let monkeys = [];
let supermod = 1;

let rounds = part === 1 ? 20 : 10000;


try {
    data = fs.readFileSync('input.txt', 'utf8').split('\n\n');
} catch (err) {
    console.error(err);
}


for (let line of data) {

    let [name, items, operation, test, positive, negative] = line.split('\n');

    test = test.match(/\d+/g).map(Number)[0];

    supermod *= test;

    const monkey = {
        items: items.match(/\d+/g).map(Number),
        inspectionsDone: 0,
        operation: operation.split('= ')[1],
        test: test,
        positive: positive.match(/\d+/g).map(Number)[0],
        negative: negative.match(/\d+/g).map(Number)[0],
    }

    // if monkeys are missing or out of order / otherwise push.
    monkeys[+name.match(/\d+/g).map(Number)[0]] = monkey;
}

for (let i = 0; i < rounds; i++) {
   for(let monkey of monkeys) {
        while(monkey.items.length > 0) {
            let item = monkey.items.splice(0, 1);

            if(part === 2){
                item %= supermod;
            }

            // Eval is not good, but working
            item = eval(monkey.operation.replace(/old/g, item));
            
            // Enable for Part1
            if(part === 1) {
                item = Math.floor(item / 3);
            } else {

            }

            monkey.inspectionsDone++;

            if (item % monkey.test === 0) {
                monkeys[monkey.positive].items.push(item);
            } else {
                monkeys[monkey.negative].items.push(item);
            }
        }
   } 
}

console.log(monkeys);

monkeys.sort((a, b) => b.inspectionsDone - a.inspectionsDone);

console.log('Solution: ' + monkeys[0].inspectionsDone * monkeys[1].inspectionsDone);