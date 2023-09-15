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

// checks if the CustomerCertificateJuris type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerCertificateJuris{}

// CustomerCertificateJuris struct for CustomerCertificateJuris
type CustomerCertificateJuris struct {
	// ID of jurisdiction
	JurisId interface{} `json:"jurisId"`
	// Registration ID for jurisdiction
	RegistrationId interface{} `json:"registrationId,omitempty"`
	// Expiration date of certificate in this jurisdiction
	EffectiveDateEndi interface{} `json:"effectiveDateEndi,omitempty"`
	// Optional internal notes
	Notes interface{} `json:"notes,omitempty"`
}

// NewCustomerCertificateJuris instantiates a new CustomerCertificateJuris object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerCertificateJuris(jurisId interface{}) *CustomerCertificateJuris {
	this := CustomerCertificateJuris{}
	this.JurisId = jurisId
	return &this
}

// NewCustomerCertificateJurisWithDefaults instantiates a new CustomerCertificateJuris object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerCertificateJurisWithDefaults() *CustomerCertificateJuris {
	this := CustomerCertificateJuris{}
	return &this
}

// GetJurisId returns the JurisId field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CustomerCertificateJuris) GetJurisId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.JurisId
}

// GetJurisIdOk returns a tuple with the JurisId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerCertificateJuris) GetJurisIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.JurisId) {
		return nil, false
	}
	return &o.JurisId, true
}

// SetJurisId sets field value
func (o *CustomerCertificateJuris) SetJurisId(v interface{}) {
	o.JurisId = v
}

// GetRegistrationId returns the RegistrationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerCertificateJuris) GetRegistrationId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RegistrationId
}

// GetRegistrationIdOk returns a tuple with the RegistrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerCertificateJuris) GetRegistrationIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RegistrationId) {
		return nil, false
	}
	return &o.RegistrationId, true
}

// HasRegistrationId returns a boolean if a field has been set.
func (o *CustomerCertificateJuris) HasRegistrationId() bool {
	if o != nil && IsNil(o.RegistrationId) {
		return true
	}

	return false
}

// SetRegistrationId gets a reference to the given interface{} and assigns it to the RegistrationId field.
func (o *CustomerCertificateJuris) SetRegistrationId(v interface{}) {
	o.RegistrationId = v
}

// GetEffectiveDateEndi returns the EffectiveDateEndi field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerCertificateJuris) GetEffectiveDateEndi() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.EffectiveDateEndi
}

// GetEffectiveDateEndiOk returns a tuple with the EffectiveDateEndi field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerCertificateJuris) GetEffectiveDateEndiOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EffectiveDateEndi) {
		return nil, false
	}
	return &o.EffectiveDateEndi, true
}

// HasEffectiveDateEndi returns a boolean if a field has been set.
func (o *CustomerCertificateJuris) HasEffectiveDateEndi() bool {
	if o != nil && IsNil(o.EffectiveDateEndi) {
		return true
	}

	return false
}

// SetEffectiveDateEndi gets a reference to the given interface{} and assigns it to the EffectiveDateEndi field.
func (o *CustomerCertificateJuris) SetEffectiveDateEndi(v interface{}) {
	o.EffectiveDateEndi = v
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerCertificateJuris) GetNotes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerCertificateJuris) GetNotesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return &o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *CustomerCertificateJuris) HasNotes() bool {
	if o != nil && IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given interface{} and assigns it to the Notes field.
func (o *CustomerCertificateJuris) SetNotes(v interface{}) {
	o.Notes = v
}

func (o CustomerCertificateJuris) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerCertificateJuris) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.JurisId != nil {
		toSerialize["jurisId"] = o.JurisId
	}
	if o.RegistrationId != nil {
		toSerialize["registrationId"] = o.RegistrationId
	}
	if o.EffectiveDateEndi != nil {
		toSerialize["effectiveDateEndi"] = o.EffectiveDateEndi
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	return toSerialize, nil
}

type NullableCustomerCertificateJuris struct {
	value *CustomerCertificateJuris
	isSet bool
}

func (v NullableCustomerCertificateJuris) Get() *CustomerCertificateJuris {
	return v.value
}

func (v *NullableCustomerCertificateJuris) Set(val *CustomerCertificateJuris) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerCertificateJuris) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerCertificateJuris) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerCertificateJuris(val *CustomerCertificateJuris) *NullableCustomerCertificateJuris {
	return &NullableCustomerCertificateJuris{value: val, isSet: true}
}

func (v NullableCustomerCertificateJuris) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerCertificateJuris) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


