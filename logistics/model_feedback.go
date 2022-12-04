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

// Feedback Feedback for a service
type Feedback struct {
	// Describes a feedback form that a Seller App can send to get feedback from the Buyer App
	FeedbackForm []FeedbackFormElement `json:"feedback_form,omitempty"`
	FeedbackUrl *FeedbackUrl `json:"feedback_url,omitempty"`
}

// NewFeedback instantiates a new Feedback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedback() *Feedback {
	this := Feedback{}
	return &this
}

// NewFeedbackWithDefaults instantiates a new Feedback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedbackWithDefaults() *Feedback {
	this := Feedback{}
	return &this
}

// GetFeedbackForm returns the FeedbackForm field value if set, zero value otherwise.
func (o *Feedback) GetFeedbackForm() []FeedbackFormElement {
	if o == nil || isNil(o.FeedbackForm) {
		var ret []FeedbackFormElement
		return ret
	}
	return o.FeedbackForm
}

// GetFeedbackFormOk returns a tuple with the FeedbackForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Feedback) GetFeedbackFormOk() ([]FeedbackFormElement, bool) {
	if o == nil || isNil(o.FeedbackForm) {
    return nil, false
	}
	return o.FeedbackForm, true
}

// HasFeedbackForm returns a boolean if a field has been set.
func (o *Feedback) HasFeedbackForm() bool {
	if o != nil && !isNil(o.FeedbackForm) {
		return true
	}

	return false
}

// SetFeedbackForm gets a reference to the given []FeedbackFormElement and assigns it to the FeedbackForm field.
func (o *Feedback) SetFeedbackForm(v []FeedbackFormElement) {
	o.FeedbackForm = v
}

// GetFeedbackUrl returns the FeedbackUrl field value if set, zero value otherwise.
func (o *Feedback) GetFeedbackUrl() FeedbackUrl {
	if o == nil || isNil(o.FeedbackUrl) {
		var ret FeedbackUrl
		return ret
	}
	return *o.FeedbackUrl
}

// GetFeedbackUrlOk returns a tuple with the FeedbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Feedback) GetFeedbackUrlOk() (*FeedbackUrl, bool) {
	if o == nil || isNil(o.FeedbackUrl) {
    return nil, false
	}
	return o.FeedbackUrl, true
}

// HasFeedbackUrl returns a boolean if a field has been set.
func (o *Feedback) HasFeedbackUrl() bool {
	if o != nil && !isNil(o.FeedbackUrl) {
		return true
	}

	return false
}

// SetFeedbackUrl gets a reference to the given FeedbackUrl and assigns it to the FeedbackUrl field.
func (o *Feedback) SetFeedbackUrl(v FeedbackUrl) {
	o.FeedbackUrl = &v
}

func (o Feedback) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FeedbackForm) {
		toSerialize["feedback_form"] = o.FeedbackForm
	}
	if !isNil(o.FeedbackUrl) {
		toSerialize["feedback_url"] = o.FeedbackUrl
	}
	return json.Marshal(toSerialize)
}

type NullableFeedback struct {
	value *Feedback
	isSet bool
}

func (v NullableFeedback) Get() *Feedback {
	return v.value
}

func (v *NullableFeedback) Set(val *Feedback) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedback) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedback(val *Feedback) *NullableFeedback {
	return &NullableFeedback{value: val, isSet: true}
}

func (v NullableFeedback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


