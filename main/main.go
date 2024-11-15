package main

import (
	"fmt"

	"github.com/Douirat/Heap-DST/heap"
)

func main(){
	people := &heap.People{}
	// people.InsertMaxHeap("Bennacer Douirat", 33)
	// people.InsertMaxHeap("Oumaima Fahsi", 25)
	// people.InsertMaxHeap("Yousef Zone01", 22)
	// people.InsertMaxHeap("Joe Roegan", 52)
	// people.InsertMaxHeap("Marry Jane", 7)
	// people.InsertMaxHeap("Jordan Peterson", 62)
	people.InsertMinHeap("Bennacer Douirat", 33)
	people.InsertMinHeap("Oumaima Fahsi", 25)
	people.InsertMinHeap("Yousef Zone01", 22)
	people.InsertMinHeap("Joe Roegan", 52)
	people.InsertMinHeap("Marry Jane", 7)
	people.InsertMinHeap("Jordan Peterson", 62)
	// max := people.GetMax()
	min := people.GetMin()
	// fmt.Printf("The oldest person is: %s he is: %d\n", max.Name, max.Age)
	fmt.Printf("The youngest person is: %s he is: %d\n", min.Name, min.Age)
	people.Display()
}