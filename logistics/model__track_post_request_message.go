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

// TrackPostRequestMessage struct for TrackPostRequestMessage
type TrackPostRequestMessage struct {
	OrderId Id `json:"order_id"`
	CallbackUrl *string `json:"callback_url,omitempty"`
}

// NewTrackPostRequestMessage instantiates a new TrackPostRequestMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackPostRequestMessage(orderId Id) *TrackPostRequestMessage {
	this := TrackPostRequestMessage{}
	this.OrderId = orderId
	return &this
}

// NewTrackPostRequestMessageWithDefaults instantiates a new TrackPostRequestMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackPostRequestMessageWithDefaults() *TrackPostRequestMessage {
	this := TrackPostRequestMessage{}
	return &this
}

// GetOrderId returns the OrderId field value
func (o *TrackPostRequestMessage) GetOrderId() Id {
	if o == nil {
		var ret Id
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *TrackPostRequestMessage) GetOrderIdOk() (*Id, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *TrackPostRequestMessage) SetOrderId(v Id) {
	o.OrderId = v
}

// GetCallbackUrl returns the CallbackUrl field value if set, zero value otherwise.
func (o *TrackPostRequestMessage) GetCallbackUrl() string {
	if o == nil || isNil(o.CallbackUrl) {
		var ret string
		return ret
	}
	return *o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackPostRequestMessage) GetCallbackUrlOk() (*string, bool) {
	if o == nil || isNil(o.CallbackUrl) {
    return nil, false
	}
	return o.CallbackUrl, true
}

// HasCallbackUrl returns a boolean if a field has been set.
func (o *TrackPostRequestMessage) HasCallbackUrl() bool {
	if o != nil && !isNil(o.CallbackUrl) {
		return true
	}

	return false
}

// SetCallbackUrl gets a reference to the given string and assigns it to the CallbackUrl field.
func (o *TrackPostRequestMessage) SetCallbackUrl(v string) {
	o.CallbackUrl = &v
}

func (o TrackPostRequestMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["order_id"] = o.OrderId
	}
	if !isNil(o.CallbackUrl) {
		toSerialize["callback_url"] = o.CallbackUrl
	}
	return json.Marshal(toSerialize)
}

type NullableTrackPostRequestMessage struct {
	value *TrackPostRequestMessage
	isSet bool
}

func (v NullableTrackPostRequestMessage) Get() *TrackPostRequestMessage {
	return v.value
}

func (v *NullableTrackPostRequestMessage) Set(val *TrackPostRequestMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackPostRequestMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackPostRequestMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackPostRequestMessage(val *TrackPostRequestMessage) *NullableTrackPostRequestMessage {
	return &NullableTrackPostRequestMessage{value: val, isSet: true}
}

func (v NullableTrackPostRequestMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackPostRequestMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


