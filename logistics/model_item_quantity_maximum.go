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

// ItemQuantityMaximum struct for ItemQuantityMaximum
type ItemQuantityMaximum struct {
	Count *int32 `json:"count,omitempty"`
	Measure *Scalar `json:"measure,omitempty"`
}

// NewItemQuantityMaximum instantiates a new ItemQuantityMaximum object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemQuantityMaximum() *ItemQuantityMaximum {
	this := ItemQuantityMaximum{}
	return &this
}

// NewItemQuantityMaximumWithDefaults instantiates a new ItemQuantityMaximum object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemQuantityMaximumWithDefaults() *ItemQuantityMaximum {
	this := ItemQuantityMaximum{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ItemQuantityMaximum) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemQuantityMaximum) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ItemQuantityMaximum) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ItemQuantityMaximum) SetCount(v int32) {
	o.Count = &v
}

// GetMeasure returns the Measure field value if set, zero value otherwise.
func (o *ItemQuantityMaximum) GetMeasure() Scalar {
	if o == nil || isNil(o.Measure) {
		var ret Scalar
		return ret
	}
	return *o.Measure
}

// GetMeasureOk returns a tuple with the Measure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemQuantityMaximum) GetMeasureOk() (*Scalar, bool) {
	if o == nil || isNil(o.Measure) {
    return nil, false
	}
	return o.Measure, true
}

// HasMeasure returns a boolean if a field has been set.
func (o *ItemQuantityMaximum) HasMeasure() bool {
	if o != nil && !isNil(o.Measure) {
		return true
	}

	return false
}

// SetMeasure gets a reference to the given Scalar and assigns it to the Measure field.
func (o *ItemQuantityMaximum) SetMeasure(v Scalar) {
	o.Measure = &v
}

func (o ItemQuantityMaximum) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !isNil(o.Measure) {
		toSerialize["measure"] = o.Measure
	}
	return json.Marshal(toSerialize)
}

type NullableItemQuantityMaximum struct {
	value *ItemQuantityMaximum
	isSet bool
}

func (v NullableItemQuantityMaximum) Get() *ItemQuantityMaximum {
	return v.value
}

func (v *NullableItemQuantityMaximum) Set(val *ItemQuantityMaximum) {
	v.value = val
	v.isSet = true
}

func (v NullableItemQuantityMaximum) IsSet() bool {
	return v.isSet
}

func (v *NullableItemQuantityMaximum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemQuantityMaximum(val *ItemQuantityMaximum) *NullableItemQuantityMaximum {
	return &NullableItemQuantityMaximum{value: val, isSet: true}
}

func (v NullableItemQuantityMaximum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemQuantityMaximum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


