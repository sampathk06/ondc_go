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

// CancelPostRequestMessage struct for CancelPostRequestMessage
type CancelPostRequestMessage struct {
	OrderId Id `json:"order_id"`
	CancellationReasonId *Id `json:"cancellation_reason_id,omitempty"`
	Descriptor *Descriptor `json:"descriptor,omitempty"`
}

// NewCancelPostRequestMessage instantiates a new CancelPostRequestMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelPostRequestMessage(orderId Id) *CancelPostRequestMessage {
	this := CancelPostRequestMessage{}
	this.OrderId = orderId
	return &this
}

// NewCancelPostRequestMessageWithDefaults instantiates a new CancelPostRequestMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelPostRequestMessageWithDefaults() *CancelPostRequestMessage {
	this := CancelPostRequestMessage{}
	return &this
}

// GetOrderId returns the OrderId field value
func (o *CancelPostRequestMessage) GetOrderId() Id {
	if o == nil {
		var ret Id
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *CancelPostRequestMessage) GetOrderIdOk() (*Id, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *CancelPostRequestMessage) SetOrderId(v Id) {
	o.OrderId = v
}

// GetCancellationReasonId returns the CancellationReasonId field value if set, zero value otherwise.
func (o *CancelPostRequestMessage) GetCancellationReasonId() Id {
	if o == nil || isNil(o.CancellationReasonId) {
		var ret Id
		return ret
	}
	return *o.CancellationReasonId
}

// GetCancellationReasonIdOk returns a tuple with the CancellationReasonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelPostRequestMessage) GetCancellationReasonIdOk() (*Id, bool) {
	if o == nil || isNil(o.CancellationReasonId) {
    return nil, false
	}
	return o.CancellationReasonId, true
}

// HasCancellationReasonId returns a boolean if a field has been set.
func (o *CancelPostRequestMessage) HasCancellationReasonId() bool {
	if o != nil && !isNil(o.CancellationReasonId) {
		return true
	}

	return false
}

// SetCancellationReasonId gets a reference to the given Id and assigns it to the CancellationReasonId field.
func (o *CancelPostRequestMessage) SetCancellationReasonId(v Id) {
	o.CancellationReasonId = &v
}

// GetDescriptor returns the Descriptor field value if set, zero value otherwise.
func (o *CancelPostRequestMessage) GetDescriptor() Descriptor {
	if o == nil || isNil(o.Descriptor) {
		var ret Descriptor
		return ret
	}
	return *o.Descriptor
}

// GetDescriptorOk returns a tuple with the Descriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelPostRequestMessage) GetDescriptorOk() (*Descriptor, bool) {
	if o == nil || isNil(o.Descriptor) {
    return nil, false
	}
	return o.Descriptor, true
}

// HasDescriptor returns a boolean if a field has been set.
func (o *CancelPostRequestMessage) HasDescriptor() bool {
	if o != nil && !isNil(o.Descriptor) {
		return true
	}

	return false
}

// SetDescriptor gets a reference to the given Descriptor and assigns it to the Descriptor field.
func (o *CancelPostRequestMessage) SetDescriptor(v Descriptor) {
	o.Descriptor = &v
}

func (o CancelPostRequestMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["order_id"] = o.OrderId
	}
	if !isNil(o.CancellationReasonId) {
		toSerialize["cancellation_reason_id"] = o.CancellationReasonId
	}
	if !isNil(o.Descriptor) {
		toSerialize["descriptor"] = o.Descriptor
	}
	return json.Marshal(toSerialize)
}

type NullableCancelPostRequestMessage struct {
	value *CancelPostRequestMessage
	isSet bool
}

func (v NullableCancelPostRequestMessage) Get() *CancelPostRequestMessage {
	return v.value
}

func (v *NullableCancelPostRequestMessage) Set(val *CancelPostRequestMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelPostRequestMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelPostRequestMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelPostRequestMessage(val *CancelPostRequestMessage) *NullableCancelPostRequestMessage {
	return &NullableCancelPostRequestMessage{value: val, isSet: true}
}

func (v NullableCancelPostRequestMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelPostRequestMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

