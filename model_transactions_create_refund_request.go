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
)

// checks if the TransactionsCreateRefundRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionsCreateRefundRequest{}

// TransactionsCreateRefundRequest struct for TransactionsCreateRefundRequest
type TransactionsCreateRefundRequest struct {
	// The ID of the transaction you want to create a refund of
	OriginalTransactionId string `json:"originalTransactionId"`
	// The ID for the new transaction
	NewTransactionId string `json:"newTransactionId"`
	// The expected transaction version. The refund will fail if this is not the latest version of the transaction.
	OriginalTransactionExpectedVersion *int32 `json:"originalTransactionExpectedVersion,omitempty"`
}

// NewTransactionsCreateRefundRequest instantiates a new TransactionsCreateRefundRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsCreateRefundRequest(originalTransactionId string, newTransactionId string) *TransactionsCreateRefundRequest {
	this := TransactionsCreateRefundRequest{}
	this.OriginalTransactionId = originalTransactionId
	this.NewTransactionId = newTransactionId
	return &this
}

// NewTransactionsCreateRefundRequestWithDefaults instantiates a new TransactionsCreateRefundRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsCreateRefundRequestWithDefaults() *TransactionsCreateRefundRequest {
	this := TransactionsCreateRefundRequest{}
	return &this
}

// GetOriginalTransactionId returns the OriginalTransactionId field value
func (o *TransactionsCreateRefundRequest) GetOriginalTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginalTransactionId
}

// GetOriginalTransactionIdOk returns a tuple with the OriginalTransactionId field value
// and a boolean to check if the value has been set.
func (o *TransactionsCreateRefundRequest) GetOriginalTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalTransactionId, true
}

// SetOriginalTransactionId sets field value
func (o *TransactionsCreateRefundRequest) SetOriginalTransactionId(v string) {
	o.OriginalTransactionId = v
}

// GetNewTransactionId returns the NewTransactionId field value
func (o *TransactionsCreateRefundRequest) GetNewTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewTransactionId
}

// GetNewTransactionIdOk returns a tuple with the NewTransactionId field value
// and a boolean to check if the value has been set.
func (o *TransactionsCreateRefundRequest) GetNewTransactionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewTransactionId, true
}

// SetNewTransactionId sets field value
func (o *TransactionsCreateRefundRequest) SetNewTransactionId(v string) {
	o.NewTransactionId = v
}

// GetOriginalTransactionExpectedVersion returns the OriginalTransactionExpectedVersion field value if set, zero value otherwise.
func (o *TransactionsCreateRefundRequest) GetOriginalTransactionExpectedVersion() int32 {
	if o == nil || IsNil(o.OriginalTransactionExpectedVersion) {
		var ret int32
		return ret
	}
	return *o.OriginalTransactionExpectedVersion
}

// GetOriginalTransactionExpectedVersionOk returns a tuple with the OriginalTransactionExpectedVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsCreateRefundRequest) GetOriginalTransactionExpectedVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.OriginalTransactionExpectedVersion) {
		return nil, false
	}
	return o.OriginalTransactionExpectedVersion, true
}

// HasOriginalTransactionExpectedVersion returns a boolean if a field has been set.
func (o *TransactionsCreateRefundRequest) HasOriginalTransactionExpectedVersion() bool {
	if o != nil && !IsNil(o.OriginalTransactionExpectedVersion) {
		return true
	}

	return false
}

// SetOriginalTransactionExpectedVersion gets a reference to the given int32 and assigns it to the OriginalTransactionExpectedVersion field.
func (o *TransactionsCreateRefundRequest) SetOriginalTransactionExpectedVersion(v int32) {
	o.OriginalTransactionExpectedVersion = &v
}

func (o TransactionsCreateRefundRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionsCreateRefundRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["originalTransactionId"] = o.OriginalTransactionId
	toSerialize["newTransactionId"] = o.NewTransactionId
	if !IsNil(o.OriginalTransactionExpectedVersion) {
		toSerialize["originalTransactionExpectedVersion"] = o.OriginalTransactionExpectedVersion
	}
	return toSerialize, nil
}

type NullableTransactionsCreateRefundRequest struct {
	value *TransactionsCreateRefundRequest
	isSet bool
}

func (v NullableTransactionsCreateRefundRequest) Get() *TransactionsCreateRefundRequest {
	return v.value
}

func (v *NullableTransactionsCreateRefundRequest) Set(val *TransactionsCreateRefundRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsCreateRefundRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsCreateRefundRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsCreateRefundRequest(val *TransactionsCreateRefundRequest) *NullableTransactionsCreateRefundRequest {
	return &NullableTransactionsCreateRefundRequest{value: val, isSet: true}
}

func (v NullableTransactionsCreateRefundRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsCreateRefundRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


