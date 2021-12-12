package main

import (
    "fmt"
    "math"
    "math/big"
)

func nthMagicalNumber(n int, a int, b int) int {
    var z big.Int
    x := big.NewInt(int64(a))
    y := big.NewInt(int64(b))
    lcm := z.Mul(z.Div(x, z.GCD(nil, nil, x, y)), y)
    L := lcm.Int64()
    l := int(L)
    C := (l/a) + (l/b) - 1
    k := (n-1)/C
    ans := 0
    for i, j, magical_needed := 1, 1, n-k*C; magical_needed != 0; magical_needed-- {
        if (a*i < b*j) {
            ans = a * i
            i++
        } else {
            ans = b * j
            j++
        }
    }
    return (k*l + ans) % int(math.Pow(10,9)+7.0)
}

func main () {
    res := nthMagicalNumber(4,2,3)
    fmt.Println(res)
}
