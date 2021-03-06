package subscription

import (
	"encoding/json"
	"github.com/chargebee/chargebee-go/enum"
	"github.com/chargebee/chargebee-go/filter"
	subscriptionEnum "github.com/chargebee/chargebee-go/models/subscription/enum"
)

type Subscription struct {
	Id                     string                             `json:"id"`
	CustomerId             string                             `json:"customer_id"`
	CurrencyCode           string                             `json:"currency_code"`
	PlanId                 string                             `json:"plan_id"`
	PlanQuantity           int32                              `json:"plan_quantity"`
	PlanUnitPrice          int32                              `json:"plan_unit_price"`
	SetupFee               int32                              `json:"setup_fee"`
	BillingPeriod          int32                              `json:"billing_period"`
	BillingPeriodUnit      subscriptionEnum.BillingPeriodUnit `json:"billing_period_unit"`
	PlanFreeQuantity       int32                              `json:"plan_free_quantity"`
	Status                 subscriptionEnum.Status            `json:"status"`
	StartDate              int64                              `json:"start_date"`
	TrialStart             int64                              `json:"trial_start"`
	TrialEnd               int64                              `json:"trial_end"`
	CurrentTermStart       int64                              `json:"current_term_start"`
	CurrentTermEnd         int64                              `json:"current_term_end"`
	NextBillingAt          int64                              `json:"next_billing_at"`
	RemainingBillingCycles int32                              `json:"remaining_billing_cycles"`
	PoNumber               string                             `json:"po_number"`
	CreatedAt              int64                              `json:"created_at"`
	StartedAt              int64                              `json:"started_at"`
	ActivatedAt            int64                              `json:"activated_at"`
	PauseDate              int64                              `json:"pause_date"`
	ResumeDate             int64                              `json:"resume_date"`
	CancelledAt            int64                              `json:"cancelled_at"`
	CancelReason           subscriptionEnum.CancelReason      `json:"cancel_reason"`
	AffiliateToken         string                             `json:"affiliate_token"`
	CreatedFromIp          string                             `json:"created_from_ip"`
	ResourceVersion        int64                              `json:"resource_version"`
	UpdatedAt              int64                              `json:"updated_at"`
	HasScheduledChanges    bool                               `json:"has_scheduled_changes"`
	PaymentSourceId        string                             `json:"payment_source_id"`
	AutoCollection         enum.AutoCollection                `json:"auto_collection"`
	DueInvoicesCount       int32                              `json:"due_invoices_count"`
	DueSince               int64                              `json:"due_since"`
	TotalDues              int32                              `json:"total_dues"`
	Mrr                    int32                              `json:"mrr"`
	ExchangeRate           float64                            `json:"exchange_rate"`
	BaseCurrencyCode       string                             `json:"base_currency_code"`
	Addons                 []*Addon                           `json:"addons"`
	Coupon                 string                             `json:"coupon"`
	Coupons                []*Coupon                          `json:"coupons"`
	ShippingAddress        *ShippingAddress                   `json:"shipping_address"`
	ReferralInfo           *ReferralInfo                      `json:"referral_info"`
	InvoiceNotes           string                             `json:"invoice_notes"`
	MetaData               json.RawMessage                    `json:"meta_data"`
	Deleted                bool                               `json:"deleted"`
	CustomField            map[string]interface{}             `json:"custom_field"`
	Object                 string                             `json:"object"`
}
type Addon struct {
	Id                     string `json:"id"`
	Quantity               int32  `json:"quantity"`
	UnitPrice              int32  `json:"unit_price"`
	TrialEnd               int64  `json:"trial_end"`
	RemainingBillingCycles int32  `json:"remaining_billing_cycles"`
	Object                 string `json:"object"`
}
type Coupon struct {
	CouponId     string `json:"coupon_id"`
	ApplyTill    int64  `json:"apply_till"`
	AppliedCount int32  `json:"applied_count"`
	CouponCode   string `json:"coupon_code"`
	Object       string `json:"object"`
}
type ShippingAddress struct {
	FirstName        string                `json:"first_name"`
	LastName         string                `json:"last_name"`
	Email            string                `json:"email"`
	Company          string                `json:"company"`
	Phone            string                `json:"phone"`
	Line1            string                `json:"line1"`
	Line2            string                `json:"line2"`
	Line3            string                `json:"line3"`
	City             string                `json:"city"`
	StateCode        string                `json:"state_code"`
	State            string                `json:"state"`
	Country          string                `json:"country"`
	Zip              string                `json:"zip"`
	ValidationStatus enum.ValidationStatus `json:"validation_status"`
	Object           string                `json:"object"`
}
type ReferralInfo struct {
	ReferralCode              string                                    `json:"referral_code"`
	CouponCode                string                                    `json:"coupon_code"`
	ReferrerId                string                                    `json:"referrer_id"`
	ExternalReferenceId       string                                    `json:"external_reference_id"`
	RewardStatus              subscriptionEnum.ReferralInfoRewardStatus `json:"reward_status"`
	ReferralSystem            enum.ReferralSystem                       `json:"referral_system"`
	AccountId                 string                                    `json:"account_id"`
	CampaignId                string                                    `json:"campaign_id"`
	ExternalCampaignId        string                                    `json:"external_campaign_id"`
	FriendOfferType           enum.FriendOfferType                      `json:"friend_offer_type"`
	ReferrerRewardType        enum.ReferrerRewardType                   `json:"referrer_reward_type"`
	NotifyReferralSystem      enum.NotifyReferralSystem                 `json:"notify_referral_system"`
	DestinationUrl            string                                    `json:"destination_url"`
	PostPurchaseWidgetEnabled bool                                      `json:"post_purchase_widget_enabled"`
	Object                    string                                    `json:"object"`
}
type CreateRequestParams struct {
	Id                   string                       `json:"id,omitempty"`
	Customer             *CreateCustomerParams        `json:"customer,omitempty"`
	PlanId               string                       `json:"plan_id"`
	PlanQuantity         *int32                       `json:"plan_quantity,omitempty"`
	PlanUnitPrice        *int32                       `json:"plan_unit_price,omitempty"`
	SetupFee             *int32                       `json:"setup_fee,omitempty"`
	StartDate            *int64                       `json:"start_date,omitempty"`
	TrialEnd             *int64                       `json:"trial_end,omitempty"`
	BillingCycles        *int32                       `json:"billing_cycles,omitempty"`
	Addons               []*CreateAddonParams         `json:"addons,omitempty"`
	Coupon               string                       `json:"coupon,omitempty"`
	AutoCollection       enum.AutoCollection          `json:"auto_collection,omitempty"`
	TermsToCharge        *int32                       `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode enum.BillingAlignmentMode    `json:"billing_alignment_mode,omitempty"`
	PoNumber             string                       `json:"po_number,omitempty"`
	CouponIds            []string                     `json:"coupon_ids,omitempty"`
	Card                 *CreateCardParams            `json:"card,omitempty"`
	PaymentMethod        *CreatePaymentMethodParams   `json:"payment_method,omitempty"`
	BillingAddress       *CreateBillingAddressParams  `json:"billing_address,omitempty"`
	ShippingAddress      *CreateShippingAddressParams `json:"shipping_address,omitempty"`
	AffiliateToken       string                       `json:"affiliate_token,omitempty"`
	CreatedFromIp        string                       `json:"created_from_ip,omitempty"`
	InvoiceNotes         string                       `json:"invoice_notes,omitempty"`
	MetaData             map[string]interface{}       `json:"meta_data,omitempty"`
	InvoiceImmediately   *bool                        `json:"invoice_immediately,omitempty"`
}
type CreateCustomerParams struct {
	Id                    string              `json:"id,omitempty"`
	Email                 string              `json:"email,omitempty"`
	FirstName             string              `json:"first_name,omitempty"`
	LastName              string              `json:"last_name,omitempty"`
	Company               string              `json:"company,omitempty"`
	Taxability            enum.Taxability     `json:"taxability,omitempty"`
	Locale                string              `json:"locale,omitempty"`
	EntityCode            enum.EntityCode     `json:"entity_code,omitempty"`
	ExemptNumber          string              `json:"exempt_number,omitempty"`
	NetTermDays           *int32              `json:"net_term_days,omitempty"`
	Phone                 string              `json:"phone,omitempty"`
	AutoCollection        enum.AutoCollection `json:"auto_collection,omitempty"`
	AllowDirectDebit      *bool               `json:"allow_direct_debit,omitempty"`
	ConsolidatedInvoicing *bool               `json:"consolidated_invoicing,omitempty"`
	VatNumber             string              `json:"vat_number,omitempty"`
	RegisteredForGst      *bool               `json:"registered_for_gst,omitempty"`
}
type CreateAddonParams struct {
	Id        string `json:"id,omitempty"`
	Quantity  *int32 `json:"quantity,omitempty"`
	UnitPrice *int32 `json:"unit_price,omitempty"`
	TrialEnd  *int64 `json:"trial_end,omitempty"`
}
type CreateCardParams struct {
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
	TmpToken         string       `json:"tmp_token,omitempty"`
	FirstName        string       `json:"first_name,omitempty"`
	LastName         string       `json:"last_name,omitempty"`
	Number           string       `json:"number,omitempty"`
	ExpiryMonth      *int32       `json:"expiry_month,omitempty"`
	ExpiryYear       *int32       `json:"expiry_year,omitempty"`
	Cvv              string       `json:"cvv,omitempty"`
	BillingAddr1     string       `json:"billing_addr1,omitempty"`
	BillingAddr2     string       `json:"billing_addr2,omitempty"`
	BillingCity      string       `json:"billing_city,omitempty"`
	BillingStateCode string       `json:"billing_state_code,omitempty"`
	BillingState     string       `json:"billing_state,omitempty"`
	BillingZip       string       `json:"billing_zip,omitempty"`
	BillingCountry   string       `json:"billing_country,omitempty"`
	IpAddress        string       `json:"ip_address,omitempty"`
}
type CreatePaymentMethodParams struct {
	Type             enum.Type    `json:"type,omitempty"`
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
	ReferenceId      string       `json:"reference_id,omitempty"`
	TmpToken         string       `json:"tmp_token,omitempty"`
	IssuingCountry   string       `json:"issuing_country,omitempty"`
}
type CreateBillingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type CreateShippingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}

type CreateForCustomerRequestParams struct {
	Id                   string                                  `json:"id,omitempty"`
	PlanId               string                                  `json:"plan_id"`
	PlanQuantity         *int32                                  `json:"plan_quantity,omitempty"`
	PlanUnitPrice        *int32                                  `json:"plan_unit_price,omitempty"`
	SetupFee             *int32                                  `json:"setup_fee,omitempty"`
	StartDate            *int64                                  `json:"start_date,omitempty"`
	TrialEnd             *int64                                  `json:"trial_end,omitempty"`
	BillingCycles        *int32                                  `json:"billing_cycles,omitempty"`
	Addons               []*CreateForCustomerAddonParams         `json:"addons,omitempty"`
	Coupon               string                                  `json:"coupon,omitempty"`
	AutoCollection       enum.AutoCollection                     `json:"auto_collection,omitempty"`
	TermsToCharge        *int32                                  `json:"terms_to_charge,omitempty"`
	BillingAlignmentMode enum.BillingAlignmentMode               `json:"billing_alignment_mode,omitempty"`
	PoNumber             string                                  `json:"po_number,omitempty"`
	CouponIds            []string                                `json:"coupon_ids,omitempty"`
	PaymentSourceId      string                                  `json:"payment_source_id,omitempty"`
	ShippingAddress      *CreateForCustomerShippingAddressParams `json:"shipping_address,omitempty"`
	InvoiceNotes         string                                  `json:"invoice_notes,omitempty"`
	MetaData             map[string]interface{}                  `json:"meta_data,omitempty"`
	InvoiceImmediately   *bool                                   `json:"invoice_immediately,omitempty"`
}
type CreateForCustomerAddonParams struct {
	Id        string `json:"id,omitempty"`
	Quantity  *int32 `json:"quantity,omitempty"`
	UnitPrice *int32 `json:"unit_price,omitempty"`
	TrialEnd  *int64 `json:"trial_end,omitempty"`
}
type CreateForCustomerShippingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}

type ListRequestParams struct {
	Limit                  *int32                  `json:"limit,omitempty"`
	Offset                 string                  `json:"offset,omitempty"`
	IncludeDeleted         *bool                   `json:"include_deleted,omitempty"`
	Id                     *filter.StringFilter    `json:"id,omitempty"`
	CustomerId             *filter.StringFilter    `json:"customer_id,omitempty"`
	PlanId                 *filter.StringFilter    `json:"plan_id,omitempty"`
	Status                 *filter.EnumFilter      `json:"status,omitempty"`
	CancelReason           *filter.EnumFilter      `json:"cancel_reason,omitempty"`
	RemainingBillingCycles *filter.NumberFilter    `json:"remaining_billing_cycles,omitempty"`
	CreatedAt              *filter.TimestampFilter `json:"created_at,omitempty"`
	ActivatedAt            *filter.TimestampFilter `json:"activated_at,omitempty"`
	NextBillingAt          *filter.TimestampFilter `json:"next_billing_at,omitempty"`
	CancelledAt            *filter.TimestampFilter `json:"cancelled_at,omitempty"`
	HasScheduledChanges    *filter.BooleanFilter   `json:"has_scheduled_changes,omitempty"`
	UpdatedAt              *filter.TimestampFilter `json:"updated_at,omitempty"`
	SortBy                 *filter.SortFilter      `json:"sort_by,omitempty"`
}

type SubscriptionsForCustomerRequestParams struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset string `json:"offset,omitempty"`
}

type RemoveScheduledCancellationRequestParams struct {
	BillingCycles *int32 `json:"billing_cycles,omitempty"`
}

type RemoveCouponsRequestParams struct {
	CouponIds []string `json:"coupon_ids,omitempty"`
}

type UpdateRequestParams struct {
	PlanId               string                       `json:"plan_id,omitempty"`
	PlanQuantity         *int32                       `json:"plan_quantity,omitempty"`
	PlanUnitPrice        *int32                       `json:"plan_unit_price,omitempty"`
	SetupFee             *int32                       `json:"setup_fee,omitempty"`
	StartDate            *int64                       `json:"start_date,omitempty"`
	TrialEnd             *int64                       `json:"trial_end,omitempty"`
	BillingCycles        *int32                       `json:"billing_cycles,omitempty"`
	Addons               []*UpdateAddonParams         `json:"addons,omitempty"`
	ReplaceAddonList     *bool                        `json:"replace_addon_list,omitempty"`
	Coupon               string                       `json:"coupon,omitempty"`
	TermsToCharge        *int32                       `json:"terms_to_charge,omitempty"`
	ReactivateFrom       *int64                       `json:"reactivate_from,omitempty"`
	BillingAlignmentMode enum.BillingAlignmentMode    `json:"billing_alignment_mode,omitempty"`
	PoNumber             string                       `json:"po_number,omitempty"`
	CouponIds            []string                     `json:"coupon_ids,omitempty"`
	ReplaceCouponList    *bool                        `json:"replace_coupon_list,omitempty"`
	Prorate              *bool                        `json:"prorate,omitempty"`
	EndOfTerm            *bool                        `json:"end_of_term,omitempty"`
	ForceTermReset       *bool                        `json:"force_term_reset,omitempty"`
	Reactivate           *bool                        `json:"reactivate,omitempty"`
	Card                 *UpdateCardParams            `json:"card,omitempty"`
	PaymentMethod        *UpdatePaymentMethodParams   `json:"payment_method,omitempty"`
	BillingAddress       *UpdateBillingAddressParams  `json:"billing_address,omitempty"`
	ShippingAddress      *UpdateShippingAddressParams `json:"shipping_address,omitempty"`
	Customer             *UpdateCustomerParams        `json:"customer,omitempty"`
	InvoiceNotes         string                       `json:"invoice_notes,omitempty"`
	MetaData             map[string]interface{}       `json:"meta_data,omitempty"`
	InvoiceImmediately   *bool                        `json:"invoice_immediately,omitempty"`
}
type UpdateAddonParams struct {
	Id        string `json:"id,omitempty"`
	Quantity  *int32 `json:"quantity,omitempty"`
	UnitPrice *int32 `json:"unit_price,omitempty"`
	TrialEnd  *int64 `json:"trial_end,omitempty"`
}
type UpdateCardParams struct {
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
	TmpToken         string       `json:"tmp_token,omitempty"`
	FirstName        string       `json:"first_name,omitempty"`
	LastName         string       `json:"last_name,omitempty"`
	Number           string       `json:"number,omitempty"`
	ExpiryMonth      *int32       `json:"expiry_month,omitempty"`
	ExpiryYear       *int32       `json:"expiry_year,omitempty"`
	Cvv              string       `json:"cvv,omitempty"`
	BillingAddr1     string       `json:"billing_addr1,omitempty"`
	BillingAddr2     string       `json:"billing_addr2,omitempty"`
	BillingCity      string       `json:"billing_city,omitempty"`
	BillingStateCode string       `json:"billing_state_code,omitempty"`
	BillingState     string       `json:"billing_state,omitempty"`
	BillingZip       string       `json:"billing_zip,omitempty"`
	BillingCountry   string       `json:"billing_country,omitempty"`
	IpAddress        string       `json:"ip_address,omitempty"`
}
type UpdatePaymentMethodParams struct {
	Type             enum.Type    `json:"type,omitempty"`
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
	ReferenceId      string       `json:"reference_id,omitempty"`
	TmpToken         string       `json:"tmp_token,omitempty"`
	IssuingCountry   string       `json:"issuing_country,omitempty"`
}
type UpdateBillingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type UpdateShippingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type UpdateCustomerParams struct {
	VatNumber        string `json:"vat_number,omitempty"`
	RegisteredForGst *bool  `json:"registered_for_gst,omitempty"`
}

type ChangeTermEndRequestParams struct {
	TermEndsAt         *int64 `json:"term_ends_at"`
	Prorate            *bool  `json:"prorate,omitempty"`
	InvoiceImmediately *bool  `json:"invoice_immediately,omitempty"`
}

type CancelRequestParams struct {
	EndOfTerm                         *bool                                  `json:"end_of_term,omitempty"`
	CreditOptionForCurrentTermCharges enum.CreditOptionForCurrentTermCharges `json:"credit_option_for_current_term_charges,omitempty"`
	UnbilledChargesOption             enum.UnbilledChargesOption             `json:"unbilled_charges_option,omitempty"`
	AccountReceivablesHandling        enum.AccountReceivablesHandling        `json:"account_receivables_handling,omitempty"`
	RefundableCreditsHandling         enum.RefundableCreditsHandling         `json:"refundable_credits_handling,omitempty"`
}

type ReactivateRequestParams struct {
	TrialEnd             *int64                    `json:"trial_end,omitempty"`
	BillingCycles        *int32                    `json:"billing_cycles,omitempty"`
	TrialPeriodDays      *int32                    `json:"trial_period_days,omitempty"`
	ReactivateFrom       *int64                    `json:"reactivate_from,omitempty"`
	InvoiceImmediately   *bool                     `json:"invoice_immediately,omitempty"`
	BillingAlignmentMode enum.BillingAlignmentMode `json:"billing_alignment_mode,omitempty"`
	TermsToCharge        *int32                    `json:"terms_to_charge,omitempty"`
}

type AddChargeAtTermEndRequestParams struct {
	Amount      *int32 `json:"amount"`
	Description string `json:"description"`
}

type ChargeAddonAtTermEndRequestParams struct {
	AddonId        string `json:"addon_id"`
	AddonQuantity  *int32 `json:"addon_quantity,omitempty"`
	AddonUnitPrice *int32 `json:"addon_unit_price,omitempty"`
}

type ChargeFutureRenewalsRequestParams struct {
	TermsToCharge *int32 `json:"terms_to_charge,omitempty"`
}

type ImportSubscriptionRequestParams struct {
	Id                       string                                   `json:"id,omitempty"`
	Customer                 *ImportSubscriptionCustomerParams        `json:"customer,omitempty"`
	PlanId                   string                                   `json:"plan_id"`
	PlanQuantity             *int32                                   `json:"plan_quantity,omitempty"`
	PlanUnitPrice            *int32                                   `json:"plan_unit_price,omitempty"`
	SetupFee                 *int32                                   `json:"setup_fee,omitempty"`
	StartDate                *int64                                   `json:"start_date,omitempty"`
	TrialEnd                 *int64                                   `json:"trial_end,omitempty"`
	BillingCycles            *int32                                   `json:"billing_cycles,omitempty"`
	Addons                   []*ImportSubscriptionAddonParams         `json:"addons,omitempty"`
	AutoCollection           enum.AutoCollection                      `json:"auto_collection,omitempty"`
	PoNumber                 string                                   `json:"po_number,omitempty"`
	CouponIds                []string                                 `json:"coupon_ids,omitempty"`
	Status                   subscriptionEnum.Status                  `json:"status"`
	CurrentTermEnd           *int64                                   `json:"current_term_end,omitempty"`
	CurrentTermStart         *int64                                   `json:"current_term_start,omitempty"`
	TrialStart               *int64                                   `json:"trial_start,omitempty"`
	CancelledAt              *int64                                   `json:"cancelled_at,omitempty"`
	StartedAt                *int64                                   `json:"started_at,omitempty"`
	CreateCurrentTermInvoice *bool                                    `json:"create_current_term_invoice,omitempty"`
	Card                     *ImportSubscriptionCardParams            `json:"card,omitempty"`
	PaymentMethod            *ImportSubscriptionPaymentMethodParams   `json:"payment_method,omitempty"`
	BillingAddress           *ImportSubscriptionBillingAddressParams  `json:"billing_address,omitempty"`
	ShippingAddress          *ImportSubscriptionShippingAddressParams `json:"shipping_address,omitempty"`
	AffiliateToken           string                                   `json:"affiliate_token,omitempty"`
	InvoiceNotes             string                                   `json:"invoice_notes,omitempty"`
	MetaData                 map[string]interface{}                   `json:"meta_data,omitempty"`
	Transaction              *ImportSubscriptionTransactionParams     `json:"transaction,omitempty"`
}
type ImportSubscriptionCustomerParams struct {
	Id               string              `json:"id,omitempty"`
	Email            string              `json:"email,omitempty"`
	FirstName        string              `json:"first_name,omitempty"`
	LastName         string              `json:"last_name,omitempty"`
	Company          string              `json:"company,omitempty"`
	Taxability       enum.Taxability     `json:"taxability,omitempty"`
	Locale           string              `json:"locale,omitempty"`
	EntityCode       enum.EntityCode     `json:"entity_code,omitempty"`
	ExemptNumber     string              `json:"exempt_number,omitempty"`
	NetTermDays      *int32              `json:"net_term_days,omitempty"`
	Phone            string              `json:"phone,omitempty"`
	AutoCollection   enum.AutoCollection `json:"auto_collection,omitempty"`
	AllowDirectDebit *bool               `json:"allow_direct_debit,omitempty"`
	VatNumber        string              `json:"vat_number,omitempty"`
}
type ImportSubscriptionAddonParams struct {
	Id        string `json:"id,omitempty"`
	Quantity  *int32 `json:"quantity,omitempty"`
	UnitPrice *int32 `json:"unit_price,omitempty"`
}
type ImportSubscriptionCardParams struct {
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
	TmpToken         string       `json:"tmp_token,omitempty"`
	FirstName        string       `json:"first_name,omitempty"`
	LastName         string       `json:"last_name,omitempty"`
	Number           string       `json:"number,omitempty"`
	ExpiryMonth      *int32       `json:"expiry_month,omitempty"`
	ExpiryYear       *int32       `json:"expiry_year,omitempty"`
	Cvv              string       `json:"cvv,omitempty"`
	BillingAddr1     string       `json:"billing_addr1,omitempty"`
	BillingAddr2     string       `json:"billing_addr2,omitempty"`
	BillingCity      string       `json:"billing_city,omitempty"`
	BillingStateCode string       `json:"billing_state_code,omitempty"`
	BillingState     string       `json:"billing_state,omitempty"`
	BillingZip       string       `json:"billing_zip,omitempty"`
	BillingCountry   string       `json:"billing_country,omitempty"`
}
type ImportSubscriptionPaymentMethodParams struct {
	Type             enum.Type    `json:"type,omitempty"`
	Gateway          enum.Gateway `json:"gateway,omitempty"`
	GatewayAccountId string       `json:"gateway_account_id,omitempty"`
	ReferenceId      string       `json:"reference_id,omitempty"`
	IssuingCountry   string       `json:"issuing_country,omitempty"`
}
type ImportSubscriptionBillingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type ImportSubscriptionShippingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}
type ImportSubscriptionTransactionParams struct {
	Amount          *int32             `json:"amount,omitempty"`
	PaymentMethod   enum.PaymentMethod `json:"payment_method,omitempty"`
	ReferenceNumber string             `json:"reference_number,omitempty"`
	Date            *int64             `json:"date,omitempty"`
}

type ImportForCustomerRequestParams struct {
	Id                       string                                  `json:"id,omitempty"`
	PlanId                   string                                  `json:"plan_id"`
	PlanQuantity             *int32                                  `json:"plan_quantity,omitempty"`
	PlanUnitPrice            *int32                                  `json:"plan_unit_price,omitempty"`
	SetupFee                 *int32                                  `json:"setup_fee,omitempty"`
	StartDate                *int64                                  `json:"start_date,omitempty"`
	TrialEnd                 *int64                                  `json:"trial_end,omitempty"`
	BillingCycles            *int32                                  `json:"billing_cycles,omitempty"`
	Addons                   []*ImportForCustomerAddonParams         `json:"addons,omitempty"`
	AutoCollection           enum.AutoCollection                     `json:"auto_collection,omitempty"`
	PoNumber                 string                                  `json:"po_number,omitempty"`
	CouponIds                []string                                `json:"coupon_ids,omitempty"`
	PaymentSourceId          string                                  `json:"payment_source_id,omitempty"`
	Status                   subscriptionEnum.Status                 `json:"status"`
	CurrentTermEnd           *int64                                  `json:"current_term_end,omitempty"`
	CurrentTermStart         *int64                                  `json:"current_term_start,omitempty"`
	TrialStart               *int64                                  `json:"trial_start,omitempty"`
	CancelledAt              *int64                                  `json:"cancelled_at,omitempty"`
	StartedAt                *int64                                  `json:"started_at,omitempty"`
	CreateCurrentTermInvoice *bool                                   `json:"create_current_term_invoice,omitempty"`
	Transaction              *ImportForCustomerTransactionParams     `json:"transaction,omitempty"`
	ShippingAddress          *ImportForCustomerShippingAddressParams `json:"shipping_address,omitempty"`
	InvoiceNotes             string                                  `json:"invoice_notes,omitempty"`
	MetaData                 map[string]interface{}                  `json:"meta_data,omitempty"`
}
type ImportForCustomerAddonParams struct {
	Id        string `json:"id,omitempty"`
	Quantity  *int32 `json:"quantity,omitempty"`
	UnitPrice *int32 `json:"unit_price,omitempty"`
}
type ImportForCustomerTransactionParams struct {
	Amount          *int32             `json:"amount,omitempty"`
	PaymentMethod   enum.PaymentMethod `json:"payment_method,omitempty"`
	ReferenceNumber string             `json:"reference_number,omitempty"`
	Date            *int64             `json:"date,omitempty"`
}
type ImportForCustomerShippingAddressParams struct {
	FirstName        string                `json:"first_name,omitempty"`
	LastName         string                `json:"last_name,omitempty"`
	Email            string                `json:"email,omitempty"`
	Company          string                `json:"company,omitempty"`
	Phone            string                `json:"phone,omitempty"`
	Line1            string                `json:"line1,omitempty"`
	Line2            string                `json:"line2,omitempty"`
	Line3            string                `json:"line3,omitempty"`
	City             string                `json:"city,omitempty"`
	StateCode        string                `json:"state_code,omitempty"`
	State            string                `json:"state,omitempty"`
	Zip              string                `json:"zip,omitempty"`
	Country          string                `json:"country,omitempty"`
	ValidationStatus enum.ValidationStatus `json:"validation_status,omitempty"`
}

type OverrideBillingProfileRequestParams struct {
	PaymentSourceId string              `json:"payment_source_id,omitempty"`
	AutoCollection  enum.AutoCollection `json:"auto_collection,omitempty"`
}

type PauseRequestParams struct {
	PauseOption             enum.PauseOption             `json:"pause_option,omitempty"`
	PauseDate               *int64                       `json:"pause_date,omitempty"`
	UnbilledChargesHandling enum.UnbilledChargesHandling `json:"unbilled_charges_handling,omitempty"`
	InvoiceDunningHandling  enum.InvoiceDunningHandling  `json:"invoice_dunning_handling,omitempty"`
	ResumeDate              *int64                       `json:"resume_date,omitempty"`
}

type ResumeRequestParams struct {
	ResumeOption           enum.ResumeOption           `json:"resume_option,omitempty"`
	ResumeDate             *int64                      `json:"resume_date,omitempty"`
	UnpaidInvoicesHandling enum.UnpaidInvoicesHandling `json:"unpaid_invoices_handling,omitempty"`
}
