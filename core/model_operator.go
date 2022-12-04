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

// Operator Describes the agent of a service
type Operator struct {
	// Describes the name of a person in format: ./{given_name}/{honorific_prefix}/{first_name}/{middle_name}/{last_name}/{honorific_suffix}
	Name *string `json:"name,omitempty"`
	// Image of an object. <br/><br/> A url based image will look like <br/><br/>```uri:http://path/to/image``` <br/><br/> An image can also be sent as a data string. For example : <br/><br/> ```data:js87y34ilhriuho84r3i4```
	Image *string `json:"image,omitempty"`
	Dob *string `json:"dob,omitempty"`
	// Gender of something, typically a Person, but possibly also fictional characters, animals, etc. While Male and Female may be used, text strings are also acceptable for people who do not identify as a binary gender
	Gender *string `json:"gender,omitempty"`
	Cred *string `json:"cred,omitempty"`
	// Describes a tag. This is a simple key-value store which is used to contain extended metadata
	Tags *map[string]string `json:"tags,omitempty"`
	Experience *OperatorAllOfExperience `json:"experience,omitempty"`
}

// NewOperator instantiates a new Operator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperator() *Operator {
	this := Operator{}
	return &this
}

// NewOperatorWithDefaults instantiates a new Operator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperatorWithDefaults() *Operator {
	this := Operator{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Operator) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Operator) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Operator) SetName(v string) {
	o.Name = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *Operator) GetImage() string {
	if o == nil || isNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetImageOk() (*string, bool) {
	if o == nil || isNil(o.Image) {
    return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *Operator) HasImage() bool {
	if o != nil && !isNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *Operator) SetImage(v string) {
	o.Image = &v
}

// GetDob returns the Dob field value if set, zero value otherwise.
func (o *Operator) GetDob() string {
	if o == nil || isNil(o.Dob) {
		var ret string
		return ret
	}
	return *o.Dob
}

// GetDobOk returns a tuple with the Dob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetDobOk() (*string, bool) {
	if o == nil || isNil(o.Dob) {
    return nil, false
	}
	return o.Dob, true
}

// HasDob returns a boolean if a field has been set.
func (o *Operator) HasDob() bool {
	if o != nil && !isNil(o.Dob) {
		return true
	}

	return false
}

// SetDob gets a reference to the given string and assigns it to the Dob field.
func (o *Operator) SetDob(v string) {
	o.Dob = &v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *Operator) GetGender() string {
	if o == nil || isNil(o.Gender) {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetGenderOk() (*string, bool) {
	if o == nil || isNil(o.Gender) {
    return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *Operator) HasGender() bool {
	if o != nil && !isNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *Operator) SetGender(v string) {
	o.Gender = &v
}

// GetCred returns the Cred field value if set, zero value otherwise.
func (o *Operator) GetCred() string {
	if o == nil || isNil(o.Cred) {
		var ret string
		return ret
	}
	return *o.Cred
}

// GetCredOk returns a tuple with the Cred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetCredOk() (*string, bool) {
	if o == nil || isNil(o.Cred) {
    return nil, false
	}
	return o.Cred, true
}

// HasCred returns a boolean if a field has been set.
func (o *Operator) HasCred() bool {
	if o != nil && !isNil(o.Cred) {
		return true
	}

	return false
}

// SetCred gets a reference to the given string and assigns it to the Cred field.
func (o *Operator) SetCred(v string) {
	o.Cred = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Operator) GetTags() map[string]string {
	if o == nil || isNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetTagsOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Operator) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *Operator) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetExperience returns the Experience field value if set, zero value otherwise.
func (o *Operator) GetExperience() OperatorAllOfExperience {
	if o == nil || isNil(o.Experience) {
		var ret OperatorAllOfExperience
		return ret
	}
	return *o.Experience
}

// GetExperienceOk returns a tuple with the Experience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Operator) GetExperienceOk() (*OperatorAllOfExperience, bool) {
	if o == nil || isNil(o.Experience) {
    return nil, false
	}
	return o.Experience, true
}

// HasExperience returns a boolean if a field has been set.
func (o *Operator) HasExperience() bool {
	if o != nil && !isNil(o.Experience) {
		return true
	}

	return false
}

// SetExperience gets a reference to the given OperatorAllOfExperience and assigns it to the Experience field.
func (o *Operator) SetExperience(v OperatorAllOfExperience) {
	o.Experience = &v
}

func (o Operator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !isNil(o.Dob) {
		toSerialize["dob"] = o.Dob
	}
	if !isNil(o.Gender) {
		toSerialize["gender"] = o.Gender
	}
	if !isNil(o.Cred) {
		toSerialize["cred"] = o.Cred
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.Experience) {
		toSerialize["experience"] = o.Experience
	}
	return json.Marshal(toSerialize)
}

type NullableOperator struct {
	value *Operator
	isSet bool
}

func (v NullableOperator) Get() *Operator {
	return v.value
}

func (v *NullableOperator) Set(val *Operator) {
	v.value = val
	v.isSet = true
}

func (v NullableOperator) IsSet() bool {
	return v.isSet
}

func (v *NullableOperator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperator(val *Operator) *NullableOperator {
	return &NullableOperator{value: val, isSet: true}
}

func (v NullableOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


