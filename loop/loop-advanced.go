package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	// =============================================================================
	// ADVANCED LOOP TRICKS AND PRO TIPS
	// =============================================================================

	// 1. LABELED BREAK AND CONTINUE - Pro Tip for Nested Loops
	fmt.Println("\n=== 1. LABELED BREAK AND CONTINUE ===")
	labeledBreakContinue()

	// 2. LOOP WITH CHANNELS AND SELECT - Concurrent Loop Patterns
	fmt.Println("\n=== 2. LOOP WITH CHANNELS AND SELECT ===")
	channelLoopPatterns()

	// 3. LOOP WITH CLOSURES - Common Gotchas and Solutions
	fmt.Println("\n=== 3. LOOP WITH CLOSURES - GOTCHAS ===")
	loopClosureGotchas()

	// 4. CUSTOM ITERATORS - Advanced Iteration Patterns
	fmt.Println("\n=== 4. CUSTOM ITERATORS ===")
	customIterators()

	// 5. PERFORMANCE OPTIMIZATIONS - Loop Optimization Tips
	fmt.Println("\n=== 5. PERFORMANCE OPTIMIZATIONS ===")
	performanceOptimizations()
}

// 1. LABELED BREAK AND CONTINUE - Pro Tip for Nested Loops
func labeledBreakContinue() {
	fmt.Println("Finding first pair that sums to 10:")

	numbers1 := []int{1, 2, 3, 4, 5}
	numbers2 := []int{6, 7, 8, 9, 10}

	// Using labeled break to exit nested loops
OuterLoop:
	for i, num1 := range numbers1 {
		for j, num2 := range numbers2 {
			fmt.Printf("Checking %d + %d = %d\n", num1, num2, num1+num2)
			if num1+num2 == 10 {
				fmt.Printf("Found pair: %d + %d = 10 at indices [%d, %d]\n", num1, num2, i, j)
				break OuterLoop // Breaks out of both loops
			}
		}
	}

	fmt.Println("\nSkipping even numbers in nested loop:")

	// Using labeled continue to skip outer loop iteration
OuterSkip:
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			fmt.Printf("Skipping outer loop iteration %d\n", i)
			continue OuterSkip
		}

		for j := 1; j <= 3; j++ {
			if j == 2 {
				fmt.Printf("  Skipping inner loop iteration %d\n", j)
				continue // Regular continue for inner loop
			}
			fmt.Printf("  Processing: i=%d, j=%d\n", i, j)
		}
	}
}

// 2. LOOP WITH CHANNELS AND SELECT - Concurrent Loop Patterns
func channelLoopPatterns() {
	// Pattern 1: Worker pool with channels
	fmt.Println("Worker Pool Pattern:")

	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// Start workers
	for w := 1; w <= 3; w++ {
		go func(id int) {
			for job := range jobs {
				fmt.Printf("Worker %d processing job %d\n", id, job)
				time.Sleep(100 * time.Millisecond)
				results <- job * 2
			}
		}(w)
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for r := 1; r <= 5; r++ {
		result := <-results
		fmt.Printf("Result: %d\n", result)
	}

	// Pattern 2: Timeout and context cancellation
	fmt.Println("\nTimeout Pattern:")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	data := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(300 * time.Millisecond)
			select {
			case data <- fmt.Sprintf("data-%d", i):
			case <-ctx.Done():
				fmt.Println("Producer cancelled")
				return
			}
		}
		close(data)
	}()

	// Loop with select for timeout handling
	for {
		select {
		case item, ok := <-data:
			if !ok {
				fmt.Println("Data channel closed")
				return
			}
			fmt.Printf("Received: %s\n", item)
		case <-ctx.Done():
			fmt.Println("Context timeout reached")
			return
		}
	}
}

// 3. LOOP WITH CLOSURES - Common Gotchas and Solutions
func loopClosureGotchas() {
	fmt.Println("Common Closure Gotcha:")

	// WRONG: Common mistake - all goroutines print the same value
	fmt.Println("❌ Wrong way (all print 3):")
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("Goroutine prints: %d\n", i) // All print 3!
		}()
	}
	wg.Wait()

	// CORRECT: Solutions
	fmt.Println("✅ Correct way 1 (pass as parameter):")
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			fmt.Printf("Goroutine prints: %d\n", val)
		}(i)
	}
	wg.Wait()

	fmt.Println("✅ Correct way 2 (capture in new variable):")
	for i := 0; i < 3; i++ {
		wg.Add(1)
		i := i // Create new variable in loop scope
		go func() {
			defer wg.Done()
			fmt.Printf("Goroutine prints: %d\n", i)
		}()
	}
	wg.Wait()

	// Advanced: Using closures for functional programming
	fmt.Println("✅ Advanced: Functional approach with closures:")
	numbers := []int{1, 2, 3, 4, 5}

	// Filter and map in one loop
	result := make([]int, 0)
	for _, num := range numbers {
		if filter := func(n int) bool { return n%2 == 0 }; filter(num) {
			mapper := func(n int) int { return n * n }
			result = append(result, mapper(num))
		}
	}
	fmt.Printf("Even numbers squared: %v\n", result)
}

