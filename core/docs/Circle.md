# Circle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gps** | **string** | Describes a gps coordinate | 
**Radius** | [**Scalar**](Scalar.md) |  | 

## Methods

### NewCircle

`func NewCircle(gps string, radius Scalar, ) *Circle`

NewCircle instantiates a new Circle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCircleWithDefaults

`func NewCircleWithDefaults() *Circle`

NewCircleWithDefaults instantiates a new Circle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGps

`func (o *Circle) GetGps() string`

GetGps returns the Gps field if non-nil, zero value otherwise.

### GetGpsOk

`func (o *Circle) GetGpsOk() (*string, bool)`

GetGpsOk returns a tuple with the Gps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGps

`func (o *Circle) SetGps(v string)`

SetGps sets Gps field to given value.


### GetRadius

`func (o *Circle) GetRadius() Scalar`

GetRadius returns the Radius field if non-nil, zero value otherwise.

### GetRadiusOk

`func (o *Circle) GetRadiusOk() (*Scalar, bool)`

GetRadiusOk returns a tuple with the Radius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadius

`func (o *Circle) SetRadius(v Scalar)`

SetRadius sets Radius field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


