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

// OnInitPostRequest struct for OnInitPostRequest
type OnInitPostRequest struct {
	Context Context `json:"context"`
	Message *OnInitPostRequestMessage `json:"message,omitempty"`
	Error *Error `json:"error,omitempty"`
}

// NewOnInitPostRequest instantiates a new OnInitPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnInitPostRequest(context Context) *OnInitPostRequest {
	this := OnInitPostRequest{}
	this.Context = context
	return &this
}

// NewOnInitPostRequestWithDefaults instantiates a new OnInitPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnInitPostRequestWithDefaults() *OnInitPostRequest {
	this := OnInitPostRequest{}
	return &this
}

// GetContext returns the Context field value
func (o *OnInitPostRequest) GetContext() Context {
	if o == nil {
		var ret Context
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *OnInitPostRequest) GetContextOk() (*Context, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *OnInitPostRequest) SetContext(v Context) {
	o.Context = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *OnInitPostRequest) GetMessage() OnInitPostRequestMessage {
	if o == nil || isNil(o.Message) {
		var ret OnInitPostRequestMessage
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnInitPostRequest) GetMessageOk() (*OnInitPostRequestMessage, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *OnInitPostRequest) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given OnInitPostRequestMessage and assigns it to the Message field.
func (o *OnInitPostRequest) SetMessage(v OnInitPostRequestMessage) {
	o.Message = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *OnInitPostRequest) GetError() Error {
	if o == nil || isNil(o.Error) {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnInitPostRequest) GetErrorOk() (*Error, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *OnInitPostRequest) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *OnInitPostRequest) SetError(v Error) {
	o.Error = &v
}

func (o OnInitPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["context"] = o.Context
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableOnInitPostRequest struct {
	value *OnInitPostRequest
	isSet bool
}

func (v NullableOnInitPostRequest) Get() *OnInitPostRequest {
	return v.value
}

func (v *NullableOnInitPostRequest) Set(val *OnInitPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOnInitPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOnInitPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnInitPostRequest(val *OnInitPostRequest) *NullableOnInitPostRequest {
	return &NullableOnInitPostRequest{value: val, isSet: true}
}

func (v NullableOnInitPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnInitPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


