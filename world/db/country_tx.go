package db

import "github.com/nickmro/world-app/world"

type CountryTx interface {
	GetCountryByCode(string) (*world.Country, error)
}
