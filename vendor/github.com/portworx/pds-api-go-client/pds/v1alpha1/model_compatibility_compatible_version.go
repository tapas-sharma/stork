/*
PDS API

Portworx Data Services API Server

API version: 121
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pds

import (
	"encoding/json"
)

// CompatibilityCompatibleVersion struct for CompatibilityCompatibleVersion
type CompatibilityCompatibleVersion struct {
	// ID of a Version entity.
	Id *string `json:"id,omitempty"`
	// Name of a Version entity.
	Name *string `json:"name,omitempty"`
}

// NewCompatibilityCompatibleVersion instantiates a new CompatibilityCompatibleVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompatibilityCompatibleVersion() *CompatibilityCompatibleVersion {
	this := CompatibilityCompatibleVersion{}
	return &this
}

// NewCompatibilityCompatibleVersionWithDefaults instantiates a new CompatibilityCompatibleVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompatibilityCompatibleVersionWithDefaults() *CompatibilityCompatibleVersion {
	this := CompatibilityCompatibleVersion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CompatibilityCompatibleVersion) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompatibilityCompatibleVersion) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CompatibilityCompatibleVersion) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CompatibilityCompatibleVersion) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CompatibilityCompatibleVersion) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompatibilityCompatibleVersion) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CompatibilityCompatibleVersion) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CompatibilityCompatibleVersion) SetName(v string) {
	o.Name = &v
}

func (o CompatibilityCompatibleVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableCompatibilityCompatibleVersion struct {
	value *CompatibilityCompatibleVersion
	isSet bool
}

func (v NullableCompatibilityCompatibleVersion) Get() *CompatibilityCompatibleVersion {
	return v.value
}

func (v *NullableCompatibilityCompatibleVersion) Set(val *CompatibilityCompatibleVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableCompatibilityCompatibleVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableCompatibilityCompatibleVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompatibilityCompatibleVersion(val *CompatibilityCompatibleVersion) *NullableCompatibilityCompatibleVersion {
	return &NullableCompatibilityCompatibleVersion{value: val, isSet: true}
}

func (v NullableCompatibilityCompatibleVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompatibilityCompatibleVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


