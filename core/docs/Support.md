# Support

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **string** |  | [optional] 
**Channels** | Pointer to **map[string]string** | Describes a tag. This is a simple key-value store which is used to contain extended metadata | [optional] 

## Methods

### NewSupport

`func NewSupport() *Support`

NewSupport instantiates a new Support object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportWithDefaults

`func NewSupportWithDefaults() *Support`

NewSupportWithDefaults instantiates a new Support object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Support) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Support) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Support) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Support) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRefId

`func (o *Support) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *Support) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *Support) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *Support) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetChannels

`func (o *Support) GetChannels() map[string]string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *Support) GetChannelsOk() (*map[string]string, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *Support) SetChannels(v map[string]string)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *Support) HasChannels() bool`

HasChannels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


