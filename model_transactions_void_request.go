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
)

// checks if the TransactionsVoidRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionsVoidRequest{}

// TransactionsVoidRequest struct for TransactionsVoidRequest
type TransactionsVoidRequest struct {
	// The expected transaction version. The void will fail if this is not the latest version of the transaction.
	TransactionExpectedVersion *int32 `json:"transactionExpectedVersion,omitempty"`
}

// NewTransactionsVoidRequest instantiates a new TransactionsVoidRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsVoidRequest() *TransactionsVoidRequest {
	this := TransactionsVoidRequest{}
	return &this
}

// NewTransactionsVoidRequestWithDefaults instantiates a new TransactionsVoidRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsVoidRequestWithDefaults() *TransactionsVoidRequest {
	this := TransactionsVoidRequest{}
	return &this
}

// GetTransactionExpectedVersion returns the TransactionExpectedVersion field value if set, zero value otherwise.
func (o *TransactionsVoidRequest) GetTransactionExpectedVersion() int32 {
	if o == nil || IsNil(o.TransactionExpectedVersion) {
		var ret int32
		return ret
	}
	return *o.TransactionExpectedVersion
}

// GetTransactionExpectedVersionOk returns a tuple with the TransactionExpectedVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsVoidRequest) GetTransactionExpectedVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.TransactionExpectedVersion) {
		return nil, false
	}
	return o.TransactionExpectedVersion, true
}

// HasTransactionExpectedVersion returns a boolean if a field has been set.
func (o *TransactionsVoidRequest) HasTransactionExpectedVersion() bool {
	if o != nil && !IsNil(o.TransactionExpectedVersion) {
		return true
	}

	return false
}

// SetTransactionExpectedVersion gets a reference to the given int32 and assigns it to the TransactionExpectedVersion field.
func (o *TransactionsVoidRequest) SetTransactionExpectedVersion(v int32) {
	o.TransactionExpectedVersion = &v
}

func (o TransactionsVoidRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionsVoidRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TransactionExpectedVersion) {
		toSerialize["transactionExpectedVersion"] = o.TransactionExpectedVersion
	}
	return toSerialize, nil
}

type NullableTransactionsVoidRequest struct {
	value *TransactionsVoidRequest
	isSet bool
}

func (v NullableTransactionsVoidRequest) Get() *TransactionsVoidRequest {
	return v.value
}

func (v *NullableTransactionsVoidRequest) Set(val *TransactionsVoidRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsVoidRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsVoidRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsVoidRequest(val *TransactionsVoidRequest) *NullableTransactionsVoidRequest {
	return &NullableTransactionsVoidRequest{value: val, isSet: true}
}

func (v NullableTransactionsVoidRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsVoidRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


