# Asset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**ThumbnailUrl** | Pointer to **string** |  | [optional] 
**PortraitUrl** | Pointer to **string** |  | [optional] 
**LandscapeUrl** | Pointer to **string** |  | [optional] 
**Landscape1610Url** | Pointer to **string** |  | [optional] 
**Portrait45Url** | Pointer to **string** |  | [optional] 
**VideoUrl** | Pointer to **string** |  | [optional] 
**LandscapeRect** | Pointer to **string** |  | [optional] 
**PortraitRect** | Pointer to **string** |  | [optional] 
**UserLandscapeRect** | Pointer to **string** |  | [optional] 
**UserPortraitRect** | Pointer to **string** |  | [optional] 
**AutoLandscape1610Rect** | Pointer to **string** |  | [optional] 
**UserLandscape1610Rect** | Pointer to **string** |  | [optional] 
**AutoPortrait45Rect** | Pointer to **string** |  | [optional] 
**UserPortrait45Rect** | Pointer to **string** |  | [optional] 
**ExifOrientation** | Pointer to **string** |  | [optional] 
**HandledAt** | Pointer to **string** |  | [optional] 
**UploadedAt** | Pointer to **string** |  | [optional] 
**GoodResolution** | Pointer to **bool** |  | [optional] 
**SourceId** | Pointer to **string** |  | [optional] 
**DuplicateOfId** | Pointer to **string** |  | [optional] 
**RotationCw** | Pointer to **string** |  | [optional] 
**LocationName** | Pointer to **string** |  | [optional] 
**Md5Hash** | Pointer to **string** |  | [optional] 
**IsSubscription** | Pointer to **bool** |  | [optional] 
**GlacieredAt** | Pointer to **string** |  | [optional] 
**Unglacierable** | Pointer to **bool** |  | [optional] 
**Duration** | Pointer to **string** |  | [optional] 
**LivePhotoOff** | Pointer to **string** |  | [optional] 
**LocalIdentifier** | Pointer to **string** |  | [optional] 
**CreatedAtOnClient** | Pointer to **string** |  | [optional] 
**Selected** | Pointer to **bool** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**RawFileName** | Pointer to **string** |  | [optional] 
**VideoFileName** | Pointer to **string** |  | [optional] 
**Width** | Pointer to **string** |  | [optional] 
**Height** | Pointer to **string** |  | [optional] 
**TakenAt** | Pointer to **string** |  | [optional] 
**HorizontalAccuracy** | Pointer to **string** |  | [optional] 
**Favorite** | Pointer to **bool** |  | [optional] 
**Orientation** | Pointer to **string** |  | [optional] 
**Hdr** | Pointer to **bool** |  | [optional] 
**Panorama** | Pointer to **bool** |  | [optional] 
**IsLive** | Pointer to **string** |  | [optional] 
**BurstId** | Pointer to **string** |  | [optional] 
**BurstSelectionTypes** | Pointer to **string** |  | [optional] 
**RepresentsBurst** | Pointer to **string** |  | [optional] 
**DataUti** | Pointer to **string** |  | [optional] 
**OriginalFileName** | Pointer to **string** |  | [optional] 
**UploadPriority** | Pointer to **string** |  | [optional] 
**IosMediaSubtypes** | Pointer to **string** |  | [optional] 
**TakenAtUserOverrideAt** | Pointer to **string** |  | [optional] 
**DurationUnclipped** | Pointer to **string** |  | [optional] 
**VideoClipStart** | Pointer to **string** |  | [optional] 
**VideoClipExcludesAudio** | Pointer to **string** |  | [optional] 
**VideoClippedByUserAt** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 

## Methods

### NewAsset

`func NewAsset() *Asset`

NewAsset instantiates a new Asset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWithDefaults

`func NewAssetWithDefaults() *Asset`

NewAssetWithDefaults instantiates a new Asset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Asset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Asset) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Asset) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Asset) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *Asset) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Asset) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Asset) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Asset) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *Asset) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *Asset) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *Asset) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *Asset) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetPortraitUrl

`func (o *Asset) GetPortraitUrl() string`

GetPortraitUrl returns the PortraitUrl field if non-nil, zero value otherwise.

### GetPortraitUrlOk

`func (o *Asset) GetPortraitUrlOk() (*string, bool)`

GetPortraitUrlOk returns a tuple with the PortraitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortraitUrl

`func (o *Asset) SetPortraitUrl(v string)`

