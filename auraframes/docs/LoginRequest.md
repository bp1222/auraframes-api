# LoginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppIdentifier** | **string** |  | 
**ClientDeviceId** | **string** |  | 
**IdentifierForVendor** | **string** |  | 
**Locale** | **string** |  | 
**User** | [**LoginRequestUser**](LoginRequestUser.md) |  | 

## Methods

### NewLoginRequest

`func NewLoginRequest(appIdentifier string, clientDeviceId string, identifierForVendor string, locale string, user LoginRequestUser, ) *LoginRequest`

NewLoginRequest instantiates a new LoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginRequestWithDefaults

`func NewLoginRequestWithDefaults() *LoginRequest`

NewLoginRequestWithDefaults instantiates a new LoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppIdentifier

`func (o *LoginRequest) GetAppIdentifier() string`

GetAppIdentifier returns the AppIdentifier field if non-nil, zero value otherwise.

### GetAppIdentifierOk

`func (o *LoginRequest) GetAppIdentifierOk() (*string, bool)`

GetAppIdentifierOk returns a tuple with the AppIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIdentifier

`func (o *LoginRequest) SetAppIdentifier(v string)`

SetAppIdentifier sets AppIdentifier field to given value.


### GetClientDeviceId

`func (o *LoginRequest) GetClientDeviceId() string`

GetClientDeviceId returns the ClientDeviceId field if non-nil, zero value otherwise.

### GetClientDeviceIdOk

`func (o *LoginRequest) GetClientDeviceIdOk() (*string, bool)`

GetClientDeviceIdOk returns a tuple with the ClientDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientDeviceId

`func (o *LoginRequest) SetClientDeviceId(v string)`

SetClientDeviceId sets ClientDeviceId field to given value.


### GetIdentifierForVendor

`func (o *LoginRequest) GetIdentifierForVendor() string`

GetIdentifierForVendor returns the IdentifierForVendor field if non-nil, zero value otherwise.

### GetIdentifierForVendorOk

`func (o *LoginRequest) GetIdentifierForVendorOk() (*string, bool)`

GetIdentifierForVendorOk returns a tuple with the IdentifierForVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierForVendor

`func (o *LoginRequest) SetIdentifierForVendor(v string)`

SetIdentifierForVendor sets IdentifierForVendor field to given value.


### GetLocale

`func (o *LoginRequest) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *LoginRequest) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *LoginRequest) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetUser

`func (o *LoginRequest) GetUser() LoginRequestUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LoginRequest) GetUserOk() (*LoginRequestUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LoginRequest) SetUser(v LoginRequestUser)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


