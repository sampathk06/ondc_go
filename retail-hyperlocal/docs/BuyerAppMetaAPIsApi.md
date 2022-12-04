# \BuyerAppMetaAPIsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancellationReasonsPost**](BuyerAppMetaAPIsApi.md#CancellationReasonsPost) | **Post** /cancellation_reasons | 
[**FeedbackCategoriesPost**](BuyerAppMetaAPIsApi.md#FeedbackCategoriesPost) | **Post** /feedback_categories | 
[**FeedbackFormPost**](BuyerAppMetaAPIsApi.md#FeedbackFormPost) | **Post** /feedback_form | 
[**RatingCategoriesPost**](BuyerAppMetaAPIsApi.md#RatingCategoriesPost) | **Post** /rating_categories | 
[**ReturnReasonsPost**](BuyerAppMetaAPIsApi.md#ReturnReasonsPost) | **Post** /return_reasons | 



## CancellationReasonsPost

> SearchPost200Response CancellationReasonsPost(ctx).CancellationReasonsPostRequest(cancellationReasonsPostRequest).Execute()





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
    cancellationReasonsPostRequest := *openapiclient.NewCancellationReasonsPostRequest() // CancellationReasonsPostRequest | List of cancellation reasons. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerAppMetaAPIsApi.CancellationReasonsPost(context.Background()).CancellationReasonsPostRequest(cancellationReasonsPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerAppMetaAPIsApi.CancellationReasonsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancellationReasonsPost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `BuyerAppMetaAPIsApi.CancellationReasonsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancellationReasonsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cancellationReasonsPostRequest** | [**CancellationReasonsPostRequest**](CancellationReasonsPostRequest.md) | List of cancellation reasons. | 

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


## FeedbackCategoriesPost

> SearchPost200Response FeedbackCategoriesPost(ctx).FeedbackCategoriesPostRequest(feedbackCategoriesPostRequest).Execute()





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
    feedbackCategoriesPostRequest := *openapiclient.NewFeedbackCategoriesPostRequest() // FeedbackCategoriesPostRequest | Array of categories for which feedback can be given by the Buyer App (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerAppMetaAPIsApi.FeedbackCategoriesPost(context.Background()).FeedbackCategoriesPostRequest(feedbackCategoriesPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerAppMetaAPIsApi.FeedbackCategoriesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FeedbackCategoriesPost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `BuyerAppMetaAPIsApi.FeedbackCategoriesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFeedbackCategoriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feedbackCategoriesPostRequest** | [**FeedbackCategoriesPostRequest**](FeedbackCategoriesPostRequest.md) | Array of categories for which feedback can be given by the Buyer App | 

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


## FeedbackFormPost

> SearchPost200Response FeedbackFormPost(ctx).FeedbackFormPostRequest(feedbackFormPostRequest).Execute()





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
    feedbackFormPostRequest := *openapiclient.NewFeedbackFormPostRequest() // FeedbackFormPostRequest | Feedback form sent by the Buyer App (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerAppMetaAPIsApi.FeedbackFormPost(context.Background()).FeedbackFormPostRequest(feedbackFormPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerAppMetaAPIsApi.FeedbackFormPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FeedbackFormPost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `BuyerAppMetaAPIsApi.FeedbackFormPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFeedbackFormPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feedbackFormPostRequest** | [**FeedbackFormPostRequest**](FeedbackFormPostRequest.md) | Feedback form sent by the Buyer App | 

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


## RatingCategoriesPost

> SearchPost200Response RatingCategoriesPost(ctx).RatingCategoriesPostRequest(ratingCategoriesPostRequest).Execute()





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
    ratingCategoriesPostRequest := *openapiclient.NewRatingCategoriesPostRequest() // RatingCategoriesPostRequest | Array of categories which can be rated (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerAppMetaAPIsApi.RatingCategoriesPost(context.Background()).RatingCategoriesPostRequest(ratingCategoriesPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerAppMetaAPIsApi.RatingCategoriesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RatingCategoriesPost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `BuyerAppMetaAPIsApi.RatingCategoriesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRatingCategoriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ratingCategoriesPostRequest** | [**RatingCategoriesPostRequest**](RatingCategoriesPostRequest.md) | Array of categories which can be rated | 

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


## ReturnReasonsPost

> SearchPost200Response ReturnReasonsPost(ctx).ReturnReasonsPostRequest(returnReasonsPostRequest).Execute()





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
    returnReasonsPostRequest := *openapiclient.NewReturnReasonsPostRequest() // ReturnReasonsPostRequest | List of return reasons (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerAppMetaAPIsApi.ReturnReasonsPost(context.Background()).ReturnReasonsPostRequest(returnReasonsPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerAppMetaAPIsApi.ReturnReasonsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReturnReasonsPost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `BuyerAppMetaAPIsApi.ReturnReasonsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReturnReasonsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **returnReasonsPostRequest** | [**ReturnReasonsPostRequest**](ReturnReasonsPostRequest.md) | List of return reasons | 

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

