/*
ONDC Protocol API for retail (grocery, f&b)

ONDC Protocol API specification, which includes statutory requirements for packaged commodities and pre-packaged food This is an adaptation of Beckn Core version 0.9.3

API version: 1.0.27
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// PaymentParams struct for PaymentParams
type PaymentParams struct {
	// This value will be placed in the the $transaction_id url param in case of http/get and in the requestBody http/post requests
	TransactionId *string `json:"transaction_id,omitempty"`
	TransactionStatus *string `json:"transaction_status,omitempty"`
	Amount *Value `json:"amount,omitempty"`
	Currency Currency `json:"currency"`
}

// NewPaymentParams instantiates a new PaymentParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentParams(currency Currency) *PaymentParams {
	this := PaymentParams{}
	this.Currency = currency
	return &this
}

// NewPaymentParamsWithDefaults instantiates a new PaymentParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentParamsWithDefaults() *PaymentParams {
	this := PaymentParams{}
	return &this
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *PaymentParams) GetTransactionId() string {
	if o == nil || isNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentParams) GetTransactionIdOk() (*string, bool) {
	if o == nil || isNil(o.TransactionId) {
    return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *PaymentParams) HasTransactionId() bool {
	if o != nil && !isNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *PaymentParams) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetTransactionStatus returns the TransactionStatus field value if set, zero value otherwise.
func (o *PaymentParams) GetTransactionStatus() string {
	if o == nil || isNil(o.TransactionStatus) {
		var ret string
		return ret
	}
	return *o.TransactionStatus
}

// GetTransactionStatusOk returns a tuple with the TransactionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentParams) GetTransactionStatusOk() (*string, bool) {
	if o == nil || isNil(o.TransactionStatus) {
    return nil, false
	}
	return o.TransactionStatus, true
}

// HasTransactionStatus returns a boolean if a field has been set.
func (o *PaymentParams) HasTransactionStatus() bool {
	if o != nil && !isNil(o.TransactionStatus) {
		return true
	}

	return false
}

// SetTransactionStatus gets a reference to the given string and assigns it to the TransactionStatus field.
func (o *PaymentParams) SetTransactionStatus(v string) {
	o.TransactionStatus = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PaymentParams) GetAmount() Value {
	if o == nil || isNil(o.Amount) {
		var ret Value
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentParams) GetAmountOk() (*Value, bool) {
	if o == nil || isNil(o.Amount) {
    return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PaymentParams) HasAmount() bool {
	if o != nil && !isNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given Value and assigns it to the Amount field.
func (o *PaymentParams) SetAmount(v Value) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value
func (o *PaymentParams) GetCurrency() Currency {
	if o == nil {
		var ret Currency
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *PaymentParams) GetCurrencyOk() (*Currency, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *PaymentParams) SetCurrency(v Currency) {
	o.Currency = v
}

func (o PaymentParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TransactionId) {
		toSerialize["transaction_id"] = o.TransactionId
	}
	if !isNil(o.TransactionStatus) {
		toSerialize["transaction_status"] = o.TransactionStatus
	}
	if !isNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentParams struct {
	value *PaymentParams
	isSet bool
}

func (v NullablePaymentParams) Get() *PaymentParams {
	return v.value
}

func (v *NullablePaymentParams) Set(val *PaymentParams) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentParams) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentParams(val *PaymentParams) *NullablePaymentParams {
	return &NullablePaymentParams{value: val, isSet: true}
}

func (v NullablePaymentParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


