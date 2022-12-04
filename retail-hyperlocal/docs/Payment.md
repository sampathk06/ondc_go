# Payment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | Pointer to **string** | A payment uri to be called by the Buyer App. If empty, then the payment is to be done offline. The details of payment should be present in the params object. If &#x60;&#x60;&#x60;tl_method&#x60;&#x60;&#x60; &#x3D; http/get, then the payment details will be sent as url params. Two url param values, &#x60;&#x60;&#x60;$transaction_id&#x60;&#x60;&#x60; and &#x60;&#x60;&#x60;$amount&#x60;&#x60;&#x60; are mandatory. And example url would be : https://www.example.com/pay?txid&#x3D;$transaction_id&amp;amount&#x3D;$amount&amp;vpa&#x3D;upiid&amp;payee&#x3D;shopez&amp;billno&#x3D;1234 | [optional] 
**TlMethod** | Pointer to **string** |  | [optional] 
**Params** | Pointer to [**PaymentParams**](PaymentParams.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Time** | Pointer to [**Time**](Time.md) |  | [optional] 
**CollectedBy** | Pointer to **string** |  | [optional] 
**OndcOrgCollectedByStatus** | Pointer to **string** |  | [optional] 
**OndcOrgBuyerAppFinderFeeType** | Pointer to **string** |  | [optional] 
**OndcOrgBuyerAppFinderFeeAmount** | Pointer to **string** | Describes a decimal value | [optional] 
**OndcOrgWithholdingAmount** | Pointer to **string** | Describes a decimal value | [optional] 
**OndcOrgWithholdingAmountStatus** | Pointer to **string** |  | [optional] 
**OndcOrgReturnWindow** | Pointer to **string** | return window for withholding amount in ISO8601 durations format e.g. &#39;PT24H&#39; indicates 24 hour return window | [optional] 
**OndcOrgReturnWindowStatus** | Pointer to **string** |  | [optional] 
**OndcOrgSettlementBasis** | Pointer to **string** | In case of prepaid payment, whether settlement between counterparties should be on the basis of collection, shipment or delivery | [optional] 
**OndcOrgSettlementBasisStatus** | Pointer to **string** |  | [optional] 
**OndcOrgSettlementWindow** | Pointer to **string** | settlement window for the counterparty in ISO8601 durations format e.g. &#39;PT48H&#39; indicates 48 hour return window | [optional] 
**OndcOrgSettlementWindowStatus** | Pointer to **string** |  | [optional] 
**OndcOrgSettlementDetails** | Pointer to [**[]PaymentOndcOrgSettlementDetailsInner**](PaymentOndcOrgSettlementDetailsInner.md) |  | [optional] 

## Methods

### NewPayment

`func NewPayment() *Payment`

NewPayment instantiates a new Payment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentWithDefaults

`func NewPaymentWithDefaults() *Payment`

NewPaymentWithDefaults instantiates a new Payment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUri

`func (o *Payment) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Payment) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Payment) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *Payment) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetTlMethod

`func (o *Payment) GetTlMethod() string`

GetTlMethod returns the TlMethod field if non-nil, zero value otherwise.

### GetTlMethodOk

`func (o *Payment) GetTlMethodOk() (*string, bool)`

GetTlMethodOk returns a tuple with the TlMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlMethod

`func (o *Payment) SetTlMethod(v string)`

SetTlMethod sets TlMethod field to given value.

### HasTlMethod

`func (o *Payment) HasTlMethod() bool`

HasTlMethod returns a boolean if a field has been set.

### GetParams

`func (o *Payment) GetParams() PaymentParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *Payment) GetParamsOk() (*PaymentParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *Payment) SetParams(v PaymentParams)`

SetParams sets Params field to given value.

### HasParams

`func (o *Payment) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetType

`func (o *Payment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Payment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Payment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Payment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *Payment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Payment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Payment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Payment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTime

`func (o *Payment) GetTime() Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Payment) GetTimeOk() (*Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Payment) SetTime(v Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *Payment) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetCollectedBy

