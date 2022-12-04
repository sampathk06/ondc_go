/*
ONDC Protocol API for logistics

ONDC Protocol Core API specification. This is an adaptation of Beckn Core version 0.9.3

API version: 1.0.19
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Payment Describes a payment
type Payment struct {
	// A payment uri to be called by the Buyer App. If empty, then the payment is to be done offline. The details of payment should be present in the params object. If ```tl_method``` = http/get, then the payment details will be sent as url params. Two url param values, ```$transaction_id``` and ```$amount``` are mandatory. And example url would be : https://www.example.com/pay?txid=$transaction_id&amount=$amount&vpa=upiid&payee=shopez&billno=1234
	Uri *string `json:"uri,omitempty"`
	TlMethod *string `json:"tl_method,omitempty"`
	Params *PaymentParams `json:"params,omitempty"`
	Type *string `json:"type,omitempty"`
	Status *string `json:"status,omitempty"`
	Time *Time `json:"time,omitempty"`
	CollectedBy *string `json:"collected_by,omitempty"`
	// CoD collection amount
	OndcOrgCollectionAmount *string `json:"@ondc/org/collection_amount,omitempty"`
	// settlement window in ISO8601 durations format e.g. 'PT48H' indicates T+2 settlement
	OndcOrgSettlementWindow *string `json:"@ondc/org/settlement_window,omitempty"`
	OndcOrgSettlementWindowStatus *string `json:"@ondc/org/settlement_window_status,omitempty"`
	OndcOrgSettlementDetails []PaymentOndcOrgSettlementDetailsInner `json:"@ondc/org/settlement_details,omitempty"`
}

// NewPayment instantiates a new Payment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayment() *Payment {
	this := Payment{}
	return &this
}

// NewPaymentWithDefaults instantiates a new Payment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentWithDefaults() *Payment {
	this := Payment{}
	return &this
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *Payment) GetUri() string {
	if o == nil || isNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetUriOk() (*string, bool) {
	if o == nil || isNil(o.Uri) {
    return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *Payment) HasUri() bool {
	if o != nil && !isNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *Payment) SetUri(v string) {
	o.Uri = &v
}

// GetTlMethod returns the TlMethod field value if set, zero value otherwise.
func (o *Payment) GetTlMethod() string {
	if o == nil || isNil(o.TlMethod) {
		var ret string
		return ret
	}
	return *o.TlMethod
}

// GetTlMethodOk returns a tuple with the TlMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetTlMethodOk() (*string, bool) {
	if o == nil || isNil(o.TlMethod) {
    return nil, false
	}
	return o.TlMethod, true
}

// HasTlMethod returns a boolean if a field has been set.
func (o *Payment) HasTlMethod() bool {
	if o != nil && !isNil(o.TlMethod) {
		return true
	}

	return false
}

// SetTlMethod gets a reference to the given string and assigns it to the TlMethod field.
func (o *Payment) SetTlMethod(v string) {
	o.TlMethod = &v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *Payment) GetParams() PaymentParams {
	if o == nil || isNil(o.Params) {
		var ret PaymentParams
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetParamsOk() (*PaymentParams, bool) {
	if o == nil || isNil(o.Params) {
    return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *Payment) HasParams() bool {
	if o != nil && !isNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given PaymentParams and assigns it to the Params field.
func (o *Payment) SetParams(v PaymentParams) {
	o.Params = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Payment) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Payment) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Payment) SetType(v string) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Payment) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Payment) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Payment) SetStatus(v string) {
	o.Status = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *Payment) GetTime() Time {
	if o == nil || isNil(o.Time) {
		var ret Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetTimeOk() (*Time, bool) {
	if o == nil || isNil(o.Time) {
    return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *Payment) HasTime() bool {
	if o != nil && !isNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given Time and assigns it to the Time field.
func (o *Payment) SetTime(v Time) {
	o.Time = &v
}

// GetCollectedBy returns the CollectedBy field value if set, zero value otherwise.
func (o *Payment) GetCollectedBy() string {
	if o == nil || isNil(o.CollectedBy) {
		var ret string
		return ret
	}
	return *o.CollectedBy
}

// GetCollectedByOk returns a tuple with the CollectedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetCollectedByOk() (*string, bool) {
	if o == nil || isNil(o.CollectedBy) {
    return nil, false
	}
	return o.CollectedBy, true
}

// HasCollectedBy returns a boolean if a field has been set.
func (o *Payment) HasCollectedBy() bool {
	if o != nil && !isNil(o.CollectedBy) {
		return true
	}

	return false
}

// SetCollectedBy gets a reference to the given string and assigns it to the CollectedBy field.
func (o *Payment) SetCollectedBy(v string) {
	o.CollectedBy = &v
}

// GetOndcOrgCollectionAmount returns the OndcOrgCollectionAmount field value if set, zero value otherwise.
func (o *Payment) GetOndcOrgCollectionAmount() string {
	if o == nil || isNil(o.OndcOrgCollectionAmount) {
		var ret string
		return ret
	}
	return *o.OndcOrgCollectionAmount
}

// GetOndcOrgCollectionAmountOk returns a tuple with the OndcOrgCollectionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetOndcOrgCollectionAmountOk() (*string, bool) {
	if o == nil || isNil(o.OndcOrgCollectionAmount) {
    return nil, false
	}
	return o.OndcOrgCollectionAmount, true
}

// HasOndcOrgCollectionAmount returns a boolean if a field has been set.
func (o *Payment) HasOndcOrgCollectionAmount() bool {
	if o != nil && !isNil(o.OndcOrgCollectionAmount) {
		return true
	}

	return false
}

// SetOndcOrgCollectionAmount gets a reference to the given string and assigns it to the OndcOrgCollectionAmount field.
func (o *Payment) SetOndcOrgCollectionAmount(v string) {
	o.OndcOrgCollectionAmount = &v
}

// GetOndcOrgSettlementWindow returns the OndcOrgSettlementWindow field value if set, zero value otherwise.
func (o *Payment) GetOndcOrgSettlementWindow() string {
	if o == nil || isNil(o.OndcOrgSettlementWindow) {
		var ret string
		return ret
	}
	return *o.OndcOrgSettlementWindow
}

// GetOndcOrgSettlementWindowOk returns a tuple with the OndcOrgSettlementWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetOndcOrgSettlementWindowOk() (*string, bool) {
	if o == nil || isNil(o.OndcOrgSettlementWindow) {
    return nil, false
	}
	return o.OndcOrgSettlementWindow, true
}

// HasOndcOrgSettlementWindow returns a boolean if a field has been set.
func (o *Payment) HasOndcOrgSettlementWindow() bool {
	if o != nil && !isNil(o.OndcOrgSettlementWindow) {
		return true
	}

	return false
}

// SetOndcOrgSettlementWindow gets a reference to the given string and assigns it to the OndcOrgSettlementWindow field.
func (o *Payment) SetOndcOrgSettlementWindow(v string) {
	o.OndcOrgSettlementWindow = &v
}

// GetOndcOrgSettlementWindowStatus returns the OndcOrgSettlementWindowStatus field value if set, zero value otherwise.
func (o *Payment) GetOndcOrgSettlementWindowStatus() string {
	if o == nil || isNil(o.OndcOrgSettlementWindowStatus) {
		var ret string
		return ret
	}
	return *o.OndcOrgSettlementWindowStatus
}

// GetOndcOrgSettlementWindowStatusOk returns a tuple with the OndcOrgSettlementWindowStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetOndcOrgSettlementWindowStatusOk() (*string, bool) {
	if o == nil || isNil(o.OndcOrgSettlementWindowStatus) {
    return nil, false
	}
	return o.OndcOrgSettlementWindowStatus, true
}

// HasOndcOrgSettlementWindowStatus returns a boolean if a field has been set.
func (o *Payment) HasOndcOrgSettlementWindowStatus() bool {
	if o != nil && !isNil(o.OndcOrgSettlementWindowStatus) {
		return true
	}

	return false
}

// SetOndcOrgSettlementWindowStatus gets a reference to the given string and assigns it to the OndcOrgSettlementWindowStatus field.
func (o *Payment) SetOndcOrgSettlementWindowStatus(v string) {
	o.OndcOrgSettlementWindowStatus = &v
}

// GetOndcOrgSettlementDetails returns the OndcOrgSettlementDetails field value if set, zero value otherwise.
func (o *Payment) GetOndcOrgSettlementDetails() []PaymentOndcOrgSettlementDetailsInner {
	if o == nil || isNil(o.OndcOrgSettlementDetails) {
		var ret []PaymentOndcOrgSettlementDetailsInner
		return ret
	}
	return o.OndcOrgSettlementDetails
}

// GetOndcOrgSettlementDetailsOk returns a tuple with the OndcOrgSettlementDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetOndcOrgSettlementDetailsOk() ([]PaymentOndcOrgSettlementDetailsInner, bool) {
	if o == nil || isNil(o.OndcOrgSettlementDetails) {
    return nil, false
	}
	return o.OndcOrgSettlementDetails, true
}

// HasOndcOrgSettlementDetails returns a boolean if a field has been set.
func (o *Payment) HasOndcOrgSettlementDetails() bool {
	if o != nil && !isNil(o.OndcOrgSettlementDetails) {
		return true
	}

	return false
}

// SetOndcOrgSettlementDetails gets a reference to the given []PaymentOndcOrgSettlementDetailsInner and assigns it to the OndcOrgSettlementDetails field.
func (o *Payment) SetOndcOrgSettlementDetails(v []PaymentOndcOrgSettlementDetailsInner) {
	o.OndcOrgSettlementDetails = v
}

func (o Payment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	if !isNil(o.TlMethod) {
		toSerialize["tl_method"] = o.TlMethod
	}
	if !isNil(o.Params) {
		toSerialize["params"] = o.Params
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !isNil(o.CollectedBy) {
		toSerialize["collected_by"] = o.CollectedBy
	}
	if !isNil(o.OndcOrgCollectionAmount) {
		toSerialize["@ondc/org/collection_amount"] = o.OndcOrgCollectionAmount
	}
	if !isNil(o.OndcOrgSettlementWindow) {
		toSerialize["@ondc/org/settlement_window"] = o.OndcOrgSettlementWindow
	}
	if !isNil(o.OndcOrgSettlementWindowStatus) {
		toSerialize["@ondc/org/settlement_window_status"] = o.OndcOrgSettlementWindowStatus
	}
	if !isNil(o.OndcOrgSettlementDetails) {
		toSerialize["@ondc/org/settlement_details"] = o.OndcOrgSettlementDetails
	}
	return json.Marshal(toSerialize)
}

type NullablePayment struct {
	value *Payment
	isSet bool
}

func (v NullablePayment) Get() *Payment {
	return v.value
}

func (v *NullablePayment) Set(val *Payment) {
	v.value = val
	v.isSet = true
}

func (v NullablePayment) IsSet() bool {
	return v.isSet
}

func (v *NullablePayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayment(val *Payment) *NullablePayment {
	return &NullablePayment{value: val, isSet: true}
}

func (v NullablePayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


