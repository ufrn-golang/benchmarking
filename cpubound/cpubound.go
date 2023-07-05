package cpubound

import (
	"runtime"
	"sync"
)

// Sequential version of the merge sort algorithm
func mergeSortSequential(items []int) []int {
	if len(items) < 2 {
		return items
	}

	middle := len(items) / 2
	firstHalf := mergeSortSequential(items[:middle])
	secondHalf := mergeSortSequential(items[middle:])
	return merge(firstHalf, secondHalf)
}


// Concurrent version of the merge sort algorithm
func mergeSortConcurrent(items []int) []int {
	if len(items) < 2 {
		return items
	}

	var firstHalf, secondHalf []int
	middle := len(items) / 2

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		defer waitGroup.Done()
		firstHalf = mergeSortConcurrent(items[:middle])
	}()

	go func() {
		defer waitGroup.Done()
		secondHalf = mergeSortConcurrent(items[middle:])
	}()

	waitGroup.Wait()
	return merge(firstHalf, secondHalf)
}


// Auxiliary function for merge sort
func merge(first, second []int) []int {
	var seq []int
	var i, j int = 0, 0
	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			seq = append(seq, first[i])
			i++
		} else {
			seq = append(seq, second[j])
			j++
		}
	}

	for ; i < len(first); i++ {
		seq = append(seq, first[i])
	}
	for ; j < len(second); j++ {
		seq = append(seq, second[j])
	}
	return seq
}


// Sequential version of adding numbers to each other
func multipleSumSequential(numbers[] int) int64 {
	var sum int64 = 0
	for _, n := range numbers {
		sum += int64(n)
	}
	return sum
}

// Concurrent version of adding numbers to each other
func multipleSumConcurrent(numbers []int) int64 {
	goroutines := runtime.NumCPU()
	slice := len(numbers) / goroutines
	var sum int64 = 0
	var mutex sync.Mutex

	var waitGroup sync.WaitGroup
	waitGroup.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(g int) {
			start := g * slice
			end := start + slice
			if g == goroutines-1 {
				end = len(numbers)
			}

			var part int
			for _, n := range numbers[start:end] {
				part += n
			}
			
			mutex.Lock()
			sum = sum + int64(part)
			mutex.Unlock()
			waitGroup.Done()
		}(i)
	}

	waitGroup.Wait()
	return sum
}