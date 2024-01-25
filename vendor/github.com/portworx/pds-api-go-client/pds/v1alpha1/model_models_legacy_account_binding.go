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

// ModelsLegacyAccountBinding struct for ModelsLegacyAccountBinding
type ModelsLegacyAccountBinding struct {
	AccountId *string `json:"account_id,omitempty"`
	ActorId *string `json:"actor_id,omitempty"`
	ActorType *string `json:"actor_type,omitempty"`
	RoleName *string `json:"role_name,omitempty"`
}

// NewModelsLegacyAccountBinding instantiates a new ModelsLegacyAccountBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsLegacyAccountBinding() *ModelsLegacyAccountBinding {
	this := ModelsLegacyAccountBinding{}
	return &this
}

// NewModelsLegacyAccountBindingWithDefaults instantiates a new ModelsLegacyAccountBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsLegacyAccountBindingWithDefaults() *ModelsLegacyAccountBinding {
	this := ModelsLegacyAccountBinding{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ModelsLegacyAccountBinding) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsLegacyAccountBinding) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *ModelsLegacyAccountBinding) HasAccountId() bool {
	if o != nil && o.AccountId != nil {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *ModelsLegacyAccountBinding) SetAccountId(v string) {
	o.AccountId = &v
}

// GetActorId returns the ActorId field value if set, zero value otherwise.
func (o *ModelsLegacyAccountBinding) GetActorId() string {
	if o == nil || o.ActorId == nil {
		var ret string
		return ret
	}
	return *o.ActorId
}

// GetActorIdOk returns a tuple with the ActorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsLegacyAccountBinding) GetActorIdOk() (*string, bool) {
	if o == nil || o.ActorId == nil {
		return nil, false
	}
	return o.ActorId, true
}

// HasActorId returns a boolean if a field has been set.
func (o *ModelsLegacyAccountBinding) HasActorId() bool {
	if o != nil && o.ActorId != nil {
		return true
	}

	return false
}

// SetActorId gets a reference to the given string and assigns it to the ActorId field.
func (o *ModelsLegacyAccountBinding) SetActorId(v string) {
	o.ActorId = &v
}

// GetActorType returns the ActorType field value if set, zero value otherwise.
func (o *ModelsLegacyAccountBinding) GetActorType() string {
	if o == nil || o.ActorType == nil {
		var ret string
		return ret
	}
	return *o.ActorType
}

// GetActorTypeOk returns a tuple with the ActorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsLegacyAccountBinding) GetActorTypeOk() (*string, bool) {
	if o == nil || o.ActorType == nil {
		return nil, false
	}
	return o.ActorType, true
}

// HasActorType returns a boolean if a field has been set.
func (o *ModelsLegacyAccountBinding) HasActorType() bool {
	if o != nil && o.ActorType != nil {
		return true
	}

	return false
}

// SetActorType gets a reference to the given string and assigns it to the ActorType field.
func (o *ModelsLegacyAccountBinding) SetActorType(v string) {
	o.ActorType = &v
}

// GetRoleName returns the RoleName field value if set, zero value otherwise.
func (o *ModelsLegacyAccountBinding) GetRoleName() string {
	if o == nil || o.RoleName == nil {
		var ret string
		return ret
	}
	return *o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsLegacyAccountBinding) GetRoleNameOk() (*string, bool) {
	if o == nil || o.RoleName == nil {
		return nil, false
	}
	return o.RoleName, true
}

// HasRoleName returns a boolean if a field has been set.
func (o *ModelsLegacyAccountBinding) HasRoleName() bool {
	if o != nil && o.RoleName != nil {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given string and assigns it to the RoleName field.
func (o *ModelsLegacyAccountBinding) SetRoleName(v string) {
	o.RoleName = &v
}

func (o ModelsLegacyAccountBinding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.ActorId != nil {
		toSerialize["actor_id"] = o.ActorId
	}
	if o.ActorType != nil {
		toSerialize["actor_type"] = o.ActorType
	}
	if o.RoleName != nil {
		toSerialize["role_name"] = o.RoleName
	}
	return json.Marshal(toSerialize)
}

type NullableModelsLegacyAccountBinding struct {
	value *ModelsLegacyAccountBinding
	isSet bool
}

func (v NullableModelsLegacyAccountBinding) Get() *ModelsLegacyAccountBinding {
	return v.value
}

func (v *NullableModelsLegacyAccountBinding) Set(val *ModelsLegacyAccountBinding) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsLegacyAccountBinding) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsLegacyAccountBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsLegacyAccountBinding(val *ModelsLegacyAccountBinding) *NullableModelsLegacyAccountBinding {
	return &NullableModelsLegacyAccountBinding{value: val, isSet: true}
}

func (v NullableModelsLegacyAccountBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsLegacyAccountBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