SetPortraitUrl sets PortraitUrl field to given value.

### HasPortraitUrl

`func (o *Asset) HasPortraitUrl() bool`

HasPortraitUrl returns a boolean if a field has been set.

### GetLandscapeUrl

`func (o *Asset) GetLandscapeUrl() string`

GetLandscapeUrl returns the LandscapeUrl field if non-nil, zero value otherwise.

### GetLandscapeUrlOk

`func (o *Asset) GetLandscapeUrlOk() (*string, bool)`

GetLandscapeUrlOk returns a tuple with the LandscapeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandscapeUrl

`func (o *Asset) SetLandscapeUrl(v string)`

SetLandscapeUrl sets LandscapeUrl field to given value.

### HasLandscapeUrl

`func (o *Asset) HasLandscapeUrl() bool`

HasLandscapeUrl returns a boolean if a field has been set.

### GetLandscape1610Url

`func (o *Asset) GetLandscape1610Url() string`

GetLandscape1610Url returns the Landscape1610Url field if non-nil, zero value otherwise.

### GetLandscape1610UrlOk

`func (o *Asset) GetLandscape1610UrlOk() (*string, bool)`

GetLandscape1610UrlOk returns a tuple with the Landscape1610Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandscape1610Url

`func (o *Asset) SetLandscape1610Url(v string)`

SetLandscape1610Url sets Landscape1610Url field to given value.

### HasLandscape1610Url

`func (o *Asset) HasLandscape1610Url() bool`

HasLandscape1610Url returns a boolean if a field has been set.

### GetPortrait45Url

`func (o *Asset) GetPortrait45Url() string`

GetPortrait45Url returns the Portrait45Url field if non-nil, zero value otherwise.

### GetPortrait45UrlOk

`func (o *Asset) GetPortrait45UrlOk() (*string, bool)`

GetPortrait45UrlOk returns a tuple with the Portrait45Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortrait45Url

`func (o *Asset) SetPortrait45Url(v string)`

SetPortrait45Url sets Portrait45Url field to given value.

### HasPortrait45Url

`func (o *Asset) HasPortrait45Url() bool`

HasPortrait45Url returns a boolean if a field has been set.

### GetVideoUrl

`func (o *Asset) GetVideoUrl() string`

GetVideoUrl returns the VideoUrl field if non-nil, zero value otherwise.

### GetVideoUrlOk

`func (o *Asset) GetVideoUrlOk() (*string, bool)`

GetVideoUrlOk returns a tuple with the VideoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoUrl

`func (o *Asset) SetVideoUrl(v string)`

SetVideoUrl sets VideoUrl field to given value.

### HasVideoUrl

`func (o *Asset) HasVideoUrl() bool`

HasVideoUrl returns a boolean if a field has been set.

### GetLandscapeRect

`func (o *Asset) GetLandscapeRect() string`

GetLandscapeRect returns the LandscapeRect field if non-nil, zero value otherwise.

### GetLandscapeRectOk

`func (o *Asset) GetLandscapeRectOk() (*string, bool)`

GetLandscapeRectOk returns a tuple with the LandscapeRect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandscapeRect

`func (o *Asset) SetLandscapeRect(v string)`

SetLandscapeRect sets LandscapeRect field to given value.

### HasLandscapeRect

`func (o *Asset) HasLandscapeRect() bool`

HasLandscapeRect returns a boolean if a field has been set.

### GetPortraitRect

`func (o *Asset) GetPortraitRect() string`

GetPortraitRect returns the PortraitRect field if non-nil, zero value otherwise.

### GetPortraitRectOk

`func (o *Asset) GetPortraitRectOk() (*string, bool)`

GetPortraitRectOk returns a tuple with the PortraitRect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortraitRect

`func (o *Asset) SetPortraitRect(v string)`

SetPortraitRect sets PortraitRect field to given value.

### HasPortraitRect

`func (o *Asset) HasPortraitRect() bool`

HasPortraitRect returns a boolean if a field has been set.

### GetUserLandscapeRect

`func (o *Asset) GetUserLandscapeRect() string`

GetUserLandscapeRect returns the UserLandscapeRect field if non-nil, zero value otherwise.

### GetUserLandscapeRectOk

`func (o *Asset) GetUserLandscapeRectOk() (*string, bool)`

GetUserLandscapeRectOk returns a tuple with the UserLandscapeRect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLandscapeRect

`func (o *Asset) SetUserLandscapeRect(v string)`

