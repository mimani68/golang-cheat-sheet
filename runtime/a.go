package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
	"runtime/trace"
	"strings"
	"sync"
	"time"
)

// =============================================================================
// GOLANG RUNTIME PACKAGE - ADVANCED USE CASES CHEAT SHEET
// =============================================================================

// 1. ADVANCED GOROUTINE MANAGEMENT & MONITORING
// Use Case: Monitor goroutine leaks and manage goroutine lifecycle
func goroutineManagement() {
	fmt.Println("=== 1. ADVANCED GOROUTINE MANAGEMENT ===")

	// Get initial goroutine count
	initialGoroutines := runtime.NumGoroutine()
	fmt.Printf("Initial goroutines: %d\n", initialGoroutines)

	// Create a goroutine leak detector
	var wg sync.WaitGroup
	leakDetector := make(chan struct{})

	// Spawn multiple goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			select {
			case <-leakDetector:
				fmt.Printf("Goroutine %d: received shutdown signal\n", id)
			case <-time.After(2 * time.Second):
				fmt.Printf("Goroutine %d: timeout completed\n", id)
			}
		}(i)
	}

	// Monitor goroutine count
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()
		for range ticker.C {
			current := runtime.NumGoroutine()
			fmt.Printf("Current goroutines: %d\n", current)
			if current <= initialGoroutines+1 { // +1 for this monitoring goroutine
				break
			}
		}
	}()

	// Graceful shutdown
	time.Sleep(1 * time.Second)
	close(leakDetector)
	wg.Wait()

	fmt.Printf("Final goroutines: %d\n", runtime.NumGoroutine())
}

// 2. MEMORY MANAGEMENT & GARBAGE COLLECTION CONTROL
// Use Case: Fine-tune GC behavior for high-performance applications
func memoryManagement() {
	fmt.Println("\n=== 2. MEMORY MANAGEMENT & GC CONTROL ===")

	// Get initial memory stats
	var m1, m2 runtime.MemStats
	runtime.ReadMemStats(&m1)
	fmt.Printf("Initial memory usage: %d KB\n", m1.Alloc/1024)

	// Allocate memory intentionally
	data := make([][]byte, 1000)
	for i := range data {
		data[i] = make([]byte, 1024) // 1KB each
	}

	runtime.ReadMemStats(&m2)
	fmt.Printf("After allocation: %d KB\n", m2.Alloc/1024)

	// Force garbage collection
	runtime.GC()

	// Set GC target percentage (lower = more frequent GC)
	oldPercent := debug.SetGCPercent(50)
	fmt.Printf("Old GC percent: %d, New: 50\n", oldPercent)

	// Monitor GC stats
	var m3 runtime.MemStats
	runtime.ReadMemStats(&m3)
	fmt.Printf("After GC - Alloc: %d KB, NumGC: %d\n", m3.Alloc/1024, m3.NumGC)

	// Restore original GC settings
	debug.SetGCPercent(oldPercent)

	// Clear reference to allow GC
	data = nil
	runtime.GC()
}

// 3. STACK TRACE ANALYSIS & DEBUGGING
// Use Case: Advanced debugging and error reporting with stack traces
func stackTraceAnalysis() {
	fmt.Println("\n=== 3. STACK TRACE ANALYSIS ===")

	// Get stack trace programmatically
	funcA := func() {
		funcB := func() {
			funcC := func() {
				// Capture stack trace
				buf := make([]byte, 4096)
				n := runtime.Stack(buf, false)
				fmt.Printf("Stack trace (current goroutine):\n%s\n", buf[:n])

				// Get caller information
				pc, file, line, ok := runtime.Caller(0)
				if ok {
					fn := runtime.FuncForPC(pc)
					fmt.Printf("Current function: %s\n", fn.Name())
					fmt.Printf("File: %s, Line: %d\n", file, line)
				}

				// Get multiple callers
				pcs := make([]uintptr, 10)
				n = runtime.Callers(0, pcs)
				frames := runtime.CallersFrames(pcs[:n])

				fmt.Println("Call stack:")
				for {
					frame, more := frames.Next()
					fmt.Printf("  %s:%d %s\n", frame.File, frame.Line, frame.Function)
					if !more {
						break
					}
				}
			}
			funcC()
		}
		funcB()
	}
	funcA()
}

// 4. CPU PROFILING & PERFORMANCE MONITORING
// Use Case: Profile CPU usage for performance optimization
func cpuProfiling() {
	fmt.Println("\n=== 4. CPU PROFILING ===")

	// Create CPU profile
	cpuFile, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Printf("Error creating CPU profile: %v\n", err)
		return
	}
	defer cpuFile.Close()

	// Start CPU profiling
	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		fmt.Printf("Error starting CPU profile: %v\n", err)
		return
	}
	defer pprof.StopCPUProfile()

	// Simulate CPU-intensive work
	fmt.Println("Performing CPU-intensive work...")
	result := 0
	for i := 0; i < 1000000; i++ {
		result += i * i
	}

	// Get CPU information
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	fmt.Printf("Result: %d\n", result)

	fmt.Println("CPU profile saved to cpu.prof")
}

