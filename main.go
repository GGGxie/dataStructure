package main

import "fmt"

type DHResp struct {
	IsDivide  bool        `json:"is_divide"` //是否需要拆单
	IsMerge   bool        `json:"is_merge"`  //是否需要合并单据
	Key       string      `json:"key"`
	MapFields string      `json:"map_fields"`
	Auth      string      `json:"auth,omitempty"`
	A         interface{} `json:"a,omitempty"`
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func main() {
	a := []int{1, 2, 3, 3, 3, 4}
	i := removeDuplicates(a)
	fmt.Println(a, i)
}
