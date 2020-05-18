package main

import (
	"testing"
	"time"
)

func TestWriteCsv(t *testing.T) {
	startDate, _ := time.Parse("2006-01-02", "2020-05-15")

	writeCSV("jllombart", startDate)
}

func TestGroupRecords(t *testing.T) {
	description := "Work with Luigi on docker-compose and cassandra resource usage; Work with Andrei on On-Prem to ACX Connectivity"

	records := [][]string{
		[]string{"05/08/20", "8.00", "", "Started to convert the acx-operator watcher to java8"},
		[]string{"05/11/20", "2.00", "", "Work with Luigi on docker-compose and cassandra resource usage"},
		[]string{"05/11/20", "5.00", "", "Work with Andrei on On-Prem to ACX Connectivity"},
	}

	grouped := groupRecords(records)

	if len(grouped) != 2 {
		t.Errorf("len(grouped) = %d; want 1", len(grouped))
	}

	if grouped[1][3] != description {
		t.Errorf("Description not grouped: %s", grouped[0][3])
	}

	if grouped[1][1] != "7.00" {
		t.Errorf("Work hours = %s", grouped[0][1])
	}
}
