# OnInitPostRequestMessageOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to [**OnInitPostRequestMessageOrderProvider**](OnInitPostRequestMessageOrderProvider.md) |  | [optional] 
**ProviderLocation** | Pointer to [**OnInitPostRequestMessageOrderProviderLocation**](OnInitPostRequestMessageOrderProviderLocation.md) |  | [optional] 
**Items** | Pointer to [**[]OnInitPostRequestMessageOrderItemsInner**](OnInitPostRequestMessageOrderItemsInner.md) |  | [optional] 
**AddOns** | Pointer to [**[]OnInitPostRequestMessageOrderAddOnsInner**](OnInitPostRequestMessageOrderAddOnsInner.md) |  | [optional] 
**Offers** | Pointer to [**[]OnInitPostRequestMessageOrderOffersInner**](OnInitPostRequestMessageOrderOffersInner.md) |  | [optional] 
**Billing** | Pointer to [**Billing**](Billing.md) |  | [optional] 
**Fulfillment** | Pointer to [**Fulfillment**](Fulfillment.md) |  | [optional] 
**Quote** | Pointer to [**Quotation**](Quotation.md) |  | [optional] 
**Payment** | Pointer to [**Payment**](Payment.md) |  | [optional] 

## Methods

### NewOnInitPostRequestMessageOrder

`func NewOnInitPostRequestMessageOrder() *OnInitPostRequestMessageOrder`

NewOnInitPostRequestMessageOrder instantiates a new OnInitPostRequestMessageOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnInitPostRequestMessageOrderWithDefaults

`func NewOnInitPostRequestMessageOrderWithDefaults() *OnInitPostRequestMessageOrder`

NewOnInitPostRequestMessageOrderWithDefaults instantiates a new OnInitPostRequestMessageOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *OnInitPostRequestMessageOrder) GetProvider() OnInitPostRequestMessageOrderProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *OnInitPostRequestMessageOrder) GetProviderOk() (*OnInitPostRequestMessageOrderProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *OnInitPostRequestMessageOrder) SetProvider(v OnInitPostRequestMessageOrderProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *OnInitPostRequestMessageOrder) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetProviderLocation

`func (o *OnInitPostRequestMessageOrder) GetProviderLocation() OnInitPostRequestMessageOrderProviderLocation`

GetProviderLocation returns the ProviderLocation field if non-nil, zero value otherwise.

### GetProviderLocationOk

`func (o *OnInitPostRequestMessageOrder) GetProviderLocationOk() (*OnInitPostRequestMessageOrderProviderLocation, bool)`

GetProviderLocationOk returns a tuple with the ProviderLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderLocation

`func (o *OnInitPostRequestMessageOrder) SetProviderLocation(v OnInitPostRequestMessageOrderProviderLocation)`

SetProviderLocation sets ProviderLocation field to given value.

### HasProviderLocation

`func (o *OnInitPostRequestMessageOrder) HasProviderLocation() bool`

HasProviderLocation returns a boolean if a field has been set.

### GetItems

`func (o *OnInitPostRequestMessageOrder) GetItems() []OnInitPostRequestMessageOrderItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *OnInitPostRequestMessageOrder) GetItemsOk() (*[]OnInitPostRequestMessageOrderItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *OnInitPostRequestMessageOrder) SetItems(v []OnInitPostRequestMessageOrderItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *OnInitPostRequestMessageOrder) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetAddOns

`func (o *OnInitPostRequestMessageOrder) GetAddOns() []OnInitPostRequestMessageOrderAddOnsInner`

GetAddOns returns the AddOns field if non-nil, zero value otherwise.

### GetAddOnsOk

`func (o *OnInitPostRequestMessageOrder) GetAddOnsOk() (*[]OnInitPostRequestMessageOrderAddOnsInner, bool)`

GetAddOnsOk returns a tuple with the AddOns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOns

`func (o *OnInitPostRequestMessageOrder) SetAddOns(v []OnInitPostRequestMessageOrderAddOnsInner)`

SetAddOns sets AddOns field to given value.

### HasAddOns

`func (o *OnInitPostRequestMessageOrder) HasAddOns() bool`

HasAddOns returns a boolean if a field has been set.

### GetOffers

`func (o *OnInitPostRequestMessageOrder) GetOffers() []OnInitPostRequestMessageOrderOffersInner`

GetOffers returns the Offers field if non-nil, zero value otherwise.

### GetOffersOk

`func (o *OnInitPostRequestMessageOrder) GetOffersOk() (*[]OnInitPostRequestMessageOrderOffersInner, bool)`

GetOffersOk returns a tuple with the Offers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffers

`func (o *OnInitPostRequestMessageOrder) SetOffers(v []OnInitPostRequestMessageOrderOffersInner)`

SetOffers sets Offers field to given value.

### HasOffers

`func (o *OnInitPostRequestMessageOrder) HasOffers() bool`

HasOffers returns a boolean if a field has been set.

### GetBilling

`func (o *OnInitPostRequestMessageOrder) GetBilling() Billing`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *OnInitPostRequestMessageOrder) GetBillingOk() (*Billing, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *OnInitPostRequestMessageOrder) SetBilling(v Billing)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *OnInitPostRequestMessageOrder) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetFulfillment

`func (o *OnInitPostRequestMessageOrder) GetFulfillment() Fulfillment`

GetFulfillment returns the Fulfillment field if non-nil, zero value otherwise.

### GetFulfillmentOk

`func (o *OnInitPostRequestMessageOrder) GetFulfillmentOk() (*Fulfillment, bool)`

GetFulfillmentOk returns a tuple with the Fulfillment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillment

`func (o *OnInitPostRequestMessageOrder) SetFulfillment(v Fulfillment)`

SetFulfillment sets Fulfillment field to given value.

### HasFulfillment

`func (o *OnInitPostRequestMessageOrder) HasFulfillment() bool`

HasFulfillment returns a boolean if a field has been set.

### GetQuote

`func (o *OnInitPostRequestMessageOrder) GetQuote() Quotation`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *OnInitPostRequestMessageOrder) GetQuoteOk() (*Quotation, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *OnInitPostRequestMessageOrder) SetQuote(v Quotation)`

SetQuote sets Quote field to given value.

### HasQuote

`func (o *OnInitPostRequestMessageOrder) HasQuote() bool`

HasQuote returns a boolean if a field has been set.

### GetPayment

`func (o *OnInitPostRequestMessageOrder) GetPayment() Payment`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *OnInitPostRequestMessageOrder) GetPaymentOk() (*Payment, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *OnInitPostRequestMessageOrder) SetPayment(v Payment)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *OnInitPostRequestMessageOrder) HasPayment() bool`

HasPayment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


