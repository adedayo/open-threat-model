/*
Open Threat Model Specification

A Schema Definition for Open Threat Model https://github.com/iriusrisk/OpenThreatModel\"

API version: 0.1.0
Contact: info@cysoh.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package otm

import (
	"encoding/json"
)

// checks if the Parent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Parent{}

// Parent struct for Parent
type Parent struct {
	TrustZone *string `json:"trustZone,omitempty" yaml:"trustZone,omitempty"`
	Component *string `json:"component,omitempty" yaml:"component,omitempty"`
}

// NewParent instantiates a new Parent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParent() *Parent {
	this := Parent{}
	return &this
}

// NewParentWithDefaults instantiates a new Parent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParentWithDefaults() *Parent {
	this := Parent{}
	return &this
}

// GetTrustZone returns the TrustZone field value if set, zero value otherwise.
func (o *Parent) GetTrustZone() string {
	if o == nil || IsNil(o.TrustZone) {
		var ret string
		return ret
	}
	return *o.TrustZone
}

// GetTrustZoneOk returns a tuple with the TrustZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Parent) GetTrustZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TrustZone) {
		return nil, false
	}
	return o.TrustZone, true
}

// HasTrustZone returns a boolean if a field has been set.
func (o *Parent) HasTrustZone() bool {
	if o != nil && !IsNil(o.TrustZone) {
		return true
	}

	return false
}

// SetTrustZone gets a reference to the given string and assigns it to the TrustZone field.
func (o *Parent) SetTrustZone(v string) {
	o.TrustZone = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *Parent) GetComponent() string {
	if o == nil || IsNil(o.Component) {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Parent) GetComponentOk() (*string, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *Parent) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *Parent) SetComponent(v string) {
	o.Component = &v
}

func (o Parent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Parent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TrustZone) {
		toSerialize["trustZone"] = o.TrustZone
	}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	return toSerialize, nil
}

type NullableParent struct {
	value *Parent
	isSet bool
}

func (v NullableParent) Get() *Parent {
	return v.value
}

func (v *NullableParent) Set(val *Parent) {
	v.value = val
	v.isSet = true
}

func (v NullableParent) IsSet() bool {
	return v.isSet
}

func (v *NullableParent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParent(val *Parent) *NullableParent {
	return &NullableParent{value: val, isSet: true}
}

func (v NullableParent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


