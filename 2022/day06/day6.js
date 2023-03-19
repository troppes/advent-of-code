import * as fs from 'fs';

let data;

try {
    data = fs.readFileSync('input.txt', 'utf8').split('');
} catch (err) {
    console.error(err);
}

let messageFound = false;
for (let i = 0; i < data.length; i++) {

    let set = new Set(data.slice(i, i + 4));
    let messageSet = new Set(data.slice(i, i + 14));

    if(set.size === 4 && !messageFound) {
        // add the length of the marker
        console.log('Found the start marker of 4 characters at: ' + (i + 4));
        messageFound = true;
    }

    if(messageSet.size === 14) {
        console.log('Found the message marker of 14 characters at: ' + (i + 14));
        return;
    }
}


