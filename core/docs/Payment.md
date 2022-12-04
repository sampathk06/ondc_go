# Payment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | Pointer to **string** | A payment uri to be called by the Buyer App. If empty, then the payment is to be done offline. The details of payment should be present in the params object. If &#x60;&#x60;&#x60;tl_method&#x60;&#x60;&#x60; &#x3D; http/get, then the payment details will be sent as url params. Two url param values, &#x60;&#x60;&#x60;$transaction_id&#x60;&#x60;&#x60; and &#x60;&#x60;&#x60;$amount&#x60;&#x60;&#x60; are mandatory. And example url would be : https://www.example.com/pay?txid&#x3D;$transaction_id&amp;amount&#x3D;$amount&amp;vpa&#x3D;upiid&amp;payee&#x3D;shopez&amp;billno&#x3D;1234 | [optional] 
**TlMethod** | Pointer to **string** |  | [optional] 
**Params** | Pointer to [**PaymentParams**](PaymentParams.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Time** | Pointer to [**Time**](Time.md) |  | [optional] 
**CollectedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewPayment

`func NewPayment() *Payment`

NewPayment instantiates a new Payment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentWithDefaults

`func NewPaymentWithDefaults() *Payment`

NewPaymentWithDefaults instantiates a new Payment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUri

`func (o *Payment) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Payment) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Payment) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *Payment) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetTlMethod

`func (o *Payment) GetTlMethod() string`

GetTlMethod returns the TlMethod field if non-nil, zero value otherwise.

### GetTlMethodOk

`func (o *Payment) GetTlMethodOk() (*string, bool)`

GetTlMethodOk returns a tuple with the TlMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlMethod

`func (o *Payment) SetTlMethod(v string)`

SetTlMethod sets TlMethod field to given value.

### HasTlMethod

`func (o *Payment) HasTlMethod() bool`

HasTlMethod returns a boolean if a field has been set.

### GetParams

`func (o *Payment) GetParams() PaymentParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *Payment) GetParamsOk() (*PaymentParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *Payment) SetParams(v PaymentParams)`

SetParams sets Params field to given value.

### HasParams

`func (o *Payment) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetType

`func (o *Payment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Payment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Payment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Payment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *Payment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Payment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Payment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Payment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTime

`func (o *Payment) GetTime() Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Payment) GetTimeOk() (*Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Payment) SetTime(v Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *Payment) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetCollectedBy

`func (o *Payment) GetCollectedBy() string`

GetCollectedBy returns the CollectedBy field if non-nil, zero value otherwise.

### GetCollectedByOk

`func (o *Payment) GetCollectedByOk() (*string, bool)`

GetCollectedByOk returns a tuple with the CollectedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectedBy

`func (o *Payment) SetCollectedBy(v string)`

SetCollectedBy sets CollectedBy field to given value.

### HasCollectedBy

`func (o *Payment) HasCollectedBy() bool`

HasCollectedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


