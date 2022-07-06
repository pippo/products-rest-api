package services

import (
	"testing"

	"github.com/pippo/products-rest-api/internal/app/products-rest-api/models"
	"github.com/pippo/products-rest-api/internal/app/products-rest-api/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type discountTestStage struct {
	t       *testing.T
	storage storage.DiscountStorage
	service *DiscountService
	result  models.DiscountValue
}

func NewDiscountTestStage(t *testing.T) (*discountTestStage, *discountTestStage, *discountTestStage) {
	stage := &discountTestStage{
		t:       t,
		storage: storage.NewInMemoryDiscountStorage(),
	}
	stage.service = NewDiscountService(stage.storage)
	return stage, stage, stage
}

func (s *discountTestStage) and() *discountTestStage {
	return s
}

func (s *discountTestStage) an_existing_discount(discount models.Discount) *discountTestStage {
	err := s.storage.Add(discount)
	require.NoError(s.t, err)
	return s
}

func (s *discountTestStage) a_discount_lookup_made_for(p models.Product) {
	s.result = s.service.LookupDiscount(p)
}

func (s *discountTestStage) the_resulting_discount_percentage_should_be(v models.DiscountValue) {
	assert.Equal(s.t, v, s.result)
}
