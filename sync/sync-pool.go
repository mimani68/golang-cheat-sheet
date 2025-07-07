package main

import (
	"bytes"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// `sync.Pool` caches allocated but unused items for later reuse,
// reducing pressure on the garbage collector.
// Ideal for high-allocation scenarios (e.g., HTTP request handling).
//
// Mechanism:
//   - Get()/Put(): Stores and retrieves objects.
//   - Per-P Pool: Each processor (P) has a private cache and shared pool, minimizing locking.
//   - Victim Cache: Survives one GC cycle to retain frequently reused objects.

// Buffer pool for reusing byte buffers
var bufferPool = sync.Pool{
	New: func() interface{} {
		// Create a buffer with initial capacity to avoid frequent reallocations
		return bytes.NewBuffer(make([]byte, 0, 1024))
	},
}

// Simulate processing a request that needs a temporary buffer
func processRequestWithPool(id int, data string) string {
	// Get a buffer from the pool
	buffer := bufferPool.Get().(*bytes.Buffer)

	// Important: Reset the buffer before use
	buffer.Reset()

	// Use the buffer for some processing
	buffer.WriteString(fmt.Sprintf("Processing request %d: ", id))
	buffer.WriteString(data)
	buffer.WriteString(" [PROCESSED]")

	result := buffer.String()

	// Return the buffer to the pool for reuse
	bufferPool.Put(buffer)

	return result
}

// Simulate processing without pool (creates new buffer each time)
func processRequestWithoutPool(id int, data string) string {
	// Create a new buffer each time
	buffer := bytes.NewBuffer(make([]byte, 0, 1024))

	buffer.WriteString(fmt.Sprintf("Processing request %d: ", id))
	buffer.WriteString(data)
	buffer.WriteString(" [PROCESSED]")

	return buffer.String()
}

// Benchmark function to measure performance
func benchmark(name string, fn func(int, string) string, iterations int) {
	// Force garbage collection before starting
	runtime.GC()

	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)

	start := time.Now()

	// Simulate concurrent requests
	var wg sync.WaitGroup
	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			result := fn(id, "sample data for processing")
			_ = result // Use the result to prevent optimization
		}(i)
	}
	wg.Wait()

	duration := time.Since(start)

	// Force garbage collection and measure memory
	runtime.GC()
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)

	fmt.Printf("\n%s Results:\n", name)
	fmt.Printf("  Duration: %v\n", duration)
	fmt.Printf("  Allocations: %d\n", m2.TotalAlloc-m1.TotalAlloc)
	fmt.Printf("  GC Cycles: %d\n", m2.NumGC-m1.NumGC)
}

func main() {
	fmt.Println("=== sync.Pool Example: Buffer Pooling ===")
	fmt.Println("\nThis example demonstrates the benefits of sync.Pool for reusing buffers")
	fmt.Println("in high-throughput scenarios like HTTP request processing.\n")

	const iterations = 10000

	// Demonstrate basic usage
	fmt.Println("1. Basic Usage Example:")
	result1 := processRequestWithPool(1, "hello world")
	result2 := processRequestWithPool(2, "foo bar")
	fmt.Printf("   Result 1: %s\n", result1)
	fmt.Printf("   Result 2: %s\n", result2)

	// Performance comparison
	fmt.Println("\n2. Performance Comparison:")
	fmt.Printf("   Running %d concurrent operations...\n", iterations)

	benchmark("WITHOUT sync.Pool", processRequestWithoutPool, iterations)
	benchmark("WITH sync.Pool", processRequestWithPool, iterations)

	// Show pool statistics
	fmt.Println("\n3. Pool Behavior:")
	fmt.Println("   Getting and putting objects multiple times...")

	// Get several buffers
	buffers := make([]*bytes.Buffer, 5)
	for i := 0; i < 5; i++ {
		buffers[i] = bufferPool.Get().(*bytes.Buffer)
		buffers[i].Reset()
		buffers[i].WriteString(fmt.Sprintf("Buffer %d", i))
		fmt.Printf("   Got buffer %d: %s\n", i, buffers[i].String())
	}

	// Put them back
	for i, buffer := range buffers {
		bufferPool.Put(buffer)
		fmt.Printf("   Returned buffer %d to pool\n", i)
	}

	// Get them again (may reuse existing buffers)
	for i := 0; i < 3; i++ {
		buffer := bufferPool.Get().(*bytes.Buffer)
		buffer.Reset()
		buffer.WriteString(fmt.Sprintf("Reused buffer %d", i))
		fmt.Printf("   Reused buffer: %s\n", buffer.String())
		bufferPool.Put(buffer)
	}

	fmt.Println("\n=== Key Benefits of sync.Pool ===")
	fmt.Println("✓ Reduces garbage collection pressure")
	fmt.Println("✓ Improves performance in high-allocation scenarios")
	fmt.Println("✓ Automatic cleanup during GC cycles")
	fmt.Println("✓ Thread-safe object reuse")
	fmt.Println("✓ Minimal overhead for Get/Put operations")

	fmt.Println("\n=== Best Practices ===")
	fmt.Println("• Always reset/clear objects before reuse")
	fmt.Println("• Use for frequently allocated temporary objects")
	fmt.Println("• Don't rely on objects staying in the pool")
	fmt.Println("• Ideal for buffers, slices, and similar objects")
}

// === sync.Pool Example: Buffer Pooling ===
//
// This example demonstrates the benefits of sync.Pool for reusing buffers
// in high-throughput scenarios like HTTP request processing.
//
// 1. Basic Usage Example:
//    Result 1: Processing request 1: hello world [PROCESSED]
//    Result 2: Processing request 2: foo bar [PROCESSED]
//
// 2. Performance Comparison:
//    Running 10000 concurrent operations...
//
// WITHOUT sync.Pool Results:
//   Duration: 13.796845ms
//   Allocations: 3867328
//   GC Cycles: 2
//
// WITH sync.Pool Results:
//   Duration: 4.274732ms
//   Allocations: 1526208
//   GC Cycles: 1
//
// 3. Pool Behavior:
//    Getting and putting objects multiple times...
//    Got buffer 0: Buffer 0
//    Got buffer 1: Buffer 1
//    Got buffer 2: Buffer 2
//    Got buffer 3: Buffer 3
//    Got buffer 4: Buffer 4
//    Returned buffer 0 to pool
//    Returned buffer 1 to pool
//    Returned buffer 2 to pool
//    Returned buffer 3 to pool
//    Returned buffer 4 to pool
//    Reused buffer: Reused buffer 0
//    Reused buffer: Reused buffer 1
//    Reused buffer: Reused buffer 2
//
// === Key Benefits of sync.Pool ===
// ✓ Reduces garbage collection pressure
// ✓ Improves performance in high-allocation scenarios
// ✓ Automatic cleanup during GC cycles
// ✓ Thread-safe object reuse
// ✓ Minimal overhead for Get/Put operations
//
// === Best Practices ===
// • Always reset/clear objects before reuse
// • Use for frequently allocated temporary objects
// • Don't rely on objects staying in the pool
// • Ideal for buffers, slices, and similar objects
// dev@pc ~/p/g/golang-cheat-sheet (master)
