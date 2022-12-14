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

// SearchPost200Response struct for SearchPost200Response
type SearchPost200Response struct {
	Message SearchPost200ResponseMessage `json:"message"`
	Error *Error `json:"error,omitempty"`
}

// NewSearchPost200Response instantiates a new SearchPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchPost200Response(message SearchPost200ResponseMessage) *SearchPost200Response {
	this := SearchPost200Response{}
	this.Message = message
	return &this
}

// NewSearchPost200ResponseWithDefaults instantiates a new SearchPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchPost200ResponseWithDefaults() *SearchPost200Response {
	this := SearchPost200Response{}
	return &this
}

// GetMessage returns the Message field value
func (o *SearchPost200Response) GetMessage() SearchPost200ResponseMessage {
	if o == nil {
		var ret SearchPost200ResponseMessage
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SearchPost200Response) GetMessageOk() (*SearchPost200ResponseMessage, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SearchPost200Response) SetMessage(v SearchPost200ResponseMessage) {
	o.Message = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SearchPost200Response) GetError() Error {
	if o == nil || isNil(o.Error) {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPost200Response) GetErrorOk() (*Error, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SearchPost200Response) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *SearchPost200Response) SetError(v Error) {
	o.Error = &v
}

func (o SearchPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableSearchPost200Response struct {
	value *SearchPost200Response
	isSet bool
}

func (v NullableSearchPost200Response) Get() *SearchPost200Response {
	return v.value
}

func (v *NullableSearchPost200Response) Set(val *SearchPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchPost200Response(val *SearchPost200Response) *NullableSearchPost200Response {
	return &NullableSearchPost200Response{value: val, isSet: true}
}

func (v NullableSearchPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