`func (o *Payment) GetCollectedBy() string`

GetCollectedBy returns the CollectedBy field if non-nil, zero value otherwise.

### GetCollectedByOk

`func (o *Payment) GetCollectedByOk() (*string, bool)`

GetCollectedByOk returns a tuple with the CollectedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectedBy

`func (o *Payment) SetCollectedBy(v string)`

SetCollectedBy sets CollectedBy field to given value.

### HasCollectedBy

`func (o *Payment) HasCollectedBy() bool`

HasCollectedBy returns a boolean if a field has been set.

### GetOndcOrgCollectedByStatus

`func (o *Payment) GetOndcOrgCollectedByStatus() string`

GetOndcOrgCollectedByStatus returns the OndcOrgCollectedByStatus field if non-nil, zero value otherwise.

### GetOndcOrgCollectedByStatusOk

`func (o *Payment) GetOndcOrgCollectedByStatusOk() (*string, bool)`

GetOndcOrgCollectedByStatusOk returns a tuple with the OndcOrgCollectedByStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgCollectedByStatus

`func (o *Payment) SetOndcOrgCollectedByStatus(v string)`

SetOndcOrgCollectedByStatus sets OndcOrgCollectedByStatus field to given value.

### HasOndcOrgCollectedByStatus

`func (o *Payment) HasOndcOrgCollectedByStatus() bool`

HasOndcOrgCollectedByStatus returns a boolean if a field has been set.

### GetOndcOrgBuyerAppFinderFeeType

`func (o *Payment) GetOndcOrgBuyerAppFinderFeeType() string`

GetOndcOrgBuyerAppFinderFeeType returns the OndcOrgBuyerAppFinderFeeType field if non-nil, zero value otherwise.

### GetOndcOrgBuyerAppFinderFeeTypeOk

`func (o *Payment) GetOndcOrgBuyerAppFinderFeeTypeOk() (*string, bool)`

GetOndcOrgBuyerAppFinderFeeTypeOk returns a tuple with the OndcOrgBuyerAppFinderFeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgBuyerAppFinderFeeType

`func (o *Payment) SetOndcOrgBuyerAppFinderFeeType(v string)`

SetOndcOrgBuyerAppFinderFeeType sets OndcOrgBuyerAppFinderFeeType field to given value.

### HasOndcOrgBuyerAppFinderFeeType

`func (o *Payment) HasOndcOrgBuyerAppFinderFeeType() bool`

HasOndcOrgBuyerAppFinderFeeType returns a boolean if a field has been set.

### GetOndcOrgBuyerAppFinderFeeAmount

`func (o *Payment) GetOndcOrgBuyerAppFinderFeeAmount() string`

GetOndcOrgBuyerAppFinderFeeAmount returns the OndcOrgBuyerAppFinderFeeAmount field if non-nil, zero value otherwise.

### GetOndcOrgBuyerAppFinderFeeAmountOk

`func (o *Payment) GetOndcOrgBuyerAppFinderFeeAmountOk() (*string, bool)`

GetOndcOrgBuyerAppFinderFeeAmountOk returns a tuple with the OndcOrgBuyerAppFinderFeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgBuyerAppFinderFeeAmount

`func (o *Payment) SetOndcOrgBuyerAppFinderFeeAmount(v string)`

SetOndcOrgBuyerAppFinderFeeAmount sets OndcOrgBuyerAppFinderFeeAmount field to given value.

### HasOndcOrgBuyerAppFinderFeeAmount

`func (o *Payment) HasOndcOrgBuyerAppFinderFeeAmount() bool`

HasOndcOrgBuyerAppFinderFeeAmount returns a boolean if a field has been set.

### GetOndcOrgWithholdingAmount

`func (o *Payment) GetOndcOrgWithholdingAmount() string`

GetOndcOrgWithholdingAmount returns the OndcOrgWithholdingAmount field if non-nil, zero value otherwise.

