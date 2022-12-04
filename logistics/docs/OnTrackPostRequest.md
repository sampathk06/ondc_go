# OnTrackPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | [**Context**](Context.md) |  | 
**Message** | Pointer to [**OnTrackPostRequestMessage**](OnTrackPostRequestMessage.md) |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewOnTrackPostRequest

`func NewOnTrackPostRequest(context Context, ) *OnTrackPostRequest`

NewOnTrackPostRequest instantiates a new OnTrackPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnTrackPostRequestWithDefaults

`func NewOnTrackPostRequestWithDefaults() *OnTrackPostRequest`

NewOnTrackPostRequestWithDefaults instantiates a new OnTrackPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *OnTrackPostRequest) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *OnTrackPostRequest) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *OnTrackPostRequest) SetContext(v Context)`

SetContext sets Context field to given value.


### GetMessage

`func (o *OnTrackPostRequest) GetMessage() OnTrackPostRequestMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OnTrackPostRequest) GetMessageOk() (*OnTrackPostRequestMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OnTrackPostRequest) SetMessage(v OnTrackPostRequestMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OnTrackPostRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetError

`func (o *OnTrackPostRequest) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OnTrackPostRequest) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OnTrackPostRequest) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *OnTrackPostRequest) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


