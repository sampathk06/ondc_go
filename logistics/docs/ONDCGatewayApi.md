# \ONDCGatewayApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OnSearchPost**](ONDCGatewayApi.md#OnSearchPost) | **Post** /on_search | 
[**SearchPost**](ONDCGatewayApi.md#SearchPost) | **Post** /search | 



## OnSearchPost

> OnSearchPost200Response OnSearchPost(ctx).OnSearchPostRequest(onSearchPostRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    onSearchPostRequest := *openapiclient.NewOnSearchPostRequest(*openapiclient.NewContext(openapiclient.Domain("nic2004:60232"), "TODO", "TODO", "Action_example", "CoreVersion_example", "BapId_example", "BapUri_example", "TransactionId_example", "MessageId_example", time.Now())) // OnSearchPostRequest | Sellers provide their catalog in response to buyer search (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ONDCGatewayApi.OnSearchPost(context.Background()).OnSearchPostRequest(onSearchPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ONDCGatewayApi.OnSearchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OnSearchPost`: OnSearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `ONDCGatewayApi.OnSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onSearchPostRequest** | [**OnSearchPostRequest**](OnSearchPostRequest.md) | Sellers provide their catalog in response to buyer search | 

### Return type

[**OnSearchPost200Response**](OnSearchPost200Response.md)

### Authorization

[GatewaySubscriberAuth](../README.md#GatewaySubscriberAuth), [GatewaySubscriberAuthNew](../README.md#GatewaySubscriberAuthNew), [SubscriberAuth](../README.md#SubscriberAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchPost

> SearchPost200Response SearchPost(ctx).SearchPostRequest(searchPostRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    searchPostRequest := *openapiclient.NewSearchPostRequest(*openapiclient.NewContext(openapiclient.Domain("nic2004:60232"), "TODO", "TODO", "Action_example", "CoreVersion_example", "BapId_example", "BapUri_example", "TransactionId_example", "MessageId_example", time.Now()), *openapiclient.NewSearchPostRequestMessage()) // SearchPostRequest | Buyer searches for products and services (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ONDCGatewayApi.SearchPost(context.Background()).SearchPostRequest(searchPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ONDCGatewayApi.SearchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchPost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `ONDCGatewayApi.SearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchPostRequest** | [**SearchPostRequest**](SearchPostRequest.md) | Buyer searches for products and services | 

### Return type

[**SearchPost200Response**](SearchPost200Response.md)

### Authorization

[GatewaySubscriberAuth](../README.md#GatewaySubscriberAuth), [GatewaySubscriberAuthNew](../README.md#GatewaySubscriberAuthNew), [SubscriberAuth](../README.md#SubscriberAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