SetUserLandscapeRect sets UserLandscapeRect field to given value.

### HasUserLandscapeRect

`func (o *Asset) HasUserLandscapeRect() bool`

HasUserLandscapeRect returns a boolean if a field has been set.

### GetUserPortraitRect

`func (o *Asset) GetUserPortraitRect() string`

GetUserPortraitRect returns the UserPortraitRect field if non-nil, zero value otherwise.

### GetUserPortraitRectOk

`func (o *Asset) GetUserPortraitRectOk() (*string, bool)`

GetUserPortraitRectOk returns a tuple with the UserPortraitRect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPortraitRect

`func (o *Asset) SetUserPortraitRect(v string)`

SetUserPortraitRect sets UserPortraitRect field to given value.

### HasUserPortraitRect

`func (o *Asset) HasUserPortraitRect() bool`

HasUserPortraitRect returns a boolean if a field has been set.

### GetAutoLandscape1610Rect

`func (o *Asset) GetAutoLandscape1610Rect() string`

GetAutoLandscape1610Rect returns the AutoLandscape1610Rect field if non-nil, zero value otherwise.

### GetAutoLandscape1610RectOk

`func (o *Asset) GetAutoLandscape1610RectOk() (*string, bool)`

GetAutoLandscape1610RectOk returns a tuple with the AutoLandscape1610Rect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoLandscape1610Rect

`func (o *Asset) SetAutoLandscape1610Rect(v string)`

SetAutoLandscape1610Rect sets AutoLandscape1610Rect field to given value.

### HasAutoLandscape1610Rect

`func (o *Asset) HasAutoLandscape1610Rect() bool`

HasAutoLandscape1610Rect returns a boolean if a field has been set.

### GetUserLandscape1610Rect

`func (o *Asset) GetUserLandscape1610Rect() string`

GetUserLandscape1610Rect returns the UserLandscape1610Rect field if non-nil, zero value otherwise.

### GetUserLandscape1610RectOk

`func (o *Asset) GetUserLandscape1610RectOk() (*string, bool)`

GetUserLandscape1610RectOk returns a tuple with the UserLandscape1610Rect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLandscape1610Rect

`func (o *Asset) SetUserLandscape1610Rect(v string)`

SetUserLandscape1610Rect sets UserLandscape1610Rect field to given value.

### HasUserLandscape1610Rect

`func (o *Asset) HasUserLandscape1610Rect() bool`

HasUserLandscape1610Rect returns a boolean if a field has been set.

### GetAutoPortrait45Rect

`func (o *Asset) GetAutoPortrait45Rect() string`

GetAutoPortrait45Rect returns the AutoPortrait45Rect field if non-nil, zero value otherwise.

### GetAutoPortrait45RectOk

`func (o *Asset) GetAutoPortrait45RectOk() (*string, bool)`

GetAutoPortrait45RectOk returns a tuple with the AutoPortrait45Rect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPortrait45Rect

`func (o *Asset) SetAutoPortrait45Rect(v string)`

SetAutoPortrait45Rect sets AutoPortrait45Rect field to given value.

### HasAutoPortrait45Rect

`func (o *Asset) HasAutoPortrait45Rect() bool`

HasAutoPortrait45Rect returns a boolean if a field has been set.

### GetUserPortrait45Rect

`func (o *Asset) GetUserPortrait45Rect() string`

GetUserPortrait45Rect returns the UserPortrait45Rect field if non-nil, zero value otherwise.

### GetUserPortrait45RectOk

`func (o *Asset) GetUserPortrait45RectOk() (*string, bool)`

GetUserPortrait45RectOk returns a tuple with the UserPortrait45Rect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPortrait45Rect

`func (o *Asset) SetUserPortrait45Rect(v string)`

SetUserPortrait45Rect sets UserPortrait45Rect field to given value.

### HasUserPortrait45Rect

`func (o *Asset) HasUserPortrait45Rect() bool`

HasUserPortrait45Rect returns a boolean if a field has been set.

### GetExifOrientation

`func (o *Asset) GetExifOrientation() string`

GetExifOrientation returns the ExifOrientation field if non-nil, zero value otherwise.

### GetExifOrientationOk

`func (o *Asset) GetExifOrientationOk() (*string, bool)`

GetExifOrientationOk returns a tuple with the ExifOrientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExifOrientation

`func (o *Asset) SetExifOrientation(v string)`

