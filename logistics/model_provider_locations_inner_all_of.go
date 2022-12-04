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

// ProviderLocationsInnerAllOf struct for ProviderLocationsInnerAllOf
type ProviderLocationsInnerAllOf struct {
	// If the entity can be rated or not
	Rateable *bool `json:"rateable,omitempty"`
}

// NewProviderLocationsInnerAllOf instantiates a new ProviderLocationsInnerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderLocationsInnerAllOf() *ProviderLocationsInnerAllOf {
	this := ProviderLocationsInnerAllOf{}
	return &this
}

// NewProviderLocationsInnerAllOfWithDefaults instantiates a new ProviderLocationsInnerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderLocationsInnerAllOfWithDefaults() *ProviderLocationsInnerAllOf {
	this := ProviderLocationsInnerAllOf{}
	return &this
}

// GetRateable returns the Rateable field value if set, zero value otherwise.
func (o *ProviderLocationsInnerAllOf) GetRateable() bool {
	if o == nil || isNil(o.Rateable) {
		var ret bool
		return ret
	}
	return *o.Rateable
}

// GetRateableOk returns a tuple with the Rateable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderLocationsInnerAllOf) GetRateableOk() (*bool, bool) {
	if o == nil || isNil(o.Rateable) {
    return nil, false
	}
	return o.Rateable, true
}

// HasRateable returns a boolean if a field has been set.
func (o *ProviderLocationsInnerAllOf) HasRateable() bool {
	if o != nil && !isNil(o.Rateable) {
		return true
	}

	return false
}

// SetRateable gets a reference to the given bool and assigns it to the Rateable field.
func (o *ProviderLocationsInnerAllOf) SetRateable(v bool) {
	o.Rateable = &v
}

func (o ProviderLocationsInnerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Rateable) {
		toSerialize["rateable"] = o.Rateable
	}
	return json.Marshal(toSerialize)
}

type NullableProviderLocationsInnerAllOf struct {
	value *ProviderLocationsInnerAllOf
	isSet bool
}

func (v NullableProviderLocationsInnerAllOf) Get() *ProviderLocationsInnerAllOf {
	return v.value
}

func (v *NullableProviderLocationsInnerAllOf) Set(val *ProviderLocationsInnerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderLocationsInnerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderLocationsInnerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderLocationsInnerAllOf(val *ProviderLocationsInnerAllOf) *NullableProviderLocationsInnerAllOf {
	return &NullableProviderLocationsInnerAllOf{value: val, isSet: true}
}

func (v NullableProviderLocationsInnerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderLocationsInnerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


