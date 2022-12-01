package exercise1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func createSliceOfSummedCalories(fileName string) []int {

	readFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var entries []int
	var sum = 0
	for fileScanner.Scan() {

		entry := fileScanner.Text()
		if entry == "" {
			entries = append(entries, sum)
			sum = 0
			continue
		}

		number, err := strconv.Atoi(entry)
		if err != nil {
			fmt.Println("Error during conversion")
			return nil
		}

		sum += number

	}
	// don't forget the last entry!
	if sum != 0 {
		entries = append(entries, sum)
	}
	// sort reverse
	sort.Sort(sort.Reverse(sort.IntSlice(entries)))

	readFile.Close()
	return entries

}

func sumNEntriesOfSlice(entries []int, n int) int {

	sum := 0
	for i := 1; i <= n; i++ {
		sum += entries[i-1]
	}

	return sum

}
