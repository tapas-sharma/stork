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

// ModelsServiceIdentity struct for ModelsServiceIdentity
type ModelsServiceIdentity struct {
	AccountId *string `json:"account_id,omitempty"`
	ClientId *string `json:"client_id,omitempty"`
	// CreatedAt is autogenerated on creation
	CreatedAt *string `json:"created_at,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	Description *string `json:"description,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	// ID is auto generated on creation
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	SecretGenerationCount *int32 `json:"secret_generation_count,omitempty"`
	SecretUpdatedAt *string `json:"secret_updated_at,omitempty"`
	// UpdatedAt is autogenerated on update
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewModelsServiceIdentity instantiates a new ModelsServiceIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsServiceIdentity() *ModelsServiceIdentity {
	this := ModelsServiceIdentity{}
	return &this
}

// NewModelsServiceIdentityWithDefaults instantiates a new ModelsServiceIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsServiceIdentityWithDefaults() *ModelsServiceIdentity {
	this := ModelsServiceIdentity{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ModelsServiceIdentity) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsServiceIdentity) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *ModelsServiceIdentity) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *ModelsServiceIdentity) SetAccountId(v string) {
	o.AccountId = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ModelsServiceIdentity) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsServiceIdentity) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ModelsServiceIdentity) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ModelsServiceIdentity) SetClientId(v string) {
	o.ClientId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ModelsServiceIdentity) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsServiceIdentity) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ModelsServiceIdentity) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ModelsServiceIdentity) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ModelsServiceIdentity) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsServiceIdentity) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ModelsServiceIdentity) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *ModelsServiceIdentity) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModelsServiceIdentity) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsServiceIdentity) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModelsServiceIdentity) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModelsServiceIdentity) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ModelsServiceIdentity) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsServiceIdentity) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ModelsServiceIdentity) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ModelsServiceIdentity) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsServiceIdentity) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsServiceIdentity) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsServiceIdentity) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelsServiceIdentity) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelsServiceIdentity) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsServiceIdentity) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelsServiceIdentity) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelsServiceIdentity) SetName(v string) {
	o.Name = &v
}

// GetSecretGenerationCount returns the SecretGenerationCount field value if set, zero value otherwise.
func (o *ModelsServiceIdentity) GetSecretGenerationCount() int32 {
	if o == nil || o.SecretGenerationCount == nil {
		var ret int32
		return ret
	}
	return *o.SecretGenerationCount
}

// GetSecretGenerationCountOk returns a tuple with the SecretGenerationCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsServiceIdentity) GetSecretGenerationCountOk() (*int32, bool) {
	if o == nil || o.SecretGenerationCount == nil {
		return nil, false
	}
	return o.SecretGenerationCount, true
}

// HasSecretGenerationCount returns a boolean if a field has been set.
func (o *ModelsServiceIdentity) HasSecretGenerationCount() bool {
	if o != nil && o.SecretGenerationCount != nil {
		return true
	}

	return false
}

// SetSecretGenerationCount gets a reference to the given int32 and assigns it to the SecretGenerationCount field.
func (o *ModelsServiceIdentity) SetSecretGenerationCount(v int32) {
	o.SecretGenerationCount = &v
}

// GetSecretUpdatedAt returns the SecretUpdatedAt field value if set, zero value otherwise.
func (o *ModelsServiceIdentity) GetSecretUpdatedAt() string {
	if o == nil || o.SecretUpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.SecretUpdatedAt
}

// GetSecretUpdatedAtOk returns a tuple with the SecretUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsServiceIdentity) GetSecretUpdatedAtOk() (*string, bool) {
	if o == nil || o.SecretUpdatedAt == nil {
		return nil, false
	}
	return o.SecretUpdatedAt, true
}

// HasSecretUpdatedAt returns a boolean if a field has been set.
func (o *ModelsServiceIdentity) HasSecretUpdatedAt() bool {
	if o != nil && o.SecretUpdatedAt != nil {
		return true
	}

	return false
}

// SetSecretUpdatedAt gets a reference to the given string and assigns it to the SecretUpdatedAt field.
func (o *ModelsServiceIdentity) SetSecretUpdatedAt(v string) {
	o.SecretUpdatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ModelsServiceIdentity) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsServiceIdentity) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ModelsServiceIdentity) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ModelsServiceIdentity) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o ModelsServiceIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SecretGenerationCount != nil {
		toSerialize["secret_generation_count"] = o.SecretGenerationCount
	}
	if o.SecretUpdatedAt != nil {
		toSerialize["secret_updated_at"] = o.SecretUpdatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableModelsServiceIdentity struct {
	value *ModelsServiceIdentity
	isSet bool
}

func (v NullableModelsServiceIdentity) Get() *ModelsServiceIdentity {
	return v.value
}

func (v *NullableModelsServiceIdentity) Set(val *ModelsServiceIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsServiceIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsServiceIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsServiceIdentity(val *ModelsServiceIdentity) *NullableModelsServiceIdentity {
	return &NullableModelsServiceIdentity{value: val, isSet: true}
}

func (v NullableModelsServiceIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsServiceIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


