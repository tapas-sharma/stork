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

// ControllersOIDCInfoResponse struct for ControllersOIDCInfoResponse
type ControllersOIDCInfoResponse struct {
	AuthUrl *string `json:"authUrl,omitempty"`
	ClientId *string `json:"clientId,omitempty"`
	Issuer *string `json:"issuer,omitempty"`
	RedirectUrl *string `json:"redirectUrl,omitempty"`
}

// NewControllersOIDCInfoResponse instantiates a new ControllersOIDCInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControllersOIDCInfoResponse() *ControllersOIDCInfoResponse {
	this := ControllersOIDCInfoResponse{}
	return &this
}

// NewControllersOIDCInfoResponseWithDefaults instantiates a new ControllersOIDCInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControllersOIDCInfoResponseWithDefaults() *ControllersOIDCInfoResponse {
	this := ControllersOIDCInfoResponse{}
	return &this
}

// GetAuthUrl returns the AuthUrl field value if set, zero value otherwise.
func (o *ControllersOIDCInfoResponse) GetAuthUrl() string {
	if o == nil || o.AuthUrl == nil {
		var ret string
		return ret
	}
	return *o.AuthUrl
}

// GetAuthUrlOk returns a tuple with the AuthUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersOIDCInfoResponse) GetAuthUrlOk() (*string, bool) {
	if o == nil || o.AuthUrl == nil {
		return nil, false
	}
	return o.AuthUrl, true
}

// HasAuthUrl returns a boolean if a field has been set.
func (o *ControllersOIDCInfoResponse) HasAuthUrl() bool {
	if o != nil && o.AuthUrl != nil {
		return true
	}

	return false
}

// SetAuthUrl gets a reference to the given string and assigns it to the AuthUrl field.
func (o *ControllersOIDCInfoResponse) SetAuthUrl(v string) {
	o.AuthUrl = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ControllersOIDCInfoResponse) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersOIDCInfoResponse) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ControllersOIDCInfoResponse) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ControllersOIDCInfoResponse) SetClientId(v string) {
	o.ClientId = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *ControllersOIDCInfoResponse) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersOIDCInfoResponse) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *ControllersOIDCInfoResponse) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *ControllersOIDCInfoResponse) SetIssuer(v string) {
	o.Issuer = &v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *ControllersOIDCInfoResponse) GetRedirectUrl() string {
	if o == nil || o.RedirectUrl == nil {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControllersOIDCInfoResponse) GetRedirectUrlOk() (*string, bool) {
	if o == nil || o.RedirectUrl == nil {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *ControllersOIDCInfoResponse) HasRedirectUrl() bool {
	if o != nil && o.RedirectUrl != nil {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *ControllersOIDCInfoResponse) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

func (o ControllersOIDCInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthUrl != nil {
		toSerialize["authUrl"] = o.AuthUrl
	}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.RedirectUrl != nil {
		toSerialize["redirectUrl"] = o.RedirectUrl
	}
	return json.Marshal(toSerialize)
}

type NullableControllersOIDCInfoResponse struct {
	value *ControllersOIDCInfoResponse
	isSet bool
}

func (v NullableControllersOIDCInfoResponse) Get() *ControllersOIDCInfoResponse {
	return v.value
}

func (v *NullableControllersOIDCInfoResponse) Set(val *ControllersOIDCInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableControllersOIDCInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableControllersOIDCInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControllersOIDCInfoResponse(val *ControllersOIDCInfoResponse) *NullableControllersOIDCInfoResponse {
	return &NullableControllersOIDCInfoResponse{value: val, isSet: true}
}

func (v NullableControllersOIDCInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControllersOIDCInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


