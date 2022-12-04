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

// OrderOndcOrgLinkedOrderProvider struct for OrderOndcOrgLinkedOrderProvider
type OrderOndcOrgLinkedOrderProvider struct {
	Id *Id `json:"id,omitempty"`
	Descriptor *Descriptor `json:"descriptor,omitempty"`
	Address *Address `json:"address,omitempty"`
}

// NewOrderOndcOrgLinkedOrderProvider instantiates a new OrderOndcOrgLinkedOrderProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderOndcOrgLinkedOrderProvider() *OrderOndcOrgLinkedOrderProvider {
	this := OrderOndcOrgLinkedOrderProvider{}
	return &this
}

// NewOrderOndcOrgLinkedOrderProviderWithDefaults instantiates a new OrderOndcOrgLinkedOrderProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderOndcOrgLinkedOrderProviderWithDefaults() *OrderOndcOrgLinkedOrderProvider {
	this := OrderOndcOrgLinkedOrderProvider{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderOndcOrgLinkedOrderProvider) GetId() Id {
	if o == nil || isNil(o.Id) {
		var ret Id
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOndcOrgLinkedOrderProvider) GetIdOk() (*Id, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderOndcOrgLinkedOrderProvider) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given Id and assigns it to the Id field.
func (o *OrderOndcOrgLinkedOrderProvider) SetId(v Id) {
	o.Id = &v
}

// GetDescriptor returns the Descriptor field value if set, zero value otherwise.
func (o *OrderOndcOrgLinkedOrderProvider) GetDescriptor() Descriptor {
	if o == nil || isNil(o.Descriptor) {
		var ret Descriptor
		return ret
	}
	return *o.Descriptor
}

// GetDescriptorOk returns a tuple with the Descriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOndcOrgLinkedOrderProvider) GetDescriptorOk() (*Descriptor, bool) {
	if o == nil || isNil(o.Descriptor) {
    return nil, false
	}
	return o.Descriptor, true
}

// HasDescriptor returns a boolean if a field has been set.
func (o *OrderOndcOrgLinkedOrderProvider) HasDescriptor() bool {
	if o != nil && !isNil(o.Descriptor) {
		return true
	}

	return false
}

// SetDescriptor gets a reference to the given Descriptor and assigns it to the Descriptor field.
func (o *OrderOndcOrgLinkedOrderProvider) SetDescriptor(v Descriptor) {
	o.Descriptor = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *OrderOndcOrgLinkedOrderProvider) GetAddress() Address {
	if o == nil || isNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOndcOrgLinkedOrderProvider) GetAddressOk() (*Address, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *OrderOndcOrgLinkedOrderProvider) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *OrderOndcOrgLinkedOrderProvider) SetAddress(v Address) {
	o.Address = &v
}

func (o OrderOndcOrgLinkedOrderProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Descriptor) {
		toSerialize["descriptor"] = o.Descriptor
	}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableOrderOndcOrgLinkedOrderProvider struct {
	value *OrderOndcOrgLinkedOrderProvider
	isSet bool
}

func (v NullableOrderOndcOrgLinkedOrderProvider) Get() *OrderOndcOrgLinkedOrderProvider {
	return v.value
}

func (v *NullableOrderOndcOrgLinkedOrderProvider) Set(val *OrderOndcOrgLinkedOrderProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderOndcOrgLinkedOrderProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderOndcOrgLinkedOrderProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderOndcOrgLinkedOrderProvider(val *OrderOndcOrgLinkedOrderProvider) *NullableOrderOndcOrgLinkedOrderProvider {
	return &NullableOrderOndcOrgLinkedOrderProvider{value: val, isSet: true}
}

func (v NullableOrderOndcOrgLinkedOrderProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderOndcOrgLinkedOrderProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


