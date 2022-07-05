# \SandboxesApi

All URIs are relative to *http://api.signadot.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewSandbox**](SandboxesApi.md#CreateNewSandbox) | **Post** /orgs/{orgName}/sandboxes | Create a new sandbox
[**DeleteSandboxById**](SandboxesApi.md#DeleteSandboxById) | **Delete** /orgs/{orgName}/sandboxes/{sandboxID} | Delete a Sandbox by ID
[**DeleteSandboxByName**](SandboxesApi.md#DeleteSandboxByName) | **Delete** /orgs/{orgName}/sandboxes/by-name/{name} | Delete Sandbox By Name
[**GetSandboxById**](SandboxesApi.md#GetSandboxById) | **Get** /orgs/{orgName}/sandboxes/{sandboxID} | Get a Sandbox by ID
[**GetSandboxReady**](SandboxesApi.md#GetSandboxReady) | **Get** /orgs/{orgName}/sandboxes/{sandboxID}/ready | Check sandbox readiness
[**GetSandboxStatusById**](SandboxesApi.md#GetSandboxStatusById) | **Get** /orgs/{orgName}/sandboxes/{sandboxID}/status | Get Sandbox Status by Sandbox ID with rate limiting.
[**GetSandboxes**](SandboxesApi.md#GetSandboxes) | **Get** /orgs/{orgName}/sandboxes | List Sandboxes
[**UpsertPrWorkspace**](SandboxesApi.md#UpsertPrWorkspace) | **Post** /repos/{githubOrg}/{githubRepo}/pulls/{prNumber}/workspaces | Create or Update workspace from Pull Request



## CreateNewSandbox

> CreateSandboxResponse CreateNewSandbox(ctx, orgName).Data(data).Execute()

Create a new sandbox



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgName := "my-company" // string | Signadot Org Name
    data := *openapiclient.NewCreateSandboxRequest("Cluster_example", []openapiclient.SandboxFork{*openapiclient.NewSandboxFork()}) // CreateSandboxRequest | Request to create sandbox

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SandboxesApi.CreateNewSandbox(context.Background(), orgName).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SandboxesApi.CreateNewSandbox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNewSandbox`: CreateSandboxResponse
    fmt.Fprintf(os.Stdout, "Response from `SandboxesApi.CreateNewSandbox`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgName** | **string** | Signadot Org Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewSandboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**CreateSandboxRequest**](CreateSandboxRequest.md) | Request to create sandbox | 

### Return type

[**CreateSandboxResponse**](CreateSandboxResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSandboxById

> map[string]interface{} DeleteSandboxById(ctx, orgName, sandboxID).Execute()

Delete a Sandbox by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgName := "my-company" // string | Signadot Org Name
    sandboxID := "8r32dkdgdg" // string | Sandbox ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SandboxesApi.DeleteSandboxById(context.Background(), orgName, sandboxID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SandboxesApi.DeleteSandboxById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSandboxById`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SandboxesApi.DeleteSandboxById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgName** | **string** | Signadot Org Name | 
**sandboxID** | **string** | Sandbox ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSandboxByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSandboxByName

> map[string]interface{} DeleteSandboxByName(ctx, orgName, name).Execute()

Delete Sandbox By Name



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgName := "my-company" // string | Signadot Org Name
    name := "name_example" // string | Sandbox Name to search for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SandboxesApi.DeleteSandboxByName(context.Background(), orgName, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SandboxesApi.DeleteSandboxByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSandboxByName`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SandboxesApi.DeleteSandboxByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgName** | **string** | Signadot Org Name | 
**name** | **string** | Sandbox Name to search for | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSandboxByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSandboxById

> GetSandboxByIdResponse GetSandboxById(ctx, orgName, sandboxID).Execute()

Get a Sandbox by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgName := "my-company" // string | Signadot Org Name
    sandboxID := "8r32dkdgdg" // string | Sandbox ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SandboxesApi.GetSandboxById(context.Background(), orgName, sandboxID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SandboxesApi.GetSandboxById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSandboxById`: GetSandboxByIdResponse
    fmt.Fprintf(os.Stdout, "Response from `SandboxesApi.GetSandboxById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgName** | **string** | Signadot Org Name | 
**sandboxID** | **string** | Sandbox ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSandboxByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetSandboxByIdResponse**](GetSandboxByIdResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSandboxReady

> SandboxReadyResponse GetSandboxReady(ctx, orgName, sandboxID).Execute()

Check sandbox readiness



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgName := "my-company" // string | Signadot Org Name
    sandboxID := "sandboxID_example" // string | Sandbox ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SandboxesApi.GetSandboxReady(context.Background(), orgName, sandboxID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SandboxesApi.GetSandboxReady``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSandboxReady`: SandboxReadyResponse
    fmt.Fprintf(os.Stdout, "Response from `SandboxesApi.GetSandboxReady`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgName** | **string** | Signadot Org Name | 
**sandboxID** | **string** | Sandbox ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSandboxReadyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SandboxReadyResponse**](SandboxReadyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSandboxStatusById

> SandboxStatusResponse GetSandboxStatusById(ctx, orgName, sandboxID).Execute()

Get Sandbox Status by Sandbox ID with rate limiting.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgName := "my-company" // string | Signadot Org Name
    sandboxID := "sandboxID_example" // string | Sandbox ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SandboxesApi.GetSandboxStatusById(context.Background(), orgName, sandboxID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SandboxesApi.GetSandboxStatusById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSandboxStatusById`: SandboxStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `SandboxesApi.GetSandboxStatusById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgName** | **string** | Signadot Org Name | 
**sandboxID** | **string** | Sandbox ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSandboxStatusByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SandboxStatusResponse**](SandboxStatusResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSandboxes

> GetSandboxesResponse GetSandboxes(ctx, orgName).Execute()

List Sandboxes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgName := "my-company" // string | Signadot Org Name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SandboxesApi.GetSandboxes(context.Background(), orgName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SandboxesApi.GetSandboxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSandboxes`: GetSandboxesResponse
    fmt.Fprintf(os.Stdout, "Response from `SandboxesApi.GetSandboxes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgName** | **string** | Signadot Org Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSandboxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSandboxesResponse**](GetSandboxesResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertPrWorkspace

> UpsertWorkspaceResponse UpsertPrWorkspace(ctx, githubOrg, githubRepo, prNumber).Data(data).Execute()

Create or Update workspace from Pull Request



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    githubOrg := "githubOrg_example" // string | GitHub Org Name
    githubRepo := "githubRepo_example" // string | GitHub Repository Name
    prNumber := int32(123) // int32 | Pull Request Number
    data := *openapiclient.NewUpsertPRWorkspacesRequest("Cluster_example", "Namespace_example") // UpsertPRWorkspacesRequest | Request to upsert workspace

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SandboxesApi.UpsertPrWorkspace(context.Background(), githubOrg, githubRepo, prNumber).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SandboxesApi.UpsertPrWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpsertPrWorkspace`: UpsertWorkspaceResponse
    fmt.Fprintf(os.Stdout, "Response from `SandboxesApi.UpsertPrWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**githubOrg** | **string** | GitHub Org Name | 
**githubRepo** | **string** | GitHub Repository Name | 
**prNumber** | **int32** | Pull Request Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpsertPrWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **data** | [**UpsertPRWorkspacesRequest**](UpsertPRWorkspacesRequest.md) | Request to upsert workspace | 

### Return type

[**UpsertWorkspaceResponse**](UpsertWorkspaceResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

