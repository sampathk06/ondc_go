# TrackPostRequestMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | [**Id**](Id.md) |  | 
**CallbackUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewTrackPostRequestMessage

`func NewTrackPostRequestMessage(orderId Id, ) *TrackPostRequestMessage`

NewTrackPostRequestMessage instantiates a new TrackPostRequestMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackPostRequestMessageWithDefaults

`func NewTrackPostRequestMessageWithDefaults() *TrackPostRequestMessage`

NewTrackPostRequestMessageWithDefaults instantiates a new TrackPostRequestMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *TrackPostRequestMessage) GetOrderId() Id`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *TrackPostRequestMessage) GetOrderIdOk() (*Id, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *TrackPostRequestMessage) SetOrderId(v Id)`

SetOrderId sets OrderId field to given value.


### GetCallbackUrl

`func (o *TrackPostRequestMessage) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *TrackPostRequestMessage) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *TrackPostRequestMessage) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *TrackPostRequestMessage) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