// 4. CUSTOM ITERATORS - Advanced Iteration Patterns
func customIterators() {
	fmt.Println("Custom Iterator Pattern:")

	// Iterator function that yields values
	fibonacci := func(max int) <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			a, b := 0, 1
			for a < max {
				ch <- a
				a, b = b, a+b
			}
		}()
		return ch
	}

	fmt.Print("Fibonacci sequence: ")
	for num := range fibonacci(100) {
		fmt.Printf("%d ", num)
	}
	fmt.Println()

	// Tree traversal iterator
	type Node struct {
		Value int
		Left  *Node
		Right *Node
	}

	// Build a sample tree
	root := &Node{
		Value: 4,
		Left: &Node{
			Value: 2,
			Left:  &Node{Value: 1},
			Right: &Node{Value: 3},
		},
		Right: &Node{
			Value: 6,
			Left:  &Node{Value: 5},
			Right: &Node{Value: 7},
		},
	}

	// In-order traversal iterator
	inOrderTraversal := func(node *Node) <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			var traverse func(*Node)
			traverse = func(n *Node) {
				if n != nil {
					traverse(n.Left)
					ch <- n.Value
					traverse(n.Right)
				}
			}
			traverse(node)
		}()
		return ch
	}

	fmt.Print("Tree in-order traversal: ")
	for value := range inOrderTraversal(root) {
		fmt.Printf("%d ", value)
	}
	fmt.Println()

	// Batch iterator - process items in batches
	fmt.Println("Batch Iterator Pattern:")
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	batchSize := 3

	for i := 0; i < len(items); i += batchSize {
		end := i + batchSize
		if end > len(items) {
			end = len(items)
		}
		batch := items[i:end]
		fmt.Printf("Processing batch %d: %v\n", i/batchSize+1, batch)

		// Process batch
		sum := 0
		for _, item := range batch {
			sum += item
		}
		fmt.Printf("  Batch sum: %d\n", sum)
	}
}

// 5. PERFORMANCE OPTIMIZATIONS - Loop Optimization Tips
func performanceOptimizations() {
	fmt.Println("Performance Optimization Tips:")

	// Tip 1: Pre-allocate slices when size is known
	fmt.Println("✅ Tip 1: Pre-allocate slices")

	// Inefficient: growing slice
	inefficientSlice := make([]int, 0)
	for i := 0; i < 1000; i++ {
		inefficientSlice = append(inefficientSlice, i)
	}

	// Efficient: pre-allocated slice
	efficientSlice := make([]int, 0, 1000) // capacity 1000
	for i := 0; i < 1000; i++ {
		efficientSlice = append(efficientSlice, i)
	}

	// Even better: direct indexing
	bestSlice := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		bestSlice[i] = i
	}

	fmt.Printf("Created slices of length: %d, %d, %d\n",
		len(inefficientSlice), len(efficientSlice), len(bestSlice))

	// Tip 2: Minimize work in loop condition
	fmt.Println("✅ Tip 2: Cache expensive operations")

	data := make([][]int, 100)
	for i := range data {
		data[i] = make([]int, 100)
		for j := range data[i] {
			data[i][j] = rand.Intn(100)
		}
	}

	// Inefficient: calling len() every iteration
	start := time.Now()
	sum1 := 0
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			sum1 += data[i][j]
		}
	}
	inefficientTime := time.Since(start)

	// Efficient: cache the length
	start = time.Now()
	sum2 := 0
	dataLen := len(data)
	for i := 0; i < dataLen; i++ {
		rowLen := len(data[i])
		for j := 0; j < rowLen; j++ {
			sum2 += data[i][j]
		}
	}
	efficientTime := time.Since(start)

	fmt.Printf("Inefficient loop time: %v, Efficient loop time: %v\n",
		inefficientTime, efficientTime)

	// Tip 3: Use range for slices when you don't need the index
	fmt.Println("✅ Tip 3: Use range appropriately")

	numbers := []int{1, 2, 3, 4, 5}

	// When you need both index and value
	for i, num := range numbers {
		if i > 0 {
			fmt.Printf("Index %d: %d\n", i, num)
		}
	}

	// When you only need the value
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Printf("Sum: %d\n", sum)

	// When you only need the index
	for i := range numbers {
		numbers[i] *= 2
	}
	fmt.Printf("Doubled: %v\n", numbers)

	// Tip 4: Early termination patterns
	fmt.Println("✅ Tip 4: Early termination")

	// Find first element matching condition
	target := 8
	found := false
	for i, num := range numbers {
		if num == target {
			fmt.Printf("Found %d at index %d\n", target, i)
			found = true
			break // Early termination
		}
	}
	if !found {
		fmt.Printf("Target %d not found\n", target)
	}

	// Tip 5: Loop unrolling for performance-critical code
	fmt.Println("✅ Tip 5: Consider loop unrolling for hot paths")

	// Regular loop
	arr := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	sum3 := 0
	for i := 0; i < len(arr); i++ {
		sum3 += arr[i]
	}

	// Unrolled loop (for performance-critical code)
	sum4 := 0
	for i := 0; i < len(arr); i += 4 {
		sum4 += arr[i] + arr[i+1] + arr[i+2] + arr[i+3]
	}

	fmt.Printf("Regular loop sum: %d, Unrolled loop sum: %d\n", sum3, sum4)

	// Tip 6: Memory-efficient iteration for large datasets
	fmt.Println("✅ Tip 6: Memory-efficient iteration")

	// Instead of loading all data into memory
	processLargeDataset := func(batchSize int) {
		// Simulate processing large dataset in batches
		for batch := 0; batch < 5; batch++ {
			fmt.Printf("Processing batch %d of large dataset\n", batch)
			// Process batch, then release memory
			time.Sleep(10 * time.Millisecond)
		}
	}

	processLargeDataset(1000)

	fmt.Println("All performance tips demonstrated!")
}
