package main

import (
    "fmt"
    "math"
)

func rob(nums []int) int {
    n := len(nums)
    rob := make([]float64, n + 2)
    for i := n-1; i>=0; i-- {
        rob[i] = math.Max( float64(nums[i]) + rob[i+2], rob[i+1])
    }
    return int(rob[0])
}

func main() {
    nums := []int{1,2,3,1}
    res := rob(nums)
    fmt.Println(res)
}
