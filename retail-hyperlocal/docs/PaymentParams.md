# PaymentParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | Pointer to **string** | This value will be placed in the the $transaction_id url param in case of http/get and in the requestBody http/post requests | [optional] 
**TransactionStatus** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to [**Value**](Value.md) |  | [optional] 
**Currency** | [**Currency**](Currency.md) |  | 

## Methods

### NewPaymentParams

`func NewPaymentParams(currency Currency, ) *PaymentParams`

NewPaymentParams instantiates a new PaymentParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentParamsWithDefaults

`func NewPaymentParamsWithDefaults() *PaymentParams`

NewPaymentParamsWithDefaults instantiates a new PaymentParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *PaymentParams) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *PaymentParams) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *PaymentParams) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *PaymentParams) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTransactionStatus

`func (o *PaymentParams) GetTransactionStatus() string`

GetTransactionStatus returns the TransactionStatus field if non-nil, zero value otherwise.

### GetTransactionStatusOk

`func (o *PaymentParams) GetTransactionStatusOk() (*string, bool)`

GetTransactionStatusOk returns a tuple with the TransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatus

`func (o *PaymentParams) SetTransactionStatus(v string)`

SetTransactionStatus sets TransactionStatus field to given value.

### HasTransactionStatus

`func (o *PaymentParams) HasTransactionStatus() bool`

HasTransactionStatus returns a boolean if a field has been set.

### GetAmount

`func (o *PaymentParams) GetAmount() Value`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentParams) GetAmountOk() (*Value, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentParams) SetAmount(v Value)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaymentParams) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *PaymentParams) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentParams) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentParams) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


