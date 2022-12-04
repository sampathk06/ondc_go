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

// OnSearchPost200Response struct for OnSearchPost200Response
type OnSearchPost200Response struct {
	Message *SearchPost200ResponseMessage `json:"message,omitempty"`
	Error *Error `json:"error,omitempty"`
}

// NewOnSearchPost200Response instantiates a new OnSearchPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnSearchPost200Response() *OnSearchPost200Response {
	this := OnSearchPost200Response{}
	return &this
}

// NewOnSearchPost200ResponseWithDefaults instantiates a new OnSearchPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnSearchPost200ResponseWithDefaults() *OnSearchPost200Response {
	this := OnSearchPost200Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *OnSearchPost200Response) GetMessage() SearchPost200ResponseMessage {
	if o == nil || isNil(o.Message) {
		var ret SearchPost200ResponseMessage
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnSearchPost200Response) GetMessageOk() (*SearchPost200ResponseMessage, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *OnSearchPost200Response) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given SearchPost200ResponseMessage and assigns it to the Message field.
func (o *OnSearchPost200Response) SetMessage(v SearchPost200ResponseMessage) {
	o.Message = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *OnSearchPost200Response) GetError() Error {
	if o == nil || isNil(o.Error) {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnSearchPost200Response) GetErrorOk() (*Error, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *OnSearchPost200Response) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *OnSearchPost200Response) SetError(v Error) {
	o.Error = &v
}

func (o OnSearchPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableOnSearchPost200Response struct {
	value *OnSearchPost200Response
	isSet bool
}

func (v NullableOnSearchPost200Response) Get() *OnSearchPost200Response {
	return v.value
}

func (v *NullableOnSearchPost200Response) Set(val *OnSearchPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableOnSearchPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableOnSearchPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnSearchPost200Response(val *OnSearchPost200Response) *NullableOnSearchPost200Response {
	return &NullableOnSearchPost200Response{value: val, isSet: true}
}

func (v NullableOnSearchPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnSearchPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