SetExifOrientation sets ExifOrientation field to given value.

### HasExifOrientation

`func (o *Asset) HasExifOrientation() bool`

HasExifOrientation returns a boolean if a field has been set.

### GetHandledAt

`func (o *Asset) GetHandledAt() string`

GetHandledAt returns the HandledAt field if non-nil, zero value otherwise.

### GetHandledAtOk

`func (o *Asset) GetHandledAtOk() (*string, bool)`

GetHandledAtOk returns a tuple with the HandledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandledAt

`func (o *Asset) SetHandledAt(v string)`

SetHandledAt sets HandledAt field to given value.

### HasHandledAt

`func (o *Asset) HasHandledAt() bool`

HasHandledAt returns a boolean if a field has been set.

### GetUploadedAt

`func (o *Asset) GetUploadedAt() string`

GetUploadedAt returns the UploadedAt field if non-nil, zero value otherwise.

### GetUploadedAtOk

`func (o *Asset) GetUploadedAtOk() (*string, bool)`

GetUploadedAtOk returns a tuple with the UploadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedAt

`func (o *Asset) SetUploadedAt(v string)`

SetUploadedAt sets UploadedAt field to given value.

### HasUploadedAt

`func (o *Asset) HasUploadedAt() bool`

HasUploadedAt returns a boolean if a field has been set.

### GetGoodResolution

`func (o *Asset) GetGoodResolution() bool`

GetGoodResolution returns the GoodResolution field if non-nil, zero value otherwise.

### GetGoodResolutionOk

`func (o *Asset) GetGoodResolutionOk() (*bool, bool)`

GetGoodResolutionOk returns a tuple with the GoodResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodResolution

`func (o *Asset) SetGoodResolution(v bool)`

SetGoodResolution sets GoodResolution field to given value.

### HasGoodResolution

`func (o *Asset) HasGoodResolution() bool`

HasGoodResolution returns a boolean if a field has been set.

### GetSourceId

`func (o *Asset) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *Asset) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *Asset) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *Asset) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetDuplicateOfId

`func (o *Asset) GetDuplicateOfId() string`

GetDuplicateOfId returns the DuplicateOfId field if non-nil, zero value otherwise.

### GetDuplicateOfIdOk

`func (o *Asset) GetDuplicateOfIdOk() (*string, bool)`

GetDuplicateOfIdOk returns a tuple with the DuplicateOfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateOfId

`func (o *Asset) SetDuplicateOfId(v string)`

SetDuplicateOfId sets DuplicateOfId field to given value.

### HasDuplicateOfId

`func (o *Asset) HasDuplicateOfId() bool`

HasDuplicateOfId returns a boolean if a field has been set.

### GetRotationCw

`func (o *Asset) GetRotationCw() string`

GetRotationCw returns the RotationCw field if non-nil, zero value otherwise.

### GetRotationCwOk

`func (o *Asset) GetRotationCwOk() (*string, bool)`

GetRotationCwOk returns a tuple with the RotationCw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationCw

`func (o *Asset) SetRotationCw(v string)`

SetRotationCw sets RotationCw field to given value.

### HasRotationCw

`func (o *Asset) HasRotationCw() bool`

HasRotationCw returns a boolean if a field has been set.

### GetLocationName

`func (o *Asset) GetLocationName() string`

GetLocationName returns the LocationName field if non-nil, zero value otherwise.

### GetLocationNameOk

`func (o *Asset) GetLocationNameOk() (*string, bool)`

GetLocationNameOk returns a tuple with the LocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationName

`func (o *Asset) SetLocationName(v string)`

SetLocationName sets LocationName field to given value.

### HasLocationName

`func (o *Asset) HasLocationName() bool`

HasLocationName returns a boolean if a field has been set.

### GetMd5Hash

`func (o *Asset) GetMd5Hash() string`

GetMd5Hash returns the Md5Hash field if non-nil, zero value otherwise.

### GetMd5HashOk

`func (o *Asset) GetMd5HashOk() (*string, bool)`

GetMd5HashOk returns a tuple with the Md5Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5Hash

`func (o *Asset) SetMd5Hash(v string)`

SetMd5Hash sets Md5Hash field to given value.

### HasMd5Hash

`func (o *Asset) HasMd5Hash() bool`

HasMd5Hash returns a boolean if a field has been set.

### GetIsSubscription

`func (o *Asset) GetIsSubscription() bool`

GetIsSubscription returns the IsSubscription field if non-nil, zero value otherwise.

### GetIsSubscriptionOk

`func (o *Asset) GetIsSubscriptionOk() (*bool, bool)`

GetIsSubscriptionOk returns a tuple with the IsSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscription

`func (o *Asset) SetIsSubscription(v bool)`

SetIsSubscription sets IsSubscription field to given value.

### HasIsSubscription

`func (o *Asset) HasIsSubscription() bool`

HasIsSubscription returns a boolean if a field has been set.

### GetGlacieredAt

`func (o *Asset) GetGlacieredAt() string`

GetGlacieredAt returns the GlacieredAt field if non-nil, zero value otherwise.

### GetGlacieredAtOk

`func (o *Asset) GetGlacieredAtOk() (*string, bool)`

GetGlacieredAtOk returns a tuple with the GlacieredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlacieredAt

`func (o *Asset) SetGlacieredAt(v string)`

SetGlacieredAt sets GlacieredAt field to given value.

### HasGlacieredAt

`func (o *Asset) HasGlacieredAt() bool`

HasGlacieredAt returns a boolean if a field has been set.

### GetUnglacierable

`func (o *Asset) GetUnglacierable() bool`

GetUnglacierable returns the Unglacierable field if non-nil, zero value otherwise.

### GetUnglacierableOk

`func (o *Asset) GetUnglacierableOk() (*bool, bool)`

GetUnglacierableOk returns a tuple with the Unglacierable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnglacierable

`func (o *Asset) SetUnglacierable(v bool)`

SetUnglacierable sets Unglacierable field to given value.

### HasUnglacierable

`func (o *Asset) HasUnglacierable() bool`

HasUnglacierable returns a boolean if a field has been set.

### GetDuration

`func (o *Asset) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Asset) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Asset) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Asset) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetLivePhotoOff

