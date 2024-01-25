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

// ControllersPartialGoogleCredentials struct for ControllersPartialGoogleCredentials
type ControllersPartialGoogleCredentials struct {
	// Project ID of the Google Cloud project.
	ProjectId *string `json:"project_id,omitempty"`
}

// NewControllersPartialGoogleCredentials instantiates a new ControllersPartialGoogleCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersPartialGoogleCredentials() *ControllersPartialGoogleCredentials {
	this := ControllersPartialGoogleCredentials{}
	return &this
}

// NewControllersPartialGoogleCredentialsWithDefaults instantiates a new ControllersPartialGoogleCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersPartialGoogleCredentialsWithDefaults() *ControllersPartialGoogleCredentials {
	this := ControllersPartialGoogleCredentials{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ControllersPartialGoogleCredentials) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersPartialGoogleCredentials) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ControllersPartialGoogleCredentials) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *ControllersPartialGoogleCredentials) SetProjectId(v string) {
	o.ProjectId = &v
}

func (o ControllersPartialGoogleCredentials) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProjectId != nil {
		toSerialize["project_id"] = o.ProjectId
	}
	return json.Marshal(toSerialize)
}

type NullableControllersPartialGoogleCredentials struct {
	value *ControllersPartialGoogleCredentials
	isSet bool
}

func (v NullableControllersPartialGoogleCredentials) Get() *ControllersPartialGoogleCredentials {
	return v.value
}

func (v *NullableControllersPartialGoogleCredentials) Set(val *ControllersPartialGoogleCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersPartialGoogleCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersPartialGoogleCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersPartialGoogleCredentials(val *ControllersPartialGoogleCredentials) *NullableControllersPartialGoogleCredentials {
	return &NullableControllersPartialGoogleCredentials{value: val, isSet: true}
}

func (v NullableControllersPartialGoogleCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersPartialGoogleCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


