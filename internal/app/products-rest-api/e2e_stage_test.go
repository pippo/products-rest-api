package test

import (
	"testing"

	"github.com/pippo/products-rest-api/internal/app/products-rest-api/models"
	"github.com/pippo/products-rest-api/internal/app/products-rest-api/services"
	"github.com/pippo/products-rest-api/internal/app/products-rest-api/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testStage struct {
	t       *testing.T
	storage storage.ProductStorage
	service *services.ProductService
	result  []models.DiscountedProduct
	product models.DiscountedProduct
}

func NewTestStage(t *testing.T) (*testStage, *testStage, *testStage) {
	stage := testStage{
		t:       t,
		storage: storage.NewInMemoryProductStorage(),
	}
	stage.service = services.NewProductService(stage.storage)
	return &stage, &stage, &stage
}

func (s *testStage) and() *testStage {
	return s
}

func (s *testStage) a_product(product models.Product) {
	err := s.storage.Add(product)
	require.NoError(s.t, err)
}

func (s *testStage) a_list_of_products(products []models.Product) {
	for _, p := range products {
		err := s.storage.Add(p)
		require.NoError(s.t, err)
	}
}

func (s *testStage) list_of_products_is_retrieved() {
	s.list_of_products_is_retrieved_with_category_filter(models.Category(""))
}

func (s *testStage) list_of_products_is_retrieved_with_category_filter(cat models.Category) {
	var err error
	s.result, err = s.service.ListProducts(services.ProductFilterCriteria{Category: cat})
	require.NoError(s.t, err)
}

func (s *testStage) list_of_products_is_retrieved_with_price_filter(maxPrice models.Price) {
	var err error
	s.result, err = s.service.ListProducts(services.ProductFilterCriteria{MaxPrice: maxPrice})
	require.NoError(s.t, err)
}

func (s *testStage) the_result_should_be_of_length(n int) *testStage {
	require.NotEmpty(s.t, s.result)
	assert.Len(s.t, s.result, n)
	return s
}

func (s *testStage) the_result_should_contain_product(sku models.SKU) *testStage {
	for _, p := range s.result {
		if p.SKU == sku {
			s.product = p
			return s
		}
	}
	assert.Fail(s.t, "product with SKU=%s not found", sku)
	return s
}

func (s *testStage) the_product_name_should_be(name models.Name) *testStage {
	require.NotEmpty(s.t, s.product)
	assert.Equal(s.t, s.product.Name, name)
	return s
}

func (s *testStage) the_product_category_should_be(cat models.Category) *testStage {
	require.NotEmpty(s.t, s.product)
	assert.Equal(s.t, s.product.Category, cat)
	return s
}

func (s *testStage) the_product_original_price_should_be(price models.Price) *testStage {
	require.NotEmpty(s.t, s.product)
	assert.Equal(s.t, s.product.Price.Original, price)
	return s
}

func (s *testStage) the_product_final_price_should_be(price models.Price) *testStage {
	require.NotEmpty(s.t, s.product)
	assert.Equal(s.t, s.product.Price.Final, price)
	return s
}

func (s *testStage) the_product_price_currency_should_be(cur models.Currency) *testStage {
	require.NotEmpty(s.t, s.product)
	assert.Equal(s.t, s.product.Price.Currency, cur)
	return s
}

func (s *testStage) the_product_discount_should_be(discount models.Percentage) *testStage {
	require.NotEmpty(s.t, s.product)
	assert.Equal(s.t, s.product.Price.DiscountPercentage, discount)
	return s
}
