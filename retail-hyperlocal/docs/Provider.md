# Provider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the provider | 
**Descriptor** | [**Descriptor**](Descriptor.md) |  | 
**CategoryId** | Pointer to **string** | Category Id of the provider | [optional] 
**OndcOrgFssaiLicenseNo** | **string** | FSSAI license no. Mandatory for category_id \&quot;F&amp;B\&quot; | 
**Rating** | Pointer to [**Value**](Value.md) |  | [optional] 
**Time** | Pointer to [**Time**](Time.md) |  | [optional] 
**Categories** | Pointer to [**[]Category**](Category.md) |  | [optional] 
**Fulfillments** | Pointer to [**[]Fulfillment**](Fulfillment.md) |  | [optional] 
**Payments** | Pointer to [**[]Payment**](Payment.md) |  | [optional] 
**Locations** | Pointer to [**[]ProviderLocationsInner**](ProviderLocationsInner.md) |  | [optional] 
**Offers** | Pointer to [**[]Offer**](Offer.md) |  | [optional] 
**Items** | Pointer to [**[]Item**](Item.md) |  | [optional] 
**Ttl** | **string** | Validity of catalog in ISO8601 durations format after which it has to be refreshed e.g. &#39;P7D&#39; indicates validity of 7 days; value of 0 indicates catalog is not cacheable | 
**Exp** | Pointer to **time.Time** | Time after which catalog has to be refreshed | [optional] 
**Rateable** | Pointer to **bool** | If the entity can be rated or not | [optional] 
**Tags** | Pointer to **map[string]string** | Describes a tag. This is a simple key-value store which is used to contain extended metadata | [optional] 

## Methods

### NewProvider

`func NewProvider(id string, descriptor Descriptor, ondcOrgFssaiLicenseNo string, ttl string, ) *Provider`

NewProvider instantiates a new Provider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderWithDefaults

`func NewProviderWithDefaults() *Provider`

NewProviderWithDefaults instantiates a new Provider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Provider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Provider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Provider) SetId(v string)`

SetId sets Id field to given value.


### GetDescriptor

`func (o *Provider) GetDescriptor() Descriptor`

GetDescriptor returns the Descriptor field if non-nil, zero value otherwise.

### GetDescriptorOk

`func (o *Provider) GetDescriptorOk() (*Descriptor, bool)`

GetDescriptorOk returns a tuple with the Descriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptor

`func (o *Provider) SetDescriptor(v Descriptor)`

SetDescriptor sets Descriptor field to given value.


### GetCategoryId

`func (o *Provider) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *Provider) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *Provider) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *Provider) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetOndcOrgFssaiLicenseNo

`func (o *Provider) GetOndcOrgFssaiLicenseNo() string`

GetOndcOrgFssaiLicenseNo returns the OndcOrgFssaiLicenseNo field if non-nil, zero value otherwise.

### GetOndcOrgFssaiLicenseNoOk

`func (o *Provider) GetOndcOrgFssaiLicenseNoOk() (*string, bool)`

GetOndcOrgFssaiLicenseNoOk returns a tuple with the OndcOrgFssaiLicenseNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgFssaiLicenseNo

`func (o *Provider) SetOndcOrgFssaiLicenseNo(v string)`

SetOndcOrgFssaiLicenseNo sets OndcOrgFssaiLicenseNo field to given value.


### GetRating

`func (o *Provider) GetRating() Value`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *Provider) GetRatingOk() (*Value, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *Provider) SetRating(v Value)`

SetRating sets Rating field to given value.

### HasRating

`func (o *Provider) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetTime

`func (o *Provider) GetTime() Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Provider) GetTimeOk() (*Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Provider) SetTime(v Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *Provider) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetCategories

`func (o *Provider) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Provider) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Provider) SetCategories(v []Category)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *Provider) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetFulfillments

`func (o *Provider) GetFulfillments() []Fulfillment`

GetFulfillments returns the Fulfillments field if non-nil, zero value otherwise.

### GetFulfillmentsOk

`func (o *Provider) GetFulfillmentsOk() (*[]Fulfillment, bool)`

GetFulfillmentsOk returns a tuple with the Fulfillments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillments

`func (o *Provider) SetFulfillments(v []Fulfillment)`

SetFulfillments sets Fulfillments field to given value.

### HasFulfillments

`func (o *Provider) HasFulfillments() bool`

HasFulfillments returns a boolean if a field has been set.

### GetPayments

`func (o *Provider) GetPayments() []Payment`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *Provider) GetPaymentsOk() (*[]Payment, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *Provider) SetPayments(v []Payment)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *Provider) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### GetLocations

`func (o *Provider) GetLocations() []ProviderLocationsInner`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *Provider) GetLocationsOk() (*[]ProviderLocationsInner, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *Provider) SetLocations(v []ProviderLocationsInner)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *Provider) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetOffers

`func (o *Provider) GetOffers() []Offer`

GetOffers returns the Offers field if non-nil, zero value otherwise.

### GetOffersOk

`func (o *Provider) GetOffersOk() (*[]Offer, bool)`

GetOffersOk returns a tuple with the Offers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffers

`func (o *Provider) SetOffers(v []Offer)`

SetOffers sets Offers field to given value.

### HasOffers

`func (o *Provider) HasOffers() bool`

HasOffers returns a boolean if a field has been set.

### GetItems

`func (o *Provider) GetItems() []Item`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Provider) GetItemsOk() (*[]Item, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Provider) SetItems(v []Item)`

SetItems sets Items field to given value.

### HasItems

`func (o *Provider) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTtl

`func (o *Provider) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *Provider) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *Provider) SetTtl(v string)`

SetTtl sets Ttl field to given value.


### GetExp

`func (o *Provider) GetExp() time.Time`

GetExp returns the Exp field if non-nil, zero value otherwise.

### GetExpOk

`func (o *Provider) GetExpOk() (*time.Time, bool)`

GetExpOk returns a tuple with the Exp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExp

`func (o *Provider) SetExp(v time.Time)`

SetExp sets Exp field to given value.

### HasExp

`func (o *Provider) HasExp() bool`

HasExp returns a boolean if a field has been set.

### GetRateable

`func (o *Provider) GetRateable() bool`

GetRateable returns the Rateable field if non-nil, zero value otherwise.

### GetRateableOk

`func (o *Provider) GetRateableOk() (*bool, bool)`

GetRateableOk returns a tuple with the Rateable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateable

`func (o *Provider) SetRateable(v bool)`

SetRateable sets Rateable field to given value.

### HasRateable

`func (o *Provider) HasRateable() bool`

HasRateable returns a boolean if a field has been set.

### GetTags

`func (o *Provider) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Provider) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Provider) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Provider) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


