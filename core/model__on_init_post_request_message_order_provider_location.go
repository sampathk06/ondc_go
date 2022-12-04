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

// OnInitPostRequestMessageOrderProviderLocation struct for OnInitPostRequestMessageOrderProviderLocation
type OnInitPostRequestMessageOrderProviderLocation struct {
	Id *Id `json:"id,omitempty"`
}

// NewOnInitPostRequestMessageOrderProviderLocation instantiates a new OnInitPostRequestMessageOrderProviderLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnInitPostRequestMessageOrderProviderLocation() *OnInitPostRequestMessageOrderProviderLocation {
	this := OnInitPostRequestMessageOrderProviderLocation{}
	return &this
}

// NewOnInitPostRequestMessageOrderProviderLocationWithDefaults instantiates a new OnInitPostRequestMessageOrderProviderLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnInitPostRequestMessageOrderProviderLocationWithDefaults() *OnInitPostRequestMessageOrderProviderLocation {
	this := OnInitPostRequestMessageOrderProviderLocation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OnInitPostRequestMessageOrderProviderLocation) GetId() Id {
	if o == nil || isNil(o.Id) {
		var ret Id
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnInitPostRequestMessageOrderProviderLocation) GetIdOk() (*Id, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OnInitPostRequestMessageOrderProviderLocation) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given Id and assigns it to the Id field.
func (o *OnInitPostRequestMessageOrderProviderLocation) SetId(v Id) {
	o.Id = &v
}

func (o OnInitPostRequestMessageOrderProviderLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableOnInitPostRequestMessageOrderProviderLocation struct {
	value *OnInitPostRequestMessageOrderProviderLocation
	isSet bool
}

func (v NullableOnInitPostRequestMessageOrderProviderLocation) Get() *OnInitPostRequestMessageOrderProviderLocation {
	return v.value
}

func (v *NullableOnInitPostRequestMessageOrderProviderLocation) Set(val *OnInitPostRequestMessageOrderProviderLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableOnInitPostRequestMessageOrderProviderLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableOnInitPostRequestMessageOrderProviderLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnInitPostRequestMessageOrderProviderLocation(val *OnInitPostRequestMessageOrderProviderLocation) *NullableOnInitPostRequestMessageOrderProviderLocation {
	return &NullableOnInitPostRequestMessageOrderProviderLocation{value: val, isSet: true}
}

func (v NullableOnInitPostRequestMessageOrderProviderLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnInitPostRequestMessageOrderProviderLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


