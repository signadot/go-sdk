# \ClusterApi

All URIs are relative to *http://api.signadot.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectCluster**](ClusterApi.md#ConnectCluster) | **Post** /orgs/{orgName}/clusters | Connect Cluster
[**CreateClusterToken**](ClusterApi.md#CreateClusterToken) | **Post** /orgs/{orgName}/clusters/{clusterName}/tokens | Create Cluster Token
[**DeleteClusterToken**](ClusterApi.md#DeleteClusterToken) | **Delete** /orgs/{orgName}/clusters/{clusterName}/tokens/{tokenId} | Delete Cluster Token
[**GetClusters**](ClusterApi.md#GetClusters) | **Get** /orgs/{orgName}/clusters | List clusters



## ConnectCluster

> ConnectClusterResponse ConnectCluster(ctx, orgName).Data(data).Execute()

Connect Cluster



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
    data := *openapiclient.NewConnectClusterRequest("my-staging-cluster") // ConnectClusterRequest | Request to create cluster

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterApi.ConnectCluster(context.Background(), orgName).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterApi.ConnectCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectCluster`: ConnectClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `ClusterApi.ConnectCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgName** | **string** | Signadot Org Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**ConnectClusterRequest**](ConnectClusterRequest.md) | Request to create cluster | 

### Return type

[**ConnectClusterResponse**](ConnectClusterResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateClusterToken

> CreateClusterTokenResponse CreateClusterToken(ctx, orgName, clusterName).Execute()

Create Cluster Token



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
    clusterName := "my-new-cluster" // string | Cluster Name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterApi.CreateClusterToken(context.Background(), orgName, clusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterApi.CreateClusterToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateClusterToken`: CreateClusterTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `ClusterApi.CreateClusterToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgName** | **string** | Signadot Org Name | 
**clusterName** | **string** | Cluster Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateClusterTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateClusterTokenResponse**](CreateClusterTokenResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClusterToken

> map[string]interface{} DeleteClusterToken(ctx, orgName, clusterName, tokenId).Execute()

Delete Cluster Token



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
    clusterName := "my-new-cluster" // string | Cluster Name
    tokenId := "kq6mtksk7mn5" // string | Token Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterApi.DeleteClusterToken(context.Background(), orgName, clusterName, tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterApi.DeleteClusterToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteClusterToken`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ClusterApi.DeleteClusterToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgName** | **string** | Signadot Org Name | 
**clusterName** | **string** | Cluster Name | 
**tokenId** | **string** | Token Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterTokenRequest struct via the builder pattern


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


## GetClusters

> GetClustersResponse GetClusters(ctx, orgName).Execute()

List clusters



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
    resp, r, err := apiClient.ClusterApi.GetClusters(context.Background(), orgName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterApi.GetClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusters`: GetClustersResponse
    fmt.Fprintf(os.Stdout, "Response from `ClusterApi.GetClusters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgName** | **string** | Signadot Org Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetClustersResponse**](GetClustersResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

