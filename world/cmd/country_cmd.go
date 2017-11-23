package cmd

import (
	"github.com/urfave/cli"
	"fmt"
	"github.com/nickmro/world-app/world"
)

func (cli *Cli) GetCountry(c *cli.Context) {
	err := cli.DB.Open()
	if err != nil {
		fmt.Println(err)
	}
	defer cli.DB.Close()

	code := c.String("code")
	name := c.String("name")

	if code != "" {
		cli.getCountryByCode(code)
	}

	if name != "" {
		cli.getCountryByName(name)
	}

}

func (cli *Cli) getCountryByCode(code string) {
	country, err := cli.CountryService.GetCountryByCode(code)
	if err != nil {
		fmt.Println(err)
	}
	printCountry(country)
}

func (cli *Cli) getCountryByName(name string) {

}

func printCountry(country *world.Country) {
	fmt.Println("Country: ", country.Name)
	fmt.Println("Region: ", country.Region)
	fmt.Println("Population: ", country.Population)
}