### GetOndcOrgWithholdingAmountOk

`func (o *Payment) GetOndcOrgWithholdingAmountOk() (*string, bool)`

GetOndcOrgWithholdingAmountOk returns a tuple with the OndcOrgWithholdingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgWithholdingAmount

`func (o *Payment) SetOndcOrgWithholdingAmount(v string)`

SetOndcOrgWithholdingAmount sets OndcOrgWithholdingAmount field to given value.

### HasOndcOrgWithholdingAmount

`func (o *Payment) HasOndcOrgWithholdingAmount() bool`

HasOndcOrgWithholdingAmount returns a boolean if a field has been set.

### GetOndcOrgWithholdingAmountStatus

`func (o *Payment) GetOndcOrgWithholdingAmountStatus() string`

GetOndcOrgWithholdingAmountStatus returns the OndcOrgWithholdingAmountStatus field if non-nil, zero value otherwise.

### GetOndcOrgWithholdingAmountStatusOk

`func (o *Payment) GetOndcOrgWithholdingAmountStatusOk() (*string, bool)`

GetOndcOrgWithholdingAmountStatusOk returns a tuple with the OndcOrgWithholdingAmountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgWithholdingAmountStatus

`func (o *Payment) SetOndcOrgWithholdingAmountStatus(v string)`

SetOndcOrgWithholdingAmountStatus sets OndcOrgWithholdingAmountStatus field to given value.

### HasOndcOrgWithholdingAmountStatus

`func (o *Payment) HasOndcOrgWithholdingAmountStatus() bool`

HasOndcOrgWithholdingAmountStatus returns a boolean if a field has been set.

### GetOndcOrgReturnWindow

`func (o *Payment) GetOndcOrgReturnWindow() string`

GetOndcOrgReturnWindow returns the OndcOrgReturnWindow field if non-nil, zero value otherwise.

### GetOndcOrgReturnWindowOk

`func (o *Payment) GetOndcOrgReturnWindowOk() (*string, bool)`

GetOndcOrgReturnWindowOk returns a tuple with the OndcOrgReturnWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgReturnWindow

`func (o *Payment) SetOndcOrgReturnWindow(v string)`

SetOndcOrgReturnWindow sets OndcOrgReturnWindow field to given value.

### HasOndcOrgReturnWindow

`func (o *Payment) HasOndcOrgReturnWindow() bool`

HasOndcOrgReturnWindow returns a boolean if a field has been set.

### GetOndcOrgReturnWindowStatus

`func (o *Payment) GetOndcOrgReturnWindowStatus() string`

GetOndcOrgReturnWindowStatus returns the OndcOrgReturnWindowStatus field if non-nil, zero value otherwise.

### GetOndcOrgReturnWindowStatusOk

`func (o *Payment) GetOndcOrgReturnWindowStatusOk() (*string, bool)`

GetOndcOrgReturnWindowStatusOk returns a tuple with the OndcOrgReturnWindowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgReturnWindowStatus

`func (o *Payment) SetOndcOrgReturnWindowStatus(v string)`

SetOndcOrgReturnWindowStatus sets OndcOrgReturnWindowStatus field to given value.

### HasOndcOrgReturnWindowStatus

`func (o *Payment) HasOndcOrgReturnWindowStatus() bool`

HasOndcOrgReturnWindowStatus returns a boolean if a field has been set.

### GetOndcOrgSettlementBasis

`func (o *Payment) GetOndcOrgSettlementBasis() string`

GetOndcOrgSettlementBasis returns the OndcOrgSettlementBasis field if non-nil, zero value otherwise.

### GetOndcOrgSettlementBasisOk

`func (o *Payment) GetOndcOrgSettlementBasisOk() (*string, bool)`

GetOndcOrgSettlementBasisOk returns a tuple with the OndcOrgSettlementBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgSettlementBasis

`func (o *Payment) SetOndcOrgSettlementBasis(v string)`

SetOndcOrgSettlementBasis sets OndcOrgSettlementBasis field to given value.

