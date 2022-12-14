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

// UpdatePostRequest struct for UpdatePostRequest
type UpdatePostRequest struct {
	Context Context `json:"context"`
	Message UpdatePostRequestMessage `json:"message"`
}

// NewUpdatePostRequest instantiates a new UpdatePostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePostRequest(context Context, message UpdatePostRequestMessage) *UpdatePostRequest {
	this := UpdatePostRequest{}
	this.Context = context
	this.Message = message
	return &this
}

// NewUpdatePostRequestWithDefaults instantiates a new UpdatePostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePostRequestWithDefaults() *UpdatePostRequest {
	this := UpdatePostRequest{}
	return &this
}

// GetContext returns the Context field value
func (o *UpdatePostRequest) GetContext() Context {
	if o == nil {
		var ret Context
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *UpdatePostRequest) GetContextOk() (*Context, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *UpdatePostRequest) SetContext(v Context) {
	o.Context = v
}

// GetMessage returns the Message field value
func (o *UpdatePostRequest) GetMessage() UpdatePostRequestMessage {
	if o == nil {
		var ret UpdatePostRequestMessage
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *UpdatePostRequest) GetMessageOk() (*UpdatePostRequestMessage, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *UpdatePostRequest) SetMessage(v UpdatePostRequestMessage) {
	o.Message = v
}

func (o UpdatePostRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["context"] = o.Context
	}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableUpdatePostRequest struct {
	value *UpdatePostRequest
	isSet bool
}

func (v NullableUpdatePostRequest) Get() *UpdatePostRequest {
	return v.value
}

func (v *NullableUpdatePostRequest) Set(val *UpdatePostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePostRequest(val *UpdatePostRequest) *NullableUpdatePostRequest {
	return &NullableUpdatePostRequest{value: val, isSet: true}
}

func (v NullableUpdatePostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


