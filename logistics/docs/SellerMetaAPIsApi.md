# \SellerMetaAPIsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFeedbackCategoriesPost**](SellerMetaAPIsApi.md#GetFeedbackCategoriesPost) | **Post** /get_feedback_categories | 



## GetFeedbackCategoriesPost

> SearchPost200Response GetFeedbackCategoriesPost(ctx).GetCancellationReasonsPostRequest(getCancellationReasonsPostRequest).Execute()





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
    resp, r, err := apiClient.SellerMetaAPIsApi.GetFeedbackCategoriesPost(context.Background()).GetCancellationReasonsPostRequest(getCancellationReasonsPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SellerMetaAPIsApi.GetFeedbackCategoriesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeedbackCategoriesPost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `SellerMetaAPIsApi.GetFeedbackCategoriesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFeedbackCategoriesPostRequest struct via the builder pattern


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

