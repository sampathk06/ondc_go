# Fulfillment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique reference ID to the fulfillment of an order | [optional] 
**Type** | Pointer to **string** | This describes the type of fulfillment | [optional] 
**OndcOrgAwbNo** | Pointer to **string** |  | [optional] 
**OndcOrgEwaybillno** | Pointer to **string** |  | [optional] 
**OndcOrgEbnexpirydate** | Pointer to **time.Time** |  | [optional] 
**ProviderId** | Pointer to [**Id**](Id.md) |  | [optional] 
**Rating** | Pointer to [**Value**](Value.md) |  | [optional] 
**State** | Pointer to [**State**](State.md) |  | [optional] 
**Tracking** | Pointer to **bool** | Indicates whether the fulfillment allows tracking | [optional] [default to false]
**Customer** | Pointer to [**FulfillmentCustomer**](FulfillmentCustomer.md) |  | [optional] 
**Agent** | Pointer to [**Agent**](Agent.md) |  | [optional] 
**Person** | Pointer to [**Person**](Person.md) |  | [optional] 
**Contact** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**Vehicle** | Pointer to [**Vehicle**](Vehicle.md) |  | [optional] 
**Start** | Pointer to [**FulfillmentStart**](FulfillmentStart.md) |  | [optional] 
**End** | Pointer to [**FulfillmentEnd**](FulfillmentEnd.md) |  | [optional] 
**Rateable** | Pointer to **bool** | If the entity can be rated or not | [optional] 
**Tags** | Pointer to **map[string]string** | Describes a tag. This is a simple key-value store which is used to contain extended metadata | [optional] 

## Methods

### NewFulfillment

`func NewFulfillment() *Fulfillment`

NewFulfillment instantiates a new Fulfillment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFulfillmentWithDefaults

`func NewFulfillmentWithDefaults() *Fulfillment`

NewFulfillmentWithDefaults instantiates a new Fulfillment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Fulfillment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Fulfillment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Fulfillment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Fulfillment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Fulfillment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Fulfillment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Fulfillment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Fulfillment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOndcOrgAwbNo

`func (o *Fulfillment) GetOndcOrgAwbNo() string`

GetOndcOrgAwbNo returns the OndcOrgAwbNo field if non-nil, zero value otherwise.

### GetOndcOrgAwbNoOk

`func (o *Fulfillment) GetOndcOrgAwbNoOk() (*string, bool)`

GetOndcOrgAwbNoOk returns a tuple with the OndcOrgAwbNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgAwbNo

`func (o *Fulfillment) SetOndcOrgAwbNo(v string)`

SetOndcOrgAwbNo sets OndcOrgAwbNo field to given value.

### HasOndcOrgAwbNo

`func (o *Fulfillment) HasOndcOrgAwbNo() bool`

HasOndcOrgAwbNo returns a boolean if a field has been set.

### GetOndcOrgEwaybillno

`func (o *Fulfillment) GetOndcOrgEwaybillno() string`

GetOndcOrgEwaybillno returns the OndcOrgEwaybillno field if non-nil, zero value otherwise.

### GetOndcOrgEwaybillnoOk

`func (o *Fulfillment) GetOndcOrgEwaybillnoOk() (*string, bool)`

GetOndcOrgEwaybillnoOk returns a tuple with the OndcOrgEwaybillno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgEwaybillno

`func (o *Fulfillment) SetOndcOrgEwaybillno(v string)`

SetOndcOrgEwaybillno sets OndcOrgEwaybillno field to given value.

### HasOndcOrgEwaybillno

`func (o *Fulfillment) HasOndcOrgEwaybillno() bool`

HasOndcOrgEwaybillno returns a boolean if a field has been set.

### GetOndcOrgEbnexpirydate

`func (o *Fulfillment) GetOndcOrgEbnexpirydate() time.Time`

GetOndcOrgEbnexpirydate returns the OndcOrgEbnexpirydate field if non-nil, zero value otherwise.

### GetOndcOrgEbnexpirydateOk

`func (o *Fulfillment) GetOndcOrgEbnexpirydateOk() (*time.Time, bool)`

GetOndcOrgEbnexpirydateOk returns a tuple with the OndcOrgEbnexpirydate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgEbnexpirydate

