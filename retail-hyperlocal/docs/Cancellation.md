# Cancellation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **string** |  | [optional] 
**Policies** | Pointer to [**[]Policy**](Policy.md) |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**CancelledBy** | Pointer to **string** |  | [optional] 
**Reasons** | Pointer to [**Option**](Option.md) |  | [optional] 
**SelectedReason** | Pointer to [**CancellationSelectedReason**](CancellationSelectedReason.md) |  | [optional] 
**AdditionalDescription** | Pointer to [**Descriptor**](Descriptor.md) |  | [optional] 

## Methods

### NewCancellation

`func NewCancellation() *Cancellation`

NewCancellation instantiates a new Cancellation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancellationWithDefaults

`func NewCancellationWithDefaults() *Cancellation`

NewCancellationWithDefaults instantiates a new Cancellation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Cancellation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Cancellation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Cancellation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Cancellation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRefId

`func (o *Cancellation) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *Cancellation) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *Cancellation) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *Cancellation) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetPolicies

`func (o *Cancellation) GetPolicies() []Policy`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *Cancellation) GetPoliciesOk() (*[]Policy, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *Cancellation) SetPolicies(v []Policy)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *Cancellation) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetTime

`func (o *Cancellation) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Cancellation) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Cancellation) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *Cancellation) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetCancelledBy

`func (o *Cancellation) GetCancelledBy() string`

GetCancelledBy returns the CancelledBy field if non-nil, zero value otherwise.

### GetCancelledByOk

`func (o *Cancellation) GetCancelledByOk() (*string, bool)`

GetCancelledByOk returns a tuple with the CancelledBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledBy

`func (o *Cancellation) SetCancelledBy(v string)`

SetCancelledBy sets CancelledBy field to given value.

### HasCancelledBy

`func (o *Cancellation) HasCancelledBy() bool`

HasCancelledBy returns a boolean if a field has been set.

### GetReasons

`func (o *Cancellation) GetReasons() Option`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *Cancellation) GetReasonsOk() (*Option, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *Cancellation) SetReasons(v Option)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *Cancellation) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### GetSelectedReason

`func (o *Cancellation) GetSelectedReason() CancellationSelectedReason`

GetSelectedReason returns the SelectedReason field if non-nil, zero value otherwise.

### GetSelectedReasonOk

`func (o *Cancellation) GetSelectedReasonOk() (*CancellationSelectedReason, bool)`

GetSelectedReasonOk returns a tuple with the SelectedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedReason

`func (o *Cancellation) SetSelectedReason(v CancellationSelectedReason)`

SetSelectedReason sets SelectedReason field to given value.

### HasSelectedReason

`func (o *Cancellation) HasSelectedReason() bool`

HasSelectedReason returns a boolean if a field has been set.

### GetAdditionalDescription

`func (o *Cancellation) GetAdditionalDescription() Descriptor`

GetAdditionalDescription returns the AdditionalDescription field if non-nil, zero value otherwise.

### GetAdditionalDescriptionOk

`func (o *Cancellation) GetAdditionalDescriptionOk() (*Descriptor, bool)`

GetAdditionalDescriptionOk returns a tuple with the AdditionalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDescription

`func (o *Cancellation) SetAdditionalDescription(v Descriptor)`

SetAdditionalDescription sets AdditionalDescription field to given value.

### HasAdditionalDescription

`func (o *Cancellation) HasAdditionalDescription() bool`

HasAdditionalDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


