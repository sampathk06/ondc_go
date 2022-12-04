# Person

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Describes the name of a person in format: ./{given_name}/{honorific_prefix}/{first_name}/{middle_name}/{last_name}/{honorific_suffix} | [optional] 
**Image** | Pointer to **string** | Image of an object. &lt;br/&gt;&lt;br/&gt; A url based image will look like &lt;br/&gt;&lt;br/&gt;&#x60;&#x60;&#x60;uri:http://path/to/image&#x60;&#x60;&#x60; &lt;br/&gt;&lt;br/&gt; An image can also be sent as a data string. For example : &lt;br/&gt;&lt;br/&gt; &#x60;&#x60;&#x60;data:js87y34ilhriuho84r3i4&#x60;&#x60;&#x60; | [optional] 
**Dob** | Pointer to **string** |  | [optional] 
**Gender** | Pointer to **string** | Gender of something, typically a Person, but possibly also fictional characters, animals, etc. While Male and Female may be used, text strings are also acceptable for people who do not identify as a binary gender | [optional] 
**Cred** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]string** | Describes a tag. This is a simple key-value store which is used to contain extended metadata | [optional] 

## Methods

### NewPerson

`func NewPerson() *Person`

NewPerson instantiates a new Person object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonWithDefaults

`func NewPersonWithDefaults() *Person`

NewPersonWithDefaults instantiates a new Person object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Person) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Person) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Person) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Person) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImage

`func (o *Person) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Person) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Person) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *Person) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetDob

`func (o *Person) GetDob() string`

GetDob returns the Dob field if non-nil, zero value otherwise.

### GetDobOk

`func (o *Person) GetDobOk() (*string, bool)`

GetDobOk returns a tuple with the Dob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDob

`func (o *Person) SetDob(v string)`

SetDob sets Dob field to given value.

### HasDob

`func (o *Person) HasDob() bool`

HasDob returns a boolean if a field has been set.

### GetGender

`func (o *Person) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *Person) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *Person) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *Person) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetCred

`func (o *Person) GetCred() string`

GetCred returns the Cred field if non-nil, zero value otherwise.

### GetCredOk

`func (o *Person) GetCredOk() (*string, bool)`

GetCredOk returns a tuple with the Cred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCred

`func (o *Person) SetCred(v string)`

SetCred sets Cred field to given value.

### HasCred

`func (o *Person) HasCred() bool`

HasCred returns a boolean if a field has been set.

### GetTags

`func (o *Person) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Person) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Person) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Person) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


