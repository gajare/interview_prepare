package main

import (
	"fmt"
)

// MySlice simulates how slices work internally
type MySlice struct {
	data []int
	len  int
	cap  int
}

// NewMySlice creates a new slice with initial capacity
func NewMySlice(initialCap int) *MySlice {
	return &MySlice{
		data: make([]int, initialCap),
		len:  0,
		cap:  initialCap,
	}
}

// Len returns current length
func (s *MySlice) Len() int {
	return s.len
}

// Cap returns current capacity
func (s *MySlice) Cap() int {
	return s.cap
}

// Append adds a new element, reallocating if needed
func (s *MySlice) Append(value int) {
	if s.len == s.cap {
		// ❗ Capacity full → grow the slice
		newCap := s.growCapacity()
		newData := make([]int, newCap)

		// Copy old data
		for i := 0; i < s.len; i++ {
			newData[i] = s.data[i]
		}

		fmt.Printf("⚠️ Capacity full — reallocating from %d → %d\n", s.cap, newCap)
		s.data = newData
		s.cap = newCap
	}

	s.data[s.len] = value
	s.len++
}

// growCapacity decides how new capacity grows (mimics Go runtime)
func (s *MySlice) growCapacity() int {
	if s.cap < 1024 {
		return s.cap * 2 // double until 1024
	}
	return s.cap + s.cap/4 // grow by 25% beyond 1024
}

// Print shows current slice state
func (s *MySlice) Print() {
	fmt.Printf("Slice: %v | Len: %d | Cap: %d | Addr: %p\n", s.data[:s.len], s.len, s.cap, s.data)
}

func main() {
	// Initial slice with capacity 3
	mySlice := NewMySlice(3)

	fmt.Println("👉 Initial slice state:")
	mySlice.Print()

	// Append elements and watch behavior
	mySlice.Append(1)
	mySlice.Append(2)
	mySlice.Append(3) // still same array

	fmt.Println("\n✅ After 3 appends (capacity not exceeded):")
	mySlice.Print()

	// This append will exceed capacity → reallocation
	mySlice.Append(4)

	fmt.Println("\n🚀 After 4th append (capacity exceeded):")
	mySlice.Print()

	// More appends to trigger another reallocation
	mySlice.Append(5)
	mySlice.Append(6)
	mySlice.Append(7)

	fmt.Println("\n🔥 After multiple appends:")
	mySlice.Print()
}
