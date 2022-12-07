import * as fs from 'fs';
import {FSTree} from './filesystem.js';

let data;

try {
    data = fs.readFileSync('input.txt', 'utf8').split('\n');
} catch (err) {
    console.error(err);
}

const firstLine = data.splice(0, 1)[0].split(' ');
const fileSystem = new FSTree(firstLine[2], 'folder');

for (let i = 0; i < data.length; i++) {
    let lineInformation = data[i].split(' ');

    if (lineInformation[0] === '$') {
        switch (lineInformation[1]) {
            case 'cd':
                if (lineInformation[2] === '..') {
                    fileSystem.moveOut();
                } else {
                    fileSystem.insertNewFolder(lineInformation[2]);
                }
                break;
            case 'ls':
                while ((i + 1) < data.length && !(data[i + 1].startsWith('$'))) {
                    if (!data[i + 1].startsWith('dir')) {
                        let file = data[i + 1].split(' ');
                        fileSystem.insertNewFile(file[0]);
                    }
                    i++;
                }
                break;
        }
    }
}

fileSystem.calcSize();
console.log('Calc Part 1: ' + fileSystem.part1);
// 70000000 total - 30000000 needed = 40000000 be full // substract from root to find how much is necesary to delete
fileSystem.calcSmallestFolderToDeleteForUpdate(fileSystem.root, (fileSystem.root.size - 40000000))
console.log('Calc Part 2: ' + fileSystem.minSizeFolder);


