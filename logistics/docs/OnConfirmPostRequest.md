# OnConfirmPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | [**Context**](Context.md) |  | 
**Message** | Pointer to [**SelectPostRequestMessage**](SelectPostRequestMessage.md) |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewOnConfirmPostRequest

`func NewOnConfirmPostRequest(context Context, ) *OnConfirmPostRequest`

NewOnConfirmPostRequest instantiates a new OnConfirmPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnConfirmPostRequestWithDefaults

`func NewOnConfirmPostRequestWithDefaults() *OnConfirmPostRequest`

NewOnConfirmPostRequestWithDefaults instantiates a new OnConfirmPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *OnConfirmPostRequest) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *OnConfirmPostRequest) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *OnConfirmPostRequest) SetContext(v Context)`

SetContext sets Context field to given value.


### GetMessage

`func (o *OnConfirmPostRequest) GetMessage() SelectPostRequestMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OnConfirmPostRequest) GetMessageOk() (*SelectPostRequestMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OnConfirmPostRequest) SetMessage(v SelectPostRequestMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OnConfirmPostRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetError

`func (o *OnConfirmPostRequest) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OnConfirmPostRequest) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OnConfirmPostRequest) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *OnConfirmPostRequest) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


