# FeedbackCategoriesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**Context**](Context.md) |  | [optional] 
**FeedbackCategories** | Pointer to [**[]RatingCategory**](RatingCategory.md) |  | [optional] 

## Methods

### NewFeedbackCategoriesPostRequest

`func NewFeedbackCategoriesPostRequest() *FeedbackCategoriesPostRequest`

NewFeedbackCategoriesPostRequest instantiates a new FeedbackCategoriesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackCategoriesPostRequestWithDefaults

`func NewFeedbackCategoriesPostRequestWithDefaults() *FeedbackCategoriesPostRequest`

NewFeedbackCategoriesPostRequestWithDefaults instantiates a new FeedbackCategoriesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *FeedbackCategoriesPostRequest) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *FeedbackCategoriesPostRequest) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *FeedbackCategoriesPostRequest) SetContext(v Context)`

SetContext sets Context field to given value.

### HasContext

`func (o *FeedbackCategoriesPostRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetFeedbackCategories

`func (o *FeedbackCategoriesPostRequest) GetFeedbackCategories() []RatingCategory`

GetFeedbackCategories returns the FeedbackCategories field if non-nil, zero value otherwise.

### GetFeedbackCategoriesOk

`func (o *FeedbackCategoriesPostRequest) GetFeedbackCategoriesOk() (*[]RatingCategory, bool)`

GetFeedbackCategoriesOk returns a tuple with the FeedbackCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackCategories

`func (o *FeedbackCategoriesPostRequest) SetFeedbackCategories(v []RatingCategory)`

SetFeedbackCategories sets FeedbackCategories field to given value.

### HasFeedbackCategories

`func (o *FeedbackCategoriesPostRequest) HasFeedbackCategories() bool`

HasFeedbackCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


