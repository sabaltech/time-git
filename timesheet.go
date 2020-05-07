package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	commitCmd := flag.NewFlagSet("commit", flag.ExitOnError)

	commitDescription := commitCmd.String("desc", "", "Description of Work Done")
	commitHours := commitCmd.Int("hours", 0, "Time Spend in Whole Hours")
	commitMinutes := commitCmd.Int("minutes", 0, "Time in Minutes")

	if len(os.Args) < 2 {
		fmt.Println("expected 'commit' or 'today' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "commit":
		commitCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'commit")
		fmt.Println("  desc:", *commitDescription)
		fmt.Println("  hours:", *commitHours)
		fmt.Println("  minutes:", *commitMinutes)

		save(*commitDescription, *commitHours, *commitMinutes)
	case "today":
		scanToday()
	default:
		fmt.Println("expected 'commit' or 'today' subcommands")
		os.Exit(1)
	}
}
