# Item

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | This is the most unique identifier of a service item. An example of an Item ID could be the SKU of a product. | [optional] 
**ParentItemId** | Pointer to [**Id**](Id.md) |  | [optional] 
**Descriptor** | Pointer to [**Descriptor**](Descriptor.md) |  | [optional] 
**Price** | Pointer to [**Price**](Price.md) |  | [optional] 
**CategoryId** | Pointer to [**Id**](Id.md) |  | [optional] 
**FulfillmentId** | Pointer to [**Id**](Id.md) |  | [optional] 
**Rating** | Pointer to [**Value**](Value.md) |  | [optional] 
**LocationId** | Pointer to [**Id**](Id.md) |  | [optional] 
**Time** | Pointer to [**Time**](Time.md) |  | [optional] 
**Rateable** | Pointer to **bool** | If the entity can be rated or not | [optional] 
**Matched** | Pointer to **bool** |  | [optional] 
**Related** | Pointer to **bool** |  | [optional] 
**Recommended** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **map[string]string** | Describes a tag. This is a simple key-value store which is used to contain extended metadata | [optional] 

## Methods

### NewItem

`func NewItem() *Item`

NewItem instantiates a new Item object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemWithDefaults

`func NewItemWithDefaults() *Item`

NewItemWithDefaults instantiates a new Item object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Item) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Item) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Item) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Item) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParentItemId

`func (o *Item) GetParentItemId() Id`

GetParentItemId returns the ParentItemId field if non-nil, zero value otherwise.

### GetParentItemIdOk

`func (o *Item) GetParentItemIdOk() (*Id, bool)`

GetParentItemIdOk returns a tuple with the ParentItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentItemId

`func (o *Item) SetParentItemId(v Id)`

SetParentItemId sets ParentItemId field to given value.

### HasParentItemId

`func (o *Item) HasParentItemId() bool`

HasParentItemId returns a boolean if a field has been set.

### GetDescriptor

`func (o *Item) GetDescriptor() Descriptor`

GetDescriptor returns the Descriptor field if non-nil, zero value otherwise.

### GetDescriptorOk

`func (o *Item) GetDescriptorOk() (*Descriptor, bool)`

GetDescriptorOk returns a tuple with the Descriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptor

`func (o *Item) SetDescriptor(v Descriptor)`

SetDescriptor sets Descriptor field to given value.

### HasDescriptor

`func (o *Item) HasDescriptor() bool`

HasDescriptor returns a boolean if a field has been set.

### GetPrice

`func (o *Item) GetPrice() Price`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Item) GetPriceOk() (*Price, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Item) SetPrice(v Price)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Item) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCategoryId

`func (o *Item) GetCategoryId() Id`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *Item) GetCategoryIdOk() (*Id, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *Item) SetCategoryId(v Id)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *Item) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetFulfillmentId

`func (o *Item) GetFulfillmentId() Id`

GetFulfillmentId returns the FulfillmentId field if non-nil, zero value otherwise.

### GetFulfillmentIdOk

`func (o *Item) GetFulfillmentIdOk() (*Id, bool)`

GetFulfillmentIdOk returns a tuple with the FulfillmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentId

`func (o *Item) SetFulfillmentId(v Id)`

SetFulfillmentId sets FulfillmentId field to given value.

### HasFulfillmentId

`func (o *Item) HasFulfillmentId() bool`

HasFulfillmentId returns a boolean if a field has been set.

### GetRating

`func (o *Item) GetRating() Value`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *Item) GetRatingOk() (*Value, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *Item) SetRating(v Value)`

SetRating sets Rating field to given value.

### HasRating

`func (o *Item) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetLocationId

`func (o *Item) GetLocationId() Id`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *Item) GetLocationIdOk() (*Id, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *Item) SetLocationId(v Id)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *Item) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetTime

`func (o *Item) GetTime() Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Item) GetTimeOk() (*Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Item) SetTime(v Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *Item) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetRateable

`func (o *Item) GetRateable() bool`

GetRateable returns the Rateable field if non-nil, zero value otherwise.

### GetRateableOk

`func (o *Item) GetRateableOk() (*bool, bool)`

GetRateableOk returns a tuple with the Rateable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateable

`func (o *Item) SetRateable(v bool)`

SetRateable sets Rateable field to given value.

### HasRateable

`func (o *Item) HasRateable() bool`

HasRateable returns a boolean if a field has been set.

### GetMatched

`func (o *Item) GetMatched() bool`

GetMatched returns the Matched field if non-nil, zero value otherwise.

### GetMatchedOk

`func (o *Item) GetMatchedOk() (*bool, bool)`

GetMatchedOk returns a tuple with the Matched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatched

`func (o *Item) SetMatched(v bool)`

SetMatched sets Matched field to given value.

### HasMatched

`func (o *Item) HasMatched() bool`

HasMatched returns a boolean if a field has been set.

### GetRelated

`func (o *Item) GetRelated() bool`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *Item) GetRelatedOk() (*bool, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *Item) SetRelated(v bool)`

SetRelated sets Related field to given value.

### HasRelated

`func (o *Item) HasRelated() bool`

HasRelated returns a boolean if a field has been set.

### GetRecommended

`func (o *Item) GetRecommended() bool`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *Item) GetRecommendedOk() (*bool, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *Item) SetRecommended(v bool)`

SetRecommended sets Recommended field to given value.

### HasRecommended

`func (o *Item) HasRecommended() bool`

HasRecommended returns a boolean if a field has been set.

### GetTags

`func (o *Item) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Item) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Item) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Item) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


