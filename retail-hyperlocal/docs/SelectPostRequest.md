# SelectPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | [**Context**](Context.md) |  | 
**Message** | [**SelectPostRequestMessage**](SelectPostRequestMessage.md) |  | 

## Methods

### NewSelectPostRequest

`func NewSelectPostRequest(context Context, message SelectPostRequestMessage, ) *SelectPostRequest`

NewSelectPostRequest instantiates a new SelectPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectPostRequestWithDefaults

`func NewSelectPostRequestWithDefaults() *SelectPostRequest`

NewSelectPostRequestWithDefaults instantiates a new SelectPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *SelectPostRequest) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *SelectPostRequest) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *SelectPostRequest) SetContext(v Context)`

SetContext sets Context field to given value.


### GetMessage

`func (o *SelectPostRequest) GetMessage() SelectPostRequestMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SelectPostRequest) GetMessageOk() (*SelectPostRequestMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SelectPostRequest) SetMessage(v SelectPostRequestMessage)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


