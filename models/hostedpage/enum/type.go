package enum

type Type string

const (
	TypeCheckoutNew          Type = "checkout_new"
	TypeCheckoutExisting     Type = "checkout_existing"
	TypeUpdateCard           Type = "update_card"
	TypeUpdatePaymentMethod  Type = "update_payment_method"
	TypeManagePaymentSources Type = "manage_payment_sources"
	TypeCollectNow           Type = "collect_now"
	TypeExtendSubscription   Type = "extend_subscription"
)
