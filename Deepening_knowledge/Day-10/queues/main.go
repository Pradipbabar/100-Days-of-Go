package main 

import ( 
	"errors"
	"fmt"
) 

type Queue struct { 
	Elements []int 
	Size	 int 
} 

func (q *Queue) Enqueue(elem int) { 
	if q.GetLength() == q.Size { 
		fmt.Println("Overflow") 
		return
	} 
	q.Elements = append(q.Elements, elem) 
} 

func (q *Queue) Dequeue() int { 
	if q.IsEmpty() { 
		fmt.Println("UnderFlow") 
		return 0
	} 
	element := q.Elements[0] 
	if q.GetLength() == 1 { 
		q.Elements = nil 
		return element 
	} 
	q.Elements = q.Elements[1:] 
	return element // Slice off the element once it is dequeued. 
} 

func (q *Queue) GetLength() int { 
	return len(q.Elements) 
} 

func (q *Queue) IsEmpty() bool { 
	return len(q.Elements) == 0
} 

func (q *Queue) Peek() (int, error) { 
	if q.IsEmpty() { 
		return 0, errors.New("empty queue") 
	} 
	return q.Elements[0], nil 
} 

func main() { 
	queue := Queue{Size: 3} 
	fmt.Println(queue.Elements) 
	queue.Enqueue(1) 
	fmt.Println(queue.Elements) 
	queue.Enqueue(2) 
	fmt.Println(queue.Elements) 
	queue.Enqueue(3) 
	fmt.Println(queue.Elements) 
	queue.Enqueue(5) 
	fmt.Println(queue.Elements) 
	elem := queue.Dequeue() 
	fmt.Println(elem) 
	fmt.Println(queue.Elements) 
	queue.Enqueue(9) 
	fmt.Println(queue.Elements) 
	elem = queue.Dequeue() 
	fmt.Println(elem) 
	fmt.Println(queue.Elements) 

} 
