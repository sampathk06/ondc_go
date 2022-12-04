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

// Offer Describes an offer
type Offer struct {
	Id *string `json:"id,omitempty"`
	Descriptor *Descriptor `json:"descriptor,omitempty"`
	LocationIds []Id `json:"location_ids,omitempty"`
	CategoryIds []Id `json:"category_ids,omitempty"`
	ItemIds []Id `json:"item_ids,omitempty"`
	Time *Time `json:"time,omitempty"`
}

// NewOffer instantiates a new Offer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOffer() *Offer {
	this := Offer{}
	return &this
}

// NewOfferWithDefaults instantiates a new Offer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferWithDefaults() *Offer {
	this := Offer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Offer) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Offer) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Offer) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Offer) SetId(v string) {
	o.Id = &v
}

// GetDescriptor returns the Descriptor field value if set, zero value otherwise.
func (o *Offer) GetDescriptor() Descriptor {
	if o == nil || isNil(o.Descriptor) {
		var ret Descriptor
		return ret
	}
	return *o.Descriptor
}

// GetDescriptorOk returns a tuple with the Descriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Offer) GetDescriptorOk() (*Descriptor, bool) {
	if o == nil || isNil(o.Descriptor) {
    return nil, false
	}
	return o.Descriptor, true
}

// HasDescriptor returns a boolean if a field has been set.
func (o *Offer) HasDescriptor() bool {
	if o != nil && !isNil(o.Descriptor) {
		return true
	}

	return false
}

// SetDescriptor gets a reference to the given Descriptor and assigns it to the Descriptor field.
func (o *Offer) SetDescriptor(v Descriptor) {
	o.Descriptor = &v
}

// GetLocationIds returns the LocationIds field value if set, zero value otherwise.
func (o *Offer) GetLocationIds() []Id {
	if o == nil || isNil(o.LocationIds) {
		var ret []Id
		return ret
	}
	return o.LocationIds
}

// GetLocationIdsOk returns a tuple with the LocationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Offer) GetLocationIdsOk() ([]Id, bool) {
	if o == nil || isNil(o.LocationIds) {
    return nil, false
	}
	return o.LocationIds, true
}

// HasLocationIds returns a boolean if a field has been set.
func (o *Offer) HasLocationIds() bool {
	if o != nil && !isNil(o.LocationIds) {
		return true
	}

	return false
}

// SetLocationIds gets a reference to the given []Id and assigns it to the LocationIds field.
func (o *Offer) SetLocationIds(v []Id) {
	o.LocationIds = v
}

// GetCategoryIds returns the CategoryIds field value if set, zero value otherwise.
func (o *Offer) GetCategoryIds() []Id {
	if o == nil || isNil(o.CategoryIds) {
		var ret []Id
		return ret
	}
	return o.CategoryIds
}

// GetCategoryIdsOk returns a tuple with the CategoryIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Offer) GetCategoryIdsOk() ([]Id, bool) {
	if o == nil || isNil(o.CategoryIds) {
    return nil, false
	}
	return o.CategoryIds, true
}

// HasCategoryIds returns a boolean if a field has been set.
func (o *Offer) HasCategoryIds() bool {
	if o != nil && !isNil(o.CategoryIds) {
		return true
	}

	return false
}

// SetCategoryIds gets a reference to the given []Id and assigns it to the CategoryIds field.
func (o *Offer) SetCategoryIds(v []Id) {
	o.CategoryIds = v
}

// GetItemIds returns the ItemIds field value if set, zero value otherwise.
func (o *Offer) GetItemIds() []Id {
	if o == nil || isNil(o.ItemIds) {
		var ret []Id
		return ret
	}
	return o.ItemIds
}

// GetItemIdsOk returns a tuple with the ItemIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Offer) GetItemIdsOk() ([]Id, bool) {
	if o == nil || isNil(o.ItemIds) {
    return nil, false
	}
	return o.ItemIds, true
}

// HasItemIds returns a boolean if a field has been set.
func (o *Offer) HasItemIds() bool {
	if o != nil && !isNil(o.ItemIds) {
		return true
	}

	return false
}

// SetItemIds gets a reference to the given []Id and assigns it to the ItemIds field.
func (o *Offer) SetItemIds(v []Id) {
	o.ItemIds = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *Offer) GetTime() Time {
	if o == nil || isNil(o.Time) {
		var ret Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Offer) GetTimeOk() (*Time, bool) {
	if o == nil || isNil(o.Time) {
    return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *Offer) HasTime() bool {
	if o != nil && !isNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given Time and assigns it to the Time field.
func (o *Offer) SetTime(v Time) {
	o.Time = &v
}

func (o Offer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Descriptor) {
		toSerialize["descriptor"] = o.Descriptor
	}
	if !isNil(o.LocationIds) {
		toSerialize["location_ids"] = o.LocationIds
	}
	if !isNil(o.CategoryIds) {
		toSerialize["category_ids"] = o.CategoryIds
	}
	if !isNil(o.ItemIds) {
		toSerialize["item_ids"] = o.ItemIds
	}
	if !isNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	return json.Marshal(toSerialize)
}

type NullableOffer struct {
	value *Offer
	isSet bool
}

func (v NullableOffer) Get() *Offer {
	return v.value
}

func (v *NullableOffer) Set(val *Offer) {
	v.value = val
	v.isSet = true
}

func (v NullableOffer) IsSet() bool {
	return v.isSet
}

func (v *NullableOffer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOffer(val *Offer) *NullableOffer {
	return &NullableOffer{value: val, isSet: true}
}

func (v NullableOffer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOffer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


