package services

import (
	"testing"

	"github.com/pippo/products-rest-api/internal/app/products-rest-api/models"
)

func Test_ListingProducts_Simple(t *testing.T) {
	given, when, then := NewTestStage(t)

	given.
		an_existing_product(models.Product{
			SKU:      "000001",
			Name:     "BV Lean leather ankle boots",
			Category: models.Category("boots"),
			Price:    models.Price(89000),
		})

	when.
		list_of_products_is_retrieved()

	then.
		the_result_should_be_of_length(1).and().
		the_result_should_contain_product("000001").and().
		the_product_category_should_be(models.Category("boots")).and().
		the_product_original_price_should_be(models.Price(89000)).and().
		the_product_final_price_should_be(models.Price(89000)).and().
		the_product_discount_should_be(models.Percentage("0%")).and().
		the_product_price_currency_should_be(models.CurrencyEUR).and().
		the_product_name_should_be(models.Name("BV Lean leather ankle boots"))
}

func Test_ListingProducts_Filtered_ByCategory(t *testing.T) {
	given, when, then := NewTestStage(t)

	given.
		a_list_of_products([]models.Product{
			{SKU: "000001", Name: "BV Lean leather ankle boots", Category: models.Category("boots"), Price: models.Price(89000)},
			{SKU: "000002", Name: "Naima embellished suede sandals", Category: models.Category("sandals"), Price: models.Price(79500)},
			{SKU: "000003", Name: "Nathane leather sneakers", Category: models.Category("snickers"), Price: models.Price(59000)},
		})

	when.
		list_of_products_is_retrieved_with_category_filter(models.Category("sandals"))

	then.
		the_result_should_be_of_length(1).and().
		the_result_should_contain_product("000002")
}

func Test_ListingProducts_Filtered_ByPrice(t *testing.T) {
	given, when, then := NewTestStage(t)

	given.
		a_list_of_products([]models.Product{
			{SKU: "000001", Name: "BV Lean leather ankle boots", Category: models.Category("boots"), Price: models.Price(89000)},
			{SKU: "000002", Name: "Naima embellished suede sandals", Category: models.Category("sandals"), Price: models.Price(79500)},
			{SKU: "000003", Name: "Nathane leather sneakers", Category: models.Category("snickers"), Price: models.Price(59000)},
		})

	when.
		list_of_products_is_retrieved_with_price_filter(models.Price(80000))

	then.
		the_result_should_be_of_length(2).and().
		the_result_should_contain_product("000002").and().
		the_result_should_contain_product("000003")
}

func Test_ListingProducts_Truncated(t *testing.T) {
	given, when, then := NewTestStage(t)

	given.
		a_list_of_products([]models.Product{
			{SKU: "000001", Name: "Product 1", Category: models.Category("boots"), Price: models.Price(10000)},
			{SKU: "000002", Name: "Product 2", Category: models.Category("boots"), Price: models.Price(10000)},
			{SKU: "000003", Name: "Product 3", Category: models.Category("boots"), Price: models.Price(10000)},
			{SKU: "000004", Name: "Product 4", Category: models.Category("boots"), Price: models.Price(10000)},
			{SKU: "000005", Name: "Product 5", Category: models.Category("boots"), Price: models.Price(10000)},
			{SKU: "000006", Name: "Product 6", Category: models.Category("boots"), Price: models.Price(10000)},
			{SKU: "000007", Name: "Product 7", Category: models.Category("boots"), Price: models.Price(10000)},
		})

	when.
		list_of_products_is_retrieved()

	then.
		the_result_should_be_of_length(5)
}

func Test_ListingProducts_WithDiscount(t *testing.T) {
	given, when, then := NewTestStage(t)

	given.
		an_existing_discount(models.Discount{SKU: "000001", Value: 10}).and().
		a_list_of_products([]models.Product{
			{SKU: "000001", Name: "BV Lean leather ankle boots", Category: models.Category("boots"), Price: models.Price(89000)},
			{SKU: "000002", Name: "Naima embellished suede sandals", Category: models.Category("sandals"), Price: models.Price(79500)},
			{SKU: "000003", Name: "Nathane leather sneakers", Category: models.Category("snickers"), Price: models.Price(59000)},
		})

	when.
		list_of_products_is_retrieved()

	then.
		the_result_should_contain_product("000001").and().
		the_product_original_price_should_be(models.Price(89000)).and().
		the_product_final_price_should_be(models.Price(80100)).and().
		the_product_discount_should_be(models.Percentage("10%"))
}
