package main

import(
	"github.com/Douirat/Heap-DST/heap"
)

func main(){
	people := &heap.People{}
	people.Insert("Bennacer Douirat", 33)
	people.Insert("Oumaima Fahsi", 25)
	people.Insert("Yousef Zone01", 22)
	people.Insert("Joe Roegan", 52)
	people.Insert("Marry Jane", 7)
	people.Insert("Jordan Peterson", 62)
	// people.MinHeap()
	// people.Heap = people.Heap[1:]
	// people.MinHeap()
	people.Display()
}