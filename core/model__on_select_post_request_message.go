/*
ONDC Protocol Core API

ONDC Protocol Core API specification. This is an adaptation of Beckn Core version 0.9.3

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// OnSelectPostRequestMessage struct for OnSelectPostRequestMessage
type OnSelectPostRequestMessage struct {
	Order OnSelectPostRequestMessageOrder `json:"order"`
}

// NewOnSelectPostRequestMessage instantiates a new OnSelectPostRequestMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnSelectPostRequestMessage(order OnSelectPostRequestMessageOrder) *OnSelectPostRequestMessage {
	this := OnSelectPostRequestMessage{}
	this.Order = order
	return &this
}

// NewOnSelectPostRequestMessageWithDefaults instantiates a new OnSelectPostRequestMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnSelectPostRequestMessageWithDefaults() *OnSelectPostRequestMessage {
	this := OnSelectPostRequestMessage{}
	return &this
}

// GetOrder returns the Order field value
func (o *OnSelectPostRequestMessage) GetOrder() OnSelectPostRequestMessageOrder {
	if o == nil {
		var ret OnSelectPostRequestMessageOrder
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *OnSelectPostRequestMessage) GetOrderOk() (*OnSelectPostRequestMessageOrder, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *OnSelectPostRequestMessage) SetOrder(v OnSelectPostRequestMessageOrder) {
	o.Order = v
}

func (o OnSelectPostRequestMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["order"] = o.Order
	}
	return json.Marshal(toSerialize)
}

type NullableOnSelectPostRequestMessage struct {
	value *OnSelectPostRequestMessage
	isSet bool
}

func (v NullableOnSelectPostRequestMessage) Get() *OnSelectPostRequestMessage {
	return v.value
}

func (v *NullableOnSelectPostRequestMessage) Set(val *OnSelectPostRequestMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableOnSelectPostRequestMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableOnSelectPostRequestMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnSelectPostRequestMessage(val *OnSelectPostRequestMessage) *NullableOnSelectPostRequestMessage {
	return &NullableOnSelectPostRequestMessage{value: val, isSet: true}
}

func (v NullableOnSelectPostRequestMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnSelectPostRequestMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


