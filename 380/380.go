package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	hm       map[int]int
	valIndex []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{hm: map[int]int{}, valIndex: []int{}}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.hm[val]; ok {
		return false
	}
	this.valIndex = append(this.valIndex, val)
	this.hm[val] = len(this.valIndex) - 1
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if index, ok := this.hm[val]; !ok {
		return false
	} else {

		this.valIndex[index], this.valIndex[len(this.valIndex)-1] = this.valIndex[len(this.valIndex)-1], this.valIndex[index]
		this.hm[this.valIndex[index]] = index
		this.valIndex = this.valIndex[:len(this.valIndex)-1]
		delete(this.hm, val)
		return true
	}
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	if len(this.valIndex) == 0 {
		return -1
	}
	index := rand.Intn(len(this.valIndex))
	return this.valIndex[index]
}
func main() {
	randomSet := Constructor()
	fmt.Println(randomSet.Insert(0))
	fmt.Println(randomSet.Insert(1))
	fmt.Println(randomSet.Remove(0))
	fmt.Println(randomSet.Insert(2))
	fmt.Println(randomSet.Remove(1))
	fmt.Println(randomSet.GetRandom())
}
