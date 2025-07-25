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
	"bytes"
	"fmt"
)

// checks if the CustomerTaxId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerTaxId{}

// CustomerTaxId The customer VAT registration number for a non-US country
type CustomerTaxId struct {
	// This setting is purely metadata and does not affect tax calculation.
	Type string `json:"type"`
	// The customer VAT registration number for a non-US country
	Value string `json:"value"`
}

type _CustomerTaxId CustomerTaxId

// NewCustomerTaxId instantiates a new CustomerTaxId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerTaxId(type_ string, value string) *CustomerTaxId {
	this := CustomerTaxId{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewCustomerTaxIdWithDefaults instantiates a new CustomerTaxId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerTaxIdWithDefaults() *CustomerTaxId {
	this := CustomerTaxId{}
	return &this
}

// GetType returns the Type field value
func (o *CustomerTaxId) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomerTaxId) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CustomerTaxId) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *CustomerTaxId) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CustomerTaxId) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CustomerTaxId) SetValue(v string) {
	o.Value = v
}

func (o CustomerTaxId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerTaxId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *CustomerTaxId) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"value",
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

	varCustomerTaxId := _CustomerTaxId{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomerTaxId)

	if err != nil {
		return err
	}

	*o = CustomerTaxId(varCustomerTaxId)

	return err
}

type NullableCustomerTaxId struct {
	value *CustomerTaxId
	isSet bool
}

func (v NullableCustomerTaxId) Get() *CustomerTaxId {
	return v.value
}

func (v *NullableCustomerTaxId) Set(val *CustomerTaxId) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerTaxId) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerTaxId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerTaxId(val *CustomerTaxId) *NullableCustomerTaxId {
	return &NullableCustomerTaxId{value: val, isSet: true}
}

func (v NullableCustomerTaxId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerTaxId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


