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

// checks if the TrustZoneRisk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrustZoneRisk{}

// TrustZoneRisk struct for TrustZoneRisk
type TrustZoneRisk struct {
	TrustRating int `json:"trustRating" yaml:"trustRating"`
}

// NewTrustZoneRisk instantiates a new TrustZoneRisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrustZoneRisk(trustRating int) *TrustZoneRisk {
	this := TrustZoneRisk{}
	this.TrustRating = trustRating
	return &this
}

// NewTrustZoneRiskWithDefaults instantiates a new TrustZoneRisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrustZoneRiskWithDefaults() *TrustZoneRisk {
	this := TrustZoneRisk{}
	return &this
}

// GetTrustRating returns the TrustRating field value
func (o *TrustZoneRisk) GetTrustRating() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.TrustRating
}

// GetTrustRatingOk returns a tuple with the TrustRating field value
// and a boolean to check if the value has been set.
func (o *TrustZoneRisk) GetTrustRatingOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrustRating, true
}

// SetTrustRating sets field value
func (o *TrustZoneRisk) SetTrustRating(v int) {
	o.TrustRating = v
}

func (o TrustZoneRisk) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrustZoneRisk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["trustRating"] = o.TrustRating
	return toSerialize, nil
}

type NullableTrustZoneRisk struct {
	value *TrustZoneRisk
	isSet bool
}

func (v NullableTrustZoneRisk) Get() *TrustZoneRisk {
	return v.value
}

func (v *NullableTrustZoneRisk) Set(val *TrustZoneRisk) {
	v.value = val
	v.isSet = true
}

func (v NullableTrustZoneRisk) IsSet() bool {
	return v.isSet
}

func (v *NullableTrustZoneRisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrustZoneRisk(val *TrustZoneRisk) *NullableTrustZoneRisk {
	return &NullableTrustZoneRisk{value: val, isSet: true}
}

func (v NullableTrustZoneRisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrustZoneRisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


