/*
Anrok API

# API reference  The Anrok API server is accessible at “https://api.anrok.com”.  All requests are HTTP POSTs with JSON in the body.  Authentication is via an HTTP header “Authorization: Bearer {sellerId}/{apiKeyId}/secret.{apiKeySecret}”.  The default rate limit for a seller account is 10 API requests per second. 

API version: 1.0.0
Contact: support@anrok.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the CreateOrUpdateTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrUpdateTransaction{}

// CreateOrUpdateTransaction struct for CreateOrUpdateTransaction
type CreateOrUpdateTransaction struct {
	LineItems []TransactionLineItem `json:"lineItems"`
	// Three letter ISO currency code (case insensitive).
	CurrencyCode string `json:"currencyCode"`
	CustomerAddress CreateEphemeralTransactionCustomerAddress `json:"customerAddress"`
	// The Anrok customer ID used to link transactions for the same customer and to look up tax exemption certificates for a customer. This is typically the billing system's customer ID with a prefix to disambiguate. - If customerId is provided without customerName, that customer object must already exist in Anrok. - If both customerId and customerName are provided, the customer object will be created if it is not already present. - Customer IDs are unique across the entire seller account. 
	CustomerId *string `json:"customerId,omitempty"`
	// The name of the customer
	CustomerName *string `json:"customerName,omitempty"`
	CustomerTaxIds []CustomerTaxId `json:"customerTaxIds,omitempty"`
	// The date that this transaction occurred, for accounting purposes. Accounting date will typically correspond to the invoice date. This is used to determine which tax return the transaction belongs to.
	AccountingDate *string `json:"accountingDate,omitempty"`
	// The time that this transaction occurred, for accounting purposes. If accountingDate is not provided, accountingTime is required to compute an accounting date for the transaction.
	AccountingTime *time.Time `json:"accountingTime,omitempty"`
	// A “tz database” string used to compute an accounting date from the request's accountingTime. The request cannot provide both an accountingDate and an accountingTimeZone. If accountingTime is provided without specifying an accountingTimeZone, the time zone configured on the seller account will be used.
	AccountingTimeZone *string `json:"accountingTimeZone,omitempty"`
	// The date to use for tax calculations. If omitted, Anrok will use the the minimum of the accounting date and two days in the future.
	TaxDate *string `json:"taxDate,omitempty"`
	// The ID of the new transaction. This must be unique across the entire seller account. This is typically the billing system's invoice ID with some prefix to disambiguate different systems.
	Id string `json:"id"`
}

// NewCreateOrUpdateTransaction instantiates a new CreateOrUpdateTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdateTransaction(lineItems []TransactionLineItem, currencyCode string, customerAddress CreateEphemeralTransactionCustomerAddress, id string) *CreateOrUpdateTransaction {
	this := CreateOrUpdateTransaction{}
	this.LineItems = lineItems
	this.CurrencyCode = currencyCode
	this.CustomerAddress = customerAddress
	this.Id = id
	return &this
}

// NewCreateOrUpdateTransactionWithDefaults instantiates a new CreateOrUpdateTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdateTransactionWithDefaults() *CreateOrUpdateTransaction {
	this := CreateOrUpdateTransaction{}
	return &this
}

// GetLineItems returns the LineItems field value
func (o *CreateOrUpdateTransaction) GetLineItems() []TransactionLineItem {
	if o == nil {
		var ret []TransactionLineItem
		return ret
	}

	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateTransaction) GetLineItemsOk() ([]TransactionLineItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.LineItems, true
}

// SetLineItems sets field value
func (o *CreateOrUpdateTransaction) SetLineItems(v []TransactionLineItem) {
	o.LineItems = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *CreateOrUpdateTransaction) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateTransaction) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *CreateOrUpdateTransaction) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetCustomerAddress returns the CustomerAddress field value
func (o *CreateOrUpdateTransaction) GetCustomerAddress() CreateEphemeralTransactionCustomerAddress {
	if o == nil {
		var ret CreateEphemeralTransactionCustomerAddress
		return ret
	}

	return o.CustomerAddress
}

// GetCustomerAddressOk returns a tuple with the CustomerAddress field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateTransaction) GetCustomerAddressOk() (*CreateEphemeralTransactionCustomerAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerAddress, true
}

// SetCustomerAddress sets field value
func (o *CreateOrUpdateTransaction) SetCustomerAddress(v CreateEphemeralTransactionCustomerAddress) {
	o.CustomerAddress = v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *CreateOrUpdateTransaction) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateTransaction) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *CreateOrUpdateTransaction) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *CreateOrUpdateTransaction) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetCustomerName returns the CustomerName field value if set, zero value otherwise.
func (o *CreateOrUpdateTransaction) GetCustomerName() string {
	if o == nil || IsNil(o.CustomerName) {
		var ret string
		return ret
	}
	return *o.CustomerName
}

// GetCustomerNameOk returns a tuple with the CustomerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateTransaction) GetCustomerNameOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerName) {
		return nil, false
	}
	return o.CustomerName, true
}

// HasCustomerName returns a boolean if a field has been set.
func (o *CreateOrUpdateTransaction) HasCustomerName() bool {
	if o != nil && !IsNil(o.CustomerName) {
		return true
	}

	return false
}

// SetCustomerName gets a reference to the given string and assigns it to the CustomerName field.
func (o *CreateOrUpdateTransaction) SetCustomerName(v string) {
	o.CustomerName = &v
}

// GetCustomerTaxIds returns the CustomerTaxIds field value if set, zero value otherwise.
func (o *CreateOrUpdateTransaction) GetCustomerTaxIds() []CustomerTaxId {
	if o == nil || IsNil(o.CustomerTaxIds) {
		var ret []CustomerTaxId
		return ret
	}
	return o.CustomerTaxIds
}

// GetCustomerTaxIdsOk returns a tuple with the CustomerTaxIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateTransaction) GetCustomerTaxIdsOk() ([]CustomerTaxId, bool) {
	if o == nil || IsNil(o.CustomerTaxIds) {
		return nil, false
	}
	return o.CustomerTaxIds, true
}

// HasCustomerTaxIds returns a boolean if a field has been set.
func (o *CreateOrUpdateTransaction) HasCustomerTaxIds() bool {
	if o != nil && !IsNil(o.CustomerTaxIds) {
		return true
	}

	return false
}

// SetCustomerTaxIds gets a reference to the given []CustomerTaxId and assigns it to the CustomerTaxIds field.
func (o *CreateOrUpdateTransaction) SetCustomerTaxIds(v []CustomerTaxId) {
	o.CustomerTaxIds = v
}

// GetAccountingDate returns the AccountingDate field value if set, zero value otherwise.
func (o *CreateOrUpdateTransaction) GetAccountingDate() string {
	if o == nil || IsNil(o.AccountingDate) {
		var ret string
		return ret
	}
	return *o.AccountingDate
}

// GetAccountingDateOk returns a tuple with the AccountingDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateTransaction) GetAccountingDateOk() (*string, bool) {
	if o == nil || IsNil(o.AccountingDate) {
		return nil, false
	}
	return o.AccountingDate, true
}

// HasAccountingDate returns a boolean if a field has been set.
func (o *CreateOrUpdateTransaction) HasAccountingDate() bool {
	if o != nil && !IsNil(o.AccountingDate) {
		return true
	}

	return false
}

// SetAccountingDate gets a reference to the given string and assigns it to the AccountingDate field.
func (o *CreateOrUpdateTransaction) SetAccountingDate(v string) {
	o.AccountingDate = &v
}

// GetAccountingTime returns the AccountingTime field value if set, zero value otherwise.
func (o *CreateOrUpdateTransaction) GetAccountingTime() time.Time {
	if o == nil || IsNil(o.AccountingTime) {
		var ret time.Time
		return ret
	}
	return *o.AccountingTime
}

// GetAccountingTimeOk returns a tuple with the AccountingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateTransaction) GetAccountingTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AccountingTime) {
		return nil, false
	}
	return o.AccountingTime, true
}

// HasAccountingTime returns a boolean if a field has been set.
func (o *CreateOrUpdateTransaction) HasAccountingTime() bool {
	if o != nil && !IsNil(o.AccountingTime) {
		return true
	}

	return false
}

// SetAccountingTime gets a reference to the given time.Time and assigns it to the AccountingTime field.
func (o *CreateOrUpdateTransaction) SetAccountingTime(v time.Time) {
	o.AccountingTime = &v
}

// GetAccountingTimeZone returns the AccountingTimeZone field value if set, zero value otherwise.
func (o *CreateOrUpdateTransaction) GetAccountingTimeZone() string {
	if o == nil || IsNil(o.AccountingTimeZone) {
		var ret string
		return ret
	}
	return *o.AccountingTimeZone
}

// GetAccountingTimeZoneOk returns a tuple with the AccountingTimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateTransaction) GetAccountingTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.AccountingTimeZone) {
		return nil, false
	}
	return o.AccountingTimeZone, true
}

// HasAccountingTimeZone returns a boolean if a field has been set.
func (o *CreateOrUpdateTransaction) HasAccountingTimeZone() bool {
	if o != nil && !IsNil(o.AccountingTimeZone) {
		return true
	}

	return false
}

// SetAccountingTimeZone gets a reference to the given string and assigns it to the AccountingTimeZone field.
func (o *CreateOrUpdateTransaction) SetAccountingTimeZone(v string) {
	o.AccountingTimeZone = &v
}

// GetTaxDate returns the TaxDate field value if set, zero value otherwise.
func (o *CreateOrUpdateTransaction) GetTaxDate() string {
	if o == nil || IsNil(o.TaxDate) {
		var ret string
		return ret
	}
	return *o.TaxDate
}

// GetTaxDateOk returns a tuple with the TaxDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateTransaction) GetTaxDateOk() (*string, bool) {
	if o == nil || IsNil(o.TaxDate) {
		return nil, false
	}
	return o.TaxDate, true
}

// HasTaxDate returns a boolean if a field has been set.
func (o *CreateOrUpdateTransaction) HasTaxDate() bool {
	if o != nil && !IsNil(o.TaxDate) {
		return true
	}

	return false
}

// SetTaxDate gets a reference to the given string and assigns it to the TaxDate field.
func (o *CreateOrUpdateTransaction) SetTaxDate(v string) {
	o.TaxDate = &v
}

// GetId returns the Id field value
func (o *CreateOrUpdateTransaction) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateTransaction) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateOrUpdateTransaction) SetId(v string) {
	o.Id = v
}

func (o CreateOrUpdateTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrUpdateTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lineItems"] = o.LineItems
	toSerialize["currencyCode"] = o.CurrencyCode
	toSerialize["customerAddress"] = o.CustomerAddress
	if !IsNil(o.CustomerId) {
		toSerialize["customerId"] = o.CustomerId
	}
	if !IsNil(o.CustomerName) {
		toSerialize["customerName"] = o.CustomerName
	}
	if !IsNil(o.CustomerTaxIds) {
		toSerialize["customerTaxIds"] = o.CustomerTaxIds
	}
	if !IsNil(o.AccountingDate) {
		toSerialize["accountingDate"] = o.AccountingDate
	}
	if !IsNil(o.AccountingTime) {
		toSerialize["accountingTime"] = o.AccountingTime
	}
	if !IsNil(o.AccountingTimeZone) {
		toSerialize["accountingTimeZone"] = o.AccountingTimeZone
	}
	if !IsNil(o.TaxDate) {
		toSerialize["taxDate"] = o.TaxDate
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableCreateOrUpdateTransaction struct {
	value *CreateOrUpdateTransaction
	isSet bool
}

func (v NullableCreateOrUpdateTransaction) Get() *CreateOrUpdateTransaction {
	return v.value
}

func (v *NullableCreateOrUpdateTransaction) Set(val *CreateOrUpdateTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdateTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdateTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdateTransaction(val *CreateOrUpdateTransaction) *NullableCreateOrUpdateTransaction {
	return &NullableCreateOrUpdateTransaction{value: val, isSet: true}
}

func (v NullableCreateOrUpdateTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdateTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


