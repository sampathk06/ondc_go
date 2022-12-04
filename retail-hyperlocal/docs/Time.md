# Time

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Duration** | Pointer to **string** | Describes duration as per ISO8601 format | [optional] 
**Range** | Pointer to [**TimeRange**](TimeRange.md) |  | [optional] 
**Days** | Pointer to **string** | comma separated values representing days of the week | [optional] 
**Schedule** | Pointer to [**Schedule**](Schedule.md) |  | [optional] 

## Methods

### NewTime

`func NewTime() *Time`

NewTime instantiates a new Time object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeWithDefaults

`func NewTimeWithDefaults() *Time`

NewTimeWithDefaults instantiates a new Time object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *Time) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Time) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Time) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Time) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetTimestamp

`func (o *Time) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Time) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Time) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Time) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetDuration

`func (o *Time) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Time) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Time) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Time) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetRange

`func (o *Time) GetRange() TimeRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *Time) GetRangeOk() (*TimeRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *Time) SetRange(v TimeRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *Time) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetDays

`func (o *Time) GetDays() string`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *Time) GetDaysOk() (*string, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *Time) SetDays(v string)`

SetDays sets Days field to given value.

### HasDays

`func (o *Time) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetSchedule

`func (o *Time) GetSchedule() Schedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *Time) GetScheduleOk() (*Schedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *Time) SetSchedule(v Schedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *Time) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


