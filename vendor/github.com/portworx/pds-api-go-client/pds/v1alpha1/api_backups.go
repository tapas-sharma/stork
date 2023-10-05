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
	"strings"
)

// Linger please
var (
	_ context.Context
)

// BackupsApiService BackupsApi service
type BackupsApiService service

type ApiApiBackupTargetsIdBackupsGetRequest struct {
	ctx context.Context
	ApiService *BackupsApiService
	id string
	sortBy *string
	limit *string
	continuation *string
	id2 *string
	clusterResourceName *string
	state *string
	suspend *string
	backupType *string
	backupLevel *string
}

// A given Backup attribute to sort results by (one of: id, cluster_resource_name, created_at, backup_type, backup_level)
func (r ApiApiBackupTargetsIdBackupsGetRequest) SortBy(sortBy string) ApiApiBackupTargetsIdBackupsGetRequest {
	r.sortBy = &sortBy
	return r
}
// Maximum number of rows to return (could be less)
func (r ApiApiBackupTargetsIdBackupsGetRequest) Limit(limit string) ApiApiBackupTargetsIdBackupsGetRequest {
	r.limit = &limit
	return r
}
// Use a token returned by a previous query to continue listing with the next batch of rows
func (r ApiApiBackupTargetsIdBackupsGetRequest) Continuation(continuation string) ApiApiBackupTargetsIdBackupsGetRequest {
	r.continuation = &continuation
	return r
}
// Filter results by Backup id
func (r ApiApiBackupTargetsIdBackupsGetRequest) Id2(id2 string) ApiApiBackupTargetsIdBackupsGetRequest {
	r.id2 = &id2
	return r
}
// Filter results by Backup cluster_resource_name
func (r ApiApiBackupTargetsIdBackupsGetRequest) ClusterResourceName(clusterResourceName string) ApiApiBackupTargetsIdBackupsGetRequest {
	r.clusterResourceName = &clusterResourceName
	return r
}
// Filter results by Backup state
func (r ApiApiBackupTargetsIdBackupsGetRequest) State(state string) ApiApiBackupTargetsIdBackupsGetRequest {
	r.state = &state
	return r
}
// Filter results by Backup suspend flag
func (r ApiApiBackupTargetsIdBackupsGetRequest) Suspend(suspend string) ApiApiBackupTargetsIdBackupsGetRequest {
	r.suspend = &suspend
	return r
}
// Filter results by Backup type (one of: adhoc,scheduled)
func (r ApiApiBackupTargetsIdBackupsGetRequest) BackupType(backupType string) ApiApiBackupTargetsIdBackupsGetRequest {
	r.backupType = &backupType
	return r
}
// Filter results by Backup type (one of: snapshot,incremental)
func (r ApiApiBackupTargetsIdBackupsGetRequest) BackupLevel(backupLevel string) ApiApiBackupTargetsIdBackupsGetRequest {
	r.backupLevel = &backupLevel
	return r
}

func (r ApiApiBackupTargetsIdBackupsGetRequest) Execute() (*ModelsPaginatedResultModelsBackup, *http.Response, error) {
	return r.ApiService.ApiBackupTargetsIdBackupsGetExecute(r)
}

