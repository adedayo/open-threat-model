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

// checks if the TrustZoneParent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrustZoneParent{}

// TrustZoneParent struct for TrustZoneParent
type TrustZoneParent struct {
	TrustZone string `json:"trustZone" yaml:"trustZone"`
}

// NewTrustZoneParent instantiates a new TrustZoneParent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrustZoneParent(trustZone string) *TrustZoneParent {
	this := TrustZoneParent{}
	this.TrustZone = trustZone
	return &this
}

// NewTrustZoneParentWithDefaults instantiates a new TrustZoneParent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrustZoneParentWithDefaults() *TrustZoneParent {
	this := TrustZoneParent{}
	return &this
}

// GetTrustZone returns the TrustZone field value
func (o *TrustZoneParent) GetTrustZone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TrustZone
}

// GetTrustZoneOk returns a tuple with the TrustZone field value
// and a boolean to check if the value has been set.
func (o *TrustZoneParent) GetTrustZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrustZone, true
}

// SetTrustZone sets field value
func (o *TrustZoneParent) SetTrustZone(v string) {
	o.TrustZone = v
}

func (o TrustZoneParent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrustZoneParent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["trustZone"] = o.TrustZone
	return toSerialize, nil
}

type NullableTrustZoneParent struct {
	value *TrustZoneParent
	isSet bool
}

func (v NullableTrustZoneParent) Get() *TrustZoneParent {
	return v.value
}

func (v *NullableTrustZoneParent) Set(val *TrustZoneParent) {
	v.value = val
	v.isSet = true
}

func (v NullableTrustZoneParent) IsSet() bool {
	return v.isSet
}

func (v *NullableTrustZoneParent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrustZoneParent(val *TrustZoneParent) *NullableTrustZoneParent {
	return &NullableTrustZoneParent{value: val, isSet: true}
}

func (v NullableTrustZoneParent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrustZoneParent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


