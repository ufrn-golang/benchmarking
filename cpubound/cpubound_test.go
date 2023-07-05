package cpubound

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func generateRandomNumbers(n int) []int {
	numbers := make([]int, n)
	seed := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(seed)
	
	for i := 0; i < n; i++ {
		numbers[i] = rand.Intn(n)
	}
	return numbers
}

// Simple unit test functions
func TestMergeSortSequential(t *testing.T) {
	unsorted := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
    sorted := mergeSortSequential(unsorted)
	fmt.Println(sorted)
}

func TestMergeSortConcurrent(t *testing.T) {
	unsorted := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
    sorted := mergeSortConcurrent(unsorted)
	fmt.Println(sorted)
}

func TestMultipleSumSequential(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(multipleSumSequential(numbers))
}

func TestMultipleSumConcurrent(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(multipleSumConcurrent(numbers))
}


// Benchmarking functions
func BenchmarkMergeSortSequential(b *testing.B) {
	numbers := generateRandomNumbers(1<<20)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergeSortSequential(numbers)
	}
	b.StopTimer()
	b.ReportMetric(b.Elapsed().Seconds()/float64(b.N), "s/op")
}

func BenchmarkMergeSortConcurrent(b *testing.B) {
	numbers := generateRandomNumbers(1<<20)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mergeSortConcurrent(numbers)
	}
	b.StopTimer()
	b.ReportMetric(b.Elapsed().Seconds()/float64(b.N), "s/op")
}