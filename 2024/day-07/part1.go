package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var total int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		valueAndNumbers := strings.Split(scanner.Text(), ": ")
		value, _ := strconv.Atoi(valueAndNumbers[0])
		numbers := convertToIntSlice(strings.Split(valueAndNumbers[1], " "))

		if canEvaluateToValue(value, numbers[1:], numbers[0]) {
			total += value
		}
	}

	fmt.Println(total)
}

func canEvaluateToValue(value int, numbers []int, currentTotal int) bool {
	if len(numbers) == 1 {
		return currentTotal*numbers[0] == value || currentTotal+numbers[0] == value
	}

	return canEvaluateToValue(value, numbers[1:], currentTotal*numbers[0]) || canEvaluateToValue(value, numbers[1:], currentTotal+numbers[0])
}

func convertToIntSlice(rawNumbers []string) []int {
	numbers := make([]int, len(rawNumbers))
	for i, rawNumber := range rawNumbers {
		numbers[i], _ = strconv.Atoi(rawNumber)
	}
	return numbers
}
