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

// SupportPostRequestMessage struct for SupportPostRequestMessage
type SupportPostRequestMessage struct {
	// ID of the element for which support is needed
	RefId *string `json:"ref_id,omitempty"`
}

// NewSupportPostRequestMessage instantiates a new SupportPostRequestMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportPostRequestMessage() *SupportPostRequestMessage {
	this := SupportPostRequestMessage{}
	return &this
}

// NewSupportPostRequestMessageWithDefaults instantiates a new SupportPostRequestMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportPostRequestMessageWithDefaults() *SupportPostRequestMessage {
	this := SupportPostRequestMessage{}
	return &this
}

// GetRefId returns the RefId field value if set, zero value otherwise.
func (o *SupportPostRequestMessage) GetRefId() string {
	if o == nil || isNil(o.RefId) {
		var ret string
		return ret
	}
	return *o.RefId
}

// GetRefIdOk returns a tuple with the RefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportPostRequestMessage) GetRefIdOk() (*string, bool) {
	if o == nil || isNil(o.RefId) {
    return nil, false
	}
	return o.RefId, true
}

// HasRefId returns a boolean if a field has been set.
func (o *SupportPostRequestMessage) HasRefId() bool {
	if o != nil && !isNil(o.RefId) {
		return true
	}

	return false
}

// SetRefId gets a reference to the given string and assigns it to the RefId field.
func (o *SupportPostRequestMessage) SetRefId(v string) {
	o.RefId = &v
}

func (o SupportPostRequestMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RefId) {
		toSerialize["ref_id"] = o.RefId
	}
	return json.Marshal(toSerialize)
}

type NullableSupportPostRequestMessage struct {
	value *SupportPostRequestMessage
	isSet bool
}

func (v NullableSupportPostRequestMessage) Get() *SupportPostRequestMessage {
	return v.value
}

func (v *NullableSupportPostRequestMessage) Set(val *SupportPostRequestMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportPostRequestMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportPostRequestMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportPostRequestMessage(val *SupportPostRequestMessage) *NullableSupportPostRequestMessage {
	return &NullableSupportPostRequestMessage{value: val, isSet: true}
}

func (v NullableSupportPostRequestMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportPostRequestMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


