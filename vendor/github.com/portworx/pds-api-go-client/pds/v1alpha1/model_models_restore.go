/*
PDS API

Portworx Data Services API Server

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pds

import (
	"encoding/json"
)

// ModelsRestore struct for ModelsRestore
type ModelsRestore struct {
	// Status of sending Restore CR to the Target Cluster
	ApplyStatus *string `json:"apply_status,omitempty"`
	// ID of the backup to be restored, the job might be already deleted
	BackupJobId *string `json:"backup_job_id,omitempty"`
	// Identifier of the PX cloud credentials of the storage with the backup
	CloudCredentialName *string `json:"cloud_credential_name,omitempty"`
	// ID of the PX cloud snapshot with the backup
	CloudSnapId *string `json:"cloud_snap_id,omitempty"`
	// ClusterResourceName k8s resource name for deployment, built from [restore + name + id].
	ClusterResourceName *string `json:"cluster_resource_name,omitempty"`
	// Completion time of the restore process
	CompletionTime *string `json:"completion_time,omitempty"`
	// CreatedAt is autogenerated on creation
	CreatedAt *string `json:"created_at,omitempty"`
	// ID of the new deployment that will contain restored data
	DeploymentId *string `json:"deployment_id,omitempty"`
	// ID of the deployment target where the restore is applied
	DeploymentTargetId *string `json:"deployment_target_id,omitempty"`
	// Error code of the restore from Target Cluster
	ErrorCode *string `json:"error_code,omitempty"`
	// ID is auto generated on creation
	Id *string `json:"id,omitempty"`
	// Name of the restore
	Name *string `json:"name,omitempty"`
	// ID of the namespace where the restore is created
	NamespaceId *string `json:"namespace_id,omitempty"`
	// Starting time of the restore process
	StartTime *string `json:"start_time,omitempty"`
	// Status of the restore process on the Target Cluster
	Status *string `json:"status,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
	// UpdatedAt is autogenerated on update
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewModelsRestore instantiates a new ModelsRestore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsRestore() *ModelsRestore {
	this := ModelsRestore{}
	return &this
}

// NewModelsRestoreWithDefaults instantiates a new ModelsRestore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsRestoreWithDefaults() *ModelsRestore {
	this := ModelsRestore{}
	return &this
}

// GetApplyStatus returns the ApplyStatus field value if set, zero value otherwise.
func (o *ModelsRestore) GetApplyStatus() string {
	if o == nil || o.ApplyStatus == nil {
		var ret string
		return ret
	}
	return *o.ApplyStatus
}

// GetApplyStatusOk returns a tuple with the ApplyStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRestore) GetApplyStatusOk() (*string, bool) {
	if o == nil || o.ApplyStatus == nil {
		return nil, false
	}
	return o.ApplyStatus, true
}

// HasApplyStatus returns a boolean if a field has been set.
func (o *ModelsRestore) HasApplyStatus() bool {
	if o != nil && o.ApplyStatus != nil {
		return true
	}

	return false
}

// SetApplyStatus gets a reference to the given string and assigns it to the ApplyStatus field.
func (o *ModelsRestore) SetApplyStatus(v string) {
	o.ApplyStatus = &v
}

// GetBackupJobId returns the BackupJobId field value if set, zero value otherwise.
func (o *ModelsRestore) GetBackupJobId() string {
	if o == nil || o.BackupJobId == nil {
		var ret string
		return ret
	}
	return *o.BackupJobId
}

// GetBackupJobIdOk returns a tuple with the BackupJobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRestore) GetBackupJobIdOk() (*string, bool) {
	if o == nil || o.BackupJobId == nil {
		return nil, false
	}
	return o.BackupJobId, true
}

// HasBackupJobId returns a boolean if a field has been set.
func (o *ModelsRestore) HasBackupJobId() bool {
	if o != nil && o.BackupJobId != nil {
		return true
	}

	return false
}

// SetBackupJobId gets a reference to the given string and assigns it to the BackupJobId field.
func (o *ModelsRestore) SetBackupJobId(v string) {
	o.BackupJobId = &v
}

// GetCloudCredentialName returns the CloudCredentialName field value if set, zero value otherwise.
func (o *ModelsRestore) GetCloudCredentialName() string {
	if o == nil || o.CloudCredentialName == nil {
		var ret string
		return ret
	}
	return *o.CloudCredentialName
}

// GetCloudCredentialNameOk returns a tuple with the CloudCredentialName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRestore) GetCloudCredentialNameOk() (*string, bool) {
	if o == nil || o.CloudCredentialName == nil {
		return nil, false
	}
	return o.CloudCredentialName, true
}

// HasCloudCredentialName returns a boolean if a field has been set.
func (o *ModelsRestore) HasCloudCredentialName() bool {
	if o != nil && o.CloudCredentialName != nil {
		return true
	}

	return false
}

// SetCloudCredentialName gets a reference to the given string and assigns it to the CloudCredentialName field.
func (o *ModelsRestore) SetCloudCredentialName(v string) {
	o.CloudCredentialName = &v
}

// GetCloudSnapId returns the CloudSnapId field value if set, zero value otherwise.
func (o *ModelsRestore) GetCloudSnapId() string {
	if o == nil || o.CloudSnapId == nil {
		var ret string
		return ret
	}
	return *o.CloudSnapId
}

// GetCloudSnapIdOk returns a tuple with the CloudSnapId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRestore) GetCloudSnapIdOk() (*string, bool) {
	if o == nil || o.CloudSnapId == nil {
		return nil, false
	}
	return o.CloudSnapId, true
}

// HasCloudSnapId returns a boolean if a field has been set.
func (o *ModelsRestore) HasCloudSnapId() bool {
	if o != nil && o.CloudSnapId != nil {
		return true
	}

	return false
}

// SetCloudSnapId gets a reference to the given string and assigns it to the CloudSnapId field.
func (o *ModelsRestore) SetCloudSnapId(v string) {
	o.CloudSnapId = &v
}

// GetClusterResourceName returns the ClusterResourceName field value if set, zero value otherwise.
func (o *ModelsRestore) GetClusterResourceName() string {
	if o == nil || o.ClusterResourceName == nil {
		var ret string
		return ret
	}
	return *o.ClusterResourceName
}

// GetClusterResourceNameOk returns a tuple with the ClusterResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRestore) GetClusterResourceNameOk() (*string, bool) {
	if o == nil || o.ClusterResourceName == nil {
		return nil, false
	}
	return o.ClusterResourceName, true
}

// HasClusterResourceName returns a boolean if a field has been set.
func (o *ModelsRestore) HasClusterResourceName() bool {
	if o != nil && o.ClusterResourceName != nil {
		return true
	}

	return false
}

// SetClusterResourceName gets a reference to the given string and assigns it to the ClusterResourceName field.
func (o *ModelsRestore) SetClusterResourceName(v string) {
	o.ClusterResourceName = &v
}

// GetCompletionTime returns the CompletionTime field value if set, zero value otherwise.
func (o *ModelsRestore) GetCompletionTime() string {
	if o == nil || o.CompletionTime == nil {
		var ret string
		return ret
	}
	return *o.CompletionTime
}

// GetCompletionTimeOk returns a tuple with the CompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRestore) GetCompletionTimeOk() (*string, bool) {
	if o == nil || o.CompletionTime == nil {
		return nil, false
	}
	return o.CompletionTime, true
}

// HasCompletionTime returns a boolean if a field has been set.
func (o *ModelsRestore) HasCompletionTime() bool {
	if o != nil && o.CompletionTime != nil {
		return true
	}

	return false
}

// SetCompletionTime gets a reference to the given string and assigns it to the CompletionTime field.
func (o *ModelsRestore) SetCompletionTime(v string) {
	o.CompletionTime = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ModelsRestore) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRestore) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ModelsRestore) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ModelsRestore) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDeploymentId returns the DeploymentId field value if set, zero value otherwise.
func (o *ModelsRestore) GetDeploymentId() string {
	if o == nil || o.DeploymentId == nil {
		var ret string
		return ret
	}
	return *o.DeploymentId
}

// GetDeploymentIdOk returns a tuple with the DeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRestore) GetDeploymentIdOk() (*string, bool) {
	if o == nil || o.DeploymentId == nil {
		return nil, false
	}
	return o.DeploymentId, true
}

// HasDeploymentId returns a boolean if a field has been set.
func (o *ModelsRestore) HasDeploymentId() bool {
	if o != nil && o.DeploymentId != nil {
		return true
	}

	return false
}

// SetDeploymentId gets a reference to the given string and assigns it to the DeploymentId field.
func (o *ModelsRestore) SetDeploymentId(v string) {
	o.DeploymentId = &v
}

// GetDeploymentTargetId returns the DeploymentTargetId field value if set, zero value otherwise.
func (o *ModelsRestore) GetDeploymentTargetId() string {
	if o == nil || o.DeploymentTargetId == nil {
		var ret string
		return ret
	}
	return *o.DeploymentTargetId
}

// GetDeploymentTargetIdOk returns a tuple with the DeploymentTargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRestore) GetDeploymentTargetIdOk() (*string, bool) {
	if o == nil || o.DeploymentTargetId == nil {
		return nil, false
	}
	return o.DeploymentTargetId, true
}

// HasDeploymentTargetId returns a boolean if a field has been set.
func (o *ModelsRestore) HasDeploymentTargetId() bool {
	if o != nil && o.DeploymentTargetId != nil {
		return true
	}

	return false
}

// SetDeploymentTargetId gets a reference to the given string and assigns it to the DeploymentTargetId field.
func (o *ModelsRestore) SetDeploymentTargetId(v string) {
	o.DeploymentTargetId = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *ModelsRestore) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRestore) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *ModelsRestore) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *ModelsRestore) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsRestore) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRestore) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsRestore) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelsRestore) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelsRestore) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRestore) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelsRestore) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelsRestore) SetName(v string) {
	o.Name = &v
}

// GetNamespaceId returns the NamespaceId field value if set, zero value otherwise.
func (o *ModelsRestore) GetNamespaceId() string {
	if o == nil || o.NamespaceId == nil {
		var ret string
		return ret
	}
	return *o.NamespaceId
}

// GetNamespaceIdOk returns a tuple with the NamespaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRestore) GetNamespaceIdOk() (*string, bool) {
	if o == nil || o.NamespaceId == nil {
		return nil, false
	}
	return o.NamespaceId, true
}

// HasNamespaceId returns a boolean if a field has been set.
func (o *ModelsRestore) HasNamespaceId() bool {
	if o != nil && o.NamespaceId != nil {
		return true
	}

	return false
}

// SetNamespaceId gets a reference to the given string and assigns it to the NamespaceId field.
func (o *ModelsRestore) SetNamespaceId(v string) {
	o.NamespaceId = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ModelsRestore) GetStartTime() string {
	if o == nil || o.StartTime == nil {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRestore) GetStartTimeOk() (*string, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ModelsRestore) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *ModelsRestore) SetStartTime(v string) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ModelsRestore) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRestore) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ModelsRestore) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ModelsRestore) SetStatus(v string) {
	o.Status = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ModelsRestore) GetTimestamp() string {
	if o == nil || o.Timestamp == nil {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRestore) GetTimestampOk() (*string, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ModelsRestore) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *ModelsRestore) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ModelsRestore) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsRestore) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ModelsRestore) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ModelsRestore) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o ModelsRestore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplyStatus != nil {
		toSerialize["apply_status"] = o.ApplyStatus
	}
	if o.BackupJobId != nil {
		toSerialize["backup_job_id"] = o.BackupJobId
	}
	if o.CloudCredentialName != nil {
		toSerialize["cloud_credential_name"] = o.CloudCredentialName
	}
	if o.CloudSnapId != nil {
		toSerialize["cloud_snap_id"] = o.CloudSnapId
	}
	if o.ClusterResourceName != nil {
		toSerialize["cluster_resource_name"] = o.ClusterResourceName
	}
	if o.CompletionTime != nil {
		toSerialize["completion_time"] = o.CompletionTime
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.DeploymentId != nil {
		toSerialize["deployment_id"] = o.DeploymentId
	}
	if o.DeploymentTargetId != nil {
		toSerialize["deployment_target_id"] = o.DeploymentTargetId
	}
	if o.ErrorCode != nil {
		toSerialize["error_code"] = o.ErrorCode
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NamespaceId != nil {
		toSerialize["namespace_id"] = o.NamespaceId
	}
	if o.StartTime != nil {
		toSerialize["start_time"] = o.StartTime
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

type NullableModelsRestore struct {
	value *ModelsRestore
	isSet bool
}

func (v NullableModelsRestore) Get() *ModelsRestore {
	return v.value
}

func (v *NullableModelsRestore) Set(val *ModelsRestore) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsRestore) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsRestore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsRestore(val *ModelsRestore) *NullableModelsRestore {
	return &NullableModelsRestore{value: val, isSet: true}
}

func (v NullableModelsRestore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsRestore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


