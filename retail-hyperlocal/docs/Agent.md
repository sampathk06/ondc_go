# Agent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Describes the name of a person in format: ./{given_name}/{honorific_prefix}/{first_name}/{middle_name}/{last_name}/{honorific_suffix} | [optional] 
**Image** | Pointer to **string** | Image of an object. &lt;br/&gt;&lt;br/&gt; A url based image will look like &lt;br/&gt;&lt;br/&gt;&#x60;&#x60;&#x60;uri:http://path/to/image&#x60;&#x60;&#x60; &lt;br/&gt;&lt;br/&gt; An image can also be sent as a data string. For example : &lt;br/&gt;&lt;br/&gt; &#x60;&#x60;&#x60;data:js87y34ilhriuho84r3i4&#x60;&#x60;&#x60; | [optional] 
**Dob** | Pointer to **string** |  | [optional] 
**Gender** | Pointer to **string** | Gender of something, typically a Person, but possibly also fictional characters, animals, etc. While Male and Female may be used, text strings are also acceptable for people who do not identify as a binary gender | [optional] 
**Cred** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]string** | Describes a tag. This is a simple key-value store which is used to contain extended metadata | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Rateable** | Pointer to **bool** | If the entity can be rated or not | [optional] 

## Methods

### NewAgent

`func NewAgent() *Agent`

NewAgent instantiates a new Agent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentWithDefaults

`func NewAgentWithDefaults() *Agent`

NewAgentWithDefaults instantiates a new Agent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Agent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Agent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Agent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Agent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImage

`func (o *Agent) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Agent) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Agent) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *Agent) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetDob

`func (o *Agent) GetDob() string`

GetDob returns the Dob field if non-nil, zero value otherwise.

### GetDobOk

`func (o *Agent) GetDobOk() (*string, bool)`

GetDobOk returns a tuple with the Dob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDob

`func (o *Agent) SetDob(v string)`

SetDob sets Dob field to given value.

### HasDob

`func (o *Agent) HasDob() bool`

HasDob returns a boolean if a field has been set.

### GetGender

`func (o *Agent) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *Agent) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *Agent) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *Agent) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetCred

`func (o *Agent) GetCred() string`

GetCred returns the Cred field if non-nil, zero value otherwise.

### GetCredOk

`func (o *Agent) GetCredOk() (*string, bool)`

GetCredOk returns a tuple with the Cred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCred

`func (o *Agent) SetCred(v string)`

SetCred sets Cred field to given value.

### HasCred

`func (o *Agent) HasCred() bool`

HasCred returns a boolean if a field has been set.

### GetTags

`func (o *Agent) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Agent) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Agent) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Agent) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPhone

`func (o *Agent) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Agent) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Agent) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Agent) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *Agent) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Agent) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Agent) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Agent) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetRateable

`func (o *Agent) GetRateable() bool`

GetRateable returns the Rateable field if non-nil, zero value otherwise.

### GetRateableOk

`func (o *Agent) GetRateableOk() (*bool, bool)`

GetRateableOk returns a tuple with the Rateable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateable

`func (o *Agent) SetRateable(v bool)`

SetRateable sets Rateable field to given value.

### HasRateable

`func (o *Agent) HasRateable() bool`

HasRateable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


