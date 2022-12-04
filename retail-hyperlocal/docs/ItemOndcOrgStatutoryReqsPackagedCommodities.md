# ItemOndcOrgStatutoryReqsPackagedCommodities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManufacturerOrPackerName** | Pointer to **string** | name of manufacturer or packer (in case manufacturer is not the packer) or name of importer for imported goods | [optional] 
**ManufacturerOrPackerAddress** | Pointer to **string** | address of manufacturer or packer (in case manufacturer is not the packer) or address of importer for imported goods | [optional] 
**CommonOrGenericNameOfCommodity** | Pointer to **string** | common or generic name of commodity | [optional] 
**MultipleProductsNameNumberOrQty** | Pointer to **string** | for packages with multiple products, the name and number of quantity of each (can be shown as \&quot;name1-number_or_quantity; name2-number_or_quantity..\&quot;) | [optional] 
**NetQuantityOrMeasureOfCommodityInPkg** | Pointer to **string** | net quantity of commodity in terms of standard unit of weight or measure of commodity contained in package | [optional] 
**MonthYearOfManufacturePackingImport** | Pointer to **string** | month and year of manufacture or packing or import | [optional] 
**ImportedProductCountryOfOrigin** | Pointer to **string** | country of origin for imported products (ISO 3166 Alpha-3 code format) | [optional] 

## Methods

### NewItemOndcOrgStatutoryReqsPackagedCommodities

`func NewItemOndcOrgStatutoryReqsPackagedCommodities() *ItemOndcOrgStatutoryReqsPackagedCommodities`

NewItemOndcOrgStatutoryReqsPackagedCommodities instantiates a new ItemOndcOrgStatutoryReqsPackagedCommodities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemOndcOrgStatutoryReqsPackagedCommoditiesWithDefaults

`func NewItemOndcOrgStatutoryReqsPackagedCommoditiesWithDefaults() *ItemOndcOrgStatutoryReqsPackagedCommodities`

NewItemOndcOrgStatutoryReqsPackagedCommoditiesWithDefaults instantiates a new ItemOndcOrgStatutoryReqsPackagedCommodities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManufacturerOrPackerName

`func (o *ItemOndcOrgStatutoryReqsPackagedCommodities) GetManufacturerOrPackerName() string`

GetManufacturerOrPackerName returns the ManufacturerOrPackerName field if non-nil, zero value otherwise.

### GetManufacturerOrPackerNameOk

`func (o *ItemOndcOrgStatutoryReqsPackagedCommodities) GetManufacturerOrPackerNameOk() (*string, bool)`

GetManufacturerOrPackerNameOk returns a tuple with the ManufacturerOrPackerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerOrPackerName

`func (o *ItemOndcOrgStatutoryReqsPackagedCommodities) SetManufacturerOrPackerName(v string)`

SetManufacturerOrPackerName sets ManufacturerOrPackerName field to given value.

### HasManufacturerOrPackerName

`func (o *ItemOndcOrgStatutoryReqsPackagedCommodities) HasManufacturerOrPackerName() bool`

HasManufacturerOrPackerName returns a boolean if a field has been set.

### GetManufacturerOrPackerAddress

`func (o *ItemOndcOrgStatutoryReqsPackagedCommodities) GetManufacturerOrPackerAddress() string`

GetManufacturerOrPackerAddress returns the ManufacturerOrPackerAddress field if non-nil, zero value otherwise.

### GetManufacturerOrPackerAddressOk

`func (o *ItemOndcOrgStatutoryReqsPackagedCommodities) GetManufacturerOrPackerAddressOk() (*string, bool)`

GetManufacturerOrPackerAddressOk returns a tuple with the ManufacturerOrPackerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerOrPackerAddress

`func (o *ItemOndcOrgStatutoryReqsPackagedCommodities) SetManufacturerOrPackerAddress(v string)`

SetManufacturerOrPackerAddress sets ManufacturerOrPackerAddress field to given value.

### HasManufacturerOrPackerAddress

`func (o *ItemOndcOrgStatutoryReqsPackagedCommodities) HasManufacturerOrPackerAddress() bool`

HasManufacturerOrPackerAddress returns a boolean if a field has been set.

### GetCommonOrGenericNameOfCommodity

`func (o *ItemOndcOrgStatutoryReqsPackagedCommodities) GetCommonOrGenericNameOfCommodity() string`

GetCommonOrGenericNameOfCommodity returns the CommonOrGenericNameOfCommodity field if non-nil, zero value otherwise.

### GetCommonOrGenericNameOfCommodityOk

`func (o *ItemOndcOrgStatutoryReqsPackagedCommodities) GetCommonOrGenericNameOfCommodityOk() (*string, bool)`

GetCommonOrGenericNameOfCommodityOk returns a tuple with the CommonOrGenericNameOfCommodity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonOrGenericNameOfCommodity

`func (o *ItemOndcOrgStatutoryReqsPackagedCommodities) SetCommonOrGenericNameOfCommodity(v string)`

SetCommonOrGenericNameOfCommodity sets CommonOrGenericNameOfCommodity field to given value.

