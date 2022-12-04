# ProviderLocationsInner

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
**Rateable** | Pointer to **bool** | If the entity can be rated or not | [optional] 

## Methods

### NewProviderLocationsInner

`func NewProviderLocationsInner() *ProviderLocationsInner`

NewProviderLocationsInner instantiates a new ProviderLocationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderLocationsInnerWithDefaults

`func NewProviderLocationsInnerWithDefaults() *ProviderLocationsInner`

NewProviderLocationsInnerWithDefaults instantiates a new ProviderLocationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProviderLocationsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProviderLocationsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProviderLocationsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProviderLocationsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescriptor

`func (o *ProviderLocationsInner) GetDescriptor() Descriptor`

GetDescriptor returns the Descriptor field if non-nil, zero value otherwise.

### GetDescriptorOk

`func (o *ProviderLocationsInner) GetDescriptorOk() (*Descriptor, bool)`

GetDescriptorOk returns a tuple with the Descriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptor

`func (o *ProviderLocationsInner) SetDescriptor(v Descriptor)`

SetDescriptor sets Descriptor field to given value.

### HasDescriptor

`func (o *ProviderLocationsInner) HasDescriptor() bool`

HasDescriptor returns a boolean if a field has been set.

### GetGps

`func (o *ProviderLocationsInner) GetGps() string`

GetGps returns the Gps field if non-nil, zero value otherwise.

### GetGpsOk

`func (o *ProviderLocationsInner) GetGpsOk() (*string, bool)`

GetGpsOk returns a tuple with the Gps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGps

`func (o *ProviderLocationsInner) SetGps(v string)`

SetGps sets Gps field to given value.

### HasGps

`func (o *ProviderLocationsInner) HasGps() bool`

HasGps returns a boolean if a field has been set.

### GetAddress

`func (o *ProviderLocationsInner) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ProviderLocationsInner) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ProviderLocationsInner) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ProviderLocationsInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetStationCode

`func (o *ProviderLocationsInner) GetStationCode() string`

GetStationCode returns the StationCode field if non-nil, zero value otherwise.

### GetStationCodeOk

`func (o *ProviderLocationsInner) GetStationCodeOk() (*string, bool)`

GetStationCodeOk returns a tuple with the StationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationCode

`func (o *ProviderLocationsInner) SetStationCode(v string)`

SetStationCode sets StationCode field to given value.

### HasStationCode

`func (o *ProviderLocationsInner) HasStationCode() bool`

HasStationCode returns a boolean if a field has been set.

### GetCity

`func (o *ProviderLocationsInner) GetCity() City`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ProviderLocationsInner) GetCityOk() (*City, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ProviderLocationsInner) SetCity(v City)`

SetCity sets City field to given value.

### HasCity

`func (o *ProviderLocationsInner) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *ProviderLocationsInner) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ProviderLocationsInner) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ProviderLocationsInner) SetCountry(v Country)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ProviderLocationsInner) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCircle

`func (o *ProviderLocationsInner) GetCircle() Circle`

GetCircle returns the Circle field if non-nil, zero value otherwise.

### GetCircleOk

`func (o *ProviderLocationsInner) GetCircleOk() (*Circle, bool)`

GetCircleOk returns a tuple with the Circle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircle

`func (o *ProviderLocationsInner) SetCircle(v Circle)`

SetCircle sets Circle field to given value.

### HasCircle

`func (o *ProviderLocationsInner) HasCircle() bool`

HasCircle returns a boolean if a field has been set.

### GetPolygon

`func (o *ProviderLocationsInner) GetPolygon() string`

GetPolygon returns the Polygon field if non-nil, zero value otherwise.

### GetPolygonOk

`func (o *ProviderLocationsInner) GetPolygonOk() (*string, bool)`

GetPolygonOk returns a tuple with the Polygon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolygon

`func (o *ProviderLocationsInner) SetPolygon(v string)`

SetPolygon sets Polygon field to given value.

### HasPolygon

`func (o *ProviderLocationsInner) HasPolygon() bool`

HasPolygon returns a boolean if a field has been set.

### GetVar3dspace

`func (o *ProviderLocationsInner) GetVar3dspace() string`

GetVar3dspace returns the Var3dspace field if non-nil, zero value otherwise.

### GetVar3dspaceOk

`func (o *ProviderLocationsInner) GetVar3dspaceOk() (*string, bool)`

GetVar3dspaceOk returns a tuple with the Var3dspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar3dspace

`func (o *ProviderLocationsInner) SetVar3dspace(v string)`

SetVar3dspace sets Var3dspace field to given value.

### HasVar3dspace

`func (o *ProviderLocationsInner) HasVar3dspace() bool`

HasVar3dspace returns a boolean if a field has been set.

### GetTime

`func (o *ProviderLocationsInner) GetTime() Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ProviderLocationsInner) GetTimeOk() (*Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ProviderLocationsInner) SetTime(v Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ProviderLocationsInner) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetRateable

`func (o *ProviderLocationsInner) GetRateable() bool`

GetRateable returns the Rateable field if non-nil, zero value otherwise.

### GetRateableOk

`func (o *ProviderLocationsInner) GetRateableOk() (*bool, bool)`

GetRateableOk returns a tuple with the Rateable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateable

`func (o *ProviderLocationsInner) SetRateable(v bool)`

SetRateable sets Rateable field to given value.

### HasRateable

`func (o *ProviderLocationsInner) HasRateable() bool`

HasRateable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


