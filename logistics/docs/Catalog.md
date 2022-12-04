# Catalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BppDescriptor** | Pointer to [**Descriptor**](Descriptor.md) |  | [optional] 
**BppCategories** | Pointer to [**[]Category**](Category.md) |  | [optional] 
**BppFulfillments** | Pointer to [**[]Fulfillment**](Fulfillment.md) |  | [optional] 
**BppPayments** | Pointer to [**[]Payment**](Payment.md) |  | [optional] 
**BppOffers** | Pointer to [**[]Offer**](Offer.md) |  | [optional] 
**BppProviders** | [**[]Provider**](Provider.md) |  | 
**Exp** | Pointer to **time.Time** | Time after which catalog has to be refreshed | [optional] 

## Methods

### NewCatalog

`func NewCatalog(bppProviders []Provider, ) *Catalog`

NewCatalog instantiates a new Catalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogWithDefaults

`func NewCatalogWithDefaults() *Catalog`

NewCatalogWithDefaults instantiates a new Catalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBppDescriptor

`func (o *Catalog) GetBppDescriptor() Descriptor`

GetBppDescriptor returns the BppDescriptor field if non-nil, zero value otherwise.

### GetBppDescriptorOk

`func (o *Catalog) GetBppDescriptorOk() (*Descriptor, bool)`

GetBppDescriptorOk returns a tuple with the BppDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBppDescriptor

`func (o *Catalog) SetBppDescriptor(v Descriptor)`

SetBppDescriptor sets BppDescriptor field to given value.

### HasBppDescriptor

`func (o *Catalog) HasBppDescriptor() bool`

HasBppDescriptor returns a boolean if a field has been set.

### GetBppCategories

`func (o *Catalog) GetBppCategories() []Category`

GetBppCategories returns the BppCategories field if non-nil, zero value otherwise.

### GetBppCategoriesOk

`func (o *Catalog) GetBppCategoriesOk() (*[]Category, bool)`

GetBppCategoriesOk returns a tuple with the BppCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBppCategories

`func (o *Catalog) SetBppCategories(v []Category)`

SetBppCategories sets BppCategories field to given value.

### HasBppCategories

`func (o *Catalog) HasBppCategories() bool`

HasBppCategories returns a boolean if a field has been set.

### GetBppFulfillments

`func (o *Catalog) GetBppFulfillments() []Fulfillment`

GetBppFulfillments returns the BppFulfillments field if non-nil, zero value otherwise.

### GetBppFulfillmentsOk

`func (o *Catalog) GetBppFulfillmentsOk() (*[]Fulfillment, bool)`

GetBppFulfillmentsOk returns a tuple with the BppFulfillments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBppFulfillments

`func (o *Catalog) SetBppFulfillments(v []Fulfillment)`

SetBppFulfillments sets BppFulfillments field to given value.

### HasBppFulfillments

`func (o *Catalog) HasBppFulfillments() bool`

HasBppFulfillments returns a boolean if a field has been set.

### GetBppPayments

`func (o *Catalog) GetBppPayments() []Payment`

GetBppPayments returns the BppPayments field if non-nil, zero value otherwise.

### GetBppPaymentsOk

`func (o *Catalog) GetBppPaymentsOk() (*[]Payment, bool)`

GetBppPaymentsOk returns a tuple with the BppPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBppPayments

`func (o *Catalog) SetBppPayments(v []Payment)`

SetBppPayments sets BppPayments field to given value.

### HasBppPayments

`func (o *Catalog) HasBppPayments() bool`

HasBppPayments returns a boolean if a field has been set.

### GetBppOffers

`func (o *Catalog) GetBppOffers() []Offer`

GetBppOffers returns the BppOffers field if non-nil, zero value otherwise.

### GetBppOffersOk

`func (o *Catalog) GetBppOffersOk() (*[]Offer, bool)`

GetBppOffersOk returns a tuple with the BppOffers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBppOffers

`func (o *Catalog) SetBppOffers(v []Offer)`

SetBppOffers sets BppOffers field to given value.

### HasBppOffers

`func (o *Catalog) HasBppOffers() bool`

HasBppOffers returns a boolean if a field has been set.

### GetBppProviders

`func (o *Catalog) GetBppProviders() []Provider`

GetBppProviders returns the BppProviders field if non-nil, zero value otherwise.

### GetBppProvidersOk

`func (o *Catalog) GetBppProvidersOk() (*[]Provider, bool)`

GetBppProvidersOk returns a tuple with the BppProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBppProviders

`func (o *Catalog) SetBppProviders(v []Provider)`

SetBppProviders sets BppProviders field to given value.


### GetExp

`func (o *Catalog) GetExp() time.Time`

GetExp returns the Exp field if non-nil, zero value otherwise.

### GetExpOk

`func (o *Catalog) GetExpOk() (*time.Time, bool)`

GetExpOk returns a tuple with the Exp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExp

`func (o *Catalog) SetExp(v time.Time)`

SetExp sets Exp field to given value.

### HasExp

`func (o *Catalog) HasExp() bool`

HasExp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


