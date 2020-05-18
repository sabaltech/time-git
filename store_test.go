package main

import (
	"testing"
	"time"
)

const userID = "jllombart"

// TestGetRecords checks that records are returned
func TestGetRecords(t *testing.T) {
	startDate, _ := time.Parse("2006-01-02", "2020-05-15")
	got := getRecords(userID, startDate, startDate)

	if len(got) != 1 {
		t.Errorf("len(got) = %d; want 1", len(got))
	}

	record := got[0]

	if len(record) != 4 {
		t.Errorf("len(record) = %d; want 4", len(record))
	}

	if record[2] != "" {
		t.Errorf("Expected invoiced hours to be blank; got %s", record[2])
	}

	if record[1] != "6.50" {
		t.Errorf("Expected actual hours to be 6.50; got %s", record[1])
	}

	workDate, err := time.Parse("01/02/06", record[0])

	if err != nil {
		t.Error("Error with correct date layout")
	}

	if workDate.Day() != 15 {
		t.Errorf("workDate.Day() = %d; want 15", workDate.Day())
	}

	if workDate.Month() != 5 {
		t.Errorf("workDate.Month() = %d; want 5", workDate.Month())
	}

	if workDate.Year() != 2020 {
		t.Errorf("workDate.Year() = %d; want 2020", workDate.Year())
	}
}
