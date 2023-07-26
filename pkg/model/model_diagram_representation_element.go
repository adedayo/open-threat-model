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

// checks if the DiagramRepresentationElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiagramRepresentationElement{}

// DiagramRepresentationElement struct for DiagramRepresentationElement
type DiagramRepresentationElement struct {
	Representation string `json:"representation" yaml:"representation"`
	Id string `json:"id" yaml:"id"`
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`
	Position *Position `json:"position,omitempty" yaml:"position,omitempty"`
	Size *Size `json:"size,omitempty" yaml:"size,omitempty"`
	Attributes *map[string]string `json:"attributes,omitempty" yaml:"attributes,omitempty"`
}

// NewDiagramRepresentationElement instantiates a new DiagramRepresentationElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiagramRepresentationElement(representation string, id string) *DiagramRepresentationElement {
	this := DiagramRepresentationElement{}
	this.Representation = representation
	this.Id = id
	return &this
}

// NewDiagramRepresentationElementWithDefaults instantiates a new DiagramRepresentationElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiagramRepresentationElementWithDefaults() *DiagramRepresentationElement {
	this := DiagramRepresentationElement{}
	return &this
}

// GetRepresentation returns the Representation field value
func (o *DiagramRepresentationElement) GetRepresentation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Representation
}

// GetRepresentationOk returns a tuple with the Representation field value
// and a boolean to check if the value has been set.
func (o *DiagramRepresentationElement) GetRepresentationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Representation, true
}

// SetRepresentation sets field value
func (o *DiagramRepresentationElement) SetRepresentation(v string) {
	o.Representation = v
}

// GetId returns the Id field value
func (o *DiagramRepresentationElement) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DiagramRepresentationElement) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DiagramRepresentationElement) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DiagramRepresentationElement) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagramRepresentationElement) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DiagramRepresentationElement) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DiagramRepresentationElement) SetName(v string) {
	o.Name = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *DiagramRepresentationElement) GetPosition() Position {
	if o == nil || IsNil(o.Position) {
		var ret Position
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagramRepresentationElement) GetPositionOk() (*Position, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *DiagramRepresentationElement) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given Position and assigns it to the Position field.
func (o *DiagramRepresentationElement) SetPosition(v Position) {
	o.Position = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *DiagramRepresentationElement) GetSize() Size {
	if o == nil || IsNil(o.Size) {
		var ret Size
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagramRepresentationElement) GetSizeOk() (*Size, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *DiagramRepresentationElement) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given Size and assigns it to the Size field.
func (o *DiagramRepresentationElement) SetSize(v Size) {
	o.Size = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *DiagramRepresentationElement) GetAttributes() map[string]string {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagramRepresentationElement) GetAttributesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *DiagramRepresentationElement) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]string and assigns it to the Attributes field.
func (o *DiagramRepresentationElement) SetAttributes(v map[string]string) {
	o.Attributes = &v
}

func (o DiagramRepresentationElement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiagramRepresentationElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["representation"] = o.Representation
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableDiagramRepresentationElement struct {
	value *DiagramRepresentationElement
	isSet bool
}

func (v NullableDiagramRepresentationElement) Get() *DiagramRepresentationElement {
	return v.value
}

func (v *NullableDiagramRepresentationElement) Set(val *DiagramRepresentationElement) {
	v.value = val
	v.isSet = true
}

func (v NullableDiagramRepresentationElement) IsSet() bool {
	return v.isSet
}

func (v *NullableDiagramRepresentationElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiagramRepresentationElement(val *DiagramRepresentationElement) *NullableDiagramRepresentationElement {
	return &NullableDiagramRepresentationElement{value: val, isSet: true}
}

func (v NullableDiagramRepresentationElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiagramRepresentationElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


