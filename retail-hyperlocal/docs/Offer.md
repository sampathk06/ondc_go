# Offer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Descriptor** | Pointer to [**Descriptor**](Descriptor.md) |  | [optional] 
**LocationIds** | Pointer to [**[]Id**](Id.md) |  | [optional] 
**CategoryIds** | Pointer to [**[]Id**](Id.md) |  | [optional] 
**ItemIds** | Pointer to [**[]Id**](Id.md) |  | [optional] 
**Time** | Pointer to [**Time**](Time.md) |  | [optional] 

## Methods

### NewOffer

`func NewOffer() *Offer`

NewOffer instantiates a new Offer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferWithDefaults

`func NewOfferWithDefaults() *Offer`

NewOfferWithDefaults instantiates a new Offer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Offer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Offer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Offer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Offer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescriptor

`func (o *Offer) GetDescriptor() Descriptor`

GetDescriptor returns the Descriptor field if non-nil, zero value otherwise.

### GetDescriptorOk

`func (o *Offer) GetDescriptorOk() (*Descriptor, bool)`

GetDescriptorOk returns a tuple with the Descriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptor

`func (o *Offer) SetDescriptor(v Descriptor)`

SetDescriptor sets Descriptor field to given value.

### HasDescriptor

`func (o *Offer) HasDescriptor() bool`

HasDescriptor returns a boolean if a field has been set.

### GetLocationIds

`func (o *Offer) GetLocationIds() []Id`

GetLocationIds returns the LocationIds field if non-nil, zero value otherwise.

### GetLocationIdsOk

`func (o *Offer) GetLocationIdsOk() (*[]Id, bool)`

GetLocationIdsOk returns a tuple with the LocationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationIds

`func (o *Offer) SetLocationIds(v []Id)`

SetLocationIds sets LocationIds field to given value.

### HasLocationIds

`func (o *Offer) HasLocationIds() bool`

HasLocationIds returns a boolean if a field has been set.

### GetCategoryIds

`func (o *Offer) GetCategoryIds() []Id`

GetCategoryIds returns the CategoryIds field if non-nil, zero value otherwise.

### GetCategoryIdsOk

`func (o *Offer) GetCategoryIdsOk() (*[]Id, bool)`

GetCategoryIdsOk returns a tuple with the CategoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryIds

`func (o *Offer) SetCategoryIds(v []Id)`

SetCategoryIds sets CategoryIds field to given value.

### HasCategoryIds

`func (o *Offer) HasCategoryIds() bool`

HasCategoryIds returns a boolean if a field has been set.

### GetItemIds

`func (o *Offer) GetItemIds() []Id`

GetItemIds returns the ItemIds field if non-nil, zero value otherwise.

### GetItemIdsOk

`func (o *Offer) GetItemIdsOk() (*[]Id, bool)`

GetItemIdsOk returns a tuple with the ItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemIds

`func (o *Offer) SetItemIds(v []Id)`

SetItemIds sets ItemIds field to given value.

### HasItemIds

`func (o *Offer) HasItemIds() bool`

HasItemIds returns a boolean if a field has been set.

### GetTime

`func (o *Offer) GetTime() Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Offer) GetTimeOk() (*Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Offer) SetTime(v Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *Offer) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


