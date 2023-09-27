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

// checks if the CustomerCertificateFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerCertificateFile{}

// CustomerCertificateFile struct for CustomerCertificateFile
type CustomerCertificateFile struct {
	// File name
	Name string `json:"name"`
	// Base64 encoded certificate image contents
	ContentsBase64 string `json:"contentsBase64"`
}

// NewCustomerCertificateFile instantiates a new CustomerCertificateFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerCertificateFile(name string, contentsBase64 string) *CustomerCertificateFile {
	this := CustomerCertificateFile{}
	this.Name = name
	this.ContentsBase64 = contentsBase64
	return &this
}

// NewCustomerCertificateFileWithDefaults instantiates a new CustomerCertificateFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerCertificateFileWithDefaults() *CustomerCertificateFile {
	this := CustomerCertificateFile{}
	return &this
}

// GetName returns the Name field value
func (o *CustomerCertificateFile) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomerCertificateFile) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomerCertificateFile) SetName(v string) {
	o.Name = v
}

// GetContentsBase64 returns the ContentsBase64 field value
func (o *CustomerCertificateFile) GetContentsBase64() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentsBase64
}

// GetContentsBase64Ok returns a tuple with the ContentsBase64 field value
// and a boolean to check if the value has been set.
func (o *CustomerCertificateFile) GetContentsBase64Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentsBase64, true
}

// SetContentsBase64 sets field value
func (o *CustomerCertificateFile) SetContentsBase64(v string) {
	o.ContentsBase64 = v
}

func (o CustomerCertificateFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerCertificateFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["contentsBase64"] = o.ContentsBase64
	return toSerialize, nil
}

type NullableCustomerCertificateFile struct {
	value *CustomerCertificateFile
	isSet bool
}

func (v NullableCustomerCertificateFile) Get() *CustomerCertificateFile {
	return v.value
}

func (v *NullableCustomerCertificateFile) Set(val *CustomerCertificateFile) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerCertificateFile) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerCertificateFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerCertificateFile(val *CustomerCertificateFile) *NullableCustomerCertificateFile {
	return &NullableCustomerCertificateFile{value: val, isSet: true}
}

func (v NullableCustomerCertificateFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerCertificateFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


