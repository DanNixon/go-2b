package main

import (
	"fmt"
	"os"

	"github.com/DanNixon/go-2b/pkg/powerbox"
	"github.com/DanNixon/go-2b/pkg/types"
	"github.com/urfave/cli/v2"
)

func openPowerbox(c *cli.Context) powerbox.Powerbox {
	pb, err := powerbox.NewRestPowerbox(c.String("address"))
	if err != nil {
		fmt.Println("Failed to open remote powerbox:", err)
		os.Exit(2)
	}
	return pb
}

func main() {
	app := &cli.App{
		Usage: "Control a 2B powerbox over REST API",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "address",
				Value: "http://localhost:8080",
				Usage: "Address of powerbox server",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "get",
				Aliases: []string{"g"},
				Usage:   "Report powerbox status",
				Action: func(c *cli.Context) error {
					pb := openPowerbox(c)
					if s, err := pb.Get(); err != nil {
						return err
					} else {
						s.PrettyPrint()
					}
					return nil
				},
			},
			{
				Name:  "set",
				Usage: "Set powerbox output",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "mode",
						Aliases: []string{"m"},
						Value:   "",
						Usage:   "Mode to select",
					},
					&cli.StringFlag{
						Name:  "power",
						Value: "low",
						Usage: "Power level to select",
					},
					&cli.IntFlag{
						Name:  "a",
						Value: 0,
						Usage: "Channel A output level",
					},
					&cli.IntFlag{
						Name:  "b",
						Value: 0,
						Usage: "Channel B output level",
					},
					&cli.IntFlag{
						Name:  "c",
						Value: 50,
						Usage: "Parameter C value",
					},
					&cli.IntFlag{
						Name:  "d",
						Value: 50,
						Usage: "Parameter D value",
					},
				},
				Action: func(c *cli.Context) error {
					newSettings := types.Settings{
						Mode:       types.Mode(c.String("mode")),
						PowerLevel: types.PowerLevel(c.String("power")),
						Channels: types.Channels{
							A:      types.NumericSetting(c.Int("a")),
							B:      types.NumericSetting(c.Int("b")),
							Linked: false,
						},
						Parameters: types.Parameters{
							C: types.NumericSetting(c.Int("c")),
							D: types.NumericSetting(c.Int("d")),
						},
					}
					if err := newSettings.Validate(); err != nil {
						return err
					}

					pb := openPowerbox(c)

					// Apply new setings
					if s, err := pb.Set(newSettings); err != nil {
						return err
					} else {
						s.PrettyPrint()
					}
					return nil
				},
			},
			{
				Name:    "kill",
				Aliases: []string{"k"},
				Usage:   "Disable powerbox output",
				Action: func(c *cli.Context) error {
					pb := openPowerbox(c)
					if s, err := pb.Kill(); err != nil {
						return err
					} else {
						s.PrettyPrint()
					}
					return nil
				},
			},
			{
				Name:    "reset",
				Aliases: []string{"r"},
				Usage:   "Reset powerbox",
				Action: func(c *cli.Context) error {
					pb := openPowerbox(c)
					if s, err := pb.Reset(); err != nil {
						return err
					} else {
						s.PrettyPrint()
					}
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
