/*
ONDC Protocol Core API

ONDC Protocol Core API specification. This is an adaptation of Beckn Core version 0.9.3

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// Time Describes time in its various forms. It can be a single point in time; duration; or a structured timetable of operations
type Time struct {
	Label *string `json:"label,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// Describes duration as per ISO8601 format
	Duration *string `json:"duration,omitempty"`
	Range *TimeRange `json:"range,omitempty"`
	// comma separated values representing days of the week
	Days *string `json:"days,omitempty"`
	Schedule *Schedule `json:"schedule,omitempty"`
}

// NewTime instantiates a new Time object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTime() *Time {
	this := Time{}
	return &this
}

// NewTimeWithDefaults instantiates a new Time object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeWithDefaults() *Time {
	this := Time{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *Time) GetLabel() string {
	if o == nil || isNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Time) GetLabelOk() (*string, bool) {
	if o == nil || isNil(o.Label) {
    return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *Time) HasLabel() bool {
	if o != nil && !isNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *Time) SetLabel(v string) {
	o.Label = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *Time) GetTimestamp() time.Time {
	if o == nil || isNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Time) GetTimestampOk() (*time.Time, bool) {
	if o == nil || isNil(o.Timestamp) {
    return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *Time) HasTimestamp() bool {
	if o != nil && !isNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *Time) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *Time) GetDuration() string {
	if o == nil || isNil(o.Duration) {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Time) GetDurationOk() (*string, bool) {
	if o == nil || isNil(o.Duration) {
    return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *Time) HasDuration() bool {
	if o != nil && !isNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *Time) SetDuration(v string) {
	o.Duration = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *Time) GetRange() TimeRange {
	if o == nil || isNil(o.Range) {
		var ret TimeRange
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Time) GetRangeOk() (*TimeRange, bool) {
	if o == nil || isNil(o.Range) {
    return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *Time) HasRange() bool {
	if o != nil && !isNil(o.Range) {
		return true
	}

	return false
}

// SetRange gets a reference to the given TimeRange and assigns it to the Range field.
func (o *Time) SetRange(v TimeRange) {
	o.Range = &v
}

// GetDays returns the Days field value if set, zero value otherwise.
func (o *Time) GetDays() string {
	if o == nil || isNil(o.Days) {
		var ret string
		return ret
	}
	return *o.Days
}

// GetDaysOk returns a tuple with the Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Time) GetDaysOk() (*string, bool) {
	if o == nil || isNil(o.Days) {
    return nil, false
	}
	return o.Days, true
}

// HasDays returns a boolean if a field has been set.
func (o *Time) HasDays() bool {
	if o != nil && !isNil(o.Days) {
		return true
	}

	return false
}

// SetDays gets a reference to the given string and assigns it to the Days field.
func (o *Time) SetDays(v string) {
	o.Days = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *Time) GetSchedule() Schedule {
	if o == nil || isNil(o.Schedule) {
		var ret Schedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Time) GetScheduleOk() (*Schedule, bool) {
	if o == nil || isNil(o.Schedule) {
    return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *Time) HasSchedule() bool {
	if o != nil && !isNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given Schedule and assigns it to the Schedule field.
func (o *Time) SetSchedule(v Schedule) {
	o.Schedule = &v
}

func (o Time) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !isNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !isNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !isNil(o.Range) {
		toSerialize["range"] = o.Range
	}
	if !isNil(o.Days) {
		toSerialize["days"] = o.Days
	}
	if !isNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	return json.Marshal(toSerialize)
}

type NullableTime struct {
	value *Time
	isSet bool
}

func (v NullableTime) Get() *Time {
	return v.value
}

func (v *NullableTime) Set(val *Time) {
	v.value = val
	v.isSet = true
}

func (v NullableTime) IsSet() bool {
	return v.isSet
}

func (v *NullableTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTime(val *Time) *NullableTime {
	return &NullableTime{value: val, isSet: true}
}

func (v NullableTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

