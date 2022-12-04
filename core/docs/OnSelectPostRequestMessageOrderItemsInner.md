# OnSelectPostRequestMessageOrderItemsInner

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
**Quantity** | Pointer to [**ItemQuantity**](ItemQuantity.md) |  | [optional] 

## Methods

### NewOnSelectPostRequestMessageOrderItemsInner

`func NewOnSelectPostRequestMessageOrderItemsInner() *OnSelectPostRequestMessageOrderItemsInner`

NewOnSelectPostRequestMessageOrderItemsInner instantiates a new OnSelectPostRequestMessageOrderItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnSelectPostRequestMessageOrderItemsInnerWithDefaults

`func NewOnSelectPostRequestMessageOrderItemsInnerWithDefaults() *OnSelectPostRequestMessageOrderItemsInner`

NewOnSelectPostRequestMessageOrderItemsInnerWithDefaults instantiates a new OnSelectPostRequestMessageOrderItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OnSelectPostRequestMessageOrderItemsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OnSelectPostRequestMessageOrderItemsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParentItemId

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetParentItemId() Id`

GetParentItemId returns the ParentItemId field if non-nil, zero value otherwise.

### GetParentItemIdOk

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetParentItemIdOk() (*Id, bool)`

GetParentItemIdOk returns a tuple with the ParentItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentItemId

`func (o *OnSelectPostRequestMessageOrderItemsInner) SetParentItemId(v Id)`

SetParentItemId sets ParentItemId field to given value.

### HasParentItemId

`func (o *OnSelectPostRequestMessageOrderItemsInner) HasParentItemId() bool`

HasParentItemId returns a boolean if a field has been set.

### GetDescriptor

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetDescriptor() Descriptor`

GetDescriptor returns the Descriptor field if non-nil, zero value otherwise.

### GetDescriptorOk

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetDescriptorOk() (*Descriptor, bool)`

GetDescriptorOk returns a tuple with the Descriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptor

`func (o *OnSelectPostRequestMessageOrderItemsInner) SetDescriptor(v Descriptor)`

SetDescriptor sets Descriptor field to given value.

### HasDescriptor

`func (o *OnSelectPostRequestMessageOrderItemsInner) HasDescriptor() bool`

HasDescriptor returns a boolean if a field has been set.

### GetPrice

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetPrice() Price`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetPriceOk() (*Price, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OnSelectPostRequestMessageOrderItemsInner) SetPrice(v Price)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OnSelectPostRequestMessageOrderItemsInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCategoryId

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetCategoryId() Id`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetCategoryIdOk() (*Id, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *OnSelectPostRequestMessageOrderItemsInner) SetCategoryId(v Id)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *OnSelectPostRequestMessageOrderItemsInner) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetFulfillmentId

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetFulfillmentId() Id`

GetFulfillmentId returns the FulfillmentId field if non-nil, zero value otherwise.

### GetFulfillmentIdOk

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetFulfillmentIdOk() (*Id, bool)`

GetFulfillmentIdOk returns a tuple with the FulfillmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentId

`func (o *OnSelectPostRequestMessageOrderItemsInner) SetFulfillmentId(v Id)`

SetFulfillmentId sets FulfillmentId field to given value.

### HasFulfillmentId

`func (o *OnSelectPostRequestMessageOrderItemsInner) HasFulfillmentId() bool`

HasFulfillmentId returns a boolean if a field has been set.

### GetRating

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetRating() Value`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetRatingOk() (*Value, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *OnSelectPostRequestMessageOrderItemsInner) SetRating(v Value)`

SetRating sets Rating field to given value.

### HasRating

`func (o *OnSelectPostRequestMessageOrderItemsInner) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetLocationId

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetLocationId() Id`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetLocationIdOk() (*Id, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *OnSelectPostRequestMessageOrderItemsInner) SetLocationId(v Id)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *OnSelectPostRequestMessageOrderItemsInner) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetTime

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetTime() Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetTimeOk() (*Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *OnSelectPostRequestMessageOrderItemsInner) SetTime(v Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *OnSelectPostRequestMessageOrderItemsInner) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetRateable

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetRateable() bool`

GetRateable returns the Rateable field if non-nil, zero value otherwise.

### GetRateableOk

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetRateableOk() (*bool, bool)`

GetRateableOk returns a tuple with the Rateable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateable

`func (o *OnSelectPostRequestMessageOrderItemsInner) SetRateable(v bool)`

SetRateable sets Rateable field to given value.

### HasRateable

`func (o *OnSelectPostRequestMessageOrderItemsInner) HasRateable() bool`

HasRateable returns a boolean if a field has been set.

### GetMatched

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetMatched() bool`

GetMatched returns the Matched field if non-nil, zero value otherwise.

### GetMatchedOk

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetMatchedOk() (*bool, bool)`

GetMatchedOk returns a tuple with the Matched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatched

`func (o *OnSelectPostRequestMessageOrderItemsInner) SetMatched(v bool)`

SetMatched sets Matched field to given value.

### HasMatched

`func (o *OnSelectPostRequestMessageOrderItemsInner) HasMatched() bool`

HasMatched returns a boolean if a field has been set.

### GetRelated

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetRelated() bool`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetRelatedOk() (*bool, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *OnSelectPostRequestMessageOrderItemsInner) SetRelated(v bool)`

SetRelated sets Related field to given value.

### HasRelated

`func (o *OnSelectPostRequestMessageOrderItemsInner) HasRelated() bool`

HasRelated returns a boolean if a field has been set.

### GetRecommended

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetRecommended() bool`

GetRecommended returns the Recommended field if non-nil, zero value otherwise.

### GetRecommendedOk

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetRecommendedOk() (*bool, bool)`

GetRecommendedOk returns a tuple with the Recommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommended

`func (o *OnSelectPostRequestMessageOrderItemsInner) SetRecommended(v bool)`

SetRecommended sets Recommended field to given value.

### HasRecommended

`func (o *OnSelectPostRequestMessageOrderItemsInner) HasRecommended() bool`

HasRecommended returns a boolean if a field has been set.

### GetTags

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OnSelectPostRequestMessageOrderItemsInner) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OnSelectPostRequestMessageOrderItemsInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetQuantity

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetQuantity() ItemQuantity`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OnSelectPostRequestMessageOrderItemsInner) GetQuantityOk() (*ItemQuantity, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OnSelectPostRequestMessageOrderItemsInner) SetQuantity(v ItemQuantity)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OnSelectPostRequestMessageOrderItemsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


