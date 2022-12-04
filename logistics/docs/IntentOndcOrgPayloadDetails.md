# IntentOndcOrgPayloadDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Weight** | Pointer to [**Scalar**](Scalar.md) |  | [optional] 
**Dimensions** | Pointer to [**Dimensions**](Dimensions.md) |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Value** | Pointer to [**Price**](Price.md) |  | [optional] 
**DangerousGoods** | Pointer to **bool** |  | [optional] 

## Methods

### NewIntentOndcOrgPayloadDetails

`func NewIntentOndcOrgPayloadDetails() *IntentOndcOrgPayloadDetails`

NewIntentOndcOrgPayloadDetails instantiates a new IntentOndcOrgPayloadDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntentOndcOrgPayloadDetailsWithDefaults

`func NewIntentOndcOrgPayloadDetailsWithDefaults() *IntentOndcOrgPayloadDetails`

NewIntentOndcOrgPayloadDetailsWithDefaults instantiates a new IntentOndcOrgPayloadDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWeight

`func (o *IntentOndcOrgPayloadDetails) GetWeight() Scalar`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *IntentOndcOrgPayloadDetails) GetWeightOk() (*Scalar, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *IntentOndcOrgPayloadDetails) SetWeight(v Scalar)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *IntentOndcOrgPayloadDetails) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetDimensions

`func (o *IntentOndcOrgPayloadDetails) GetDimensions() Dimensions`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *IntentOndcOrgPayloadDetails) GetDimensionsOk() (*Dimensions, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *IntentOndcOrgPayloadDetails) SetDimensions(v Dimensions)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *IntentOndcOrgPayloadDetails) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetCategory

`func (o *IntentOndcOrgPayloadDetails) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *IntentOndcOrgPayloadDetails) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *IntentOndcOrgPayloadDetails) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *IntentOndcOrgPayloadDetails) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetValue

`func (o *IntentOndcOrgPayloadDetails) GetValue() Price`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IntentOndcOrgPayloadDetails) GetValueOk() (*Price, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IntentOndcOrgPayloadDetails) SetValue(v Price)`

SetValue sets Value field to given value.

### HasValue

`func (o *IntentOndcOrgPayloadDetails) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDangerousGoods

`func (o *IntentOndcOrgPayloadDetails) GetDangerousGoods() bool`

GetDangerousGoods returns the DangerousGoods field if non-nil, zero value otherwise.

### GetDangerousGoodsOk

`func (o *IntentOndcOrgPayloadDetails) GetDangerousGoodsOk() (*bool, bool)`

GetDangerousGoodsOk returns a tuple with the DangerousGoods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDangerousGoods

`func (o *IntentOndcOrgPayloadDetails) SetDangerousGoods(v bool)`

SetDangerousGoods sets DangerousGoods field to given value.

### HasDangerousGoods

`func (o *IntentOndcOrgPayloadDetails) HasDangerousGoods() bool`

HasDangerousGoods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


