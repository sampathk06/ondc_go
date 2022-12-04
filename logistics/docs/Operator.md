# Operator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Describes the name of a person in format: ./{given_name}/{honorific_prefix}/{first_name}/{middle_name}/{last_name}/{honorific_suffix} | [optional] 
**Image** | Pointer to **string** | Image of an object. &lt;br/&gt;&lt;br/&gt; A url based image will look like &lt;br/&gt;&lt;br/&gt;&#x60;&#x60;&#x60;uri:http://path/to/image&#x60;&#x60;&#x60; &lt;br/&gt;&lt;br/&gt; An image can also be sent as a data string. For example : &lt;br/&gt;&lt;br/&gt; &#x60;&#x60;&#x60;data:js87y34ilhriuho84r3i4&#x60;&#x60;&#x60; | [optional] 
**Dob** | Pointer to **string** |  | [optional] 
**Gender** | Pointer to **string** | Gender of something, typically a Person, but possibly also fictional characters, animals, etc. While Male and Female may be used, text strings are also acceptable for people who do not identify as a binary gender | [optional] 
**Cred** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]string** | Describes a tag. This is a simple key-value store which is used to contain extended metadata | [optional] 
**Experience** | Pointer to [**OperatorAllOfExperience**](OperatorAllOfExperience.md) |  | [optional] 

## Methods

### NewOperator

`func NewOperator() *Operator`

NewOperator instantiates a new Operator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatorWithDefaults

`func NewOperatorWithDefaults() *Operator`

NewOperatorWithDefaults instantiates a new Operator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Operator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Operator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Operator) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Operator) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImage

`func (o *Operator) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Operator) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Operator) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *Operator) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetDob

`func (o *Operator) GetDob() string`

GetDob returns the Dob field if non-nil, zero value otherwise.

### GetDobOk

`func (o *Operator) GetDobOk() (*string, bool)`

GetDobOk returns a tuple with the Dob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDob

`func (o *Operator) SetDob(v string)`

SetDob sets Dob field to given value.

### HasDob

`func (o *Operator) HasDob() bool`

HasDob returns a boolean if a field has been set.

### GetGender

`func (o *Operator) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *Operator) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *Operator) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *Operator) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetCred

`func (o *Operator) GetCred() string`

GetCred returns the Cred field if non-nil, zero value otherwise.

### GetCredOk

`func (o *Operator) GetCredOk() (*string, bool)`

GetCredOk returns a tuple with the Cred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCred

`func (o *Operator) SetCred(v string)`

SetCred sets Cred field to given value.

### HasCred

`func (o *Operator) HasCred() bool`

HasCred returns a boolean if a field has been set.

### GetTags

`func (o *Operator) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Operator) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Operator) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Operator) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetExperience

`func (o *Operator) GetExperience() OperatorAllOfExperience`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *Operator) GetExperienceOk() (*OperatorAllOfExperience, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *Operator) SetExperience(v OperatorAllOfExperience)`

SetExperience sets Experience field to given value.

### HasExperience

`func (o *Operator) HasExperience() bool`

HasExperience returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