### HasCommonOrGenericNameOfCommodity

`func (o *ItemOndcOrgStatutoryReqsPackagedCommodities) HasCommonOrGenericNameOfCommodity() bool`

HasCommonOrGenericNameOfCommodity returns a boolean if a field has been set.

### GetMultipleProductsNameNumberOrQty

`func (o *ItemOndcOrgStatutoryReqsPackagedCommodities) GetMultipleProductsNameNumberOrQty() string`

GetMultipleProductsNameNumberOrQty returns the MultipleProductsNameNumberOrQty field if non-nil, zero value otherwise.

### GetMultipleProductsNameNumberOrQtyOk

`func (o *ItemOndcOrgStatutoryReqsPackagedCommodities) GetMultipleProductsNameNumberOrQtyOk() (*string, bool)`

GetMultipleProductsNameNumberOrQtyOk returns a tuple with the MultipleProductsNameNumberOrQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleProductsNameNumberOrQty

`func (o *ItemOndcOrgStatutoryReqsPackagedCommodities) SetMultipleProductsNameNumberOrQty(v string)`

SetMultipleProductsNameNumberOrQty sets MultipleProductsNameNumberOrQty field to given value.

### HasMultipleProductsNameNumberOrQty

`func (o *ItemOndcOrgStatutoryReqsPackagedCommodities) HasMultipleProductsNameNumberOrQty() bool`

HasMultipleProductsNameNumberOrQty returns a boolean if a field has been set.

### GetNetQuantityOrMeasureOfCommodityInPkg

`func (o *ItemOndcOrgStatutoryReqsPackagedCommodities) GetNetQuantityOrMeasureOfCommodityInPkg() string`

GetNetQuantityOrMeasureOfCommodityInPkg returns the NetQuantityOrMeasureOfCommodityInPkg field if non-nil, zero value otherwise.

### GetNetQuantityOrMeasureOfCommodityInPkgOk

`func (o *ItemOndcOrgStatutoryReqsPackagedCommodities) GetNetQuantityOrMeasureOfCommodityInPkgOk() (*string, bool)`

GetNetQuantityOrMeasureOfCommodityInPkgOk returns a tuple with the NetQuantityOrMeasureOfCommodityInPkg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetQuantityOrMeasureOfCommodityInPkg

`func (o *ItemOndcOrgStatutoryReqsPackagedCommodities) SetNetQuantityOrMeasureOfCommodityInPkg(v string)`

SetNetQuantityOrMeasureOfCommodityInPkg sets NetQuantityOrMeasureOfCommodityInPkg field to given value.

### HasNetQuantityOrMeasureOfCommodityInPkg

`func (o *ItemOndcOrgStatutoryReqsPackagedCommodities) HasNetQuantityOrMeasureOfCommodityInPkg() bool`

HasNetQuantityOrMeasureOfCommodityInPkg returns a boolean if a field has been set.

### GetMonthYearOfManufacturePackingImport

`func (o *ItemOndcOrgStatutoryReqsPackagedCommodities) GetMonthYearOfManufacturePackingImport() string`

GetMonthYearOfManufacturePackingImport returns the MonthYearOfManufacturePackingImport field if non-nil, zero value otherwise.

### GetMonthYearOfManufacturePackingImportOk

`func (o *ItemOndcOrgStatutoryReqsPackagedCommodities) GetMonthYearOfManufacturePackingImportOk() (*string, bool)`

GetMonthYearOfManufacturePackingImportOk returns a tuple with the MonthYearOfManufacturePackingImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthYearOfManufacturePackingImport

`func (o *ItemOndcOrgStatutoryReqsPackagedCommodities) SetMonthYearOfManufacturePackingImport(v string)`

SetMonthYearOfManufacturePackingImport sets MonthYearOfManufacturePackingImport field to given value.

### HasMonthYearOfManufacturePackingImport

`func (o *ItemOndcOrgStatutoryReqsPackagedCommodities) HasMonthYearOfManufacturePackingImport() bool`

HasMonthYearOfManufacturePackingImport returns a boolean if a field has been set.

### GetImportedProductCountryOfOrigin

`func (o *ItemOndcOrgStatutoryReqsPackagedCommodities) GetImportedProductCountryOfOrigin() string`

GetImportedProductCountryOfOrigin returns the ImportedProductCountryOfOrigin field if non-nil, zero value otherwise.

### GetImportedProductCountryOfOriginOk

`func (o *ItemOndcOrgStatutoryReqsPackagedCommodities) GetImportedProductCountryOfOriginOk() (*string, bool)`

GetImportedProductCountryOfOriginOk returns a tuple with the ImportedProductCountryOfOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedProductCountryOfOrigin

`func (o *ItemOndcOrgStatutoryReqsPackagedCommodities) SetImportedProductCountryOfOrigin(v string)`

SetImportedProductCountryOfOrigin sets ImportedProductCountryOfOrigin field to given value.

### HasImportedProductCountryOfOrigin

`func (o *ItemOndcOrgStatutoryReqsPackagedCommodities) HasImportedProductCountryOfOrigin() bool`

HasImportedProductCountryOfOrigin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