// 5. MEMORY PROFILING & LEAK DETECTION
// Use Case: Detect memory leaks and optimize memory usage
func memoryProfiling() {
	fmt.Println("\n=== 5. MEMORY PROFILING ===")

	// Create memory profile
	memFile, err := os.Create("mem.prof")
	if err != nil {
		fmt.Printf("Error creating memory profile: %v\n", err)
		return
	}
	defer memFile.Close()

	// Allocate memory to profile
	var memoryHogs [][]byte
	for i := 0; i < 100; i++ {
		memoryHogs = append(memoryHogs, make([]byte, 1024*1024)) // 1MB each
	}

	// Force GC before profiling
	runtime.GC()

	// Write heap profile
	if err := pprof.WriteHeapProfile(memFile); err != nil {
		fmt.Printf("Error writing heap profile: %v\n", err)
		return
	}

	// Get detailed memory stats
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Heap alloc: %d MB\n", m.HeapAlloc/1024/1024)
	fmt.Printf("Heap sys: %d MB\n", m.HeapSys/1024/1024)
	fmt.Printf("Heap objects: %d\n", m.HeapObjects)
	fmt.Printf("GC cycles: %d\n", m.NumGC)

	fmt.Println("Memory profile saved to mem.prof")
}

// 6. GOROUTINE PROFILING & DEADLOCK DETECTION
// Use Case: Analyze goroutine states and detect potential deadlocks
func goroutineProfiling() {
	fmt.Println("\n=== 6. GOROUTINE PROFILING ===")

	// Create goroutine profile
	goroutineFile, err := os.Create("goroutine.prof")
	if err != nil {
		fmt.Printf("Error creating goroutine profile: %v\n", err)
		return
	}
	defer goroutineFile.Close()

	// Create various goroutine states
	var wg sync.WaitGroup
	ch := make(chan int)

	// Blocked goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ch // This will block
	}()

	// CPU-bound goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			_ = i * i
		}
	}()

	// Sleeping goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
	}()

	// Wait a bit for goroutines to reach their states
	time.Sleep(10 * time.Millisecond)

	// Write goroutine profile
	if err := pprof.Lookup("goroutine").WriteTo(goroutineFile, 1); err != nil {
		fmt.Printf("Error writing goroutine profile: %v\n", err)
		return
	}

	fmt.Printf("Active goroutines: %d\n", runtime.NumGoroutine())

	// Unblock and wait for completion
	close(ch)
	wg.Wait()

	fmt.Println("Goroutine profile saved to goroutine.prof")
}

// 7. BUILD INFORMATION & VERSION DETECTION
// Use Case: Runtime version detection and build information
func buildInformation() {
	fmt.Println("\n=== 7. BUILD INFORMATION ===")

	// Get build information
	buildInfo, ok := debug.ReadBuildInfo()
	if ok {
		fmt.Printf("Module path: %s\n", buildInfo.Path)
		fmt.Printf("Go version: %s\n", buildInfo.GoVersion)

		// Print build settings
		fmt.Println("Build settings:")
		for _, setting := range buildInfo.Settings {
			fmt.Printf("  %s: %s\n", setting.Key, setting.Value)
		}

		// Print dependencies
		fmt.Println("Dependencies:")
		for _, dep := range buildInfo.Deps {
			fmt.Printf("  %s %s\n", dep.Path, dep.Version)
		}
	}

	// Get runtime information
	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Printf("Go root: %s\n", runtime.GOROOT())
	fmt.Printf("Compiler: %s\n", runtime.Compiler)
	fmt.Printf("Architecture: %s\n", runtime.GOARCH)
	fmt.Printf("OS: %s\n", runtime.GOOS)
}

