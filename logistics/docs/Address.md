# Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Door** | Pointer to **string** | Door / Shop number of the address | [optional] 
**Name** | Pointer to **string** | Name of address if applicable. Example, shop name | [optional] 
**Building** | Pointer to **string** | Name of the building or block | [optional] 
**Street** | Pointer to **string** | Street name or number | [optional] 
**Locality** | Pointer to **string** | Name of the locality, apartments | [optional] 
**Ward** | Pointer to **string** | Name or number of the ward if applicable | [optional] 
**City** | Pointer to **string** | City name | [optional] 
**State** | Pointer to **string** | State name | [optional] 
**Country** | Pointer to **string** | Country name | [optional] 
**AreaCode** | Pointer to **string** | Area code. This can be Pincode, ZIP code or any equivalent | [optional] 

## Methods

### NewAddress

`func NewAddress() *Address`

NewAddress instantiates a new Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWithDefaults

`func NewAddressWithDefaults() *Address`

NewAddressWithDefaults instantiates a new Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDoor

`func (o *Address) GetDoor() string`

GetDoor returns the Door field if non-nil, zero value otherwise.

### GetDoorOk

`func (o *Address) GetDoorOk() (*string, bool)`

GetDoorOk returns a tuple with the Door field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoor

`func (o *Address) SetDoor(v string)`

SetDoor sets Door field to given value.

### HasDoor

`func (o *Address) HasDoor() bool`

HasDoor returns a boolean if a field has been set.

### GetName

`func (o *Address) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Address) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Address) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Address) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBuilding

`func (o *Address) GetBuilding() string`

GetBuilding returns the Building field if non-nil, zero value otherwise.

### GetBuildingOk

`func (o *Address) GetBuildingOk() (*string, bool)`

GetBuildingOk returns a tuple with the Building field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilding

`func (o *Address) SetBuilding(v string)`

SetBuilding sets Building field to given value.

### HasBuilding

`func (o *Address) HasBuilding() bool`

HasBuilding returns a boolean if a field has been set.

### GetStreet

`func (o *Address) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *Address) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *Address) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *Address) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetLocality

`func (o *Address) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *Address) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *Address) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *Address) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetWard

`func (o *Address) GetWard() string`

GetWard returns the Ward field if non-nil, zero value otherwise.

### GetWardOk

`func (o *Address) GetWardOk() (*string, bool)`

GetWardOk returns a tuple with the Ward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWard

`func (o *Address) SetWard(v string)`

SetWard sets Ward field to given value.

### HasWard

`func (o *Address) HasWard() bool`

HasWard returns a boolean if a field has been set.

### GetCity

`func (o *Address) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Address) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Address) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Address) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *Address) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Address) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Address) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Address) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *Address) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Address) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Address) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Address) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetAreaCode

`func (o *Address) GetAreaCode() string`

GetAreaCode returns the AreaCode field if non-nil, zero value otherwise.

### GetAreaCodeOk

`func (o *Address) GetAreaCodeOk() (*string, bool)`

GetAreaCodeOk returns a tuple with the AreaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaCode

`func (o *Address) SetAreaCode(v string)`

SetAreaCode sets AreaCode field to given value.

### HasAreaCode

`func (o *Address) HasAreaCode() bool`

HasAreaCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


