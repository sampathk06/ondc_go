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

// Support Customer support
type Support struct {
	Type *string `json:"type,omitempty"`
	RefId *string `json:"ref_id,omitempty"`
	// Describes a tag. This is a simple key-value store which is used to contain extended metadata
	Channels *map[string]string `json:"channels,omitempty"`
}

// NewSupport instantiates a new Support object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupport() *Support {
	this := Support{}
	return &this
}

// NewSupportWithDefaults instantiates a new Support object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportWithDefaults() *Support {
	this := Support{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Support) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Support) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Support) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Support) SetType(v string) {
	o.Type = &v
}

// GetRefId returns the RefId field value if set, zero value otherwise.
func (o *Support) GetRefId() string {
	if o == nil || isNil(o.RefId) {
		var ret string
		return ret
	}
	return *o.RefId
}

// GetRefIdOk returns a tuple with the RefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Support) GetRefIdOk() (*string, bool) {
	if o == nil || isNil(o.RefId) {
    return nil, false
	}
	return o.RefId, true
}

// HasRefId returns a boolean if a field has been set.
func (o *Support) HasRefId() bool {
	if o != nil && !isNil(o.RefId) {
		return true
	}

	return false
}

// SetRefId gets a reference to the given string and assigns it to the RefId field.
func (o *Support) SetRefId(v string) {
	o.RefId = &v
}

// GetChannels returns the Channels field value if set, zero value otherwise.
func (o *Support) GetChannels() map[string]string {
	if o == nil || isNil(o.Channels) {
		var ret map[string]string
		return ret
	}
	return *o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Support) GetChannelsOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Channels) {
    return nil, false
	}
	return o.Channels, true
}

// HasChannels returns a boolean if a field has been set.
func (o *Support) HasChannels() bool {
	if o != nil && !isNil(o.Channels) {
		return true
	}

	return false
}

// SetChannels gets a reference to the given map[string]string and assigns it to the Channels field.
func (o *Support) SetChannels(v map[string]string) {
	o.Channels = &v
}

func (o Support) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.RefId) {
		toSerialize["ref_id"] = o.RefId
	}
	if !isNil(o.Channels) {
		toSerialize["channels"] = o.Channels
	}
	return json.Marshal(toSerialize)
}

type NullableSupport struct {
	value *Support
	isSet bool
}

func (v NullableSupport) Get() *Support {
	return v.value
}

func (v *NullableSupport) Set(val *Support) {
	v.value = val
	v.isSet = true
}

func (v NullableSupport) IsSet() bool {
	return v.isSet
}

func (v *NullableSupport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupport(val *Support) *NullableSupport {
	return &NullableSupport{value: val, isSet: true}
}

func (v NullableSupport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

