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

// Vehicle Describes the properties of a vehicle used in a mobility service
type Vehicle struct {
	Category *string `json:"category,omitempty"`
	Capacity *int32 `json:"capacity,omitempty"`
	Make *string `json:"make,omitempty"`
	Model *string `json:"model,omitempty"`
	Size *string `json:"size,omitempty"`
	Variant *string `json:"variant,omitempty"`
	Color *string `json:"color,omitempty"`
	EnergyType *string `json:"energy_type,omitempty"`
	Registration *string `json:"registration,omitempty"`
}

// NewVehicle instantiates a new Vehicle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVehicle() *Vehicle {
	this := Vehicle{}
	return &this
}

// NewVehicleWithDefaults instantiates a new Vehicle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVehicleWithDefaults() *Vehicle {
	this := Vehicle{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *Vehicle) GetCategory() string {
	if o == nil || isNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetCategoryOk() (*string, bool) {
	if o == nil || isNil(o.Category) {
    return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *Vehicle) HasCategory() bool {
	if o != nil && !isNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *Vehicle) SetCategory(v string) {
	o.Category = &v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *Vehicle) GetCapacity() int32 {
	if o == nil || isNil(o.Capacity) {
		var ret int32
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetCapacityOk() (*int32, bool) {
	if o == nil || isNil(o.Capacity) {
    return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *Vehicle) HasCapacity() bool {
	if o != nil && !isNil(o.Capacity) {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int32 and assigns it to the Capacity field.
func (o *Vehicle) SetCapacity(v int32) {
	o.Capacity = &v
}

// GetMake returns the Make field value if set, zero value otherwise.
func (o *Vehicle) GetMake() string {
	if o == nil || isNil(o.Make) {
		var ret string
		return ret
	}
	return *o.Make
}

// GetMakeOk returns a tuple with the Make field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetMakeOk() (*string, bool) {
	if o == nil || isNil(o.Make) {
    return nil, false
	}
	return o.Make, true
}

// HasMake returns a boolean if a field has been set.
func (o *Vehicle) HasMake() bool {
	if o != nil && !isNil(o.Make) {
		return true
	}

	return false
}

// SetMake gets a reference to the given string and assigns it to the Make field.
func (o *Vehicle) SetMake(v string) {
	o.Make = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *Vehicle) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *Vehicle) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *Vehicle) SetModel(v string) {
	o.Model = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Vehicle) GetSize() string {
	if o == nil || isNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetSizeOk() (*string, bool) {
	if o == nil || isNil(o.Size) {
    return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Vehicle) HasSize() bool {
	if o != nil && !isNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *Vehicle) SetSize(v string) {
	o.Size = &v
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *Vehicle) GetVariant() string {
	if o == nil || isNil(o.Variant) {
		var ret string
		return ret
	}
	return *o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetVariantOk() (*string, bool) {
	if o == nil || isNil(o.Variant) {
    return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *Vehicle) HasVariant() bool {
	if o != nil && !isNil(o.Variant) {
		return true
	}

	return false
}

// SetVariant gets a reference to the given string and assigns it to the Variant field.
func (o *Vehicle) SetVariant(v string) {
	o.Variant = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *Vehicle) GetColor() string {
	if o == nil || isNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetColorOk() (*string, bool) {
	if o == nil || isNil(o.Color) {
    return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *Vehicle) HasColor() bool {
	if o != nil && !isNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *Vehicle) SetColor(v string) {
	o.Color = &v
}

// GetEnergyType returns the EnergyType field value if set, zero value otherwise.
func (o *Vehicle) GetEnergyType() string {
	if o == nil || isNil(o.EnergyType) {
		var ret string
		return ret
	}
	return *o.EnergyType
}

// GetEnergyTypeOk returns a tuple with the EnergyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetEnergyTypeOk() (*string, bool) {
	if o == nil || isNil(o.EnergyType) {
    return nil, false
	}
	return o.EnergyType, true
}

// HasEnergyType returns a boolean if a field has been set.
func (o *Vehicle) HasEnergyType() bool {
	if o != nil && !isNil(o.EnergyType) {
		return true
	}

	return false
}

// SetEnergyType gets a reference to the given string and assigns it to the EnergyType field.
func (o *Vehicle) SetEnergyType(v string) {
	o.EnergyType = &v
}

// GetRegistration returns the Registration field value if set, zero value otherwise.
func (o *Vehicle) GetRegistration() string {
	if o == nil || isNil(o.Registration) {
		var ret string
		return ret
	}
	return *o.Registration
}

// GetRegistrationOk returns a tuple with the Registration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vehicle) GetRegistrationOk() (*string, bool) {
	if o == nil || isNil(o.Registration) {
    return nil, false
	}
	return o.Registration, true
}

// HasRegistration returns a boolean if a field has been set.
func (o *Vehicle) HasRegistration() bool {
	if o != nil && !isNil(o.Registration) {
		return true
	}

	return false
}

// SetRegistration gets a reference to the given string and assigns it to the Registration field.
func (o *Vehicle) SetRegistration(v string) {
	o.Registration = &v
}

func (o Vehicle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !isNil(o.Capacity) {
		toSerialize["capacity"] = o.Capacity
	}
	if !isNil(o.Make) {
		toSerialize["make"] = o.Make
	}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !isNil(o.Variant) {
		toSerialize["variant"] = o.Variant
	}
	if !isNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !isNil(o.EnergyType) {
		toSerialize["energy_type"] = o.EnergyType
	}
	if !isNil(o.Registration) {
		toSerialize["registration"] = o.Registration
	}
	return json.Marshal(toSerialize)
}

type NullableVehicle struct {
	value *Vehicle
	isSet bool
}

func (v NullableVehicle) Get() *Vehicle {
	return v.value
}

func (v *NullableVehicle) Set(val *Vehicle) {
	v.value = val
	v.isSet = true
}

func (v NullableVehicle) IsSet() bool {
	return v.isSet
}

func (v *NullableVehicle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVehicle(val *Vehicle) *NullableVehicle {
	return &NullableVehicle{value: val, isSet: true}
}

func (v NullableVehicle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVehicle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


