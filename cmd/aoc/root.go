package aoc

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type Runner interface {
	Run() error
}

var (
	year int
	day  int
)

var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "Advent of Code runner",
	Long:  "Run Advent of Code solutions by year and day, e.g. `go run cmd/aoc -y 2022 -d 4`",
	RunE: func(cmd *cobra.Command, args []string) error {
		if year < 2015 || year > time.Now().Year() {
			return fmt.Errorf("invalid --year: must be between 2015 and today's year")
		}
		if day < 1 || day > 25 {
			return fmt.Errorf("invalid --day: must be between 1 and 25")
		}
		return dispatch(year, day)
	},
}

func init() {
	rootCmd.Flags().IntVarP(&year, "year", "y", 0, "Year of the challenge (e.g. 2022)")
	rootCmd.Flags().IntVarP(&day, "day", "d", 0, "Day of the challenge (1-25)")
	rootCmd.MarkFlagRequired("year")
	rootCmd.MarkFlagRequired("day")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}

func dispatch(year int, day int) error {
	if days, ok := registry[year]; ok {
		if run, ok := days[day]; ok {
			return run()
		}
	}
	return fmt.Errorf("solution not found for y=%d d=%d", year, day)
}
