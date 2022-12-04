# Intent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Descriptor** | Pointer to [**Descriptor**](Descriptor.md) |  | [optional] 
**Provider** | Pointer to [**Provider**](Provider.md) |  | [optional] 
**Fulfillment** | Pointer to [**Fulfillment**](Fulfillment.md) |  | [optional] 
**Payment** | Pointer to [**Payment**](Payment.md) |  | [optional] 
**Category** | Pointer to [**Category**](Category.md) |  | [optional] 
**Offer** | Pointer to [**Offer**](Offer.md) |  | [optional] 
**Item** | Pointer to [**Item**](Item.md) |  | [optional] 
**OndcOrgPayloadDetails** | Pointer to [**IntentOndcOrgPayloadDetails**](IntentOndcOrgPayloadDetails.md) |  | [optional] 
**Tags** | Pointer to **map[string]string** | Describes a tag. This is a simple key-value store which is used to contain extended metadata | [optional] 

## Methods

### NewIntent

`func NewIntent() *Intent`

NewIntent instantiates a new Intent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntentWithDefaults

`func NewIntentWithDefaults() *Intent`

NewIntentWithDefaults instantiates a new Intent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescriptor

`func (o *Intent) GetDescriptor() Descriptor`

GetDescriptor returns the Descriptor field if non-nil, zero value otherwise.

### GetDescriptorOk

`func (o *Intent) GetDescriptorOk() (*Descriptor, bool)`

GetDescriptorOk returns a tuple with the Descriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptor

`func (o *Intent) SetDescriptor(v Descriptor)`

SetDescriptor sets Descriptor field to given value.

### HasDescriptor

`func (o *Intent) HasDescriptor() bool`

HasDescriptor returns a boolean if a field has been set.

### GetProvider

`func (o *Intent) GetProvider() Provider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Intent) GetProviderOk() (*Provider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Intent) SetProvider(v Provider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *Intent) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetFulfillment

`func (o *Intent) GetFulfillment() Fulfillment`

GetFulfillment returns the Fulfillment field if non-nil, zero value otherwise.

### GetFulfillmentOk

`func (o *Intent) GetFulfillmentOk() (*Fulfillment, bool)`

GetFulfillmentOk returns a tuple with the Fulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillment

`func (o *Intent) SetFulfillment(v Fulfillment)`

SetFulfillment sets Fulfillment field to given value.

### HasFulfillment

`func (o *Intent) HasFulfillment() bool`

HasFulfillment returns a boolean if a field has been set.

### GetPayment

`func (o *Intent) GetPayment() Payment`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *Intent) GetPaymentOk() (*Payment, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *Intent) SetPayment(v Payment)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *Intent) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetCategory

`func (o *Intent) GetCategory() Category`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Intent) GetCategoryOk() (*Category, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Intent) SetCategory(v Category)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Intent) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetOffer

`func (o *Intent) GetOffer() Offer`

GetOffer returns the Offer field if non-nil, zero value otherwise.

### GetOfferOk

`func (o *Intent) GetOfferOk() (*Offer, bool)`

GetOfferOk returns a tuple with the Offer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffer

`func (o *Intent) SetOffer(v Offer)`

SetOffer sets Offer field to given value.

### HasOffer

`func (o *Intent) HasOffer() bool`

HasOffer returns a boolean if a field has been set.

### GetItem

`func (o *Intent) GetItem() Item`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *Intent) GetItemOk() (*Item, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *Intent) SetItem(v Item)`

SetItem sets Item field to given value.

### HasItem

`func (o *Intent) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetOndcOrgPayloadDetails

`func (o *Intent) GetOndcOrgPayloadDetails() IntentOndcOrgPayloadDetails`

GetOndcOrgPayloadDetails returns the OndcOrgPayloadDetails field if non-nil, zero value otherwise.

### GetOndcOrgPayloadDetailsOk

`func (o *Intent) GetOndcOrgPayloadDetailsOk() (*IntentOndcOrgPayloadDetails, bool)`

GetOndcOrgPayloadDetailsOk returns a tuple with the OndcOrgPayloadDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgPayloadDetails

`func (o *Intent) SetOndcOrgPayloadDetails(v IntentOndcOrgPayloadDetails)`

SetOndcOrgPayloadDetails sets OndcOrgPayloadDetails field to given value.

### HasOndcOrgPayloadDetails

`func (o *Intent) HasOndcOrgPayloadDetails() bool`

HasOndcOrgPayloadDetails returns a boolean if a field has been set.

### GetTags

`func (o *Intent) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Intent) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Intent) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Intent) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


