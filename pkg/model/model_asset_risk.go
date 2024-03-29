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

// checks if the AssetRisk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetRisk{}

// AssetRisk struct for AssetRisk
type AssetRisk struct {
	Confidentiality int `json:"confidentiality" yaml:"confidentiality"`
	Integrity int `json:"integrity" yaml:"integrity"`
	Availability int `json:"availability" yaml:"availability"`
	Comment *string `json:"comment,omitempty" yaml:"comment,omitempty"`
}

// NewAssetRisk instantiates a new AssetRisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetRisk(confidentiality int, integrity int, availability int) *AssetRisk {
	this := AssetRisk{}
	this.Confidentiality = confidentiality
	this.Integrity = integrity
	this.Availability = availability
	return &this
}

// NewAssetRiskWithDefaults instantiates a new AssetRisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetRiskWithDefaults() *AssetRisk {
	this := AssetRisk{}
	return &this
}

// GetConfidentiality returns the Confidentiality field value
func (o *AssetRisk) GetConfidentiality() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.Confidentiality
}

// GetConfidentialityOk returns a tuple with the Confidentiality field value
// and a boolean to check if the value has been set.
func (o *AssetRisk) GetConfidentialityOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Confidentiality, true
}

// SetConfidentiality sets field value
func (o *AssetRisk) SetConfidentiality(v int) {
	o.Confidentiality = v
}

// GetIntegrity returns the Integrity field value
func (o *AssetRisk) GetIntegrity() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.Integrity
}

// GetIntegrityOk returns a tuple with the Integrity field value
// and a boolean to check if the value has been set.
func (o *AssetRisk) GetIntegrityOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Integrity, true
}

// SetIntegrity sets field value
func (o *AssetRisk) SetIntegrity(v int) {
	o.Integrity = v
}

// GetAvailability returns the Availability field value
func (o *AssetRisk) GetAvailability() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.Availability
}

// GetAvailabilityOk returns a tuple with the Availability field value
// and a boolean to check if the value has been set.
func (o *AssetRisk) GetAvailabilityOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Availability, true
}

// SetAvailability sets field value
func (o *AssetRisk) SetAvailability(v int) {
	o.Availability = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *AssetRisk) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetRisk) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *AssetRisk) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *AssetRisk) SetComment(v string) {
	o.Comment = &v
}

func (o AssetRisk) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetRisk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["confidentiality"] = o.Confidentiality
	toSerialize["integrity"] = o.Integrity
	toSerialize["availability"] = o.Availability
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return toSerialize, nil
}

type NullableAssetRisk struct {
	value *AssetRisk
	isSet bool
}

func (v NullableAssetRisk) Get() *AssetRisk {
	return v.value
}

func (v *NullableAssetRisk) Set(val *AssetRisk) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetRisk) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetRisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetRisk(val *AssetRisk) *NullableAssetRisk {
	return &NullableAssetRisk{value: val, isSet: true}
}

func (v NullableAssetRisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetRisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


