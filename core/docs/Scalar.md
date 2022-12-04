# Scalar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Value** | **float32** |  | 
**EstimatedValue** | Pointer to **float32** |  | [optional] 
**ComputedValue** | Pointer to **float32** |  | [optional] 
**Range** | Pointer to [**ScalarRange**](ScalarRange.md) |  | [optional] 
**Unit** | **string** |  | 

## Methods

### NewScalar

`func NewScalar(value float32, unit string, ) *Scalar`

NewScalar instantiates a new Scalar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScalarWithDefaults

`func NewScalarWithDefaults() *Scalar`

NewScalarWithDefaults instantiates a new Scalar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Scalar) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Scalar) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Scalar) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Scalar) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *Scalar) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Scalar) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Scalar) SetValue(v float32)`

SetValue sets Value field to given value.


### GetEstimatedValue

`func (o *Scalar) GetEstimatedValue() float32`

GetEstimatedValue returns the EstimatedValue field if non-nil, zero value otherwise.

### GetEstimatedValueOk

`func (o *Scalar) GetEstimatedValueOk() (*float32, bool)`

GetEstimatedValueOk returns a tuple with the EstimatedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedValue

`func (o *Scalar) SetEstimatedValue(v float32)`

SetEstimatedValue sets EstimatedValue field to given value.

### HasEstimatedValue

`func (o *Scalar) HasEstimatedValue() bool`

HasEstimatedValue returns a boolean if a field has been set.

### GetComputedValue

`func (o *Scalar) GetComputedValue() float32`

GetComputedValue returns the ComputedValue field if non-nil, zero value otherwise.

### GetComputedValueOk

`func (o *Scalar) GetComputedValueOk() (*float32, bool)`

GetComputedValueOk returns a tuple with the ComputedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputedValue

`func (o *Scalar) SetComputedValue(v float32)`

SetComputedValue sets ComputedValue field to given value.

### HasComputedValue

`func (o *Scalar) HasComputedValue() bool`

HasComputedValue returns a boolean if a field has been set.

### GetRange

`func (o *Scalar) GetRange() ScalarRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *Scalar) GetRangeOk() (*ScalarRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *Scalar) SetRange(v ScalarRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *Scalar) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetUnit

`func (o *Scalar) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Scalar) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Scalar) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


