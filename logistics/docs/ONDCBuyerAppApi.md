# \ONDCBuyerAppApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OnCancelPost**](ONDCBuyerAppApi.md#OnCancelPost) | **Post** /on_cancel | 
[**OnConfirmPost**](ONDCBuyerAppApi.md#OnConfirmPost) | **Post** /on_confirm | 
[**OnInitPost**](ONDCBuyerAppApi.md#OnInitPost) | **Post** /on_init | 
[**OnRatingPost**](ONDCBuyerAppApi.md#OnRatingPost) | **Post** /on_rating | 
[**OnSearchPost**](ONDCBuyerAppApi.md#OnSearchPost) | **Post** /on_search | 
[**OnSelectPost**](ONDCBuyerAppApi.md#OnSelectPost) | **Post** /on_select | 
[**OnStatusPost**](ONDCBuyerAppApi.md#OnStatusPost) | **Post** /on_status | 
[**OnSupportPost**](ONDCBuyerAppApi.md#OnSupportPost) | **Post** /on_support | 
[**OnTrackPost**](ONDCBuyerAppApi.md#OnTrackPost) | **Post** /on_track | 
[**OnUpdatePost**](ONDCBuyerAppApi.md#OnUpdatePost) | **Post** /on_update | 



## OnCancelPost

> SearchPost200Response OnCancelPost(ctx).OnConfirmPostRequest(onConfirmPostRequest).Execute()





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
    onConfirmPostRequest := *openapiclient.NewOnConfirmPostRequest(*openapiclient.NewContext(openapiclient.Domain("nic2004:60232"), "TODO", "TODO", "Action_example", "CoreVersion_example", "BapId_example", "BapUri_example", "TransactionId_example", "MessageId_example", time.Now())) // OnConfirmPostRequest | Seller provides response to cancellation request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ONDCBuyerAppApi.OnCancelPost(context.Background()).OnConfirmPostRequest(onConfirmPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ONDCBuyerAppApi.OnCancelPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OnCancelPost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `ONDCBuyerAppApi.OnCancelPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onConfirmPostRequest** | [**OnConfirmPostRequest**](OnConfirmPostRequest.md) | Seller provides response to cancellation request | 

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


## OnConfirmPost

> SearchPost200Response OnConfirmPost(ctx).OnConfirmPostRequest(onConfirmPostRequest).Execute()





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
    onConfirmPostRequest := *openapiclient.NewOnConfirmPostRequest(*openapiclient.NewContext(openapiclient.Domain("nic2004:60232"), "TODO", "TODO", "Action_example", "CoreVersion_example", "BapId_example", "BapUri_example", "TransactionId_example", "MessageId_example", time.Now())) // OnConfirmPostRequest | Seller confirms order (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ONDCBuyerAppApi.OnConfirmPost(context.Background()).OnConfirmPostRequest(onConfirmPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ONDCBuyerAppApi.OnConfirmPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OnConfirmPost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `ONDCBuyerAppApi.OnConfirmPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnConfirmPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onConfirmPostRequest** | [**OnConfirmPostRequest**](OnConfirmPostRequest.md) | Seller confirms order | 

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


## OnInitPost

> SearchPost200Response OnInitPost(ctx).OnInitPostRequest(onInitPostRequest).Execute()





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
    onInitPostRequest := *openapiclient.NewOnInitPostRequest(*openapiclient.NewContext(openapiclient.Domain("nic2004:60232"), "TODO", "TODO", "Action_example", "CoreVersion_example", "BapId_example", "BapUri_example", "TransactionId_example", "MessageId_example", time.Now())) // OnInitPostRequest | Seller provides terms and conditions for an order (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ONDCBuyerAppApi.OnInitPost(context.Background()).OnInitPostRequest(onInitPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ONDCBuyerAppApi.OnInitPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OnInitPost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `ONDCBuyerAppApi.OnInitPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnInitPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onInitPostRequest** | [**OnInitPostRequest**](OnInitPostRequest.md) | Seller provides terms and conditions for an order | 

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


## OnRatingPost

> SearchPost200Response OnRatingPost(ctx).OnRatingPostRequest(onRatingPostRequest).Execute()





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
    onRatingPostRequest := *openapiclient.NewOnRatingPostRequest(*openapiclient.NewContext(openapiclient.Domain("nic2004:60232"), "TODO", "TODO", "Action_example", "CoreVersion_example", "BapId_example", "BapUri_example", "TransactionId_example", "MessageId_example", time.Now())) // OnRatingPostRequest | Seller provides response to rating provided by buyer (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ONDCBuyerAppApi.OnRatingPost(context.Background()).OnRatingPostRequest(onRatingPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ONDCBuyerAppApi.OnRatingPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OnRatingPost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `ONDCBuyerAppApi.OnRatingPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnRatingPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onRatingPostRequest** | [**OnRatingPostRequest**](OnRatingPostRequest.md) | Seller provides response to rating provided by buyer | 

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
    resp, r, err := apiClient.ONDCBuyerAppApi.OnSearchPost(context.Background()).OnSearchPostRequest(onSearchPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ONDCBuyerAppApi.OnSearchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OnSearchPost`: OnSearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `ONDCBuyerAppApi.OnSearchPost`: %v\n", resp)
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


## OnSelectPost

> SearchPost200Response OnSelectPost(ctx).OnSelectPostRequest(onSelectPostRequest).Execute()





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
    onSelectPostRequest := *openapiclient.NewOnSelectPostRequest(*openapiclient.NewContext(openapiclient.Domain("nic2004:60232"), "TODO", "TODO", "Action_example", "CoreVersion_example", "BapId_example", "BapUri_example", "TransactionId_example", "MessageId_example", time.Now())) // OnSelectPostRequest | Seller provides quote for selected items (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ONDCBuyerAppApi.OnSelectPost(context.Background()).OnSelectPostRequest(onSelectPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ONDCBuyerAppApi.OnSelectPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OnSelectPost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `ONDCBuyerAppApi.OnSelectPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnSelectPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onSelectPostRequest** | [**OnSelectPostRequest**](OnSelectPostRequest.md) | Seller provides quote for selected items | 

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


## OnStatusPost

> SearchPost200Response OnStatusPost(ctx).OnConfirmPostRequest(onConfirmPostRequest).Execute()





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
    onConfirmPostRequest := *openapiclient.NewOnConfirmPostRequest(*openapiclient.NewContext(openapiclient.Domain("nic2004:60232"), "TODO", "TODO", "Action_example", "CoreVersion_example", "BapId_example", "BapUri_example", "TransactionId_example", "MessageId_example", time.Now())) // OnConfirmPostRequest | Seller provides status information for order (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ONDCBuyerAppApi.OnStatusPost(context.Background()).OnConfirmPostRequest(onConfirmPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ONDCBuyerAppApi.OnStatusPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OnStatusPost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `ONDCBuyerAppApi.OnStatusPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnStatusPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onConfirmPostRequest** | [**OnConfirmPostRequest**](OnConfirmPostRequest.md) | Seller provides status information for order | 

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


## OnSupportPost

> SearchPost200Response OnSupportPost(ctx).OnSupportPostRequest(onSupportPostRequest).Execute()





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
    onSupportPostRequest := *openapiclient.NewOnSupportPostRequest(*openapiclient.NewContext(openapiclient.Domain("nic2004:60232"), "TODO", "TODO", "Action_example", "CoreVersion_example", "BapId_example", "BapUri_example", "TransactionId_example", "MessageId_example", time.Now())) // OnSupportPostRequest | Seller provides Contact Support details (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ONDCBuyerAppApi.OnSupportPost(context.Background()).OnSupportPostRequest(onSupportPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ONDCBuyerAppApi.OnSupportPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OnSupportPost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `ONDCBuyerAppApi.OnSupportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnSupportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onSupportPostRequest** | [**OnSupportPostRequest**](OnSupportPostRequest.md) | Seller provides Contact Support details | 

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


## OnTrackPost

> SearchPost200Response OnTrackPost(ctx).OnTrackPostRequest(onTrackPostRequest).Execute()





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
    onTrackPostRequest := *openapiclient.NewOnTrackPostRequest(*openapiclient.NewContext(openapiclient.Domain("nic2004:60232"), "TODO", "TODO", "Action_example", "CoreVersion_example", "BapId_example", "BapUri_example", "TransactionId_example", "MessageId_example", time.Now())) // OnTrackPostRequest | Seller provides tracking details for an order (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ONDCBuyerAppApi.OnTrackPost(context.Background()).OnTrackPostRequest(onTrackPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ONDCBuyerAppApi.OnTrackPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OnTrackPost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `ONDCBuyerAppApi.OnTrackPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnTrackPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onTrackPostRequest** | [**OnTrackPostRequest**](OnTrackPostRequest.md) | Seller provides tracking details for an order | 

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


## OnUpdatePost

> SearchPost200Response OnUpdatePost(ctx).OnConfirmPostRequest(onConfirmPostRequest).Execute()





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
    onConfirmPostRequest := *openapiclient.NewOnConfirmPostRequest(*openapiclient.NewContext(openapiclient.Domain("nic2004:60232"), "TODO", "TODO", "Action_example", "CoreVersion_example", "BapId_example", "BapUri_example", "TransactionId_example", "MessageId_example", time.Now())) // OnConfirmPostRequest | Seller provides response to order update (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ONDCBuyerAppApi.OnUpdatePost(context.Background()).OnConfirmPostRequest(onConfirmPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ONDCBuyerAppApi.OnUpdatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OnUpdatePost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `ONDCBuyerAppApi.OnUpdatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOnUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onConfirmPostRequest** | [**OnConfirmPostRequest**](OnConfirmPostRequest.md) | Seller provides response to order update | 

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

