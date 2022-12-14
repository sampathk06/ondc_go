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

// Scalar An object representing a scalar quantity.
type Scalar struct {
	Type *string `json:"type,omitempty"`
	Value float32 `json:"value"`
	EstimatedValue *float32 `json:"estimated_value,omitempty"`
	ComputedValue *float32 `json:"computed_value,omitempty"`
	Range *ScalarRange `json:"range,omitempty"`
	Unit string `json:"unit"`
}

// NewScalar instantiates a new Scalar object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScalar(value float32, unit string) *Scalar {
	this := Scalar{}
	this.Value = value
	this.Unit = unit
	return &this
}

// NewScalarWithDefaults instantiates a new Scalar object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScalarWithDefaults() *Scalar {
	this := Scalar{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Scalar) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scalar) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Scalar) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Scalar) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value
func (o *Scalar) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Scalar) GetValueOk() (*float32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Scalar) SetValue(v float32) {
	o.Value = v
}

// GetEstimatedValue returns the EstimatedValue field value if set, zero value otherwise.
func (o *Scalar) GetEstimatedValue() float32 {
	if o == nil || isNil(o.EstimatedValue) {
		var ret float32
		return ret
	}
	return *o.EstimatedValue
}

// GetEstimatedValueOk returns a tuple with the EstimatedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scalar) GetEstimatedValueOk() (*float32, bool) {
	if o == nil || isNil(o.EstimatedValue) {
    return nil, false
	}
	return o.EstimatedValue, true
}

// HasEstimatedValue returns a boolean if a field has been set.
func (o *Scalar) HasEstimatedValue() bool {
	if o != nil && !isNil(o.EstimatedValue) {
		return true
	}

	return false
}

// SetEstimatedValue gets a reference to the given float32 and assigns it to the EstimatedValue field.
func (o *Scalar) SetEstimatedValue(v float32) {
	o.EstimatedValue = &v
}

// GetComputedValue returns the ComputedValue field value if set, zero value otherwise.
func (o *Scalar) GetComputedValue() float32 {
	if o == nil || isNil(o.ComputedValue) {
		var ret float32
		return ret
	}
	return *o.ComputedValue
}

// GetComputedValueOk returns a tuple with the ComputedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scalar) GetComputedValueOk() (*float32, bool) {
	if o == nil || isNil(o.ComputedValue) {
    return nil, false
	}
	return o.ComputedValue, true
}

// HasComputedValue returns a boolean if a field has been set.
func (o *Scalar) HasComputedValue() bool {
	if o != nil && !isNil(o.ComputedValue) {
		return true
	}

	return false
}

// SetComputedValue gets a reference to the given float32 and assigns it to the ComputedValue field.
func (o *Scalar) SetComputedValue(v float32) {
	o.ComputedValue = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *Scalar) GetRange() ScalarRange {
	if o == nil || isNil(o.Range) {
		var ret ScalarRange
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scalar) GetRangeOk() (*ScalarRange, bool) {
	if o == nil || isNil(o.Range) {
    return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *Scalar) HasRange() bool {
	if o != nil && !isNil(o.Range) {
		return true
	}

	return false
}

// SetRange gets a reference to the given ScalarRange and assigns it to the Range field.
func (o *Scalar) SetRange(v ScalarRange) {
	o.Range = &v
}

// GetUnit returns the Unit field value
func (o *Scalar) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *Scalar) GetUnitOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *Scalar) SetUnit(v string) {
	o.Unit = v
}

func (o Scalar) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if !isNil(o.EstimatedValue) {
		toSerialize["estimated_value"] = o.EstimatedValue
	}
	if !isNil(o.ComputedValue) {
		toSerialize["computed_value"] = o.ComputedValue
	}
	if !isNil(o.Range) {
		toSerialize["range"] = o.Range
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableScalar struct {
	value *Scalar
	isSet bool
}

func (v NullableScalar) Get() *Scalar {
	return v.value
}

func (v *NullableScalar) Set(val *Scalar) {
	v.value = val
	v.isSet = true
}

func (v NullableScalar) IsSet() bool {
	return v.isSet
}

func (v *NullableScalar) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScalar(val *Scalar) *NullableScalar {
	return &NullableScalar{value: val, isSet: true}
}

func (v NullableScalar) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScalar) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


