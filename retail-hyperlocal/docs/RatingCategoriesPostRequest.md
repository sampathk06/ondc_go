# RatingCategoriesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**Context**](Context.md) |  | [optional] 
**RatingCategories** | Pointer to [**[]RatingCategory**](RatingCategory.md) |  | [optional] 

## Methods

### NewRatingCategoriesPostRequest

`func NewRatingCategoriesPostRequest() *RatingCategoriesPostRequest`

NewRatingCategoriesPostRequest instantiates a new RatingCategoriesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRatingCategoriesPostRequestWithDefaults

`func NewRatingCategoriesPostRequestWithDefaults() *RatingCategoriesPostRequest`

NewRatingCategoriesPostRequestWithDefaults instantiates a new RatingCategoriesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *RatingCategoriesPostRequest) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *RatingCategoriesPostRequest) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *RatingCategoriesPostRequest) SetContext(v Context)`

SetContext sets Context field to given value.

### HasContext

`func (o *RatingCategoriesPostRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetRatingCategories

`func (o *RatingCategoriesPostRequest) GetRatingCategories() []RatingCategory`

GetRatingCategories returns the RatingCategories field if non-nil, zero value otherwise.

### GetRatingCategoriesOk

`func (o *RatingCategoriesPostRequest) GetRatingCategoriesOk() (*[]RatingCategory, bool)`

GetRatingCategoriesOk returns a tuple with the RatingCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingCategories

`func (o *RatingCategoriesPostRequest) SetRatingCategories(v []RatingCategory)`

SetRatingCategories sets RatingCategories field to given value.

### HasRatingCategories

`func (o *RatingCategoriesPostRequest) HasRatingCategories() bool`

HasRatingCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


