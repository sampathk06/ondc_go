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

// Agent Describes an order executor
type Agent struct {
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
	Phone *string `json:"phone,omitempty"`
	Email *string `json:"email,omitempty"`
	// If the entity can be rated or not
	Rateable *bool `json:"rateable,omitempty"`
}

// NewAgent instantiates a new Agent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgent() *Agent {
	this := Agent{}
	return &this
}

// NewAgentWithDefaults instantiates a new Agent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentWithDefaults() *Agent {
	this := Agent{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Agent) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Agent) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Agent) SetName(v string) {
	o.Name = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *Agent) GetImage() string {
	if o == nil || isNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetImageOk() (*string, bool) {
	if o == nil || isNil(o.Image) {
    return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *Agent) HasImage() bool {
	if o != nil && !isNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *Agent) SetImage(v string) {
	o.Image = &v
}

// GetDob returns the Dob field value if set, zero value otherwise.
func (o *Agent) GetDob() string {
	if o == nil || isNil(o.Dob) {
		var ret string
		return ret
	}
	return *o.Dob
}

// GetDobOk returns a tuple with the Dob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetDobOk() (*string, bool) {
	if o == nil || isNil(o.Dob) {
    return nil, false
	}
	return o.Dob, true
}

// HasDob returns a boolean if a field has been set.
func (o *Agent) HasDob() bool {
	if o != nil && !isNil(o.Dob) {
		return true
	}

	return false
}

// SetDob gets a reference to the given string and assigns it to the Dob field.
func (o *Agent) SetDob(v string) {
	o.Dob = &v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *Agent) GetGender() string {
	if o == nil || isNil(o.Gender) {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetGenderOk() (*string, bool) {
	if o == nil || isNil(o.Gender) {
    return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *Agent) HasGender() bool {
	if o != nil && !isNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *Agent) SetGender(v string) {
	o.Gender = &v
}

// GetCred returns the Cred field value if set, zero value otherwise.
func (o *Agent) GetCred() string {
	if o == nil || isNil(o.Cred) {
		var ret string
		return ret
	}
	return *o.Cred
}

// GetCredOk returns a tuple with the Cred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetCredOk() (*string, bool) {
	if o == nil || isNil(o.Cred) {
    return nil, false
	}
	return o.Cred, true
}

// HasCred returns a boolean if a field has been set.
func (o *Agent) HasCred() bool {
	if o != nil && !isNil(o.Cred) {
		return true
	}

	return false
}

// SetCred gets a reference to the given string and assigns it to the Cred field.
func (o *Agent) SetCred(v string) {
	o.Cred = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Agent) GetTags() map[string]string {
	if o == nil || isNil(o.Tags) {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetTagsOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Agent) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *Agent) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *Agent) GetPhone() string {
	if o == nil || isNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetPhoneOk() (*string, bool) {
	if o == nil || isNil(o.Phone) {
    return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *Agent) HasPhone() bool {
	if o != nil && !isNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *Agent) SetPhone(v string) {
	o.Phone = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *Agent) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *Agent) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *Agent) SetEmail(v string) {
	o.Email = &v
}

// GetRateable returns the Rateable field value if set, zero value otherwise.
func (o *Agent) GetRateable() bool {
	if o == nil || isNil(o.Rateable) {
		var ret bool
		return ret
	}
	return *o.Rateable
}

// GetRateableOk returns a tuple with the Rateable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Agent) GetRateableOk() (*bool, bool) {
	if o == nil || isNil(o.Rateable) {
    return nil, false
	}
	return o.Rateable, true
}

// HasRateable returns a boolean if a field has been set.
func (o *Agent) HasRateable() bool {
	if o != nil && !isNil(o.Rateable) {
		return true
	}

	return false
}

// SetRateable gets a reference to the given bool and assigns it to the Rateable field.
func (o *Agent) SetRateable(v bool) {
	o.Rateable = &v
}

func (o Agent) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.Rateable) {
		toSerialize["rateable"] = o.Rateable
	}
	return json.Marshal(toSerialize)
}

type NullableAgent struct {
	value *Agent
	isSet bool
}

func (v NullableAgent) Get() *Agent {
	return v.value
}

func (v *NullableAgent) Set(val *Agent) {
	v.value = val
	v.isSet = true
}

func (v NullableAgent) IsSet() bool {
	return v.isSet
}

func (v *NullableAgent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgent(val *Agent) *NullableAgent {
	return &NullableAgent{value: val, isSet: true}
}

func (v NullableAgent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


