/*
ONDC Protocol Core API

ONDC Protocol Core API specification. This is an adaptation of Beckn Core version 0.9.3

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// GetFeedbackFormPostRequestMessage struct for GetFeedbackFormPostRequestMessage
type GetFeedbackFormPostRequestMessage struct {
	RatingValue *Value `json:"rating_value,omitempty"`
	RatingCategory *RatingCategory `json:"rating_category,omitempty"`
}

// NewGetFeedbackFormPostRequestMessage instantiates a new GetFeedbackFormPostRequestMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFeedbackFormPostRequestMessage() *GetFeedbackFormPostRequestMessage {
	this := GetFeedbackFormPostRequestMessage{}
	return &this
}

// NewGetFeedbackFormPostRequestMessageWithDefaults instantiates a new GetFeedbackFormPostRequestMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFeedbackFormPostRequestMessageWithDefaults() *GetFeedbackFormPostRequestMessage {
	this := GetFeedbackFormPostRequestMessage{}
	return &this
}

// GetRatingValue returns the RatingValue field value if set, zero value otherwise.
func (o *GetFeedbackFormPostRequestMessage) GetRatingValue() Value {
	if o == nil || isNil(o.RatingValue) {
		var ret Value
		return ret
	}
	return *o.RatingValue
}

// GetRatingValueOk returns a tuple with the RatingValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFeedbackFormPostRequestMessage) GetRatingValueOk() (*Value, bool) {
	if o == nil || isNil(o.RatingValue) {
    return nil, false
	}
	return o.RatingValue, true
}

// HasRatingValue returns a boolean if a field has been set.
func (o *GetFeedbackFormPostRequestMessage) HasRatingValue() bool {
	if o != nil && !isNil(o.RatingValue) {
		return true
	}

	return false
}

// SetRatingValue gets a reference to the given Value and assigns it to the RatingValue field.
func (o *GetFeedbackFormPostRequestMessage) SetRatingValue(v Value) {
	o.RatingValue = &v
}

// GetRatingCategory returns the RatingCategory field value if set, zero value otherwise.
func (o *GetFeedbackFormPostRequestMessage) GetRatingCategory() RatingCategory {
	if o == nil || isNil(o.RatingCategory) {
		var ret RatingCategory
		return ret
	}
	return *o.RatingCategory
}

// GetRatingCategoryOk returns a tuple with the RatingCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFeedbackFormPostRequestMessage) GetRatingCategoryOk() (*RatingCategory, bool) {
	if o == nil || isNil(o.RatingCategory) {
    return nil, false
	}
	return o.RatingCategory, true
}

// HasRatingCategory returns a boolean if a field has been set.
func (o *GetFeedbackFormPostRequestMessage) HasRatingCategory() bool {
	if o != nil && !isNil(o.RatingCategory) {
		return true
	}

	return false
}

// SetRatingCategory gets a reference to the given RatingCategory and assigns it to the RatingCategory field.
func (o *GetFeedbackFormPostRequestMessage) SetRatingCategory(v RatingCategory) {
	o.RatingCategory = &v
}

func (o GetFeedbackFormPostRequestMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RatingValue) {
		toSerialize["rating_value"] = o.RatingValue
	}
	if !isNil(o.RatingCategory) {
		toSerialize["rating_category"] = o.RatingCategory
	}
	return json.Marshal(toSerialize)
}

type NullableGetFeedbackFormPostRequestMessage struct {
	value *GetFeedbackFormPostRequestMessage
	isSet bool
}

func (v NullableGetFeedbackFormPostRequestMessage) Get() *GetFeedbackFormPostRequestMessage {
	return v.value
}

func (v *NullableGetFeedbackFormPostRequestMessage) Set(val *GetFeedbackFormPostRequestMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFeedbackFormPostRequestMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFeedbackFormPostRequestMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFeedbackFormPostRequestMessage(val *GetFeedbackFormPostRequestMessage) *NullableGetFeedbackFormPostRequestMessage {
	return &NullableGetFeedbackFormPostRequestMessage{value: val, isSet: true}
}

func (v NullableGetFeedbackFormPostRequestMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFeedbackFormPostRequestMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


