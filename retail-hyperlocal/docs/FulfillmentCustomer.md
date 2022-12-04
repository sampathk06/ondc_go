# FulfillmentCustomer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Person** | Pointer to [**Person**](Person.md) |  | [optional] 
**Contact** | Pointer to [**Contact**](Contact.md) |  | [optional] 

## Methods

### NewFulfillmentCustomer

`func NewFulfillmentCustomer() *FulfillmentCustomer`

NewFulfillmentCustomer instantiates a new FulfillmentCustomer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFulfillmentCustomerWithDefaults

`func NewFulfillmentCustomerWithDefaults() *FulfillmentCustomer`

NewFulfillmentCustomerWithDefaults instantiates a new FulfillmentCustomer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPerson

`func (o *FulfillmentCustomer) GetPerson() Person`

GetPerson returns the Person field if non-nil, zero value otherwise.

### GetPersonOk

`func (o *FulfillmentCustomer) GetPersonOk() (*Person, bool)`

GetPersonOk returns a tuple with the Person field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerson

`func (o *FulfillmentCustomer) SetPerson(v Person)`

SetPerson sets Person field to given value.

### HasPerson

`func (o *FulfillmentCustomer) HasPerson() bool`

HasPerson returns a boolean if a field has been set.

### GetContact

`func (o *FulfillmentCustomer) GetContact() Contact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *FulfillmentCustomer) GetContactOk() (*Contact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *FulfillmentCustomer) SetContact(v Contact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *FulfillmentCustomer) HasContact() bool`

HasContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


