package cmd

import (
	"github.com/urfave/cli"
	"github.com/nickmro/world-app/world/db/postgresql"
	"github.com/nickmro/world-app/world/app"
	"github.com/nickmro/world-app/world"
	"github.com/drinkin/di/env"
)

// Cli is a command line interface for the world app.
type Cli struct {
	*cli.App
	DB *postgresql.DB
	CountryService world.CountryService
	CityService world.CityService
}

// NewCli returns a new Cli.
func NewCli() *Cli {
	db := postgresql.NewDB(env.MustGet("DB_URL"))

	countryService := &app.CountryService{
		DB: db,
	}

	c := &Cli{
		App: cli.NewApp(),
		DB: db,
		CountryService: countryService,
	}

	getCountryCMD := cli.Command{
		Name: "get",
		Subcommands: []cli.Command{
			{
				Name: "country",
				Action: c.GetCountry,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name: "code",
						Usage: "code of country to get",
					},
					cli.StringFlag{
						Name: "name",
						Usage: "name of the country to get",
					},
				},
			},
		},
	}

	c.App.Commands = []cli.Command{getCountryCMD}

	return c
}
