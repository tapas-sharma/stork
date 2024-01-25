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

// RequestsCreateRestoreRequest struct for RequestsCreateRestoreRequest
type RequestsCreateRestoreRequest struct {
	DeploymentTargetId *string `json:"deployment_target_id,omitempty"`
	Name *string `json:"name,omitempty"`
	NamespaceId *string `json:"namespace_id,omitempty"`
}

// NewRequestsCreateRestoreRequest instantiates a new RequestsCreateRestoreRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestsCreateRestoreRequest() *RequestsCreateRestoreRequest {
	this := RequestsCreateRestoreRequest{}
	return &this
}

// NewRequestsCreateRestoreRequestWithDefaults instantiates a new RequestsCreateRestoreRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestsCreateRestoreRequestWithDefaults() *RequestsCreateRestoreRequest {
	this := RequestsCreateRestoreRequest{}
	return &this
}

// GetDeploymentTargetId returns the DeploymentTargetId field value if set, zero value otherwise.
func (o *RequestsCreateRestoreRequest) GetDeploymentTargetId() string {
	if o == nil || o.DeploymentTargetId == nil {
		var ret string
		return ret
	}
	return *o.DeploymentTargetId
}

// GetDeploymentTargetIdOk returns a tuple with the DeploymentTargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsCreateRestoreRequest) GetDeploymentTargetIdOk() (*string, bool) {
	if o == nil || o.DeploymentTargetId == nil {
		return nil, false
	}
	return o.DeploymentTargetId, true
}

// HasDeploymentTargetId returns a boolean if a field has been set.
func (o *RequestsCreateRestoreRequest) HasDeploymentTargetId() bool {
	if o != nil && o.DeploymentTargetId != nil {
		return true
	}

	return false
}

// SetDeploymentTargetId gets a reference to the given string and assigns it to the DeploymentTargetId field.
func (o *RequestsCreateRestoreRequest) SetDeploymentTargetId(v string) {
	o.DeploymentTargetId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RequestsCreateRestoreRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsCreateRestoreRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RequestsCreateRestoreRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RequestsCreateRestoreRequest) SetName(v string) {
	o.Name = &v
}

// GetNamespaceId returns the NamespaceId field value if set, zero value otherwise.
func (o *RequestsCreateRestoreRequest) GetNamespaceId() string {
	if o == nil || o.NamespaceId == nil {
		var ret string
		return ret
	}
	return *o.NamespaceId
}

// GetNamespaceIdOk returns a tuple with the NamespaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsCreateRestoreRequest) GetNamespaceIdOk() (*string, bool) {
	if o == nil || o.NamespaceId == nil {
		return nil, false
	}
	return o.NamespaceId, true
}

// HasNamespaceId returns a boolean if a field has been set.
func (o *RequestsCreateRestoreRequest) HasNamespaceId() bool {
	if o != nil && o.NamespaceId != nil {
		return true
	}

	return false
}

// SetNamespaceId gets a reference to the given string and assigns it to the NamespaceId field.
func (o *RequestsCreateRestoreRequest) SetNamespaceId(v string) {
	o.NamespaceId = &v
}

func (o RequestsCreateRestoreRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeploymentTargetId != nil {
		toSerialize["deployment_target_id"] = o.DeploymentTargetId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NamespaceId != nil {
		toSerialize["namespace_id"] = o.NamespaceId
	}
	return json.Marshal(toSerialize)
}

type NullableRequestsCreateRestoreRequest struct {
	value *RequestsCreateRestoreRequest
	isSet bool
}

func (v NullableRequestsCreateRestoreRequest) Get() *RequestsCreateRestoreRequest {
	return v.value
}

func (v *NullableRequestsCreateRestoreRequest) Set(val *RequestsCreateRestoreRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestsCreateRestoreRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestsCreateRestoreRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestsCreateRestoreRequest(val *RequestsCreateRestoreRequest) *NullableRequestsCreateRestoreRequest {
	return &NullableRequestsCreateRestoreRequest{value: val, isSet: true}
}

func (v NullableRequestsCreateRestoreRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestsCreateRestoreRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


