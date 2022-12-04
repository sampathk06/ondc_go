# SelectPostRequestMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | [**Order**](Order.md) |  | 

## Methods

### NewSelectPostRequestMessage

`func NewSelectPostRequestMessage(order Order, ) *SelectPostRequestMessage`

NewSelectPostRequestMessage instantiates a new SelectPostRequestMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectPostRequestMessageWithDefaults

`func NewSelectPostRequestMessageWithDefaults() *SelectPostRequestMessage`

NewSelectPostRequestMessageWithDefaults instantiates a new SelectPostRequestMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *SelectPostRequestMessage) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *SelectPostRequestMessage) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *SelectPostRequestMessage) SetOrder(v Order)`

SetOrder sets Order field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


