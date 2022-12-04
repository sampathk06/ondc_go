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

// IntentOndcOrgPayloadDetails payload details that will allow logistics provider to determine serviceability. For weight, enums for unit are - \"Kilogram\", \"Gram\" For dimensions, enums for length.unit, breadth.unit and height.unit are - \"meter\", \"centimeter\"
type IntentOndcOrgPayloadDetails struct {
	Weight *Scalar `json:"weight,omitempty"`
	Dimensions *Dimensions `json:"dimensions,omitempty"`
	Category *string `json:"category,omitempty"`
	Value *Price `json:"value,omitempty"`
	DangerousGoods *bool `json:"dangerous_goods,omitempty"`
}

// NewIntentOndcOrgPayloadDetails instantiates a new IntentOndcOrgPayloadDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntentOndcOrgPayloadDetails() *IntentOndcOrgPayloadDetails {
	this := IntentOndcOrgPayloadDetails{}
	return &this
}

// NewIntentOndcOrgPayloadDetailsWithDefaults instantiates a new IntentOndcOrgPayloadDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntentOndcOrgPayloadDetailsWithDefaults() *IntentOndcOrgPayloadDetails {
	this := IntentOndcOrgPayloadDetails{}
	return &this
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *IntentOndcOrgPayloadDetails) GetWeight() Scalar {
	if o == nil || isNil(o.Weight) {
		var ret Scalar
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntentOndcOrgPayloadDetails) GetWeightOk() (*Scalar, bool) {
	if o == nil || isNil(o.Weight) {
    return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *IntentOndcOrgPayloadDetails) HasWeight() bool {
	if o != nil && !isNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given Scalar and assigns it to the Weight field.
func (o *IntentOndcOrgPayloadDetails) SetWeight(v Scalar) {
	o.Weight = &v
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *IntentOndcOrgPayloadDetails) GetDimensions() Dimensions {
	if o == nil || isNil(o.Dimensions) {
		var ret Dimensions
		return ret
	}
	return *o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntentOndcOrgPayloadDetails) GetDimensionsOk() (*Dimensions, bool) {
	if o == nil || isNil(o.Dimensions) {
    return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *IntentOndcOrgPayloadDetails) HasDimensions() bool {
	if o != nil && !isNil(o.Dimensions) {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given Dimensions and assigns it to the Dimensions field.
func (o *IntentOndcOrgPayloadDetails) SetDimensions(v Dimensions) {
	o.Dimensions = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *IntentOndcOrgPayloadDetails) GetCategory() string {
	if o == nil || isNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntentOndcOrgPayloadDetails) GetCategoryOk() (*string, bool) {
	if o == nil || isNil(o.Category) {
    return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *IntentOndcOrgPayloadDetails) HasCategory() bool {
	if o != nil && !isNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *IntentOndcOrgPayloadDetails) SetCategory(v string) {
	o.Category = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *IntentOndcOrgPayloadDetails) GetValue() Price {
	if o == nil || isNil(o.Value) {
		var ret Price
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntentOndcOrgPayloadDetails) GetValueOk() (*Price, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *IntentOndcOrgPayloadDetails) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given Price and assigns it to the Value field.
func (o *IntentOndcOrgPayloadDetails) SetValue(v Price) {
	o.Value = &v
}

// GetDangerousGoods returns the DangerousGoods field value if set, zero value otherwise.
func (o *IntentOndcOrgPayloadDetails) GetDangerousGoods() bool {
	if o == nil || isNil(o.DangerousGoods) {
		var ret bool
		return ret
	}
	return *o.DangerousGoods
}

// GetDangerousGoodsOk returns a tuple with the DangerousGoods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntentOndcOrgPayloadDetails) GetDangerousGoodsOk() (*bool, bool) {
	if o == nil || isNil(o.DangerousGoods) {
    return nil, false
	}
	return o.DangerousGoods, true
}

// HasDangerousGoods returns a boolean if a field has been set.
func (o *IntentOndcOrgPayloadDetails) HasDangerousGoods() bool {
	if o != nil && !isNil(o.DangerousGoods) {
		return true
	}

	return false
}

// SetDangerousGoods gets a reference to the given bool and assigns it to the DangerousGoods field.
func (o *IntentOndcOrgPayloadDetails) SetDangerousGoods(v bool) {
	o.DangerousGoods = &v
}

func (o IntentOndcOrgPayloadDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !isNil(o.Dimensions) {
		toSerialize["dimensions"] = o.Dimensions
	}
	if !isNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !isNil(o.DangerousGoods) {
		toSerialize["dangerous_goods"] = o.DangerousGoods
	}
	return json.Marshal(toSerialize)
}

type NullableIntentOndcOrgPayloadDetails struct {
	value *IntentOndcOrgPayloadDetails
	isSet bool
}

func (v NullableIntentOndcOrgPayloadDetails) Get() *IntentOndcOrgPayloadDetails {
	return v.value
}

func (v *NullableIntentOndcOrgPayloadDetails) Set(val *IntentOndcOrgPayloadDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableIntentOndcOrgPayloadDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableIntentOndcOrgPayloadDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntentOndcOrgPayloadDetails(val *IntentOndcOrgPayloadDetails) *NullableIntentOndcOrgPayloadDetails {
	return &NullableIntentOndcOrgPayloadDetails{value: val, isSet: true}
}

func (v NullableIntentOndcOrgPayloadDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntentOndcOrgPayloadDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


