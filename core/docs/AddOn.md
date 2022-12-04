# AddOn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the add-on. This follows the syntax {item.id}/add-on/{add-on unique id} for item specific add-on OR  | [optional] 
**Descriptor** | Pointer to [**Descriptor**](Descriptor.md) |  | [optional] 
**Price** | Pointer to [**Price**](Price.md) |  | [optional] 

## Methods

### NewAddOn

`func NewAddOn() *AddOn`

NewAddOn instantiates a new AddOn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddOnWithDefaults

`func NewAddOnWithDefaults() *AddOn`

NewAddOnWithDefaults instantiates a new AddOn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddOn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddOn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddOn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AddOn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescriptor

`func (o *AddOn) GetDescriptor() Descriptor`

GetDescriptor returns the Descriptor field if non-nil, zero value otherwise.

### GetDescriptorOk

`func (o *AddOn) GetDescriptorOk() (*Descriptor, bool)`

GetDescriptorOk returns a tuple with the Descriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptor

`func (o *AddOn) SetDescriptor(v Descriptor)`

SetDescriptor sets Descriptor field to given value.

### HasDescriptor

`func (o *AddOn) HasDescriptor() bool`

HasDescriptor returns a boolean if a field has been set.

### GetPrice

`func (o *AddOn) GetPrice() Price`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AddOn) GetPriceOk() (*Price, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AddOn) SetPrice(v Price)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *AddOn) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


