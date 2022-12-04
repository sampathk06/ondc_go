/*
ONDC Protocol Core API

ONDC Protocol Core API specification. This is an adaptation of Beckn Core version 0.9.3

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// Order Describes the details of an order
type Order struct {
	// Hash of order object without id
	Id *string `json:"id,omitempty"`
	State *string `json:"state,omitempty"`
	Provider *OrderProvider `json:"provider,omitempty"`
	Items []OrderItemsInner `json:"items,omitempty"`
	AddOns []OrderAddOnsInner `json:"add_ons,omitempty"`
	Offers []OrderOffersInner `json:"offers,omitempty"`
	Documents []Document `json:"documents,omitempty"`
	Billing *Billing `json:"billing,omitempty"`
	Fulfillment *Fulfillment `json:"fulfillment,omitempty"`
	Quote *Quotation `json:"quote,omitempty"`
	Payment *Payment `json:"payment,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewOrder instantiates a new Order object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrder() *Order {
	this := Order{}
	return &this
}

// NewOrderWithDefaults instantiates a new Order object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderWithDefaults() *Order {
	this := Order{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Order) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Order) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Order) SetId(v string) {
	o.Id = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Order) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
    return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Order) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Order) SetState(v string) {
	o.State = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *Order) GetProvider() OrderProvider {
	if o == nil || isNil(o.Provider) {
		var ret OrderProvider
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetProviderOk() (*OrderProvider, bool) {
	if o == nil || isNil(o.Provider) {
    return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *Order) HasProvider() bool {
	if o != nil && !isNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given OrderProvider and assigns it to the Provider field.
func (o *Order) SetProvider(v OrderProvider) {
	o.Provider = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *Order) GetItems() []OrderItemsInner {
	if o == nil || isNil(o.Items) {
		var ret []OrderItemsInner
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetItemsOk() ([]OrderItemsInner, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *Order) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []OrderItemsInner and assigns it to the Items field.
func (o *Order) SetItems(v []OrderItemsInner) {
	o.Items = v
}

// GetAddOns returns the AddOns field value if set, zero value otherwise.
func (o *Order) GetAddOns() []OrderAddOnsInner {
	if o == nil || isNil(o.AddOns) {
		var ret []OrderAddOnsInner
		return ret
	}
	return o.AddOns
}

// GetAddOnsOk returns a tuple with the AddOns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetAddOnsOk() ([]OrderAddOnsInner, bool) {
	if o == nil || isNil(o.AddOns) {
    return nil, false
	}
	return o.AddOns, true
}

// HasAddOns returns a boolean if a field has been set.
func (o *Order) HasAddOns() bool {
	if o != nil && !isNil(o.AddOns) {
		return true
	}

	return false
}

// SetAddOns gets a reference to the given []OrderAddOnsInner and assigns it to the AddOns field.
func (o *Order) SetAddOns(v []OrderAddOnsInner) {
	o.AddOns = v
}

// GetOffers returns the Offers field value if set, zero value otherwise.
func (o *Order) GetOffers() []OrderOffersInner {
	if o == nil || isNil(o.Offers) {
		var ret []OrderOffersInner
		return ret
	}
	return o.Offers
}

// GetOffersOk returns a tuple with the Offers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetOffersOk() ([]OrderOffersInner, bool) {
	if o == nil || isNil(o.Offers) {
    return nil, false
	}
	return o.Offers, true
}

// HasOffers returns a boolean if a field has been set.
func (o *Order) HasOffers() bool {
	if o != nil && !isNil(o.Offers) {
		return true
	}

	return false
}

// SetOffers gets a reference to the given []OrderOffersInner and assigns it to the Offers field.
func (o *Order) SetOffers(v []OrderOffersInner) {
	o.Offers = v
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *Order) GetDocuments() []Document {
	if o == nil || isNil(o.Documents) {
		var ret []Document
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetDocumentsOk() ([]Document, bool) {
	if o == nil || isNil(o.Documents) {
    return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *Order) HasDocuments() bool {
	if o != nil && !isNil(o.Documents) {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []Document and assigns it to the Documents field.
func (o *Order) SetDocuments(v []Document) {
	o.Documents = v
}

// GetBilling returns the Billing field value if set, zero value otherwise.
func (o *Order) GetBilling() Billing {
	if o == nil || isNil(o.Billing) {
		var ret Billing
		return ret
	}
	return *o.Billing
}

// GetBillingOk returns a tuple with the Billing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetBillingOk() (*Billing, bool) {
	if o == nil || isNil(o.Billing) {
    return nil, false
	}
	return o.Billing, true
}

// HasBilling returns a boolean if a field has been set.
func (o *Order) HasBilling() bool {
	if o != nil && !isNil(o.Billing) {
		return true
	}

	return false
}

// SetBilling gets a reference to the given Billing and assigns it to the Billing field.
func (o *Order) SetBilling(v Billing) {
	o.Billing = &v
}

// GetFulfillment returns the Fulfillment field value if set, zero value otherwise.
func (o *Order) GetFulfillment() Fulfillment {
	if o == nil || isNil(o.Fulfillment) {
		var ret Fulfillment
		return ret
	}
	return *o.Fulfillment
}

// GetFulfillmentOk returns a tuple with the Fulfillment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetFulfillmentOk() (*Fulfillment, bool) {
	if o == nil || isNil(o.Fulfillment) {
    return nil, false
	}
	return o.Fulfillment, true
}

// HasFulfillment returns a boolean if a field has been set.
func (o *Order) HasFulfillment() bool {
	if o != nil && !isNil(o.Fulfillment) {
		return true
	}

	return false
}

// SetFulfillment gets a reference to the given Fulfillment and assigns it to the Fulfillment field.
func (o *Order) SetFulfillment(v Fulfillment) {
	o.Fulfillment = &v
}

// GetQuote returns the Quote field value if set, zero value otherwise.
func (o *Order) GetQuote() Quotation {
	if o == nil || isNil(o.Quote) {
		var ret Quotation
		return ret
	}
	return *o.Quote
}

// GetQuoteOk returns a tuple with the Quote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetQuoteOk() (*Quotation, bool) {
	if o == nil || isNil(o.Quote) {
    return nil, false
	}
	return o.Quote, true
}

// HasQuote returns a boolean if a field has been set.
func (o *Order) HasQuote() bool {
	if o != nil && !isNil(o.Quote) {
		return true
	}

	return false
}

// SetQuote gets a reference to the given Quotation and assigns it to the Quote field.
func (o *Order) SetQuote(v Quotation) {
	o.Quote = &v
}

// GetPayment returns the Payment field value if set, zero value otherwise.
func (o *Order) GetPayment() Payment {
	if o == nil || isNil(o.Payment) {
		var ret Payment
		return ret
	}
	return *o.Payment
}

// GetPaymentOk returns a tuple with the Payment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetPaymentOk() (*Payment, bool) {
	if o == nil || isNil(o.Payment) {
    return nil, false
	}
	return o.Payment, true
}

// HasPayment returns a boolean if a field has been set.
func (o *Order) HasPayment() bool {
	if o != nil && !isNil(o.Payment) {
		return true
	}

	return false
}

// SetPayment gets a reference to the given Payment and assigns it to the Payment field.
func (o *Order) SetPayment(v Payment) {
	o.Payment = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Order) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Order) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Order) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Order) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Order) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Order) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o Order) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.Provider) {
		toSerialize["provider"] = o.Provider
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
	if !isNil(o.Documents) {
		toSerialize["documents"] = o.Documents
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
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableOrder struct {
	value *Order
	isSet bool
}

func (v NullableOrder) Get() *Order {
	return v.value
}

func (v *NullableOrder) Set(val *Order) {
	v.value = val
	v.isSet = true
}

func (v NullableOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrder(val *Order) *NullableOrder {
	return &NullableOrder{value: val, isSet: true}
}

func (v NullableOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

