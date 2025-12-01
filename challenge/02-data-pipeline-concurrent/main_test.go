package main

import (
	"testing"
	"time"
)

// Helper function to reset the global state before each test that uses it.
func resetGlobalState() {
	mu.Lock()
	total = 0
	mu.Unlock()
}

// TestIsEvenNumber tests the pure helper function. (No changes needed)
func TestIsEvenNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"positive_even", 2, true},
		{"positive_odd", 3, false},
		{"zero", 0, true},
		{"negative_even", -4, true},
		{"negative_odd", -5, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := isEvenNumber(tc.input)
			if result != tc.expected {
				t.Errorf("isEvenNumber(%d) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}

// TestPip1 tests the first stage of the pipeline.
func TestPip1(t *testing.T) {
	t.Run("sends_even_number", func(t *testing.T) {
		// FIX: We must add to the WaitGroup because pip1 will call Done().
		wg.Add(1)
		inputCh := make(chan int, 1)
		outputCh := make(chan int, 1)

		go pip1(inputCh, outputCh)

		inputCh <- 4
		close(inputCh)

		select {
		case result := <-outputCh:
			if result != 4 {
				t.Errorf("Expected 4 on output channel, got %d", result)
			}
		case <-time.After(100 * time.Millisecond):
			t.Error("Timed out waiting for result on output channel")
		}

		// Verify the output channel is closed
		_, ok := <-outputCh
		if ok {
			t.Error("Expected output channel to be closed")
		}
		// FIX: Wait for the goroutine to finish and call Done().
		wg.Wait()
	})

	t.Run("does_not_send_odd_number", func(t *testing.T) {
		// FIX: Add to the WaitGroup.
		wg.Add(1)
		inputCh := make(chan int, 1)
		outputCh := make(chan int, 1)

		go pip1(inputCh, outputCh)

		inputCh <- 5
		close(inputCh)

		select {
		case result, ok := <-outputCh:
			if ok {
				t.Errorf("Expected nothing, but received value %d", result)
			}
		case <-time.After(100 * time.Millisecond):
			t.Error("Timed out waiting for channel to close")
		}
		// FIX: Wait for the goroutine to finish.
		wg.Wait()
	})
}

// TestPip2 tests the second stage of the pipeline.
func TestPip2(t *testing.T) {
	t.Run("filters_and_squares_numbers", func(t *testing.T) {
		// FIX: Add to the WaitGroup.
		wg.Add(1)
		inputCh := make(chan int, 3)
		outputCh := make(chan int, 2) // Buffered for expected results

		go pip2(inputCh, outputCh)

		inputCh <- 10  // square is 100 (< 100_000)
		inputCh <- 400 // square is 160_000 (> 100_000)
		inputCh <- 20  // square is 400 (< 100_000)
		close(inputCh)

		var results []int
		for v := range outputCh {
			results = append(results, v)
		}

		expected := []int{100, 400}
		if len(results) != len(expected) {
			t.Errorf("Expected %d results, got %d", len(expected), len(results))
		}
		for i := range expected {
			if results[i] != expected[i] {
				t.Errorf("Result at index %d is %d, want %d", i, results[i], expected[i])
			}
		}
		// FIX: Wait for the goroutine to finish.
		wg.Wait()
	})

	t.Run("handles_empty_input", func(t *testing.T) {
		// FIX: Add to the WaitGroup.
		wg.Add(1)
		inputCh := make(chan int)
		outputCh := make(chan int)

		go pip2(inputCh, outputCh)

		close(inputCh) // Close immediately

		_, ok := <-outputCh
		if ok {
			t.Error("Expected output channel to be closed without sending any value")
		}
		// FIX: Wait for the goroutine to finish.
		wg.Wait()
	})
}

// TestPip3 tests the third stage which has a side effect.
func TestPip3(t *testing.T) {
	t.Run("updates_total_and_sends_result", func(t *testing.T) {
		resetGlobalState() // Ensure a clean state

		// FIX: Add to the WaitGroup.
		wg.Add(1)
		inputCh := make(chan int, 1)
		outputCh := make(chan int, 1)

		go pip3(inputCh, outputCh)

		inputCh <- 100
		close(inputCh)

		select {
		case result := <-outputCh:
			if result != 100 {
				t.Errorf("Expected 100 on output channel, got %d", result)
			}
		case <-time.After(100 * time.Millisecond):
			t.Error("Timed out waiting for result on output channel")
		}

		mu.Lock()
		currentTotal := total
		mu.Unlock()

		if currentTotal != 100 {
			t.Errorf("Expected global total to be 100, got %d", currentTotal)
		}
		// FIX: Wait for the goroutine to finish.
		wg.Wait()
	})

	t.Run("handles_closed_input_channel", func(t *testing.T) {
		resetGlobalState() // Ensure a clean state

		// FIX: Add to the WaitGroup.
		wg.Add(1)
		inputCh := make(chan int)
		outputCh := make(chan int)

		go pip3(inputCh, outputCh)

		close(inputCh) // Close input without sending a value

		_, ok := <-outputCh
		if ok {
			t.Error("Expected output channel to be closed without sending a value")
		}

		mu.Lock()
		currentTotal := total
		mu.Unlock()

		if currentTotal != 0 {
			t.Errorf("Expected global total to remain 0, got %d", currentTotal)
		}
		// FIX: Wait for the goroutine to finish.
		wg.Wait()
	})
}

// BenchmarkIsEvenNumber provides a simple performance benchmark. (No changes needed)
func BenchmarkIsEvenNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isEvenNumber(i)
	}
}

//
// # Run all unit tests
// go test -v
//
// # Run tests with race condition detection (good practice)
// go test -race -v
//
// # Run the benchmark
// go test -bench=.
//
//
// === RUN   TestIsEvenNumber
// === RUN   TestIsEvenNumber/positive_even
// === RUN   TestIsEvenNumber/positive_odd
// === RUN   TestIsEvenNumber/zero
// === RUN   TestIsEvenNumber/negative_even
// === RUN   TestIsEvenNumber/negative_odd
// --- PASS: TestIsEvenNumber (0.00s)
//     --- PASS: TestIsEvenNumber/positive_even (0.00s)
//     --- PASS: TestIsEvenNumber/positive_odd (0.00s)
//     --- PASS: TestIsEvenNumber/zero (0.00s)
//     --- PASS: TestIsEvenNumber/negative_even (0.00s)
//     --- PASS: TestIsEvenNumber/negative_odd (0.00s)
// === RUN   TestPip1
// === RUN   TestPip1/sends_even_number
// === RUN   TestPip1/does_not_send_odd_number
// --- PASS: TestPip1 (0.00s)
//     --- PASS: TestPip1/sends_even_number (0.00s)
//     --- PASS: TestPip1/does_not_send_odd_number (0.00s)
// === RUN   TestPip2
// === RUN   TestPip2/filters_and_squares_numbers
// === RUN   TestPip2/handles_empty_input
// --- PASS: TestPip2 (0.00s)
//     --- PASS: TestPip2/filters_and_squares_numbers (0.00s)
//     --- PASS: TestPip2/handles_empty_input (0.00s)
// === RUN   TestPip3
// === RUN   TestPip3/updates_total_and_sends_result
// === RUN   TestPip3/handles_closed_input_channel
// --- PASS: TestPip3 (0.00s)
//     --- PASS: TestPip3/updates_total_and_sends_result (0.00s)
//     --- PASS: TestPip3/handles_closed_input_channel (0.00s)
// PASS
// ok      example.com/m/v2        0.004s
