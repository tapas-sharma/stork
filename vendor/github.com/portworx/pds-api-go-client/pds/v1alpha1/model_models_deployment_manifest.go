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

// ModelsDeploymentManifest struct for ModelsDeploymentManifest
type ModelsDeploymentManifest struct {
	// CreatedAt is autogenerated on creation
	CreatedAt *string `json:"created_at,omitempty"`
	Deployment *ModelsDeployment `json:"deployment,omitempty"`
	DeploymentId *string `json:"deployment_id,omitempty"`
	// Health represents the real state on the target cluster.
	Health *string `json:"health,omitempty"`
	// ID is auto generated on creation
	Id *string `json:"id,omitempty"`
	Manifest map[string]interface{} `json:"manifest,omitempty"`
	ReadyReplicas *int32 `json:"readyReplicas,omitempty"`
	Replicas *int32 `json:"replicas,omitempty"`
	// Status of the restore process on the Target Cluster
	RestoreStatus *string `json:"restore_status,omitempty"`
	// Restoring represents whether restore is happening or not.
	Restoring *bool `json:"restoring,omitempty"`
	// Status represents a granular state of the deployment. This includes transitional states like 'Deploying', which can only be inferred by the PDS control plane.
	Status *string `json:"status,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
	// UpdatedAt is autogenerated on update
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewModelsDeploymentManifest instantiates a new ModelsDeploymentManifest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsDeploymentManifest() *ModelsDeploymentManifest {
	this := ModelsDeploymentManifest{}
	return &this
}

// NewModelsDeploymentManifestWithDefaults instantiates a new ModelsDeploymentManifest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsDeploymentManifestWithDefaults() *ModelsDeploymentManifest {
	this := ModelsDeploymentManifest{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ModelsDeploymentManifest) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentManifest) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ModelsDeploymentManifest) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ModelsDeploymentManifest) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDeployment returns the Deployment field value if set, zero value otherwise.
func (o *ModelsDeploymentManifest) GetDeployment() ModelsDeployment {
	if o == nil || o.Deployment == nil {
		var ret ModelsDeployment
		return ret
	}
	return *o.Deployment
}

// GetDeploymentOk returns a tuple with the Deployment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentManifest) GetDeploymentOk() (*ModelsDeployment, bool) {
	if o == nil || o.Deployment == nil {
		return nil, false
	}
	return o.Deployment, true
}

// HasDeployment returns a boolean if a field has been set.
func (o *ModelsDeploymentManifest) HasDeployment() bool {
	if o != nil && o.Deployment != nil {
		return true
	}

	return false
}

// SetDeployment gets a reference to the given ModelsDeployment and assigns it to the Deployment field.
func (o *ModelsDeploymentManifest) SetDeployment(v ModelsDeployment) {
	o.Deployment = &v
}

// GetDeploymentId returns the DeploymentId field value if set, zero value otherwise.
func (o *ModelsDeploymentManifest) GetDeploymentId() string {
	if o == nil || o.DeploymentId == nil {
		var ret string
		return ret
	}
	return *o.DeploymentId
}

// GetDeploymentIdOk returns a tuple with the DeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentManifest) GetDeploymentIdOk() (*string, bool) {
	if o == nil || o.DeploymentId == nil {
		return nil, false
	}
	return o.DeploymentId, true
}

// HasDeploymentId returns a boolean if a field has been set.
func (o *ModelsDeploymentManifest) HasDeploymentId() bool {
	if o != nil && o.DeploymentId != nil {
		return true
	}

	return false
}

// SetDeploymentId gets a reference to the given string and assigns it to the DeploymentId field.
func (o *ModelsDeploymentManifest) SetDeploymentId(v string) {
	o.DeploymentId = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *ModelsDeploymentManifest) GetHealth() string {
	if o == nil || o.Health == nil {
		var ret string
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentManifest) GetHealthOk() (*string, bool) {
	if o == nil || o.Health == nil {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *ModelsDeploymentManifest) HasHealth() bool {
	if o != nil && o.Health != nil {
		return true
	}

	return false
}

// SetHealth gets a reference to the given string and assigns it to the Health field.
func (o *ModelsDeploymentManifest) SetHealth(v string) {
	o.Health = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsDeploymentManifest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentManifest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsDeploymentManifest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelsDeploymentManifest) SetId(v string) {
	o.Id = &v
}

// GetManifest returns the Manifest field value if set, zero value otherwise.
func (o *ModelsDeploymentManifest) GetManifest() map[string]interface{} {
	if o == nil || o.Manifest == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Manifest
}

// GetManifestOk returns a tuple with the Manifest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentManifest) GetManifestOk() (map[string]interface{}, bool) {
	if o == nil || o.Manifest == nil {
		return nil, false
	}
	return o.Manifest, true
}

// HasManifest returns a boolean if a field has been set.
func (o *ModelsDeploymentManifest) HasManifest() bool {
	if o != nil && o.Manifest != nil {
		return true
	}

	return false
}

// SetManifest gets a reference to the given map[string]interface{} and assigns it to the Manifest field.
func (o *ModelsDeploymentManifest) SetManifest(v map[string]interface{}) {
	o.Manifest = v
}

// GetReadyReplicas returns the ReadyReplicas field value if set, zero value otherwise.
func (o *ModelsDeploymentManifest) GetReadyReplicas() int32 {
	if o == nil || o.ReadyReplicas == nil {
		var ret int32
		return ret
	}
	return *o.ReadyReplicas
}

// GetReadyReplicasOk returns a tuple with the ReadyReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentManifest) GetReadyReplicasOk() (*int32, bool) {
	if o == nil || o.ReadyReplicas == nil {
		return nil, false
	}
	return o.ReadyReplicas, true
}

// HasReadyReplicas returns a boolean if a field has been set.
func (o *ModelsDeploymentManifest) HasReadyReplicas() bool {
	if o != nil && o.ReadyReplicas != nil {
		return true
	}

	return false
}

// SetReadyReplicas gets a reference to the given int32 and assigns it to the ReadyReplicas field.
func (o *ModelsDeploymentManifest) SetReadyReplicas(v int32) {
	o.ReadyReplicas = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *ModelsDeploymentManifest) GetReplicas() int32 {
	if o == nil || o.Replicas == nil {
		var ret int32
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentManifest) GetReplicasOk() (*int32, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *ModelsDeploymentManifest) HasReplicas() bool {
	if o != nil && o.Replicas != nil {
		return true
	}

	return false
}

// SetReplicas gets a reference to the given int32 and assigns it to the Replicas field.
func (o *ModelsDeploymentManifest) SetReplicas(v int32) {
	o.Replicas = &v
}

// GetRestoreStatus returns the RestoreStatus field value if set, zero value otherwise.
func (o *ModelsDeploymentManifest) GetRestoreStatus() string {
	if o == nil || o.RestoreStatus == nil {
		var ret string
		return ret
	}
	return *o.RestoreStatus
}

// GetRestoreStatusOk returns a tuple with the RestoreStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentManifest) GetRestoreStatusOk() (*string, bool) {
	if o == nil || o.RestoreStatus == nil {
		return nil, false
	}
	return o.RestoreStatus, true
}

// HasRestoreStatus returns a boolean if a field has been set.
func (o *ModelsDeploymentManifest) HasRestoreStatus() bool {
	if o != nil && o.RestoreStatus != nil {
		return true
	}

	return false
}

// SetRestoreStatus gets a reference to the given string and assigns it to the RestoreStatus field.
func (o *ModelsDeploymentManifest) SetRestoreStatus(v string) {
	o.RestoreStatus = &v
}

// GetRestoring returns the Restoring field value if set, zero value otherwise.
func (o *ModelsDeploymentManifest) GetRestoring() bool {
	if o == nil || o.Restoring == nil {
		var ret bool
		return ret
	}
	return *o.Restoring
}

// GetRestoringOk returns a tuple with the Restoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentManifest) GetRestoringOk() (*bool, bool) {
	if o == nil || o.Restoring == nil {
		return nil, false
	}
	return o.Restoring, true
}

// HasRestoring returns a boolean if a field has been set.
func (o *ModelsDeploymentManifest) HasRestoring() bool {
	if o != nil && o.Restoring != nil {
		return true
	}

	return false
}

// SetRestoring gets a reference to the given bool and assigns it to the Restoring field.
func (o *ModelsDeploymentManifest) SetRestoring(v bool) {
	o.Restoring = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ModelsDeploymentManifest) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentManifest) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ModelsDeploymentManifest) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ModelsDeploymentManifest) SetStatus(v string) {
	o.Status = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ModelsDeploymentManifest) GetTimestamp() string {
	if o == nil || o.Timestamp == nil {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentManifest) GetTimestampOk() (*string, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ModelsDeploymentManifest) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *ModelsDeploymentManifest) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ModelsDeploymentManifest) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsDeploymentManifest) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ModelsDeploymentManifest) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ModelsDeploymentManifest) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o ModelsDeploymentManifest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Deployment != nil {
		toSerialize["deployment"] = o.Deployment
	}
	if o.DeploymentId != nil {
		toSerialize["deployment_id"] = o.DeploymentId
	}
	if o.Health != nil {
		toSerialize["health"] = o.Health
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Manifest != nil {
		toSerialize["manifest"] = o.Manifest
	}
	if o.ReadyReplicas != nil {
		toSerialize["readyReplicas"] = o.ReadyReplicas
	}
	if o.Replicas != nil {
		toSerialize["replicas"] = o.Replicas
	}
	if o.RestoreStatus != nil {
		toSerialize["restore_status"] = o.RestoreStatus
	}
	if o.Restoring != nil {
		toSerialize["restoring"] = o.Restoring
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableModelsDeploymentManifest struct {
	value *ModelsDeploymentManifest
	isSet bool
}

func (v NullableModelsDeploymentManifest) Get() *ModelsDeploymentManifest {
	return v.value
}

func (v *NullableModelsDeploymentManifest) Set(val *ModelsDeploymentManifest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsDeploymentManifest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsDeploymentManifest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsDeploymentManifest(val *ModelsDeploymentManifest) *NullableModelsDeploymentManifest {
	return &NullableModelsDeploymentManifest{value: val, isSet: true}
}

func (v NullableModelsDeploymentManifest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsDeploymentManifest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