`func (o *Fulfillment) SetOndcOrgEbnexpirydate(v time.Time)`

SetOndcOrgEbnexpirydate sets OndcOrgEbnexpirydate field to given value.

### HasOndcOrgEbnexpirydate

`func (o *Fulfillment) HasOndcOrgEbnexpirydate() bool`

HasOndcOrgEbnexpirydate returns a boolean if a field has been set.

### GetProviderId

`func (o *Fulfillment) GetProviderId() Id`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *Fulfillment) GetProviderIdOk() (*Id, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *Fulfillment) SetProviderId(v Id)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *Fulfillment) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetRating

`func (o *Fulfillment) GetRating() Value`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *Fulfillment) GetRatingOk() (*Value, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *Fulfillment) SetRating(v Value)`

SetRating sets Rating field to given value.

### HasRating

`func (o *Fulfillment) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetState

`func (o *Fulfillment) GetState() State`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Fulfillment) GetStateOk() (*State, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Fulfillment) SetState(v State)`

SetState sets State field to given value.

### HasState

`func (o *Fulfillment) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTracking

`func (o *Fulfillment) GetTracking() bool`

GetTracking returns the Tracking field if non-nil, zero value otherwise.

### GetTrackingOk

`func (o *Fulfillment) GetTrackingOk() (*bool, bool)`

GetTrackingOk returns a tuple with the Tracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracking

`func (o *Fulfillment) SetTracking(v bool)`

SetTracking sets Tracking field to given value.

### HasTracking

`func (o *Fulfillment) HasTracking() bool`

HasTracking returns a boolean if a field has been set.

### GetCustomer

`func (o *Fulfillment) GetCustomer() FulfillmentCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *Fulfillment) GetCustomerOk() (*FulfillmentCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *Fulfillment) SetCustomer(v FulfillmentCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *Fulfillment) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetAgent

`func (o *Fulfillment) GetAgent() Agent`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *Fulfillment) GetAgentOk() (*Agent, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *Fulfillment) SetAgent(v Agent)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *Fulfillment) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetPerson

`func (o *Fulfillment) GetPerson() Person`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *Fulfillment) GetPersonOk() (*Person, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *Fulfillment) SetPerson(v Person)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *Fulfillment) HasPerson() bool`

HasPerson returns a boolean if a field has been set.

### GetContact

`func (o *Fulfillment) GetContact() Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Fulfillment) GetContactOk() (*Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Fulfillment) SetContact(v Contact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *Fulfillment) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetVehicle

`func (o *Fulfillment) GetVehicle() Vehicle`

GetVehicle returns the Vehicle field if non-nil, zero value otherwise.

### GetVehicleOk

`func (o *Fulfillment) GetVehicleOk() (*Vehicle, bool)`

GetVehicleOk returns a tuple with the Vehicle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicle

`func (o *Fulfillment) SetVehicle(v Vehicle)`

SetVehicle sets Vehicle field to given value.

### HasVehicle

`func (o *Fulfillment) HasVehicle() bool`

HasVehicle returns a boolean if a field has been set.

### GetStart

`func (o *Fulfillment) GetStart() FulfillmentStart`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Fulfillment) GetStartOk() (*FulfillmentStart, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Fulfillment) SetStart(v FulfillmentStart)`

SetStart sets Start field to given value.

### HasStart

`func (o *Fulfillment) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *Fulfillment) GetEnd() FulfillmentEnd`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Fulfillment) GetEndOk() (*FulfillmentEnd, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Fulfillment) SetEnd(v FulfillmentEnd)`

SetEnd sets End field to given value.

### HasEnd

`func (o *Fulfillment) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetRateable

`func (o *Fulfillment) GetRateable() bool`

GetRateable returns the Rateable field if non-nil, zero value otherwise.

### GetRateableOk

`func (o *Fulfillment) GetRateableOk() (*bool, bool)`

GetRateableOk returns a tuple with the Rateable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateable

`func (o *Fulfillment) SetRateable(v bool)`

SetRateable sets Rateable field to given value.

### HasRateable

`func (o *Fulfillment) HasRateable() bool`

HasRateable returns a boolean if a field has been set.

### GetTags

`func (o *Fulfillment) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Fulfillment) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Fulfillment) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Fulfillment) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


