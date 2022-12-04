# OrderOndcOrgLinkedOrderOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**Id**](Id.md) |  | [optional] 
**Weight** | Pointer to [**Scalar**](Scalar.md) |  | [optional] 
**Dimensions** | Pointer to [**Dimensions**](Dimensions.md) |  | [optional] 
**DeclaredValue** | Pointer to [**Price**](Price.md) |  | [optional] 
**TaxableValue** | Pointer to [**Price**](Price.md) |  | [optional] 
**HsnCode** | Pointer to **string** |  | [optional] 
**SgstAmount** | Pointer to **string** | Describes a decimal value | [optional] 
**CgstAmount** | Pointer to **string** | Describes a decimal value | [optional] 
**IgstAmount** | Pointer to **string** | Describes a decimal value | [optional] 

## Methods

### NewOrderOndcOrgLinkedOrderOrder

`func NewOrderOndcOrgLinkedOrderOrder() *OrderOndcOrgLinkedOrderOrder`

NewOrderOndcOrgLinkedOrderOrder instantiates a new OrderOndcOrgLinkedOrderOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderOndcOrgLinkedOrderOrderWithDefaults

`func NewOrderOndcOrgLinkedOrderOrderWithDefaults() *OrderOndcOrgLinkedOrderOrder`

NewOrderOndcOrgLinkedOrderOrderWithDefaults instantiates a new OrderOndcOrgLinkedOrderOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderOndcOrgLinkedOrderOrder) GetId() Id`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderOndcOrgLinkedOrderOrder) GetIdOk() (*Id, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderOndcOrgLinkedOrderOrder) SetId(v Id)`

SetId sets Id field to given value.

### HasId

`func (o *OrderOndcOrgLinkedOrderOrder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWeight

`func (o *OrderOndcOrgLinkedOrderOrder) GetWeight() Scalar`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *OrderOndcOrgLinkedOrderOrder) GetWeightOk() (*Scalar, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *OrderOndcOrgLinkedOrderOrder) SetWeight(v Scalar)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *OrderOndcOrgLinkedOrderOrder) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetDimensions

`func (o *OrderOndcOrgLinkedOrderOrder) GetDimensions() Dimensions`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *OrderOndcOrgLinkedOrderOrder) GetDimensionsOk() (*Dimensions, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *OrderOndcOrgLinkedOrderOrder) SetDimensions(v Dimensions)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *OrderOndcOrgLinkedOrderOrder) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetDeclaredValue

`func (o *OrderOndcOrgLinkedOrderOrder) GetDeclaredValue() Price`

GetDeclaredValue returns the DeclaredValue field if non-nil, zero value otherwise.

### GetDeclaredValueOk

`func (o *OrderOndcOrgLinkedOrderOrder) GetDeclaredValueOk() (*Price, bool)`

GetDeclaredValueOk returns a tuple with the DeclaredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredValue

`func (o *OrderOndcOrgLinkedOrderOrder) SetDeclaredValue(v Price)`

SetDeclaredValue sets DeclaredValue field to given value.

### HasDeclaredValue

`func (o *OrderOndcOrgLinkedOrderOrder) HasDeclaredValue() bool`

HasDeclaredValue returns a boolean if a field has been set.

### GetTaxableValue

`func (o *OrderOndcOrgLinkedOrderOrder) GetTaxableValue() Price`

GetTaxableValue returns the TaxableValue field if non-nil, zero value otherwise.

### GetTaxableValueOk

`func (o *OrderOndcOrgLinkedOrderOrder) GetTaxableValueOk() (*Price, bool)`

GetTaxableValueOk returns a tuple with the TaxableValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableValue

`func (o *OrderOndcOrgLinkedOrderOrder) SetTaxableValue(v Price)`

SetTaxableValue sets TaxableValue field to given value.

### HasTaxableValue

`func (o *OrderOndcOrgLinkedOrderOrder) HasTaxableValue() bool`

HasTaxableValue returns a boolean if a field has been set.

### GetHsnCode

`func (o *OrderOndcOrgLinkedOrderOrder) GetHsnCode() string`

GetHsnCode returns the HsnCode field if non-nil, zero value otherwise.

### GetHsnCodeOk

`func (o *OrderOndcOrgLinkedOrderOrder) GetHsnCodeOk() (*string, bool)`

GetHsnCodeOk returns a tuple with the HsnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsnCode

`func (o *OrderOndcOrgLinkedOrderOrder) SetHsnCode(v string)`

SetHsnCode sets HsnCode field to given value.

### HasHsnCode

`func (o *OrderOndcOrgLinkedOrderOrder) HasHsnCode() bool`

HasHsnCode returns a boolean if a field has been set.

### GetSgstAmount

`func (o *OrderOndcOrgLinkedOrderOrder) GetSgstAmount() string`

GetSgstAmount returns the SgstAmount field if non-nil, zero value otherwise.

### GetSgstAmountOk

`func (o *OrderOndcOrgLinkedOrderOrder) GetSgstAmountOk() (*string, bool)`

GetSgstAmountOk returns a tuple with the SgstAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgstAmount

`func (o *OrderOndcOrgLinkedOrderOrder) SetSgstAmount(v string)`

SetSgstAmount sets SgstAmount field to given value.

### HasSgstAmount

`func (o *OrderOndcOrgLinkedOrderOrder) HasSgstAmount() bool`

HasSgstAmount returns a boolean if a field has been set.

### GetCgstAmount

`func (o *OrderOndcOrgLinkedOrderOrder) GetCgstAmount() string`

GetCgstAmount returns the CgstAmount field if non-nil, zero value otherwise.

### GetCgstAmountOk

`func (o *OrderOndcOrgLinkedOrderOrder) GetCgstAmountOk() (*string, bool)`

GetCgstAmountOk returns a tuple with the CgstAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCgstAmount

`func (o *OrderOndcOrgLinkedOrderOrder) SetCgstAmount(v string)`

SetCgstAmount sets CgstAmount field to given value.

### HasCgstAmount

`func (o *OrderOndcOrgLinkedOrderOrder) HasCgstAmount() bool`

HasCgstAmount returns a boolean if a field has been set.

### GetIgstAmount

`func (o *OrderOndcOrgLinkedOrderOrder) GetIgstAmount() string`

GetIgstAmount returns the IgstAmount field if non-nil, zero value otherwise.

### GetIgstAmountOk

`func (o *OrderOndcOrgLinkedOrderOrder) GetIgstAmountOk() (*string, bool)`

GetIgstAmountOk returns a tuple with the IgstAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgstAmount

`func (o *OrderOndcOrgLinkedOrderOrder) SetIgstAmount(v string)`

SetIgstAmount sets IgstAmount field to given value.

### HasIgstAmount

`func (o *OrderOndcOrgLinkedOrderOrder) HasIgstAmount() bool`

HasIgstAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


