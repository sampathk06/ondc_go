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

// Price Describes the price of an item. Allows for domain extension.
type Price struct {
	// ISO 4217 alphabetic currency code e.g. 'INR'
	Currency string `json:"currency"`
	// Describes a decimal value
	Value string `json:"value"`
	// Describes a decimal value
	EstimatedValue *string `json:"estimated_value,omitempty"`
	// Describes a decimal value
	ComputedValue *string `json:"computed_value,omitempty"`
	// Describes a decimal value
	ListedValue *string `json:"listed_value,omitempty"`
	// Describes a decimal value
	OfferedValue *string `json:"offered_value,omitempty"`
	// Describes a decimal value
	MinimumValue *string `json:"minimum_value,omitempty"`
	// Describes a decimal value
	MaximumValue *string `json:"maximum_value,omitempty"`
}

// NewPrice instantiates a new Price object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrice(currency string, value string) *Price {
	this := Price{}
	this.Currency = currency
	this.Value = value
	return &this
}

// NewPriceWithDefaults instantiates a new Price object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceWithDefaults() *Price {
	this := Price{}
	return &this
}

// GetCurrency returns the Currency field value
func (o *Price) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Price) GetCurrencyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Price) SetCurrency(v string) {
	o.Currency = v
}

// GetValue returns the Value field value
func (o *Price) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Price) GetValueOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Price) SetValue(v string) {
	o.Value = v
}

// GetEstimatedValue returns the EstimatedValue field value if set, zero value otherwise.
func (o *Price) GetEstimatedValue() string {
	if o == nil || isNil(o.EstimatedValue) {
		var ret string
		return ret
	}
	return *o.EstimatedValue
}

// GetEstimatedValueOk returns a tuple with the EstimatedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetEstimatedValueOk() (*string, bool) {
	if o == nil || isNil(o.EstimatedValue) {
    return nil, false
	}
	return o.EstimatedValue, true
}

// HasEstimatedValue returns a boolean if a field has been set.
func (o *Price) HasEstimatedValue() bool {
	if o != nil && !isNil(o.EstimatedValue) {
		return true
	}

	return false
}

// SetEstimatedValue gets a reference to the given string and assigns it to the EstimatedValue field.
func (o *Price) SetEstimatedValue(v string) {
	o.EstimatedValue = &v
}

// GetComputedValue returns the ComputedValue field value if set, zero value otherwise.
func (o *Price) GetComputedValue() string {
	if o == nil || isNil(o.ComputedValue) {
		var ret string
		return ret
	}
	return *o.ComputedValue
}

// GetComputedValueOk returns a tuple with the ComputedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetComputedValueOk() (*string, bool) {
	if o == nil || isNil(o.ComputedValue) {
    return nil, false
	}
	return o.ComputedValue, true
}

// HasComputedValue returns a boolean if a field has been set.
func (o *Price) HasComputedValue() bool {
	if o != nil && !isNil(o.ComputedValue) {
		return true
	}

	return false
}

// SetComputedValue gets a reference to the given string and assigns it to the ComputedValue field.
func (o *Price) SetComputedValue(v string) {
	o.ComputedValue = &v
}

// GetListedValue returns the ListedValue field value if set, zero value otherwise.
func (o *Price) GetListedValue() string {
	if o == nil || isNil(o.ListedValue) {
		var ret string
		return ret
	}
	return *o.ListedValue
}

// GetListedValueOk returns a tuple with the ListedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetListedValueOk() (*string, bool) {
	if o == nil || isNil(o.ListedValue) {
    return nil, false
	}
	return o.ListedValue, true
}

// HasListedValue returns a boolean if a field has been set.
func (o *Price) HasListedValue() bool {
	if o != nil && !isNil(o.ListedValue) {
		return true
	}

	return false
}

// SetListedValue gets a reference to the given string and assigns it to the ListedValue field.
func (o *Price) SetListedValue(v string) {
	o.ListedValue = &v
}

// GetOfferedValue returns the OfferedValue field value if set, zero value otherwise.
func (o *Price) GetOfferedValue() string {
	if o == nil || isNil(o.OfferedValue) {
		var ret string
		return ret
	}
	return *o.OfferedValue
}

// GetOfferedValueOk returns a tuple with the OfferedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetOfferedValueOk() (*string, bool) {
	if o == nil || isNil(o.OfferedValue) {
    return nil, false
	}
	return o.OfferedValue, true
}

// HasOfferedValue returns a boolean if a field has been set.
func (o *Price) HasOfferedValue() bool {
	if o != nil && !isNil(o.OfferedValue) {
		return true
	}

	return false
}

// SetOfferedValue gets a reference to the given string and assigns it to the OfferedValue field.
func (o *Price) SetOfferedValue(v string) {
	o.OfferedValue = &v
}

// GetMinimumValue returns the MinimumValue field value if set, zero value otherwise.
func (o *Price) GetMinimumValue() string {
	if o == nil || isNil(o.MinimumValue) {
		var ret string
		return ret
	}
	return *o.MinimumValue
}

// GetMinimumValueOk returns a tuple with the MinimumValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetMinimumValueOk() (*string, bool) {
	if o == nil || isNil(o.MinimumValue) {
    return nil, false
	}
	return o.MinimumValue, true
}

// HasMinimumValue returns a boolean if a field has been set.
func (o *Price) HasMinimumValue() bool {
	if o != nil && !isNil(o.MinimumValue) {
		return true
	}

	return false
}

// SetMinimumValue gets a reference to the given string and assigns it to the MinimumValue field.
func (o *Price) SetMinimumValue(v string) {
	o.MinimumValue = &v
}

// GetMaximumValue returns the MaximumValue field value if set, zero value otherwise.
func (o *Price) GetMaximumValue() string {
	if o == nil || isNil(o.MaximumValue) {
		var ret string
		return ret
	}
	return *o.MaximumValue
}

// GetMaximumValueOk returns a tuple with the MaximumValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Price) GetMaximumValueOk() (*string, bool) {
	if o == nil || isNil(o.MaximumValue) {
    return nil, false
	}
	return o.MaximumValue, true
}

// HasMaximumValue returns a boolean if a field has been set.
func (o *Price) HasMaximumValue() bool {
	if o != nil && !isNil(o.MaximumValue) {
		return true
	}

	return false
}

// SetMaximumValue gets a reference to the given string and assigns it to the MaximumValue field.
func (o *Price) SetMaximumValue(v string) {
	o.MaximumValue = &v
}

func (o Price) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if !isNil(o.EstimatedValue) {
		toSerialize["estimated_value"] = o.EstimatedValue
	}
	if !isNil(o.ComputedValue) {
		toSerialize["computed_value"] = o.ComputedValue
	}
	if !isNil(o.ListedValue) {
		toSerialize["listed_value"] = o.ListedValue
	}
	if !isNil(o.OfferedValue) {
		toSerialize["offered_value"] = o.OfferedValue
	}
	if !isNil(o.MinimumValue) {
		toSerialize["minimum_value"] = o.MinimumValue
	}
	if !isNil(o.MaximumValue) {
		toSerialize["maximum_value"] = o.MaximumValue
	}
	return json.Marshal(toSerialize)
}

type NullablePrice struct {
	value *Price
	isSet bool
}

func (v NullablePrice) Get() *Price {
	return v.value
}

func (v *NullablePrice) Set(val *Price) {
	v.value = val
	v.isSet = true
}

func (v NullablePrice) IsSet() bool {
	return v.isSet
}

func (v *NullablePrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrice(val *Price) *NullablePrice {
	return &NullablePrice{value: val, isSet: true}
}

func (v NullablePrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


