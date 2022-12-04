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

// Category Describes a category
type Category struct {
	// Unique id of the category
	Id *string `json:"id,omitempty"`
	// Unique id of the parent category
	ParentCategoryId *string `json:"parent_category_id,omitempty"`
	Descriptor *Descriptor `json:"descriptor,omitempty"`
	Time *Time `json:"time,omitempty"`
	// Describes a tag. This is a simple key-value store which is used to contain extended metadata
	Tags *map[string]string `json:"tags,omitempty"`
}

// NewCategory instantiates a new Category object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategory() *Category {
	this := Category{}
	return &this
}

// NewCategoryWithDefaults instantiates a new Category object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryWithDefaults() *Category {
	this := Category{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Category) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Category) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Category) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Category) SetId(v string) {
	o.Id = &v
}

// GetParentCategoryId returns the ParentCategoryId field value if set, zero value otherwise.
func (o *Category) GetParentCategoryId() string {
	if o == nil || isNil(o.ParentCategoryId) {
		var ret string
		return ret
	}
	return *o.ParentCategoryId
}

// GetParentCategoryIdOk returns a tuple with the ParentCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Category) GetParentCategoryIdOk() (*string, bool) {
	if o == nil || isNil(o.ParentCategoryId) {
    return nil, false
	}
	return o.ParentCategoryId, true
}

// HasParentCategoryId returns a boolean if a field has been set.
func (o *Category) HasParentCategoryId() bool {
	if o != nil && !isNil(o.ParentCategoryId) {
		return true
	}

	return false
}

// SetParentCategoryId gets a reference to the given string and assigns it to the ParentCategoryId field.
func (o *Category) SetParentCategoryId(v string) {
	o.ParentCategoryId = &v
}

// GetDescriptor returns the Descriptor field value if set, zero value otherwise.
func (o *Category) GetDescriptor() Descriptor {
	if o == nil || isNil(o.Descriptor) {
		var ret Descriptor
		return ret
	}
	return *o.Descriptor
}

// GetDescriptorOk returns a tuple with the Descriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Category) GetDescriptorOk() (*Descriptor, bool) {
	if o == nil || isNil(o.Descriptor) {
    return nil, false
	}
	return o.Descriptor, true
}

// HasDescriptor returns a boolean if a field has been set.
func (o *Category) HasDescriptor() bool {
	if o != nil && !isNil(o.Descriptor) {
		return true
	}

	return false
}

// SetDescriptor gets a reference to the given Descriptor and assigns it to the Descriptor field.
func (o *Category) SetDescriptor(v Descriptor) {
	o.Descriptor = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *Category) GetTime() Time {
	if o == nil || isNil(o.Time) {
		var ret Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Category) GetTimeOk() (*Time, bool) {
	if o == nil || isNil(o.Time) {
    return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *Category) HasTime() bool {
	if o != nil && !isNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given Time and assigns it to the Time field.
func (o *Category) SetTime(v Time) {
	o.Time = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Category) GetTags() map[string]string {
	if o == nil || isNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Category) GetTagsOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Category) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *Category) SetTags(v map[string]string) {
	o.Tags = &v
}

func (o Category) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.ParentCategoryId) {
		toSerialize["parent_category_id"] = o.ParentCategoryId
	}
	if !isNil(o.Descriptor) {
		toSerialize["descriptor"] = o.Descriptor
	}
	if !isNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableCategory struct {
	value *Category
	isSet bool
}

func (v NullableCategory) Get() *Category {
	return v.value
}

func (v *NullableCategory) Set(val *Category) {
	v.value = val
	v.isSet = true
}

func (v NullableCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategory(val *Category) *NullableCategory {
	return &NullableCategory{value: val, isSet: true}
}

func (v NullableCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


