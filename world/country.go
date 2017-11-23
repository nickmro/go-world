package world

type Country struct {
	Code             string   `db:"code" json:"code"`
	Name             string   `db:"name" json:"name"`
	Continent        string   `db:"continent" json:"continent"`
	Region           string   `db:"region" json:"region"`
	SurfaceArea      float64  `db:"surfacearea" json:"-"`
	IndependenceYear *int16   `db:"indepyear" json:"independence_year"`
	Population       int32    `db:"population" json:"population"`
	LifeExpectancy   *float64 `db:"lifeexpectancy" json:"life_expectancy"`
	GNP              *float64 `db:"gnp" json:"gnp"`
	OldGNP           *float64 `db:"gnpold" json:"gnpold"`
	LocalName        string   `db:"localname" json:"local_name"`
	GovernmentForm   string   `db:"governmentform" json:"government_form"`
	HeadOfState      *string  `db:"headofstate" json:"head_of_state"`
	Capital          int64    `db:"capital" json:"capital"`
	Code2            string   `db:"code2" json:"code2"`
}

type CountryService interface {
	GetCountryByCode(string) (*Country, error)
}
