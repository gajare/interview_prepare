package main

import "fmt"

type Duplicate struct {
	value int
	count int
}

func main() {
	s1 := []int{1, 2, 4, 1, 4, 2, 5, 2}
	skipIndex := []int{}
	results := []Duplicate{}

	for i := 0; i < len(s1); i++ {
		shouldSkip := false
		for k := 0; k < len(skipIndex); k++ {
			if skipIndex[k] == i {
				shouldSkip = true // Found 'i' in the skip list
				break             // No need to check the rest of skipIndex
			}
		}

		if shouldSkip {
			fmt.Println("Skipping element (already counted):", s1[i])
			continue // Skip to the next iteration of the outer loop (i++)
		}

		// --- 2. Process the element (It's a unique start of a count) ---
		d := Duplicate{
			value: s1[i],
			count: 1, // Start count at 1 for the current element s1[i]
		}

		// Inner loop checks for duplicates *after* the current index 'i'
		for j := i + 1; j < len(s1); j++ {
			if s1[i] == s1[j] {
				d.count += 1
				skipIndex = append(skipIndex, j)
			}
		}

		// Store the result
		results = append(results, d)
	}

	// --- Print Final Results ---
	fmt.Println("\n--- Final Counts ---")
	for _, item := range results {
		fmt.Printf("Value %d: Count %d\n", item.value, item.count)
	}
}
