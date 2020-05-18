package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// Headers mm/dd/yy , actual hours, invoiced hours, description

func writeCSV(userID string, date time.Time) {
	file, _ := os.Create("result.csv")
	year, month, _ := date.Date()
	startDate := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	records := getRecords(userID, startDate, date)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	groupRecords := groupRecords(records)

	for _, record := range groupRecords {
		if err := writer.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}
}

func groupRecords(records [][]string) [][]string {
	recordMap := make(map[string][]string)

	for _, record := range records {
		if recordMap[record[0]] != nil {
			recordMap[record[0]] = combineRows(recordMap[record[0]], record)
		} else {
			recordMap[record[0]] = record
		}
	}

	results := [][]string{}

	for key := range recordMap {
		results = append(results, recordMap[key])
	}

	return results
}

func combineRows(acc []string, curr []string) []string {
	date, accHours, _, accDesc := desconstructRecord(acc)
	_, curHours, _, curDesc := desconstructRecord(curr)

	fAccHours, _ := strconv.ParseFloat(accHours, 64)
	fCurHours, _ := strconv.ParseFloat(curHours, 64)

	hours := fAccHours + fCurHours

	return []string{date, fmt.Sprintf("%.2f", hours), "", fmt.Sprintf("%s; %s", accDesc, curDesc)}
}

func desconstructRecord(record []string) (string, string, string, string) {
	return record[0], record[1], record[2], record[3]
}
