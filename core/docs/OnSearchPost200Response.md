# OnSearchPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to [**SearchPost200ResponseMessage**](SearchPost200ResponseMessage.md) |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewOnSearchPost200Response

`func NewOnSearchPost200Response() *OnSearchPost200Response`

NewOnSearchPost200Response instantiates a new OnSearchPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnSearchPost200ResponseWithDefaults

`func NewOnSearchPost200ResponseWithDefaults() *OnSearchPost200Response`

NewOnSearchPost200ResponseWithDefaults instantiates a new OnSearchPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *OnSearchPost200Response) GetMessage() SearchPost200ResponseMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OnSearchPost200Response) GetMessageOk() (*SearchPost200ResponseMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OnSearchPost200Response) SetMessage(v SearchPost200ResponseMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OnSearchPost200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetError

`func (o *OnSearchPost200Response) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OnSearchPost200Response) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OnSearchPost200Response) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *OnSearchPost200Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