`func (o *Asset) GetLivePhotoOff() string`

GetLivePhotoOff returns the LivePhotoOff field if non-nil, zero value otherwise.

### GetLivePhotoOffOk

`func (o *Asset) GetLivePhotoOffOk() (*string, bool)`

GetLivePhotoOffOk returns a tuple with the LivePhotoOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivePhotoOff

`func (o *Asset) SetLivePhotoOff(v string)`

SetLivePhotoOff sets LivePhotoOff field to given value.

### HasLivePhotoOff

`func (o *Asset) HasLivePhotoOff() bool`

HasLivePhotoOff returns a boolean if a field has been set.

### GetLocalIdentifier

`func (o *Asset) GetLocalIdentifier() string`

GetLocalIdentifier returns the LocalIdentifier field if non-nil, zero value otherwise.

### GetLocalIdentifierOk

`func (o *Asset) GetLocalIdentifierOk() (*string, bool)`

GetLocalIdentifierOk returns a tuple with the LocalIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIdentifier

`func (o *Asset) SetLocalIdentifier(v string)`

SetLocalIdentifier sets LocalIdentifier field to given value.

### HasLocalIdentifier

`func (o *Asset) HasLocalIdentifier() bool`

HasLocalIdentifier returns a boolean if a field has been set.

### GetCreatedAtOnClient

`func (o *Asset) GetCreatedAtOnClient() string`

GetCreatedAtOnClient returns the CreatedAtOnClient field if non-nil, zero value otherwise.

### GetCreatedAtOnClientOk

`func (o *Asset) GetCreatedAtOnClientOk() (*string, bool)`

GetCreatedAtOnClientOk returns a tuple with the CreatedAtOnClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtOnClient

`func (o *Asset) SetCreatedAtOnClient(v string)`

SetCreatedAtOnClient sets CreatedAtOnClient field to given value.

### HasCreatedAtOnClient

`func (o *Asset) HasCreatedAtOnClient() bool`

HasCreatedAtOnClient returns a boolean if a field has been set.

### GetSelected

`func (o *Asset) GetSelected() bool`

GetSelected returns the Selected field if non-nil, zero value otherwise.

### GetSelectedOk

`func (o *Asset) GetSelectedOk() (*bool, bool)`

GetSelectedOk returns a tuple with the Selected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelected

`func (o *Asset) SetSelected(v bool)`

SetSelected sets Selected field to given value.

### HasSelected

`func (o *Asset) HasSelected() bool`

HasSelected returns a boolean if a field has been set.

### GetFileName

