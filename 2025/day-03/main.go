package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var batteryBanks [][]int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var batteryBank []int
		for battery := range strings.SplitSeq(scanner.Text(), "") {
			joltage, err := strconv.Atoi(battery)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
			batteryBank = append(batteryBank, joltage)
		}
		batteryBanks = append(batteryBanks, batteryBank)
	}

	fmt.Println(part1(batteryBanks))
	fmt.Println(part2(batteryBanks))
}

func part1(batteryBanks [][]int) (total int) {
	for i := range batteryBanks {
		total += largestBatteryCombo(batteryBanks[i], 2)
	}
	return
}

func part2(batteryBanks [][]int) (total int) {
	for i := range batteryBanks {
		total += largestBatteryCombo(batteryBanks[i], 12)
	}
	return
}

func highestJoltageIndex(batteries []int) (index int) {
	highestJoltage := batteries[0]
	for i := 1; i < len(batteries); i++ {
		if batteries[i] > highestJoltage {
			highestJoltage = batteries[i]
			index = i
		}
	}
	return
}

func largestBatteryCombo(batteryBank []int, batteriesToUse int) (total int) {
	var offset int = 0
	var index int = 0

	for batteriesNeeded := batteriesToUse; batteriesNeeded > 0; batteriesNeeded-- {
		index = highestJoltageIndex(batteryBank[offset:len(batteryBank)-batteriesNeeded+1]) + offset
		offset = index + 1
		total = total*10 + batteryBank[index]
	}

	return
}
