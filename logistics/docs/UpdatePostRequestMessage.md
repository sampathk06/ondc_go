# UpdatePostRequestMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdateTarget** | **string** | Comma separated values of order objects being updated. For example: &#x60;&#x60;&#x60;\&quot;update_target\&quot;:\&quot;item,billing,fulfillment\&quot;&#x60;&#x60;&#x60; | 
**Order** | [**Order**](Order.md) |  | 

## Methods

### NewUpdatePostRequestMessage

`func NewUpdatePostRequestMessage(updateTarget string, order Order, ) *UpdatePostRequestMessage`

NewUpdatePostRequestMessage instantiates a new UpdatePostRequestMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePostRequestMessageWithDefaults

`func NewUpdatePostRequestMessageWithDefaults() *UpdatePostRequestMessage`

NewUpdatePostRequestMessageWithDefaults instantiates a new UpdatePostRequestMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdateTarget

`func (o *UpdatePostRequestMessage) GetUpdateTarget() string`

GetUpdateTarget returns the UpdateTarget field if non-nil, zero value otherwise.

### GetUpdateTargetOk

`func (o *UpdatePostRequestMessage) GetUpdateTargetOk() (*string, bool)`

GetUpdateTargetOk returns a tuple with the UpdateTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTarget

`func (o *UpdatePostRequestMessage) SetUpdateTarget(v string)`

SetUpdateTarget sets UpdateTarget field to given value.


### GetOrder

`func (o *UpdatePostRequestMessage) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *UpdatePostRequestMessage) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *UpdatePostRequestMessage) SetOrder(v Order)`

SetOrder sets Order field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


