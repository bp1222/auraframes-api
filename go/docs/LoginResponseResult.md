# LoginResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentUser** | Pointer to [**CurrentUser**](CurrentUser.md) |  | [optional] 

## Methods

### NewLoginResponseResult

`func NewLoginResponseResult() *LoginResponseResult`

NewLoginResponseResult instantiates a new LoginResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginResponseResultWithDefaults

`func NewLoginResponseResultWithDefaults() *LoginResponseResult`

NewLoginResponseResultWithDefaults instantiates a new LoginResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentUser

`func (o *LoginResponseResult) GetCurrentUser() CurrentUser`

GetCurrentUser returns the CurrentUser field if non-nil, zero value otherwise.

### GetCurrentUserOk

`func (o *LoginResponseResult) GetCurrentUserOk() (*CurrentUser, bool)`

GetCurrentUserOk returns a tuple with the CurrentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUser

`func (o *LoginResponseResult) SetCurrentUser(v CurrentUser)`

SetCurrentUser sets CurrentUser field to given value.

### HasCurrentUser

`func (o *LoginResponseResult) HasCurrentUser() bool`

HasCurrentUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


