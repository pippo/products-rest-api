package models_test

import (
	"testing"

	"github.com/pippo/products-rest-api/internal/app/products-rest-api/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestApplyDiscount_Happy(t *testing.T) {
	p := models.Product{SKU: "1", Name: "Product #1", Category: "category1", Price: 1000}
	dp, err := p.ApplyDiscount(models.DiscountValue(10))

	require.NoError(t, err)
	assert.Equal(t, models.Price(1000), dp.Price.Original)
	assert.Equal(t, models.Price(900), dp.Price.Final)
	assert.Equal(t, models.Percentage("10%"), *dp.Price.DiscountPercentage)
}

func TestApplyDiscount_Zero(t *testing.T) {
	p := models.Product{SKU: "1", Name: "Product #1", Category: "category1", Price: 1000}
	dp, err := p.ApplyDiscount(models.DiscountValue(0))

	require.NoError(t, err)
	assert.Equal(t, models.Price(1000), dp.Price.Original)
	assert.Equal(t, models.Price(1000), dp.Price.Final)
	assert.Nil(t, dp.Price.DiscountPercentage)
}

func TestApplyDiscount_Negative(t *testing.T) {
	p := models.Product{SKU: "1", Name: "Product #1", Category: "category1", Price: 1000}
	_, err := p.ApplyDiscount(models.DiscountValue(-10))

	assert.ErrorIs(t, err, models.ErrDiscountOutOfBounds)
}

func TestApplyDiscount_ToHigh(t *testing.T) {
	p := models.Product{SKU: "1", Name: "Product #1", Category: "category1", Price: 1000}
	_, err := p.ApplyDiscount(models.DiscountValue(101))

	assert.ErrorIs(t, err, models.ErrDiscountOutOfBounds)
}
