import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	values map[int]int
	arr    []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{make(map[int]int), make([]int, 0, 100)}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	_, exists := this.values[val]
	if exists {
		return false
	}
	this.values[val] = len(this.arr)
	this.arr = append(this.arr, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	currIndex, exists := this.values[val]
	if !exists {
		return false
	}
	lastEle := this.arr[len(this.arr)-1]
	this.arr[currIndex] = this.arr[len(this.arr)-1]
	this.arr = this.arr[:(len(this.arr) - 1)]
	this.values[lastEle] = currIndex
	delete(this.values, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	temp := rand.Intn((len(this.arr) - 1) | 1)
	fmt.Println(temp)
	return this.arr[temp]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */