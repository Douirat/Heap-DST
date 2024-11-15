package heap

import "fmt"

// Create a structure to represent People:
type Person struct {
	Name string
	Age  int
}

// Create a structure to represet the heap:
type People struct {
	Heap []*Person
}

// Ins

// Insert a new user respecting the max heap property:
func (people *People) InsertMaxHeap(name string, age int) {
	person := &Person{Name: name, Age: age}
	people.Heap = append(people.Heap, person)
	i := len(people.Heap) - 1
	for i != 0 && people.Heap[(i-1)/2].Age < people.Heap[i].Age {
		fmt.Println(i)
		Temp := people.Heap[(i-1)/2]
		people.Heap[(i-1)/2] = people.Heap[i]
		people.Heap[i] = Temp
		i = (i - 1) / 2
	}
}

// Insert a new user respecting the min heap property:
func (people *People) InsertMinHeap(name string, age int) {
	person := &Person{Name: name, Age: age}
	people.Heap = append(people.Heap, person)
	i := len(people.Heap) - 1
	for i != 0 && people.Heap[(i-1)/2].Age > people.Heap[i].Age {
		Temp := people.Heap[(i-1)/2]
		people.Heap[(i-1)/2] = people.Heap[i]
		people.Heap[i] = Temp
		i = (i - 1) / 2
	}
}

// Display the content of my heap:
func (people *People) Display() {
	for _, v := range people.Heap {
		fmt.Printf("The user name is: %s and his age is: %d\n", v.Name, v.Age)
	}
}

// Heapify down:
func (people *People) HeapifyDownMax(ind int) {
	Max := ind
	Left := (ind * 2) + 1
	Right := (ind * 2) + 2
	if Left < len(people.Heap) && people.Heap[Left].Age > people.Heap[Max].Age {
		Max = Left
	}
	if Right < len(people.Heap) && people.Heap[Right].Age > people.Heap[Max].Age {
		Max = Right
	}
	if Max != ind {
		people.Heap[Max], people.Heap[ind] = people.Heap[ind], people.Heap[Max]
		people.HeapifyDownMax(Max)
	}
}

// Heapify down after each extraction:
func (people *People) HeapifyDownMin(ind int) {
	Min := ind
	Left := (ind * 2) + 1
	Right := (ind * 2) + 2
	if Left < len(people.Heap) && people.Heap[Left].Age < people.Heap[Min].Age {
		Min = Left
	}
	if Right < len(people.Heap) && people.Heap[Right].Age < people.Heap[Min].Age {
		Min = Right
	}
	if Min != ind {
		people.Heap[Min], people.Heap[ind] = people.Heap[ind], people.Heap[Min]
		people.HeapifyDownMin(Min)
	}
}

// Extrract the Max:
func (people *People) GetMax() *Person {
	max := people.Heap[0]
	people.Heap[0] = people.Heap[len(people.Heap)-1]
	people.Heap = people.Heap[:len(people.Heap)-1]
	people.HeapifyDownMax(0)
	return max
}

// Extract the mean:
func (people *People) GetMin() *Person {
	min := people.Heap[0]
	people.Heap[0] = people.Heap[len(people.Heap)-1]
	people.Heap = people.Heap[:len(people.Heap)-1]
	people.HeapifyDownMin(0)
	return min
}