package main

import "fmt"

// jobs
type number struct {
	a int
	b int
}

// result
type sum struct {
	result int
}

func worker(jobsCh <-chan number, resultsCh chan<- sum) {
	for job := range jobsCh {
		resultsCh <- sum{result: job.a + job.b}
	}
}

func main() {

	// Basic Sync
	// var wg sync.WaitGroup
	// for i := range a {
	// 	wg.Add(1)
	// 	go func(i int) {
	// 		defer wg.Done()
	// 		fmt.Printf("%v ", a[i])
	// 	}(i)
	// }
	// wg.Wait()

	// Woker pool
	nums := []number{
		{a: 1, b: 2},
		{a: 3, b: 4},
		{a: 5, b: 6},
		{a: 7, b: 8},
		{a: 9, b: 10},
		{a: 1, b: 2},
		{a: 3, b: 4},
		{a: 5, b: 6},
		{a: 7, b: 8},
		{a: 9, b: 10},
	}

	jobsCh := make(chan number, len(nums))
	resultsCh := make(chan sum, len(nums))

	for _, num := range nums {
		jobsCh <- num
	}
	close(jobsCh)

	numberWorker := 2
	for w := 0; w < numberWorker; w++ {
		// go routines
		go worker(jobsCh, resultsCh)
	}

	results := make([]sum, 0)
	for a := 0; a < len(nums); a++ {
		temp := <-resultsCh
		results = append(results, temp)
	}
	fmt.Println(results)
}
