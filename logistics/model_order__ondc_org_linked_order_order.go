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

// OrderOndcOrgLinkedOrderOrder use same units for weight and dimensions as defined for Intent
type OrderOndcOrgLinkedOrderOrder struct {
	Id *Id `json:"id,omitempty"`
	Weight *Scalar `json:"weight,omitempty"`
	Dimensions *Dimensions `json:"dimensions,omitempty"`
	DeclaredValue *Price `json:"declared_value,omitempty"`
	TaxableValue *Price `json:"taxable_value,omitempty"`
	HsnCode *string `json:"hsn_code,omitempty"`
	// Describes a decimal value
	SgstAmount *string `json:"sgst_amount,omitempty"`
	// Describes a decimal value
	CgstAmount *string `json:"cgst_amount,omitempty"`
	// Describes a decimal value
	IgstAmount *string `json:"igst_amount,omitempty"`
}

// NewOrderOndcOrgLinkedOrderOrder instantiates a new OrderOndcOrgLinkedOrderOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderOndcOrgLinkedOrderOrder() *OrderOndcOrgLinkedOrderOrder {
	this := OrderOndcOrgLinkedOrderOrder{}
	return &this
}

// NewOrderOndcOrgLinkedOrderOrderWithDefaults instantiates a new OrderOndcOrgLinkedOrderOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderOndcOrgLinkedOrderOrderWithDefaults() *OrderOndcOrgLinkedOrderOrder {
	this := OrderOndcOrgLinkedOrderOrder{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderOndcOrgLinkedOrderOrder) GetId() Id {
	if o == nil || isNil(o.Id) {
		var ret Id
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOndcOrgLinkedOrderOrder) GetIdOk() (*Id, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderOndcOrgLinkedOrderOrder) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given Id and assigns it to the Id field.
func (o *OrderOndcOrgLinkedOrderOrder) SetId(v Id) {
	o.Id = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *OrderOndcOrgLinkedOrderOrder) GetWeight() Scalar {
	if o == nil || isNil(o.Weight) {
		var ret Scalar
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOndcOrgLinkedOrderOrder) GetWeightOk() (*Scalar, bool) {
	if o == nil || isNil(o.Weight) {
    return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *OrderOndcOrgLinkedOrderOrder) HasWeight() bool {
	if o != nil && !isNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given Scalar and assigns it to the Weight field.
func (o *OrderOndcOrgLinkedOrderOrder) SetWeight(v Scalar) {
	o.Weight = &v
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *OrderOndcOrgLinkedOrderOrder) GetDimensions() Dimensions {
	if o == nil || isNil(o.Dimensions) {
		var ret Dimensions
		return ret
	}
	return *o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOndcOrgLinkedOrderOrder) GetDimensionsOk() (*Dimensions, bool) {
	if o == nil || isNil(o.Dimensions) {
    return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *OrderOndcOrgLinkedOrderOrder) HasDimensions() bool {
	if o != nil && !isNil(o.Dimensions) {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given Dimensions and assigns it to the Dimensions field.
func (o *OrderOndcOrgLinkedOrderOrder) SetDimensions(v Dimensions) {
	o.Dimensions = &v
}

// GetDeclaredValue returns the DeclaredValue field value if set, zero value otherwise.
func (o *OrderOndcOrgLinkedOrderOrder) GetDeclaredValue() Price {
	if o == nil || isNil(o.DeclaredValue) {
		var ret Price
		return ret
	}
	return *o.DeclaredValue
}

// GetDeclaredValueOk returns a tuple with the DeclaredValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOndcOrgLinkedOrderOrder) GetDeclaredValueOk() (*Price, bool) {
	if o == nil || isNil(o.DeclaredValue) {
    return nil, false
	}
	return o.DeclaredValue, true
}

// HasDeclaredValue returns a boolean if a field has been set.
func (o *OrderOndcOrgLinkedOrderOrder) HasDeclaredValue() bool {
	if o != nil && !isNil(o.DeclaredValue) {
		return true
	}

	return false
}

// SetDeclaredValue gets a reference to the given Price and assigns it to the DeclaredValue field.
func (o *OrderOndcOrgLinkedOrderOrder) SetDeclaredValue(v Price) {
	o.DeclaredValue = &v
}

// GetTaxableValue returns the TaxableValue field value if set, zero value otherwise.
func (o *OrderOndcOrgLinkedOrderOrder) GetTaxableValue() Price {
	if o == nil || isNil(o.TaxableValue) {
		var ret Price
		return ret
	}
	return *o.TaxableValue
}

// GetTaxableValueOk returns a tuple with the TaxableValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOndcOrgLinkedOrderOrder) GetTaxableValueOk() (*Price, bool) {
	if o == nil || isNil(o.TaxableValue) {
    return nil, false
	}
	return o.TaxableValue, true
}

// HasTaxableValue returns a boolean if a field has been set.
func (o *OrderOndcOrgLinkedOrderOrder) HasTaxableValue() bool {
	if o != nil && !isNil(o.TaxableValue) {
		return true
	}

	return false
}

// SetTaxableValue gets a reference to the given Price and assigns it to the TaxableValue field.
func (o *OrderOndcOrgLinkedOrderOrder) SetTaxableValue(v Price) {
	o.TaxableValue = &v
}

// GetHsnCode returns the HsnCode field value if set, zero value otherwise.
func (o *OrderOndcOrgLinkedOrderOrder) GetHsnCode() string {
	if o == nil || isNil(o.HsnCode) {
		var ret string
		return ret
	}
	return *o.HsnCode
}

// GetHsnCodeOk returns a tuple with the HsnCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOndcOrgLinkedOrderOrder) GetHsnCodeOk() (*string, bool) {
	if o == nil || isNil(o.HsnCode) {
    return nil, false
	}
	return o.HsnCode, true
}

// HasHsnCode returns a boolean if a field has been set.
func (o *OrderOndcOrgLinkedOrderOrder) HasHsnCode() bool {
	if o != nil && !isNil(o.HsnCode) {
		return true
	}

	return false
}

// SetHsnCode gets a reference to the given string and assigns it to the HsnCode field.
func (o *OrderOndcOrgLinkedOrderOrder) SetHsnCode(v string) {
	o.HsnCode = &v
}

// GetSgstAmount returns the SgstAmount field value if set, zero value otherwise.
func (o *OrderOndcOrgLinkedOrderOrder) GetSgstAmount() string {
	if o == nil || isNil(o.SgstAmount) {
		var ret string
		return ret
	}
	return *o.SgstAmount
}

// GetSgstAmountOk returns a tuple with the SgstAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOndcOrgLinkedOrderOrder) GetSgstAmountOk() (*string, bool) {
	if o == nil || isNil(o.SgstAmount) {
    return nil, false
	}
	return o.SgstAmount, true
}

// HasSgstAmount returns a boolean if a field has been set.
func (o *OrderOndcOrgLinkedOrderOrder) HasSgstAmount() bool {
	if o != nil && !isNil(o.SgstAmount) {
		return true
	}

	return false
}

// SetSgstAmount gets a reference to the given string and assigns it to the SgstAmount field.
func (o *OrderOndcOrgLinkedOrderOrder) SetSgstAmount(v string) {
	o.SgstAmount = &v
}

// GetCgstAmount returns the CgstAmount field value if set, zero value otherwise.
func (o *OrderOndcOrgLinkedOrderOrder) GetCgstAmount() string {
	if o == nil || isNil(o.CgstAmount) {
		var ret string
		return ret
	}
	return *o.CgstAmount
}

// GetCgstAmountOk returns a tuple with the CgstAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOndcOrgLinkedOrderOrder) GetCgstAmountOk() (*string, bool) {
	if o == nil || isNil(o.CgstAmount) {
    return nil, false
	}
	return o.CgstAmount, true
}

// HasCgstAmount returns a boolean if a field has been set.
func (o *OrderOndcOrgLinkedOrderOrder) HasCgstAmount() bool {
	if o != nil && !isNil(o.CgstAmount) {
		return true
	}

	return false
}

// SetCgstAmount gets a reference to the given string and assigns it to the CgstAmount field.
func (o *OrderOndcOrgLinkedOrderOrder) SetCgstAmount(v string) {
	o.CgstAmount = &v
}

// GetIgstAmount returns the IgstAmount field value if set, zero value otherwise.
func (o *OrderOndcOrgLinkedOrderOrder) GetIgstAmount() string {
	if o == nil || isNil(o.IgstAmount) {
		var ret string
		return ret
	}
	return *o.IgstAmount
}

// GetIgstAmountOk returns a tuple with the IgstAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderOndcOrgLinkedOrderOrder) GetIgstAmountOk() (*string, bool) {
	if o == nil || isNil(o.IgstAmount) {
    return nil, false
	}
	return o.IgstAmount, true
}

// HasIgstAmount returns a boolean if a field has been set.
func (o *OrderOndcOrgLinkedOrderOrder) HasIgstAmount() bool {
	if o != nil && !isNil(o.IgstAmount) {
		return true
	}

	return false
}

// SetIgstAmount gets a reference to the given string and assigns it to the IgstAmount field.
func (o *OrderOndcOrgLinkedOrderOrder) SetIgstAmount(v string) {
	o.IgstAmount = &v
}

func (o OrderOndcOrgLinkedOrderOrder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !isNil(o.Dimensions) {
		toSerialize["dimensions"] = o.Dimensions
	}
	if !isNil(o.DeclaredValue) {
		toSerialize["declared_value"] = o.DeclaredValue
	}
	if !isNil(o.TaxableValue) {
		toSerialize["taxable_value"] = o.TaxableValue
	}
	if !isNil(o.HsnCode) {
		toSerialize["hsn_code"] = o.HsnCode
	}
	if !isNil(o.SgstAmount) {
		toSerialize["sgst_amount"] = o.SgstAmount
	}
	if !isNil(o.CgstAmount) {
		toSerialize["cgst_amount"] = o.CgstAmount
	}
	if !isNil(o.IgstAmount) {
		toSerialize["igst_amount"] = o.IgstAmount
	}
	return json.Marshal(toSerialize)
}

type NullableOrderOndcOrgLinkedOrderOrder struct {
	value *OrderOndcOrgLinkedOrderOrder
	isSet bool
}

func (v NullableOrderOndcOrgLinkedOrderOrder) Get() *OrderOndcOrgLinkedOrderOrder {
	return v.value
}

func (v *NullableOrderOndcOrgLinkedOrderOrder) Set(val *OrderOndcOrgLinkedOrderOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderOndcOrgLinkedOrderOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderOndcOrgLinkedOrderOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderOndcOrgLinkedOrderOrder(val *OrderOndcOrgLinkedOrderOrder) *NullableOrderOndcOrgLinkedOrderOrder {
	return &NullableOrderOndcOrgLinkedOrderOrder{value: val, isSet: true}
}

func (v NullableOrderOndcOrgLinkedOrderOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderOndcOrgLinkedOrderOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

