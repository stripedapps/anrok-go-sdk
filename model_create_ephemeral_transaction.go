/*
Anrok API

# API reference  The Anrok API server is accessible at `https://api.anrok.com`.  All requests are HTTP POSTs with JSON in the body.  Authentication is via an HTTP header `Authorization: Bearer {apiKey}`.  The default rate limit for a seller account is 10 API requests per second. 

API version: 1.1
Contact: support@anrok.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the CreateEphemeralTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateEphemeralTransaction{}

// CreateEphemeralTransaction struct for CreateEphemeralTransaction
type CreateEphemeralTransaction struct {
	// The line items in the transaction.
	LineItems []TransactionLineItem `json:"lineItems"`
	// Three letter ISO currency code (case insensitive).
	CurrencyCode string `json:"currencyCode"`
	CustomerAddress TransactionCustomerAddress `json:"customerAddress"`
	// The name of the customer. This is used for display purposes only.
	CustomerName *string `json:"customerName,omitempty"`
	// Tax IDs for the customer receiving the product
	CustomerTaxIds []CustomerTaxId `json:"customerTaxIds,omitempty"`
	ShipFromAddress *TransactionShipFromAddress `json:"shipFromAddress,omitempty"`
	// The date that this transaction occurred, for accounting purposes. Accounting date will typically correspond to the invoice date. This is used to determine which tax return the transaction belongs to.
	AccountingDate *string `json:"accountingDate,omitempty"`
	// The time that this transaction occurred, for accounting purposes. If `accountingDate` is not provided, `accountingTime` is required to compute an accounting date for the transaction.
	AccountingTime *time.Time `json:"accountingTime,omitempty"`
	// A “tz database” string used to compute an accounting date from the request's `accountingTime`. The request cannot provide both an `accountingDate` and an `accountingTimeZone`. If `accountingTime` is provided without specifying an `accountingTimeZone`, the time zone configured on the seller account will be used.
	AccountingTimeZone *string `json:"accountingTimeZone,omitempty"`
	// The date to use for tax calculations. If omitted, Anrok will use the the accounting date.
	TaxDate *string `json:"taxDate,omitempty"`
	// The Anrok customer ID used to link transactions for the same customer and to look up tax exemption certificates for a customer. This is typically the billing system's customer ID with a prefix to disambiguate. - If customerId is provided without customerName, that customer object must   already exist in Anrok. - Customer IDs are unique across the entire seller account. 
	CustomerId *string `json:"customerId,omitempty"`
}

type _CreateEphemeralTransaction CreateEphemeralTransaction

// NewCreateEphemeralTransaction instantiates a new CreateEphemeralTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEphemeralTransaction(lineItems []TransactionLineItem, currencyCode string, customerAddress TransactionCustomerAddress) *CreateEphemeralTransaction {
	this := CreateEphemeralTransaction{}
	this.LineItems = lineItems
	this.CurrencyCode = currencyCode
	this.CustomerAddress = customerAddress
	return &this
}

// NewCreateEphemeralTransactionWithDefaults instantiates a new CreateEphemeralTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEphemeralTransactionWithDefaults() *CreateEphemeralTransaction {
	this := CreateEphemeralTransaction{}
	return &this
}

// GetLineItems returns the LineItems field value
func (o *CreateEphemeralTransaction) GetLineItems() []TransactionLineItem {
	if o == nil {
		var ret []TransactionLineItem
		return ret
	}

	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value
// and a boolean to check if the value has been set.
func (o *CreateEphemeralTransaction) GetLineItemsOk() ([]TransactionLineItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.LineItems, true
}

// SetLineItems sets field value
func (o *CreateEphemeralTransaction) SetLineItems(v []TransactionLineItem) {
	o.LineItems = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *CreateEphemeralTransaction) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *CreateEphemeralTransaction) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *CreateEphemeralTransaction) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetCustomerAddress returns the CustomerAddress field value
func (o *CreateEphemeralTransaction) GetCustomerAddress() TransactionCustomerAddress {
	if o == nil {
		var ret TransactionCustomerAddress
		return ret
	}

	return o.CustomerAddress
}

// GetCustomerAddressOk returns a tuple with the CustomerAddress field value
// and a boolean to check if the value has been set.
func (o *CreateEphemeralTransaction) GetCustomerAddressOk() (*TransactionCustomerAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerAddress, true
}

// SetCustomerAddress sets field value
func (o *CreateEphemeralTransaction) SetCustomerAddress(v TransactionCustomerAddress) {
	o.CustomerAddress = v
}

// GetCustomerName returns the CustomerName field value if set, zero value otherwise.
func (o *CreateEphemeralTransaction) GetCustomerName() string {
	if o == nil || IsNil(o.CustomerName) {
		var ret string
		return ret
	}
	return *o.CustomerName
}

// GetCustomerNameOk returns a tuple with the CustomerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEphemeralTransaction) GetCustomerNameOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerName) {
		return nil, false
	}
	return o.CustomerName, true
}

// HasCustomerName returns a boolean if a field has been set.
func (o *CreateEphemeralTransaction) HasCustomerName() bool {
	if o != nil && !IsNil(o.CustomerName) {
		return true
	}

	return false
}

// SetCustomerName gets a reference to the given string and assigns it to the CustomerName field.
func (o *CreateEphemeralTransaction) SetCustomerName(v string) {
	o.CustomerName = &v
}

// GetCustomerTaxIds returns the CustomerTaxIds field value if set, zero value otherwise.
func (o *CreateEphemeralTransaction) GetCustomerTaxIds() []CustomerTaxId {
	if o == nil || IsNil(o.CustomerTaxIds) {
		var ret []CustomerTaxId
		return ret
	}
	return o.CustomerTaxIds
}

// GetCustomerTaxIdsOk returns a tuple with the CustomerTaxIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEphemeralTransaction) GetCustomerTaxIdsOk() ([]CustomerTaxId, bool) {
	if o == nil || IsNil(o.CustomerTaxIds) {
		return nil, false
	}
	return o.CustomerTaxIds, true
}

// HasCustomerTaxIds returns a boolean if a field has been set.
func (o *CreateEphemeralTransaction) HasCustomerTaxIds() bool {
	if o != nil && !IsNil(o.CustomerTaxIds) {
		return true
	}

	return false
}

// SetCustomerTaxIds gets a reference to the given []CustomerTaxId and assigns it to the CustomerTaxIds field.
func (o *CreateEphemeralTransaction) SetCustomerTaxIds(v []CustomerTaxId) {
	o.CustomerTaxIds = v
}

// GetShipFromAddress returns the ShipFromAddress field value if set, zero value otherwise.
func (o *CreateEphemeralTransaction) GetShipFromAddress() TransactionShipFromAddress {
	if o == nil || IsNil(o.ShipFromAddress) {
		var ret TransactionShipFromAddress
		return ret
	}
	return *o.ShipFromAddress
}

// GetShipFromAddressOk returns a tuple with the ShipFromAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEphemeralTransaction) GetShipFromAddressOk() (*TransactionShipFromAddress, bool) {
	if o == nil || IsNil(o.ShipFromAddress) {
		return nil, false
	}
	return o.ShipFromAddress, true
}

// HasShipFromAddress returns a boolean if a field has been set.
func (o *CreateEphemeralTransaction) HasShipFromAddress() bool {
	if o != nil && !IsNil(o.ShipFromAddress) {
		return true
	}

	return false
}

// SetShipFromAddress gets a reference to the given TransactionShipFromAddress and assigns it to the ShipFromAddress field.
func (o *CreateEphemeralTransaction) SetShipFromAddress(v TransactionShipFromAddress) {
	o.ShipFromAddress = &v
}

// GetAccountingDate returns the AccountingDate field value if set, zero value otherwise.
func (o *CreateEphemeralTransaction) GetAccountingDate() string {
	if o == nil || IsNil(o.AccountingDate) {
		var ret string
		return ret
	}
	return *o.AccountingDate
}

// GetAccountingDateOk returns a tuple with the AccountingDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEphemeralTransaction) GetAccountingDateOk() (*string, bool) {
	if o == nil || IsNil(o.AccountingDate) {
		return nil, false
	}
	return o.AccountingDate, true
}

// HasAccountingDate returns a boolean if a field has been set.
func (o *CreateEphemeralTransaction) HasAccountingDate() bool {
	if o != nil && !IsNil(o.AccountingDate) {
		return true
	}

	return false
}

// SetAccountingDate gets a reference to the given string and assigns it to the AccountingDate field.
func (o *CreateEphemeralTransaction) SetAccountingDate(v string) {
	o.AccountingDate = &v
}

// GetAccountingTime returns the AccountingTime field value if set, zero value otherwise.
func (o *CreateEphemeralTransaction) GetAccountingTime() time.Time {
	if o == nil || IsNil(o.AccountingTime) {
		var ret time.Time
		return ret
	}
	return *o.AccountingTime
}

// GetAccountingTimeOk returns a tuple with the AccountingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEphemeralTransaction) GetAccountingTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AccountingTime) {
		return nil, false
	}
	return o.AccountingTime, true
}

// HasAccountingTime returns a boolean if a field has been set.
func (o *CreateEphemeralTransaction) HasAccountingTime() bool {
	if o != nil && !IsNil(o.AccountingTime) {
		return true
	}

	return false
}

// SetAccountingTime gets a reference to the given time.Time and assigns it to the AccountingTime field.
func (o *CreateEphemeralTransaction) SetAccountingTime(v time.Time) {
	o.AccountingTime = &v
}

// GetAccountingTimeZone returns the AccountingTimeZone field value if set, zero value otherwise.
func (o *CreateEphemeralTransaction) GetAccountingTimeZone() string {
	if o == nil || IsNil(o.AccountingTimeZone) {
		var ret string
		return ret
	}
	return *o.AccountingTimeZone
}

// GetAccountingTimeZoneOk returns a tuple with the AccountingTimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEphemeralTransaction) GetAccountingTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.AccountingTimeZone) {
		return nil, false
	}
	return o.AccountingTimeZone, true
}

// HasAccountingTimeZone returns a boolean if a field has been set.
func (o *CreateEphemeralTransaction) HasAccountingTimeZone() bool {
	if o != nil && !IsNil(o.AccountingTimeZone) {
		return true
	}

	return false
}

// SetAccountingTimeZone gets a reference to the given string and assigns it to the AccountingTimeZone field.
func (o *CreateEphemeralTransaction) SetAccountingTimeZone(v string) {
	o.AccountingTimeZone = &v
}

// GetTaxDate returns the TaxDate field value if set, zero value otherwise.
func (o *CreateEphemeralTransaction) GetTaxDate() string {
	if o == nil || IsNil(o.TaxDate) {
		var ret string
		return ret
	}
	return *o.TaxDate
}

// GetTaxDateOk returns a tuple with the TaxDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEphemeralTransaction) GetTaxDateOk() (*string, bool) {
	if o == nil || IsNil(o.TaxDate) {
		return nil, false
	}
	return o.TaxDate, true
}

// HasTaxDate returns a boolean if a field has been set.
func (o *CreateEphemeralTransaction) HasTaxDate() bool {
	if o != nil && !IsNil(o.TaxDate) {
		return true
	}

	return false
}

// SetTaxDate gets a reference to the given string and assigns it to the TaxDate field.
func (o *CreateEphemeralTransaction) SetTaxDate(v string) {
	o.TaxDate = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *CreateEphemeralTransaction) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEphemeralTransaction) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *CreateEphemeralTransaction) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *CreateEphemeralTransaction) SetCustomerId(v string) {
	o.CustomerId = &v
}

func (o CreateEphemeralTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateEphemeralTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lineItems"] = o.LineItems
	toSerialize["currencyCode"] = o.CurrencyCode
	toSerialize["customerAddress"] = o.CustomerAddress
	if !IsNil(o.CustomerName) {
		toSerialize["customerName"] = o.CustomerName
	}
	if !IsNil(o.CustomerTaxIds) {
		toSerialize["customerTaxIds"] = o.CustomerTaxIds
	}
	if !IsNil(o.ShipFromAddress) {
		toSerialize["shipFromAddress"] = o.ShipFromAddress
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
	if !IsNil(o.CustomerId) {
		toSerialize["customerId"] = o.CustomerId
	}
	return toSerialize, nil
}

func (o *CreateEphemeralTransaction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"lineItems",
		"currencyCode",
		"customerAddress",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateEphemeralTransaction := _CreateEphemeralTransaction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateEphemeralTransaction)

	if err != nil {
		return err
	}

	*o = CreateEphemeralTransaction(varCreateEphemeralTransaction)

	return err
}

type NullableCreateEphemeralTransaction struct {
	value *CreateEphemeralTransaction
	isSet bool
}

func (v NullableCreateEphemeralTransaction) Get() *CreateEphemeralTransaction {
	return v.value
}

func (v *NullableCreateEphemeralTransaction) Set(val *CreateEphemeralTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEphemeralTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEphemeralTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEphemeralTransaction(val *CreateEphemeralTransaction) *NullableCreateEphemeralTransaction {
	return &NullableCreateEphemeralTransaction{value: val, isSet: true}
}

func (v NullableCreateEphemeralTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEphemeralTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


