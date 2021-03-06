package paymentsource

import (
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
	paymentSourceEnum "github.com/chargebee/chargebee-go/models/paymentsource/enum"
)

type PaymentSource struct {
	Id               string                   `json:"id"`
	CustomerId       string                   `json:"customer_id"`
	Type             enum.Type                `json:"type"`
	ReferenceId      string                   `json:"reference_id"`
	Status           paymentSourceEnum.Status `json:"status"`
	Gateway          enum.Gateway             `json:"gateway"`
	GatewayAccountId string                   `json:"gateway_account_id"`
	IpAddress        string                   `json:"ip_address"`
	IssuingCountry   string                   `json:"issuing_country"`
	Card             *Card                    `json:"card"`
	BankAccount      *BankAccount             `json:"bank_account"`
	AmazonPayment    *AmazonPayment           `json:"amazon_payment"`
	Paypal           *Paypal                  `json:"paypal"`
	Deleted          bool                     `json:"deleted"`
	Object           string                   `json:"object"`
}
type Card struct {
	FirstName        string                            `json:"first_name"`
	LastName         string                            `json:"last_name"`
	Iin              string                            `json:"iin"`
	Last4            string                            `json:"last4"`
	Brand            paymentSourceEnum.CardBrand       `json:"brand"`
	FundingType      paymentSourceEnum.CardFundingType `json:"funding_type"`
	ExpiryMonth      int32                             `json:"expiry_month"`
	ExpiryYear       int32                             `json:"expiry_year"`
	BillingAddr1     string                            `json:"billing_addr1"`
	BillingAddr2     string                            `json:"billing_addr2"`
	BillingCity      string                            `json:"billing_city"`
	BillingStateCode string                            `json:"billing_state_code"`
	BillingState     string                            `json:"billing_state"`
	BillingCountry   string                            `json:"billing_country"`
	BillingZip       string                            `json:"billing_zip"`
	MaskedNumber     string                            `json:"masked_number"`
	Object           string                            `json:"object"`
}
type BankAccount struct {
	NameOnAccount string                                   `json:"name_on_account"`
	BankName      string                                   `json:"bank_name"`
	MandateId     string                                   `json:"mandate_id"`
	AccountType   paymentSourceEnum.BankAccountAccountType `json:"account_type"`
	Object        string                                   `json:"object"`
}
type AmazonPayment struct {
	Email       string `json:"email"`
	AgreementId string `json:"agreement_id"`
	Object      string `json:"object"`
}
type Paypal struct {
	Email       string `json:"email"`
	AgreementId string `json:"agreement_id"`
	Object      string `json:"object"`
}
type CreateUsingTempTokenRequestParams struct {
	CustomerId                  string    `json:"customer_id"`
	GatewayAccountId            string    `json:"gateway_account_id,omitempty"`
	Type                        enum.Type `json:"type"`
	TmpToken                    string    `json:"tmp_token"`
	IssuingCountry              string    `json:"issuing_country,omitempty"`
	ReplacePrimaryPaymentSource *bool     `json:"replace_primary_payment_source,omitempty"`
}

type CreateUsingPermanentTokenRequestParams struct {
	CustomerId                  string    `json:"customer_id"`
	Type                        enum.Type `json:"type"`
	GatewayAccountId            string    `json:"gateway_account_id,omitempty"`
	ReferenceId                 string    `json:"reference_id"`
	IssuingCountry              string    `json:"issuing_country,omitempty"`
	ReplacePrimaryPaymentSource *bool     `json:"replace_primary_payment_source,omitempty"`
}

type CreateCardRequestParams struct {
	CustomerId                  string                `json:"customer_id"`
	Card                        *CreateCardCardParams `json:"card,omitempty"`
	ReplacePrimaryPaymentSource *bool                 `json:"replace_primary_payment_source,omitempty"`
}
type CreateCardCardParams struct {
	GatewayAccountId string `json:"gateway_account_id,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	Number           string `json:"number"`
	ExpiryMonth      *int32 `json:"expiry_month"`
	ExpiryYear       *int32 `json:"expiry_year"`
	Cvv              string `json:"cvv,omitempty"`
	BillingAddr1     string `json:"billing_addr1,omitempty"`
	BillingAddr2     string `json:"billing_addr2,omitempty"`
	BillingCity      string `json:"billing_city,omitempty"`
	BillingStateCode string `json:"billing_state_code,omitempty"`
	BillingState     string `json:"billing_state,omitempty"`
	BillingZip       string `json:"billing_zip,omitempty"`
	BillingCountry   string `json:"billing_country,omitempty"`
}

type UpdateCardRequestParams struct {
	Card            *UpdateCardCardParams  `json:"card,omitempty"`
	GatewayMetaData map[string]interface{} `json:"gateway_meta_data,omitempty"`
}
type UpdateCardCardParams struct {
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	ExpiryMonth      *int32 `json:"expiry_month,omitempty"`
	ExpiryYear       *int32 `json:"expiry_year,omitempty"`
	BillingAddr1     string `json:"billing_addr1,omitempty"`
	BillingAddr2     string `json:"billing_addr2,omitempty"`
	BillingCity      string `json:"billing_city,omitempty"`
	BillingZip       string `json:"billing_zip,omitempty"`
	BillingStateCode string `json:"billing_state_code,omitempty"`
	BillingState     string `json:"billing_state,omitempty"`
	BillingCountry   string `json:"billing_country,omitempty"`
}

type ListRequestParams struct {
	Limit      *int32               `json:"limit,omitempty"`
	Offset     string               `json:"offset,omitempty"`
	CustomerId *filter.StringFilter `json:"customer_id,omitempty"`
	Type       *filter.EnumFilter   `json:"type,omitempty"`
	Status     *filter.EnumFilter   `json:"status,omitempty"`
}

type SwitchGatewayAccountRequestParams struct {
	GatewayAccountId string `json:"gateway_account_id"`
}

type ExportPaymentSourceRequestParams struct {
	GatewayAccountId string `json:"gateway_account_id"`
}
