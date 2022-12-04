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

// OperatorAllOfExperience struct for OperatorAllOfExperience
type OperatorAllOfExperience struct {
	Label *string `json:"label,omitempty"`
	Value *string `json:"value,omitempty"`
	Unit *string `json:"unit,omitempty"`
}

// NewOperatorAllOfExperience instantiates a new OperatorAllOfExperience object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperatorAllOfExperience() *OperatorAllOfExperience {
	this := OperatorAllOfExperience{}
	return &this
}

// NewOperatorAllOfExperienceWithDefaults instantiates a new OperatorAllOfExperience object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperatorAllOfExperienceWithDefaults() *OperatorAllOfExperience {
	this := OperatorAllOfExperience{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *OperatorAllOfExperience) GetLabel() string {
	if o == nil || isNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorAllOfExperience) GetLabelOk() (*string, bool) {
	if o == nil || isNil(o.Label) {
    return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *OperatorAllOfExperience) HasLabel() bool {
	if o != nil && !isNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *OperatorAllOfExperience) SetLabel(v string) {
	o.Label = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *OperatorAllOfExperience) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorAllOfExperience) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *OperatorAllOfExperience) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *OperatorAllOfExperience) SetValue(v string) {
	o.Value = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *OperatorAllOfExperience) GetUnit() string {
	if o == nil || isNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorAllOfExperience) GetUnitOk() (*string, bool) {
	if o == nil || isNil(o.Unit) {
    return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *OperatorAllOfExperience) HasUnit() bool {
	if o != nil && !isNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *OperatorAllOfExperience) SetUnit(v string) {
	o.Unit = &v
}

func (o OperatorAllOfExperience) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !isNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableOperatorAllOfExperience struct {
	value *OperatorAllOfExperience
	isSet bool
}

func (v NullableOperatorAllOfExperience) Get() *OperatorAllOfExperience {
	return v.value
}

func (v *NullableOperatorAllOfExperience) Set(val *OperatorAllOfExperience) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatorAllOfExperience) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatorAllOfExperience) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatorAllOfExperience(val *OperatorAllOfExperience) *NullableOperatorAllOfExperience {
	return &NullableOperatorAllOfExperience{value: val, isSet: true}
}

func (v NullableOperatorAllOfExperience) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperatorAllOfExperience) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


