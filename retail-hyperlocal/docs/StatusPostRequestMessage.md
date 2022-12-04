# StatusPostRequestMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | [**Id**](Id.md) |  | 

## Methods

### NewStatusPostRequestMessage

`func NewStatusPostRequestMessage(orderId Id, ) *StatusPostRequestMessage`

NewStatusPostRequestMessage instantiates a new StatusPostRequestMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusPostRequestMessageWithDefaults

`func NewStatusPostRequestMessageWithDefaults() *StatusPostRequestMessage`

NewStatusPostRequestMessageWithDefaults instantiates a new StatusPostRequestMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *StatusPostRequestMessage) GetOrderId() Id`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *StatusPostRequestMessage) GetOrderIdOk() (*Id, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *StatusPostRequestMessage) SetOrderId(v Id)`

SetOrderId sets OrderId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


