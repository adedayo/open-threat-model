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

// checks if the CodeRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CodeRepresentation{}

// CodeRepresentation struct for CodeRepresentation
type CodeRepresentation struct {
	Name string `json:"name" yaml:"name"`
	Id string `json:"id" yaml:"id"`
	Type string `json:"type" yaml:"type"`
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
	Attributes *map[string]string `json:"attributes,omitempty" yaml:"attributes,omitempty"`
	Repository *Repository `json:"repository,omitempty" yaml:"repository,omitempty"`
}

// NewCodeRepresentation instantiates a new CodeRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodeRepresentation(name string, id string, type_ string) *CodeRepresentation {
	this := CodeRepresentation{}
	this.Name = name
	this.Id = id
	this.Type = type_
	return &this
}

// NewCodeRepresentationWithDefaults instantiates a new CodeRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodeRepresentationWithDefaults() *CodeRepresentation {
	this := CodeRepresentation{}
	return &this
}

// GetName returns the Name field value
func (o *CodeRepresentation) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CodeRepresentation) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CodeRepresentation) SetName(v string) {
	o.Name = v
}

// GetId returns the Id field value
func (o *CodeRepresentation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CodeRepresentation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CodeRepresentation) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *CodeRepresentation) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CodeRepresentation) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CodeRepresentation) SetType(v string) {
	o.Type = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CodeRepresentation) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodeRepresentation) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CodeRepresentation) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CodeRepresentation) SetDescription(v string) {
	o.Description = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CodeRepresentation) GetAttributes() map[string]string {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodeRepresentation) GetAttributesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CodeRepresentation) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]string and assigns it to the Attributes field.
func (o *CodeRepresentation) SetAttributes(v map[string]string) {
	o.Attributes = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *CodeRepresentation) GetRepository() Repository {
	if o == nil || IsNil(o.Repository) {
		var ret Repository
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodeRepresentation) GetRepositoryOk() (*Repository, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *CodeRepresentation) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given Repository and assigns it to the Repository field.
func (o *CodeRepresentation) SetRepository(v Repository) {
	o.Repository = &v
}

func (o CodeRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CodeRepresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	return toSerialize, nil
}

type NullableCodeRepresentation struct {
	value *CodeRepresentation
	isSet bool
}

func (v NullableCodeRepresentation) Get() *CodeRepresentation {
	return v.value
}

func (v *NullableCodeRepresentation) Set(val *CodeRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullableCodeRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullableCodeRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodeRepresentation(val *CodeRepresentation) *NullableCodeRepresentation {
	return &NullableCodeRepresentation{value: val, isSet: true}
}

func (v NullableCodeRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodeRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

