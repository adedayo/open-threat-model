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

// checks if the TrustZone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrustZone{}

// TrustZone struct for TrustZone
type TrustZone struct {
	Id string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
	Type string `json:"type" yaml:"type"`
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
	Risk TrustZoneRisk `json:"risk" yaml:"risk"`
	Parent *Parent `json:"parent,omitempty" yaml:"parent,omitempty"`
	Representations []TrustZoneRepresentationsInner `json:"representations,omitempty" yaml:"representations,omitempty"`
	Attributes *map[string]string `json:"attributes,omitempty" yaml:"attributes,omitempty"`
}

// NewTrustZone instantiates a new TrustZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrustZone(id string, name string, type_ string, risk TrustZoneRisk) *TrustZone {
	this := TrustZone{}
	this.Id = id
	this.Name = name
	this.Type = type_
	this.Risk = risk
	return &this
}

// NewTrustZoneWithDefaults instantiates a new TrustZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrustZoneWithDefaults() *TrustZone {
	this := TrustZone{}
	return &this
}

// GetId returns the Id field value
func (o *TrustZone) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TrustZone) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TrustZone) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TrustZone) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TrustZone) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TrustZone) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *TrustZone) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TrustZone) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TrustZone) SetType(v string) {
	o.Type = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TrustZone) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustZone) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TrustZone) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TrustZone) SetDescription(v string) {
	o.Description = &v
}

// GetRisk returns the Risk field value
func (o *TrustZone) GetRisk() TrustZoneRisk {
	if o == nil {
		var ret TrustZoneRisk
		return ret
	}

	return o.Risk
}

// GetRiskOk returns a tuple with the Risk field value
// and a boolean to check if the value has been set.
func (o *TrustZone) GetRiskOk() (*TrustZoneRisk, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Risk, true
}

// SetRisk sets field value
func (o *TrustZone) SetRisk(v TrustZoneRisk) {
	o.Risk = v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *TrustZone) GetParent() Parent {
	if o == nil || IsNil(o.Parent) {
		var ret Parent
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustZone) GetParentOk() (*Parent, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *TrustZone) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given Parent and assigns it to the Parent field.
func (o *TrustZone) SetParent(v Parent) {
	o.Parent = &v
}

// GetRepresentations returns the Representations field value if set, zero value otherwise.
func (o *TrustZone) GetRepresentations() []TrustZoneRepresentationsInner {
	if o == nil || IsNil(o.Representations) {
		var ret []TrustZoneRepresentationsInner
		return ret
	}
	return o.Representations
}

// GetRepresentationsOk returns a tuple with the Representations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustZone) GetRepresentationsOk() ([]TrustZoneRepresentationsInner, bool) {
	if o == nil || IsNil(o.Representations) {
		return nil, false
	}
	return o.Representations, true
}

// HasRepresentations returns a boolean if a field has been set.
func (o *TrustZone) HasRepresentations() bool {
	if o != nil && !IsNil(o.Representations) {
		return true
	}

	return false
}

// SetRepresentations gets a reference to the given []TrustZoneRepresentationsInner and assigns it to the Representations field.
func (o *TrustZone) SetRepresentations(v []TrustZoneRepresentationsInner) {
	o.Representations = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TrustZone) GetAttributes() map[string]string {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustZone) GetAttributesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TrustZone) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]string and assigns it to the Attributes field.
func (o *TrustZone) SetAttributes(v map[string]string) {
	o.Attributes = &v
}

func (o TrustZone) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrustZone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["risk"] = o.Risk
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.Representations) {
		toSerialize["representations"] = o.Representations
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableTrustZone struct {
	value *TrustZone
	isSet bool
}

func (v NullableTrustZone) Get() *TrustZone {
	return v.value
}

func (v *NullableTrustZone) Set(val *TrustZone) {
	v.value = val
	v.isSet = true
}

func (v NullableTrustZone) IsSet() bool {
	return v.isSet
}

func (v *NullableTrustZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrustZone(val *TrustZone) *NullableTrustZone {
	return &NullableTrustZone{value: val, isSet: true}
}

func (v NullableTrustZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrustZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

