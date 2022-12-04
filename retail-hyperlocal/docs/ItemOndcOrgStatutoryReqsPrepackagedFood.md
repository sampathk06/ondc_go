# ItemOndcOrgStatutoryReqsPrepackagedFood

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IngredientsInfo** | Pointer to **string** | list of ingredients (except single ingredient foods), can be shown as ingredient (with percentage); ingredient (with percentage);..) e.g. \&quot;Puffed Rice (40%); Split Green Gram (20%); Ground Nuts (20%);..\&quot; | [optional] 
**NutritionalInfo** | Pointer to **string** | nutritional info (can be shown as nutritional info (with unit, per standard unit, per serving);..) e.g. \&quot;Energy(KCal) - (per 100kg) 420, (per serving 50g) 250; Protein(g) - (per 100kg) 12, (per serving 50g)6;..\&quot; | [optional] 
**AdditivesInfo** | Pointer to **string** | food additives together with specific name or recognized International Numbering System (can be shown as additive1-name or number;additive2-name or number;..) | [optional] 
**ManufacturerOrPackerName** | Pointer to **string** | name of manufacturer or packer (for non-retail containers) | [optional] 
**ManufacturerOrPackerAddress** | Pointer to **string** | address of manufacturer or packer (for non-retail containers) | [optional] 
**BrandOwnerName** | Pointer to **string** | name of brand owner | [optional] 
**BrandOwnerAddress** | Pointer to **string** | address of brand owner | [optional] 
**BrandOwnerFSSAILogo** | Pointer to **string** | FSSAI logo of brand owner (url based image e.g. uri:http://path/to/image) | [optional] 
**BrandOwnerFSSAILicenseNo** | Pointer to **string** | FSSAI license no of brand owner | [optional] 
**OtherFSSAILicenseNo** | Pointer to **string** | FSSAI license no of manufacturer or marketer or packer or bottler if different from brand owner | [optional] 
**NetQuantity** | Pointer to **string** | net quantity | [optional] 
**ImporterName** | Pointer to **string** | name of importer | [optional] 
**ImporterAddress** | Pointer to **string** | address of importer | [optional] 
**ImporterFSSAILogo** | Pointer to **string** | FSSAI logo of importer (url based image e.g. uri:http://path/to/image) | [optional] 
**ImporterFSSAILicenseNo** | Pointer to **string** | FSSAI license no of importer | [optional] 
**ImportedProductCountryOfOrigin** | Pointer to **string** | country of origin for imported products (ISO 3166 Alpha-3 code format) | [optional] 
**OtherImporterName** | Pointer to **string** | name of importer for product manufactured outside but packaged or bottled in India | [optional] 
**OtherImporterAddress** | Pointer to **string** | address of importer for product manufactured outside but packaged or bottled in India | [optional] 
**OtherPremises** | Pointer to **string** | premises where product manufactured outside are packaged or bottled in India | [optional] 
**OtherImporterCountryOfOrigin** | Pointer to **string** | country of origin for product manufactured outside but packaged or bottled in India (ISO 3166 Alpha-3 code format) | [optional] 

## Methods

### NewItemOndcOrgStatutoryReqsPrepackagedFood

`func NewItemOndcOrgStatutoryReqsPrepackagedFood() *ItemOndcOrgStatutoryReqsPrepackagedFood`

NewItemOndcOrgStatutoryReqsPrepackagedFood instantiates a new ItemOndcOrgStatutoryReqsPrepackagedFood object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemOndcOrgStatutoryReqsPrepackagedFoodWithDefaults

`func NewItemOndcOrgStatutoryReqsPrepackagedFoodWithDefaults() *ItemOndcOrgStatutoryReqsPrepackagedFood`

NewItemOndcOrgStatutoryReqsPrepackagedFoodWithDefaults instantiates a new ItemOndcOrgStatutoryReqsPrepackagedFood object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngredientsInfo

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetIngredientsInfo() string`

GetIngredientsInfo returns the IngredientsInfo field if non-nil, zero value otherwise.

### GetIngredientsInfoOk

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetIngredientsInfoOk() (*string, bool)`

GetIngredientsInfoOk returns a tuple with the IngredientsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngredientsInfo

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetIngredientsInfo(v string)`

SetIngredientsInfo sets IngredientsInfo field to given value.

### HasIngredientsInfo

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasIngredientsInfo() bool`

HasIngredientsInfo returns a boolean if a field has been set.

### GetNutritionalInfo

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetNutritionalInfo() string`

GetNutritionalInfo returns the NutritionalInfo field if non-nil, zero value otherwise.

### GetNutritionalInfoOk

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetNutritionalInfoOk() (*string, bool)`

GetNutritionalInfoOk returns a tuple with the NutritionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNutritionalInfo

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetNutritionalInfo(v string)`

SetNutritionalInfo sets NutritionalInfo field to given value.

### HasNutritionalInfo

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasNutritionalInfo() bool`

HasNutritionalInfo returns a boolean if a field has been set.

### GetAdditivesInfo

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetAdditivesInfo() string`

GetAdditivesInfo returns the AdditivesInfo field if non-nil, zero value otherwise.

### GetAdditivesInfoOk

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetAdditivesInfoOk() (*string, bool)`

GetAdditivesInfoOk returns a tuple with the AdditivesInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditivesInfo

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetAdditivesInfo(v string)`

SetAdditivesInfo sets AdditivesInfo field to given value.

### HasAdditivesInfo

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasAdditivesInfo() bool`

HasAdditivesInfo returns a boolean if a field has been set.

### GetManufacturerOrPackerName

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetManufacturerOrPackerName() string`

GetManufacturerOrPackerName returns the ManufacturerOrPackerName field if non-nil, zero value otherwise.

### GetManufacturerOrPackerNameOk

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetManufacturerOrPackerNameOk() (*string, bool)`

GetManufacturerOrPackerNameOk returns a tuple with the ManufacturerOrPackerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerOrPackerName

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetManufacturerOrPackerName(v string)`

SetManufacturerOrPackerName sets ManufacturerOrPackerName field to given value.

### HasManufacturerOrPackerName

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasManufacturerOrPackerName() bool`

HasManufacturerOrPackerName returns a boolean if a field has been set.

### GetManufacturerOrPackerAddress

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetManufacturerOrPackerAddress() string`

GetManufacturerOrPackerAddress returns the ManufacturerOrPackerAddress field if non-nil, zero value otherwise.

### GetManufacturerOrPackerAddressOk

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetManufacturerOrPackerAddressOk() (*string, bool)`

GetManufacturerOrPackerAddressOk returns a tuple with the ManufacturerOrPackerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerOrPackerAddress

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetManufacturerOrPackerAddress(v string)`

SetManufacturerOrPackerAddress sets ManufacturerOrPackerAddress field to given value.

### HasManufacturerOrPackerAddress

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasManufacturerOrPackerAddress() bool`

HasManufacturerOrPackerAddress returns a boolean if a field has been set.

### GetBrandOwnerName

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetBrandOwnerName() string`

GetBrandOwnerName returns the BrandOwnerName field if non-nil, zero value otherwise.

### GetBrandOwnerNameOk

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetBrandOwnerNameOk() (*string, bool)`

GetBrandOwnerNameOk returns a tuple with the BrandOwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandOwnerName

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetBrandOwnerName(v string)`

SetBrandOwnerName sets BrandOwnerName field to given value.

### HasBrandOwnerName

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasBrandOwnerName() bool`

HasBrandOwnerName returns a boolean if a field has been set.

### GetBrandOwnerAddress

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetBrandOwnerAddress() string`

GetBrandOwnerAddress returns the BrandOwnerAddress field if non-nil, zero value otherwise.

### GetBrandOwnerAddressOk

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetBrandOwnerAddressOk() (*string, bool)`

GetBrandOwnerAddressOk returns a tuple with the BrandOwnerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandOwnerAddress

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetBrandOwnerAddress(v string)`

SetBrandOwnerAddress sets BrandOwnerAddress field to given value.

### HasBrandOwnerAddress

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasBrandOwnerAddress() bool`

HasBrandOwnerAddress returns a boolean if a field has been set.

### GetBrandOwnerFSSAILogo

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetBrandOwnerFSSAILogo() string`

GetBrandOwnerFSSAILogo returns the BrandOwnerFSSAILogo field if non-nil, zero value otherwise.

### GetBrandOwnerFSSAILogoOk

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetBrandOwnerFSSAILogoOk() (*string, bool)`

GetBrandOwnerFSSAILogoOk returns a tuple with the BrandOwnerFSSAILogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandOwnerFSSAILogo

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetBrandOwnerFSSAILogo(v string)`

SetBrandOwnerFSSAILogo sets BrandOwnerFSSAILogo field to given value.

### HasBrandOwnerFSSAILogo

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasBrandOwnerFSSAILogo() bool`

HasBrandOwnerFSSAILogo returns a boolean if a field has been set.

### GetBrandOwnerFSSAILicenseNo

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetBrandOwnerFSSAILicenseNo() string`

GetBrandOwnerFSSAILicenseNo returns the BrandOwnerFSSAILicenseNo field if non-nil, zero value otherwise.

### GetBrandOwnerFSSAILicenseNoOk

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetBrandOwnerFSSAILicenseNoOk() (*string, bool)`

GetBrandOwnerFSSAILicenseNoOk returns a tuple with the BrandOwnerFSSAILicenseNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandOwnerFSSAILicenseNo

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetBrandOwnerFSSAILicenseNo(v string)`

SetBrandOwnerFSSAILicenseNo sets BrandOwnerFSSAILicenseNo field to given value.

### HasBrandOwnerFSSAILicenseNo

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasBrandOwnerFSSAILicenseNo() bool`

HasBrandOwnerFSSAILicenseNo returns a boolean if a field has been set.

### GetOtherFSSAILicenseNo

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetOtherFSSAILicenseNo() string`

GetOtherFSSAILicenseNo returns the OtherFSSAILicenseNo field if non-nil, zero value otherwise.

### GetOtherFSSAILicenseNoOk

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetOtherFSSAILicenseNoOk() (*string, bool)`

GetOtherFSSAILicenseNoOk returns a tuple with the OtherFSSAILicenseNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherFSSAILicenseNo

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetOtherFSSAILicenseNo(v string)`

SetOtherFSSAILicenseNo sets OtherFSSAILicenseNo field to given value.

### HasOtherFSSAILicenseNo

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasOtherFSSAILicenseNo() bool`

HasOtherFSSAILicenseNo returns a boolean if a field has been set.

### GetNetQuantity

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetNetQuantity() string`

GetNetQuantity returns the NetQuantity field if non-nil, zero value otherwise.

### GetNetQuantityOk

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetNetQuantityOk() (*string, bool)`

GetNetQuantityOk returns a tuple with the NetQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetQuantity

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetNetQuantity(v string)`

SetNetQuantity sets NetQuantity field to given value.

### HasNetQuantity

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasNetQuantity() bool`

HasNetQuantity returns a boolean if a field has been set.

### GetImporterName

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetImporterName() string`

GetImporterName returns the ImporterName field if non-nil, zero value otherwise.

### GetImporterNameOk

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetImporterNameOk() (*string, bool)`

GetImporterNameOk returns a tuple with the ImporterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImporterName

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetImporterName(v string)`

SetImporterName sets ImporterName field to given value.

### HasImporterName

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasImporterName() bool`

HasImporterName returns a boolean if a field has been set.

### GetImporterAddress

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetImporterAddress() string`

GetImporterAddress returns the ImporterAddress field if non-nil, zero value otherwise.

### GetImporterAddressOk

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetImporterAddressOk() (*string, bool)`

GetImporterAddressOk returns a tuple with the ImporterAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImporterAddress

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetImporterAddress(v string)`

SetImporterAddress sets ImporterAddress field to given value.

### HasImporterAddress

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasImporterAddress() bool`

HasImporterAddress returns a boolean if a field has been set.

### GetImporterFSSAILogo

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetImporterFSSAILogo() string`

GetImporterFSSAILogo returns the ImporterFSSAILogo field if non-nil, zero value otherwise.

### GetImporterFSSAILogoOk

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetImporterFSSAILogoOk() (*string, bool)`

GetImporterFSSAILogoOk returns a tuple with the ImporterFSSAILogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImporterFSSAILogo

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetImporterFSSAILogo(v string)`

SetImporterFSSAILogo sets ImporterFSSAILogo field to given value.

### HasImporterFSSAILogo

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasImporterFSSAILogo() bool`

HasImporterFSSAILogo returns a boolean if a field has been set.

### GetImporterFSSAILicenseNo

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetImporterFSSAILicenseNo() string`

GetImporterFSSAILicenseNo returns the ImporterFSSAILicenseNo field if non-nil, zero value otherwise.

### GetImporterFSSAILicenseNoOk

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetImporterFSSAILicenseNoOk() (*string, bool)`

GetImporterFSSAILicenseNoOk returns a tuple with the ImporterFSSAILicenseNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImporterFSSAILicenseNo

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetImporterFSSAILicenseNo(v string)`

SetImporterFSSAILicenseNo sets ImporterFSSAILicenseNo field to given value.

### HasImporterFSSAILicenseNo

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasImporterFSSAILicenseNo() bool`

HasImporterFSSAILicenseNo returns a boolean if a field has been set.

### GetImportedProductCountryOfOrigin

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetImportedProductCountryOfOrigin() string`

GetImportedProductCountryOfOrigin returns the ImportedProductCountryOfOrigin field if non-nil, zero value otherwise.

### GetImportedProductCountryOfOriginOk

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetImportedProductCountryOfOriginOk() (*string, bool)`

GetImportedProductCountryOfOriginOk returns a tuple with the ImportedProductCountryOfOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedProductCountryOfOrigin

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetImportedProductCountryOfOrigin(v string)`

SetImportedProductCountryOfOrigin sets ImportedProductCountryOfOrigin field to given value.

### HasImportedProductCountryOfOrigin

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasImportedProductCountryOfOrigin() bool`

HasImportedProductCountryOfOrigin returns a boolean if a field has been set.

### GetOtherImporterName

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetOtherImporterName() string`

GetOtherImporterName returns the OtherImporterName field if non-nil, zero value otherwise.

### GetOtherImporterNameOk

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetOtherImporterNameOk() (*string, bool)`

GetOtherImporterNameOk returns a tuple with the OtherImporterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherImporterName

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetOtherImporterName(v string)`

SetOtherImporterName sets OtherImporterName field to given value.

### HasOtherImporterName

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasOtherImporterName() bool`

HasOtherImporterName returns a boolean if a field has been set.

### GetOtherImporterAddress

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetOtherImporterAddress() string`

GetOtherImporterAddress returns the OtherImporterAddress field if non-nil, zero value otherwise.

### GetOtherImporterAddressOk

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetOtherImporterAddressOk() (*string, bool)`

GetOtherImporterAddressOk returns a tuple with the OtherImporterAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherImporterAddress

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetOtherImporterAddress(v string)`

SetOtherImporterAddress sets OtherImporterAddress field to given value.

### HasOtherImporterAddress

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasOtherImporterAddress() bool`

HasOtherImporterAddress returns a boolean if a field has been set.

### GetOtherPremises

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetOtherPremises() string`

GetOtherPremises returns the OtherPremises field if non-nil, zero value otherwise.

### GetOtherPremisesOk

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetOtherPremisesOk() (*string, bool)`

GetOtherPremisesOk returns a tuple with the OtherPremises field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherPremises

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetOtherPremises(v string)`

SetOtherPremises sets OtherPremises field to given value.

### HasOtherPremises

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasOtherPremises() bool`

HasOtherPremises returns a boolean if a field has been set.

### GetOtherImporterCountryOfOrigin

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetOtherImporterCountryOfOrigin() string`

GetOtherImporterCountryOfOrigin returns the OtherImporterCountryOfOrigin field if non-nil, zero value otherwise.

### GetOtherImporterCountryOfOriginOk

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetOtherImporterCountryOfOriginOk() (*string, bool)`

GetOtherImporterCountryOfOriginOk returns a tuple with the OtherImporterCountryOfOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherImporterCountryOfOrigin

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetOtherImporterCountryOfOrigin(v string)`

SetOtherImporterCountryOfOrigin sets OtherImporterCountryOfOrigin field to given value.

### HasOtherImporterCountryOfOrigin

`func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasOtherImporterCountryOfOrigin() bool`

HasOtherImporterCountryOfOrigin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


