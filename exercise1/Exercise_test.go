package exercise1

import "testing"

// test function with example input
func TestForExampleInputGetHighestCalories(t *testing.T) {

	entries := createSliceOfSummedCalories("data/exercise1_example.txt")

	if sumNEntriesOfSlice(entries, 1) != 24000 {
		t.Errorf("Expected String(%d) is not same as"+
			" actual (%d)", 24000, sumNEntriesOfSlice(entries, 1))
	}
}

// test function with real input
func TestForRealInputGetHighestCalories(t *testing.T) {

	entries := createSliceOfSummedCalories("data/exercise1_input.txt")

	sumOfEntries := sumNEntriesOfSlice(entries, 1)

	if sumOfEntries != 69836 {
		t.Errorf("Expected (%d) is not same as"+
			" actual (%d)", 69836, sumOfEntries)
	}
}

// test function with example input
func TestForExampleGetHighestCaloriesTopN(t *testing.T) {

	entries := createSliceOfSummedCalories("data/exercise1_example.txt")

	sumOfEntries := sumNEntriesOfSlice(entries, 3)

	if sumOfEntries != 45000 {
		t.Errorf("Expected (%d) is not same as"+
			" actual (%d)", 45000, sumOfEntries)
	}
}

// test function with real input
func TestForRealInputGetHighestCaloriesTopN(t *testing.T) {

	entries := createSliceOfSummedCalories("data/exercise1_input.txt")

	sumOfEntries := sumNEntriesOfSlice(entries, 3)

	if sumOfEntries != 207968 {
		t.Errorf("Expected (%d) is not same as"+
			" actual (%d)", 207968, sumOfEntries)
	}
}