// 8. EXECUTION TRACING & PERFORMANCE ANALYSIS
// Use Case: Detailed execution tracing for performance analysis
func executionTracing() {
	fmt.Println("\n=== 8. EXECUTION TRACING ===")

	// Create trace file
	traceFile, err := os.Create("trace.out")
	if err != nil {
		fmt.Printf("Error creating trace file: %v\n", err)
		return
	}
	defer traceFile.Close()

	// Start tracing
	if err := trace.Start(traceFile); err != nil {
		fmt.Printf("Error starting trace: %v\n", err)
		return
	}
	defer trace.Stop()

	// Create tasks and regions for tracing
	ctx := context.Background()

	// Task 1: CPU-intensive work
	task1 := trace.NewTask(ctx, "cpu-intensive")
	func() {
		defer task1.End()

		region := trace.StartRegion(ctx, "calculation")
		defer region.End()

		result := 0
		for i := 0; i < 100000; i++ {
			result += i
		}
		fmt.Printf("CPU task result: %d\n", result)
	}()

	// Task 2: Concurrent work
	task2 := trace.NewTask(ctx, "concurrent-work")
	func() {
		defer task2.End()

		var wg sync.WaitGroup
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				region := trace.StartRegion(ctx, fmt.Sprintf("worker-%d", id))
				defer region.End()

				time.Sleep(10 * time.Millisecond)
				fmt.Printf("Worker %d completed\n", id)
			}(i)
		}
		wg.Wait()
	}()

	fmt.Println("Execution trace saved to trace.out")
	fmt.Println("View with: go tool trace trace.out")
}

// 9. RUNTIME MONITORING & HEALTH CHECKS
// Use Case: Real-time runtime monitoring for production systems
func runtimeMonitoring() {
	fmt.Println("\n=== 9. RUNTIME MONITORING ===")

	// Create a monitoring system
	monitor := &RuntimeMonitor{
		interval: 1 * time.Second,
		done:     make(chan bool),
	}

	// Start monitoring
	go monitor.Start()

	// Simulate some work
	time.Sleep(3 * time.Second)

	// Stop monitoring
	monitor.Stop()
}

type RuntimeMonitor struct {
	interval time.Duration
	done     chan bool
}

func (rm *RuntimeMonitor) Start() {
	ticker := time.NewTicker(rm.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			rm.collectMetrics()
		case <-rm.done:
			return
		}
	}
}

func (rm *RuntimeMonitor) Stop() {
	rm.done <- true
}

func (rm *RuntimeMonitor) collectMetrics() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Runtime Metrics - Goroutines: %d, Memory: %d KB, GC: %d\n",
		runtime.NumGoroutine(), m.Alloc/1024, m.NumGC)
}

// 10. ADVANCED RUNTIME HOOKS & FINALIZERS
// Use Case: Resource cleanup and advanced lifecycle management
func runtimeHooks() {
	fmt.Println("\n=== 10. RUNTIME HOOKS & FINALIZERS ===")

	// Resource with finalizer
	type Resource struct {
		name string
		data []byte
	}

	// Create resource with finalizer
	r := &Resource{
		name: "test-resource",
		data: make([]byte, 1024),
	}

	// Set finalizer
	runtime.SetFinalizer(r, (*Resource).cleanup)
	fmt.Printf("Created resource: %s\n", r.name)

	// Demonstrate finalizer behavior
	r = nil // Remove reference

	// Force GC to trigger finalizer
	runtime.GC()
	runtime.GC() // Call twice to ensure finalizer runs

	// Lock OS Thread (useful for C interop)
	runtime.LockOSThread()
	fmt.Println("OS thread locked")

	// Gosched yields the processor
	runtime.Gosched()
	fmt.Println("Yielded processor")

	// Unlock OS Thread
	runtime.UnlockOSThread()
	fmt.Println("OS thread unlocked")

	// Demonstrate keepalive
	keepAliveDemo()
}

func (r *Resource) cleanup() {
	fmt.Printf("Finalizer called for resource: %s\n", r.name)
	// Cleanup resource
	r.data = nil

	// Clear finalizer
	runtime.SetFinalizer(r, nil)
}

func keepAliveDemo() {
	// KeepAlive prevents GC of object
	data := make([]byte, 1024)
	ptr := &data[0]

	// Without KeepAlive, data might be GC'd here
	// since we only use ptr, not data
	fmt.Printf("Pointer: %p\n", ptr)

	// Keep data alive until this point
	runtime.KeepAlive(data)
	fmt.Println("Data kept alive")
}

// =============================================================================
// MAIN FUNCTION - DEMONSTRATES ALL USE CASES
// =============================================================================

func main() {
	fmt.Println("GOLANG RUNTIME PACKAGE - ADVANCED USE CASES")
	fmt.Println("=" + strings.Repeat("=", 50))

	// Run all examples
	goroutineManagement()
	memoryManagement()
	stackTraceAnalysis()
	cpuProfiling()
	memoryProfiling()
	goroutineProfiling()
	buildInformation()
	executionTracing()
	runtimeMonitoring()
	runtimeHooks()

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("All runtime examples completed!")
	fmt.Println("Profile files generated: cpu.prof, mem.prof, goroutine.prof, trace.out")
	fmt.Println("View profiles with: go tool pprof <profile-file>")
	fmt.Println("View trace with: go tool trace trace.out")
}
