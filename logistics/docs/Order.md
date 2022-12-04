# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for Order across network, will be created by buyer app in confirm API | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to [**OrderProvider**](OrderProvider.md) |  | [optional] 
**Items** | Pointer to [**[]OrderItemsInner**](OrderItemsInner.md) |  | [optional] 
**AddOns** | Pointer to [**[]OrderAddOnsInner**](OrderAddOnsInner.md) |  | [optional] 
**Offers** | Pointer to [**[]OrderOffersInner**](OrderOffersInner.md) |  | [optional] 
**Documents** | Pointer to [**[]Document**](Document.md) |  | [optional] 
**Billing** | Pointer to [**Billing**](Billing.md) |  | [optional] 
**Fulfillments** | Pointer to [**[]Fulfillment**](Fulfillment.md) |  | [optional] 
**Quote** | Pointer to [**Quotation**](Quotation.md) |  | [optional] 
**Payment** | Pointer to [**Payment**](Payment.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**OndcOrgCreatedBy** | Pointer to **string** | order created by | [optional] 
**OndcOrgCancellation** | Pointer to [**Cancellation**](Cancellation.md) |  | [optional] 
**OndcOrgLinkedOrder** | Pointer to [**OrderOndcOrgLinkedOrder**](OrderOndcOrgLinkedOrder.md) |  | [optional] 

## Methods

### NewOrder

`func NewOrder() *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Order) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Order) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Order) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Order) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *Order) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Order) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Order) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Order) HasState() bool`

HasState returns a boolean if a field has been set.

### GetProvider

`func (o *Order) GetProvider() OrderProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Order) GetProviderOk() (*OrderProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Order) SetProvider(v OrderProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *Order) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetItems

`func (o *Order) GetItems() []OrderItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Order) GetItemsOk() (*[]OrderItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Order) SetItems(v []OrderItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *Order) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetAddOns

`func (o *Order) GetAddOns() []OrderAddOnsInner`

GetAddOns returns the AddOns field if non-nil, zero value otherwise.

### GetAddOnsOk

`func (o *Order) GetAddOnsOk() (*[]OrderAddOnsInner, bool)`

GetAddOnsOk returns a tuple with the AddOns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddOns

`func (o *Order) SetAddOns(v []OrderAddOnsInner)`

SetAddOns sets AddOns field to given value.

### HasAddOns

`func (o *Order) HasAddOns() bool`

HasAddOns returns a boolean if a field has been set.

### GetOffers

`func (o *Order) GetOffers() []OrderOffersInner`

GetOffers returns the Offers field if non-nil, zero value otherwise.

### GetOffersOk

`func (o *Order) GetOffersOk() (*[]OrderOffersInner, bool)`

GetOffersOk returns a tuple with the Offers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffers

`func (o *Order) SetOffers(v []OrderOffersInner)`

SetOffers sets Offers field to given value.

### HasOffers

`func (o *Order) HasOffers() bool`

HasOffers returns a boolean if a field has been set.

### GetDocuments

`func (o *Order) GetDocuments() []Document`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *Order) GetDocumentsOk() (*[]Document, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *Order) SetDocuments(v []Document)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *Order) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetBilling

`func (o *Order) GetBilling() Billing`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *Order) GetBillingOk() (*Billing, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *Order) SetBilling(v Billing)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *Order) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetFulfillments

`func (o *Order) GetFulfillments() []Fulfillment`

GetFulfillments returns the Fulfillments field if non-nil, zero value otherwise.

### GetFulfillmentsOk

`func (o *Order) GetFulfillmentsOk() (*[]Fulfillment, bool)`

GetFulfillmentsOk returns a tuple with the Fulfillments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillments

`func (o *Order) SetFulfillments(v []Fulfillment)`

SetFulfillments sets Fulfillments field to given value.

### HasFulfillments

`func (o *Order) HasFulfillments() bool`

HasFulfillments returns a boolean if a field has been set.

### GetQuote

`func (o *Order) GetQuote() Quotation`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *Order) GetQuoteOk() (*Quotation, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *Order) SetQuote(v Quotation)`

SetQuote sets Quote field to given value.

### HasQuote

`func (o *Order) HasQuote() bool`

HasQuote returns a boolean if a field has been set.

### GetPayment

`func (o *Order) GetPayment() Payment`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *Order) GetPaymentOk() (*Payment, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *Order) SetPayment(v Payment)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *Order) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Order) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Order) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Order) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Order) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Order) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Order) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Order) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Order) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetOndcOrgCreatedBy

`func (o *Order) GetOndcOrgCreatedBy() string`

GetOndcOrgCreatedBy returns the OndcOrgCreatedBy field if non-nil, zero value otherwise.

### GetOndcOrgCreatedByOk

`func (o *Order) GetOndcOrgCreatedByOk() (*string, bool)`

GetOndcOrgCreatedByOk returns a tuple with the OndcOrgCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgCreatedBy

`func (o *Order) SetOndcOrgCreatedBy(v string)`

SetOndcOrgCreatedBy sets OndcOrgCreatedBy field to given value.

### HasOndcOrgCreatedBy

`func (o *Order) HasOndcOrgCreatedBy() bool`

HasOndcOrgCreatedBy returns a boolean if a field has been set.

### GetOndcOrgCancellation

`func (o *Order) GetOndcOrgCancellation() Cancellation`

GetOndcOrgCancellation returns the OndcOrgCancellation field if non-nil, zero value otherwise.

### GetOndcOrgCancellationOk

`func (o *Order) GetOndcOrgCancellationOk() (*Cancellation, bool)`

GetOndcOrgCancellationOk returns a tuple with the OndcOrgCancellation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgCancellation

`func (o *Order) SetOndcOrgCancellation(v Cancellation)`

SetOndcOrgCancellation sets OndcOrgCancellation field to given value.

### HasOndcOrgCancellation

`func (o *Order) HasOndcOrgCancellation() bool`

HasOndcOrgCancellation returns a boolean if a field has been set.

### GetOndcOrgLinkedOrder

`func (o *Order) GetOndcOrgLinkedOrder() OrderOndcOrgLinkedOrder`

GetOndcOrgLinkedOrder returns the OndcOrgLinkedOrder field if non-nil, zero value otherwise.

### GetOndcOrgLinkedOrderOk

`func (o *Order) GetOndcOrgLinkedOrderOk() (*OrderOndcOrgLinkedOrder, bool)`

GetOndcOrgLinkedOrderOk returns a tuple with the OndcOrgLinkedOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgLinkedOrder

`func (o *Order) SetOndcOrgLinkedOrder(v OrderOndcOrgLinkedOrder)`

SetOndcOrgLinkedOrder sets OndcOrgLinkedOrder field to given value.

### HasOndcOrgLinkedOrder

`func (o *Order) HasOndcOrgLinkedOrder() bool`

HasOndcOrgLinkedOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


