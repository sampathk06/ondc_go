# Rating

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RatingCategory** | Pointer to **string** | Category of the object being rated | [optional] 
**Id** | Pointer to **string** | Id of the object being rated | [optional] 
**Value** | Pointer to **float32** | Rating value given to the object (1 - Poor; 2 - Needs improvement; 3 - Satisfactory; 4 - Good; 5 - Excellent) | [optional] 
**FeedbackForm** | Pointer to [**[]FeedbackFormElement**](FeedbackFormElement.md) | Describes a feedback form that a Seller App can send to get feedback from the Buyer App | [optional] 
**FeedbackId** | Pointer to [**FeedbackId**](FeedbackId.md) |  | [optional] 

## Methods

### NewRating

`func NewRating() *Rating`

NewRating instantiates a new Rating object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRatingWithDefaults

`func NewRatingWithDefaults() *Rating`

NewRatingWithDefaults instantiates a new Rating object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRatingCategory

`func (o *Rating) GetRatingCategory() string`

GetRatingCategory returns the RatingCategory field if non-nil, zero value otherwise.

### GetRatingCategoryOk

`func (o *Rating) GetRatingCategoryOk() (*string, bool)`

GetRatingCategoryOk returns a tuple with the RatingCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingCategory

`func (o *Rating) SetRatingCategory(v string)`

SetRatingCategory sets RatingCategory field to given value.

### HasRatingCategory

`func (o *Rating) HasRatingCategory() bool`

HasRatingCategory returns a boolean if a field has been set.

### GetId

`func (o *Rating) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Rating) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Rating) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Rating) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValue

`func (o *Rating) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Rating) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Rating) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *Rating) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetFeedbackForm

`func (o *Rating) GetFeedbackForm() []FeedbackFormElement`

GetFeedbackForm returns the FeedbackForm field if non-nil, zero value otherwise.

### GetFeedbackFormOk

`func (o *Rating) GetFeedbackFormOk() (*[]FeedbackFormElement, bool)`

GetFeedbackFormOk returns a tuple with the FeedbackForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackForm

`func (o *Rating) SetFeedbackForm(v []FeedbackFormElement)`

SetFeedbackForm sets FeedbackForm field to given value.

### HasFeedbackForm

`func (o *Rating) HasFeedbackForm() bool`

HasFeedbackForm returns a boolean if a field has been set.

### GetFeedbackId

`func (o *Rating) GetFeedbackId() FeedbackId`

GetFeedbackId returns the FeedbackId field if non-nil, zero value otherwise.

### GetFeedbackIdOk

`func (o *Rating) GetFeedbackIdOk() (*FeedbackId, bool)`

GetFeedbackIdOk returns a tuple with the FeedbackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackId

`func (o *Rating) SetFeedbackId(v FeedbackId)`

SetFeedbackId sets FeedbackId field to given value.

### HasFeedbackId

`func (o *Rating) HasFeedbackId() bool`

HasFeedbackId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


