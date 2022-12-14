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

// OrderAddOnsInner struct for OrderAddOnsInner
type OrderAddOnsInner struct {
	Id Id `json:"id"`
}

// NewOrderAddOnsInner instantiates a new OrderAddOnsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderAddOnsInner(id Id) *OrderAddOnsInner {
	this := OrderAddOnsInner{}
	this.Id = id
	return &this
}

// NewOrderAddOnsInnerWithDefaults instantiates a new OrderAddOnsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderAddOnsInnerWithDefaults() *OrderAddOnsInner {
	this := OrderAddOnsInner{}
	return &this
}

// GetId returns the Id field value
func (o *OrderAddOnsInner) GetId() Id {
	if o == nil {
		var ret Id
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrderAddOnsInner) GetIdOk() (*Id, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OrderAddOnsInner) SetId(v Id) {
	o.Id = v
}

func (o OrderAddOnsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableOrderAddOnsInner struct {
	value *OrderAddOnsInner
	isSet bool
}

func (v NullableOrderAddOnsInner) Get() *OrderAddOnsInner {
	return v.value
}

func (v *NullableOrderAddOnsInner) Set(val *OrderAddOnsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderAddOnsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderAddOnsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderAddOnsInner(val *OrderAddOnsInner) *NullableOrderAddOnsInner {
	return &NullableOrderAddOnsInner{value: val, isSet: true}
}

func (v NullableOrderAddOnsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderAddOnsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


