# FeedbackUrlParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeedbackId** | **string** | This value will be placed in the the $feedback_id url param in case of http/get and in the requestBody http/post requests | 

## Methods

### NewFeedbackUrlParams

`func NewFeedbackUrlParams(feedbackId string, ) *FeedbackUrlParams`

NewFeedbackUrlParams instantiates a new FeedbackUrlParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackUrlParamsWithDefaults

`func NewFeedbackUrlParamsWithDefaults() *FeedbackUrlParams`

NewFeedbackUrlParamsWithDefaults instantiates a new FeedbackUrlParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeedbackId

`func (o *FeedbackUrlParams) GetFeedbackId() string`

GetFeedbackId returns the FeedbackId field if non-nil, zero value otherwise.

### GetFeedbackIdOk

`func (o *FeedbackUrlParams) GetFeedbackIdOk() (*string, bool)`

GetFeedbackIdOk returns a tuple with the FeedbackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackId

`func (o *FeedbackUrlParams) SetFeedbackId(v string)`

SetFeedbackId sets FeedbackId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


