package main

import (
    "fmt"
)

func twoSum(nums []int, target int) []int {
    some_map := make(map[int]int)
    var res []int

    if len(nums) < 2 {
        return res
    }

    for i := 0; i<len(nums); i++ {
        j := target - nums[i]
        if _, ok := some_map[j]; ok {
            res = append(res, some_map[j])
            res = append(res, i)
            return res
        } else {
            some_map[nums[i]] = i
        }
    }
    return res
}

func main() {
    input := []int{2,7,11,15}
    res := twoSum(input, 9)
    fmt.Println(res)
}
