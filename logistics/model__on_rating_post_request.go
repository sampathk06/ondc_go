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

// OnRatingPostRequest struct for OnRatingPostRequest
type OnRatingPostRequest struct {
	Context Context `json:"context"`
	Message *RatingAck `json:"message,omitempty"`
	Error *Error `json:"error,omitempty"`
}

// NewOnRatingPostRequest instantiates a new OnRatingPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnRatingPostRequest(context Context) *OnRatingPostRequest {
	this := OnRatingPostRequest{}
	this.Context = context
	return &this
}

// NewOnRatingPostRequestWithDefaults instantiates a new OnRatingPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnRatingPostRequestWithDefaults() *OnRatingPostRequest {
	this := OnRatingPostRequest{}
	return &this
}

// GetContext returns the Context field value
func (o *OnRatingPostRequest) GetContext() Context {
	if o == nil {
		var ret Context
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *OnRatingPostRequest) GetContextOk() (*Context, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *OnRatingPostRequest) SetContext(v Context) {
	o.Context = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *OnRatingPostRequest) GetMessage() RatingAck {
	if o == nil || isNil(o.Message) {
		var ret RatingAck
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnRatingPostRequest) GetMessageOk() (*RatingAck, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *OnRatingPostRequest) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given RatingAck and assigns it to the Message field.
func (o *OnRatingPostRequest) SetMessage(v RatingAck) {
	o.Message = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *OnRatingPostRequest) GetError() Error {
	if o == nil || isNil(o.Error) {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnRatingPostRequest) GetErrorOk() (*Error, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *OnRatingPostRequest) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *OnRatingPostRequest) SetError(v Error) {
	o.Error = &v
}

func (o OnRatingPostRequest) MarshalJSON() ([]byte, error) {
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

type NullableOnRatingPostRequest struct {
	value *OnRatingPostRequest
	isSet bool
}

func (v NullableOnRatingPostRequest) Get() *OnRatingPostRequest {
	return v.value
}

func (v *NullableOnRatingPostRequest) Set(val *OnRatingPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOnRatingPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOnRatingPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnRatingPostRequest(val *OnRatingPostRequest) *NullableOnRatingPostRequest {
	return &NullableOnRatingPostRequest{value: val, isSet: true}
}

func (v NullableOnRatingPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnRatingPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


