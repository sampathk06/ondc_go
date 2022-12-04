# PaymentOndcOrgSettlementDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SettlementCounterparty** | Pointer to **string** |  | [optional] 
**SettlementPhase** | Pointer to **string** |  | [optional] 
**SettlementAmount** | Pointer to **int32** |  | [optional] 
**SettlementType** | Pointer to **string** |  | [optional] 
**SettlementBankAccountNo** | Pointer to **string** |  | [optional] 
**SettlementIfscCode** | Pointer to **string** |  | [optional] 
**UpiAddress** | Pointer to **string** | UPI payment address e.g. VPA | [optional] 
**BankName** | Pointer to **string** | Bank name | [optional] 
**BranchName** | Pointer to **string** | Branch name | [optional] 
**BeneficiaryAddress** | Pointer to **string** | Beneficiary Address | [optional] 
**SettlementStatus** | Pointer to **string** |  | [optional] 
**SettlementReference** | Pointer to **string** | Settlement transaction reference number | [optional] 
**SettlementTimestamp** | Pointer to **time.Time** | Settlement transaction timestamp | [optional] 

## Methods

### NewPaymentOndcOrgSettlementDetailsInner

`func NewPaymentOndcOrgSettlementDetailsInner() *PaymentOndcOrgSettlementDetailsInner`

NewPaymentOndcOrgSettlementDetailsInner instantiates a new PaymentOndcOrgSettlementDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentOndcOrgSettlementDetailsInnerWithDefaults

`func NewPaymentOndcOrgSettlementDetailsInnerWithDefaults() *PaymentOndcOrgSettlementDetailsInner`

NewPaymentOndcOrgSettlementDetailsInnerWithDefaults instantiates a new PaymentOndcOrgSettlementDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettlementCounterparty

`func (o *PaymentOndcOrgSettlementDetailsInner) GetSettlementCounterparty() string`

GetSettlementCounterparty returns the SettlementCounterparty field if non-nil, zero value otherwise.

### GetSettlementCounterpartyOk

`func (o *PaymentOndcOrgSettlementDetailsInner) GetSettlementCounterpartyOk() (*string, bool)`

GetSettlementCounterpartyOk returns a tuple with the SettlementCounterparty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementCounterparty

`func (o *PaymentOndcOrgSettlementDetailsInner) SetSettlementCounterparty(v string)`

SetSettlementCounterparty sets SettlementCounterparty field to given value.

### HasSettlementCounterparty

`func (o *PaymentOndcOrgSettlementDetailsInner) HasSettlementCounterparty() bool`

HasSettlementCounterparty returns a boolean if a field has been set.

### GetSettlementPhase

`func (o *PaymentOndcOrgSettlementDetailsInner) GetSettlementPhase() string`

GetSettlementPhase returns the SettlementPhase field if non-nil, zero value otherwise.

### GetSettlementPhaseOk

`func (o *PaymentOndcOrgSettlementDetailsInner) GetSettlementPhaseOk() (*string, bool)`

GetSettlementPhaseOk returns a tuple with the SettlementPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementPhase

`func (o *PaymentOndcOrgSettlementDetailsInner) SetSettlementPhase(v string)`

SetSettlementPhase sets SettlementPhase field to given value.

### HasSettlementPhase

`func (o *PaymentOndcOrgSettlementDetailsInner) HasSettlementPhase() bool`

HasSettlementPhase returns a boolean if a field has been set.

### GetSettlementAmount

`func (o *PaymentOndcOrgSettlementDetailsInner) GetSettlementAmount() int32`

GetSettlementAmount returns the SettlementAmount field if non-nil, zero value otherwise.

### GetSettlementAmountOk

`func (o *PaymentOndcOrgSettlementDetailsInner) GetSettlementAmountOk() (*int32, bool)`

GetSettlementAmountOk returns a tuple with the SettlementAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementAmount

`func (o *PaymentOndcOrgSettlementDetailsInner) SetSettlementAmount(v int32)`

SetSettlementAmount sets SettlementAmount field to given value.

### HasSettlementAmount

`func (o *PaymentOndcOrgSettlementDetailsInner) HasSettlementAmount() bool`

HasSettlementAmount returns a boolean if a field has been set.

### GetSettlementType

`func (o *PaymentOndcOrgSettlementDetailsInner) GetSettlementType() string`

GetSettlementType returns the SettlementType field if non-nil, zero value otherwise.

### GetSettlementTypeOk

`func (o *PaymentOndcOrgSettlementDetailsInner) GetSettlementTypeOk() (*string, bool)`

GetSettlementTypeOk returns a tuple with the SettlementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementType

`func (o *PaymentOndcOrgSettlementDetailsInner) SetSettlementType(v string)`

SetSettlementType sets SettlementType field to given value.

### HasSettlementType

`func (o *PaymentOndcOrgSettlementDetailsInner) HasSettlementType() bool`

HasSettlementType returns a boolean if a field has been set.

### GetSettlementBankAccountNo

`func (o *PaymentOndcOrgSettlementDetailsInner) GetSettlementBankAccountNo() string`

GetSettlementBankAccountNo returns the SettlementBankAccountNo field if non-nil, zero value otherwise.

### GetSettlementBankAccountNoOk

`func (o *PaymentOndcOrgSettlementDetailsInner) GetSettlementBankAccountNoOk() (*string, bool)`

GetSettlementBankAccountNoOk returns a tuple with the SettlementBankAccountNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementBankAccountNo

`func (o *PaymentOndcOrgSettlementDetailsInner) SetSettlementBankAccountNo(v string)`

SetSettlementBankAccountNo sets SettlementBankAccountNo field to given value.

### HasSettlementBankAccountNo

