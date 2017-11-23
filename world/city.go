package world

type City struct {
	Id         int64  `db:"id" json:"id"`
	Name       string `db:"name" json:"name"`
	Code       string `db:"code" json:"code"`
	District   string `db:"district" json:"district"`
	Population int64  `db:"population" json:"population"`
}

type CityService interface {

}