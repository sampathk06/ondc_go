/*
ONDC Protocol API for retail (grocery, f&b)

ONDC Protocol API specification, which includes statutory requirements for packaged commodities and pre-packaged food This is an adaptation of Beckn Core version 0.9.3

API version: 1.0.27
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ItemQuantityAllocated struct for ItemQuantityAllocated
type ItemQuantityAllocated struct {
	Count *int32 `json:"count,omitempty"`
	Measure *Scalar `json:"measure,omitempty"`
}

// NewItemQuantityAllocated instantiates a new ItemQuantityAllocated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemQuantityAllocated() *ItemQuantityAllocated {
	this := ItemQuantityAllocated{}
	return &this
}

// NewItemQuantityAllocatedWithDefaults instantiates a new ItemQuantityAllocated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemQuantityAllocatedWithDefaults() *ItemQuantityAllocated {
	this := ItemQuantityAllocated{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ItemQuantityAllocated) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemQuantityAllocated) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ItemQuantityAllocated) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ItemQuantityAllocated) SetCount(v int32) {
	o.Count = &v
}

// GetMeasure returns the Measure field value if set, zero value otherwise.
func (o *ItemQuantityAllocated) GetMeasure() Scalar {
	if o == nil || isNil(o.Measure) {
		var ret Scalar
		return ret
	}
	return *o.Measure
}

// GetMeasureOk returns a tuple with the Measure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemQuantityAllocated) GetMeasureOk() (*Scalar, bool) {
	if o == nil || isNil(o.Measure) {
    return nil, false
	}
	return o.Measure, true
}

// HasMeasure returns a boolean if a field has been set.
func (o *ItemQuantityAllocated) HasMeasure() bool {
	if o != nil && !isNil(o.Measure) {
		return true
	}

	return false
}

// SetMeasure gets a reference to the given Scalar and assigns it to the Measure field.
func (o *ItemQuantityAllocated) SetMeasure(v Scalar) {
	o.Measure = &v
}

func (o ItemQuantityAllocated) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !isNil(o.Measure) {
		toSerialize["measure"] = o.Measure
	}
	return json.Marshal(toSerialize)
}

type NullableItemQuantityAllocated struct {
	value *ItemQuantityAllocated
	isSet bool
}

func (v NullableItemQuantityAllocated) Get() *ItemQuantityAllocated {
	return v.value
}

func (v *NullableItemQuantityAllocated) Set(val *ItemQuantityAllocated) {
	v.value = val
	v.isSet = true
}

func (v NullableItemQuantityAllocated) IsSet() bool {
	return v.isSet
}

func (v *NullableItemQuantityAllocated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemQuantityAllocated(val *ItemQuantityAllocated) *NullableItemQuantityAllocated {
	return &NullableItemQuantityAllocated{value: val, isSet: true}
}

func (v NullableItemQuantityAllocated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemQuantityAllocated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


