# Subscriber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriberId** | Pointer to **string** | Registered domain name of the subscriber. Must have a valid SSL certificate issued by a Certificate Authority of the operating region | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**CbUrl** | Pointer to **string** | Callback URL of the subscriber. The Registry will call this URL&#39;s on_subscribe API to validate the subscriber\\&#39;s credentials | [optional] 
**Domain** | Pointer to [**Domain**](Domain.md) |  | [optional] 
**City** | Pointer to [**Code**](Code.md) |  | [optional] 
**Country** | Pointer to [**Code**](Code.md) |  | [optional] 
**SigningPublicKey** | Pointer to **string** | Signing Public key of the subscriber. &lt;br/&gt;&lt;br/&gt;Any subscriber platform (Buyer App, Seller App, Gateway) who wants to transact on the network must digitally sign the &#x60;&#x60;&#x60;requestBody&#x60;&#x60;&#x60; using the corresponding private key of this public key and send it in the transport layer header. In case of &#x60;&#x60;&#x60;HTTP&#x60;&#x60;&#x60; it is the &#x60;&#x60;&#x60;Authorization&#x60;&#x60;&#x60; header. &lt;br&gt;&lt;br/&gt;The &#x60;&#x60;&#x60;Authorization&#x60;&#x60;&#x60; will be used to validate the signature of a Buyer App or Seller App.&lt;br/&gt;&lt;br/&gt;Furthermore, if an API call is being proxied or multicast by a ONDC Gateway, the Gateway must use it\\&#39;s signing key to digitally sign the &#x60;&#x60;&#x60;requestBody&#x60;&#x60;&#x60; using the corresponding private key of this public key and send it in the &#x60;&#x60;&#x60;X-Gateway-Authorization&#x60;&#x60;&#x60; header. | [optional] 
**EncryptionPublicKey** | Pointer to **string** | Encryption public key of the Buyer App. Any Seller App must encrypt the &#x60;&#x60;&#x60;requestBody.message&#x60;&#x60;&#x60; value of the &#x60;&#x60;&#x60;on_search&#x60;&#x60;&#x60; API using this public key. | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** | Timestamp when a subscriber was added to the registry with status &#x3D; INITIATED | [optional] 
**Updated** | Pointer to **time.Time** |  | [optional] 
**Expires** | Pointer to **time.Time** | Expiry timestamp in UTC derived from the &#x60;&#x60;&#x60;lease_time&#x60;&#x60;&#x60; of the subscriber | [optional] 

## Methods

### NewSubscriber

`func NewSubscriber() *Subscriber`

NewSubscriber instantiates a new Subscriber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriberWithDefaults

`func NewSubscriberWithDefaults() *Subscriber`

NewSubscriberWithDefaults instantiates a new Subscriber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriberId

`func (o *Subscriber) GetSubscriberId() string`

GetSubscriberId returns the SubscriberId field if non-nil, zero value otherwise.

### GetSubscriberIdOk

`func (o *Subscriber) GetSubscriberIdOk() (*string, bool)`

GetSubscriberIdOk returns a tuple with the SubscriberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberId

`func (o *Subscriber) SetSubscriberId(v string)`

SetSubscriberId sets SubscriberId field to given value.

### HasSubscriberId

`func (o *Subscriber) HasSubscriberId() bool`

HasSubscriberId returns a boolean if a field has been set.

### GetType

`func (o *Subscriber) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Subscriber) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Subscriber) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Subscriber) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCbUrl

`func (o *Subscriber) GetCbUrl() string`

GetCbUrl returns the CbUrl field if non-nil, zero value otherwise.

### GetCbUrlOk

`func (o *Subscriber) GetCbUrlOk() (*string, bool)`

GetCbUrlOk returns a tuple with the CbUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCbUrl

`func (o *Subscriber) SetCbUrl(v string)`

SetCbUrl sets CbUrl field to given value.

### HasCbUrl

`func (o *Subscriber) HasCbUrl() bool`

HasCbUrl returns a boolean if a field has been set.

### GetDomain

`func (o *Subscriber) GetDomain() Domain`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Subscriber) GetDomainOk() (*Domain, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Subscriber) SetDomain(v Domain)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Subscriber) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetCity

`func (o *Subscriber) GetCity() Code`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Subscriber) GetCityOk() (*Code, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Subscriber) SetCity(v Code)`

SetCity sets City field to given value.

### HasCity

`func (o *Subscriber) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *Subscriber) GetCountry() Code`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Subscriber) GetCountryOk() (*Code, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Subscriber) SetCountry(v Code)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Subscriber) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetSigningPublicKey

`func (o *Subscriber) GetSigningPublicKey() string`

GetSigningPublicKey returns the SigningPublicKey field if non-nil, zero value otherwise.

### GetSigningPublicKeyOk

`func (o *Subscriber) GetSigningPublicKeyOk() (*string, bool)`

GetSigningPublicKeyOk returns a tuple with the SigningPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningPublicKey

`func (o *Subscriber) SetSigningPublicKey(v string)`

SetSigningPublicKey sets SigningPublicKey field to given value.

### HasSigningPublicKey

`func (o *Subscriber) HasSigningPublicKey() bool`

HasSigningPublicKey returns a boolean if a field has been set.

### GetEncryptionPublicKey

`func (o *Subscriber) GetEncryptionPublicKey() string`

GetEncryptionPublicKey returns the EncryptionPublicKey field if non-nil, zero value otherwise.

### GetEncryptionPublicKeyOk

`func (o *Subscriber) GetEncryptionPublicKeyOk() (*string, bool)`

GetEncryptionPublicKeyOk returns a tuple with the EncryptionPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionPublicKey

`func (o *Subscriber) SetEncryptionPublicKey(v string)`

SetEncryptionPublicKey sets EncryptionPublicKey field to given value.

### HasEncryptionPublicKey

`func (o *Subscriber) HasEncryptionPublicKey() bool`

HasEncryptionPublicKey returns a boolean if a field has been set.

### GetStatus

`func (o *Subscriber) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Subscriber) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Subscriber) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Subscriber) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreated

`func (o *Subscriber) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Subscriber) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Subscriber) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Subscriber) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *Subscriber) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Subscriber) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Subscriber) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Subscriber) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetExpires

`func (o *Subscriber) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *Subscriber) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *Subscriber) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *Subscriber) HasExpires() bool`

HasExpires returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


