/*
Anrok API

# API reference  The Anrok API server is accessible at “https://api.anrok.com”.  All requests are HTTP POSTs with JSON in the body.  Authentication is via an HTTP header “Authorization: Bearer {sellerId}/{apiKeyId}/secret.{apiKeySecret}”.  The default rate limit for a seller account is 10 API requests per second. 

API version: 0.0.1
Contact: support@anrok.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ProductIdMappingsAdd409Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductIdMappingsAdd409Response{}

// ProductIdMappingsAdd409Response struct for ProductIdMappingsAdd409Response
type ProductIdMappingsAdd409Response struct {
	Type interface{} `json:"type,omitempty"`
}

// NewProductIdMappingsAdd409Response instantiates a new ProductIdMappingsAdd409Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductIdMappingsAdd409Response() *ProductIdMappingsAdd409Response {
	this := ProductIdMappingsAdd409Response{}
	return &this
}

// NewProductIdMappingsAdd409ResponseWithDefaults instantiates a new ProductIdMappingsAdd409Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductIdMappingsAdd409ResponseWithDefaults() *ProductIdMappingsAdd409Response {
	this := ProductIdMappingsAdd409Response{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProductIdMappingsAdd409Response) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProductIdMappingsAdd409Response) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProductIdMappingsAdd409Response) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *ProductIdMappingsAdd409Response) SetType(v interface{}) {
	o.Type = v
}

func (o ProductIdMappingsAdd409Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductIdMappingsAdd409Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableProductIdMappingsAdd409Response struct {
	value *ProductIdMappingsAdd409Response
	isSet bool
}

func (v NullableProductIdMappingsAdd409Response) Get() *ProductIdMappingsAdd409Response {
	return v.value
}

func (v *NullableProductIdMappingsAdd409Response) Set(val *ProductIdMappingsAdd409Response) {
	v.value = val
	v.isSet = true
}

func (v NullableProductIdMappingsAdd409Response) IsSet() bool {
	return v.isSet
}

func (v *NullableProductIdMappingsAdd409Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductIdMappingsAdd409Response(val *ProductIdMappingsAdd409Response) *NullableProductIdMappingsAdd409Response {
	return &NullableProductIdMappingsAdd409Response{value: val, isSet: true}
}

func (v NullableProductIdMappingsAdd409Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductIdMappingsAdd409Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


