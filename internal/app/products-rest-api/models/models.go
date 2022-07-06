package models

import "fmt"

type SKU string

type Name string

type Category string

type Currency string

type Price int64

type DiscountValue int8

type Percentage string

type Product struct {
	SKU      SKU      `json:"sku"`
	Name     Name     `json:"name"`
	Category Category `json:"category"`
	Price    Price    `json:"price"`
}

type PriceWithDiscount struct {
	Original           Price      `json:"original"`
	Final              Price      `json:"final"`
	DiscountPercentage Percentage `json:"discount_percentage"`
	Currency           Currency   `json:"currency"`
}

type DiscountedProduct struct {
	SKU      SKU               `json:"sku"`
	Name     Name              `json:"name"`
	Category Category          `json:"category"`
	Price    PriceWithDiscount `json:"price"`
}

type Discount struct {
	SKU      SKU
	Category Category
	Value    DiscountValue
}

const (
	CurrencyEUR = "EUR"
)

var (
	ErrDiscountOutOfBounds = fmt.Errorf("discount out of bounds")
)

func (p *Product) ApplyDiscount(dv DiscountValue) (*DiscountedProduct, error) {
	if dv < 0 || dv > 100 {
		return nil, fmt.Errorf("bad discount: %w: %d", ErrDiscountOutOfBounds, dv)
	}

	finalPrice := Price(int(p.Price) * (100 - int(dv)) / 100)
	return &DiscountedProduct{
		SKU:      p.SKU,
		Category: p.Category,
		Name:     p.Name,
		Price: PriceWithDiscount{
			Original:           p.Price,
			Final:              finalPrice,
			DiscountPercentage: Percentage(fmt.Sprintf("%d%%", dv)),
			Currency:           CurrencyEUR,
		},
	}, nil
}