`func (o *Asset) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *Asset) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *Asset) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *Asset) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetRawFileName

`func (o *Asset) GetRawFileName() string`

GetRawFileName returns the RawFileName field if non-nil, zero value otherwise.

### GetRawFileNameOk

`func (o *Asset) GetRawFileNameOk() (*string, bool)`

GetRawFileNameOk returns a tuple with the RawFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawFileName

`func (o *Asset) SetRawFileName(v string)`

SetRawFileName sets RawFileName field to given value.

### HasRawFileName

`func (o *Asset) HasRawFileName() bool`

HasRawFileName returns a boolean if a field has been set.

### GetVideoFileName

`func (o *Asset) GetVideoFileName() string`

GetVideoFileName returns the VideoFileName field if non-nil, zero value otherwise.

### GetVideoFileNameOk

`func (o *Asset) GetVideoFileNameOk() (*string, bool)`

GetVideoFileNameOk returns a tuple with the VideoFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoFileName

`func (o *Asset) SetVideoFileName(v string)`

SetVideoFileName sets VideoFileName field to given value.

### HasVideoFileName

`func (o *Asset) HasVideoFileName() bool`

HasVideoFileName returns a boolean if a field has been set.

### GetWidth

`func (o *Asset) GetWidth() string`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *Asset) GetWidthOk() (*string, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *Asset) SetWidth(v string)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *Asset) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *Asset) GetHeight() string`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *Asset) GetHeightOk() (*string, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *Asset) SetHeight(v string)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *Asset) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetTakenAt

`func (o *Asset) GetTakenAt() string`

GetTakenAt returns the TakenAt field if non-nil, zero value otherwise.

### GetTakenAtOk

`func (o *Asset) GetTakenAtOk() (*string, bool)`

GetTakenAtOk returns a tuple with the TakenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakenAt

`func (o *Asset) SetTakenAt(v string)`

SetTakenAt sets TakenAt field to given value.

### HasTakenAt

`func (o *Asset) HasTakenAt() bool`

HasTakenAt returns a boolean if a field has been set.

### GetHorizontalAccuracy

`func (o *Asset) GetHorizontalAccuracy() string`

GetHorizontalAccuracy returns the HorizontalAccuracy field if non-nil, zero value otherwise.

### GetHorizontalAccuracyOk

`func (o *Asset) GetHorizontalAccuracyOk() (*string, bool)`

GetHorizontalAccuracyOk returns a tuple with the HorizontalAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizontalAccuracy

`func (o *Asset) SetHorizontalAccuracy(v string)`

SetHorizontalAccuracy sets HorizontalAccuracy field to given value.

### HasHorizontalAccuracy

`func (o *Asset) HasHorizontalAccuracy() bool`

HasHorizontalAccuracy returns a boolean if a field has been set.

### GetFavorite

`func (o *Asset) GetFavorite() bool`

GetFavorite returns the Favorite field if non-nil, zero value otherwise.

### GetFavoriteOk

`func (o *Asset) GetFavoriteOk() (*bool, bool)`

GetFavoriteOk returns a tuple with the Favorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorite

`func (o *Asset) SetFavorite(v bool)`

SetFavorite sets Favorite field to given value.

### HasFavorite

`func (o *Asset) HasFavorite() bool`

HasFavorite returns a boolean if a field has been set.

### GetOrientation

`func (o *Asset) GetOrientation() string`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *Asset) GetOrientationOk() (*string, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *Asset) SetOrientation(v string)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *Asset) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetHdr

`func (o *Asset) GetHdr() bool`

GetHdr returns the Hdr field if non-nil, zero value otherwise.

### GetHdrOk

`func (o *Asset) GetHdrOk() (*bool, bool)`

GetHdrOk returns a tuple with the Hdr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdr

`func (o *Asset) SetHdr(v bool)`

SetHdr sets Hdr field to given value.

### HasHdr

`func (o *Asset) HasHdr() bool`

HasHdr returns a boolean if a field has been set.

### GetPanorama

`func (o *Asset) GetPanorama() bool`

GetPanorama returns the Panorama field if non-nil, zero value otherwise.

### GetPanoramaOk

`func (o *Asset) GetPanoramaOk() (*bool, bool)`

GetPanoramaOk returns a tuple with the Panorama field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanorama

`func (o *Asset) SetPanorama(v bool)`

SetPanorama sets Panorama field to given value.

### HasPanorama

`func (o *Asset) HasPanorama() bool`

HasPanorama returns a boolean if a field has been set.

