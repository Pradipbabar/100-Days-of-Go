package list

import "fmt"

type Node struct {
    Data int
    Next *Node
}

type List struct {
    Head *Node
}


func (l *List) Append(data int) {

    node := &Node{ Data: data, Next: nil }

    if l.Head == nil {
        l.Head = node
    } else {
        p := l.Head
        for p.Next != nil {
            p = p.Next
        }

        p.Next = node
    }
}


func (l List) Print() {
    p := l.Head
    fmt.Print("head")
    for p != nil {
        fmt.Print("->", p.Data)
        p = p.Next
    }
    fmt.Println()
}

func (l *List) Prepend(data int) {

    node := &Node{Data: data, Next: nil}

    if l.Head == nil {
        l.Head = node
    } else {
        temp := l.Head
        l.Head = node
        node.Next = temp
    }
}

func (l List) Length() int {
    p := l.Head
    var len int
    for p != nil {
        len++
        p = p.Next
    }
    return len
}

func (l List) Search(data int) bool {
 
    p := l.Head
    for p != nil {
        if p.Data == data {
            return true
        }
        p = p.Next
    }

    return false
}

func (l *List) DeleteFirst() {

    if l.Head == nil {
        return
    }
    p := l.Head

    l.Head = p.Next

    p = nil

}

func (l *List) DeleteLast() {

    if l.Head == nil {
        return
    }

    p := l.Head
    cur := &Node{}

    for p.Next != nil {
        cur = p
        p = p.Next
    }
    cur.Next = nil

}

func (l *List) Delete(data int) {

    if l.Head == nil {
        return
    }

    curr := l.Head
    prev := l.Head

    if l.Head.Data == data {
        l.Head = l.Head.Next
        return
    }

    for curr != nil {

        if curr.Data == data {
            break
        }
        prev = curr
        curr = curr.Next
    }
    
    if curr == nil {
        return
    }

    prev.Next = curr.Next
}
