# StatusPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | [**Context**](Context.md) |  | 
**Message** | [**StatusPostRequestMessage**](StatusPostRequestMessage.md) |  | 

## Methods

### NewStatusPostRequest

`func NewStatusPostRequest(context Context, message StatusPostRequestMessage, ) *StatusPostRequest`

NewStatusPostRequest instantiates a new StatusPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusPostRequestWithDefaults

`func NewStatusPostRequestWithDefaults() *StatusPostRequest`

NewStatusPostRequestWithDefaults instantiates a new StatusPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *StatusPostRequest) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *StatusPostRequest) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *StatusPostRequest) SetContext(v Context)`

SetContext sets Context field to given value.


### GetMessage

`func (o *StatusPostRequest) GetMessage() StatusPostRequestMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *StatusPostRequest) GetMessageOk() (*StatusPostRequestMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *StatusPostRequest) SetMessage(v StatusPostRequestMessage)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


