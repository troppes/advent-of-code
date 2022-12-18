import * as fs from 'fs';

// really struggled here, a lot of help came from https://github.com/CodingAP/advent-of-code/blob/main/profiles/github/2022/day15/solution.js

let data;
try {
    data = fs.readFileSync('input.txt', 'utf8');
} catch (err) {
    console.error(err);
}

let beacons = [];
let allBeacons = new Set();

const calcManDistance = (a, b) => Math.abs(a[0] - b[0]) + Math.abs(a[1] - b[1]);

data.split("\n").map(line => {
    let lineData = line.match(/-?\d+/g).map(Number);
    let nearestSensor = [lineData[0], lineData[1]];
    let pos = [lineData[2], lineData[3]];
    const distance = calcManDistance(nearestSensor, pos);

    allBeacons.add(pos.toString());
    beacons.push({ nearestSensor, pos, distance })
});

const part1 = (line) => {
    let notBeacons = new Set();

    for (const beacon of beacons) {

        // find the distance of sensor to line
        let distanceToGoal = Math.abs(beacon.nearestSensor[1] - line);

        // check if the beacon can hit the line we are looking for
        if (distanceToGoal <= beacon.distance) {

            // simulate the beacon 
            let distance = beacon.distance - distanceToGoal;

            for (let i = beacon.nearestSensor[0] - distance; i <= beacon.nearestSensor[0] + distance; i++) {
                if (!allBeacons.has([i, line].toString())) {
                    notBeacons.add([i, line].toString());
                }
            }
        }
    }
    return notBeacons.size;
}

const part2 = (lineMin, lineMax) => {

    for (let line = lineMin; line <= lineMax; line++) {
        let ranges = [];
        for (const beacon of beacons) {

            // find the distance of sensor to line
            let distanceToGoal = Math.abs(beacon.nearestSensor[1] - line);

            // check if the beacon can hit the line we are looking for
            if (distanceToGoal <= beacon.distance) {

                let distance = beacon.distance - distanceToGoal;

                let minX = Math.max(lineMin, beacon.nearestSensor[0] - distance);
                let maxX = Math.min(lineMax, beacon.nearestSensor[0] + distance);

                let currentRange = [minX, maxX];
                    
                for (let i = ranges.length - 1; i >= 0; i--) {

                    // check if current range is not included in ranges
                    if (currentRange[0] <= ranges[i][1] && currentRange[1] >= ranges[i][0] ) {
                        
                        // create a new range that covers the new as well as the old ground
                        currentRange[0] = Math.min(currentRange[0], ranges[i][0]);
                        currentRange[1] = Math.max(currentRange[1], ranges[i][1]);

                        // remove old range
                        ranges.splice(i, 1);
                    }
                }
                // add the new range
                ranges.push(currentRange);
            }

        }

        if (!(ranges[0][0] == lineMin && ranges[0][1] == lineMax)) {
            return (ranges[ranges.length - 1][1] + 1) * 4000000 + line;
        } 
    }


}

console.log('Part 1: ' + part1(2000000));
console.log('Part 2: ' + part2(0, 4000000));