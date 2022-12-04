# OrderProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**Id**](Id.md) |  | [optional] 
**Locations** | Pointer to [**[]OrderProviderLocationsInner**](OrderProviderLocationsInner.md) |  | [optional] 

## Methods

### NewOrderProvider

`func NewOrderProvider() *OrderProvider`

NewOrderProvider instantiates a new OrderProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderProviderWithDefaults

`func NewOrderProviderWithDefaults() *OrderProvider`

NewOrderProviderWithDefaults instantiates a new OrderProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderProvider) GetId() Id`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderProvider) GetIdOk() (*Id, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderProvider) SetId(v Id)`

SetId sets Id field to given value.

### HasId

`func (o *OrderProvider) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocations

`func (o *OrderProvider) GetLocations() []OrderProviderLocationsInner`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *OrderProvider) GetLocationsOk() (*[]OrderProviderLocationsInner, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *OrderProvider) SetLocations(v []OrderProviderLocationsInner)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *OrderProvider) HasLocations() bool`

HasLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


