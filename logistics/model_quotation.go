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

// Quotation Describes a quote
type Quotation struct {
	Price *Price `json:"price,omitempty"`
	Breakup []QuotationBreakupInner `json:"breakup,omitempty"`
	// Describes duration as per ISO8601 format
	Ttl *string `json:"ttl,omitempty"`
}

// NewQuotation instantiates a new Quotation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotation() *Quotation {
	this := Quotation{}
	return &this
}

// NewQuotationWithDefaults instantiates a new Quotation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotationWithDefaults() *Quotation {
	this := Quotation{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *Quotation) GetPrice() Price {
	if o == nil || isNil(o.Price) {
		var ret Price
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quotation) GetPriceOk() (*Price, bool) {
	if o == nil || isNil(o.Price) {
    return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *Quotation) HasPrice() bool {
	if o != nil && !isNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given Price and assigns it to the Price field.
func (o *Quotation) SetPrice(v Price) {
	o.Price = &v
}

// GetBreakup returns the Breakup field value if set, zero value otherwise.
func (o *Quotation) GetBreakup() []QuotationBreakupInner {
	if o == nil || isNil(o.Breakup) {
		var ret []QuotationBreakupInner
		return ret
	}
	return o.Breakup
}

// GetBreakupOk returns a tuple with the Breakup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quotation) GetBreakupOk() ([]QuotationBreakupInner, bool) {
	if o == nil || isNil(o.Breakup) {
    return nil, false
	}
	return o.Breakup, true
}

// HasBreakup returns a boolean if a field has been set.
func (o *Quotation) HasBreakup() bool {
	if o != nil && !isNil(o.Breakup) {
		return true
	}

	return false
}

// SetBreakup gets a reference to the given []QuotationBreakupInner and assigns it to the Breakup field.
func (o *Quotation) SetBreakup(v []QuotationBreakupInner) {
	o.Breakup = v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *Quotation) GetTtl() string {
	if o == nil || isNil(o.Ttl) {
		var ret string
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quotation) GetTtlOk() (*string, bool) {
	if o == nil || isNil(o.Ttl) {
    return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *Quotation) HasTtl() bool {
	if o != nil && !isNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given string and assigns it to the Ttl field.
func (o *Quotation) SetTtl(v string) {
	o.Ttl = &v
}

func (o Quotation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !isNil(o.Breakup) {
		toSerialize["breakup"] = o.Breakup
	}
	if !isNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	return json.Marshal(toSerialize)
}

type NullableQuotation struct {
	value *Quotation
	isSet bool
}

func (v NullableQuotation) Get() *Quotation {
	return v.value
}

func (v *NullableQuotation) Set(val *Quotation) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotation) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotation(val *Quotation) *NullableQuotation {
	return &NullableQuotation{value: val, isSet: true}
}

func (v NullableQuotation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

