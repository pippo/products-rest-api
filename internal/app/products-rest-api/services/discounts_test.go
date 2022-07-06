package services

import (
	"testing"

	"github.com/pippo/products-rest-api/internal/app/products-rest-api/models"
)

func Test_LookupDiscount_BySKU(t *testing.T) {
	given, when, then := NewDiscountTestStage(t)

	given.
		an_existing_discount(models.Discount{
			SKU:   "000001",
			Value: 10,
		})

	when.
		a_discount_lookup_made_for(models.Product{
			SKU:      "000001",
			Name:     "BV Lean leather ankle boots",
			Category: models.Category("boots"),
			Price:    models.Price(89000),
		})

	then.
		the_resulting_discount_percentage_should_be(models.DiscountValue(10))
}

func Test_LookupDiscount_ByCategory(t *testing.T) {
	given, when, then := NewDiscountTestStage(t)

	given.
		an_existing_discount(models.Discount{
			Category: "boots",
			Value:    10,
		})

	when.
		a_discount_lookup_made_for(models.Product{
			SKU:      "000001",
			Name:     "BV Lean leather ankle boots",
			Category: models.Category("boots"),
			Price:    models.Price(89000),
		})

	then.
		the_resulting_discount_percentage_should_be(models.DiscountValue(10))
}

func Test_LookupDiscount_None(t *testing.T) {
	given, when, then := NewDiscountTestStage(t)

	given.
		an_existing_discount(models.Discount{
			SKU:   "000002",
			Value: 10,
		})

	when.
		a_discount_lookup_made_for(models.Product{
			SKU:      "000001",
			Name:     "BV Lean leather ankle boots",
			Category: models.Category("boots"),
			Price:    models.Price(89000),
		})

	then.
		the_resulting_discount_percentage_should_be(models.DiscountValue(0))
}

func Test_LookupDiscount_Multiple(t *testing.T) {
	given, when, then := NewDiscountTestStage(t)

	given.
		an_existing_discount(models.Discount{SKU: "000001", Value: 20}).and().
		an_existing_discount(models.Discount{Category: "boots", Value: 10})

	when.
		a_discount_lookup_made_for(models.Product{
			SKU:      "000001",
			Name:     "BV Lean leather ankle boots",
			Category: models.Category("boots"),
			Price:    models.Price(89000),
		})

	then.
		the_resulting_discount_percentage_should_be(models.DiscountValue(20))
}
