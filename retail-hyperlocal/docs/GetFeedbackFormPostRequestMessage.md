# GetFeedbackFormPostRequestMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RatingValue** | Pointer to [**Value**](Value.md) |  | [optional] 
**RatingCategory** | Pointer to [**RatingCategory**](RatingCategory.md) |  | [optional] 

## Methods

### NewGetFeedbackFormPostRequestMessage

`func NewGetFeedbackFormPostRequestMessage() *GetFeedbackFormPostRequestMessage`

NewGetFeedbackFormPostRequestMessage instantiates a new GetFeedbackFormPostRequestMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFeedbackFormPostRequestMessageWithDefaults

`func NewGetFeedbackFormPostRequestMessageWithDefaults() *GetFeedbackFormPostRequestMessage`

NewGetFeedbackFormPostRequestMessageWithDefaults instantiates a new GetFeedbackFormPostRequestMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRatingValue

`func (o *GetFeedbackFormPostRequestMessage) GetRatingValue() Value`

GetRatingValue returns the RatingValue field if non-nil, zero value otherwise.

### GetRatingValueOk

`func (o *GetFeedbackFormPostRequestMessage) GetRatingValueOk() (*Value, bool)`

GetRatingValueOk returns a tuple with the RatingValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingValue

`func (o *GetFeedbackFormPostRequestMessage) SetRatingValue(v Value)`

SetRatingValue sets RatingValue field to given value.

### HasRatingValue

`func (o *GetFeedbackFormPostRequestMessage) HasRatingValue() bool`

HasRatingValue returns a boolean if a field has been set.

### GetRatingCategory

`func (o *GetFeedbackFormPostRequestMessage) GetRatingCategory() RatingCategory`

GetRatingCategory returns the RatingCategory field if non-nil, zero value otherwise.

### GetRatingCategoryOk

`func (o *GetFeedbackFormPostRequestMessage) GetRatingCategoryOk() (*RatingCategory, bool)`

GetRatingCategoryOk returns a tuple with the RatingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingCategory

`func (o *GetFeedbackFormPostRequestMessage) SetRatingCategory(v RatingCategory)`

SetRatingCategory sets RatingCategory field to given value.

### HasRatingCategory

`func (o *GetFeedbackFormPostRequestMessage) HasRatingCategory() bool`

HasRatingCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


