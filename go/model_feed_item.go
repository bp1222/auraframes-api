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

// FeedItem struct for FeedItem
type FeedItem struct {
	Assets *[]Asset `json:"assets,omitempty"`
	Metadata *Metadata `json:"metadata,omitempty"`
	Message *string `json:"message,omitempty"`
	StickFor *string `json:"stick_for,omitempty"`
}

// NewFeedItem instantiates a new FeedItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedItem() *FeedItem {
	this := FeedItem{}
	return &this
}

// NewFeedItemWithDefaults instantiates a new FeedItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedItemWithDefaults() *FeedItem {
	this := FeedItem{}
	return &this
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *FeedItem) GetAssets() []Asset {
	if o == nil || o.Assets == nil {
		var ret []Asset
		return ret
	}
	return *o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedItem) GetAssetsOk() (*[]Asset, bool) {
	if o == nil || o.Assets == nil {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *FeedItem) HasAssets() bool {
	if o != nil && o.Assets != nil {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []Asset and assigns it to the Assets field.
func (o *FeedItem) SetAssets(v []Asset) {
	o.Assets = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *FeedItem) GetMetadata() Metadata {
	if o == nil || o.Metadata == nil {
		var ret Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedItem) GetMetadataOk() (*Metadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *FeedItem) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Metadata and assigns it to the Metadata field.
func (o *FeedItem) SetMetadata(v Metadata) {
	o.Metadata = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *FeedItem) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedItem) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *FeedItem) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *FeedItem) SetMessage(v string) {
	o.Message = &v
}

// GetStickFor returns the StickFor field value if set, zero value otherwise.
func (o *FeedItem) GetStickFor() string {
	if o == nil || o.StickFor == nil {
		var ret string
		return ret
	}
	return *o.StickFor
}

// GetStickForOk returns a tuple with the StickFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedItem) GetStickForOk() (*string, bool) {
	if o == nil || o.StickFor == nil {
		return nil, false
	}
	return o.StickFor, true
}

// HasStickFor returns a boolean if a field has been set.
func (o *FeedItem) HasStickFor() bool {
	if o != nil && o.StickFor != nil {
		return true
	}

	return false
}

// SetStickFor gets a reference to the given string and assigns it to the StickFor field.
func (o *FeedItem) SetStickFor(v string) {
	o.StickFor = &v
}

func (o FeedItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Assets != nil {
		toSerialize["assets"] = o.Assets
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.StickFor != nil {
		toSerialize["stick_for"] = o.StickFor
	}
	return json.Marshal(toSerialize)
}

type NullableFeedItem struct {
	value *FeedItem
	isSet bool
}

func (v NullableFeedItem) Get() *FeedItem {
	return v.value
}

func (v *NullableFeedItem) Set(val *FeedItem) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedItem) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedItem(val *FeedItem) *NullableFeedItem {
	return &NullableFeedItem{value: val, isSet: true}
}

func (v NullableFeedItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

