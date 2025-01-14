/*
 * Service Registry Service - Fleet Manager - v1
 *
 * Main entry point for the system, responsible for all sorts of management operations for the whole service of managed service registry.
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serviceregistrymgmt

import (
	"encoding/json"
	"time"
)

// RegistryStatus struct for RegistryStatus
type RegistryStatus struct {
	// ISO 8601 UTC timestamp.
	LastUpdated time.Time `json:"lastUpdated"`
	Value RegistryStatusValue `json:"value"`
}

// NewRegistryStatus instantiates a new RegistryStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistryStatus(lastUpdated time.Time, value RegistryStatusValue) *RegistryStatus {
	this := RegistryStatus{}
	this.LastUpdated = lastUpdated
	this.Value = value
	return &this
}

// NewRegistryStatusWithDefaults instantiates a new RegistryStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistryStatusWithDefaults() *RegistryStatus {
	this := RegistryStatus{}
	return &this
}

// GetLastUpdated returns the LastUpdated field value
func (o *RegistryStatus) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *RegistryStatus) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *RegistryStatus) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetValue returns the Value field value
func (o *RegistryStatus) GetValue() RegistryStatusValue {
	if o == nil {
		var ret RegistryStatusValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *RegistryStatus) GetValueOk() (*RegistryStatusValue, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *RegistryStatus) SetValue(v RegistryStatusValue) {
	o.Value = v
}

func (o RegistryStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableRegistryStatus struct {
	value *RegistryStatus
	isSet bool
}

func (v NullableRegistryStatus) Get() *RegistryStatus {
	return v.value
}

func (v *NullableRegistryStatus) Set(val *RegistryStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistryStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistryStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistryStatus(val *RegistryStatus) *NullableRegistryStatus {
	return &NullableRegistryStatus{value: val, isSet: true}
}

func (v NullableRegistryStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistryStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


