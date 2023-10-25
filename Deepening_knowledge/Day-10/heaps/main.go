package main

import (
   "fmt"
)
type BinaryHeap struct {
   array []int
   size  int
}
func NewBinaryHeap() *BinaryHeap {
   return &BinaryHeap{}
}
func (h *BinaryHeap) Insert(value int) {
   h.array = append(h.array, value)
   h.size++
   h.percolateUp(h.size - 1)
}
func (h *BinaryHeap) Delete() int {
   if h.size == 0 {
      panic("Heap is empty")
   }

   min := h.array[0]
   h.array[0] = h.array[h.size-1]
   h.array = h.array[:h.size-1]
   h.size--
   h.percolateDown(0)

   return min
}

func (h *BinaryHeap) percolateUp(index int) {
   parentIndex := (index - 1) / 2

   for index > 0 && h.array[index] < h.array[parentIndex] {
      h.array[index], h.array[parentIndex] = h.array[parentIndex], h.array[index]
      index = parentIndex
      parentIndex = (index - 1) / 2
   }
}

func (h *BinaryHeap) percolateDown(index int) {
   for {
      leftChildIndex := 2*index + 1
      rightChildIndex := 2*index + 2
      smallestIndex := index

      if leftChildIndex < h.size && h.array[leftChildIndex] < h.array[smallestIndex] {
         smallestIndex = leftChildIndex
      }

      if rightChildIndex < h.size && h.array[rightChildIndex] < h.array[smallestIndex] {
         smallestIndex = rightChildIndex
      }

      if smallestIndex == index {
         break
      }

      h.array[index], h.array[smallestIndex] = h.array[smallestIndex], h.array[index]
      index = smallestIndex
   }
}

func main() {
   heap := NewBinaryHeap()

   heap.Insert(50)
   heap.Insert(30)
   heap.Insert(80)
   heap.Insert(20)
   heap.Insert(10)

   fmt.Println("The Heap elements given here are:", heap.array)
   fmt.Println("The minimum element deleted here is:", heap.Delete())
   fmt.Println("The Heap elements after deletion are:", heap.array)
}