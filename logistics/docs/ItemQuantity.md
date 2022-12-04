# ItemQuantity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allocated** | Pointer to [**ItemQuantityAllocated**](ItemQuantityAllocated.md) |  | [optional] 
**Available** | Pointer to [**ItemQuantityAllocated**](ItemQuantityAllocated.md) |  | [optional] 
**Maximum** | Pointer to [**ItemQuantityMaximum**](ItemQuantityMaximum.md) |  | [optional] 
**Minimum** | Pointer to [**ItemQuantityAllocated**](ItemQuantityAllocated.md) |  | [optional] 
**Selected** | Pointer to [**ItemQuantityAllocated**](ItemQuantityAllocated.md) |  | [optional] 

## Methods

### NewItemQuantity

`func NewItemQuantity() *ItemQuantity`

NewItemQuantity instantiates a new ItemQuantity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemQuantityWithDefaults

`func NewItemQuantityWithDefaults() *ItemQuantity`

NewItemQuantityWithDefaults instantiates a new ItemQuantity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocated

`func (o *ItemQuantity) GetAllocated() ItemQuantityAllocated`

GetAllocated returns the Allocated field if non-nil, zero value otherwise.

### GetAllocatedOk

`func (o *ItemQuantity) GetAllocatedOk() (*ItemQuantityAllocated, bool)`

GetAllocatedOk returns a tuple with the Allocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocated

`func (o *ItemQuantity) SetAllocated(v ItemQuantityAllocated)`

SetAllocated sets Allocated field to given value.

### HasAllocated

`func (o *ItemQuantity) HasAllocated() bool`

HasAllocated returns a boolean if a field has been set.

### GetAvailable

`func (o *ItemQuantity) GetAvailable() ItemQuantityAllocated`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *ItemQuantity) GetAvailableOk() (*ItemQuantityAllocated, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *ItemQuantity) SetAvailable(v ItemQuantityAllocated)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *ItemQuantity) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetMaximum

`func (o *ItemQuantity) GetMaximum() ItemQuantityMaximum`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *ItemQuantity) GetMaximumOk() (*ItemQuantityMaximum, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *ItemQuantity) SetMaximum(v ItemQuantityMaximum)`

SetMaximum sets Maximum field to given value.

### HasMaximum

`func (o *ItemQuantity) HasMaximum() bool`

HasMaximum returns a boolean if a field has been set.

### GetMinimum

`func (o *ItemQuantity) GetMinimum() ItemQuantityAllocated`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *ItemQuantity) GetMinimumOk() (*ItemQuantityAllocated, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *ItemQuantity) SetMinimum(v ItemQuantityAllocated)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *ItemQuantity) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.

### GetSelected

`func (o *ItemQuantity) GetSelected() ItemQuantityAllocated`

GetSelected returns the Selected field if non-nil, zero value otherwise.

### GetSelectedOk

`func (o *ItemQuantity) GetSelectedOk() (*ItemQuantityAllocated, bool)`

GetSelectedOk returns a tuple with the Selected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelected

`func (o *ItemQuantity) SetSelected(v ItemQuantityAllocated)`

SetSelected sets Selected field to given value.

### HasSelected

`func (o *ItemQuantity) HasSelected() bool`

HasSelected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


