# \ONDCSellerAppApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelPost**](ONDCSellerAppApi.md#CancelPost) | **Post** /cancel | 
[**ConfirmPost**](ONDCSellerAppApi.md#ConfirmPost) | **Post** /confirm | 
[**InitPost**](ONDCSellerAppApi.md#InitPost) | **Post** /init | 
[**RatingPost**](ONDCSellerAppApi.md#RatingPost) | **Post** /rating | 
[**SearchPost**](ONDCSellerAppApi.md#SearchPost) | **Post** /search | 
[**SelectPost**](ONDCSellerAppApi.md#SelectPost) | **Post** /select | 
[**StatusPost**](ONDCSellerAppApi.md#StatusPost) | **Post** /status | 
[**SupportPost**](ONDCSellerAppApi.md#SupportPost) | **Post** /support | 
[**TrackPost**](ONDCSellerAppApi.md#TrackPost) | **Post** /track | 
[**UpdatePost**](ONDCSellerAppApi.md#UpdatePost) | **Post** /update | 



## CancelPost

> SearchPost200Response CancelPost(ctx).CancelPostRequest(cancelPostRequest).Execute()





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
    cancelPostRequest := *openapiclient.NewCancelPostRequest(*openapiclient.NewContext(openapiclient.Domain("nic2004:52110"), "TODO", "TODO", "Action_example", "CoreVersion_example", "BapId_example", "BapUri_example", "TransactionId_example", "MessageId_example", time.Now()), *openapiclient.NewCancelPostRequestMessage("TODO")) // CancelPostRequest | Buyer cancels an order (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ONDCSellerAppApi.CancelPost(context.Background()).CancelPostRequest(cancelPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ONDCSellerAppApi.CancelPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelPost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `ONDCSellerAppApi.CancelPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cancelPostRequest** | [**CancelPostRequest**](CancelPostRequest.md) | Buyer cancels an order | 

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


## ConfirmPost

> SearchPost200Response ConfirmPost(ctx).SelectPostRequest(selectPostRequest).Execute()





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
    selectPostRequest := *openapiclient.NewSelectPostRequest(*openapiclient.NewContext(openapiclient.Domain("nic2004:52110"), "TODO", "TODO", "Action_example", "CoreVersion_example", "BapId_example", "BapUri_example", "TransactionId_example", "MessageId_example", time.Now()), *openapiclient.NewSelectPostRequestMessage(*openapiclient.NewOrder())) // SelectPostRequest | Buyer confirms an order (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ONDCSellerAppApi.ConfirmPost(context.Background()).SelectPostRequest(selectPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ONDCSellerAppApi.ConfirmPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfirmPost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `ONDCSellerAppApi.ConfirmPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfirmPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selectPostRequest** | [**SelectPostRequest**](SelectPostRequest.md) | Buyer confirms an order | 

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


## InitPost

> SearchPost200Response InitPost(ctx).SelectPostRequest(selectPostRequest).Execute()





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
    selectPostRequest := *openapiclient.NewSelectPostRequest(*openapiclient.NewContext(openapiclient.Domain("nic2004:52110"), "TODO", "TODO", "Action_example", "CoreVersion_example", "BapId_example", "BapUri_example", "TransactionId_example", "MessageId_example", time.Now()), *openapiclient.NewSelectPostRequestMessage(*openapiclient.NewOrder())) // SelectPostRequest | Buyer initializes order checkout (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ONDCSellerAppApi.InitPost(context.Background()).SelectPostRequest(selectPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ONDCSellerAppApi.InitPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InitPost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `ONDCSellerAppApi.InitPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selectPostRequest** | [**SelectPostRequest**](SelectPostRequest.md) | Buyer initializes order checkout | 

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


## RatingPost

> SearchPost200Response RatingPost(ctx).RatingPostRequest(ratingPostRequest).Execute()





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
    ratingPostRequest := *openapiclient.NewRatingPostRequest(*openapiclient.NewContext(openapiclient.Domain("nic2004:52110"), "TODO", "TODO", "Action_example", "CoreVersion_example", "BapId_example", "BapUri_example", "TransactionId_example", "MessageId_example", time.Now()), *openapiclient.NewRating()) // RatingPostRequest | Buyer rates for one or more rating categories (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ONDCSellerAppApi.RatingPost(context.Background()).RatingPostRequest(ratingPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ONDCSellerAppApi.RatingPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RatingPost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `ONDCSellerAppApi.RatingPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRatingPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ratingPostRequest** | [**RatingPostRequest**](RatingPostRequest.md) | Buyer rates for one or more rating categories | 

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
    searchPostRequest := *openapiclient.NewSearchPostRequest(*openapiclient.NewContext(openapiclient.Domain("nic2004:52110"), "TODO", "TODO", "Action_example", "CoreVersion_example", "BapId_example", "BapUri_example", "TransactionId_example", "MessageId_example", time.Now()), *openapiclient.NewSearchPostRequestMessage()) // SearchPostRequest | Buyer searches for products and services (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ONDCSellerAppApi.SearchPost(context.Background()).SearchPostRequest(searchPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ONDCSellerAppApi.SearchPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchPost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `ONDCSellerAppApi.SearchPost`: %v\n", resp)
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


## SelectPost

> SearchPost200Response SelectPost(ctx).SelectPostRequest(selectPostRequest).Execute()





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
    selectPostRequest := *openapiclient.NewSelectPostRequest(*openapiclient.NewContext(openapiclient.Domain("nic2004:52110"), "TODO", "TODO", "Action_example", "CoreVersion_example", "BapId_example", "BapUri_example", "TransactionId_example", "MessageId_example", time.Now()), *openapiclient.NewSelectPostRequestMessage(*openapiclient.NewOrder())) // SelectPostRequest | Buyer selects one or more catalog items (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ONDCSellerAppApi.SelectPost(context.Background()).SelectPostRequest(selectPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ONDCSellerAppApi.SelectPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SelectPost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `ONDCSellerAppApi.SelectPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSelectPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selectPostRequest** | [**SelectPostRequest**](SelectPostRequest.md) | Buyer selects one or more catalog items | 

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


## StatusPost

> SearchPost200Response StatusPost(ctx).StatusPostRequest(statusPostRequest).Execute()





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
    statusPostRequest := *openapiclient.NewStatusPostRequest(*openapiclient.NewContext(openapiclient.Domain("nic2004:52110"), "TODO", "TODO", "Action_example", "CoreVersion_example", "BapId_example", "BapUri_example", "TransactionId_example", "MessageId_example", time.Now()), *openapiclient.NewStatusPostRequestMessage("TODO")) // StatusPostRequest | Buyer checks for status of order (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ONDCSellerAppApi.StatusPost(context.Background()).StatusPostRequest(statusPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ONDCSellerAppApi.StatusPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StatusPost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `ONDCSellerAppApi.StatusPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStatusPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **statusPostRequest** | [**StatusPostRequest**](StatusPostRequest.md) | Buyer checks for status of order | 

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


## SupportPost

> SearchPost200Response SupportPost(ctx).SupportPostRequest(supportPostRequest).Execute()





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
    supportPostRequest := *openapiclient.NewSupportPostRequest(*openapiclient.NewContext(openapiclient.Domain("nic2004:52110"), "TODO", "TODO", "Action_example", "CoreVersion_example", "BapId_example", "BapUri_example", "TransactionId_example", "MessageId_example", time.Now()), *openapiclient.NewSupportPostRequestMessage()) // SupportPostRequest | Buyer searches for Support Contact details (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ONDCSellerAppApi.SupportPost(context.Background()).SupportPostRequest(supportPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ONDCSellerAppApi.SupportPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SupportPost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `ONDCSellerAppApi.SupportPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSupportPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **supportPostRequest** | [**SupportPostRequest**](SupportPostRequest.md) | Buyer searches for Support Contact details | 

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


## TrackPost

> SearchPost200Response TrackPost(ctx).TrackPostRequest(trackPostRequest).Execute()





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
    trackPostRequest := *openapiclient.NewTrackPostRequest(*openapiclient.NewContext(openapiclient.Domain("nic2004:52110"), "TODO", "TODO", "Action_example", "CoreVersion_example", "BapId_example", "BapUri_example", "TransactionId_example", "MessageId_example", time.Now()), *openapiclient.NewTrackPostRequestMessage("TODO")) // TrackPostRequest | Buyer tracks fulfillment of an order (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ONDCSellerAppApi.TrackPost(context.Background()).TrackPostRequest(trackPostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ONDCSellerAppApi.TrackPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TrackPost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `ONDCSellerAppApi.TrackPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrackPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trackPostRequest** | [**TrackPostRequest**](TrackPostRequest.md) | Buyer tracks fulfillment of an order | 

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


## UpdatePost

> SearchPost200Response UpdatePost(ctx).UpdatePostRequest(updatePostRequest).Execute()





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
    updatePostRequest := *openapiclient.NewUpdatePostRequest(*openapiclient.NewContext(openapiclient.Domain("nic2004:52110"), "TODO", "TODO", "Action_example", "CoreVersion_example", "BapId_example", "BapUri_example", "TransactionId_example", "MessageId_example", time.Now()), *openapiclient.NewUpdatePostRequestMessage("UpdateTarget_example", *openapiclient.NewOrder())) // UpdatePostRequest | Buyer updates an order (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ONDCSellerAppApi.UpdatePost(context.Background()).UpdatePostRequest(updatePostRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ONDCSellerAppApi.UpdatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePost`: SearchPost200Response
    fmt.Fprintf(os.Stdout, "Response from `ONDCSellerAppApi.UpdatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatePostRequest** | [**UpdatePostRequest**](UpdatePostRequest.md) | Buyer updates an order | 

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

