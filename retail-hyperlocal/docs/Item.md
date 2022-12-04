# Item

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | This is the most unique identifier of a service item. An example of an Item ID could be the SKU of a product. | 
**ParentItemId** | Pointer to [**Id**](Id.md) |  | [optional] 
**Descriptor** | [**Descriptor**](Descriptor.md) |  | 
**Price** | [**Price**](Price.md) |  | 
**Quantity** | Pointer to [**ItemQuantity**](ItemQuantity.md) |  | [optional] 
**CategoryId** | [**Id**](Id.md) |  | 
**FulfillmentId** | [**Id**](Id.md) |  | 
**Rating** | Pointer to [**Value**](Value.md) |  | [optional] 
**LocationId** | [**Id**](Id.md) |  | 
**Time** | Pointer to [**Time**](Time.md) |  | [optional] 
**Rateable** | Pointer to **bool** | If the entity can be rated or not | [optional] 
**Matched** | **bool** |  | 
**Related** | Pointer to **bool** |  | [optional] 
**Recommended** | Pointer to **bool** |  | [optional] 
**OndcOrgReturnable** | **bool** | whether the item is returnable | 
**OndcOrgSellerPickupReturn** | **bool** | in case of return, whether the item should be picked up by seller | 
**OndcOrgReturnWindow** | **string** | return window for the item in ISO8601 durations format e.g. &#39;PT24H&#39; indicates 24 hour return window. Mandatory if \&quot;@ondc/org/returnable\&quot; is \&quot;true\&quot; | 
**OndcOrgCancellable** | **bool** | whether the item is cancellable | 
**OndcOrgTimeToShip** | **string** | time from order confirmation by which item ready to ship in ISO8601 durations format (e.g. &#39;PT30M&#39; indicates item ready to ship in 30 mins). Mandatory for category_id \&quot;F&amp;B\&quot; | 
**OndcOrgAvailableOnCod** | **bool** | whether the catalog item is available on COD | 
**OndcOrgContactDetailsConsumerCare** | **string** | contact details for consumer care | 
**OndcOrgStatutoryReqsPackagedCommodities** | [**ItemOndcOrgStatutoryReqsPackagedCommodities**](ItemOndcOrgStatutoryReqsPackagedCommodities.md) |  | 
**OndcOrgStatutoryReqsPrepackagedFood** | [**ItemOndcOrgStatutoryReqsPrepackagedFood**](ItemOndcOrgStatutoryReqsPrepackagedFood.md) |  | 
**OndcOrgMandatoryReqsVeggiesFruits** | **map[string]interface{}** | mandatory for category_id \&quot;Fruits and Vegetables\&quot; required attributes include the following - net_quantity | 
**Tags** | Pointer to **map[string]string** | Describes a tag. This is a simple key-value store which is used to contain extended metadata | [optional] 

## Methods

### NewItem

`func NewItem(id string, descriptor Descriptor, price Price, categoryId Id, fulfillmentId Id, locationId Id, matched bool, ondcOrgReturnable bool, ondcOrgSellerPickupReturn bool, ondcOrgReturnWindow string, ondcOrgCancellable bool, ondcOrgTimeToShip string, ondcOrgAvailableOnCod bool, ondcOrgContactDetailsConsumerCare string, ondcOrgStatutoryReqsPackagedCommodities ItemOndcOrgStatutoryReqsPackagedCommodities, ondcOrgStatutoryReqsPrepackagedFood ItemOndcOrgStatutoryReqsPrepackagedFood, ondcOrgMandatoryReqsVeggiesFruits map[string]interface{}, ) *Item`

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


### GetQuantity

`func (o *Item) GetQuantity() ItemQuantity`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Item) GetQuantityOk() (*ItemQuantity, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Item) SetQuantity(v ItemQuantity)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *Item) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

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

### GetOndcOrgReturnable

`func (o *Item) GetOndcOrgReturnable() bool`

GetOndcOrgReturnable returns the OndcOrgReturnable field if non-nil, zero value otherwise.

### GetOndcOrgReturnableOk

`func (o *Item) GetOndcOrgReturnableOk() (*bool, bool)`

GetOndcOrgReturnableOk returns a tuple with the OndcOrgReturnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgReturnable

`func (o *Item) SetOndcOrgReturnable(v bool)`

SetOndcOrgReturnable sets OndcOrgReturnable field to given value.


### GetOndcOrgSellerPickupReturn

`func (o *Item) GetOndcOrgSellerPickupReturn() bool`

GetOndcOrgSellerPickupReturn returns the OndcOrgSellerPickupReturn field if non-nil, zero value otherwise.

### GetOndcOrgSellerPickupReturnOk

`func (o *Item) GetOndcOrgSellerPickupReturnOk() (*bool, bool)`

GetOndcOrgSellerPickupReturnOk returns a tuple with the OndcOrgSellerPickupReturn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgSellerPickupReturn

`func (o *Item) SetOndcOrgSellerPickupReturn(v bool)`

SetOndcOrgSellerPickupReturn sets OndcOrgSellerPickupReturn field to given value.


### GetOndcOrgReturnWindow

`func (o *Item) GetOndcOrgReturnWindow() string`

GetOndcOrgReturnWindow returns the OndcOrgReturnWindow field if non-nil, zero value otherwise.

### GetOndcOrgReturnWindowOk

`func (o *Item) GetOndcOrgReturnWindowOk() (*string, bool)`

GetOndcOrgReturnWindowOk returns a tuple with the OndcOrgReturnWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgReturnWindow

`func (o *Item) SetOndcOrgReturnWindow(v string)`

SetOndcOrgReturnWindow sets OndcOrgReturnWindow field to given value.


### GetOndcOrgCancellable

`func (o *Item) GetOndcOrgCancellable() bool`

