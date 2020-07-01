class RandomizedSet {
    indexMap: Map<number, number>;
    arr: number[];
    constructor() {
        this.indexMap = new Map<number, number>();
        this.arr = [];
    }

    insert(val: number): boolean {
        if (this.indexMap.has(val)) {
            return false;
        }
        this.indexMap.set(val, this.arr.length);
        this.arr.push(val)
        return true;
    }

    remove(val: number): boolean {
        if (!this.indexMap.has(val)) {
            return false;
        }
        let currIndex: number = this.indexMap.get(val) || 0;
        this.arr[currIndex] = this.arr[this.arr.length - 1];
        this.indexMap.set(this.arr[currIndex], currIndex);
        this.indexMap.delete(val);
        this.arr.pop();
        return true;
    }

    getRandom(): number {
        // let temp = Math.floor(Math.random()*this.arr.length);
        return this.arr[Math.floor(Math.random() * this.arr.length)];
    }
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * var obj = new RandomizedSet()
 * var param_1 = obj.insert(val)
 * var param_2 = obj.remove(val)
 * var param_3 = obj.getRandom()
 */