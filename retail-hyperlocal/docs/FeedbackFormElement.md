# FeedbackFormElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to [**Id**](Id.md) |  | [optional] 
**Question** | Pointer to **string** | Specifies the question to which the answer options will be contained in the child FeedbackFormElements | [optional] 
**Answer** | Pointer to **string** | Specifies an answer option to which the question will be in the FeedbackFormElement specified in parent_id | [optional] 
**AnswerType** | Pointer to **string** | Specifies how the answer option should be rendered. | [optional] 

## Methods

### NewFeedbackFormElement

`func NewFeedbackFormElement() *FeedbackFormElement`

NewFeedbackFormElement instantiates a new FeedbackFormElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackFormElementWithDefaults

`func NewFeedbackFormElementWithDefaults() *FeedbackFormElement`

NewFeedbackFormElementWithDefaults instantiates a new FeedbackFormElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FeedbackFormElement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FeedbackFormElement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FeedbackFormElement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FeedbackFormElement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParentId

`func (o *FeedbackFormElement) GetParentId() Id`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *FeedbackFormElement) GetParentIdOk() (*Id, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *FeedbackFormElement) SetParentId(v Id)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *FeedbackFormElement) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetQuestion

`func (o *FeedbackFormElement) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *FeedbackFormElement) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *FeedbackFormElement) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *FeedbackFormElement) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetAnswer

`func (o *FeedbackFormElement) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *FeedbackFormElement) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *FeedbackFormElement) SetAnswer(v string)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *FeedbackFormElement) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetAnswerType

`func (o *FeedbackFormElement) GetAnswerType() string`

GetAnswerType returns the AnswerType field if non-nil, zero value otherwise.

### GetAnswerTypeOk

`func (o *FeedbackFormElement) GetAnswerTypeOk() (*string, bool)`

GetAnswerTypeOk returns a tuple with the AnswerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerType

`func (o *FeedbackFormElement) SetAnswerType(v string)`

SetAnswerType sets AnswerType field to given value.

### HasAnswerType

`func (o *FeedbackFormElement) HasAnswerType() bool`

HasAnswerType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


