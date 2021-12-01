package main

import (
    "fmt"
    "math"
)

func rob(nums []int) int {
    n := len(nums)
    rob := []int{nums[0]}
    no_rob := []int{0}
    for i := 1; i<n; i++ {
        rob    = append(rob, nums[i] + no_rob[i-1])
        no_rob = append(no_rob, int(math.Max(float64(rob[i-1]), float64(no_rob[i-1]))))
    }
    return int(math.Max(float64(rob[n-1]), float64(no_rob[n-1])))
}

func main() {
    nums := []int{1,2,3,1}
    res := rob(nums)
    fmt.Println(res)
}
