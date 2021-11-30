package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

type List struct {
    head *ListNode
    tail *ListNode
}

func merge( a *ListNode, b *ListNode ) {
    if a.Next == nil {
        a.Next = b
        return
    }

    if b == nil {
        return
    }

    if a.Next.Val <= b.Val {
        merge(a.Next, b)
    } else {
        tmp := a.Next
        a.Next = b
        merge(a.Next, tmp)
    }
}

func mergeTwoLists( a *ListNode, b *ListNode ) *ListNode {
    if a == nil {
        return b
    }
    if b == nil {
        return a
    }

    if a.Val > b.Val {
        merge(b,a)
        return b
    }
    merge(a,b)
    return a
}

func (L *List) Insert( val int ) {
    node := &ListNode{ Val: val, Next: nil }

    if L.head == nil {
        L.head = node
    } else {
        l := L.head
        for l.Next != nil {
            l = l.Next
        }

        l.Next = node
        L.tail = l
    }
}

func (l *List) Display () {
    list := l.head
    for list != nil {
        fmt.Printf("%+v ", list.Val)
        list = list.Next
    }
    fmt.Println()
}

func main() {
    a := List{ }
    a.Insert( 1 )
    a.Insert( 2 )
    a.Insert( 4 )

    b := List{ }
    b.Insert( 1 )
    b.Insert( 4 )
    b.Insert( 3 )

    // a.Display()
    c := mergeTwoLists(a.head, b.head)
    for c != nil {
        fmt.Printf("%d", c.Val)
        c = c.Next
    }
    fmt.Println()
}

