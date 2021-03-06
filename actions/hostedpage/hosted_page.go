package hostedpage

import (
	"encoding/json"
	"fmt"
	"github.com/chargebee/chargebee-go"
	"github.com/chargebee/chargebee-go/models/hostedpage"
)

func CheckoutNew(params *hostedpage.CheckoutNewRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/checkout_new"), params)
}

func CheckoutExisting(params *hostedpage.CheckoutExistingRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/checkout_existing"), params)
}

func UpdateCard(params *hostedpage.UpdateCardRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/update_card"), params)
}

func UpdatePaymentMethod(params *hostedpage.UpdatePaymentMethodRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/update_payment_method"), params)
}

func ManagePaymentSources(params *hostedpage.ManagePaymentSourcesRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/manage_payment_sources"), params)
}

func CollectNow(params *hostedpage.CollectNowRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/collect_now"), params)
}

func ExtendSubscription(params *hostedpage.ExtendSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/extend_subscription"), params)
}

func Acknowledge(id string) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/hosted_pages/%v/acknowledge", id), nil)
}

func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/hosted_pages/%v", id), nil)
}

func List(params *hostedpage.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/hosted_pages"), params)
}
func Content(page hostedpage.HostedPage) *chargebee.Result {
	content := &chargebee.Result{}
	err1 := json.Unmarshal(page.Content, content)
	if err1 != nil {
		panic(err1)
	}
	return content
}
