package heap

import "fmt"

// Create a structure to represent People:
type Person struct {
Name string
Age int
}

// Create a structure to represet the heap:
type People struct {
	Heap []*Person
}

//Insert a new user:
func (people *People)Insert(name string, age int) {
	person := &Person{Name: name, Age: age}
	people.Heap = append(people.Heap, person)
}

// Display the content of my heap:
func (people *People)Display() {
	for _, v := range people.Heap {
		fmt.Printf("The user name is: %s and his age is: %d\n", v.Name, v.Age)
	}
}

// Heapify my heap:
func (people *People)MinHeap(){
	size := len(people.Heap)-1
	for i:=(size-2); i>=0; i--{
		people.ExtractMinHeap(i)
	}
}

// Heapify the heap to maintain min-heap property
func (people *People) ExtractMinHeap(i int) {
    Min := i
    Left := (i * 2) + 1
    Right := (i * 2) + 2
    size := len(people.Heap)
    
    if Left < size && people.Heap[Left].Age < people.Heap[Min].Age {
        Min = Left
    }
    
    if Right < size && people.Heap[Right].Age < people.Heap[Min].Age {
        Min = Right
    }
    
    // Swap and continue heapifying if needed
    if Min != i {
        people.Heap[Min], people.Heap[i] = people.Heap[i], people.Heap[Min]
        people.ExtractMinHeap(Min) // Recursive call to fix the heap property further down
    }
}