### HasOndcOrgSettlementBasis

`func (o *Payment) HasOndcOrgSettlementBasis() bool`

HasOndcOrgSettlementBasis returns a boolean if a field has been set.

### GetOndcOrgSettlementBasisStatus

`func (o *Payment) GetOndcOrgSettlementBasisStatus() string`

GetOndcOrgSettlementBasisStatus returns the OndcOrgSettlementBasisStatus field if non-nil, zero value otherwise.

### GetOndcOrgSettlementBasisStatusOk

`func (o *Payment) GetOndcOrgSettlementBasisStatusOk() (*string, bool)`

GetOndcOrgSettlementBasisStatusOk returns a tuple with the OndcOrgSettlementBasisStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgSettlementBasisStatus

`func (o *Payment) SetOndcOrgSettlementBasisStatus(v string)`

SetOndcOrgSettlementBasisStatus sets OndcOrgSettlementBasisStatus field to given value.

### HasOndcOrgSettlementBasisStatus

`func (o *Payment) HasOndcOrgSettlementBasisStatus() bool`

HasOndcOrgSettlementBasisStatus returns a boolean if a field has been set.

### GetOndcOrgSettlementWindow

`func (o *Payment) GetOndcOrgSettlementWindow() string`

GetOndcOrgSettlementWindow returns the OndcOrgSettlementWindow field if non-nil, zero value otherwise.

### GetOndcOrgSettlementWindowOk

`func (o *Payment) GetOndcOrgSettlementWindowOk() (*string, bool)`

GetOndcOrgSettlementWindowOk returns a tuple with the OndcOrgSettlementWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgSettlementWindow

`func (o *Payment) SetOndcOrgSettlementWindow(v string)`

SetOndcOrgSettlementWindow sets OndcOrgSettlementWindow field to given value.

### HasOndcOrgSettlementWindow

`func (o *Payment) HasOndcOrgSettlementWindow() bool`

HasOndcOrgSettlementWindow returns a boolean if a field has been set.

### GetOndcOrgSettlementWindowStatus

`func (o *Payment) GetOndcOrgSettlementWindowStatus() string`

GetOndcOrgSettlementWindowStatus returns the OndcOrgSettlementWindowStatus field if non-nil, zero value otherwise.

### GetOndcOrgSettlementWindowStatusOk

`func (o *Payment) GetOndcOrgSettlementWindowStatusOk() (*string, bool)`

GetOndcOrgSettlementWindowStatusOk returns a tuple with the OndcOrgSettlementWindowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgSettlementWindowStatus

`func (o *Payment) SetOndcOrgSettlementWindowStatus(v string)`

SetOndcOrgSettlementWindowStatus sets OndcOrgSettlementWindowStatus field to given value.

### HasOndcOrgSettlementWindowStatus

`func (o *Payment) HasOndcOrgSettlementWindowStatus() bool`

HasOndcOrgSettlementWindowStatus returns a boolean if a field has been set.

### GetOndcOrgSettlementDetails

`func (o *Payment) GetOndcOrgSettlementDetails() []PaymentOndcOrgSettlementDetailsInner`

GetOndcOrgSettlementDetails returns the OndcOrgSettlementDetails field if non-nil, zero value otherwise.

### GetOndcOrgSettlementDetailsOk

`func (o *Payment) GetOndcOrgSettlementDetailsOk() (*[]PaymentOndcOrgSettlementDetailsInner, bool)`

GetOndcOrgSettlementDetailsOk returns a tuple with the OndcOrgSettlementDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOndcOrgSettlementDetails

`func (o *Payment) SetOndcOrgSettlementDetails(v []PaymentOndcOrgSettlementDetailsInner)`

SetOndcOrgSettlementDetails sets OndcOrgSettlementDetails field to given value.

### HasOndcOrgSettlementDetails

`func (o *Payment) HasOndcOrgSettlementDetails() bool`

HasOndcOrgSettlementDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