/*
ApiBackupTargetsIdBackupsGet List BackupTarget's Backups

Lists Backups belonging to the BackupTarget.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id BackupTarget ID (must be valid UUID)
 @return ApiApiBackupTargetsIdBackupsGetRequest
*/
func (a *BackupsApiService) ApiBackupTargetsIdBackupsGet(ctx context.Context, id string) ApiApiBackupTargetsIdBackupsGetRequest {
	return ApiApiBackupTargetsIdBackupsGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsPaginatedResultModelsBackup
func (a *BackupsApiService) ApiBackupTargetsIdBackupsGetExecute(r ApiApiBackupTargetsIdBackupsGetRequest) (*ModelsPaginatedResultModelsBackup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsPaginatedResultModelsBackup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupsApiService.ApiBackupTargetsIdBackupsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/backup-targets/{id}/backups"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sortBy != nil {
		localVarQueryParams.Add("sort_by", parameterToString(*r.sortBy, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.continuation != nil {
		localVarQueryParams.Add("continuation", parameterToString(*r.continuation, ""))
	}
	if r.id2 != nil {
		localVarQueryParams.Add("id", parameterToString(*r.id2, ""))
	}
	if r.clusterResourceName != nil {
		localVarQueryParams.Add("cluster_resource_name", parameterToString(*r.clusterResourceName, ""))
	}
	if r.state != nil {
		localVarQueryParams.Add("state", parameterToString(*r.state, ""))
	}
	if r.suspend != nil {
		localVarQueryParams.Add("suspend", parameterToString(*r.suspend, ""))
	}
	if r.backupType != nil {
		localVarQueryParams.Add("backup_type", parameterToString(*r.backupType, ""))
	}
	if r.backupLevel != nil {
		localVarQueryParams.Add("backup_level", parameterToString(*r.backupLevel, ""))
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

type ApiApiBackupsIdDeleteRequest struct {
	ctx context.Context
	ApiService *BackupsApiService
	id string
	localOnly *bool
}

// Set to true to only delete the Backup object in the database (does not delete any actual resources)
func (r ApiApiBackupsIdDeleteRequest) LocalOnly(localOnly bool) ApiApiBackupsIdDeleteRequest {
	r.localOnly = &localOnly
	return r
}

func (r ApiApiBackupsIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiBackupsIdDeleteExecute(r)
}

/*
ApiBackupsIdDelete Delete Backup

Deletes an existing database deployment Backup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Backup ID (must be a valid UUID)
 @return ApiApiBackupsIdDeleteRequest
*/
func (a *BackupsApiService) ApiBackupsIdDelete(ctx context.Context, id string) ApiApiBackupsIdDeleteRequest {
	return ApiApiBackupsIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *BackupsApiService) ApiBackupsIdDeleteExecute(r ApiApiBackupsIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupsApiService.ApiBackupsIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/backups/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	if r.localOnly != nil {
		localVarHeaderParams["Local-Only"] = parameterToString(*r.localOnly, "")
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

type ApiApiBackupsIdGetRequest struct {
	ctx context.Context
	ApiService *BackupsApiService
	id string
}


func (r ApiApiBackupsIdGetRequest) Execute() (*ModelsBackup, *http.Response, error) {
	return r.ApiService.ApiBackupsIdGetExecute(r)
}

/*
ApiBackupsIdGet Get Backup

Fetches a single Backup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Backup ID (must be a valid UUID)
 @return ApiApiBackupsIdGetRequest
*/
func (a *BackupsApiService) ApiBackupsIdGet(ctx context.Context, id string) ApiApiBackupsIdGetRequest {
	return ApiApiBackupsIdGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsBackup
func (a *BackupsApiService) ApiBackupsIdGetExecute(r ApiApiBackupsIdGetRequest) (*ModelsBackup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsBackup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupsApiService.ApiBackupsIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/backups/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiApiBackupsIdJobsNameDeleteRequest struct {
	ctx context.Context
	ApiService *BackupsApiService
	id string
	name string
	localOnly *bool
}

// Set to true to only delete the Backup object in the database (does not delete any actual resources)
func (r ApiApiBackupsIdJobsNameDeleteRequest) LocalOnly(localOnly bool) ApiApiBackupsIdJobsNameDeleteRequest {
	r.localOnly = &localOnly
	return r
}

func (r ApiApiBackupsIdJobsNameDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiBackupsIdJobsNameDeleteExecute(r)
}

/*
ApiBackupsIdJobsNameDelete Delete Backup jobs

Deletes an existing job for scheduled backups

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Backup ID (must be a valid UUID)
 @param name Backup job name
 @return ApiApiBackupsIdJobsNameDeleteRequest
*/
func (a *BackupsApiService) ApiBackupsIdJobsNameDelete(ctx context.Context, id string, name string) ApiApiBackupsIdJobsNameDeleteRequest {
	return ApiApiBackupsIdJobsNameDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		name: name,
	}
}

// Execute executes the request
func (a *BackupsApiService) ApiBackupsIdJobsNameDeleteExecute(r ApiApiBackupsIdJobsNameDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupsApiService.ApiBackupsIdJobsNameDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/backups/{id}/jobs/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterToString(r.name, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	if r.localOnly != nil {
		localVarHeaderParams["Local-Only"] = parameterToString(*r.localOnly, "")
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

type ApiApiBackupsIdPutRequest struct {
	ctx context.Context
	ApiService *BackupsApiService
	id string
	body *ControllersUpdateBackupRequest
	localOnly *bool
}

// Request body containing updated backup
func (r ApiApiBackupsIdPutRequest) Body(body ControllersUpdateBackupRequest) ApiApiBackupsIdPutRequest {
	r.body = &body
	return r
}
// Set to true to only update the Backup object in the database (does not create any actual resources)
func (r ApiApiBackupsIdPutRequest) LocalOnly(localOnly bool) ApiApiBackupsIdPutRequest {
	r.localOnly = &localOnly
	return r
}

func (r ApiApiBackupsIdPutRequest) Execute() (*ModelsBackup, *http.Response, error) {
	return r.ApiService.ApiBackupsIdPutExecute(r)
}

/*
ApiBackupsIdPut Update Backup

Updates an existing database Backup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Backup ID (must be UUID)
 @return ApiApiBackupsIdPutRequest
*/
func (a *BackupsApiService) ApiBackupsIdPut(ctx context.Context, id string) ApiApiBackupsIdPutRequest {
	return ApiApiBackupsIdPutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsBackup
func (a *BackupsApiService) ApiBackupsIdPutExecute(r ApiApiBackupsIdPutRequest) (*ModelsBackup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsBackup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupsApiService.ApiBackupsIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/backups/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

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
	if r.localOnly != nil {
		localVarHeaderParams["Local-Only"] = parameterToString(*r.localOnly, "")
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

type ApiApiDeploymentsIdBackupsGetRequest struct {
	ctx context.Context
	ApiService *BackupsApiService
	id string
	sortBy *string
	limit *string
	continuation *string
	id2 *string
	clusterResourceName *string
	state *string
	suspend *string
	backupType *string
	backupLevel *string
}

// A given Backup attribute to sort results by (one of: id, cluster_resource_name, created_at, backup_type, backup_level)
func (r ApiApiDeploymentsIdBackupsGetRequest) SortBy(sortBy string) ApiApiDeploymentsIdBackupsGetRequest {
	r.sortBy = &sortBy
	return r
}
// Maximum number of rows to return (could be less)
func (r ApiApiDeploymentsIdBackupsGetRequest) Limit(limit string) ApiApiDeploymentsIdBackupsGetRequest {
	r.limit = &limit
	return r
}
// Use a token returned by a previous query to continue listing with the next batch of rows
func (r ApiApiDeploymentsIdBackupsGetRequest) Continuation(continuation string) ApiApiDeploymentsIdBackupsGetRequest {
	r.continuation = &continuation
	return r
}
// Filter results by Backup id
func (r ApiApiDeploymentsIdBackupsGetRequest) Id2(id2 string) ApiApiDeploymentsIdBackupsGetRequest {
	r.id2 = &id2
	return r
}
// Filter results by Backup cluster_resource_name
func (r ApiApiDeploymentsIdBackupsGetRequest) ClusterResourceName(clusterResourceName string) ApiApiDeploymentsIdBackupsGetRequest {
	r.clusterResourceName = &clusterResourceName
	return r
}
// Filter results by Backup state
func (r ApiApiDeploymentsIdBackupsGetRequest) State(state string) ApiApiDeploymentsIdBackupsGetRequest {
	r.state = &state
	return r
}
// Filter results by Backup suspend flag
func (r ApiApiDeploymentsIdBackupsGetRequest) Suspend(suspend string) ApiApiDeploymentsIdBackupsGetRequest {
	r.suspend = &suspend
	return r
}
// Filter results by Backup type (one of: adhoc,scheduled)
func (r ApiApiDeploymentsIdBackupsGetRequest) BackupType(backupType string) ApiApiDeploymentsIdBackupsGetRequest {
	r.backupType = &backupType
	return r
}
// Filter results by Backup type (one of: snapshot,incremental)
func (r ApiApiDeploymentsIdBackupsGetRequest) BackupLevel(backupLevel string) ApiApiDeploymentsIdBackupsGetRequest {
	r.backupLevel = &backupLevel
	return r
}

func (r ApiApiDeploymentsIdBackupsGetRequest) Execute() (*ModelsPaginatedResultModelsBackup, *http.Response, error) {
	return r.ApiService.ApiDeploymentsIdBackupsGetExecute(r)
}

/*
ApiDeploymentsIdBackupsGet List Deployment's Backups

Lists Backups belonging to the Deployment.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Deployment ID (must be valid UUID)
 @return ApiApiDeploymentsIdBackupsGetRequest
*/
func (a *BackupsApiService) ApiDeploymentsIdBackupsGet(ctx context.Context, id string) ApiApiDeploymentsIdBackupsGetRequest {
	return ApiApiDeploymentsIdBackupsGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsPaginatedResultModelsBackup
func (a *BackupsApiService) ApiDeploymentsIdBackupsGetExecute(r ApiApiDeploymentsIdBackupsGetRequest) (*ModelsPaginatedResultModelsBackup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsPaginatedResultModelsBackup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupsApiService.ApiDeploymentsIdBackupsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/deployments/{id}/backups"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sortBy != nil {
		localVarQueryParams.Add("sort_by", parameterToString(*r.sortBy, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.continuation != nil {
		localVarQueryParams.Add("continuation", parameterToString(*r.continuation, ""))
	}
	if r.id2 != nil {
		localVarQueryParams.Add("id", parameterToString(*r.id2, ""))
	}
	if r.clusterResourceName != nil {
		localVarQueryParams.Add("cluster_resource_name", parameterToString(*r.clusterResourceName, ""))
	}
	if r.state != nil {
		localVarQueryParams.Add("state", parameterToString(*r.state, ""))
	}
	if r.suspend != nil {
		localVarQueryParams.Add("suspend", parameterToString(*r.suspend, ""))
	}
	if r.backupType != nil {
		localVarQueryParams.Add("backup_type", parameterToString(*r.backupType, ""))
	}
	if r.backupLevel != nil {
		localVarQueryParams.Add("backup_level", parameterToString(*r.backupLevel, ""))
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

type ApiApiDeploymentsIdBackupsPostRequest struct {
	ctx context.Context
	ApiService *BackupsApiService
	id string
	body *ControllersCreateDeploymentBackup
	localOnly *bool
}

// Request body containing the backup config
func (r ApiApiDeploymentsIdBackupsPostRequest) Body(body ControllersCreateDeploymentBackup) ApiApiDeploymentsIdBackupsPostRequest {
	r.body = &body
	return r
}
// Set to true to only create the Backup object in the database (does not create any actual resources)
func (r ApiApiDeploymentsIdBackupsPostRequest) LocalOnly(localOnly bool) ApiApiDeploymentsIdBackupsPostRequest {
	r.localOnly = &localOnly
	return r
}

func (r ApiApiDeploymentsIdBackupsPostRequest) Execute() (*ModelsBackup, *http.Response, error) {
	return r.ApiService.ApiDeploymentsIdBackupsPostExecute(r)
}

/*
ApiDeploymentsIdBackupsPost Create Backup

Creates a new database Backup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Deployment ID (must be valid UUID)
 @return ApiApiDeploymentsIdBackupsPostRequest
*/
func (a *BackupsApiService) ApiDeploymentsIdBackupsPost(ctx context.Context, id string) ApiApiDeploymentsIdBackupsPostRequest {
	return ApiApiDeploymentsIdBackupsPostRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsBackup
func (a *BackupsApiService) ApiDeploymentsIdBackupsPostExecute(r ApiApiDeploymentsIdBackupsPostRequest) (*ModelsBackup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsBackup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BackupsApiService.ApiDeploymentsIdBackupsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/deployments/{id}/backups"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

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
	if r.localOnly != nil {
		localVarHeaderParams["Local-Only"] = parameterToString(*r.localOnly, "")
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