`func (o *PaymentOndcOrgSettlementDetailsInner) HasSettlementBankAccountNo() bool`

HasSettlementBankAccountNo returns a boolean if a field has been set.

### GetSettlementIfscCode

`func (o *PaymentOndcOrgSettlementDetailsInner) GetSettlementIfscCode() string`

GetSettlementIfscCode returns the SettlementIfscCode field if non-nil, zero value otherwise.

### GetSettlementIfscCodeOk

`func (o *PaymentOndcOrgSettlementDetailsInner) GetSettlementIfscCodeOk() (*string, bool)`

GetSettlementIfscCodeOk returns a tuple with the SettlementIfscCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementIfscCode

`func (o *PaymentOndcOrgSettlementDetailsInner) SetSettlementIfscCode(v string)`

SetSettlementIfscCode sets SettlementIfscCode field to given value.

### HasSettlementIfscCode

`func (o *PaymentOndcOrgSettlementDetailsInner) HasSettlementIfscCode() bool`

HasSettlementIfscCode returns a boolean if a field has been set.

### GetUpiAddress

`func (o *PaymentOndcOrgSettlementDetailsInner) GetUpiAddress() string`

GetUpiAddress returns the UpiAddress field if non-nil, zero value otherwise.

### GetUpiAddressOk

`func (o *PaymentOndcOrgSettlementDetailsInner) GetUpiAddressOk() (*string, bool)`

GetUpiAddressOk returns a tuple with the UpiAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpiAddress

`func (o *PaymentOndcOrgSettlementDetailsInner) SetUpiAddress(v string)`

SetUpiAddress sets UpiAddress field to given value.

### HasUpiAddress

`func (o *PaymentOndcOrgSettlementDetailsInner) HasUpiAddress() bool`

HasUpiAddress returns a boolean if a field has been set.

### GetBankName

`func (o *PaymentOndcOrgSettlementDetailsInner) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *PaymentOndcOrgSettlementDetailsInner) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *PaymentOndcOrgSettlementDetailsInner) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *PaymentOndcOrgSettlementDetailsInner) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### GetBranchName

`func (o *PaymentOndcOrgSettlementDetailsInner) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *PaymentOndcOrgSettlementDetailsInner) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *PaymentOndcOrgSettlementDetailsInner) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.

### HasBranchName

`func (o *PaymentOndcOrgSettlementDetailsInner) HasBranchName() bool`

HasBranchName returns a boolean if a field has been set.

### GetBeneficiaryAddress

`func (o *PaymentOndcOrgSettlementDetailsInner) GetBeneficiaryAddress() string`

GetBeneficiaryAddress returns the BeneficiaryAddress field if non-nil, zero value otherwise.

### GetBeneficiaryAddressOk

`func (o *PaymentOndcOrgSettlementDetailsInner) GetBeneficiaryAddressOk() (*string, bool)`

GetBeneficiaryAddressOk returns a tuple with the BeneficiaryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryAddress

`func (o *PaymentOndcOrgSettlementDetailsInner) SetBeneficiaryAddress(v string)`

SetBeneficiaryAddress sets BeneficiaryAddress field to given value.

### HasBeneficiaryAddress

`func (o *PaymentOndcOrgSettlementDetailsInner) HasBeneficiaryAddress() bool`

HasBeneficiaryAddress returns a boolean if a field has been set.

### GetSettlementStatus

`func (o *PaymentOndcOrgSettlementDetailsInner) GetSettlementStatus() string`

GetSettlementStatus returns the SettlementStatus field if non-nil, zero value otherwise.

### GetSettlementStatusOk

`func (o *PaymentOndcOrgSettlementDetailsInner) GetSettlementStatusOk() (*string, bool)`

GetSettlementStatusOk returns a tuple with the SettlementStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementStatus

`func (o *PaymentOndcOrgSettlementDetailsInner) SetSettlementStatus(v string)`

SetSettlementStatus sets SettlementStatus field to given value.

### HasSettlementStatus

`func (o *PaymentOndcOrgSettlementDetailsInner) HasSettlementStatus() bool`

HasSettlementStatus returns a boolean if a field has been set.

### GetSettlementReference

`func (o *PaymentOndcOrgSettlementDetailsInner) GetSettlementReference() string`

GetSettlementReference returns the SettlementReference field if non-nil, zero value otherwise.

### GetSettlementReferenceOk

`func (o *PaymentOndcOrgSettlementDetailsInner) GetSettlementReferenceOk() (*string, bool)`

GetSettlementReferenceOk returns a tuple with the SettlementReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementReference

`func (o *PaymentOndcOrgSettlementDetailsInner) SetSettlementReference(v string)`

SetSettlementReference sets SettlementReference field to given value.

### HasSettlementReference

`func (o *PaymentOndcOrgSettlementDetailsInner) HasSettlementReference() bool`

HasSettlementReference returns a boolean if a field has been set.

### GetSettlementTimestamp

`func (o *PaymentOndcOrgSettlementDetailsInner) GetSettlementTimestamp() time.Time`

GetSettlementTimestamp returns the SettlementTimestamp field if non-nil, zero value otherwise.

### GetSettlementTimestampOk

`func (o *PaymentOndcOrgSettlementDetailsInner) GetSettlementTimestampOk() (*time.Time, bool)`

GetSettlementTimestampOk returns a tuple with the SettlementTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementTimestamp

`func (o *PaymentOndcOrgSettlementDetailsInner) SetSettlementTimestamp(v time.Time)`

SetSettlementTimestamp sets SettlementTimestamp field to given value.

### HasSettlementTimestamp

`func (o *PaymentOndcOrgSettlementDetailsInner) HasSettlementTimestamp() bool`

HasSettlementTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


