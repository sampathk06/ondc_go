# State

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Descriptor** | Pointer to [**Descriptor**](Descriptor.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedBy** | Pointer to **string** | ID of entity which changed the state | [optional] 

## Methods

### NewState

`func NewState() *State`

NewState instantiates a new State object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateWithDefaults

`func NewStateWithDefaults() *State`

NewStateWithDefaults instantiates a new State object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescriptor

`func (o *State) GetDescriptor() Descriptor`

GetDescriptor returns the Descriptor field if non-nil, zero value otherwise.

### GetDescriptorOk

`func (o *State) GetDescriptorOk() (*Descriptor, bool)`

GetDescriptorOk returns a tuple with the Descriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptor

`func (o *State) SetDescriptor(v Descriptor)`

SetDescriptor sets Descriptor field to given value.

### HasDescriptor

`func (o *State) HasDescriptor() bool`

HasDescriptor returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *State) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *State) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *State) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *State) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *State) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *State) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *State) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *State) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


