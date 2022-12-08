export class Field {
    constructor() {
        this.field = [];
        this.visibleTrees = [];
        this.maxTreesX = [];
        this.maxTreesY = [];
    }

    addRow(row) {
        this.field.push(row);
    }

    isTreeVisible(x, y) {
        let currentTree = this.field[y][x];

        if (x === 0 || this.field[y].length - 1 === x ||
            y === 0 || this.field.length - 1 === y) return true;

        const [negX, posX, negY, posY] = this.calcMaxTree(x,y);

        if (negX < currentTree || posX < currentTree || negY < currentTree || posY < currentTree) return true;

        return false;
    }

    calcMaxTree(x, y) {
        // create variables for all directions
        let negX = 0, posX = 0, negY = 0, posY = 0;

        for (let i = x - 1; i >= 0; i--) {
            if(this.field[y][i] > negX) negX = this.field[y][i];
        }

        for (let i = x + 1; i < this.field[y].length; i++) {
            if(this.field[y][i] > posX) posX = this.field[y][i];
        }

        for (let j = y - 1; j >= 0; j--) {
            if(this.field[j][x] > negY) negY = this.field[j][x];
        }

        for (let j = y + 1; j < this.field.length; j++) {
            if(this.field[j][x] > posY) posY = this.field[j][x];
        }

        return [negX, posX, negY, posY];
    }

    calcScenicScore(x, y, currentTree) {
        // create variables for all directions
        let scenicNegX = 0, scenicPosX = 0, scenicNegY = 0, scenicPosY = 0;

        for (let i = x - 1; i >= 0; i--) {
            scenicNegX++;
            if(this.field[y][i] >= currentTree) break;
        }

        for (let i = x + 1; i < this.field[y].length; i++) {
            scenicPosX++;
            if(this.field[y][i] >= currentTree) break;
        }

        for (let j = y - 1; j >= 0; j--) {
            scenicNegY++;
            if(this.field[j][x] >= currentTree) break;
        }

        for (let j = y + 1; j < this.field.length; j++) {
            scenicPosY++;
            if(this.field[j][x] >= currentTree) break;
        }

        return scenicNegX * scenicPosX * scenicNegY * scenicPosY;
    }

    findVisibleTrees() {

        let vTrees = 0;
        for (let y = 0; y < this.field.length; y++) {
            for (let x = 0; x < this.field[y].length; x++) {
                let currentTree = this.field[y][x];
                if (this.isTreeVisible(x, y)) {
                    this.visibleTrees.push(currentTree);
                    vTrees++;
                }
            }
        }
        return vTrees;
    }

    findBestSpot() {
        let bestSpot = {x: 0, y: 0, size: 0, scenicScore: 0};
        for (let y = 0; y < this.field.length; y++) {
            for (let x = 0; x < this.field[y].length; x++) {
                let currentTree = this.field[y][x];
                if (this.isTreeVisible(x, y)) {
                    const scenicScore = this.calcScenicScore(x, y, currentTree);
                    if(scenicScore > bestSpot.scenicScore) {
                        bestSpot.x = x;
                        bestSpot.y = y;
                        bestSpot.size = currentTree;
                        bestSpot.scenicScore = scenicScore;
                    }
                }
            }
        }
        return bestSpot;

    }

}