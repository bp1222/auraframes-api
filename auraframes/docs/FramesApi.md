# \FramesApi

All URIs are relative to *https://api.pushd.com/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFrames**](FramesApi.md#GetFrames) | **Get** /frames.json | Access to an Aura Frame



## GetFrames

> Frames GetFrames(ctx).IncludeSharedAlbums(includeSharedAlbums).Execute()

Access to an Aura Frame

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
    includeSharedAlbums := "includeSharedAlbums_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FramesApi.GetFrames(context.Background()).IncludeSharedAlbums(includeSharedAlbums).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FramesApi.GetFrames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFrames`: Frames
    fmt.Fprintf(os.Stdout, "Response from `FramesApi.GetFrames`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFramesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeSharedAlbums** | **string** |  | 

### Return type

[**Frames**](Frames.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [UserId](../README.md#UserId)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

