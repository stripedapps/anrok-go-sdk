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

// checks if the CreateEphemeralTransactionSuccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateEphemeralTransactionSuccess{}

// CreateEphemeralTransactionSuccess struct for CreateEphemeralTransactionSuccess
type CreateEphemeralTransactionSuccess struct {
	// The total tax amount to collect from the customer, in the smallest denomination of the currency (e.g. cents or pennies)
	TaxAmountToCollect *int64 `json:"taxAmountToCollect,omitempty"`
	LineItems []CreateEphemeralTransactionSuccessLineItemsInner `json:"lineItems,omitempty"`
}

// NewCreateEphemeralTransactionSuccess instantiates a new CreateEphemeralTransactionSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEphemeralTransactionSuccess() *CreateEphemeralTransactionSuccess {
	this := CreateEphemeralTransactionSuccess{}
	return &this
}

// NewCreateEphemeralTransactionSuccessWithDefaults instantiates a new CreateEphemeralTransactionSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEphemeralTransactionSuccessWithDefaults() *CreateEphemeralTransactionSuccess {
	this := CreateEphemeralTransactionSuccess{}
	return &this
}

// GetTaxAmountToCollect returns the TaxAmountToCollect field value if set, zero value otherwise.
func (o *CreateEphemeralTransactionSuccess) GetTaxAmountToCollect() int64 {
	if o == nil || IsNil(o.TaxAmountToCollect) {
		var ret int64
		return ret
	}
	return *o.TaxAmountToCollect
}

// GetTaxAmountToCollectOk returns a tuple with the TaxAmountToCollect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEphemeralTransactionSuccess) GetTaxAmountToCollectOk() (*int64, bool) {
	if o == nil || IsNil(o.TaxAmountToCollect) {
		return nil, false
	}
	return o.TaxAmountToCollect, true
}

// HasTaxAmountToCollect returns a boolean if a field has been set.
func (o *CreateEphemeralTransactionSuccess) HasTaxAmountToCollect() bool {
	if o != nil && !IsNil(o.TaxAmountToCollect) {
		return true
	}

	return false
}

// SetTaxAmountToCollect gets a reference to the given int64 and assigns it to the TaxAmountToCollect field.
func (o *CreateEphemeralTransactionSuccess) SetTaxAmountToCollect(v int64) {
	o.TaxAmountToCollect = &v
}

// GetLineItems returns the LineItems field value if set, zero value otherwise.
func (o *CreateEphemeralTransactionSuccess) GetLineItems() []CreateEphemeralTransactionSuccessLineItemsInner {
	if o == nil || IsNil(o.LineItems) {
		var ret []CreateEphemeralTransactionSuccessLineItemsInner
		return ret
	}
	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEphemeralTransactionSuccess) GetLineItemsOk() ([]CreateEphemeralTransactionSuccessLineItemsInner, bool) {
	if o == nil || IsNil(o.LineItems) {
		return nil, false
	}
	return o.LineItems, true
}

// HasLineItems returns a boolean if a field has been set.
func (o *CreateEphemeralTransactionSuccess) HasLineItems() bool {
	if o != nil && !IsNil(o.LineItems) {
		return true
	}

	return false
}

// SetLineItems gets a reference to the given []CreateEphemeralTransactionSuccessLineItemsInner and assigns it to the LineItems field.
func (o *CreateEphemeralTransactionSuccess) SetLineItems(v []CreateEphemeralTransactionSuccessLineItemsInner) {
	o.LineItems = v
}

func (o CreateEphemeralTransactionSuccess) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateEphemeralTransactionSuccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TaxAmountToCollect) {
		toSerialize["taxAmountToCollect"] = o.TaxAmountToCollect
	}
	if !IsNil(o.LineItems) {
		toSerialize["lineItems"] = o.LineItems
	}
	return toSerialize, nil
}

type NullableCreateEphemeralTransactionSuccess struct {
	value *CreateEphemeralTransactionSuccess
	isSet bool
}

func (v NullableCreateEphemeralTransactionSuccess) Get() *CreateEphemeralTransactionSuccess {
	return v.value
}

func (v *NullableCreateEphemeralTransactionSuccess) Set(val *CreateEphemeralTransactionSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEphemeralTransactionSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEphemeralTransactionSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEphemeralTransactionSuccess(val *CreateEphemeralTransactionSuccess) *NullableCreateEphemeralTransactionSuccess {
	return &NullableCreateEphemeralTransactionSuccess{value: val, isSet: true}
}

func (v NullableCreateEphemeralTransactionSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEphemeralTransactionSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


