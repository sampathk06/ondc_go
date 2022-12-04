# Policy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Descriptor** | Pointer to [**Descriptor**](Descriptor.md) |  | [optional] 
**ParentPolicyId** | Pointer to [**Id**](Id.md) |  | [optional] 
**Time** | Pointer to [**Time**](Time.md) |  | [optional] 

## Methods

### NewPolicy

`func NewPolicy() *Policy`

NewPolicy instantiates a new Policy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyWithDefaults

`func NewPolicyWithDefaults() *Policy`

NewPolicyWithDefaults instantiates a new Policy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Policy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Policy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Policy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Policy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescriptor

`func (o *Policy) GetDescriptor() Descriptor`

GetDescriptor returns the Descriptor field if non-nil, zero value otherwise.

### GetDescriptorOk

`func (o *Policy) GetDescriptorOk() (*Descriptor, bool)`

GetDescriptorOk returns a tuple with the Descriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptor

`func (o *Policy) SetDescriptor(v Descriptor)`

SetDescriptor sets Descriptor field to given value.

### HasDescriptor

`func (o *Policy) HasDescriptor() bool`

HasDescriptor returns a boolean if a field has been set.

### GetParentPolicyId

`func (o *Policy) GetParentPolicyId() Id`

GetParentPolicyId returns the ParentPolicyId field if non-nil, zero value otherwise.

### GetParentPolicyIdOk

`func (o *Policy) GetParentPolicyIdOk() (*Id, bool)`

GetParentPolicyIdOk returns a tuple with the ParentPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPolicyId

`func (o *Policy) SetParentPolicyId(v Id)`

SetParentPolicyId sets ParentPolicyId field to given value.

### HasParentPolicyId

`func (o *Policy) HasParentPolicyId() bool`

HasParentPolicyId returns a boolean if a field has been set.

### GetTime

`func (o *Policy) GetTime() Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Policy) GetTimeOk() (*Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Policy) SetTime(v Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *Policy) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


