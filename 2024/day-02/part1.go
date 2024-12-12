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
	var safeReport bool
	var increasing bool

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		rawReportData := strings.Split(scanner.Text(), " ")

		lastReportValue, _ := strconv.Atoi(rawReportData[0])
		for i := 1; i < len(rawReportData); i++ {
			reportValue, _ := strconv.Atoi(rawReportData[i])

			if i == 1 {
				increasing = reportValue > lastReportValue
			}

			if increasing {
				safeReport = reportValue >= lastReportValue+1 && reportValue <= lastReportValue+3
			} else {
				safeReport = reportValue <= lastReportValue-1 && reportValue >= lastReportValue-3
			}

			if !safeReport {
				break
			}

			lastReportValue = reportValue
		}

		if safeReport {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}
