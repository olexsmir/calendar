package cli

import (
	"context"

	"github.com/urfave/cli/v3"

	"codeberg.org/cream-cal/calendar/internal/version"
)

type Cli struct{}

func New() *Cli {
	return &Cli{}
}

func (c *Cli) Run(ctx context.Context, args []string) error {
	cmd := &cli.Command{
		Name:                  version.Name,
		Usage:                 "the best calendar app you'll ever use",
		Version:               version.Version,
		EnableShellCompletion: true,
		Commands: []*cli.Command{
			{
				Name:   "today",
				Usage:  "list events for today",
				Action: c.todayAction,
			},
		},
	}
	return cmd.Run(ctx, args)
}
