package main

import "math/rand"

// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]
func main() {
}

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{
		nums: nums,
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	ret := make([]int, len(this.nums))
	copy(ret, this.nums)
	return ret
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int { //模拟洗牌算法
	ret := make([]int, len(this.nums))
	copy(ret, this.nums)
	for i := range ret {
		random := rand.Intn(i + 1) //返回一个[0,i+1)的随机值
		ret[i], ret[random] = ret[random], ret[i]
	}
	return ret
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
