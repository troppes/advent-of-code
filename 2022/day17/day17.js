import * as fs from 'fs';

let vents;
try {
    vents = fs.readFileSync('input.txt', 'utf8').split('');
} catch (err) {
    console.error(err);
}

let rockLengths = [1,3,3,4,2];

let rocksPos = [
    [[2, 0], [3, 0], [4, 0], [5, 0]],
    [[3, 0], [2, 1], [3, 1], [4, 1], [3, 2]],
    [[4, 0], [4, 1], [4, 2], [3, 2], [2, 2]],
    [[2, 0], [2, 1], [2, 2], [2, 3]],
    [[2, 0], [3, 0], [2, 1], [3, 1]]
];

const simulateRocks = (field, rocks, vents, rockSpawnMax) => {

    let currentRockSpawn = 0;
    let ventIndex = 0;
    let rockIndex = 0;


    while (currentRockSpawn < rockSpawnMax) {

        let rowsFree = field.findIndex(array => array.some(e => e === 1));

        // remove free rows
        field.splice(0, rowsFree);

        // Add 0 Padding of three lines + the length of the piece
        for (let i = 0; i < (3 + rockLengths[rockIndex]); i++) field.unshift([0, 0, 0, 0, 0, 0, 0]);

        // Add Stone
        let currentRock = JSON.parse(JSON.stringify(rocks[rockIndex]));

        // printField(field, currentRock);


        rockIndex++;
        rockIndex %= rocks.length;
        currentRockSpawn++;


        while (1) {

            let possible = true;

            currentRock = moveStone(vents[ventIndex], currentRock, field);

            ventIndex++ 
            ventIndex %= vents.length;

            if (!possible) break;


            [possible, currentRock] = calcGravity(currentRock, field);

            if (!possible) break;

        }



        // draw stones into field
        currentRock.forEach(pos => {
            if (field[pos[1]][pos[0]] === 1) {
                console.log('Logic Error');
            }
            field[pos[1]][pos[0]] = 1;
        })

        // truncate if line is full
        field.splice(field.findIndex(array => array.every(e => e === 1)), field.length);



    }
    // remove free rows
    field.splice(0, field.findIndex(array => array.some(e => e === 1)));
    return field.length - 1;


}

const printField = (field, currentRock) => {
    let fieldCopy = JSON.parse(JSON.stringify(field));
    currentRock.forEach(pos => {
        // logic error
        if (fieldCopy[pos[1]][pos[0]] === 1) {
            fieldCopy[pos[1]][pos[0]] = 3;
        } else {
            fieldCopy[pos[1]][pos[0]] = 2;
        }
    });

    for (let y = 0; y < fieldCopy.length; y++) {
        let row = '';
        for (let x = 0; x < fieldCopy[y].length; x++) {
            let currentChar = fieldCopy[y][x];
            switch (currentChar) {
                case 0:
                    row += '.';
                    break;
                case 1:
                    row += '#';
                    break
                case 2:
                    row += '@';
                    break
                case 3:
                    row += 'X';
                    break
                default:
                    row += 'E'
            }
        }
        console.log(row);
    }
    console.log('\n');


}

const isCollision = (rock, field) => {
    return rock.some(pos => {
        let result = field[pos[1]][pos[0]] !== 0;
        //console.log(pos);
        //console.log(field[pos[1]][pos[0]])
        //console.log(result);
        return result;
    });
}

const calcGravity = (rock, field) => {
    let newRock = rock.map(pos => [pos[0], pos[1] + 1]);

    // if position already full, hit is found
    if (isCollision(newRock, field)) {
        // console.log('Hit ground');
        return [false, rock];
    }
    return [true, newRock];
}

const moveStone = (direction, rock, field) => {
    let newRock;
    if (direction === '<') {
        // console.log('Left');
        newRock = rock.map(pos => [pos[0] - 1, pos[1]]);
    } else {
        // console.log('Right');
        newRock = rock.map(pos => [pos[0] + 1, pos[1]])
    }

    // detect wall
    for (let pos of newRock) {
        if (pos[0] < 0 || pos[0] > 6 || isCollision([pos], field)) {
            // console.log('vent hit something');
            return rock;
        }
    }


    return newRock;
}



const part1 = () => {

    const rockSpawnMax = 2022;
    let field = [[1, 1, 1, 1, 1, 1, 1]];

    const result = simulateRocks(field, rocksPos, vents, rockSpawnMax);

    console.log('Part 1:' + result);

}



part1();