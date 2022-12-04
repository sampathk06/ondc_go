# SearchPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | [**SearchPost200ResponseMessage**](SearchPost200ResponseMessage.md) |  | 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewSearchPost200Response

`func NewSearchPost200Response(message SearchPost200ResponseMessage, ) *SearchPost200Response`

NewSearchPost200Response instantiates a new SearchPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchPost200ResponseWithDefaults

`func NewSearchPost200ResponseWithDefaults() *SearchPost200Response`

NewSearchPost200ResponseWithDefaults instantiates a new SearchPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SearchPost200Response) GetMessage() SearchPost200ResponseMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SearchPost200Response) GetMessageOk() (*SearchPost200ResponseMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SearchPost200Response) SetMessage(v SearchPost200ResponseMessage)`

SetMessage sets Message field to given value.


### GetError

`func (o *SearchPost200Response) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SearchPost200Response) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SearchPost200Response) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *SearchPost200Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


