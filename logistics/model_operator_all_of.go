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

// OperatorAllOf struct for OperatorAllOf
type OperatorAllOf struct {
	Experience *OperatorAllOfExperience `json:"experience,omitempty"`
}

// NewOperatorAllOf instantiates a new OperatorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperatorAllOf() *OperatorAllOf {
	this := OperatorAllOf{}
	return &this
}

// NewOperatorAllOfWithDefaults instantiates a new OperatorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperatorAllOfWithDefaults() *OperatorAllOf {
	this := OperatorAllOf{}
	return &this
}

// GetExperience returns the Experience field value if set, zero value otherwise.
func (o *OperatorAllOf) GetExperience() OperatorAllOfExperience {
	if o == nil || isNil(o.Experience) {
		var ret OperatorAllOfExperience
		return ret
	}
	return *o.Experience
}

// GetExperienceOk returns a tuple with the Experience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorAllOf) GetExperienceOk() (*OperatorAllOfExperience, bool) {
	if o == nil || isNil(o.Experience) {
    return nil, false
	}
	return o.Experience, true
}

// HasExperience returns a boolean if a field has been set.
func (o *OperatorAllOf) HasExperience() bool {
	if o != nil && !isNil(o.Experience) {
		return true
	}

	return false
}

// SetExperience gets a reference to the given OperatorAllOfExperience and assigns it to the Experience field.
func (o *OperatorAllOf) SetExperience(v OperatorAllOfExperience) {
	o.Experience = &v
}

func (o OperatorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Experience) {
		toSerialize["experience"] = o.Experience
	}
	return json.Marshal(toSerialize)
}

type NullableOperatorAllOf struct {
	value *OperatorAllOf
	isSet bool
}

func (v NullableOperatorAllOf) Get() *OperatorAllOf {
	return v.value
}

func (v *NullableOperatorAllOf) Set(val *OperatorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatorAllOf(val *OperatorAllOf) *NullableOperatorAllOf {
	return &NullableOperatorAllOf{value: val, isSet: true}
}

func (v NullableOperatorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperatorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


