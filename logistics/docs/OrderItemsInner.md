# OrderItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**Id**](Id.md) |  | 
**CategoryId** | Pointer to [**Id**](Id.md) |  | [optional] 
**Quantity** | Pointer to [**Selected**](Selected.md) |  | [optional] 

## Methods

### NewOrderItemsInner

`func NewOrderItemsInner(id Id, ) *OrderItemsInner`

NewOrderItemsInner instantiates a new OrderItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderItemsInnerWithDefaults

`func NewOrderItemsInnerWithDefaults() *OrderItemsInner`

NewOrderItemsInnerWithDefaults instantiates a new OrderItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderItemsInner) GetId() Id`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderItemsInner) GetIdOk() (*Id, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderItemsInner) SetId(v Id)`

SetId sets Id field to given value.


### GetCategoryId

`func (o *OrderItemsInner) GetCategoryId() Id`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *OrderItemsInner) GetCategoryIdOk() (*Id, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *OrderItemsInner) SetCategoryId(v Id)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *OrderItemsInner) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetQuantity

`func (o *OrderItemsInner) GetQuantity() Selected`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderItemsInner) GetQuantityOk() (*Selected, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderItemsInner) SetQuantity(v Selected)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrderItemsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


