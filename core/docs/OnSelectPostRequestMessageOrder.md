# OnSelectPostRequestMessageOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to [**Provider**](Provider.md) |  | [optional] 
**ProviderLocation** | Pointer to [**Location**](Location.md) |  | [optional] 
**Items** | Pointer to [**[]OnSelectPostRequestMessageOrderItemsInner**](OnSelectPostRequestMessageOrderItemsInner.md) |  | [optional] 
**AddOns** | Pointer to [**[]AddOn**](AddOn.md) |  | [optional] 
**Offers** | Pointer to [**[]Offer**](Offer.md) |  | [optional] 
**Quote** | Pointer to [**Quotation**](Quotation.md) |  | [optional] 

## Methods

### NewOnSelectPostRequestMessageOrder

`func NewOnSelectPostRequestMessageOrder() *OnSelectPostRequestMessageOrder`

NewOnSelectPostRequestMessageOrder instantiates a new OnSelectPostRequestMessageOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnSelectPostRequestMessageOrderWithDefaults

`func NewOnSelectPostRequestMessageOrderWithDefaults() *OnSelectPostRequestMessageOrder`

NewOnSelectPostRequestMessageOrderWithDefaults instantiates a new OnSelectPostRequestMessageOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *OnSelectPostRequestMessageOrder) GetProvider() Provider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *OnSelectPostRequestMessageOrder) GetProviderOk() (*Provider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *OnSelectPostRequestMessageOrder) SetProvider(v Provider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *OnSelectPostRequestMessageOrder) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetProviderLocation

`func (o *OnSelectPostRequestMessageOrder) GetProviderLocation() Location`

GetProviderLocation returns the ProviderLocation field if non-nil, zero value otherwise.

### GetProviderLocationOk

`func (o *OnSelectPostRequestMessageOrder) GetProviderLocationOk() (*Location, bool)`

GetProviderLocationOk returns a tuple with the ProviderLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderLocation

`func (o *OnSelectPostRequestMessageOrder) SetProviderLocation(v Location)`

SetProviderLocation sets ProviderLocation field to given value.

### HasProviderLocation

`func (o *OnSelectPostRequestMessageOrder) HasProviderLocation() bool`

HasProviderLocation returns a boolean if a field has been set.

### GetItems

`func (o *OnSelectPostRequestMessageOrder) GetItems() []OnSelectPostRequestMessageOrderItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *OnSelectPostRequestMessageOrder) GetItemsOk() (*[]OnSelectPostRequestMessageOrderItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *OnSelectPostRequestMessageOrder) SetItems(v []OnSelectPostRequestMessageOrderItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *OnSelectPostRequestMessageOrder) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetAddOns

`func (o *OnSelectPostRequestMessageOrder) GetAddOns() []AddOn`

GetAddOns returns the AddOns field if non-nil, zero value otherwise.

### GetAddOnsOk

`func (o *OnSelectPostRequestMessageOrder) GetAddOnsOk() (*[]AddOn, bool)`

GetAddOnsOk returns a tuple with the AddOns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOns

`func (o *OnSelectPostRequestMessageOrder) SetAddOns(v []AddOn)`

SetAddOns sets AddOns field to given value.

### HasAddOns

`func (o *OnSelectPostRequestMessageOrder) HasAddOns() bool`

HasAddOns returns a boolean if a field has been set.

### GetOffers

`func (o *OnSelectPostRequestMessageOrder) GetOffers() []Offer`

GetOffers returns the Offers field if non-nil, zero value otherwise.

### GetOffersOk

`func (o *OnSelectPostRequestMessageOrder) GetOffersOk() (*[]Offer, bool)`

GetOffersOk returns a tuple with the Offers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffers

`func (o *OnSelectPostRequestMessageOrder) SetOffers(v []Offer)`

SetOffers sets Offers field to given value.

### HasOffers

`func (o *OnSelectPostRequestMessageOrder) HasOffers() bool`

HasOffers returns a boolean if a field has been set.

### GetQuote

`func (o *OnSelectPostRequestMessageOrder) GetQuote() Quotation`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *OnSelectPostRequestMessageOrder) GetQuoteOk() (*Quotation, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *OnSelectPostRequestMessageOrder) SetQuote(v Quotation)`

SetQuote sets Quote field to given value.

### HasQuote

`func (o *OnSelectPostRequestMessageOrder) HasQuote() bool`

HasQuote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


