package main
import "fmt"

type Stack []int   //stack

// Push adds a value to the top of the stack.
func (st *Stack) Push(v int) {
   *st = append(*st, v)
}

// Pop removes and returns the top value from the stack.
func (st *Stack) Pop() int {
   if st.IsEmpty() {
      return 0
   }
   top := (*st)[len(*st)-1]
   *st = (*st)[:len(*st)-1]
   return top
}

func (st *Stack) IsEmpty() bool {
   return len(*st) == 0
}

func main() {
   st := Stack{}
   st.Push(1)
   st.Push(2)
   fmt.Println("The value popped from the stack is given as:")
   fmt.Println(st.Pop())
   fmt.Println(st.Pop())
   fmt.Println("Is the stack empty?")
   fmt.Println(st.IsEmpty())
}