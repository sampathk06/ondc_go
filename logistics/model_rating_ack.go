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

// RatingAck struct for RatingAck
type RatingAck struct {
	// If feedback has been recorded or not
	FeedbackAck *bool `json:"feedback_ack,omitempty"`
	// If rating has been recorded or not
	RatingAck *bool `json:"rating_ack,omitempty"`
}

// NewRatingAck instantiates a new RatingAck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRatingAck() *RatingAck {
	this := RatingAck{}
	return &this
}

// NewRatingAckWithDefaults instantiates a new RatingAck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRatingAckWithDefaults() *RatingAck {
	this := RatingAck{}
	return &this
}

// GetFeedbackAck returns the FeedbackAck field value if set, zero value otherwise.
func (o *RatingAck) GetFeedbackAck() bool {
	if o == nil || isNil(o.FeedbackAck) {
		var ret bool
		return ret
	}
	return *o.FeedbackAck
}

// GetFeedbackAckOk returns a tuple with the FeedbackAck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatingAck) GetFeedbackAckOk() (*bool, bool) {
	if o == nil || isNil(o.FeedbackAck) {
    return nil, false
	}
	return o.FeedbackAck, true
}

// HasFeedbackAck returns a boolean if a field has been set.
func (o *RatingAck) HasFeedbackAck() bool {
	if o != nil && !isNil(o.FeedbackAck) {
		return true
	}

	return false
}

// SetFeedbackAck gets a reference to the given bool and assigns it to the FeedbackAck field.
func (o *RatingAck) SetFeedbackAck(v bool) {
	o.FeedbackAck = &v
}

// GetRatingAck returns the RatingAck field value if set, zero value otherwise.
func (o *RatingAck) GetRatingAck() bool {
	if o == nil || isNil(o.RatingAck) {
		var ret bool
		return ret
	}
	return *o.RatingAck
}

// GetRatingAckOk returns a tuple with the RatingAck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RatingAck) GetRatingAckOk() (*bool, bool) {
	if o == nil || isNil(o.RatingAck) {
    return nil, false
	}
	return o.RatingAck, true
}

// HasRatingAck returns a boolean if a field has been set.
func (o *RatingAck) HasRatingAck() bool {
	if o != nil && !isNil(o.RatingAck) {
		return true
	}

	return false
}

// SetRatingAck gets a reference to the given bool and assigns it to the RatingAck field.
func (o *RatingAck) SetRatingAck(v bool) {
	o.RatingAck = &v
}

func (o RatingAck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FeedbackAck) {
		toSerialize["feedback_ack"] = o.FeedbackAck
	}
	if !isNil(o.RatingAck) {
		toSerialize["rating_ack"] = o.RatingAck
	}
	return json.Marshal(toSerialize)
}

type NullableRatingAck struct {
	value *RatingAck
	isSet bool
}

func (v NullableRatingAck) Get() *RatingAck {
	return v.value
}

func (v *NullableRatingAck) Set(val *RatingAck) {
	v.value = val
	v.isSet = true
}

func (v NullableRatingAck) IsSet() bool {
	return v.isSet
}

func (v *NullableRatingAck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRatingAck(val *RatingAck) *NullableRatingAck {
	return &NullableRatingAck{value: val, isSet: true}
}

func (v NullableRatingAck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRatingAck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


