// Part 1 understood with the code of: https://github.com/surgi1/adventofcode/blob/main/2022/day16/script.js
// Part 2 solo, but unoptimised

import * as fs from 'fs';

let data;
try {
    data = fs.readFileSync('input.txt', 'utf8').split('\n');
} catch (err) {
    console.error(err);
}

let nodes = []
for (let node of data) {
    let tmp = node.split(' ');

    nodes[tmp[1]] = {
        rate: Number(tmp[4].match(/\d+/g)[0]),
        connections: tmp.slice(tmp.indexOf('to') + 2).map(valves => valves.substr(0, 2)),
        distanceMap: null,
    }

}

const createNewNodes = (exceptions) => {
    let newNodes = [];
    for (let node of data) {
        let tmp = node.split(' ');

        let rate = exceptions.includes(tmp[1]) ? 0 : Number(tmp[4].match(/\d+/g)[0]);
        newNodes[tmp[1]] = {
            rate,
            connections: tmp.slice(tmp.indexOf('to') + 2).map(valves => valves.substr(0, 2)),
            distanceMap: null,
        }
    
    }

    return newNodes;
}



// filter out all inactive nodes
const activeNodes = (nodes) => {
    let actives = [];
    for(let index in nodes){
        if(nodes[index].rate > 0) {
            actives.push(index);
        }
    }
    return actives;
}

const distanceMap = (startName, distances = {}) => {
    if (nodes[startName].distanceMap) return nodes[startName].distanceMap;

    const spread = (name, steps) => {
        if (distances[name] != undefined && distances[name] <= steps) return;
        distances[name] = steps;
        nodes[name].connections.forEach(n => spread(n, steps + 1));
    }

    spread(startName, 0);
    nodes[startName].distanceMap = distances;
    return distances;
}

const computePaths = (timeLeft, nodes) => {

    let paths = [{ 
        curr: 'AA', 
        active: activeNodes(nodes),
        timeLeft: timeLeft, 
        finished: false, 
        steps: [], 
        releasedPressure: 0,
    }]

    for (let n = 0; n < paths.length; n++) {

        let path = paths[n];


        if (path.timeLeft <= 0) {
            path.finished = true;
        }

        if(path.finished) continue;

        let distances = distanceMap(path.curr);
        let moved = false;

        path.active.forEach(activePath => {
            // its already visted
            if (activePath == path.curr) return true;
            // we cant visit the path anytmore
            if (path.timeLeft - distances[activePath] <= 1) return true;

            moved = true;

            paths.push({
                curr: activePath,
                active: path.active.filter(v => v != activePath),
                timeLeft: path.timeLeft - distances[activePath] - 1,
                finished: false,
                steps: [...path.steps, activePath],
                releasedPressure: path.releasedPressure + (path.timeLeft - distances[activePath] - 1) * nodes[activePath].rate
            })
        })

        if (!moved) path.finished = true;
    }

    return paths.filter(p => p.finished).sort((a, b) => b.releasedPressure - a.releasedPressure);
}





const part2 = () => {
    let paths = computePaths(26, nodes);
    let maxScore = 0;

    for(let path of paths) {
        let newNodes = createNewNodes(path.steps);
        let elePath = computePaths(26, newNodes)[0];

        let currentScore = elePath.releasedPressure + path.releasedPressure;

        if(currentScore > maxScore) maxScore = currentScore;
    }

    console.log('Part2: ' + maxScore);
}

console.log('Part1: ' + computePaths(30, nodes)[0].releasedPressure);

part2();