package main

import (
    "fmt"
    "linkedlist/list"
)

func main() {

    l := list.List{}

    l.Append(1)
    l.Append(2)
    l.Append(42)
    l.Print()
    l.Prepend(0)
    l.Print()
    fmt.Println("Len:", l.Length())
    fmt.Println("42?", l.Search(42))
    fmt.Println("200?", l.Search(200))
    l.DeleteFirst()
    l.Print()
    fmt.Println("--- deleted first ---")
    l.DeleteLast()
    l.Print()
    l.Append(3)
    l.Append(6)
    l.Append(9)

    l.Print()
    l.Delete(6)
    l.Print()
    l.Delete(1)
    l.Print()
    l.Delete(7)
    l.Print()
}