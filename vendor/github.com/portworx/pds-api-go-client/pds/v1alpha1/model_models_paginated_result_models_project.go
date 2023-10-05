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

// ModelsPaginatedResultModelsProject struct for ModelsPaginatedResultModelsProject
type ModelsPaginatedResultModelsProject struct {
	Data []ModelsProject `json:"data,omitempty"`
	Pagination *ConstraintPagination `json:"pagination,omitempty"`
}

// NewModelsPaginatedResultModelsProject instantiates a new ModelsPaginatedResultModelsProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsPaginatedResultModelsProject() *ModelsPaginatedResultModelsProject {
	this := ModelsPaginatedResultModelsProject{}
	return &this
}

// NewModelsPaginatedResultModelsProjectWithDefaults instantiates a new ModelsPaginatedResultModelsProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsPaginatedResultModelsProjectWithDefaults() *ModelsPaginatedResultModelsProject {
	this := ModelsPaginatedResultModelsProject{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ModelsPaginatedResultModelsProject) GetData() []ModelsProject {
	if o == nil || o.Data == nil {
		var ret []ModelsProject
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsPaginatedResultModelsProject) GetDataOk() ([]ModelsProject, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ModelsPaginatedResultModelsProject) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []ModelsProject and assigns it to the Data field.
func (o *ModelsPaginatedResultModelsProject) SetData(v []ModelsProject) {
	o.Data = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ModelsPaginatedResultModelsProject) GetPagination() ConstraintPagination {
	if o == nil || o.Pagination == nil {
		var ret ConstraintPagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsPaginatedResultModelsProject) GetPaginationOk() (*ConstraintPagination, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ModelsPaginatedResultModelsProject) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given ConstraintPagination and assigns it to the Pagination field.
func (o *ModelsPaginatedResultModelsProject) SetPagination(v ConstraintPagination) {
	o.Pagination = &v
}

func (o ModelsPaginatedResultModelsProject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableModelsPaginatedResultModelsProject struct {
	value *ModelsPaginatedResultModelsProject
	isSet bool
}

func (v NullableModelsPaginatedResultModelsProject) Get() *ModelsPaginatedResultModelsProject {
	return v.value
}

func (v *NullableModelsPaginatedResultModelsProject) Set(val *ModelsPaginatedResultModelsProject) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsPaginatedResultModelsProject) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsPaginatedResultModelsProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsPaginatedResultModelsProject(val *ModelsPaginatedResultModelsProject) *NullableModelsPaginatedResultModelsProject {
	return &NullableModelsPaginatedResultModelsProject{value: val, isSet: true}
}

func (v NullableModelsPaginatedResultModelsProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsPaginatedResultModelsProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


