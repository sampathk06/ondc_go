# ReturnReasonsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**Context**](Context.md) |  | [optional] 
**ReturnReasons** | Pointer to [**[]Option**](Option.md) |  | [optional] 

## Methods

### NewReturnReasonsPostRequest

`func NewReturnReasonsPostRequest() *ReturnReasonsPostRequest`

NewReturnReasonsPostRequest instantiates a new ReturnReasonsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnReasonsPostRequestWithDefaults

`func NewReturnReasonsPostRequestWithDefaults() *ReturnReasonsPostRequest`

NewReturnReasonsPostRequestWithDefaults instantiates a new ReturnReasonsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *ReturnReasonsPostRequest) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ReturnReasonsPostRequest) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ReturnReasonsPostRequest) SetContext(v Context)`

SetContext sets Context field to given value.

### HasContext

`func (o *ReturnReasonsPostRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetReturnReasons

`func (o *ReturnReasonsPostRequest) GetReturnReasons() []Option`

GetReturnReasons returns the ReturnReasons field if non-nil, zero value otherwise.

### GetReturnReasonsOk

`func (o *ReturnReasonsPostRequest) GetReturnReasonsOk() (*[]Option, bool)`

GetReturnReasonsOk returns a tuple with the ReturnReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnReasons

`func (o *ReturnReasonsPostRequest) SetReturnReasons(v []Option)`

SetReturnReasons sets ReturnReasons field to given value.

### HasReturnReasons

`func (o *ReturnReasonsPostRequest) HasReturnReasons() bool`

HasReturnReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


