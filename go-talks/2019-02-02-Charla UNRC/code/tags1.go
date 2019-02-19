package main

type Rate struct {
	ID            string      `json:"id" validate:"required,UUIDv4"`
	CurrencyBase  Currency    `json:"currency_base" validate:"required,min=3,max=3,currency_base"`
	CurrencyQuote Currency    `json:"currency_quote" validate:"required,min=3,max=3,currency"`
	Rate          float64     `json:"rate" validate:"required,gt=0"`
	InvRate       float64     `json:"inv_rate" validate:"-"`
	CreationDate  ISO8601Time `json:"creation_date" validate:"-"`
	HMAC          string      `json:"hmac" validate:"-"`
}
