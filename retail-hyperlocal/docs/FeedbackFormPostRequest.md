# FeedbackFormPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**Context**](Context.md) |  | [optional] 
**Message** | Pointer to [**Feedback**](Feedback.md) |  | [optional] 

## Methods

### NewFeedbackFormPostRequest

`func NewFeedbackFormPostRequest() *FeedbackFormPostRequest`

NewFeedbackFormPostRequest instantiates a new FeedbackFormPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackFormPostRequestWithDefaults

`func NewFeedbackFormPostRequestWithDefaults() *FeedbackFormPostRequest`

NewFeedbackFormPostRequestWithDefaults instantiates a new FeedbackFormPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *FeedbackFormPostRequest) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *FeedbackFormPostRequest) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *FeedbackFormPostRequest) SetContext(v Context)`

SetContext sets Context field to given value.

### HasContext

`func (o *FeedbackFormPostRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetMessage

`func (o *FeedbackFormPostRequest) GetMessage() Feedback`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FeedbackFormPostRequest) GetMessageOk() (*Feedback, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FeedbackFormPostRequest) SetMessage(v Feedback)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FeedbackFormPostRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


