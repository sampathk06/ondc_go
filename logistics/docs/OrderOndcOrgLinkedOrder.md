# OrderOndcOrgLinkedOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]OrderOndcOrgLinkedOrderItemsInner**](OrderOndcOrgLinkedOrderItemsInner.md) |  | [optional] 
**Provider** | Pointer to [**OrderOndcOrgLinkedOrderProvider**](OrderOndcOrgLinkedOrderProvider.md) |  | [optional] 
**Order** | Pointer to [**OrderOndcOrgLinkedOrderOrder**](OrderOndcOrgLinkedOrderOrder.md) |  | [optional] 

## Methods

### NewOrderOndcOrgLinkedOrder

`func NewOrderOndcOrgLinkedOrder() *OrderOndcOrgLinkedOrder`

NewOrderOndcOrgLinkedOrder instantiates a new OrderOndcOrgLinkedOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderOndcOrgLinkedOrderWithDefaults

`func NewOrderOndcOrgLinkedOrderWithDefaults() *OrderOndcOrgLinkedOrder`

NewOrderOndcOrgLinkedOrderWithDefaults instantiates a new OrderOndcOrgLinkedOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *OrderOndcOrgLinkedOrder) GetItems() []OrderOndcOrgLinkedOrderItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *OrderOndcOrgLinkedOrder) GetItemsOk() (*[]OrderOndcOrgLinkedOrderItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *OrderOndcOrgLinkedOrder) SetItems(v []OrderOndcOrgLinkedOrderItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *OrderOndcOrgLinkedOrder) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetProvider

`func (o *OrderOndcOrgLinkedOrder) GetProvider() OrderOndcOrgLinkedOrderProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *OrderOndcOrgLinkedOrder) GetProviderOk() (*OrderOndcOrgLinkedOrderProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *OrderOndcOrgLinkedOrder) SetProvider(v OrderOndcOrgLinkedOrderProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *OrderOndcOrgLinkedOrder) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetOrder

`func (o *OrderOndcOrgLinkedOrder) GetOrder() OrderOndcOrgLinkedOrderOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *OrderOndcOrgLinkedOrder) GetOrderOk() (*OrderOndcOrgLinkedOrderOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *OrderOndcOrgLinkedOrder) SetOrder(v OrderOndcOrgLinkedOrderOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *OrderOndcOrgLinkedOrder) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


