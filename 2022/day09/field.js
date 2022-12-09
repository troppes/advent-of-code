export class Field {


    constructor(knots = 2) {
        this.visited = new Set();
        this.rope = Array.from({length: knots}, (x => {
            return {x: 0, y: 0}
        }));
    }

    handleMovement(direction, distance) {
        for (let i = 0; i < distance; i++) {
            this.moveHead(direction);
            this.calculateTailMovement();
        }
    }

    moveHead(direction) {
        switch (direction) {
            case 'U':
                this.rope[0].y++;
                break;
            case 'D':
                this.rope[0].y--;
                break;
            case 'R':
                this.rope[0].x++;
                break;
            case 'L':
                this.rope[0].x--;
                break;
        }
    }

    calculateTailMovement() {

        for(let i = 1; i < this.rope.length; i++){
            let currentKnot = this.rope[i];
            let leadingKnot = this.rope[i-1];

            if(this.isDistanceIsTooHighBetweenHT(leadingKnot, currentKnot)) {
                currentKnot.x += Math.sign(leadingKnot.x - currentKnot.x);
                currentKnot.y += Math.sign(leadingKnot.y - currentKnot.y);
            }
        }

        this.addFieldVisited(this.rope[this.rope.length-1].x, this.rope[this.rope.length-1].y);

    }

    isDistanceIsTooHighBetweenHT(leadingKnot, currentKnot) {
        return (Math.abs(leadingKnot.x - currentKnot.x) > 1 || Math.abs(leadingKnot.y - currentKnot.y) > 1);
    }

    addFieldVisited(x, y) {
        this.visited.add([x, y].toString());
    }

}