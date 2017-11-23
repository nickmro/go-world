package postgresql

import (
	"github.com/nickmro/world-app/world"
	"database/sql"
	"github.com/nickmro/world-app/world/db"
)

const (
	countriesTableName = "country"
)

func (tx *Tx) GetCountryByCode(code string) (*world.Country, error) {
	var country world.Country

	err := tx.Select("*").
		From(countriesTableName).
		Where("code = $1", code).
		QueryStruct(&country)

	if err == sql.ErrNoRows {
		return nil, db.ErrCountryNotFound
	}

	if err != nil {
		tx.logger.Info(err.Error())
		tx.logger.Error(err.Error())
		return nil, db.ErrInternal
	}

	return &country, nil
}
