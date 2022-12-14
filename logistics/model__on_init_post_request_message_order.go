/*
ONDC Protocol API for logistics

ONDC Protocol Core API specification. This is an adaptation of Beckn Core version 0.9.3

API version: 1.0.19
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// OnInitPostRequestMessageOrder struct for OnInitPostRequestMessageOrder
type OnInitPostRequestMessageOrder struct {
	Provider *OnInitPostRequestMessageOrderProvider `json:"provider,omitempty"`
	ProviderLocation *OnInitPostRequestMessageOrderProviderLocation `json:"provider_location,omitempty"`
	Items []OnInitPostRequestMessageOrderItemsInner `json:"items,omitempty"`
	AddOns []OnInitPostRequestMessageOrderAddOnsInner `json:"add_ons,omitempty"`
	Offers []OnInitPostRequestMessageOrderOffersInner `json:"offers,omitempty"`
	Billing *Billing `json:"billing,omitempty"`
	Fulfillment *Fulfillment `json:"fulfillment,omitempty"`
	Quote *Quotation `json:"quote,omitempty"`
	Payment *Payment `json:"payment,omitempty"`
}

// NewOnInitPostRequestMessageOrder instantiates a new OnInitPostRequestMessageOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnInitPostRequestMessageOrder() *OnInitPostRequestMessageOrder {
	this := OnInitPostRequestMessageOrder{}
	return &this
}

// NewOnInitPostRequestMessageOrderWithDefaults instantiates a new OnInitPostRequestMessageOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnInitPostRequestMessageOrderWithDefaults() *OnInitPostRequestMessageOrder {
	this := OnInitPostRequestMessageOrder{}
	return &this
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *OnInitPostRequestMessageOrder) GetProvider() OnInitPostRequestMessageOrderProvider {
	if o == nil || isNil(o.Provider) {
		var ret OnInitPostRequestMessageOrderProvider
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnInitPostRequestMessageOrder) GetProviderOk() (*OnInitPostRequestMessageOrderProvider, bool) {
	if o == nil || isNil(o.Provider) {
    return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *OnInitPostRequestMessageOrder) HasProvider() bool {
	if o != nil && !isNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given OnInitPostRequestMessageOrderProvider and assigns it to the Provider field.
func (o *OnInitPostRequestMessageOrder) SetProvider(v OnInitPostRequestMessageOrderProvider) {
	o.Provider = &v
}

// GetProviderLocation returns the ProviderLocation field value if set, zero value otherwise.
func (o *OnInitPostRequestMessageOrder) GetProviderLocation() OnInitPostRequestMessageOrderProviderLocation {
	if o == nil || isNil(o.ProviderLocation) {
		var ret OnInitPostRequestMessageOrderProviderLocation
		return ret
	}
	return *o.ProviderLocation
}

// GetProviderLocationOk returns a tuple with the ProviderLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnInitPostRequestMessageOrder) GetProviderLocationOk() (*OnInitPostRequestMessageOrderProviderLocation, bool) {
	if o == nil || isNil(o.ProviderLocation) {
    return nil, false
	}
	return o.ProviderLocation, true
}

// HasProviderLocation returns a boolean if a field has been set.
func (o *OnInitPostRequestMessageOrder) HasProviderLocation() bool {
	if o != nil && !isNil(o.ProviderLocation) {
		return true
	}

	return false
}

// SetProviderLocation gets a reference to the given OnInitPostRequestMessageOrderProviderLocation and assigns it to the ProviderLocation field.
func (o *OnInitPostRequestMessageOrder) SetProviderLocation(v OnInitPostRequestMessageOrderProviderLocation) {
	o.ProviderLocation = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *OnInitPostRequestMessageOrder) GetItems() []OnInitPostRequestMessageOrderItemsInner {
	if o == nil || isNil(o.Items) {
		var ret []OnInitPostRequestMessageOrderItemsInner
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnInitPostRequestMessageOrder) GetItemsOk() ([]OnInitPostRequestMessageOrderItemsInner, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *OnInitPostRequestMessageOrder) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []OnInitPostRequestMessageOrderItemsInner and assigns it to the Items field.
func (o *OnInitPostRequestMessageOrder) SetItems(v []OnInitPostRequestMessageOrderItemsInner) {
	o.Items = v
}

// GetAddOns returns the AddOns field value if set, zero value otherwise.
func (o *OnInitPostRequestMessageOrder) GetAddOns() []OnInitPostRequestMessageOrderAddOnsInner {
	if o == nil || isNil(o.AddOns) {
		var ret []OnInitPostRequestMessageOrderAddOnsInner
		return ret
	}
	return o.AddOns
}

// GetAddOnsOk returns a tuple with the AddOns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnInitPostRequestMessageOrder) GetAddOnsOk() ([]OnInitPostRequestMessageOrderAddOnsInner, bool) {
	if o == nil || isNil(o.AddOns) {
    return nil, false
	}
	return o.AddOns, true
}

// HasAddOns returns a boolean if a field has been set.
func (o *OnInitPostRequestMessageOrder) HasAddOns() bool {
	if o != nil && !isNil(o.AddOns) {
		return true
	}

	return false
}

// SetAddOns gets a reference to the given []OnInitPostRequestMessageOrderAddOnsInner and assigns it to the AddOns field.
func (o *OnInitPostRequestMessageOrder) SetAddOns(v []OnInitPostRequestMessageOrderAddOnsInner) {
	o.AddOns = v
}

// GetOffers returns the Offers field value if set, zero value otherwise.
func (o *OnInitPostRequestMessageOrder) GetOffers() []OnInitPostRequestMessageOrderOffersInner {
	if o == nil || isNil(o.Offers) {
		var ret []OnInitPostRequestMessageOrderOffersInner
		return ret
	}
	return o.Offers
}

// GetOffersOk returns a tuple with the Offers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnInitPostRequestMessageOrder) GetOffersOk() ([]OnInitPostRequestMessageOrderOffersInner, bool) {
	if o == nil || isNil(o.Offers) {
    return nil, false
	}
	return o.Offers, true
}

// HasOffers returns a boolean if a field has been set.
func (o *OnInitPostRequestMessageOrder) HasOffers() bool {
	if o != nil && !isNil(o.Offers) {
		return true
	}

	return false
}

// SetOffers gets a reference to the given []OnInitPostRequestMessageOrderOffersInner and assigns it to the Offers field.
func (o *OnInitPostRequestMessageOrder) SetOffers(v []OnInitPostRequestMessageOrderOffersInner) {
	o.Offers = v
}

// GetBilling returns the Billing field value if set, zero value otherwise.
func (o *OnInitPostRequestMessageOrder) GetBilling() Billing {
	if o == nil || isNil(o.Billing) {
		var ret Billing
		return ret
	}
	return *o.Billing
}

// GetBillingOk returns a tuple with the Billing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnInitPostRequestMessageOrder) GetBillingOk() (*Billing, bool) {
	if o == nil || isNil(o.Billing) {
    return nil, false
	}
	return o.Billing, true
}

// HasBilling returns a boolean if a field has been set.
func (o *OnInitPostRequestMessageOrder) HasBilling() bool {
	if o != nil && !isNil(o.Billing) {
		return true
	}

	return false
}

// SetBilling gets a reference to the given Billing and assigns it to the Billing field.
func (o *OnInitPostRequestMessageOrder) SetBilling(v Billing) {
	o.Billing = &v
}

// GetFulfillment returns the Fulfillment field value if set, zero value otherwise.
func (o *OnInitPostRequestMessageOrder) GetFulfillment() Fulfillment {
	if o == nil || isNil(o.Fulfillment) {
		var ret Fulfillment
		return ret
	}
	return *o.Fulfillment
}

// GetFulfillmentOk returns a tuple with the Fulfillment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnInitPostRequestMessageOrder) GetFulfillmentOk() (*Fulfillment, bool) {
	if o == nil || isNil(o.Fulfillment) {
    return nil, false
	}
	return o.Fulfillment, true
}

// HasFulfillment returns a boolean if a field has been set.
func (o *OnInitPostRequestMessageOrder) HasFulfillment() bool {
	if o != nil && !isNil(o.Fulfillment) {
		return true
	}

	return false
}

// SetFulfillment gets a reference to the given Fulfillment and assigns it to the Fulfillment field.
func (o *OnInitPostRequestMessageOrder) SetFulfillment(v Fulfillment) {
	o.Fulfillment = &v
}

// GetQuote returns the Quote field value if set, zero value otherwise.
func (o *OnInitPostRequestMessageOrder) GetQuote() Quotation {
	if o == nil || isNil(o.Quote) {
		var ret Quotation
		return ret
	}
	return *o.Quote
}

// GetQuoteOk returns a tuple with the Quote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnInitPostRequestMessageOrder) GetQuoteOk() (*Quotation, bool) {
	if o == nil || isNil(o.Quote) {
    return nil, false
	}
	return o.Quote, true
}

// HasQuote returns a boolean if a field has been set.
func (o *OnInitPostRequestMessageOrder) HasQuote() bool {
	if o != nil && !isNil(o.Quote) {
		return true
	}

	return false
}

// SetQuote gets a reference to the given Quotation and assigns it to the Quote field.
func (o *OnInitPostRequestMessageOrder) SetQuote(v Quotation) {
	o.Quote = &v
}

// GetPayment returns the Payment field value if set, zero value otherwise.
func (o *OnInitPostRequestMessageOrder) GetPayment() Payment {
	if o == nil || isNil(o.Payment) {
		var ret Payment
		return ret
	}
	return *o.Payment
}

// GetPaymentOk returns a tuple with the Payment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnInitPostRequestMessageOrder) GetPaymentOk() (*Payment, bool) {
	if o == nil || isNil(o.Payment) {
    return nil, false
	}
	return o.Payment, true
}

// HasPayment returns a boolean if a field has been set.
func (o *OnInitPostRequestMessageOrder) HasPayment() bool {
	if o != nil && !isNil(o.Payment) {
		return true
	}

	return false
}

// SetPayment gets a reference to the given Payment and assigns it to the Payment field.
func (o *OnInitPostRequestMessageOrder) SetPayment(v Payment) {
	o.Payment = &v
}

func (o OnInitPostRequestMessageOrder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !isNil(o.ProviderLocation) {
		toSerialize["provider_location"] = o.ProviderLocation
	}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.AddOns) {
		toSerialize["add_ons"] = o.AddOns
	}
	if !isNil(o.Offers) {
		toSerialize["offers"] = o.Offers
	}
	if !isNil(o.Billing) {
		toSerialize["billing"] = o.Billing
	}
	if !isNil(o.Fulfillment) {
		toSerialize["fulfillment"] = o.Fulfillment
	}
	if !isNil(o.Quote) {
		toSerialize["quote"] = o.Quote
	}
	if !isNil(o.Payment) {
		toSerialize["payment"] = o.Payment
	}
	return json.Marshal(toSerialize)
}

type NullableOnInitPostRequestMessageOrder struct {
	value *OnInitPostRequestMessageOrder
	isSet bool
}

func (v NullableOnInitPostRequestMessageOrder) Get() *OnInitPostRequestMessageOrder {
	return v.value
}

func (v *NullableOnInitPostRequestMessageOrder) Set(val *OnInitPostRequestMessageOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableOnInitPostRequestMessageOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableOnInitPostRequestMessageOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnInitPostRequestMessageOrder(val *OnInitPostRequestMessageOrder) *NullableOnInitPostRequestMessageOrder {
	return &NullableOnInitPostRequestMessageOrder{value: val, isSet: true}
}

func (v NullableOnInitPostRequestMessageOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnInitPostRequestMessageOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


