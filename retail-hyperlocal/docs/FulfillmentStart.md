# FulfillmentStart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 
**Time** | Pointer to [**Time**](Time.md) |  | [optional] 
**Instructions** | Pointer to [**Descriptor**](Descriptor.md) |  | [optional] 
**Contact** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**Person** | Pointer to [**Person**](Person.md) |  | [optional] 
**Authorization** | Pointer to [**Authorization**](Authorization.md) |  | [optional] 

## Methods

### NewFulfillmentStart

`func NewFulfillmentStart() *FulfillmentStart`

NewFulfillmentStart instantiates a new FulfillmentStart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFulfillmentStartWithDefaults

`func NewFulfillmentStartWithDefaults() *FulfillmentStart`

NewFulfillmentStartWithDefaults instantiates a new FulfillmentStart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *FulfillmentStart) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *FulfillmentStart) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *FulfillmentStart) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *FulfillmentStart) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetTime

`func (o *FulfillmentStart) GetTime() Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *FulfillmentStart) GetTimeOk() (*Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *FulfillmentStart) SetTime(v Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *FulfillmentStart) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetInstructions

`func (o *FulfillmentStart) GetInstructions() Descriptor`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *FulfillmentStart) GetInstructionsOk() (*Descriptor, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *FulfillmentStart) SetInstructions(v Descriptor)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *FulfillmentStart) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetContact

`func (o *FulfillmentStart) GetContact() Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *FulfillmentStart) GetContactOk() (*Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *FulfillmentStart) SetContact(v Contact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *FulfillmentStart) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetPerson

`func (o *FulfillmentStart) GetPerson() Person`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *FulfillmentStart) GetPersonOk() (*Person, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *FulfillmentStart) SetPerson(v Person)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *FulfillmentStart) HasPerson() bool`

HasPerson returns a boolean if a field has been set.

### GetAuthorization

`func (o *FulfillmentStart) GetAuthorization() Authorization`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *FulfillmentStart) GetAuthorizationOk() (*Authorization, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *FulfillmentStart) SetAuthorization(v Authorization)`

SetAuthorization sets Authorization field to given value.

### HasAuthorization

`func (o *FulfillmentStart) HasAuthorization() bool`

HasAuthorization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


