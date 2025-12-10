package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type MathProblem struct {
	Numbers  []int
	Operator string
}

func (m MathProblem) Calculate() (total int) {
	total = m.Numbers[0]
	for i := 1; i < len(m.Numbers); i++ {
		if m.Operator == "+" {
			total += m.Numbers[i]
			continue
		}
		if m.Operator == "*" {
			total *= m.Numbers[i]
			continue
		}
	}
	return
}

func main() {
	var mathProblems []MathProblem

	if len(os.Args) < 2 || (os.Args[1] != "part1" && os.Args[1] != "part2") {
		fmt.Fprintln(os.Stderr, "Usage: main <part1|part2>")
		os.Exit(1)
	}

	var total int

	if os.Args[1] == "part1" {
		mathProblems = part1(os.Stdin)
	} else if os.Args[1] == "part2" {
		mathProblems = part2(os.Stdin)
	}

	for _, mathProblem := range mathProblems {
		total += mathProblem.Calculate()
	}

	fmt.Println(total)
}

func part1(reader io.Reader) []MathProblem {
	var mathProblems []MathProblem

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		if len(mathProblems) == 0 {
			mathProblems = make([]MathProblem, len(fields))
		}

		for index, field := range fields {
			if field == "*" || field == "+" {
				mathProblems[index].Operator = field
				continue
			}

			number, err := strconv.Atoi(field)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
			mathProblems[index].Numbers = append(mathProblems[index].Numbers, number)
		}
	}

	return mathProblems
}

func part2(reader io.Reader) []MathProblem {
	var rows [][]string
	var mathProblems []MathProblem

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		rows = append(rows, strings.Split(scanner.Text(), ""))
	}

	operators := rows[len(rows)-1]
	rows = rows[:len(rows)-1]

	var number int
	for i := range operators {
		// empty column
		if i+1 < len(operators) && (operators[i+1] == "*" || operators[i+1] == "+") {
			continue
		}

		if operators[i] == "*" || operators[i] == "+" {
			mathProblems = append(mathProblems, MathProblem{[]int{}, operators[i]})
		}

		number = 0
		for row := 0; row < len(rows); row++ {
			if rows[row][i] == " " {
				continue
			}

			digit, err := strconv.Atoi(rows[row][i])
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}

			number = number*10 + digit
		}
		mathProblems[len(mathProblems)-1].Numbers = append(mathProblems[len(mathProblems)-1].Numbers, number)
	}

	return mathProblems
}
