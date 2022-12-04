# Feedback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeedbackForm** | Pointer to [**[]FeedbackFormElement**](FeedbackFormElement.md) | Describes a feedback form that a Seller App can send to get feedback from the Buyer App | [optional] 
**FeedbackUrl** | Pointer to [**FeedbackUrl**](FeedbackUrl.md) |  | [optional] 

## Methods

### NewFeedback

`func NewFeedback() *Feedback`

NewFeedback instantiates a new Feedback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackWithDefaults

`func NewFeedbackWithDefaults() *Feedback`

NewFeedbackWithDefaults instantiates a new Feedback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeedbackForm

`func (o *Feedback) GetFeedbackForm() []FeedbackFormElement`

GetFeedbackForm returns the FeedbackForm field if non-nil, zero value otherwise.

### GetFeedbackFormOk

`func (o *Feedback) GetFeedbackFormOk() (*[]FeedbackFormElement, bool)`

GetFeedbackFormOk returns a tuple with the FeedbackForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackForm

`func (o *Feedback) SetFeedbackForm(v []FeedbackFormElement)`

SetFeedbackForm sets FeedbackForm field to given value.

### HasFeedbackForm

`func (o *Feedback) HasFeedbackForm() bool`

HasFeedbackForm returns a boolean if a field has been set.

### GetFeedbackUrl

`func (o *Feedback) GetFeedbackUrl() FeedbackUrl`

GetFeedbackUrl returns the FeedbackUrl field if non-nil, zero value otherwise.

### GetFeedbackUrlOk

`func (o *Feedback) GetFeedbackUrlOk() (*FeedbackUrl, bool)`

GetFeedbackUrlOk returns a tuple with the FeedbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackUrl

`func (o *Feedback) SetFeedbackUrl(v FeedbackUrl)`

SetFeedbackUrl sets FeedbackUrl field to given value.

### HasFeedbackUrl

`func (o *Feedback) HasFeedbackUrl() bool`

HasFeedbackUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


