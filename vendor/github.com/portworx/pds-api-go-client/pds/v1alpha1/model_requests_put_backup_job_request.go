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

// RequestsPutBackupJobRequest struct for RequestsPutBackupJobRequest
type RequestsPutBackupJobRequest struct {
	BackupCapability *string `json:"backup_capability,omitempty"`
	BackupId *string `json:"backup_id,omitempty"`
	BackupSpec map[string]interface{} `json:"backup_spec,omitempty"`
	CloudCredentialName *string `json:"cloud_credential_name,omitempty"`
	CloudSnapId *string `json:"cloud_snap_id,omitempty"`
	// CompletionStatus of the snapshot.
	CompletionStatus *string `json:"completion_status,omitempty"`
	CompletionTime *string `json:"completion_time,omitempty"`
	DataServiceSpec map[string]interface{} `json:"data_service_spec,omitempty"`
	Deleted *bool `json:"deleted,omitempty"`
	DeploymentId *string `json:"deployment_id,omitempty"`
	DeploymentTargetId *string `json:"deployment_target_id,omitempty"`
	// ErrorCode if CompletionStatus is \"Failed\"
	ErrorCode *string `json:"error_code,omitempty"`
	ErrorMessage *string `json:"error_message,omitempty"`
	FileSize *int32 `json:"file_size,omitempty"`
	ImageName *string `json:"image_name,omitempty"`
	Name *string `json:"name,omitempty"`
	NamespaceName *string `json:"namespace_name,omitempty"`
	ProjectId string `json:"project_id"`
	StartTime *string `json:"start_time,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
}

// NewRequestsPutBackupJobRequest instantiates a new RequestsPutBackupJobRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestsPutBackupJobRequest(projectId string) *RequestsPutBackupJobRequest {
	this := RequestsPutBackupJobRequest{}
	this.ProjectId = projectId
	return &this
}

// NewRequestsPutBackupJobRequestWithDefaults instantiates a new RequestsPutBackupJobRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestsPutBackupJobRequestWithDefaults() *RequestsPutBackupJobRequest {
	this := RequestsPutBackupJobRequest{}
	return &this
}

// GetBackupCapability returns the BackupCapability field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetBackupCapability() string {
	if o == nil || o.BackupCapability == nil {
		var ret string
		return ret
	}
	return *o.BackupCapability
}

// GetBackupCapabilityOk returns a tuple with the BackupCapability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetBackupCapabilityOk() (*string, bool) {
	if o == nil || o.BackupCapability == nil {
		return nil, false
	}
	return o.BackupCapability, true
}

// HasBackupCapability returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasBackupCapability() bool {
	if o != nil && o.BackupCapability != nil {
		return true
	}

	return false
}

// SetBackupCapability gets a reference to the given string and assigns it to the BackupCapability field.
func (o *RequestsPutBackupJobRequest) SetBackupCapability(v string) {
	o.BackupCapability = &v
}

// GetBackupId returns the BackupId field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetBackupId() string {
	if o == nil || o.BackupId == nil {
		var ret string
		return ret
	}
	return *o.BackupId
}

// GetBackupIdOk returns a tuple with the BackupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetBackupIdOk() (*string, bool) {
	if o == nil || o.BackupId == nil {
		return nil, false
	}
	return o.BackupId, true
}

// HasBackupId returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasBackupId() bool {
	if o != nil && o.BackupId != nil {
		return true
	}

	return false
}

// SetBackupId gets a reference to the given string and assigns it to the BackupId field.
func (o *RequestsPutBackupJobRequest) SetBackupId(v string) {
	o.BackupId = &v
}

// GetBackupSpec returns the BackupSpec field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetBackupSpec() map[string]interface{} {
	if o == nil || o.BackupSpec == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.BackupSpec
}

// GetBackupSpecOk returns a tuple with the BackupSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetBackupSpecOk() (map[string]interface{}, bool) {
	if o == nil || o.BackupSpec == nil {
		return nil, false
	}
	return o.BackupSpec, true
}

// HasBackupSpec returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasBackupSpec() bool {
	if o != nil && o.BackupSpec != nil {
		return true
	}

	return false
}

// SetBackupSpec gets a reference to the given map[string]interface{} and assigns it to the BackupSpec field.
func (o *RequestsPutBackupJobRequest) SetBackupSpec(v map[string]interface{}) {
	o.BackupSpec = v
}

// GetCloudCredentialName returns the CloudCredentialName field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetCloudCredentialName() string {
	if o == nil || o.CloudCredentialName == nil {
		var ret string
		return ret
	}
	return *o.CloudCredentialName
}

// GetCloudCredentialNameOk returns a tuple with the CloudCredentialName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetCloudCredentialNameOk() (*string, bool) {
	if o == nil || o.CloudCredentialName == nil {
		return nil, false
	}
	return o.CloudCredentialName, true
}

// HasCloudCredentialName returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasCloudCredentialName() bool {
	if o != nil && o.CloudCredentialName != nil {
		return true
	}

	return false
}

// SetCloudCredentialName gets a reference to the given string and assigns it to the CloudCredentialName field.
func (o *RequestsPutBackupJobRequest) SetCloudCredentialName(v string) {
	o.CloudCredentialName = &v
}

// GetCloudSnapId returns the CloudSnapId field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetCloudSnapId() string {
	if o == nil || o.CloudSnapId == nil {
		var ret string
		return ret
	}
	return *o.CloudSnapId
}

// GetCloudSnapIdOk returns a tuple with the CloudSnapId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetCloudSnapIdOk() (*string, bool) {
	if o == nil || o.CloudSnapId == nil {
		return nil, false
	}
	return o.CloudSnapId, true
}

// HasCloudSnapId returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasCloudSnapId() bool {
	if o != nil && o.CloudSnapId != nil {
		return true
	}

	return false
}

// SetCloudSnapId gets a reference to the given string and assigns it to the CloudSnapId field.
func (o *RequestsPutBackupJobRequest) SetCloudSnapId(v string) {
	o.CloudSnapId = &v
}

// GetCompletionStatus returns the CompletionStatus field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetCompletionStatus() string {
	if o == nil || o.CompletionStatus == nil {
		var ret string
		return ret
	}
	return *o.CompletionStatus
}

// GetCompletionStatusOk returns a tuple with the CompletionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetCompletionStatusOk() (*string, bool) {
	if o == nil || o.CompletionStatus == nil {
		return nil, false
	}
	return o.CompletionStatus, true
}

// HasCompletionStatus returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasCompletionStatus() bool {
	if o != nil && o.CompletionStatus != nil {
		return true
	}

	return false
}

// SetCompletionStatus gets a reference to the given string and assigns it to the CompletionStatus field.
func (o *RequestsPutBackupJobRequest) SetCompletionStatus(v string) {
	o.CompletionStatus = &v
}

// GetCompletionTime returns the CompletionTime field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetCompletionTime() string {
	if o == nil || o.CompletionTime == nil {
		var ret string
		return ret
	}
	return *o.CompletionTime
}

// GetCompletionTimeOk returns a tuple with the CompletionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetCompletionTimeOk() (*string, bool) {
	if o == nil || o.CompletionTime == nil {
		return nil, false
	}
	return o.CompletionTime, true
}

// HasCompletionTime returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasCompletionTime() bool {
	if o != nil && o.CompletionTime != nil {
		return true
	}

	return false
}

// SetCompletionTime gets a reference to the given string and assigns it to the CompletionTime field.
func (o *RequestsPutBackupJobRequest) SetCompletionTime(v string) {
	o.CompletionTime = &v
}

// GetDataServiceSpec returns the DataServiceSpec field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetDataServiceSpec() map[string]interface{} {
	if o == nil || o.DataServiceSpec == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.DataServiceSpec
}

// GetDataServiceSpecOk returns a tuple with the DataServiceSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetDataServiceSpecOk() (map[string]interface{}, bool) {
	if o == nil || o.DataServiceSpec == nil {
		return nil, false
	}
	return o.DataServiceSpec, true
}

// HasDataServiceSpec returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasDataServiceSpec() bool {
	if o != nil && o.DataServiceSpec != nil {
		return true
	}

	return false
}

// SetDataServiceSpec gets a reference to the given map[string]interface{} and assigns it to the DataServiceSpec field.
func (o *RequestsPutBackupJobRequest) SetDataServiceSpec(v map[string]interface{}) {
	o.DataServiceSpec = v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *RequestsPutBackupJobRequest) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetDeploymentId returns the DeploymentId field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetDeploymentId() string {
	if o == nil || o.DeploymentId == nil {
		var ret string
		return ret
	}
	return *o.DeploymentId
}

// GetDeploymentIdOk returns a tuple with the DeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetDeploymentIdOk() (*string, bool) {
	if o == nil || o.DeploymentId == nil {
		return nil, false
	}
	return o.DeploymentId, true
}

// HasDeploymentId returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasDeploymentId() bool {
	if o != nil && o.DeploymentId != nil {
		return true
	}

	return false
}

// SetDeploymentId gets a reference to the given string and assigns it to the DeploymentId field.
func (o *RequestsPutBackupJobRequest) SetDeploymentId(v string) {
	o.DeploymentId = &v
}

// GetDeploymentTargetId returns the DeploymentTargetId field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetDeploymentTargetId() string {
	if o == nil || o.DeploymentTargetId == nil {
		var ret string
		return ret
	}
	return *o.DeploymentTargetId
}

// GetDeploymentTargetIdOk returns a tuple with the DeploymentTargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetDeploymentTargetIdOk() (*string, bool) {
	if o == nil || o.DeploymentTargetId == nil {
		return nil, false
	}
	return o.DeploymentTargetId, true
}

// HasDeploymentTargetId returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasDeploymentTargetId() bool {
	if o != nil && o.DeploymentTargetId != nil {
		return true
	}

	return false
}

// SetDeploymentTargetId gets a reference to the given string and assigns it to the DeploymentTargetId field.
func (o *RequestsPutBackupJobRequest) SetDeploymentTargetId(v string) {
	o.DeploymentTargetId = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *RequestsPutBackupJobRequest) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *RequestsPutBackupJobRequest) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetFileSize() int32 {
	if o == nil || o.FileSize == nil {
		var ret int32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetFileSizeOk() (*int32, bool) {
	if o == nil || o.FileSize == nil {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasFileSize() bool {
	if o != nil && o.FileSize != nil {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int32 and assigns it to the FileSize field.
func (o *RequestsPutBackupJobRequest) SetFileSize(v int32) {
	o.FileSize = &v
}

// GetImageName returns the ImageName field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetImageName() string {
	if o == nil || o.ImageName == nil {
		var ret string
		return ret
	}
	return *o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetImageNameOk() (*string, bool) {
	if o == nil || o.ImageName == nil {
		return nil, false
	}
	return o.ImageName, true
}

// HasImageName returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasImageName() bool {
	if o != nil && o.ImageName != nil {
		return true
	}

	return false
}

// SetImageName gets a reference to the given string and assigns it to the ImageName field.
func (o *RequestsPutBackupJobRequest) SetImageName(v string) {
	o.ImageName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RequestsPutBackupJobRequest) SetName(v string) {
	o.Name = &v
}

// GetNamespaceName returns the NamespaceName field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetNamespaceName() string {
	if o == nil || o.NamespaceName == nil {
		var ret string
		return ret
	}
	return *o.NamespaceName
}

// GetNamespaceNameOk returns a tuple with the NamespaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetNamespaceNameOk() (*string, bool) {
	if o == nil || o.NamespaceName == nil {
		return nil, false
	}
	return o.NamespaceName, true
}

// HasNamespaceName returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasNamespaceName() bool {
	if o != nil && o.NamespaceName != nil {
		return true
	}

	return false
}

// SetNamespaceName gets a reference to the given string and assigns it to the NamespaceName field.
func (o *RequestsPutBackupJobRequest) SetNamespaceName(v string) {
	o.NamespaceName = &v
}

// GetProjectId returns the ProjectId field value
func (o *RequestsPutBackupJobRequest) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetProjectIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *RequestsPutBackupJobRequest) SetProjectId(v string) {
	o.ProjectId = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetStartTime() string {
	if o == nil || o.StartTime == nil {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetStartTimeOk() (*string, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *RequestsPutBackupJobRequest) SetStartTime(v string) {
	o.StartTime = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *RequestsPutBackupJobRequest) GetTimestamp() string {
	if o == nil || o.Timestamp == nil {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestsPutBackupJobRequest) GetTimestampOk() (*string, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *RequestsPutBackupJobRequest) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *RequestsPutBackupJobRequest) SetTimestamp(v string) {
	o.Timestamp = &v
}

func (o RequestsPutBackupJobRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BackupCapability != nil {
		toSerialize["backup_capability"] = o.BackupCapability
	}
	if o.BackupId != nil {
		toSerialize["backup_id"] = o.BackupId
	}
	if o.BackupSpec != nil {
		toSerialize["backup_spec"] = o.BackupSpec
	}
	if o.CloudCredentialName != nil {
		toSerialize["cloud_credential_name"] = o.CloudCredentialName
	}
	if o.CloudSnapId != nil {
		toSerialize["cloud_snap_id"] = o.CloudSnapId
	}
	if o.CompletionStatus != nil {
		toSerialize["completion_status"] = o.CompletionStatus
	}
	if o.CompletionTime != nil {
		toSerialize["completion_time"] = o.CompletionTime
	}
	if o.DataServiceSpec != nil {
		toSerialize["data_service_spec"] = o.DataServiceSpec
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
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
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if o.FileSize != nil {
		toSerialize["file_size"] = o.FileSize
	}
	if o.ImageName != nil {
		toSerialize["image_name"] = o.ImageName
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NamespaceName != nil {
		toSerialize["namespace_name"] = o.NamespaceName
	}
	if true {
		toSerialize["project_id"] = o.ProjectId
	}
	if o.StartTime != nil {
		toSerialize["start_time"] = o.StartTime
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableRequestsPutBackupJobRequest struct {
	value *RequestsPutBackupJobRequest
	isSet bool
}

func (v NullableRequestsPutBackupJobRequest) Get() *RequestsPutBackupJobRequest {
	return v.value
}

func (v *NullableRequestsPutBackupJobRequest) Set(val *RequestsPutBackupJobRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestsPutBackupJobRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestsPutBackupJobRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestsPutBackupJobRequest(val *RequestsPutBackupJobRequest) *NullableRequestsPutBackupJobRequest {
	return &NullableRequestsPutBackupJobRequest{value: val, isSet: true}
}

func (v NullableRequestsPutBackupJobRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestsPutBackupJobRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


