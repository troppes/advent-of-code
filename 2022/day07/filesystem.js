export class Folder {
    constructor(name, parent = null) {
        this.name = name;
        this.size = 0;
        this.parent = parent;
        this.children = [];
    }
}

export class FSTree {

    currentPointer = null;
    root = null;

    part1 = 0;
    minSizeFolder = Number.MAX_VALUE;

    constructor(key, value = key) {
        this.root = new Folder(key, value);
        this.currentPointer = this.root;
    }

    calcSize(folder = this.root) {
        for (const child of folder.children) {
            folder.size += this.calcSize(child);
        }
        if (folder.size < 100000) this.part1 += folder.size;

        return folder.size;
    }

    calcSmallestFolderToDeleteForUpdate(folder = this.root, spaceNeeded) {
        // ignore folders that are too small
        if (folder.size >= spaceNeeded) {
            // if folder is found, check if the current directory is smaller, than the one found before
            if (folder.size < this.minSizeFolder) {
                this.minSizeFolder = folder.size;
            }
        }

        for (const child of folder.children) {
            this.calcSmallestFolderToDeleteForUpdate(child, spaceNeeded);
        }
    }

    insertNewFile(filesize) {
        this.currentPointer.size += +filesize;
    }

    insertNewFolder(name) {
        const newNode = new Folder(name, this.currentPointer);

        this.currentPointer.children.push(newNode);
        this.currentPointer = newNode;
    }

    moveOut() {
        this.currentPointer = this.currentPointer.parent;
    }
}