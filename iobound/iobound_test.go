package iobound

import "testing"

var urls = make([]string, 10)
func init() {
	for i := range urls {
    	urls[i] = "https://www.google.com"
    }
}

func BenchmarkGetURLSequential(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		getURLSequential(urls)
	}
	b.StopTimer()
	b.ReportMetric(b.Elapsed().Seconds() / float64(b.N), "s/op") 
}

func BenchmarkGetURLConcurrent(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		getURLConcurrent(urls)
	}
	b.StopTimer()
	b.ReportMetric(b.Elapsed().Seconds() / float64(b.N), "s/op") 
}