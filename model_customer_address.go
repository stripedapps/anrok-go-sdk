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

// checks if the CustomerAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerAddress{}

// CustomerAddress struct for CustomerAddress
type CustomerAddress struct {
	// The country's full name or 2-letter ISO code
	Country NullableString `json:"country,omitempty"`
	// The \"delivery line\" of the address
	Line1 NullableString `json:"line1,omitempty"`
	// City or town name
	City NullableString `json:"city,omitempty"`
	// Other principal subdivision (i.e. province, state, county)
	Region NullableString `json:"region,omitempty"`
	// Postal code
	PostalCode NullableString `json:"postalCode,omitempty"`
}

// NewCustomerAddress instantiates a new CustomerAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerAddress() *CustomerAddress {
	this := CustomerAddress{}
	return &this
}

// NewCustomerAddressWithDefaults instantiates a new CustomerAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerAddressWithDefaults() *CustomerAddress {
	this := CustomerAddress{}
	return &this
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerAddress) GetCountry() string {
	if o == nil || IsNil(o.Country.Get()) {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerAddress) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *CustomerAddress) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *CustomerAddress) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *CustomerAddress) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *CustomerAddress) UnsetCountry() {
	o.Country.Unset()
}

// GetLine1 returns the Line1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerAddress) GetLine1() string {
	if o == nil || IsNil(o.Line1.Get()) {
		var ret string
		return ret
	}
	return *o.Line1.Get()
}

// GetLine1Ok returns a tuple with the Line1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerAddress) GetLine1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Line1.Get(), o.Line1.IsSet()
}

// HasLine1 returns a boolean if a field has been set.
func (o *CustomerAddress) HasLine1() bool {
	if o != nil && o.Line1.IsSet() {
		return true
	}

	return false
}

// SetLine1 gets a reference to the given NullableString and assigns it to the Line1 field.
func (o *CustomerAddress) SetLine1(v string) {
	o.Line1.Set(&v)
}
// SetLine1Nil sets the value for Line1 to be an explicit nil
func (o *CustomerAddress) SetLine1Nil() {
	o.Line1.Set(nil)
}

// UnsetLine1 ensures that no value is present for Line1, not even an explicit nil
func (o *CustomerAddress) UnsetLine1() {
	o.Line1.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerAddress) GetCity() string {
	if o == nil || IsNil(o.City.Get()) {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerAddress) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *CustomerAddress) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *CustomerAddress) SetCity(v string) {
	o.City.Set(&v)
}
// SetCityNil sets the value for City to be an explicit nil
func (o *CustomerAddress) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *CustomerAddress) UnsetCity() {
	o.City.Unset()
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerAddress) GetRegion() string {
	if o == nil || IsNil(o.Region.Get()) {
		var ret string
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerAddress) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *CustomerAddress) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableString and assigns it to the Region field.
func (o *CustomerAddress) SetRegion(v string) {
	o.Region.Set(&v)
}
// SetRegionNil sets the value for Region to be an explicit nil
func (o *CustomerAddress) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *CustomerAddress) UnsetRegion() {
	o.Region.Unset()
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomerAddress) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode.Get()) {
		var ret string
		return ret
	}
	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomerAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// HasPostalCode returns a boolean if a field has been set.
func (o *CustomerAddress) HasPostalCode() bool {
	if o != nil && o.PostalCode.IsSet() {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given NullableString and assigns it to the PostalCode field.
func (o *CustomerAddress) SetPostalCode(v string) {
	o.PostalCode.Set(&v)
}
// SetPostalCodeNil sets the value for PostalCode to be an explicit nil
func (o *CustomerAddress) SetPostalCodeNil() {
	o.PostalCode.Set(nil)
}

// UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
func (o *CustomerAddress) UnsetPostalCode() {
	o.PostalCode.Unset()
}

func (o CustomerAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
	}
	if o.Line1.IsSet() {
		toSerialize["line1"] = o.Line1.Get()
	}
	if o.City.IsSet() {
		toSerialize["city"] = o.City.Get()
	}
	if o.Region.IsSet() {
		toSerialize["region"] = o.Region.Get()
	}
	if o.PostalCode.IsSet() {
		toSerialize["postalCode"] = o.PostalCode.Get()
	}
	return toSerialize, nil
}

type NullableCustomerAddress struct {
	value *CustomerAddress
	isSet bool
}

func (v NullableCustomerAddress) Get() *CustomerAddress {
	return v.value
}

func (v *NullableCustomerAddress) Set(val *CustomerAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerAddress(val *CustomerAddress) *NullableCustomerAddress {
	return &NullableCustomerAddress{value: val, isSet: true}
}

func (v NullableCustomerAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


