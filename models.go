package main

// SnapRequest data model
type SnapRequest struct {
	TransactionDetails TransactionDetails `json:"transaction_details"`
	CreditCard         CreditCard         `json:"credit_card"`
	ItemDetails        []ItemDetails      `json:"item_details"`
	CustomerDetails    CustomerDetails    `json:"customer_details"`
	Expiry             Expiry             `json:"expiry"`
}

// TransactionDetails data model
type TransactionDetails struct {
	OrderID     string `json:"order_id"`
	GrossAmount int    `json:"gross_amount"`
}

// CreditCard data model
type CreditCard struct {
	Secure        bool        `json:"secure,omitempty"`
	Channel       string      `json:"channel,omitempty"`
	Bank          string      `json:"bank,omitempty"`
	Installment   Installment `json:"installment,omitempty"`
	WhitelistBins []string    `json:"whitelist_bins,omitempty"`
}

// Installment data model
type Installment struct {
	Required bool  `json:"required,omitempty"`
	Terms    Terms `json:"terms,omitempty"`
}

// Terms for Installment data model
type Terms struct {
	BNI     []int `json:"bni,omitempty"`
	Mandiri []int `json:"mandiri,omitempty"`
	CIMB    []int `json:"cimb,omitempty"`
	BCA     []int `json:"bca,omitempty"`
	Offline []int `json:"offline,omitempty"`
}

// ItemDetails data model
type ItemDetails struct {
	ID       string `json:"id"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
	Name     string `json:"name"`
}

// CustomerDetails data model
type CustomerDetails struct {
	FirstName       string  `json:"first_name"`
	LastName        string  `json:"last_name"`
	Email           string  `json:"email"`
	Phone           string  `json:"phone"`
	BillingAddress  Address `json:"billing_address"`
	ShippingAddress Address `json:"shipping_address"`
}

// Address data model
type Address struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	City        string `json:"city"`
	PostalCode  string `json:"postal_code"`
	CountryCode string `json:"country_code"`
}

// Expiry data model
type Expiry struct {
	StartTime string `json:"start_time"`
	Unit      string `json:"unit"`
	Duration  int    `json:"duration"`
}

// Card data model
type Card struct {
	UserID     string `json:"user_id,omitempty"`
	SavedToken string `json:"token_id"`
	MaskedCard string `json:"cardhash"`
	StatusCode string `json:"status_code"`
}

// JsonCard data model
type JsonCard struct {
	SavedToken string `json:"token_id"`
	MaskedCard string `json:"cardhash"`
	StatusCode string `json:"status_code"`
}