GetOndcOrgCancellable returns the OndcOrgCancellable field if non-nil, zero value otherwise.

### GetOndcOrgCancellableOk

`func (o *Item) GetOndcOrgCancellableOk() (*bool, bool)`

GetOndcOrgCancellableOk returns a tuple with the OndcOrgCancellable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgCancellable

`func (o *Item) SetOndcOrgCancellable(v bool)`

SetOndcOrgCancellable sets OndcOrgCancellable field to given value.


### GetOndcOrgTimeToShip

`func (o *Item) GetOndcOrgTimeToShip() string`

GetOndcOrgTimeToShip returns the OndcOrgTimeToShip field if non-nil, zero value otherwise.

### GetOndcOrgTimeToShipOk

`func (o *Item) GetOndcOrgTimeToShipOk() (*string, bool)`

GetOndcOrgTimeToShipOk returns a tuple with the OndcOrgTimeToShip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgTimeToShip

`func (o *Item) SetOndcOrgTimeToShip(v string)`

SetOndcOrgTimeToShip sets OndcOrgTimeToShip field to given value.


### GetOndcOrgAvailableOnCod

`func (o *Item) GetOndcOrgAvailableOnCod() bool`

GetOndcOrgAvailableOnCod returns the OndcOrgAvailableOnCod field if non-nil, zero value otherwise.

### GetOndcOrgAvailableOnCodOk

`func (o *Item) GetOndcOrgAvailableOnCodOk() (*bool, bool)`

GetOndcOrgAvailableOnCodOk returns a tuple with the OndcOrgAvailableOnCod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgAvailableOnCod

`func (o *Item) SetOndcOrgAvailableOnCod(v bool)`

SetOndcOrgAvailableOnCod sets OndcOrgAvailableOnCod field to given value.


### GetOndcOrgContactDetailsConsumerCare

`func (o *Item) GetOndcOrgContactDetailsConsumerCare() string`

GetOndcOrgContactDetailsConsumerCare returns the OndcOrgContactDetailsConsumerCare field if non-nil, zero value otherwise.

### GetOndcOrgContactDetailsConsumerCareOk

`func (o *Item) GetOndcOrgContactDetailsConsumerCareOk() (*string, bool)`

GetOndcOrgContactDetailsConsumerCareOk returns a tuple with the OndcOrgContactDetailsConsumerCare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgContactDetailsConsumerCare

`func (o *Item) SetOndcOrgContactDetailsConsumerCare(v string)`

SetOndcOrgContactDetailsConsumerCare sets OndcOrgContactDetailsConsumerCare field to given value.


### GetOndcOrgStatutoryReqsPackagedCommodities

`func (o *Item) GetOndcOrgStatutoryReqsPackagedCommodities() ItemOndcOrgStatutoryReqsPackagedCommodities`

GetOndcOrgStatutoryReqsPackagedCommodities returns the OndcOrgStatutoryReqsPackagedCommodities field if non-nil, zero value otherwise.

### GetOndcOrgStatutoryReqsPackagedCommoditiesOk

`func (o *Item) GetOndcOrgStatutoryReqsPackagedCommoditiesOk() (*ItemOndcOrgStatutoryReqsPackagedCommodities, bool)`

GetOndcOrgStatutoryReqsPackagedCommoditiesOk returns a tuple with the OndcOrgStatutoryReqsPackagedCommodities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgStatutoryReqsPackagedCommodities

`func (o *Item) SetOndcOrgStatutoryReqsPackagedCommodities(v ItemOndcOrgStatutoryReqsPackagedCommodities)`

SetOndcOrgStatutoryReqsPackagedCommodities sets OndcOrgStatutoryReqsPackagedCommodities field to given value.


### GetOndcOrgStatutoryReqsPrepackagedFood

`func (o *Item) GetOndcOrgStatutoryReqsPrepackagedFood() ItemOndcOrgStatutoryReqsPrepackagedFood`

GetOndcOrgStatutoryReqsPrepackagedFood returns the OndcOrgStatutoryReqsPrepackagedFood field if non-nil, zero value otherwise.

### GetOndcOrgStatutoryReqsPrepackagedFoodOk

`func (o *Item) GetOndcOrgStatutoryReqsPrepackagedFoodOk() (*ItemOndcOrgStatutoryReqsPrepackagedFood, bool)`

GetOndcOrgStatutoryReqsPrepackagedFoodOk returns a tuple with the OndcOrgStatutoryReqsPrepackagedFood field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgStatutoryReqsPrepackagedFood

`func (o *Item) SetOndcOrgStatutoryReqsPrepackagedFood(v ItemOndcOrgStatutoryReqsPrepackagedFood)`

SetOndcOrgStatutoryReqsPrepackagedFood sets OndcOrgStatutoryReqsPrepackagedFood field to given value.


### GetOndcOrgMandatoryReqsVeggiesFruits

`func (o *Item) GetOndcOrgMandatoryReqsVeggiesFruits() map[string]interface{}`

GetOndcOrgMandatoryReqsVeggiesFruits returns the OndcOrgMandatoryReqsVeggiesFruits field if non-nil, zero value otherwise.

### GetOndcOrgMandatoryReqsVeggiesFruitsOk

`func (o *Item) GetOndcOrgMandatoryReqsVeggiesFruitsOk() (*map[string]interface{}, bool)`

GetOndcOrgMandatoryReqsVeggiesFruitsOk returns a tuple with the OndcOrgMandatoryReqsVeggiesFruits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgMandatoryReqsVeggiesFruits

`func (o *Item) SetOndcOrgMandatoryReqsVeggiesFruits(v map[string]interface{})`

SetOndcOrgMandatoryReqsVeggiesFruits sets OndcOrgMandatoryReqsVeggiesFruits field to given value.


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


