/*
ONDC Protocol API for retail (grocery, f&b)

ONDC Protocol API specification, which includes statutory requirements for packaged commodities and pre-packaged food This is an adaptation of Beckn Core version 0.9.3

API version: 1.0.27
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ItemOndcOrgStatutoryReqsPrepackagedFood mandatory for category_id \"Packaged food\" - required attributes include the following - nutritional_info, additives_info, net_quantity; other attributes are required on case-by-case basis
type ItemOndcOrgStatutoryReqsPrepackagedFood struct {
	// list of ingredients (except single ingredient foods), can be shown as ingredient (with percentage); ingredient (with percentage);..) e.g. \"Puffed Rice (40%); Split Green Gram (20%); Ground Nuts (20%);..\"
	IngredientsInfo *string `json:"ingredients_info,omitempty"`
	// nutritional info (can be shown as nutritional info (with unit, per standard unit, per serving);..) e.g. \"Energy(KCal) - (per 100kg) 420, (per serving 50g) 250; Protein(g) - (per 100kg) 12, (per serving 50g)6;..\"
	NutritionalInfo *string `json:"nutritional_info,omitempty"`
	// food additives together with specific name or recognized International Numbering System (can be shown as additive1-name or number;additive2-name or number;..)
	AdditivesInfo *string `json:"additives_info,omitempty"`
	// name of manufacturer or packer (for non-retail containers)
	ManufacturerOrPackerName *string `json:"manufacturer_or_packer_name,omitempty"`
	// address of manufacturer or packer (for non-retail containers)
	ManufacturerOrPackerAddress *string `json:"manufacturer_or_packer_address,omitempty"`
	// name of brand owner
	BrandOwnerName *string `json:"brand_owner_name,omitempty"`
	// address of brand owner
	BrandOwnerAddress *string `json:"brand_owner_address,omitempty"`
	// FSSAI logo of brand owner (url based image e.g. uri:http://path/to/image)
	BrandOwnerFSSAILogo *string `json:"brand_owner_FSSAI_logo,omitempty"`
	// FSSAI license no of brand owner
	BrandOwnerFSSAILicenseNo *string `json:"brand_owner_FSSAI_license_no,omitempty"`
	// FSSAI license no of manufacturer or marketer or packer or bottler if different from brand owner
	OtherFSSAILicenseNo *string `json:"other_FSSAI_license_no,omitempty"`
	// net quantity
	NetQuantity *string `json:"net_quantity,omitempty"`
	// name of importer
	ImporterName *string `json:"importer_name,omitempty"`
	// address of importer
	ImporterAddress *string `json:"importer_address,omitempty"`
	// FSSAI logo of importer (url based image e.g. uri:http://path/to/image)
	ImporterFSSAILogo *string `json:"importer_FSSAI_logo,omitempty"`
	// FSSAI license no of importer
	ImporterFSSAILicenseNo *string `json:"importer_FSSAI_license_no,omitempty"`
	// country of origin for imported products (ISO 3166 Alpha-3 code format)
	ImportedProductCountryOfOrigin *string `json:"imported_product_country_of_origin,omitempty"`
	// name of importer for product manufactured outside but packaged or bottled in India
	OtherImporterName *string `json:"other_importer_name,omitempty"`
	// address of importer for product manufactured outside but packaged or bottled in India
	OtherImporterAddress *string `json:"other_importer_address,omitempty"`
	// premises where product manufactured outside are packaged or bottled in India
	OtherPremises *string `json:"other_premises,omitempty"`
	// country of origin for product manufactured outside but packaged or bottled in India (ISO 3166 Alpha-3 code format)
	OtherImporterCountryOfOrigin *string `json:"other_importer_country_of_origin,omitempty"`
}

// NewItemOndcOrgStatutoryReqsPrepackagedFood instantiates a new ItemOndcOrgStatutoryReqsPrepackagedFood object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemOndcOrgStatutoryReqsPrepackagedFood() *ItemOndcOrgStatutoryReqsPrepackagedFood {
	this := ItemOndcOrgStatutoryReqsPrepackagedFood{}
	return &this
}

// NewItemOndcOrgStatutoryReqsPrepackagedFoodWithDefaults instantiates a new ItemOndcOrgStatutoryReqsPrepackagedFood object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemOndcOrgStatutoryReqsPrepackagedFoodWithDefaults() *ItemOndcOrgStatutoryReqsPrepackagedFood {
	this := ItemOndcOrgStatutoryReqsPrepackagedFood{}
	return &this
}

// GetIngredientsInfo returns the IngredientsInfo field value if set, zero value otherwise.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetIngredientsInfo() string {
	if o == nil || isNil(o.IngredientsInfo) {
		var ret string
		return ret
	}
	return *o.IngredientsInfo
}

// GetIngredientsInfoOk returns a tuple with the IngredientsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetIngredientsInfoOk() (*string, bool) {
	if o == nil || isNil(o.IngredientsInfo) {
    return nil, false
	}
	return o.IngredientsInfo, true
}

// HasIngredientsInfo returns a boolean if a field has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasIngredientsInfo() bool {
	if o != nil && !isNil(o.IngredientsInfo) {
		return true
	}

	return false
}

// SetIngredientsInfo gets a reference to the given string and assigns it to the IngredientsInfo field.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetIngredientsInfo(v string) {
	o.IngredientsInfo = &v
}

// GetNutritionalInfo returns the NutritionalInfo field value if set, zero value otherwise.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetNutritionalInfo() string {
	if o == nil || isNil(o.NutritionalInfo) {
		var ret string
		return ret
	}
	return *o.NutritionalInfo
}

// GetNutritionalInfoOk returns a tuple with the NutritionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetNutritionalInfoOk() (*string, bool) {
	if o == nil || isNil(o.NutritionalInfo) {
    return nil, false
	}
	return o.NutritionalInfo, true
}

// HasNutritionalInfo returns a boolean if a field has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasNutritionalInfo() bool {
	if o != nil && !isNil(o.NutritionalInfo) {
		return true
	}

	return false
}

// SetNutritionalInfo gets a reference to the given string and assigns it to the NutritionalInfo field.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetNutritionalInfo(v string) {
	o.NutritionalInfo = &v
}

// GetAdditivesInfo returns the AdditivesInfo field value if set, zero value otherwise.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetAdditivesInfo() string {
	if o == nil || isNil(o.AdditivesInfo) {
		var ret string
		return ret
	}
	return *o.AdditivesInfo
}

// GetAdditivesInfoOk returns a tuple with the AdditivesInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetAdditivesInfoOk() (*string, bool) {
	if o == nil || isNil(o.AdditivesInfo) {
    return nil, false
	}
	return o.AdditivesInfo, true
}

// HasAdditivesInfo returns a boolean if a field has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasAdditivesInfo() bool {
	if o != nil && !isNil(o.AdditivesInfo) {
		return true
	}

	return false
}

// SetAdditivesInfo gets a reference to the given string and assigns it to the AdditivesInfo field.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetAdditivesInfo(v string) {
	o.AdditivesInfo = &v
}

// GetManufacturerOrPackerName returns the ManufacturerOrPackerName field value if set, zero value otherwise.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetManufacturerOrPackerName() string {
	if o == nil || isNil(o.ManufacturerOrPackerName) {
		var ret string
		return ret
	}
	return *o.ManufacturerOrPackerName
}

// GetManufacturerOrPackerNameOk returns a tuple with the ManufacturerOrPackerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetManufacturerOrPackerNameOk() (*string, bool) {
	if o == nil || isNil(o.ManufacturerOrPackerName) {
    return nil, false
	}
	return o.ManufacturerOrPackerName, true
}

// HasManufacturerOrPackerName returns a boolean if a field has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasManufacturerOrPackerName() bool {
	if o != nil && !isNil(o.ManufacturerOrPackerName) {
		return true
	}

	return false
}

// SetManufacturerOrPackerName gets a reference to the given string and assigns it to the ManufacturerOrPackerName field.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetManufacturerOrPackerName(v string) {
	o.ManufacturerOrPackerName = &v
}

// GetManufacturerOrPackerAddress returns the ManufacturerOrPackerAddress field value if set, zero value otherwise.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetManufacturerOrPackerAddress() string {
	if o == nil || isNil(o.ManufacturerOrPackerAddress) {
		var ret string
		return ret
	}
	return *o.ManufacturerOrPackerAddress
}

// GetManufacturerOrPackerAddressOk returns a tuple with the ManufacturerOrPackerAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetManufacturerOrPackerAddressOk() (*string, bool) {
	if o == nil || isNil(o.ManufacturerOrPackerAddress) {
    return nil, false
	}
	return o.ManufacturerOrPackerAddress, true
}

// HasManufacturerOrPackerAddress returns a boolean if a field has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasManufacturerOrPackerAddress() bool {
	if o != nil && !isNil(o.ManufacturerOrPackerAddress) {
		return true
	}

	return false
}

// SetManufacturerOrPackerAddress gets a reference to the given string and assigns it to the ManufacturerOrPackerAddress field.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetManufacturerOrPackerAddress(v string) {
	o.ManufacturerOrPackerAddress = &v
}

// GetBrandOwnerName returns the BrandOwnerName field value if set, zero value otherwise.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetBrandOwnerName() string {
	if o == nil || isNil(o.BrandOwnerName) {
		var ret string
		return ret
	}
	return *o.BrandOwnerName
}

// GetBrandOwnerNameOk returns a tuple with the BrandOwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetBrandOwnerNameOk() (*string, bool) {
	if o == nil || isNil(o.BrandOwnerName) {
    return nil, false
	}
	return o.BrandOwnerName, true
}

// HasBrandOwnerName returns a boolean if a field has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasBrandOwnerName() bool {
	if o != nil && !isNil(o.BrandOwnerName) {
		return true
	}

	return false
}

// SetBrandOwnerName gets a reference to the given string and assigns it to the BrandOwnerName field.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetBrandOwnerName(v string) {
	o.BrandOwnerName = &v
}

// GetBrandOwnerAddress returns the BrandOwnerAddress field value if set, zero value otherwise.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetBrandOwnerAddress() string {
	if o == nil || isNil(o.BrandOwnerAddress) {
		var ret string
		return ret
	}
	return *o.BrandOwnerAddress
}

// GetBrandOwnerAddressOk returns a tuple with the BrandOwnerAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetBrandOwnerAddressOk() (*string, bool) {
	if o == nil || isNil(o.BrandOwnerAddress) {
    return nil, false
	}
	return o.BrandOwnerAddress, true
}

// HasBrandOwnerAddress returns a boolean if a field has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasBrandOwnerAddress() bool {
	if o != nil && !isNil(o.BrandOwnerAddress) {
		return true
	}

	return false
}

// SetBrandOwnerAddress gets a reference to the given string and assigns it to the BrandOwnerAddress field.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetBrandOwnerAddress(v string) {
	o.BrandOwnerAddress = &v
}

// GetBrandOwnerFSSAILogo returns the BrandOwnerFSSAILogo field value if set, zero value otherwise.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetBrandOwnerFSSAILogo() string {
	if o == nil || isNil(o.BrandOwnerFSSAILogo) {
		var ret string
		return ret
	}
	return *o.BrandOwnerFSSAILogo
}

// GetBrandOwnerFSSAILogoOk returns a tuple with the BrandOwnerFSSAILogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetBrandOwnerFSSAILogoOk() (*string, bool) {
	if o == nil || isNil(o.BrandOwnerFSSAILogo) {
    return nil, false
	}
	return o.BrandOwnerFSSAILogo, true
}

// HasBrandOwnerFSSAILogo returns a boolean if a field has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasBrandOwnerFSSAILogo() bool {
	if o != nil && !isNil(o.BrandOwnerFSSAILogo) {
		return true
	}

	return false
}

// SetBrandOwnerFSSAILogo gets a reference to the given string and assigns it to the BrandOwnerFSSAILogo field.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetBrandOwnerFSSAILogo(v string) {
	o.BrandOwnerFSSAILogo = &v
}

// GetBrandOwnerFSSAILicenseNo returns the BrandOwnerFSSAILicenseNo field value if set, zero value otherwise.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetBrandOwnerFSSAILicenseNo() string {
	if o == nil || isNil(o.BrandOwnerFSSAILicenseNo) {
		var ret string
		return ret
	}
	return *o.BrandOwnerFSSAILicenseNo
}

// GetBrandOwnerFSSAILicenseNoOk returns a tuple with the BrandOwnerFSSAILicenseNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetBrandOwnerFSSAILicenseNoOk() (*string, bool) {
	if o == nil || isNil(o.BrandOwnerFSSAILicenseNo) {
    return nil, false
	}
	return o.BrandOwnerFSSAILicenseNo, true
}

// HasBrandOwnerFSSAILicenseNo returns a boolean if a field has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasBrandOwnerFSSAILicenseNo() bool {
	if o != nil && !isNil(o.BrandOwnerFSSAILicenseNo) {
		return true
	}

	return false
}

// SetBrandOwnerFSSAILicenseNo gets a reference to the given string and assigns it to the BrandOwnerFSSAILicenseNo field.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetBrandOwnerFSSAILicenseNo(v string) {
	o.BrandOwnerFSSAILicenseNo = &v
}

// GetOtherFSSAILicenseNo returns the OtherFSSAILicenseNo field value if set, zero value otherwise.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetOtherFSSAILicenseNo() string {
	if o == nil || isNil(o.OtherFSSAILicenseNo) {
		var ret string
		return ret
	}
	return *o.OtherFSSAILicenseNo
}

// GetOtherFSSAILicenseNoOk returns a tuple with the OtherFSSAILicenseNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetOtherFSSAILicenseNoOk() (*string, bool) {
	if o == nil || isNil(o.OtherFSSAILicenseNo) {
    return nil, false
	}
	return o.OtherFSSAILicenseNo, true
}

// HasOtherFSSAILicenseNo returns a boolean if a field has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasOtherFSSAILicenseNo() bool {
	if o != nil && !isNil(o.OtherFSSAILicenseNo) {
		return true
	}

	return false
}

// SetOtherFSSAILicenseNo gets a reference to the given string and assigns it to the OtherFSSAILicenseNo field.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetOtherFSSAILicenseNo(v string) {
	o.OtherFSSAILicenseNo = &v
}

// GetNetQuantity returns the NetQuantity field value if set, zero value otherwise.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetNetQuantity() string {
	if o == nil || isNil(o.NetQuantity) {
		var ret string
		return ret
	}
	return *o.NetQuantity
}

// GetNetQuantityOk returns a tuple with the NetQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetNetQuantityOk() (*string, bool) {
	if o == nil || isNil(o.NetQuantity) {
    return nil, false
	}
	return o.NetQuantity, true
}

// HasNetQuantity returns a boolean if a field has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasNetQuantity() bool {
	if o != nil && !isNil(o.NetQuantity) {
		return true
	}

	return false
}

// SetNetQuantity gets a reference to the given string and assigns it to the NetQuantity field.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetNetQuantity(v string) {
	o.NetQuantity = &v
}

// GetImporterName returns the ImporterName field value if set, zero value otherwise.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetImporterName() string {
	if o == nil || isNil(o.ImporterName) {
		var ret string
		return ret
	}
	return *o.ImporterName
}

// GetImporterNameOk returns a tuple with the ImporterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetImporterNameOk() (*string, bool) {
	if o == nil || isNil(o.ImporterName) {
    return nil, false
	}
	return o.ImporterName, true
}

// HasImporterName returns a boolean if a field has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasImporterName() bool {
	if o != nil && !isNil(o.ImporterName) {
		return true
	}

	return false
}

// SetImporterName gets a reference to the given string and assigns it to the ImporterName field.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetImporterName(v string) {
	o.ImporterName = &v
}

// GetImporterAddress returns the ImporterAddress field value if set, zero value otherwise.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetImporterAddress() string {
	if o == nil || isNil(o.ImporterAddress) {
		var ret string
		return ret
	}
	return *o.ImporterAddress
}

// GetImporterAddressOk returns a tuple with the ImporterAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetImporterAddressOk() (*string, bool) {
	if o == nil || isNil(o.ImporterAddress) {
    return nil, false
	}
	return o.ImporterAddress, true
}

// HasImporterAddress returns a boolean if a field has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasImporterAddress() bool {
	if o != nil && !isNil(o.ImporterAddress) {
		return true
	}

	return false
}

// SetImporterAddress gets a reference to the given string and assigns it to the ImporterAddress field.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetImporterAddress(v string) {
	o.ImporterAddress = &v
}

// GetImporterFSSAILogo returns the ImporterFSSAILogo field value if set, zero value otherwise.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetImporterFSSAILogo() string {
	if o == nil || isNil(o.ImporterFSSAILogo) {
		var ret string
		return ret
	}
	return *o.ImporterFSSAILogo
}

// GetImporterFSSAILogoOk returns a tuple with the ImporterFSSAILogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetImporterFSSAILogoOk() (*string, bool) {
	if o == nil || isNil(o.ImporterFSSAILogo) {
    return nil, false
	}
	return o.ImporterFSSAILogo, true
}

// HasImporterFSSAILogo returns a boolean if a field has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasImporterFSSAILogo() bool {
	if o != nil && !isNil(o.ImporterFSSAILogo) {
		return true
	}

	return false
}

// SetImporterFSSAILogo gets a reference to the given string and assigns it to the ImporterFSSAILogo field.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetImporterFSSAILogo(v string) {
	o.ImporterFSSAILogo = &v
}

// GetImporterFSSAILicenseNo returns the ImporterFSSAILicenseNo field value if set, zero value otherwise.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetImporterFSSAILicenseNo() string {
	if o == nil || isNil(o.ImporterFSSAILicenseNo) {
		var ret string
		return ret
	}
	return *o.ImporterFSSAILicenseNo
}

// GetImporterFSSAILicenseNoOk returns a tuple with the ImporterFSSAILicenseNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetImporterFSSAILicenseNoOk() (*string, bool) {
	if o == nil || isNil(o.ImporterFSSAILicenseNo) {
    return nil, false
	}
	return o.ImporterFSSAILicenseNo, true
}

// HasImporterFSSAILicenseNo returns a boolean if a field has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasImporterFSSAILicenseNo() bool {
	if o != nil && !isNil(o.ImporterFSSAILicenseNo) {
		return true
	}

	return false
}

// SetImporterFSSAILicenseNo gets a reference to the given string and assigns it to the ImporterFSSAILicenseNo field.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetImporterFSSAILicenseNo(v string) {
	o.ImporterFSSAILicenseNo = &v
}

// GetImportedProductCountryOfOrigin returns the ImportedProductCountryOfOrigin field value if set, zero value otherwise.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetImportedProductCountryOfOrigin() string {
	if o == nil || isNil(o.ImportedProductCountryOfOrigin) {
		var ret string
		return ret
	}
	return *o.ImportedProductCountryOfOrigin
}

// GetImportedProductCountryOfOriginOk returns a tuple with the ImportedProductCountryOfOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetImportedProductCountryOfOriginOk() (*string, bool) {
	if o == nil || isNil(o.ImportedProductCountryOfOrigin) {
    return nil, false
	}
	return o.ImportedProductCountryOfOrigin, true
}

// HasImportedProductCountryOfOrigin returns a boolean if a field has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasImportedProductCountryOfOrigin() bool {
	if o != nil && !isNil(o.ImportedProductCountryOfOrigin) {
		return true
	}

	return false
}

// SetImportedProductCountryOfOrigin gets a reference to the given string and assigns it to the ImportedProductCountryOfOrigin field.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetImportedProductCountryOfOrigin(v string) {
	o.ImportedProductCountryOfOrigin = &v
}

// GetOtherImporterName returns the OtherImporterName field value if set, zero value otherwise.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetOtherImporterName() string {
	if o == nil || isNil(o.OtherImporterName) {
		var ret string
		return ret
	}
	return *o.OtherImporterName
}

// GetOtherImporterNameOk returns a tuple with the OtherImporterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetOtherImporterNameOk() (*string, bool) {
	if o == nil || isNil(o.OtherImporterName) {
    return nil, false
	}
	return o.OtherImporterName, true
}

// HasOtherImporterName returns a boolean if a field has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasOtherImporterName() bool {
	if o != nil && !isNil(o.OtherImporterName) {
		return true
	}

	return false
}

// SetOtherImporterName gets a reference to the given string and assigns it to the OtherImporterName field.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetOtherImporterName(v string) {
	o.OtherImporterName = &v
}

// GetOtherImporterAddress returns the OtherImporterAddress field value if set, zero value otherwise.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetOtherImporterAddress() string {
	if o == nil || isNil(o.OtherImporterAddress) {
		var ret string
		return ret
	}
	return *o.OtherImporterAddress
}

// GetOtherImporterAddressOk returns a tuple with the OtherImporterAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetOtherImporterAddressOk() (*string, bool) {
	if o == nil || isNil(o.OtherImporterAddress) {
    return nil, false
	}
	return o.OtherImporterAddress, true
}

// HasOtherImporterAddress returns a boolean if a field has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasOtherImporterAddress() bool {
	if o != nil && !isNil(o.OtherImporterAddress) {
		return true
	}

	return false
}

// SetOtherImporterAddress gets a reference to the given string and assigns it to the OtherImporterAddress field.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetOtherImporterAddress(v string) {
	o.OtherImporterAddress = &v
}

// GetOtherPremises returns the OtherPremises field value if set, zero value otherwise.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetOtherPremises() string {
	if o == nil || isNil(o.OtherPremises) {
		var ret string
		return ret
	}
	return *o.OtherPremises
}

// GetOtherPremisesOk returns a tuple with the OtherPremises field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetOtherPremisesOk() (*string, bool) {
	if o == nil || isNil(o.OtherPremises) {
    return nil, false
	}
	return o.OtherPremises, true
}

// HasOtherPremises returns a boolean if a field has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasOtherPremises() bool {
	if o != nil && !isNil(o.OtherPremises) {
		return true
	}

	return false
}

// SetOtherPremises gets a reference to the given string and assigns it to the OtherPremises field.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetOtherPremises(v string) {
	o.OtherPremises = &v
}

// GetOtherImporterCountryOfOrigin returns the OtherImporterCountryOfOrigin field value if set, zero value otherwise.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetOtherImporterCountryOfOrigin() string {
	if o == nil || isNil(o.OtherImporterCountryOfOrigin) {
		var ret string
		return ret
	}
	return *o.OtherImporterCountryOfOrigin
}

// GetOtherImporterCountryOfOriginOk returns a tuple with the OtherImporterCountryOfOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) GetOtherImporterCountryOfOriginOk() (*string, bool) {
	if o == nil || isNil(o.OtherImporterCountryOfOrigin) {
    return nil, false
	}
	return o.OtherImporterCountryOfOrigin, true
}

// HasOtherImporterCountryOfOrigin returns a boolean if a field has been set.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) HasOtherImporterCountryOfOrigin() bool {
	if o != nil && !isNil(o.OtherImporterCountryOfOrigin) {
		return true
	}

	return false
}

// SetOtherImporterCountryOfOrigin gets a reference to the given string and assigns it to the OtherImporterCountryOfOrigin field.
func (o *ItemOndcOrgStatutoryReqsPrepackagedFood) SetOtherImporterCountryOfOrigin(v string) {
	o.OtherImporterCountryOfOrigin = &v
}

func (o ItemOndcOrgStatutoryReqsPrepackagedFood) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IngredientsInfo) {
		toSerialize["ingredients_info"] = o.IngredientsInfo
	}
	if !isNil(o.NutritionalInfo) {
		toSerialize["nutritional_info"] = o.NutritionalInfo
	}
	if !isNil(o.AdditivesInfo) {
		toSerialize["additives_info"] = o.AdditivesInfo
	}
	if !isNil(o.ManufacturerOrPackerName) {
		toSerialize["manufacturer_or_packer_name"] = o.ManufacturerOrPackerName
	}
	if !isNil(o.ManufacturerOrPackerAddress) {
		toSerialize["manufacturer_or_packer_address"] = o.ManufacturerOrPackerAddress
	}
	if !isNil(o.BrandOwnerName) {
		toSerialize["brand_owner_name"] = o.BrandOwnerName
	}
	if !isNil(o.BrandOwnerAddress) {
		toSerialize["brand_owner_address"] = o.BrandOwnerAddress
	}
	if !isNil(o.BrandOwnerFSSAILogo) {
		toSerialize["brand_owner_FSSAI_logo"] = o.BrandOwnerFSSAILogo
	}
	if !isNil(o.BrandOwnerFSSAILicenseNo) {
		toSerialize["brand_owner_FSSAI_license_no"] = o.BrandOwnerFSSAILicenseNo
	}
	if !isNil(o.OtherFSSAILicenseNo) {
		toSerialize["other_FSSAI_license_no"] = o.OtherFSSAILicenseNo
	}
	if !isNil(o.NetQuantity) {
		toSerialize["net_quantity"] = o.NetQuantity
	}
	if !isNil(o.ImporterName) {
		toSerialize["importer_name"] = o.ImporterName
	}
	if !isNil(o.ImporterAddress) {
		toSerialize["importer_address"] = o.ImporterAddress
	}
	if !isNil(o.ImporterFSSAILogo) {
		toSerialize["importer_FSSAI_logo"] = o.ImporterFSSAILogo
	}
	if !isNil(o.ImporterFSSAILicenseNo) {
		toSerialize["importer_FSSAI_license_no"] = o.ImporterFSSAILicenseNo
	}
	if !isNil(o.ImportedProductCountryOfOrigin) {
		toSerialize["imported_product_country_of_origin"] = o.ImportedProductCountryOfOrigin
	}
	if !isNil(o.OtherImporterName) {
		toSerialize["other_importer_name"] = o.OtherImporterName
	}
	if !isNil(o.OtherImporterAddress) {
		toSerialize["other_importer_address"] = o.OtherImporterAddress
	}
	if !isNil(o.OtherPremises) {
		toSerialize["other_premises"] = o.OtherPremises
	}
	if !isNil(o.OtherImporterCountryOfOrigin) {
		toSerialize["other_importer_country_of_origin"] = o.OtherImporterCountryOfOrigin
	}
	return json.Marshal(toSerialize)
}

type NullableItemOndcOrgStatutoryReqsPrepackagedFood struct {
	value *ItemOndcOrgStatutoryReqsPrepackagedFood
	isSet bool
}

func (v NullableItemOndcOrgStatutoryReqsPrepackagedFood) Get() *ItemOndcOrgStatutoryReqsPrepackagedFood {
	return v.value
}

func (v *NullableItemOndcOrgStatutoryReqsPrepackagedFood) Set(val *ItemOndcOrgStatutoryReqsPrepackagedFood) {
	v.value = val
	v.isSet = true
}

func (v NullableItemOndcOrgStatutoryReqsPrepackagedFood) IsSet() bool {
	return v.isSet
}

func (v *NullableItemOndcOrgStatutoryReqsPrepackagedFood) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemOndcOrgStatutoryReqsPrepackagedFood(val *ItemOndcOrgStatutoryReqsPrepackagedFood) *NullableItemOndcOrgStatutoryReqsPrepackagedFood {
	return &NullableItemOndcOrgStatutoryReqsPrepackagedFood{value: val, isSet: true}
}

func (v NullableItemOndcOrgStatutoryReqsPrepackagedFood) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemOndcOrgStatutoryReqsPrepackagedFood) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

