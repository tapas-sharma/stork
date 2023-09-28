/*
PDS API

Portworx Data Services API Server

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pds

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Linger please
var (
	_ context.Context
)

// GlobalRoleBindingsApiService GlobalRoleBindingsApi service
type GlobalRoleBindingsApiService service

type ApiApiGlobalRoleBindingsDeleteRequest struct {
	ctx context.Context
	ApiService *GlobalRoleBindingsApiService
	body *RequestsDeleteRoleBindingRequest
}

// Request body containing the global role binding
func (r ApiApiGlobalRoleBindingsDeleteRequest) Body(body RequestsDeleteRoleBindingRequest) ApiApiGlobalRoleBindingsDeleteRequest {
	r.body = &body
	return r
}

func (r ApiApiGlobalRoleBindingsDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiGlobalRoleBindingsDeleteExecute(r)
}

/*
ApiGlobalRoleBindingsDelete Delete GlobalRoleBinding

Removes a single GlobalRoleBindings

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiGlobalRoleBindingsDeleteRequest
*/
func (a *GlobalRoleBindingsApiService) ApiGlobalRoleBindingsDelete(ctx context.Context) ApiApiGlobalRoleBindingsDeleteRequest {
	return ApiApiGlobalRoleBindingsDeleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *GlobalRoleBindingsApiService) ApiGlobalRoleBindingsDeleteExecute(r ApiApiGlobalRoleBindingsDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlobalRoleBindingsApiService.ApiGlobalRoleBindingsDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/global-role-bindings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiApiGlobalRoleBindingsGetRequest struct {
	ctx context.Context
	ApiService *GlobalRoleBindingsApiService
	sortBy *string
	roleName *string
	actorId *string
	actorType *string
}

// A given GlobalRoleBinding attribute to sort results by (one of: role_name, actor_id)
func (r ApiApiGlobalRoleBindingsGetRequest) SortBy(sortBy string) ApiApiGlobalRoleBindingsGetRequest {
	r.sortBy = &sortBy
	return r
}
// Filter results by GlobalRoleBinding assigned role name
func (r ApiApiGlobalRoleBindingsGetRequest) RoleName(roleName string) ApiApiGlobalRoleBindingsGetRequest {
	r.roleName = &roleName
	return r
}
// Filter results by GlobalRoleBinding actor id
func (r ApiApiGlobalRoleBindingsGetRequest) ActorId(actorId string) ApiApiGlobalRoleBindingsGetRequest {
	r.actorId = &actorId
	return r
}
// Filter results by GlobalRoleBinding actor type
func (r ApiApiGlobalRoleBindingsGetRequest) ActorType(actorType string) ApiApiGlobalRoleBindingsGetRequest {
	r.actorType = &actorType
	return r
}

func (r ApiApiGlobalRoleBindingsGetRequest) Execute() (*ControllersPaginatedGlobalRoleBindings, *http.Response, error) {
	return r.ApiService.ApiGlobalRoleBindingsGetExecute(r)
}

/*
ApiGlobalRoleBindingsGet List GlobalRoleBindings

Lists GlobalRoleBindings

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiGlobalRoleBindingsGetRequest
*/
func (a *GlobalRoleBindingsApiService) ApiGlobalRoleBindingsGet(ctx context.Context) ApiApiGlobalRoleBindingsGetRequest {
	return ApiApiGlobalRoleBindingsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ControllersPaginatedGlobalRoleBindings
func (a *GlobalRoleBindingsApiService) ApiGlobalRoleBindingsGetExecute(r ApiApiGlobalRoleBindingsGetRequest) (*ControllersPaginatedGlobalRoleBindings, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ControllersPaginatedGlobalRoleBindings
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlobalRoleBindingsApiService.ApiGlobalRoleBindingsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/global-role-bindings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sortBy != nil {
		localVarQueryParams.Add("sort_by", parameterToString(*r.sortBy, ""))
	}
	if r.roleName != nil {
		localVarQueryParams.Add("role_name", parameterToString(*r.roleName, ""))
	}
	if r.actorId != nil {
		localVarQueryParams.Add("actor_id", parameterToString(*r.actorId, ""))
	}
	if r.actorType != nil {
		localVarQueryParams.Add("actor_type", parameterToString(*r.actorType, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiGlobalRoleBindingsPutRequest struct {
	ctx context.Context
	ApiService *GlobalRoleBindingsApiService
	body *RequestsPutLegacyBindingRequest
}

// Request body containing the global role binding
func (r ApiApiGlobalRoleBindingsPutRequest) Body(body RequestsPutLegacyBindingRequest) ApiApiGlobalRoleBindingsPutRequest {
	r.body = &body
	return r
}

func (r ApiApiGlobalRoleBindingsPutRequest) Execute() (*ModelsLegacyGlobalBinding, *http.Response, error) {
	return r.ApiService.ApiGlobalRoleBindingsPutExecute(r)
}

/*
ApiGlobalRoleBindingsPut Set GlobalRoleBinding

Creates a new GlobalRoleBinding

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiApiGlobalRoleBindingsPutRequest
*/
func (a *GlobalRoleBindingsApiService) ApiGlobalRoleBindingsPut(ctx context.Context) ApiApiGlobalRoleBindingsPutRequest {
	return ApiApiGlobalRoleBindingsPutRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ModelsLegacyGlobalBinding
func (a *GlobalRoleBindingsApiService) ApiGlobalRoleBindingsPutExecute(r ApiApiGlobalRoleBindingsPutRequest) (*ModelsLegacyGlobalBinding, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsLegacyGlobalBinding
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlobalRoleBindingsApiService.ApiGlobalRoleBindingsPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/global-role-bindings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
