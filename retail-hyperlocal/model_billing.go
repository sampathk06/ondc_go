/*
ONDC Protocol API for retail (grocery, f&b)

ONDC Protocol API specification, which includes statutory requirements for packaged commodities and pre-packaged food This is an adaptation of Beckn Core version 0.9.3

API version: 1.0.27
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// Billing Describes a billing event
type Billing struct {
	// Personal details of the customer needed for billing.
	Name string `json:"name"`
	Organization *Organization `json:"organization,omitempty"`
	Address *Address `json:"address,omitempty"`
	Email *string `json:"email,omitempty"`
	Phone string `json:"phone"`
	Time *Time `json:"time,omitempty"`
	// GST number
	TaxNumber string `json:"tax_number"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewBilling instantiates a new Billing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBilling(name string, phone string, taxNumber string) *Billing {
	this := Billing{}
	this.Name = name
	this.Phone = phone
	this.TaxNumber = taxNumber
	return &this
}

// NewBillingWithDefaults instantiates a new Billing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingWithDefaults() *Billing {
	this := Billing{}
	return &this
}

// GetName returns the Name field value
func (o *Billing) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Billing) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Billing) SetName(v string) {
	o.Name = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *Billing) GetOrganization() Organization {
	if o == nil || isNil(o.Organization) {
		var ret Organization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Billing) GetOrganizationOk() (*Organization, bool) {
	if o == nil || isNil(o.Organization) {
    return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *Billing) HasOrganization() bool {
	if o != nil && !isNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Organization and assigns it to the Organization field.
func (o *Billing) SetOrganization(v Organization) {
	o.Organization = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Billing) GetAddress() Address {
	if o == nil || isNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Billing) GetAddressOk() (*Address, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Billing) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *Billing) SetAddress(v Address) {
	o.Address = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *Billing) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Billing) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
    return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *Billing) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *Billing) SetEmail(v string) {
	o.Email = &v
}

// GetPhone returns the Phone field value
func (o *Billing) GetPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
func (o *Billing) GetPhoneOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *Billing) SetPhone(v string) {
	o.Phone = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *Billing) GetTime() Time {
	if o == nil || isNil(o.Time) {
		var ret Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Billing) GetTimeOk() (*Time, bool) {
	if o == nil || isNil(o.Time) {
    return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *Billing) HasTime() bool {
	if o != nil && !isNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given Time and assigns it to the Time field.
func (o *Billing) SetTime(v Time) {
	o.Time = &v
}

// GetTaxNumber returns the TaxNumber field value
func (o *Billing) GetTaxNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaxNumber
}

// GetTaxNumberOk returns a tuple with the TaxNumber field value
// and a boolean to check if the value has been set.
func (o *Billing) GetTaxNumberOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TaxNumber, true
}

// SetTaxNumber sets field value
func (o *Billing) SetTaxNumber(v string) {
	o.TaxNumber = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Billing) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Billing) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Billing) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Billing) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Billing) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Billing) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Billing) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Billing) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o Billing) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["phone"] = o.Phone
	}
	if !isNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if true {
		toSerialize["tax_number"] = o.TaxNumber
	}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableBilling struct {
	value *Billing
	isSet bool
}

func (v NullableBilling) Get() *Billing {
	return v.value
}

func (v *NullableBilling) Set(val *Billing) {
	v.value = val
	v.isSet = true
}

func (v NullableBilling) IsSet() bool {
	return v.isSet
}

func (v *NullableBilling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBilling(val *Billing) *NullableBilling {
	return &NullableBilling{value: val, isSet: true}
}

func (v NullableBilling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBilling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


