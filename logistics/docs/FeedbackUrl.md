# FeedbackUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | feedback URL sent by the Seller App | [optional] 
**TlMethod** | Pointer to **string** |  | [optional] 
**Params** | Pointer to [**FeedbackUrlParams**](FeedbackUrlParams.md) |  | [optional] 

## Methods

### NewFeedbackUrl

`func NewFeedbackUrl() *FeedbackUrl`

NewFeedbackUrl instantiates a new FeedbackUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackUrlWithDefaults

`func NewFeedbackUrlWithDefaults() *FeedbackUrl`

NewFeedbackUrlWithDefaults instantiates a new FeedbackUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *FeedbackUrl) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FeedbackUrl) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FeedbackUrl) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FeedbackUrl) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetTlMethod

`func (o *FeedbackUrl) GetTlMethod() string`

GetTlMethod returns the TlMethod field if non-nil, zero value otherwise.

### GetTlMethodOk

`func (o *FeedbackUrl) GetTlMethodOk() (*string, bool)`

GetTlMethodOk returns a tuple with the TlMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlMethod

`func (o *FeedbackUrl) SetTlMethod(v string)`

SetTlMethod sets TlMethod field to given value.

### HasTlMethod

`func (o *FeedbackUrl) HasTlMethod() bool`

HasTlMethod returns a boolean if a field has been set.

### GetParams

`func (o *FeedbackUrl) GetParams() FeedbackUrlParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *FeedbackUrl) GetParamsOk() (*FeedbackUrlParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *FeedbackUrl) SetParams(v FeedbackUrlParams)`

SetParams sets Params field to given value.

### HasParams

`func (o *FeedbackUrl) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


