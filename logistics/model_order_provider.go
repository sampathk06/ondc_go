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

// OrderProvider struct for OrderProvider
type OrderProvider struct {
	Id *Id `json:"id,omitempty"`
	Locations []OrderProviderLocationsInner `json:"locations,omitempty"`
}

// NewOrderProvider instantiates a new OrderProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderProvider() *OrderProvider {
	this := OrderProvider{}
	return &this
}

// NewOrderProviderWithDefaults instantiates a new OrderProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderProviderWithDefaults() *OrderProvider {
	this := OrderProvider{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderProvider) GetId() Id {
	if o == nil || isNil(o.Id) {
		var ret Id
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderProvider) GetIdOk() (*Id, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderProvider) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given Id and assigns it to the Id field.
func (o *OrderProvider) SetId(v Id) {
	o.Id = &v
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *OrderProvider) GetLocations() []OrderProviderLocationsInner {
	if o == nil || isNil(o.Locations) {
		var ret []OrderProviderLocationsInner
		return ret
	}
	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderProvider) GetLocationsOk() ([]OrderProviderLocationsInner, bool) {
	if o == nil || isNil(o.Locations) {
    return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *OrderProvider) HasLocations() bool {
	if o != nil && !isNil(o.Locations) {
		return true
	}

	return false
}

// SetLocations gets a reference to the given []OrderProviderLocationsInner and assigns it to the Locations field.
func (o *OrderProvider) SetLocations(v []OrderProviderLocationsInner) {
	o.Locations = v
}

func (o OrderProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Locations) {
		toSerialize["locations"] = o.Locations
	}
	return json.Marshal(toSerialize)
}

type NullableOrderProvider struct {
	value *OrderProvider
	isSet bool
}

func (v NullableOrderProvider) Get() *OrderProvider {
	return v.value
}

func (v *NullableOrderProvider) Set(val *OrderProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderProvider(val *OrderProvider) *NullableOrderProvider {
	return &NullableOrderProvider{value: val, isSet: true}
}

func (v NullableOrderProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


