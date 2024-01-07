package main

import "fmt"

type DivisionError struct {
	dividend int
	divisor  int
}

func (e DivisionError) Error() string {
	return fmt.Sprintf("division error: %d / %d", e.dividend, e.divisor)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, DivisionError{dividend: a, divisor: b}
	}
	return a / b, nil
}

func main() {
	fmt.Println(divide(0, 0))
}
