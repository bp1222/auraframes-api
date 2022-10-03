# \AuthApi

All URIs are relative to *https://api.pushd.com/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Login**](AuthApi.md#Login) | **Post** /login.json | Login to Aura Frames



## Login

> LoginResponse Login(ctx).XDeviceIdentifier(xDeviceIdentifier).XClientDeviceId(xClientDeviceId).LoginRequest(loginRequest).Execute()

Login to Aura Frames

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
    xDeviceIdentifier := "xDeviceIdentifier_example" // string |  (optional)
    xClientDeviceId := "xClientDeviceId_example" // string |  (optional)
    loginRequest := *openapiclient.NewLoginRequest("AppIdentifier_example", "ClientDeviceId_example", "IdentifierForVendor_example", "Locale_example", *openapiclient.NewLoginRequestUser("Email_example", "Password_example")) // LoginRequest | Login Information (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.Login(context.Background()).XDeviceIdentifier(xDeviceIdentifier).XClientDeviceId(xClientDeviceId).LoginRequest(loginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.Login``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Login`: LoginResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.Login`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xDeviceIdentifier** | **string** |  | 
 **xClientDeviceId** | **string** |  | 
 **loginRequest** | [**LoginRequest**](LoginRequest.md) | Login Information | 

### Return type

[**LoginResponse**](LoginResponse.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [UserId](../README.md#UserId)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