### GetIsLive

`func (o *Asset) GetIsLive() string`

GetIsLive returns the IsLive field if non-nil, zero value otherwise.

### GetIsLiveOk

`func (o *Asset) GetIsLiveOk() (*string, bool)`

GetIsLiveOk returns a tuple with the IsLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLive

`func (o *Asset) SetIsLive(v string)`

SetIsLive sets IsLive field to given value.

### HasIsLive

`func (o *Asset) HasIsLive() bool`

HasIsLive returns a boolean if a field has been set.

### GetBurstId

`func (o *Asset) GetBurstId() string`

GetBurstId returns the BurstId field if non-nil, zero value otherwise.

### GetBurstIdOk

`func (o *Asset) GetBurstIdOk() (*string, bool)`

GetBurstIdOk returns a tuple with the BurstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstId

`func (o *Asset) SetBurstId(v string)`

SetBurstId sets BurstId field to given value.

### HasBurstId

`func (o *Asset) HasBurstId() bool`

HasBurstId returns a boolean if a field has been set.

### GetBurstSelectionTypes

`func (o *Asset) GetBurstSelectionTypes() string`

GetBurstSelectionTypes returns the BurstSelectionTypes field if non-nil, zero value otherwise.

### GetBurstSelectionTypesOk

`func (o *Asset) GetBurstSelectionTypesOk() (*string, bool)`

GetBurstSelectionTypesOk returns a tuple with the BurstSelectionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstSelectionTypes

`func (o *Asset) SetBurstSelectionTypes(v string)`

SetBurstSelectionTypes sets BurstSelectionTypes field to given value.

### HasBurstSelectionTypes

`func (o *Asset) HasBurstSelectionTypes() bool`

HasBurstSelectionTypes returns a boolean if a field has been set.

### GetRepresentsBurst

`func (o *Asset) GetRepresentsBurst() string`

GetRepresentsBurst returns the RepresentsBurst field if non-nil, zero value otherwise.

### GetRepresentsBurstOk

`func (o *Asset) GetRepresentsBurstOk() (*string, bool)`

GetRepresentsBurstOk returns a tuple with the RepresentsBurst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepresentsBurst

`func (o *Asset) SetRepresentsBurst(v string)`

SetRepresentsBurst sets RepresentsBurst field to given value.

### HasRepresentsBurst

`func (o *Asset) HasRepresentsBurst() bool`

HasRepresentsBurst returns a boolean if a field has been set.

### GetDataUti

`func (o *Asset) GetDataUti() string`

GetDataUti returns the DataUti field if non-nil, zero value otherwise.

### GetDataUtiOk

`func (o *Asset) GetDataUtiOk() (*string, bool)`

GetDataUtiOk returns a tuple with the DataUti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataUti

`func (o *Asset) SetDataUti(v string)`

SetDataUti sets DataUti field to given value.

### HasDataUti

`func (o *Asset) HasDataUti() bool`

HasDataUti returns a boolean if a field has been set.

### GetOriginalFileName

`func (o *Asset) GetOriginalFileName() string`

GetOriginalFileName returns the OriginalFileName field if non-nil, zero value otherwise.

### GetOriginalFileNameOk

`func (o *Asset) GetOriginalFileNameOk() (*string, bool)`

GetOriginalFileNameOk returns a tuple with the OriginalFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFileName

`func (o *Asset) SetOriginalFileName(v string)`

SetOriginalFileName sets OriginalFileName field to given value.

### HasOriginalFileName

`func (o *Asset) HasOriginalFileName() bool`

HasOriginalFileName returns a boolean if a field has been set.

### GetUploadPriority

`func (o *Asset) GetUploadPriority() string`

GetUploadPriority returns the UploadPriority field if non-nil, zero value otherwise.

### GetUploadPriorityOk

`func (o *Asset) GetUploadPriorityOk() (*string, bool)`

GetUploadPriorityOk returns a tuple with the UploadPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadPriority

`func (o *Asset) SetUploadPriority(v string)`

SetUploadPriority sets UploadPriority field to given value.

### HasUploadPriority

`func (o *Asset) HasUploadPriority() bool`

HasUploadPriority returns a boolean if a field has been set.

### GetIosMediaSubtypes

`func (o *Asset) GetIosMediaSubtypes() string`

GetIosMediaSubtypes returns the IosMediaSubtypes field if non-nil, zero value otherwise.

