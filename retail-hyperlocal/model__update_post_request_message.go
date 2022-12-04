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

// UpdatePostRequestMessage struct for UpdatePostRequestMessage
type UpdatePostRequestMessage struct {
	// Comma separated values of order objects being updated. For example: ```\"update_target\":\"item,billing,fulfillment\"```
	UpdateTarget string `json:"update_target"`
	Order Order `json:"order"`
}

// NewUpdatePostRequestMessage instantiates a new UpdatePostRequestMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePostRequestMessage(updateTarget string, order Order) *UpdatePostRequestMessage {
	this := UpdatePostRequestMessage{}
	this.UpdateTarget = updateTarget
	this.Order = order
	return &this
}

// NewUpdatePostRequestMessageWithDefaults instantiates a new UpdatePostRequestMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePostRequestMessageWithDefaults() *UpdatePostRequestMessage {
	this := UpdatePostRequestMessage{}
	return &this
}

// GetUpdateTarget returns the UpdateTarget field value
func (o *UpdatePostRequestMessage) GetUpdateTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdateTarget
}

// GetUpdateTargetOk returns a tuple with the UpdateTarget field value
// and a boolean to check if the value has been set.
func (o *UpdatePostRequestMessage) GetUpdateTargetOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UpdateTarget, true
}

// SetUpdateTarget sets field value
func (o *UpdatePostRequestMessage) SetUpdateTarget(v string) {
	o.UpdateTarget = v
}

// GetOrder returns the Order field value
func (o *UpdatePostRequestMessage) GetOrder() Order {
	if o == nil {
		var ret Order
		return ret
	}

	return o.Order
}

// GetOrderOk returns a tuple with the Order field value
// and a boolean to check if the value has been set.
func (o *UpdatePostRequestMessage) GetOrderOk() (*Order, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Order, true
}

// SetOrder sets field value
func (o *UpdatePostRequestMessage) SetOrder(v Order) {
	o.Order = v
}

func (o UpdatePostRequestMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["update_target"] = o.UpdateTarget
	}
	if true {
		toSerialize["order"] = o.Order
	}
	return json.Marshal(toSerialize)
}

type NullableUpdatePostRequestMessage struct {
	value *UpdatePostRequestMessage
	isSet bool
}

func (v NullableUpdatePostRequestMessage) Get() *UpdatePostRequestMessage {
	return v.value
}

func (v *NullableUpdatePostRequestMessage) Set(val *UpdatePostRequestMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePostRequestMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePostRequestMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePostRequestMessage(val *UpdatePostRequestMessage) *NullableUpdatePostRequestMessage {
	return &NullableUpdatePostRequestMessage{value: val, isSet: true}
}

func (v NullableUpdatePostRequestMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePostRequestMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


