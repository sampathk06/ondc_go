# Quotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | [**Price**](Price.md) |  | 
**Breakup** | Pointer to [**[]QuotationBreakupInner**](QuotationBreakupInner.md) |  | [optional] 
**Ttl** | Pointer to **string** | Describes duration as per ISO8601 format | [optional] 

## Methods

### NewQuotation

`func NewQuotation(price Price, ) *Quotation`

NewQuotation instantiates a new Quotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotationWithDefaults

`func NewQuotationWithDefaults() *Quotation`

NewQuotationWithDefaults instantiates a new Quotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *Quotation) GetPrice() Price`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Quotation) GetPriceOk() (*Price, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Quotation) SetPrice(v Price)`

SetPrice sets Price field to given value.


### GetBreakup

`func (o *Quotation) GetBreakup() []QuotationBreakupInner`

GetBreakup returns the Breakup field if non-nil, zero value otherwise.

### GetBreakupOk

`func (o *Quotation) GetBreakupOk() (*[]QuotationBreakupInner, bool)`

GetBreakupOk returns a tuple with the Breakup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakup

`func (o *Quotation) SetBreakup(v []QuotationBreakupInner)`

SetBreakup sets Breakup field to given value.

### HasBreakup

`func (o *Quotation) HasBreakup() bool`

HasBreakup returns a boolean if a field has been set.

### GetTtl

`func (o *Quotation) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *Quotation) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *Quotation) SetTtl(v string)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *Quotation) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


