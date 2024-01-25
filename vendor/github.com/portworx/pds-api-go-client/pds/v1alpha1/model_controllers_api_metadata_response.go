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

// ControllersAPIMetadataResponse struct for ControllersAPIMetadataResponse
type ControllersAPIMetadataResponse struct {
	Features *map[string]string `json:"features,omitempty"`
	HelmChartVersion *string `json:"helm_chart_version,omitempty"`
	PdsBuildNumber *int32 `json:"pds_build_number,omitempty"`
}

// NewControllersAPIMetadataResponse instantiates a new ControllersAPIMetadataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersAPIMetadataResponse() *ControllersAPIMetadataResponse {
	this := ControllersAPIMetadataResponse{}
	return &this
}

// NewControllersAPIMetadataResponseWithDefaults instantiates a new ControllersAPIMetadataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersAPIMetadataResponseWithDefaults() *ControllersAPIMetadataResponse {
	this := ControllersAPIMetadataResponse{}
	return &this
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *ControllersAPIMetadataResponse) GetFeatures() map[string]string {
	if o == nil || o.Features == nil {
		var ret map[string]string
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersAPIMetadataResponse) GetFeaturesOk() (*map[string]string, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *ControllersAPIMetadataResponse) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given map[string]string and assigns it to the Features field.
func (o *ControllersAPIMetadataResponse) SetFeatures(v map[string]string) {
	o.Features = &v
}

// GetHelmChartVersion returns the HelmChartVersion field value if set, zero value otherwise.
func (o *ControllersAPIMetadataResponse) GetHelmChartVersion() string {
	if o == nil || o.HelmChartVersion == nil {
		var ret string
		return ret
	}
	return *o.HelmChartVersion
}

// GetHelmChartVersionOk returns a tuple with the HelmChartVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersAPIMetadataResponse) GetHelmChartVersionOk() (*string, bool) {
	if o == nil || o.HelmChartVersion == nil {
		return nil, false
	}
	return o.HelmChartVersion, true
}

// HasHelmChartVersion returns a boolean if a field has been set.
func (o *ControllersAPIMetadataResponse) HasHelmChartVersion() bool {
	if o != nil && o.HelmChartVersion != nil {
		return true
	}

	return false
}

// SetHelmChartVersion gets a reference to the given string and assigns it to the HelmChartVersion field.
func (o *ControllersAPIMetadataResponse) SetHelmChartVersion(v string) {
	o.HelmChartVersion = &v
}

// GetPdsBuildNumber returns the PdsBuildNumber field value if set, zero value otherwise.
func (o *ControllersAPIMetadataResponse) GetPdsBuildNumber() int32 {
	if o == nil || o.PdsBuildNumber == nil {
		var ret int32
		return ret
	}
	return *o.PdsBuildNumber
}

// GetPdsBuildNumberOk returns a tuple with the PdsBuildNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersAPIMetadataResponse) GetPdsBuildNumberOk() (*int32, bool) {
	if o == nil || o.PdsBuildNumber == nil {
		return nil, false
	}
	return o.PdsBuildNumber, true
}

// HasPdsBuildNumber returns a boolean if a field has been set.
func (o *ControllersAPIMetadataResponse) HasPdsBuildNumber() bool {
	if o != nil && o.PdsBuildNumber != nil {
		return true
	}

	return false
}

// SetPdsBuildNumber gets a reference to the given int32 and assigns it to the PdsBuildNumber field.
func (o *ControllersAPIMetadataResponse) SetPdsBuildNumber(v int32) {
	o.PdsBuildNumber = &v
}

func (o ControllersAPIMetadataResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	if o.HelmChartVersion != nil {
		toSerialize["helm_chart_version"] = o.HelmChartVersion
	}
	if o.PdsBuildNumber != nil {
		toSerialize["pds_build_number"] = o.PdsBuildNumber
	}
	return json.Marshal(toSerialize)
}

type NullableControllersAPIMetadataResponse struct {
	value *ControllersAPIMetadataResponse
	isSet bool
}

func (v NullableControllersAPIMetadataResponse) Get() *ControllersAPIMetadataResponse {
	return v.value
}

func (v *NullableControllersAPIMetadataResponse) Set(val *ControllersAPIMetadataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersAPIMetadataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersAPIMetadataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersAPIMetadataResponse(val *ControllersAPIMetadataResponse) *NullableControllersAPIMetadataResponse {
	return &NullableControllersAPIMetadataResponse{value: val, isSet: true}
}

func (v NullableControllersAPIMetadataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersAPIMetadataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


