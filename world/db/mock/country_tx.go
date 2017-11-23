package mock

import "github.com/nickmro/world-app/world"

func (tx *Tx) GetCountryByCode(code string) (*world.Country, error) {
	return tx.GetCountryByCodeFn(code)
}
