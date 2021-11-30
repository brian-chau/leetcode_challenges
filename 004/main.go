package main

import (
    "fmt"
    "sort"
)

func findMedianSortedArrays(a []int, b []int) float64 {
    c := append(a, b...)
    //sort.Slice(c, func(i, j int) bool { return c[i] < c[j] })
    sort.Ints(c)

    if len(c) % 2 == 1 {
        return float64(c[len(c) / 2])
    }

    middle_number := len(c) / 2
    return float64(c[middle_number - 1] + c[middle_number]) / 2.0
}

func main() {
    a := []int{1,2}
    b := []int{3,4}
    res := findMedianSortedArrays(a,b)
    fmt.Println(res)
}
