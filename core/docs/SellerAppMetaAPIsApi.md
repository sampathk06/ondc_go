# \SellerAppMetaAPIsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCancellationReasonsPost**](SellerAppMetaAPIsApi.md#GetCancellationReasonsPost) | **Post** /get_cancellation_reasons | 
[**GetFeedbackFormPost**](SellerAppMetaAPIsApi.md#GetFeedbackFormPost) | **Post** /get_feedback_form | 
[**GetRatingCategoriesPost**](SellerAppMetaAPIsApi.md#GetRatingCategoriesPost) | **Post** /get_rating_categories | 
[**GetReturnReasonsPost**](SellerAppMetaAPIsApi.md#GetReturnReasonsPost) | **Post** /get_return_reasons | 



## GetCancellationReasonsPost

> SearchPost200Response GetCancellationReasonsPost(ctx).GetCancellationReasonsPostRequest(getCancellationReasonsPostRequest).Execute()





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
    getCancellationReasonsPostRequest := *openapiclient.NewGetCancellationReasonsPostRequest() // GetCancellationReasonsPostRequest | Context header is sent as the request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SellerAppMetaAPIsApi.GetCancellationReasonsPost(context.Background()).GetCancellationReasonsPostRequest(getCancellationReasonsPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SellerAppMetaAPIsApi.GetCancellationReasonsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCancellationReasonsPost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `SellerAppMetaAPIsApi.GetCancellationReasonsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCancellationReasonsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getCancellationReasonsPostRequest** | [**GetCancellationReasonsPostRequest**](GetCancellationReasonsPostRequest.md) | Context header is sent as the request | 

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


## GetFeedbackFormPost

> SearchPost200Response GetFeedbackFormPost(ctx).GetFeedbackFormPostRequest(getFeedbackFormPostRequest).Execute()





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
    getFeedbackFormPostRequest := *openapiclient.NewGetFeedbackFormPostRequest() // GetFeedbackFormPostRequest | The rating value and category is sent by the Buyer App (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SellerAppMetaAPIsApi.GetFeedbackFormPost(context.Background()).GetFeedbackFormPostRequest(getFeedbackFormPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SellerAppMetaAPIsApi.GetFeedbackFormPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeedbackFormPost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `SellerAppMetaAPIsApi.GetFeedbackFormPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFeedbackFormPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getFeedbackFormPostRequest** | [**GetFeedbackFormPostRequest**](GetFeedbackFormPostRequest.md) | The rating value and category is sent by the Buyer App | 

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


## GetRatingCategoriesPost

> SearchPost200Response GetRatingCategoriesPost(ctx).GetCancellationReasonsPostRequest(getCancellationReasonsPostRequest).Execute()





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
    getCancellationReasonsPostRequest := *openapiclient.NewGetCancellationReasonsPostRequest() // GetCancellationReasonsPostRequest | Context header is sent as the request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SellerAppMetaAPIsApi.GetRatingCategoriesPost(context.Background()).GetCancellationReasonsPostRequest(getCancellationReasonsPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SellerAppMetaAPIsApi.GetRatingCategoriesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRatingCategoriesPost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `SellerAppMetaAPIsApi.GetRatingCategoriesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRatingCategoriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getCancellationReasonsPostRequest** | [**GetCancellationReasonsPostRequest**](GetCancellationReasonsPostRequest.md) | Context header is sent as the request | 

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


## GetReturnReasonsPost

> SearchPost200Response GetReturnReasonsPost(ctx).GetCancellationReasonsPostRequest(getCancellationReasonsPostRequest).Execute()





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
    getCancellationReasonsPostRequest := *openapiclient.NewGetCancellationReasonsPostRequest() // GetCancellationReasonsPostRequest | Context header is sent as the request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SellerAppMetaAPIsApi.GetReturnReasonsPost(context.Background()).GetCancellationReasonsPostRequest(getCancellationReasonsPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SellerAppMetaAPIsApi.GetReturnReasonsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReturnReasonsPost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `SellerAppMetaAPIsApi.GetReturnReasonsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReturnReasonsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getCancellationReasonsPostRequest** | [**GetCancellationReasonsPostRequest**](GetCancellationReasonsPostRequest.md) | Context header is sent as the request | 

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

