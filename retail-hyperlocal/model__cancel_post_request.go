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

// CancelPostRequest struct for CancelPostRequest
type CancelPostRequest struct {
	Context Context `json:"context"`
	Message CancelPostRequestMessage `json:"message"`
}

// NewCancelPostRequest instantiates a new CancelPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelPostRequest(context Context, message CancelPostRequestMessage) *CancelPostRequest {
	this := CancelPostRequest{}
	this.Context = context
	this.Message = message
	return &this
}

// NewCancelPostRequestWithDefaults instantiates a new CancelPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelPostRequestWithDefaults() *CancelPostRequest {
	this := CancelPostRequest{}
	return &this
}

// GetContext returns the Context field value
func (o *CancelPostRequest) GetContext() Context {
	if o == nil {
		var ret Context
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *CancelPostRequest) GetContextOk() (*Context, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *CancelPostRequest) SetContext(v Context) {
	o.Context = v
}

// GetMessage returns the Message field value
func (o *CancelPostRequest) GetMessage() CancelPostRequestMessage {
	if o == nil {
		var ret CancelPostRequestMessage
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *CancelPostRequest) GetMessageOk() (*CancelPostRequestMessage, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *CancelPostRequest) SetMessage(v CancelPostRequestMessage) {
	o.Message = v
}

func (o CancelPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["context"] = o.Context
	}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableCancelPostRequest struct {
	value *CancelPostRequest
	isSet bool
}

func (v NullableCancelPostRequest) Get() *CancelPostRequest {
	return v.value
}

func (v *NullableCancelPostRequest) Set(val *CancelPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelPostRequest(val *CancelPostRequest) *NullableCancelPostRequest {
	return &NullableCancelPostRequest{value: val, isSet: true}
}

func (v NullableCancelPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


