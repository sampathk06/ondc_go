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

// ReturnReasonsPostRequest struct for ReturnReasonsPostRequest
type ReturnReasonsPostRequest struct {
	Context *Context `json:"context,omitempty"`
	ReturnReasons []Option `json:"return_reasons,omitempty"`
}

// NewReturnReasonsPostRequest instantiates a new ReturnReasonsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnReasonsPostRequest() *ReturnReasonsPostRequest {
	this := ReturnReasonsPostRequest{}
	return &this
}

// NewReturnReasonsPostRequestWithDefaults instantiates a new ReturnReasonsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnReasonsPostRequestWithDefaults() *ReturnReasonsPostRequest {
	this := ReturnReasonsPostRequest{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *ReturnReasonsPostRequest) GetContext() Context {
	if o == nil || isNil(o.Context) {
		var ret Context
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnReasonsPostRequest) GetContextOk() (*Context, bool) {
	if o == nil || isNil(o.Context) {
    return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *ReturnReasonsPostRequest) HasContext() bool {
	if o != nil && !isNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given Context and assigns it to the Context field.
func (o *ReturnReasonsPostRequest) SetContext(v Context) {
	o.Context = &v
}

// GetReturnReasons returns the ReturnReasons field value if set, zero value otherwise.
func (o *ReturnReasonsPostRequest) GetReturnReasons() []Option {
	if o == nil || isNil(o.ReturnReasons) {
		var ret []Option
		return ret
	}
	return o.ReturnReasons
}

// GetReturnReasonsOk returns a tuple with the ReturnReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnReasonsPostRequest) GetReturnReasonsOk() ([]Option, bool) {
	if o == nil || isNil(o.ReturnReasons) {
    return nil, false
	}
	return o.ReturnReasons, true
}

// HasReturnReasons returns a boolean if a field has been set.
func (o *ReturnReasonsPostRequest) HasReturnReasons() bool {
	if o != nil && !isNil(o.ReturnReasons) {
		return true
	}

	return false
}

// SetReturnReasons gets a reference to the given []Option and assigns it to the ReturnReasons field.
func (o *ReturnReasonsPostRequest) SetReturnReasons(v []Option) {
	o.ReturnReasons = v
}

func (o ReturnReasonsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !isNil(o.ReturnReasons) {
		toSerialize["return_reasons"] = o.ReturnReasons
	}
	return json.Marshal(toSerialize)
}

type NullableReturnReasonsPostRequest struct {
	value *ReturnReasonsPostRequest
	isSet bool
}

func (v NullableReturnReasonsPostRequest) Get() *ReturnReasonsPostRequest {
	return v.value
}

func (v *NullableReturnReasonsPostRequest) Set(val *ReturnReasonsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnReasonsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnReasonsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnReasonsPostRequest(val *ReturnReasonsPostRequest) *NullableReturnReasonsPostRequest {
	return &NullableReturnReasonsPostRequest{value: val, isSet: true}
}

func (v NullableReturnReasonsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnReasonsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

