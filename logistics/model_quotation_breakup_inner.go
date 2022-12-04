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

// QuotationBreakupInner struct for QuotationBreakupInner
type QuotationBreakupInner struct {
	OndcOrgItemId *Id `json:"@ondc/org/item_id,omitempty"`
	OndcOrgTitleType *string `json:"@ondc/org/title_type,omitempty"`
	Title *string `json:"title,omitempty"`
	Price *Price `json:"price,omitempty"`
}

// NewQuotationBreakupInner instantiates a new QuotationBreakupInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotationBreakupInner() *QuotationBreakupInner {
	this := QuotationBreakupInner{}
	return &this
}

// NewQuotationBreakupInnerWithDefaults instantiates a new QuotationBreakupInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotationBreakupInnerWithDefaults() *QuotationBreakupInner {
	this := QuotationBreakupInner{}
	return &this
}

// GetOndcOrgItemId returns the OndcOrgItemId field value if set, zero value otherwise.
func (o *QuotationBreakupInner) GetOndcOrgItemId() Id {
	if o == nil || isNil(o.OndcOrgItemId) {
		var ret Id
		return ret
	}
	return *o.OndcOrgItemId
}

// GetOndcOrgItemIdOk returns a tuple with the OndcOrgItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotationBreakupInner) GetOndcOrgItemIdOk() (*Id, bool) {
	if o == nil || isNil(o.OndcOrgItemId) {
    return nil, false
	}
	return o.OndcOrgItemId, true
}

// HasOndcOrgItemId returns a boolean if a field has been set.
func (o *QuotationBreakupInner) HasOndcOrgItemId() bool {
	if o != nil && !isNil(o.OndcOrgItemId) {
		return true
	}

	return false
}

// SetOndcOrgItemId gets a reference to the given Id and assigns it to the OndcOrgItemId field.
func (o *QuotationBreakupInner) SetOndcOrgItemId(v Id) {
	o.OndcOrgItemId = &v
}

// GetOndcOrgTitleType returns the OndcOrgTitleType field value if set, zero value otherwise.
func (o *QuotationBreakupInner) GetOndcOrgTitleType() string {
	if o == nil || isNil(o.OndcOrgTitleType) {
		var ret string
		return ret
	}
	return *o.OndcOrgTitleType
}

// GetOndcOrgTitleTypeOk returns a tuple with the OndcOrgTitleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotationBreakupInner) GetOndcOrgTitleTypeOk() (*string, bool) {
	if o == nil || isNil(o.OndcOrgTitleType) {
    return nil, false
	}
	return o.OndcOrgTitleType, true
}

// HasOndcOrgTitleType returns a boolean if a field has been set.
func (o *QuotationBreakupInner) HasOndcOrgTitleType() bool {
	if o != nil && !isNil(o.OndcOrgTitleType) {
		return true
	}

	return false
}

// SetOndcOrgTitleType gets a reference to the given string and assigns it to the OndcOrgTitleType field.
func (o *QuotationBreakupInner) SetOndcOrgTitleType(v string) {
	o.OndcOrgTitleType = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *QuotationBreakupInner) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotationBreakupInner) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
    return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *QuotationBreakupInner) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *QuotationBreakupInner) SetTitle(v string) {
	o.Title = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *QuotationBreakupInner) GetPrice() Price {
	if o == nil || isNil(o.Price) {
		var ret Price
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotationBreakupInner) GetPriceOk() (*Price, bool) {
	if o == nil || isNil(o.Price) {
    return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *QuotationBreakupInner) HasPrice() bool {
	if o != nil && !isNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given Price and assigns it to the Price field.
func (o *QuotationBreakupInner) SetPrice(v Price) {
	o.Price = &v
}

func (o QuotationBreakupInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OndcOrgItemId) {
		toSerialize["@ondc/org/item_id"] = o.OndcOrgItemId
	}
	if !isNil(o.OndcOrgTitleType) {
		toSerialize["@ondc/org/title_type"] = o.OndcOrgTitleType
	}
	if !isNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !isNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	return json.Marshal(toSerialize)
}

type NullableQuotationBreakupInner struct {
	value *QuotationBreakupInner
	isSet bool
}

func (v NullableQuotationBreakupInner) Get() *QuotationBreakupInner {
	return v.value
}

func (v *NullableQuotationBreakupInner) Set(val *QuotationBreakupInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotationBreakupInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotationBreakupInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotationBreakupInner(val *QuotationBreakupInner) *NullableQuotationBreakupInner {
	return &NullableQuotationBreakupInner{value: val, isSet: true}
}

func (v NullableQuotationBreakupInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotationBreakupInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

