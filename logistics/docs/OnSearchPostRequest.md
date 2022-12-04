# OnSearchPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | [**Context**](Context.md) |  | 
**Message** | Pointer to [**OnSearchPostRequestMessage**](OnSearchPostRequestMessage.md) |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewOnSearchPostRequest

`func NewOnSearchPostRequest(context Context, ) *OnSearchPostRequest`

NewOnSearchPostRequest instantiates a new OnSearchPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnSearchPostRequestWithDefaults

`func NewOnSearchPostRequestWithDefaults() *OnSearchPostRequest`

NewOnSearchPostRequestWithDefaults instantiates a new OnSearchPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *OnSearchPostRequest) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *OnSearchPostRequest) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *OnSearchPostRequest) SetContext(v Context)`

SetContext sets Context field to given value.


### GetMessage

`func (o *OnSearchPostRequest) GetMessage() OnSearchPostRequestMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OnSearchPostRequest) GetMessageOk() (*OnSearchPostRequestMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OnSearchPostRequest) SetMessage(v OnSearchPostRequestMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OnSearchPostRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetError

`func (o *OnSearchPostRequest) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OnSearchPostRequest) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OnSearchPostRequest) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *OnSearchPostRequest) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


