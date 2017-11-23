package postgresql_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/nickmro/world-app/world/db"
)

var _ = Describe("CountryTx", func() {
	var pq *DB
	var tx *Tx

	BeforeEach(func() {
		pq = MustOpenDB()
		tx = pq.MustBegin()
	})

	AfterEach(func() {
		tx.Rollback()
		pq.Close()
	})

	Describe("GetCountryByCode", func() {
		Context("when the country exists", func() {
			It("returns a country", func() {
				country, err := tx.GetCountryByCode("USA")
				Expect(err).ToNot(HaveOccurred())
				Expect(country).ToNot(BeNil())
			})
		})

		Context("when the country does not exist", func() {
			It("returns a country not found error", func() {
				_, err := tx.GetCountryByCode("A")
				Expect(err).To(Equal(db.ErrCountryNotFound))
			})
		})
	})

})
