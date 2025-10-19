package aoc

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	year      int
	day       int
	part      int
	inputFile string
)

var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "Advent of Code runner",
	Long:  "Run Advent of Code solutions by year and day, e.g. `go run cmd/aoc -y 2022 -d 4 -p 1 -i ./inputs/2022-04-1.txt`",
	RunE: func(cmd *cobra.Command, args []string) error {
		if year < 2015 || year > time.Now().Year() {
			return fmt.Errorf("invalid --year: must be between 2015 and today's year")
		}
		if day < 1 || day > 25 {
			return fmt.Errorf("invalid --day: must be between 1 and 25")
		}
		if part != 1 && part != 2 {
			return fmt.Errorf("invalid --part: must be 1 or 2")
		}
		b, err := os.ReadFile(inputFile)
		if err != nil {
			return fmt.Errorf("reading input file: %w", err)
		}
		return dispatch(year, day, part, string(b))
	},
}

func init() {
	rootCmd.Flags().IntVarP(&year, "year", "y", 0, "Year of the challenge (e.g. 2022)")
	rootCmd.Flags().IntVarP(&day, "day", "d", 0, "Day of the challenge (1-25)")
	rootCmd.Flags().IntVarP(&part, "part", "p", 0, "Challenge part (1 or 2)")
	rootCmd.Flags().StringVarP(&inputFile, "input-file", "i", "", "File that contains the challenge's input data")
	_ = rootCmd.MarkFlagRequired("year")
	_ = rootCmd.MarkFlagRequired("day")
	_ = rootCmd.MarkFlagRequired("part")
	_ = rootCmd.MarkFlagRequired("input-file")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}

func dispatch(year int, day int, part int, input string) error {
	days, ok := registry[year]
	if !ok {
		return solutionNotFoundError(year, day, part)
	}

	parts, ok := days[day]
	if !ok {
		return solutionNotFoundError(year, day, part)
	}

	runnerFunc, ok := parts[part]
	if !ok {
		return solutionNotFoundError(year, day, part)
	}

	start := time.Now()
	out, err := runnerFunc(input)
	elapsed := time.Since(start)

	if err != nil {
		return err
	}

	slog.Info("Runner finished successfully",
		"year", year,
		"day", day,
		"part", part,
		"duration", humanDuration(elapsed),
	)
	slog.Info("Result=" + out)

	return nil
}

func solutionNotFoundError(year int, day int, part int) error {
	return fmt.Errorf("solution not found for y=%d d=%d p=%d", year, day, part)
}

// humanDuration renders a time.Duration in a friendly, precise unit.
func humanDuration(d time.Duration) string {
	if d < time.Microsecond {
		return fmt.Sprintf("%d nanoseconds", d.Nanoseconds())
	}
	if d < time.Millisecond {
		us := d.Microseconds()
		if us == 1 {
			return "1 microsecond"
		}
		return fmt.Sprintf("%d microseconds", us)
	}
	if d < time.Second {
		ms := d.Milliseconds()
		if ms == 1 {
			return "1 millisecond"
		}
		return fmt.Sprintf("%d milliseconds", ms)
	}

	secs := float64(d) / float64(time.Second)
	return fmt.Sprintf("%.3f seconds", secs)
}
