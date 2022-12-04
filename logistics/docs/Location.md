# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Descriptor** | Pointer to [**Descriptor**](Descriptor.md) |  | [optional] 
**Gps** | Pointer to **string** | Describes a gps coordinate | [optional] 
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**StationCode** | Pointer to **string** |  | [optional] 
**City** | Pointer to [**City**](City.md) |  | [optional] 
**Country** | Pointer to [**Country**](Country.md) |  | [optional] 
**Circle** | Pointer to [**Circle**](Circle.md) |  | [optional] 
**Polygon** | Pointer to **string** |  | [optional] 
**Var3dspace** | Pointer to **string** |  | [optional] 
**Time** | Pointer to [**Time**](Time.md) |  | [optional] 

## Methods

### NewLocation

`func NewLocation() *Location`

NewLocation instantiates a new Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWithDefaults

`func NewLocationWithDefaults() *Location`

NewLocationWithDefaults instantiates a new Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Location) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Location) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Location) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Location) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescriptor

`func (o *Location) GetDescriptor() Descriptor`

GetDescriptor returns the Descriptor field if non-nil, zero value otherwise.

### GetDescriptorOk

`func (o *Location) GetDescriptorOk() (*Descriptor, bool)`

GetDescriptorOk returns a tuple with the Descriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptor

`func (o *Location) SetDescriptor(v Descriptor)`

SetDescriptor sets Descriptor field to given value.

### HasDescriptor

`func (o *Location) HasDescriptor() bool`

HasDescriptor returns a boolean if a field has been set.

### GetGps

`func (o *Location) GetGps() string`

GetGps returns the Gps field if non-nil, zero value otherwise.

### GetGpsOk

`func (o *Location) GetGpsOk() (*string, bool)`

GetGpsOk returns a tuple with the Gps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGps

`func (o *Location) SetGps(v string)`

SetGps sets Gps field to given value.

### HasGps

`func (o *Location) HasGps() bool`

HasGps returns a boolean if a field has been set.

### GetAddress

`func (o *Location) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Location) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Location) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Location) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetStationCode

`func (o *Location) GetStationCode() string`

GetStationCode returns the StationCode field if non-nil, zero value otherwise.

### GetStationCodeOk

`func (o *Location) GetStationCodeOk() (*string, bool)`

GetStationCodeOk returns a tuple with the StationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationCode

`func (o *Location) SetStationCode(v string)`

SetStationCode sets StationCode field to given value.

### HasStationCode

`func (o *Location) HasStationCode() bool`

HasStationCode returns a boolean if a field has been set.

### GetCity

`func (o *Location) GetCity() City`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Location) GetCityOk() (*City, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Location) SetCity(v City)`

SetCity sets City field to given value.

### HasCity

`func (o *Location) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *Location) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Location) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Location) SetCountry(v Country)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Location) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCircle

`func (o *Location) GetCircle() Circle`

GetCircle returns the Circle field if non-nil, zero value otherwise.

### GetCircleOk

`func (o *Location) GetCircleOk() (*Circle, bool)`

GetCircleOk returns a tuple with the Circle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircle

`func (o *Location) SetCircle(v Circle)`

SetCircle sets Circle field to given value.

### HasCircle

`func (o *Location) HasCircle() bool`

HasCircle returns a boolean if a field has been set.

### GetPolygon

`func (o *Location) GetPolygon() string`

GetPolygon returns the Polygon field if non-nil, zero value otherwise.

### GetPolygonOk

`func (o *Location) GetPolygonOk() (*string, bool)`

GetPolygonOk returns a tuple with the Polygon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolygon

`func (o *Location) SetPolygon(v string)`

SetPolygon sets Polygon field to given value.

### HasPolygon

`func (o *Location) HasPolygon() bool`

HasPolygon returns a boolean if a field has been set.

### GetVar3dspace

`func (o *Location) GetVar3dspace() string`

GetVar3dspace returns the Var3dspace field if non-nil, zero value otherwise.

### GetVar3dspaceOk

`func (o *Location) GetVar3dspaceOk() (*string, bool)`

GetVar3dspaceOk returns a tuple with the Var3dspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar3dspace

`func (o *Location) SetVar3dspace(v string)`

SetVar3dspace sets Var3dspace field to given value.

### HasVar3dspace

`func (o *Location) HasVar3dspace() bool`

HasVar3dspace returns a boolean if a field has been set.

### GetTime

`func (o *Location) GetTime() Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Location) GetTimeOk() (*Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Location) SetTime(v Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *Location) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


