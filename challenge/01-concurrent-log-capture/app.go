package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

type LogRecord struct {
	Timestamp string
	Level     string
	Message   string
}

func generateLogRecords(numLogs int) <-chan LogRecord {
	logRecords := make(chan LogRecord, numLogs)

	var wg sync.WaitGroup
	for i := 1; i <= numLogs; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

			timestamp := time.Now().Format("2006-01-02 15:04:05")
			level := []string{"INFO", "WARNING", "ERROR", "DEBUG"}[rand.Intn(4)]
			message := fmt.Sprintf("Log message %d", id)

			logRecords <- LogRecord{
				Timestamp: timestamp,
				Level:     level,
				Message:   message,
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(logRecords)
	}()

	return logRecords
}

func writeLogsToFile(filename string, logRecords <-chan LogRecord, numWorkers int) error {
	// A WaitGroup waits for a collection of goroutines to finish
	var wg sync.WaitGroup
	errChan := make(chan error, numWorkers)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				errChan <- fmt.Errorf("worker %d: %w", workerID, err)
				return
			}
			defer file.Close()

			writer := bufio.NewWriter(file)
			defer writer.Flush()

			for log := range logRecords {
				logLine := fmt.Sprintf("%s [%s] %s\n", log.Timestamp, log.Level, log.Message)
				if _, err := writer.WriteString(logLine); err != nil {
					errChan <- fmt.Errorf("worker %d: %w", workerID, err)
					return
				}
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	const numLogs = 120
	const filename = "logs.txt"
	const numWorkers = 5

	logRecords := generateLogRecords(numLogs)

	if err := writeLogsToFile(filename, logRecords, numWorkers); err != nil {
		fmt.Printf("Error writing logs to file: %v\n", err)
		return
	}

	fmt.Printf("Successfully wrote %d logs to %s using %d concurrent workers\n", numLogs, filename, numWorkers)
}
