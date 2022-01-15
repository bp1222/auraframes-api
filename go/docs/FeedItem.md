# FeedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**[]Asset**](Asset.md) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**StickFor** | Pointer to **string** |  | [optional] 

## Methods

### NewFeedItem

`func NewFeedItem() *FeedItem`

NewFeedItem instantiates a new FeedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedItemWithDefaults

`func NewFeedItemWithDefaults() *FeedItem`

NewFeedItemWithDefaults instantiates a new FeedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *FeedItem) GetAssets() []Asset`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *FeedItem) GetAssetsOk() (*[]Asset, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *FeedItem) SetAssets(v []Asset)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *FeedItem) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetMetadata

`func (o *FeedItem) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FeedItem) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FeedItem) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FeedItem) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMessage

`func (o *FeedItem) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FeedItem) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FeedItem) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FeedItem) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStickFor

`func (o *FeedItem) GetStickFor() string`

GetStickFor returns the StickFor field if non-nil, zero value otherwise.

### GetStickForOk

`func (o *FeedItem) GetStickForOk() (*string, bool)`

GetStickForOk returns a tuple with the StickFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickFor

`func (o *FeedItem) SetStickFor(v string)`

SetStickFor sets StickFor field to given value.

### HasStickFor

`func (o *FeedItem) HasStickFor() bool`

HasStickFor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


