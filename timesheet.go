package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

const layout = "2006-01-02"

func main() {
	commitCmd := flag.NewFlagSet("commit", flag.ExitOnError)

	commitDescription := commitCmd.String("desc", "", "Description of Work Done")
	commitHours := commitCmd.Int("hours", 0, "Time Spend in Whole Hours")
	commitMinutes := commitCmd.Int("minutes", 0, "Time in Minutes")
	commitDate := commitCmd.String("date", "", "Date Work Completed")

	if len(os.Args) < 2 {
		fmt.Println("expected 'commit' or 'today' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "commit":
		commitCmd.Parse(os.Args[2:])

		workDate := time.Now()

		// Validate Date
		if *commitDate != "" && *commitDate != "today" && *commitDate != "timesheet" {
			parsedDate, err := time.Parse(layout, *commitDate)
			workDate = parsedDate

			if err != nil {
				fmt.Println("Invalid Date Format - ", *commitDate)
				os.Exit(1)
			}
		}

		// Validate Time Spent
		if *commitHours <= 0 && *commitMinutes <= 0 {
			fmt.Println("Work Time Spent must be greater than 0")
			os.Exit(1)
		}

		fmt.Println("subcommand 'commit")
		fmt.Println("  desc:", *commitDescription)
		fmt.Println("  hours:", *commitHours)
		fmt.Println("  minutes:", *commitMinutes)
		fmt.Println("  date:", *commitDate)

		save(*commitDescription, *commitHours, *commitMinutes, workDate)
	case "today":
		printTimesheet("jllombart", time.Now(), time.Now())
	case "timesheet":
		printTimesheet("jllombart", getMondayDate(time.Now()).AddDate(0, 0, -4), time.Now())
	default:
		fmt.Println("expected 'commit' or 'today' subcommands")
		os.Exit(1)
	}
}
