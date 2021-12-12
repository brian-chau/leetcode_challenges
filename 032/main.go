package main

import "fmt"

func longestValidParentheses(s string) int {
    left, right, maxlength := 0,0,0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            left++
        } else {
            right++
        }

        if left == right {
            if 2 * right > maxlength {
                maxlength = 2 * right
            }
        } else if right >= left {
            left, right = 0, 0
        }
    }

    left, right = 0, 0
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '(' {
            left++
        } else {
            right++
        }
        if left == right {
            if 2 * left > maxlength {
                maxlength = 2 * left
            }
        } else if left >= right {
            left, right = 0, 0
        }
    }
    return maxlength
}

func main() {
    input := "(()"
    res := longestValidParentheses(input)
    fmt.Println(res)
}
