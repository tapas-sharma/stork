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

// ModelsAWSDetails struct for ModelsAWSDetails
type ModelsAWSDetails struct {
	AccessKey *string `json:"access_key,omitempty"`
	HostedZoneId *string `json:"hosted_zone_id,omitempty"`
	SecretKey *string `json:"secret_key,omitempty"`
}

// NewModelsAWSDetails instantiates a new ModelsAWSDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsAWSDetails() *ModelsAWSDetails {
	this := ModelsAWSDetails{}
	return &this
}

// NewModelsAWSDetailsWithDefaults instantiates a new ModelsAWSDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsAWSDetailsWithDefaults() *ModelsAWSDetails {
	this := ModelsAWSDetails{}
	return &this
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *ModelsAWSDetails) GetAccessKey() string {
	if o == nil || o.AccessKey == nil {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsAWSDetails) GetAccessKeyOk() (*string, bool) {
	if o == nil || o.AccessKey == nil {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *ModelsAWSDetails) HasAccessKey() bool {
	if o != nil && o.AccessKey != nil {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *ModelsAWSDetails) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetHostedZoneId returns the HostedZoneId field value if set, zero value otherwise.
func (o *ModelsAWSDetails) GetHostedZoneId() string {
	if o == nil || o.HostedZoneId == nil {
		var ret string
		return ret
	}
	return *o.HostedZoneId
}

// GetHostedZoneIdOk returns a tuple with the HostedZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsAWSDetails) GetHostedZoneIdOk() (*string, bool) {
	if o == nil || o.HostedZoneId == nil {
		return nil, false
	}
	return o.HostedZoneId, true
}

// HasHostedZoneId returns a boolean if a field has been set.
func (o *ModelsAWSDetails) HasHostedZoneId() bool {
	if o != nil && o.HostedZoneId != nil {
		return true
	}

	return false
}

// SetHostedZoneId gets a reference to the given string and assigns it to the HostedZoneId field.
func (o *ModelsAWSDetails) SetHostedZoneId(v string) {
	o.HostedZoneId = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *ModelsAWSDetails) GetSecretKey() string {
	if o == nil || o.SecretKey == nil {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsAWSDetails) GetSecretKeyOk() (*string, bool) {
	if o == nil || o.SecretKey == nil {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *ModelsAWSDetails) HasSecretKey() bool {
	if o != nil && o.SecretKey != nil {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *ModelsAWSDetails) SetSecretKey(v string) {
	o.SecretKey = &v
}

func (o ModelsAWSDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessKey != nil {
		toSerialize["access_key"] = o.AccessKey
	}
	if o.HostedZoneId != nil {
		toSerialize["hosted_zone_id"] = o.HostedZoneId
	}
	if o.SecretKey != nil {
		toSerialize["secret_key"] = o.SecretKey
	}
	return json.Marshal(toSerialize)
}

type NullableModelsAWSDetails struct {
	value *ModelsAWSDetails
	isSet bool
}

func (v NullableModelsAWSDetails) Get() *ModelsAWSDetails {
	return v.value
}

func (v *NullableModelsAWSDetails) Set(val *ModelsAWSDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsAWSDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsAWSDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsAWSDetails(val *ModelsAWSDetails) *NullableModelsAWSDetails {
	return &NullableModelsAWSDetails{value: val, isSet: true}
}

func (v NullableModelsAWSDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsAWSDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


