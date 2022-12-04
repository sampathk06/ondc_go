# GetFeedbackFormPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**Context**](Context.md) |  | [optional] 
**Message** | Pointer to [**GetFeedbackFormPostRequestMessage**](GetFeedbackFormPostRequestMessage.md) |  | [optional] 

## Methods

### NewGetFeedbackFormPostRequest

`func NewGetFeedbackFormPostRequest() *GetFeedbackFormPostRequest`

NewGetFeedbackFormPostRequest instantiates a new GetFeedbackFormPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFeedbackFormPostRequestWithDefaults

`func NewGetFeedbackFormPostRequestWithDefaults() *GetFeedbackFormPostRequest`

NewGetFeedbackFormPostRequestWithDefaults instantiates a new GetFeedbackFormPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *GetFeedbackFormPostRequest) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *GetFeedbackFormPostRequest) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *GetFeedbackFormPostRequest) SetContext(v Context)`

SetContext sets Context field to given value.

### HasContext

`func (o *GetFeedbackFormPostRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetMessage

`func (o *GetFeedbackFormPostRequest) GetMessage() GetFeedbackFormPostRequestMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetFeedbackFormPostRequest) GetMessageOk() (*GetFeedbackFormPostRequestMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetFeedbackFormPostRequest) SetMessage(v GetFeedbackFormPostRequestMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetFeedbackFormPostRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


