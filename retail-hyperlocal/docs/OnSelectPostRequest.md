# OnSelectPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | [**Context**](Context.md) |  | 
**Message** | Pointer to [**SelectPostRequestMessage**](SelectPostRequestMessage.md) |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewOnSelectPostRequest

`func NewOnSelectPostRequest(context Context, ) *OnSelectPostRequest`

NewOnSelectPostRequest instantiates a new OnSelectPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnSelectPostRequestWithDefaults

`func NewOnSelectPostRequestWithDefaults() *OnSelectPostRequest`

NewOnSelectPostRequestWithDefaults instantiates a new OnSelectPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *OnSelectPostRequest) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *OnSelectPostRequest) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *OnSelectPostRequest) SetContext(v Context)`

SetContext sets Context field to given value.


### GetMessage

`func (o *OnSelectPostRequest) GetMessage() SelectPostRequestMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OnSelectPostRequest) GetMessageOk() (*SelectPostRequestMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OnSelectPostRequest) SetMessage(v SelectPostRequestMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OnSelectPostRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetError

`func (o *OnSelectPostRequest) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OnSelectPostRequest) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OnSelectPostRequest) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *OnSelectPostRequest) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


