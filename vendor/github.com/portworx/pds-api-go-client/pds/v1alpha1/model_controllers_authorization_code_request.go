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

// ControllersAuthorizationCodeRequest struct for ControllersAuthorizationCodeRequest
type ControllersAuthorizationCodeRequest struct {
	AuthCode *string `json:"AuthCode,omitempty"`
}

// NewControllersAuthorizationCodeRequest instantiates a new ControllersAuthorizationCodeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersAuthorizationCodeRequest() *ControllersAuthorizationCodeRequest {
	this := ControllersAuthorizationCodeRequest{}
	return &this
}

// NewControllersAuthorizationCodeRequestWithDefaults instantiates a new ControllersAuthorizationCodeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersAuthorizationCodeRequestWithDefaults() *ControllersAuthorizationCodeRequest {
	this := ControllersAuthorizationCodeRequest{}
	return &this
}

// GetAuthCode returns the AuthCode field value if set, zero value otherwise.
func (o *ControllersAuthorizationCodeRequest) GetAuthCode() string {
	if o == nil || o.AuthCode == nil {
		var ret string
		return ret
	}
	return *o.AuthCode
}

// GetAuthCodeOk returns a tuple with the AuthCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersAuthorizationCodeRequest) GetAuthCodeOk() (*string, bool) {
	if o == nil || o.AuthCode == nil {
		return nil, false
	}
	return o.AuthCode, true
}

// HasAuthCode returns a boolean if a field has been set.
func (o *ControllersAuthorizationCodeRequest) HasAuthCode() bool {
	if o != nil && o.AuthCode != nil {
		return true
	}

	return false
}

// SetAuthCode gets a reference to the given string and assigns it to the AuthCode field.
func (o *ControllersAuthorizationCodeRequest) SetAuthCode(v string) {
	o.AuthCode = &v
}

func (o ControllersAuthorizationCodeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthCode != nil {
		toSerialize["AuthCode"] = o.AuthCode
	}
	return json.Marshal(toSerialize)
}

type NullableControllersAuthorizationCodeRequest struct {
	value *ControllersAuthorizationCodeRequest
	isSet bool
}

func (v NullableControllersAuthorizationCodeRequest) Get() *ControllersAuthorizationCodeRequest {
	return v.value
}

func (v *NullableControllersAuthorizationCodeRequest) Set(val *ControllersAuthorizationCodeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersAuthorizationCodeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersAuthorizationCodeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersAuthorizationCodeRequest(val *ControllersAuthorizationCodeRequest) *NullableControllersAuthorizationCodeRequest {
	return &NullableControllersAuthorizationCodeRequest{value: val, isSet: true}
}

func (v NullableControllersAuthorizationCodeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersAuthorizationCodeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


