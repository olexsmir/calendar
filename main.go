package main

import (
	"context"
	"log/slog"
	"os"

	"codeberg.org/cream-cal/calendar/internal/cli"
)

func main() {
	if err := cli.New().Run(context.Background(), os.Args); err != nil {
		slog.Error("cal", "err", err)
		os.Exit(1)
	}
}
