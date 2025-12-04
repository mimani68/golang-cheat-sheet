package main

import (
	"fmt"
	"strconv"
	"strings"
)

func MissingDigit(str string) string {
	if !strings.Contains(str, "x") {
		fmt.Println("[ERROR] Invalid equation, please determine X variable.")
	}
	result := 0
	el := strings.Split(str, " ")
	numOneStr := el[0]
	operationSymbol := el[1]
	numTwoStr := el[2]
	numThreeStr := el[4]
	variableStr := ""
	// First scenario
	if strings.Contains(numOneStr, "x") {
		variableStr = numOneStr
		numTwo, numTwoErr := strconv.Atoi(numTwoStr)
		if numTwoErr != nil {
			fmt.Println("[ERROR] Number two is not valid DIGIT")
		}
		numThree, numThreeErr := strconv.Atoi(numThreeStr)
		if numThreeErr != nil {
			fmt.Println("[ERROR] Number three is not valid DIGIT")
		}
		result = reverseOperationHandler(operationSymbol, numThree, numTwo)
	}
	// Second scenario
	if strings.Contains(numTwoStr, "x") {
		variableStr = numTwoStr
		numOne, numOneErr := strconv.Atoi(numOneStr)
		if numOneErr != nil {
			fmt.Println("[ERROR] Number one is not valid DIGIT")
		}
		numThree, numThreeErr := strconv.Atoi(numThreeStr)
		if numThreeErr != nil {
			fmt.Println("[ERROR] Number three is not valid DIGIT")
		}
		result = reverseOperationHandler(operationSymbol, numThree, numOne)
	}
	// Third scenario
	if strings.Contains(numThreeStr, "x") {
		variableStr = numThreeStr
		numOne, numOneErr := strconv.Atoi(numOneStr)
		if numOneErr != nil {
			fmt.Println("[ERROR] Number one is not valid DIGIT")
		}
		numTwo, numTwoErr := strconv.Atoi(numTwoStr)
		if numTwoErr != nil {
			fmt.Println("[ERROR] Number two is not valid DIGIT")
		}
		result = operationHandler(operationSymbol, numOne, numTwo)
	}
	return fmt.Sprintf("%s", findVariable(variableStr, result))

}

func operationHandler(operationSymbol string, numOne, numTwo int) int {
	result := 0
	switch operationSymbol {
	case "-":
		result = numOne - numTwo
	case "+":
		result = numOne + numTwo
	case "*":
		result = numOne * numTwo
	case "/":
		result = numOne / numTwo
	}
	return result
}

func reverseOperationHandler(operationSymbol string, numOne, numTwo int) int {
	result := 0
	switch operationSymbol {
	case "-":
		result = numOne + numTwo
	case "+":
		result = numOne - numTwo
	case "*":
		result = numOne / numTwo
	case "/":
		result = numOne * numTwo
	}
	return result
}

func findVariable(variableString string, target int) string {
	variableIndex := strings.Index(variableString, "x")
	output := fmt.Sprintf("%d", target)
	return string(output[variableIndex])
}

func main() {
	fmt.Println(MissingDigit(readline()))
}
