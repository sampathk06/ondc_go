# CancelPostRequestMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | [**Id**](Id.md) |  | 
**CancellationReasonId** | Pointer to [**Id**](Id.md) |  | [optional] 
**Descriptor** | Pointer to [**Descriptor**](Descriptor.md) |  | [optional] 

## Methods

### NewCancelPostRequestMessage

`func NewCancelPostRequestMessage(orderId Id, ) *CancelPostRequestMessage`

NewCancelPostRequestMessage instantiates a new CancelPostRequestMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelPostRequestMessageWithDefaults

`func NewCancelPostRequestMessageWithDefaults() *CancelPostRequestMessage`

NewCancelPostRequestMessageWithDefaults instantiates a new CancelPostRequestMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *CancelPostRequestMessage) GetOrderId() Id`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CancelPostRequestMessage) GetOrderIdOk() (*Id, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CancelPostRequestMessage) SetOrderId(v Id)`

SetOrderId sets OrderId field to given value.


### GetCancellationReasonId

`func (o *CancelPostRequestMessage) GetCancellationReasonId() Id`

GetCancellationReasonId returns the CancellationReasonId field if non-nil, zero value otherwise.

### GetCancellationReasonIdOk

`func (o *CancelPostRequestMessage) GetCancellationReasonIdOk() (*Id, bool)`

GetCancellationReasonIdOk returns a tuple with the CancellationReasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationReasonId

`func (o *CancelPostRequestMessage) SetCancellationReasonId(v Id)`

SetCancellationReasonId sets CancellationReasonId field to given value.

### HasCancellationReasonId

`func (o *CancelPostRequestMessage) HasCancellationReasonId() bool`

HasCancellationReasonId returns a boolean if a field has been set.

### GetDescriptor

`func (o *CancelPostRequestMessage) GetDescriptor() Descriptor`

GetDescriptor returns the Descriptor field if non-nil, zero value otherwise.

### GetDescriptorOk

`func (o *CancelPostRequestMessage) GetDescriptorOk() (*Descriptor, bool)`

GetDescriptorOk returns a tuple with the Descriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptor

`func (o *CancelPostRequestMessage) SetDescriptor(v Descriptor)`

SetDescriptor sets Descriptor field to given value.

### HasDescriptor

`func (o *CancelPostRequestMessage) HasDescriptor() bool`

HasDescriptor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


