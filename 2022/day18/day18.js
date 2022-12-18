import * as fs from 'fs';

let data;
try {
    data = fs.readFileSync('input.txt', 'utf8').split('\n');
} catch (err) {
    console.error(err);
}

let drops = [];

for (let drop of data) {
    drops.push(drop.split(',').map(Number));
}

let voxels = new Map();

drops.forEach(drop => {
    voxels.set(drop.toString(), drop);
})

function getSurface(map) {
    let surface = 0;

    for (let cube of map) {

        // start with 6 sides not touching
        let sides = 6;

        for (let i = -1; i <= 1; i += 2) {
            let x = (cube[1][0] + i) + "," + cube[1][1] + "," + cube[1][2];
            let y = cube[1][0] + "," + (cube[1][1] + i) + "," + cube[1][2];
            let z = cube[1][0] + "," + cube[1][1] + "," + (cube[1][2] + i);

            // if one of the sides touches => remove it
            if(map.has(x)) sides--;
            if(map.has(y)) sides--;
            if(map.has(z)) sides--;
        }

        surface += sides;
    }

    return surface;
}


console.log('Part 1: ' + getSurface(voxels));


