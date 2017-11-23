package app_test

import (
	app "github.com/nickmro/world-app/world/app"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nickmro/world-app/world/db/mock"
	"github.com/nickmro/world-app/world"
)

type CountryService struct {
	*app.CountryService
	db *mock.DB
	tx *mock.Tx
}

func NewCountryService() *CountryService {
	svc := &CountryService{
		CountryService: &app.CountryService{},
		db: mock.NewDB(),
		tx: mock.NewTx(),
	}
	svc.db.Tx = svc.tx
	svc.CountryService.DB = svc.db
	return svc
}

var _ = Describe("CountryService", func() {
	var svc *CountryService

	BeforeEach(func() {
		svc = NewCountryService()
	})

	Describe("GetCountryByCode", func() {
		It("returns the country", func() {
			country, err := svc.GetCountryByCode("USA")
			Expect(country).ToNot(BeNil())
			Expect(err).ToNot(HaveOccurred())
		})

		Context("when getting the country from the database fails", func() {
			It("returns an error", func() {
				svc.db.Tx.GetCountryByCodeFn = func(string) (*world.Country, error) {
					return nil, app.ErrInternal
				}
				_, err := svc.GetCountryByCode("USA")
				Expect(err).To(Equal(app.ErrInternal))
			})
		})
	})

})
