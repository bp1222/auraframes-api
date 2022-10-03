/*
Aura Frame API - Unofficial

Reverse Engineered API for Aura Frames

API version: 0.0.1
Contact: dave@mudsite.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package auraframes

import (
	"encoding/json"
)

// LoginResponse struct for LoginResponse
type LoginResponse struct {
	Error *bool `json:"error,omitempty"`
	Result *LoginResponseResult `json:"result,omitempty"`
}

// NewLoginResponse instantiates a new LoginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginResponse() *LoginResponse {
	this := LoginResponse{}
	return &this
}

// NewLoginResponseWithDefaults instantiates a new LoginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginResponseWithDefaults() *LoginResponse {
	this := LoginResponse{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *LoginResponse) GetError() bool {
	if o == nil || o.Error == nil {
		var ret bool
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginResponse) GetErrorOk() (*bool, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *LoginResponse) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given bool and assigns it to the Error field.
func (o *LoginResponse) SetError(v bool) {
	o.Error = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *LoginResponse) GetResult() LoginResponseResult {
	if o == nil || o.Result == nil {
		var ret LoginResponseResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginResponse) GetResultOk() (*LoginResponseResult, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *LoginResponse) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given LoginResponseResult and assigns it to the Result field.
func (o *LoginResponse) SetResult(v LoginResponseResult) {
	o.Result = &v
}

func (o LoginResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableLoginResponse struct {
	value *LoginResponse
	isSet bool
}

func (v NullableLoginResponse) Get() *LoginResponse {
	return v.value
}

func (v *NullableLoginResponse) Set(val *LoginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginResponse(val *LoginResponse) *NullableLoginResponse {
	return &NullableLoginResponse{value: val, isSet: true}
}

func (v NullableLoginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

