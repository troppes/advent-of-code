import * as fs from 'fs';

let blueprints;

let end = null;

try {
    blueprints = fs.readFileSync('input.txt', 'utf8').split('\n').map(elem => {

        let numbers = elem.match(/\d+/g).map(Number)
        // since the materials are always the same we can hardcore here
        return {
            'index': numbers[0],
            'ore': numbers[1],
            'clay': numbers[2],
            'obsidian': { 'ore': numbers[3], 'clay': numbers[4] },
            'geode': { 'ore': numbers[5], 'obsidian': numbers[6] },
        }
    });
} catch (err) {
    console.error(err);
}

const findMaxGeodes = (blueprint, startTime) => {

    //State is amount owned of: ore, clay, obsidian, geodes (cracked), ore robots, clay robots, obsidian robots, geode crackers, and time remaining 
    let seen = new Set();
    let bfsQueue = [];

    let startState = [0, 0, 0, 0, 1, 0, 0, 0, startTime];

    let currentMaxGeodes = 0;

    let maxOreOfBlueprint = Math.max(blueprint.ore, blueprint.clay, blueprint.obsidian.ore, blueprint.geode.ore);

    bfsQueue.push(startState);

    while (bfsQueue.length > 0) {
        let [ore, clay, obsidian, geodes, oreBots, clayBots, obsBots, geoBots, timeLeft] = bfsQueue.pop();

        if (currentMaxGeodes < geodes) {
            currentMaxGeodes = geodes;
        }

        // if time is out we cant produce more
        if (timeLeft == 0) continue;

        // Destroy robots and ressoucres that we cant use
        if (oreBots >= maxOreOfBlueprint) oreBots = maxOreOfBlueprint;
        if (clayBots >= blueprint.obsidian.clay) clayBots = blueprint.obsidian.clay;
        if (obsBots >= blueprint.geode.obsidian) obsBots = blueprint.geode.obsidian;

        let hash = [ore, clay, obsidian, geodes, oreBots, clayBots, obsBots, geoBots, timeLeft].toString();

        // optimize run withs hashes to skip double calculation
        if (seen.has(hash)) {
            continue;
        }
        seen.add(hash);

        timeLeft--;

        let newOre = ore + oreBots;
        let newClay = clay + clayBots;
        let newObsidian = obsidian + obsBots;
        let newGeodes = geodes + geoBots;

        // farm resources
        bfsQueue.push([newOre, newClay, newObsidian, newGeodes, oreBots, clayBots, obsBots, geoBots, timeLeft]);

        // create new orebot
        if (ore >= blueprint.ore) {
            bfsQueue.push([(newOre - blueprint.ore), newClay, newObsidian, newGeodes, (oreBots + 1), clayBots, obsBots, geoBots, timeLeft]);
        }
        // create new claybot
        if (ore >= blueprint.clay) {
            bfsQueue.push([(newOre - blueprint.clay), newClay, newObsidian, newGeodes, oreBots, (clayBots + 1), obsBots, geoBots, timeLeft]);
        }
        // create new obsbot
        if (ore >= blueprint.obsidian.ore && clay >= blueprint.obsidian.clay) {
            bfsQueue.push([(newOre - blueprint.obsidian.ore), (newClay - blueprint.obsidian.clay), newObsidian, newGeodes, oreBots, clayBots, (obsBots + 1), geoBots, timeLeft]);
        }
        // create new geobot
        if (ore >= blueprint.geode.ore && obsidian >= blueprint.geode.obsidian) {
            bfsQueue.push([(newOre - blueprint.geode.ore), newClay, (newObsidian - blueprint.geode.obsidian), newGeodes, oreBots, clayBots, obsBots, (geoBots + 1), timeLeft]);
        }
    }

    return currentMaxGeodes;

}


const calcPart1 = (blueprints) => {
    let result = 0;
    for (let blueprint of blueprints) {
        // array ore, clay, obsidian , geode or the robots => we start with one robot and zero ressources
        result += blueprint.index * findMaxGeodes(blueprint, 24);
    }
    return result;
}

console.log('Part 1: ' + calcPart1(blueprints));