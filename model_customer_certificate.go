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

// checks if the CustomerCertificate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerCertificate{}

// CustomerCertificate struct for CustomerCertificate
type CustomerCertificate struct {
	// The Anrok customer ID used to link transactions for the same customer and to look up tax exemption certificates for a customer. This is typically the billing system's customer ID with a prefix to disambiguate. - If customerId is provided without customerName, that customer object must already exist in Anrok. - If both customerId and customerName are provided, the customer object will be created if it is not already present. - Customer IDs are unique across the entire seller account. 
	CustomerId interface{} `json:"customerId"`
	// The name of the customer
	CustomerName interface{} `json:"customerName,omitempty"`
	// Effective date of certificate
	EffectiveDateBegin interface{} `json:"effectiveDateBegin"`
	// Certificate exemption number
	ExemptionNumber interface{} `json:"exemptionNumber,omitempty"`
	// Optional internal notes
	Notes interface{} `json:"notes,omitempty"`
	CertificateFile CustomerCertificateFile `json:"certificateFile"`
	// Jurisdictions for which certificate applies
	Jurises interface{} `json:"jurises"`
}

// NewCustomerCertificate instantiates a new CustomerCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerCertificate(customerId interface{}, effectiveDateBegin interface{}, certificateFile CustomerCertificateFile, jurises interface{}) *CustomerCertificate {
	this := CustomerCertificate{}
	this.CustomerId = customerId
	this.EffectiveDateBegin = effectiveDateBegin
	this.CertificateFile = certificateFile
	this.Jurises = jurises
	return &this
}

// NewCustomerCertificateWithDefaults instantiates a new CustomerCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerCertificateWithDefaults() *CustomerCertificate {
	this := CustomerCertificate{}
	return &this
}

// GetCustomerId returns the CustomerId field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CustomerCertificate) GetCustomerId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerCertificate) GetCustomerIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *CustomerCertificate) SetCustomerId(v interface{}) {
	o.CustomerId = v
}

// GetCustomerName returns the CustomerName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerCertificate) GetCustomerName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CustomerName
}

// GetCustomerNameOk returns a tuple with the CustomerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerCertificate) GetCustomerNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CustomerName) {
		return nil, false
	}
	return &o.CustomerName, true
}

// HasCustomerName returns a boolean if a field has been set.
func (o *CustomerCertificate) HasCustomerName() bool {
	if o != nil && IsNil(o.CustomerName) {
		return true
	}

	return false
}

// SetCustomerName gets a reference to the given interface{} and assigns it to the CustomerName field.
func (o *CustomerCertificate) SetCustomerName(v interface{}) {
	o.CustomerName = v
}

// GetEffectiveDateBegin returns the EffectiveDateBegin field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CustomerCertificate) GetEffectiveDateBegin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.EffectiveDateBegin
}

// GetEffectiveDateBeginOk returns a tuple with the EffectiveDateBegin field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerCertificate) GetEffectiveDateBeginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.EffectiveDateBegin) {
		return nil, false
	}
	return &o.EffectiveDateBegin, true
}

// SetEffectiveDateBegin sets field value
func (o *CustomerCertificate) SetEffectiveDateBegin(v interface{}) {
	o.EffectiveDateBegin = v
}

// GetExemptionNumber returns the ExemptionNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerCertificate) GetExemptionNumber() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExemptionNumber
}

// GetExemptionNumberOk returns a tuple with the ExemptionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerCertificate) GetExemptionNumberOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ExemptionNumber) {
		return nil, false
	}
	return &o.ExemptionNumber, true
}

// HasExemptionNumber returns a boolean if a field has been set.
func (o *CustomerCertificate) HasExemptionNumber() bool {
	if o != nil && IsNil(o.ExemptionNumber) {
		return true
	}

	return false
}

// SetExemptionNumber gets a reference to the given interface{} and assigns it to the ExemptionNumber field.
func (o *CustomerCertificate) SetExemptionNumber(v interface{}) {
	o.ExemptionNumber = v
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerCertificate) GetNotes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerCertificate) GetNotesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return &o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *CustomerCertificate) HasNotes() bool {
	if o != nil && IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given interface{} and assigns it to the Notes field.
func (o *CustomerCertificate) SetNotes(v interface{}) {
	o.Notes = v
}

// GetCertificateFile returns the CertificateFile field value
func (o *CustomerCertificate) GetCertificateFile() CustomerCertificateFile {
	if o == nil {
		var ret CustomerCertificateFile
		return ret
	}

	return o.CertificateFile
}

// GetCertificateFileOk returns a tuple with the CertificateFile field value
// and a boolean to check if the value has been set.
func (o *CustomerCertificate) GetCertificateFileOk() (*CustomerCertificateFile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CertificateFile, true
}

// SetCertificateFile sets field value
func (o *CustomerCertificate) SetCertificateFile(v CustomerCertificateFile) {
	o.CertificateFile = v
}

// GetJurises returns the Jurises field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CustomerCertificate) GetJurises() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Jurises
}

// GetJurisesOk returns a tuple with the Jurises field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerCertificate) GetJurisesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Jurises) {
		return nil, false
	}
	return &o.Jurises, true
}

// SetJurises sets field value
func (o *CustomerCertificate) SetJurises(v interface{}) {
	o.Jurises = v
}

func (o CustomerCertificate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerCertificate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomerId != nil {
		toSerialize["customerId"] = o.CustomerId
	}
	if o.CustomerName != nil {
		toSerialize["customerName"] = o.CustomerName
	}
	if o.EffectiveDateBegin != nil {
		toSerialize["effectiveDateBegin"] = o.EffectiveDateBegin
	}
	if o.ExemptionNumber != nil {
		toSerialize["exemptionNumber"] = o.ExemptionNumber
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	toSerialize["certificateFile"] = o.CertificateFile
	if o.Jurises != nil {
		toSerialize["jurises"] = o.Jurises
	}
	return toSerialize, nil
}

type NullableCustomerCertificate struct {
	value *CustomerCertificate
	isSet bool
}

func (v NullableCustomerCertificate) Get() *CustomerCertificate {
	return v.value
}

func (v *NullableCustomerCertificate) Set(val *CustomerCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerCertificate(val *CustomerCertificate) *NullableCustomerCertificate {
	return &NullableCustomerCertificate{value: val, isSet: true}
}

func (v NullableCustomerCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


