package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var safeReports = 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		rawReportData := strings.Split(scanner.Text(), " ")

		reportData := make([]int, len(rawReportData))
		for i, reportValue := range rawReportData {
			reportData[i], _ = strconv.Atoi(reportValue)
		}

		if isReportSafe(reportData) {
			safeReports++
			continue
		}

		for i := 0; i < len(reportData); i++ {
			if isReportSafe(removeByIndex(reportData, i)) {
				safeReports++
				break
			}
		}
	}

	fmt.Println(safeReports)
}

func isReportSafe(reportData []int) bool {
	var increasing bool
	var safeReport bool

	lastReportValue := reportData[0]
	for i := 1; i < len(reportData); i++ {
		reportValue := reportData[i]

		if i == 1 {
			increasing = reportValue > lastReportValue
		}

		if increasing {
			safeReport = reportValue >= lastReportValue+1 && reportValue <= lastReportValue+3
		} else {
			safeReport = reportValue <= lastReportValue-1 && reportValue >= lastReportValue-3
		}

		if !safeReport {
			return false
		}

		lastReportValue = reportValue
	}

	return true
}

func removeByIndex(slice []int, index int) []int {
	newSlice := make([]int, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:index]...)
	return append(newSlice, slice[index+1:]...)
}
