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

// FeedbackUrlParams struct for FeedbackUrlParams
type FeedbackUrlParams struct {
	// This value will be placed in the the $feedback_id url param in case of http/get and in the requestBody http/post requests
	FeedbackId string `json:"feedback_id"`
}

// NewFeedbackUrlParams instantiates a new FeedbackUrlParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedbackUrlParams(feedbackId string) *FeedbackUrlParams {
	this := FeedbackUrlParams{}
	this.FeedbackId = feedbackId
	return &this
}

// NewFeedbackUrlParamsWithDefaults instantiates a new FeedbackUrlParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedbackUrlParamsWithDefaults() *FeedbackUrlParams {
	this := FeedbackUrlParams{}
	return &this
}

// GetFeedbackId returns the FeedbackId field value
func (o *FeedbackUrlParams) GetFeedbackId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FeedbackId
}

// GetFeedbackIdOk returns a tuple with the FeedbackId field value
// and a boolean to check if the value has been set.
func (o *FeedbackUrlParams) GetFeedbackIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FeedbackId, true
}

// SetFeedbackId sets field value
func (o *FeedbackUrlParams) SetFeedbackId(v string) {
	o.FeedbackId = v
}

func (o FeedbackUrlParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["feedback_id"] = o.FeedbackId
	}
	return json.Marshal(toSerialize)
}

type NullableFeedbackUrlParams struct {
	value *FeedbackUrlParams
	isSet bool
}

func (v NullableFeedbackUrlParams) Get() *FeedbackUrlParams {
	return v.value
}

func (v *NullableFeedbackUrlParams) Set(val *FeedbackUrlParams) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedbackUrlParams) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedbackUrlParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedbackUrlParams(val *FeedbackUrlParams) *NullableFeedbackUrlParams {
	return &NullableFeedbackUrlParams{value: val, isSet: true}
}

func (v NullableFeedbackUrlParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedbackUrlParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

