# Schedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frequency** | Pointer to **string** | Describes duration as per ISO8601 format | [optional] 
**Holidays** | Pointer to [**[]time.Time**](time.Time.md) |  | [optional] 
**Times** | Pointer to [**[]time.Time**](time.Time.md) |  | [optional] 

## Methods

### NewSchedule

`func NewSchedule() *Schedule`

NewSchedule instantiates a new Schedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleWithDefaults

`func NewScheduleWithDefaults() *Schedule`

NewScheduleWithDefaults instantiates a new Schedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrequency

`func (o *Schedule) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *Schedule) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *Schedule) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *Schedule) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetHolidays

`func (o *Schedule) GetHolidays() []time.Time`

GetHolidays returns the Holidays field if non-nil, zero value otherwise.

### GetHolidaysOk

`func (o *Schedule) GetHolidaysOk() (*[]time.Time, bool)`

GetHolidaysOk returns a tuple with the Holidays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolidays

`func (o *Schedule) SetHolidays(v []time.Time)`

SetHolidays sets Holidays field to given value.

### HasHolidays

`func (o *Schedule) HasHolidays() bool`

HasHolidays returns a boolean if a field has been set.

### GetTimes

`func (o *Schedule) GetTimes() []time.Time`

GetTimes returns the Times field if non-nil, zero value otherwise.

### GetTimesOk

`func (o *Schedule) GetTimesOk() (*[]time.Time, bool)`

GetTimesOk returns a tuple with the Times field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimes

`func (o *Schedule) SetTimes(v []time.Time)`

SetTimes sets Times field to given value.

### HasTimes

`func (o *Schedule) HasTimes() bool`

HasTimes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


