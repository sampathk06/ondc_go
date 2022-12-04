# UpdatePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | [**Context**](Context.md) |  | 
**Message** | [**UpdatePostRequestMessage**](UpdatePostRequestMessage.md) |  | 

## Methods

### NewUpdatePostRequest

`func NewUpdatePostRequest(context Context, message UpdatePostRequestMessage, ) *UpdatePostRequest`

NewUpdatePostRequest instantiates a new UpdatePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePostRequestWithDefaults

`func NewUpdatePostRequestWithDefaults() *UpdatePostRequest`

NewUpdatePostRequestWithDefaults instantiates a new UpdatePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *UpdatePostRequest) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *UpdatePostRequest) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *UpdatePostRequest) SetContext(v Context)`

SetContext sets Context field to given value.


### GetMessage

`func (o *UpdatePostRequest) GetMessage() UpdatePostRequestMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdatePostRequest) GetMessageOk() (*UpdatePostRequestMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdatePostRequest) SetMessage(v UpdatePostRequestMessage)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


