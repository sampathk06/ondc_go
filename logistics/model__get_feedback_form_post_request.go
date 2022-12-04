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

// GetFeedbackFormPostRequest struct for GetFeedbackFormPostRequest
type GetFeedbackFormPostRequest struct {
	Context *Context `json:"context,omitempty"`
	Message *GetFeedbackFormPostRequestMessage `json:"message,omitempty"`
}

// NewGetFeedbackFormPostRequest instantiates a new GetFeedbackFormPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFeedbackFormPostRequest() *GetFeedbackFormPostRequest {
	this := GetFeedbackFormPostRequest{}
	return &this
}

// NewGetFeedbackFormPostRequestWithDefaults instantiates a new GetFeedbackFormPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFeedbackFormPostRequestWithDefaults() *GetFeedbackFormPostRequest {
	this := GetFeedbackFormPostRequest{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *GetFeedbackFormPostRequest) GetContext() Context {
	if o == nil || isNil(o.Context) {
		var ret Context
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFeedbackFormPostRequest) GetContextOk() (*Context, bool) {
	if o == nil || isNil(o.Context) {
    return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *GetFeedbackFormPostRequest) HasContext() bool {
	if o != nil && !isNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given Context and assigns it to the Context field.
func (o *GetFeedbackFormPostRequest) SetContext(v Context) {
	o.Context = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetFeedbackFormPostRequest) GetMessage() GetFeedbackFormPostRequestMessage {
	if o == nil || isNil(o.Message) {
		var ret GetFeedbackFormPostRequestMessage
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFeedbackFormPostRequest) GetMessageOk() (*GetFeedbackFormPostRequestMessage, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetFeedbackFormPostRequest) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given GetFeedbackFormPostRequestMessage and assigns it to the Message field.
func (o *GetFeedbackFormPostRequest) SetMessage(v GetFeedbackFormPostRequestMessage) {
	o.Message = &v
}

func (o GetFeedbackFormPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableGetFeedbackFormPostRequest struct {
	value *GetFeedbackFormPostRequest
	isSet bool
}

func (v NullableGetFeedbackFormPostRequest) Get() *GetFeedbackFormPostRequest {
	return v.value
}

func (v *NullableGetFeedbackFormPostRequest) Set(val *GetFeedbackFormPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFeedbackFormPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFeedbackFormPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFeedbackFormPostRequest(val *GetFeedbackFormPostRequest) *NullableGetFeedbackFormPostRequest {
	return &NullableGetFeedbackFormPostRequest{value: val, isSet: true}
}

func (v NullableGetFeedbackFormPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFeedbackFormPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

