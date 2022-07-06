package services

import (
	"testing"

	"github.com/pippo/products-rest-api/internal/app/products-rest-api/models"
	"github.com/pippo/products-rest-api/internal/app/products-rest-api/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type productTestStage struct {
	t               *testing.T
	discountStorage storage.DiscountStorage
	discountService *DiscountService
	productStorage  storage.ProductStorage
	productService  *ProductService
	result          []models.DiscountedProduct
	product         models.DiscountedProduct
}

func NewTestStage(t *testing.T) (*productTestStage, *productTestStage, *productTestStage) {
	stage := &productTestStage{
		t:               t,
		discountStorage: storage.NewInMemoryDiscountStorage(),
		productStorage:  storage.NewInMemoryProductStorage(),
	}
	stage.discountService = NewDiscountService(stage.discountStorage)
	stage.productService = NewProductService(stage.productStorage, stage.discountService)
	return stage, stage, stage
}

func (s *productTestStage) and() *productTestStage {
	return s
}

func (s *productTestStage) an_existing_product(product models.Product) {
	err := s.productStorage.Add(product)
	require.NoError(s.t, err)
}

func (s *productTestStage) an_existing_discount(discount models.Discount) *productTestStage {
	err := s.discountStorage.Add(discount)
	require.NoError(s.t, err)
	return s
}

func (s *productTestStage) a_list_of_products(products []models.Product) {
	for _, p := range products {
		err := s.productStorage.Add(p)
		require.NoError(s.t, err)
	}
}

func (s *productTestStage) list_of_products_is_retrieved() {
	s.list_of_products_is_retrieved_with_category_filter(models.Category(""))
}

func (s *productTestStage) list_of_products_is_retrieved_with_category_filter(cat models.Category) {
	var err error
	s.result, err = s.productService.ListProducts(ProductFilterCriteria{Category: cat})
	require.NoError(s.t, err)
}

func (s *productTestStage) list_of_products_is_retrieved_with_price_filter(maxPrice models.Price) {
	var err error
	s.result, err = s.productService.ListProducts(ProductFilterCriteria{MaxPrice: maxPrice})
	require.NoError(s.t, err)
}

func (s *productTestStage) the_result_should_be_of_length(n int) *productTestStage {
	require.NotEmpty(s.t, s.result)
	assert.Len(s.t, s.result, n)
	return s
}

func (s *productTestStage) the_result_should_contain_product(sku models.SKU) *productTestStage {
	for _, p := range s.result {
		if p.SKU == sku {
			s.product = p
			return s
		}
	}
	assert.Fail(s.t, "product with SKU=%s not found", sku)
	return s
}

func (s *productTestStage) the_product_name_should_be(name models.Name) *productTestStage {
	require.NotEmpty(s.t, s.product)
	assert.Equal(s.t, s.product.Name, name)
	return s
}

func (s *productTestStage) the_product_category_should_be(cat models.Category) *productTestStage {
	require.NotEmpty(s.t, s.product)
	assert.Equal(s.t, s.product.Category, cat)
	return s
}

func (s *productTestStage) the_product_original_price_should_be(price models.Price) *productTestStage {
	require.NotEmpty(s.t, s.product)
	assert.Equal(s.t, s.product.Price.Original, price)
	return s
}

func (s *productTestStage) the_product_final_price_should_be(price models.Price) *productTestStage {
	require.NotEmpty(s.t, s.product)
	assert.Equal(s.t, s.product.Price.Final, price)
	return s
}

func (s *productTestStage) the_product_price_currency_should_be(cur models.Currency) *productTestStage {
	require.NotEmpty(s.t, s.product)
	assert.Equal(s.t, s.product.Price.Currency, cur)
	return s
}

func (s *productTestStage) the_product_discount_should_be(discount models.Percentage) *productTestStage {
	require.NotEmpty(s.t, s.product)
	assert.Equal(s.t, s.product.Price.DiscountPercentage, discount)
	return s
}
