package main

import "fmt"

func removeElement(nums []int, val int) int {
    i, j := 0, 0
    for ; j < len(nums); j++ {
        if nums[j] != val {
            nums[i] = nums[j]
            i++
        }
    }
    return i
}

func main() {
    input := []int{3,2,2,3}
    res := removeElement(input, 3)
    fmt.Println(res)
}
