# Context

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | [**Domain**](Domain.md) |  | 
**Country** | [**Code**](Code.md) |  | 
**City** | [**Code**](Code.md) |  | 
**Action** | **string** | Defines the ONDC API call. Any actions other than the enumerated actions are not supported by ONDC Protocol | 
**CoreVersion** | **string** | Version of ONDC core API specification being used | 
**BapId** | **string** | Unique id of the Buyer App. By default it is the fully qualified domain name of the Buyer App | 
**BapUri** | **string** | URI of the Buyer App for accepting callbacks. Must have the same domain name as the bap_id | 
**BppId** | Pointer to **string** | Unique id of the Seller App. By default it is the fully qualified domain name of the Seller App, mandatory for all peer-to-peer API requests, i.e. except search and on_search | [optional] 
**BppUri** | Pointer to **string** | URI of the Seller App. Must have the same domain name as the bap_id, mandatory for all peer-to-peer API requests, i.e. except search and on_search | [optional] 
**TransactionId** | **string** | This is a unique value which persists across all API calls from search through confirm | 
**MessageId** | **string** | This is a unique value which persists during a request / callback cycle | 
**Timestamp** | **time.Time** | Time of request generation in RFC3339 format | 
**Key** | Pointer to **string** | The encryption public key of the sender | [optional] 
**Ttl** | Pointer to **string** | Timestamp for which this message holds valid in ISO8601 durations format - Outer limit for ttl for search, select, init, confirm, status, track, cancel, update, rating, support is &#39;PT30S&#39; which is 30 seconds, different buyer apps can change this to meet their UX requirements, but it shouldn&#39;t exceed this outer limit | [optional] 

## Methods

### NewContext

`func NewContext(domain Domain, country Code, city Code, action string, coreVersion string, bapId string, bapUri string, transactionId string, messageId string, timestamp time.Time, ) *Context`

NewContext instantiates a new Context object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextWithDefaults

`func NewContextWithDefaults() *Context`

NewContextWithDefaults instantiates a new Context object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *Context) GetDomain() Domain`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Context) GetDomainOk() (*Domain, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Context) SetDomain(v Domain)`

SetDomain sets Domain field to given value.


### GetCountry

`func (o *Context) GetCountry() Code`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Context) GetCountryOk() (*Code, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Context) SetCountry(v Code)`

SetCountry sets Country field to given value.


### GetCity

`func (o *Context) GetCity() Code`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Context) GetCityOk() (*Code, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Context) SetCity(v Code)`

SetCity sets City field to given value.


### GetAction

`func (o *Context) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Context) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Context) SetAction(v string)`

SetAction sets Action field to given value.


### GetCoreVersion

`func (o *Context) GetCoreVersion() string`

GetCoreVersion returns the CoreVersion field if non-nil, zero value otherwise.

### GetCoreVersionOk

`func (o *Context) GetCoreVersionOk() (*string, bool)`

GetCoreVersionOk returns a tuple with the CoreVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreVersion

`func (o *Context) SetCoreVersion(v string)`

SetCoreVersion sets CoreVersion field to given value.


### GetBapId

`func (o *Context) GetBapId() string`

GetBapId returns the BapId field if non-nil, zero value otherwise.

### GetBapIdOk

`func (o *Context) GetBapIdOk() (*string, bool)`

GetBapIdOk returns a tuple with the BapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBapId

`func (o *Context) SetBapId(v string)`

SetBapId sets BapId field to given value.


### GetBapUri

`func (o *Context) GetBapUri() string`

GetBapUri returns the BapUri field if non-nil, zero value otherwise.

### GetBapUriOk

`func (o *Context) GetBapUriOk() (*string, bool)`

GetBapUriOk returns a tuple with the BapUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBapUri

`func (o *Context) SetBapUri(v string)`

SetBapUri sets BapUri field to given value.


### GetBppId

`func (o *Context) GetBppId() string`

GetBppId returns the BppId field if non-nil, zero value otherwise.

### GetBppIdOk

`func (o *Context) GetBppIdOk() (*string, bool)`

GetBppIdOk returns a tuple with the BppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBppId

`func (o *Context) SetBppId(v string)`

SetBppId sets BppId field to given value.

### HasBppId

`func (o *Context) HasBppId() bool`

HasBppId returns a boolean if a field has been set.

### GetBppUri

`func (o *Context) GetBppUri() string`

GetBppUri returns the BppUri field if non-nil, zero value otherwise.

### GetBppUriOk

`func (o *Context) GetBppUriOk() (*string, bool)`

GetBppUriOk returns a tuple with the BppUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBppUri

`func (o *Context) SetBppUri(v string)`

SetBppUri sets BppUri field to given value.

### HasBppUri

`func (o *Context) HasBppUri() bool`

HasBppUri returns a boolean if a field has been set.

### GetTransactionId

`func (o *Context) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *Context) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *Context) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetMessageId

`func (o *Context) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *Context) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *Context) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.


### GetTimestamp

`func (o *Context) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Context) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Context) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetKey

`func (o *Context) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Context) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Context) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Context) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetTtl

`func (o *Context) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *Context) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *Context) SetTtl(v string)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *Context) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


