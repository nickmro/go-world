package app

import (
	"github.com/nickmro/world-app/world"
	"github.com/nickmro/world-app/world/db"
)

type CountryService struct {
	DB db.DB
}

func (svc *CountryService) GetCountryByCode(code string) (*world.Country, error) {
	tx, err := svc.DB.Begin()
	if err != nil {
		return nil, ErrInternal
	}

	country, err := tx.GetCountryByCode(code)
	if err != nil {
		return nil, err
	}

	tx.Commit()

	return country, nil
}
