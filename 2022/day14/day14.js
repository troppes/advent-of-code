import * as fs from "fs";

let data;
try {
  data = fs
    .readFileSync("input.txt", "utf8")
    .split("\n")
    .map((row) => row.split(" -> ").map((r) => r.split(",").map(Number)));
} catch (err) {
  console.error(err);
}

const calcMinMax = (data) => {
  let x = [];
  let y = [];

  for (let row of data) {
    for (let tupel of row) {
      x.push(tupel[0]);
      y.push(tupel[1]);
    }
  }

  console.log(Math.max(...y));

  return [Math.min(...x), Math.max(...x), Math.max(...y)];
};

const buildField = (data) => {
  let field = {};

  const [minX, maxX, maxY] = calcMinMax(data);

  console.log(minX);
  console.log(maxX);
  console.log(maxY);
  // create field with right indices
  for (let y = 0; y <= maxY + 1; y++) {
    for (let x = minX - 1; x <= maxX + 1; x++) {
      field[[y, x].toString()] = 0;
    }
  }

  // build paths
  data.forEach((rock) => {
    for (let i = 0; i < rock.length - 1; i++) {
      let one = rock[i];
      let two = rock[i + 1];

      if (one[0] === two[0]) {
        // y changes
        let ordered = [one[1], two[1]].sort();
        for (let j = ordered[0]; j <= ordered[1]; j++)
          field[[j, one[0]].toString()] = 1;
      } else if (one[1] === two[1]) {
        // x changes
        let ordered = [one[0], two[0]].sort();
        for (let j = ordered[0]; j <= ordered[1]; j++)
          field[[one[1], j].toString()] = 1;
      } else {
        console.log("Logic Error");
      }
    }
  });

  return field;
};

const simulateSand = (field) => {
  let sandSpawn = [0, 500];
  let sandCounter = 0;

  const [minX, maxX, maxY] = calcMinMax(data);

  while (1) {
    let sand = JSON.parse(JSON.stringify(sandSpawn));
    sandCounter++;

    while (field[[sand[0], sand[1]].toString()] !== undefined) {
      if (field[[sand[0] + 1, sand[1]].toString()] === 0) {
        sand[0]++;
      } else if (field[[sand[0] + 1, sand[1] - 1].toString()] === 0) {
        sand[0]++;
        sand[1]--;
      } else if (field[[sand[0] + 1, sand[1] + 1].toString()] === 0) {
        sand[0]++;
        sand[1]++;
      } else {
        field[sand.toString()] = 2;
        break;
      }
    }

    // write down

    // negate the +1 added in the beggining
    if (sand[0] > maxY) return sandCounter - 1;
    console.log("Turn" + sandCounter);

    //if (sandCounter > 5) break;
    // printField(field);
  }
};

const printField = (field) => {
  let minX = Number.MAX_SAFE_INTEGER;
  let minY = Number.MAX_SAFE_INTEGER;
  let maxX = 0;
  let maxY = 0;
  for (let index in field) {
    // create index

    let [indexY, indexX] = index.split(",").map(Number);
    if (indexX < minX) minX = indexX;
    if (indexX > maxX) maxX = indexX;

    if (indexY < minY) minY = indexY;
    if (indexY > maxY) maxY = indexY;
  }

  for (let y = minY; y <= maxY; y++) {
    let row = "";
    for (let x = minX; x <= maxX; x++) {
      let currentChar = field[[y, x].toString()];
      switch (currentChar) {
        case 0:
          row += ".";
          break;
        case 1:
          row += "#";
          break;
        case 2:
          row += "o";
          break;
      }
    }
    console.log(row);
  }
  console.log("\n");
};
const part1 = () => {
  let field = buildField(data);
  printField(field);
  console.log("Part 1:" + simulateSand(field));
};

part1();