### GetIosMediaSubtypesOk

`func (o *Asset) GetIosMediaSubtypesOk() (*string, bool)`

GetIosMediaSubtypesOk returns a tuple with the IosMediaSubtypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosMediaSubtypes

`func (o *Asset) SetIosMediaSubtypes(v string)`

SetIosMediaSubtypes sets IosMediaSubtypes field to given value.

### HasIosMediaSubtypes

`func (o *Asset) HasIosMediaSubtypes() bool`

HasIosMediaSubtypes returns a boolean if a field has been set.

### GetTakenAtUserOverrideAt

`func (o *Asset) GetTakenAtUserOverrideAt() string`

GetTakenAtUserOverrideAt returns the TakenAtUserOverrideAt field if non-nil, zero value otherwise.

### GetTakenAtUserOverrideAtOk

`func (o *Asset) GetTakenAtUserOverrideAtOk() (*string, bool)`

GetTakenAtUserOverrideAtOk returns a tuple with the TakenAtUserOverrideAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakenAtUserOverrideAt

`func (o *Asset) SetTakenAtUserOverrideAt(v string)`

SetTakenAtUserOverrideAt sets TakenAtUserOverrideAt field to given value.

### HasTakenAtUserOverrideAt

`func (o *Asset) HasTakenAtUserOverrideAt() bool`

HasTakenAtUserOverrideAt returns a boolean if a field has been set.

### GetDurationUnclipped

`func (o *Asset) GetDurationUnclipped() string`

GetDurationUnclipped returns the DurationUnclipped field if non-nil, zero value otherwise.

### GetDurationUnclippedOk

`func (o *Asset) GetDurationUnclippedOk() (*string, bool)`

GetDurationUnclippedOk returns a tuple with the DurationUnclipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationUnclipped

`func (o *Asset) SetDurationUnclipped(v string)`

SetDurationUnclipped sets DurationUnclipped field to given value.

### HasDurationUnclipped

`func (o *Asset) HasDurationUnclipped() bool`

HasDurationUnclipped returns a boolean if a field has been set.

### GetVideoClipStart

`func (o *Asset) GetVideoClipStart() string`

GetVideoClipStart returns the VideoClipStart field if non-nil, zero value otherwise.

### GetVideoClipStartOk

`func (o *Asset) GetVideoClipStartOk() (*string, bool)`

GetVideoClipStartOk returns a tuple with the VideoClipStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoClipStart

`func (o *Asset) SetVideoClipStart(v string)`

SetVideoClipStart sets VideoClipStart field to given value.

### HasVideoClipStart

`func (o *Asset) HasVideoClipStart() bool`

HasVideoClipStart returns a boolean if a field has been set.

### GetVideoClipExcludesAudio

`func (o *Asset) GetVideoClipExcludesAudio() string`

GetVideoClipExcludesAudio returns the VideoClipExcludesAudio field if non-nil, zero value otherwise.

### GetVideoClipExcludesAudioOk

`func (o *Asset) GetVideoClipExcludesAudioOk() (*string, bool)`

GetVideoClipExcludesAudioOk returns a tuple with the VideoClipExcludesAudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoClipExcludesAudio

`func (o *Asset) SetVideoClipExcludesAudio(v string)`

SetVideoClipExcludesAudio sets VideoClipExcludesAudio field to given value.

### HasVideoClipExcludesAudio

`func (o *Asset) HasVideoClipExcludesAudio() bool`

HasVideoClipExcludesAudio returns a boolean if a field has been set.

### GetVideoClippedByUserAt

`func (o *Asset) GetVideoClippedByUserAt() string`

GetVideoClippedByUserAt returns the VideoClippedByUserAt field if non-nil, zero value otherwise.

### GetVideoClippedByUserAtOk

`func (o *Asset) GetVideoClippedByUserAtOk() (*string, bool)`

GetVideoClippedByUserAtOk returns a tuple with the VideoClippedByUserAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoClippedByUserAt

`func (o *Asset) SetVideoClippedByUserAt(v string)`

SetVideoClippedByUserAt sets VideoClippedByUserAt field to given value.

### HasVideoClippedByUserAt

`func (o *Asset) HasVideoClippedByUserAt() bool`

HasVideoClippedByUserAt returns a boolean if a field has been set.

### GetLocation

`func (o *Asset) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Asset) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Asset) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Asset) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetUser

`func (o *Asset) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Asset) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Asset) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *Asset) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


