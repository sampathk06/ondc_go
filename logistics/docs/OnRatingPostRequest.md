# OnRatingPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | [**Context**](Context.md) |  | 
**Message** | Pointer to [**RatingAck**](RatingAck.md) |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewOnRatingPostRequest

`func NewOnRatingPostRequest(context Context, ) *OnRatingPostRequest`

NewOnRatingPostRequest instantiates a new OnRatingPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnRatingPostRequestWithDefaults

`func NewOnRatingPostRequestWithDefaults() *OnRatingPostRequest`

NewOnRatingPostRequestWithDefaults instantiates a new OnRatingPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *OnRatingPostRequest) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *OnRatingPostRequest) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *OnRatingPostRequest) SetContext(v Context)`

SetContext sets Context field to given value.


### GetMessage

`func (o *OnRatingPostRequest) GetMessage() RatingAck`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OnRatingPostRequest) GetMessageOk() (*RatingAck, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OnRatingPostRequest) SetMessage(v RatingAck)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OnRatingPostRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetError

`func (o *OnRatingPostRequest) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OnRatingPostRequest) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OnRatingPostRequest) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *OnRatingPostRequest) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


