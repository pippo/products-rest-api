package models

type SKU string

type Name string

type Category string

type Currency string

type Price int64

type DiscountValue int8

type Percentage string

type Product struct {
	SKU      SKU
	Name     Name
	Category Category
	Price    Price
}

type PriceWithDiscount struct {
	Original           Price
	Final              Price
	DiscountPercentage Percentage
	Currency           Currency
}

type DiscountedProduct struct {
	SKU      SKU
	Name     Name
	Category Category
	Price    PriceWithDiscount
}

type Discount struct {
	SKU      SKU
	Category Category
	Value    DiscountValue
}

const (
	CurrencyEUR = "EUR"
)
