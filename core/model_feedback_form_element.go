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

// FeedbackFormElement An element in the feedback form. It can be question or an answer to the question.
type FeedbackFormElement struct {
	Id *string `json:"id,omitempty"`
	ParentId *Id `json:"parent_id,omitempty"`
	// Specifies the question to which the answer options will be contained in the child FeedbackFormElements
	Question *string `json:"question,omitempty"`
	// Specifies an answer option to which the question will be in the FeedbackFormElement specified in parent_id
	Answer *string `json:"answer,omitempty"`
	// Specifies how the answer option should be rendered.
	AnswerType *string `json:"answer_type,omitempty"`
}

// NewFeedbackFormElement instantiates a new FeedbackFormElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedbackFormElement() *FeedbackFormElement {
	this := FeedbackFormElement{}
	return &this
}

// NewFeedbackFormElementWithDefaults instantiates a new FeedbackFormElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedbackFormElementWithDefaults() *FeedbackFormElement {
	this := FeedbackFormElement{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FeedbackFormElement) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackFormElement) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FeedbackFormElement) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FeedbackFormElement) SetId(v string) {
	o.Id = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *FeedbackFormElement) GetParentId() Id {
	if o == nil || isNil(o.ParentId) {
		var ret Id
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackFormElement) GetParentIdOk() (*Id, bool) {
	if o == nil || isNil(o.ParentId) {
    return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *FeedbackFormElement) HasParentId() bool {
	if o != nil && !isNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given Id and assigns it to the ParentId field.
func (o *FeedbackFormElement) SetParentId(v Id) {
	o.ParentId = &v
}

// GetQuestion returns the Question field value if set, zero value otherwise.
func (o *FeedbackFormElement) GetQuestion() string {
	if o == nil || isNil(o.Question) {
		var ret string
		return ret
	}
	return *o.Question
}

// GetQuestionOk returns a tuple with the Question field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackFormElement) GetQuestionOk() (*string, bool) {
	if o == nil || isNil(o.Question) {
    return nil, false
	}
	return o.Question, true
}

// HasQuestion returns a boolean if a field has been set.
func (o *FeedbackFormElement) HasQuestion() bool {
	if o != nil && !isNil(o.Question) {
		return true
	}

	return false
}

// SetQuestion gets a reference to the given string and assigns it to the Question field.
func (o *FeedbackFormElement) SetQuestion(v string) {
	o.Question = &v
}

// GetAnswer returns the Answer field value if set, zero value otherwise.
func (o *FeedbackFormElement) GetAnswer() string {
	if o == nil || isNil(o.Answer) {
		var ret string
		return ret
	}
	return *o.Answer
}

// GetAnswerOk returns a tuple with the Answer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackFormElement) GetAnswerOk() (*string, bool) {
	if o == nil || isNil(o.Answer) {
    return nil, false
	}
	return o.Answer, true
}

// HasAnswer returns a boolean if a field has been set.
func (o *FeedbackFormElement) HasAnswer() bool {
	if o != nil && !isNil(o.Answer) {
		return true
	}

	return false
}

// SetAnswer gets a reference to the given string and assigns it to the Answer field.
func (o *FeedbackFormElement) SetAnswer(v string) {
	o.Answer = &v
}

// GetAnswerType returns the AnswerType field value if set, zero value otherwise.
func (o *FeedbackFormElement) GetAnswerType() string {
	if o == nil || isNil(o.AnswerType) {
		var ret string
		return ret
	}
	return *o.AnswerType
}

// GetAnswerTypeOk returns a tuple with the AnswerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackFormElement) GetAnswerTypeOk() (*string, bool) {
	if o == nil || isNil(o.AnswerType) {
    return nil, false
	}
	return o.AnswerType, true
}

// HasAnswerType returns a boolean if a field has been set.
func (o *FeedbackFormElement) HasAnswerType() bool {
	if o != nil && !isNil(o.AnswerType) {
		return true
	}

	return false
}

// SetAnswerType gets a reference to the given string and assigns it to the AnswerType field.
func (o *FeedbackFormElement) SetAnswerType(v string) {
	o.AnswerType = &v
}

func (o FeedbackFormElement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.ParentId) {
		toSerialize["parent_id"] = o.ParentId
	}
	if !isNil(o.Question) {
		toSerialize["question"] = o.Question
	}
	if !isNil(o.Answer) {
		toSerialize["answer"] = o.Answer
	}
	if !isNil(o.AnswerType) {
		toSerialize["answer_type"] = o.AnswerType
	}
	return json.Marshal(toSerialize)
}

type NullableFeedbackFormElement struct {
	value *FeedbackFormElement
	isSet bool
}

func (v NullableFeedbackFormElement) Get() *FeedbackFormElement {
	return v.value
}

func (v *NullableFeedbackFormElement) Set(val *FeedbackFormElement) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedbackFormElement) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedbackFormElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedbackFormElement(val *FeedbackFormElement) *NullableFeedbackFormElement {
	return &NullableFeedbackFormElement{value: val, isSet: true}
}

func (v NullableFeedbackFormElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedbackFormElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


