/*
ONDC Protocol Core API

ONDC Protocol Core API specification. This is an adaptation of Beckn Core version 0.9.3

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Address Describes an address
type Address struct {
	// Door / Shop number of the address
	Door *string `json:"door,omitempty"`
	// Name of address if applicable. Example, shop name
	Name *string `json:"name,omitempty"`
	// Name of the building or block
	Building *string `json:"building,omitempty"`
	// Street name or number
	Street *string `json:"street,omitempty"`
	// Name of the locality, apartments
	Locality *string `json:"locality,omitempty"`
	// Name or number of the ward if applicable
	Ward *string `json:"ward,omitempty"`
	// City name
	City *string `json:"city,omitempty"`
	// State name
	State *string `json:"state,omitempty"`
	// Country name
	Country *string `json:"country,omitempty"`
	// Area code. This can be Pincode, ZIP code or any equivalent
	AreaCode *string `json:"area_code,omitempty"`
}

// NewAddress instantiates a new Address object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddress() *Address {
	this := Address{}
	return &this
}

// NewAddressWithDefaults instantiates a new Address object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressWithDefaults() *Address {
	this := Address{}
	return &this
}

// GetDoor returns the Door field value if set, zero value otherwise.
func (o *Address) GetDoor() string {
	if o == nil || isNil(o.Door) {
		var ret string
		return ret
	}
	return *o.Door
}

// GetDoorOk returns a tuple with the Door field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetDoorOk() (*string, bool) {
	if o == nil || isNil(o.Door) {
    return nil, false
	}
	return o.Door, true
}

// HasDoor returns a boolean if a field has been set.
func (o *Address) HasDoor() bool {
	if o != nil && !isNil(o.Door) {
		return true
	}

	return false
}

// SetDoor gets a reference to the given string and assigns it to the Door field.
func (o *Address) SetDoor(v string) {
	o.Door = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Address) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Address) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Address) SetName(v string) {
	o.Name = &v
}

// GetBuilding returns the Building field value if set, zero value otherwise.
func (o *Address) GetBuilding() string {
	if o == nil || isNil(o.Building) {
		var ret string
		return ret
	}
	return *o.Building
}

// GetBuildingOk returns a tuple with the Building field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetBuildingOk() (*string, bool) {
	if o == nil || isNil(o.Building) {
    return nil, false
	}
	return o.Building, true
}

// HasBuilding returns a boolean if a field has been set.
func (o *Address) HasBuilding() bool {
	if o != nil && !isNil(o.Building) {
		return true
	}

	return false
}

// SetBuilding gets a reference to the given string and assigns it to the Building field.
func (o *Address) SetBuilding(v string) {
	o.Building = &v
}

// GetStreet returns the Street field value if set, zero value otherwise.
func (o *Address) GetStreet() string {
	if o == nil || isNil(o.Street) {
		var ret string
		return ret
	}
	return *o.Street
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetStreetOk() (*string, bool) {
	if o == nil || isNil(o.Street) {
    return nil, false
	}
	return o.Street, true
}

// HasStreet returns a boolean if a field has been set.
func (o *Address) HasStreet() bool {
	if o != nil && !isNil(o.Street) {
		return true
	}

	return false
}

// SetStreet gets a reference to the given string and assigns it to the Street field.
func (o *Address) SetStreet(v string) {
	o.Street = &v
}

// GetLocality returns the Locality field value if set, zero value otherwise.
func (o *Address) GetLocality() string {
	if o == nil || isNil(o.Locality) {
		var ret string
		return ret
	}
	return *o.Locality
}

// GetLocalityOk returns a tuple with the Locality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetLocalityOk() (*string, bool) {
	if o == nil || isNil(o.Locality) {
    return nil, false
	}
	return o.Locality, true
}

// HasLocality returns a boolean if a field has been set.
func (o *Address) HasLocality() bool {
	if o != nil && !isNil(o.Locality) {
		return true
	}

	return false
}

// SetLocality gets a reference to the given string and assigns it to the Locality field.
func (o *Address) SetLocality(v string) {
	o.Locality = &v
}

// GetWard returns the Ward field value if set, zero value otherwise.
func (o *Address) GetWard() string {
	if o == nil || isNil(o.Ward) {
		var ret string
		return ret
	}
	return *o.Ward
}

// GetWardOk returns a tuple with the Ward field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetWardOk() (*string, bool) {
	if o == nil || isNil(o.Ward) {
    return nil, false
	}
	return o.Ward, true
}

// HasWard returns a boolean if a field has been set.
func (o *Address) HasWard() bool {
	if o != nil && !isNil(o.Ward) {
		return true
	}

	return false
}

// SetWard gets a reference to the given string and assigns it to the Ward field.
func (o *Address) SetWard(v string) {
	o.Ward = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *Address) GetCity() string {
	if o == nil || isNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetCityOk() (*string, bool) {
	if o == nil || isNil(o.City) {
    return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *Address) HasCity() bool {
	if o != nil && !isNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *Address) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Address) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
    return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Address) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Address) SetState(v string) {
	o.State = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *Address) GetCountry() string {
	if o == nil || isNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetCountryOk() (*string, bool) {
	if o == nil || isNil(o.Country) {
    return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *Address) HasCountry() bool {
	if o != nil && !isNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *Address) SetCountry(v string) {
	o.Country = &v
}

// GetAreaCode returns the AreaCode field value if set, zero value otherwise.
func (o *Address) GetAreaCode() string {
	if o == nil || isNil(o.AreaCode) {
		var ret string
		return ret
	}
	return *o.AreaCode
}

// GetAreaCodeOk returns a tuple with the AreaCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetAreaCodeOk() (*string, bool) {
	if o == nil || isNil(o.AreaCode) {
    return nil, false
	}
	return o.AreaCode, true
}

// HasAreaCode returns a boolean if a field has been set.
func (o *Address) HasAreaCode() bool {
	if o != nil && !isNil(o.AreaCode) {
		return true
	}

	return false
}

// SetAreaCode gets a reference to the given string and assigns it to the AreaCode field.
func (o *Address) SetAreaCode(v string) {
	o.AreaCode = &v
}

func (o Address) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Door) {
		toSerialize["door"] = o.Door
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Building) {
		toSerialize["building"] = o.Building
	}
	if !isNil(o.Street) {
		toSerialize["street"] = o.Street
	}
	if !isNil(o.Locality) {
		toSerialize["locality"] = o.Locality
	}
	if !isNil(o.Ward) {
		toSerialize["ward"] = o.Ward
	}
	if !isNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !isNil(o.AreaCode) {
		toSerialize["area_code"] = o.AreaCode
	}
	return json.Marshal(toSerialize)
}

type NullableAddress struct {
	value *Address
	isSet bool
}

func (v NullableAddress) Get() *Address {
	return v.value
}

func (v *NullableAddress) Set(val *Address) {
	v.value = val
	v.isSet = true
}

func (v NullableAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddress(val *Address) *NullableAddress {
	return &NullableAddress{value: val, isSet: true}
}

func (v NullableAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


