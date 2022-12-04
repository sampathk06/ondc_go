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

// StatusPostRequest struct for StatusPostRequest
type StatusPostRequest struct {
	Context Context `json:"context"`
	Message StatusPostRequestMessage `json:"message"`
}

// NewStatusPostRequest instantiates a new StatusPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusPostRequest(context Context, message StatusPostRequestMessage) *StatusPostRequest {
	this := StatusPostRequest{}
	this.Context = context
	this.Message = message
	return &this
}

// NewStatusPostRequestWithDefaults instantiates a new StatusPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusPostRequestWithDefaults() *StatusPostRequest {
	this := StatusPostRequest{}
	return &this
}

// GetContext returns the Context field value
func (o *StatusPostRequest) GetContext() Context {
	if o == nil {
		var ret Context
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *StatusPostRequest) GetContextOk() (*Context, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *StatusPostRequest) SetContext(v Context) {
	o.Context = v
}

// GetMessage returns the Message field value
func (o *StatusPostRequest) GetMessage() StatusPostRequestMessage {
	if o == nil {
		var ret StatusPostRequestMessage
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *StatusPostRequest) GetMessageOk() (*StatusPostRequestMessage, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *StatusPostRequest) SetMessage(v StatusPostRequestMessage) {
	o.Message = v
}

func (o StatusPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["context"] = o.Context
	}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableStatusPostRequest struct {
	value *StatusPostRequest
	isSet bool
}

func (v NullableStatusPostRequest) Get() *StatusPostRequest {
	return v.value
}

func (v *NullableStatusPostRequest) Set(val *StatusPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusPostRequest(val *StatusPostRequest) *NullableStatusPostRequest {
	return &NullableStatusPostRequest{value: val, isSet: true}
}

func (v NullableStatusPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


