# Frame

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**SoftwareVersion** | Pointer to **string** |  | [optional] 
**BuildVersion** | Pointer to **string** |  | [optional] 
**HwAndroidVersion** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**HandledAt** | Pointer to **string** |  | [optional] 
**DeletedAt** | Pointer to **string** |  | [optional] 
**UpdatedAtOnClient** | Pointer to **string** |  | [optional] 
**LastImpressionAt** | Pointer to **string** |  | [optional] 
**Orientation** | Pointer to **string** |  | [optional] 
**AutoBrightness** | Pointer to **bool** |  | [optional] 
**MinBrightness** | Pointer to **string** |  | [optional] 
**MaxBrightness** | Pointer to **string** |  | [optional] 
**Brightness** | Pointer to **string** |  | [optional] 
**SenseMotion** | Pointer to **bool** |  | [optional] 
**DefaultSpeed** | Pointer to **string** |  | [optional] 
**SlideshowInterval** | Pointer to **string** |  | [optional] 
**SlideshowAuto** | Pointer to **bool** |  | [optional] 
**Digits** | Pointer to **string** |  | [optional] 
**ContributorTokens** | Pointer to [**[]ContributorTokens**](ContributorTokens.md) |  | [optional] 
**HwSerial** | Pointer to **string** |  | [optional] 
**MattingColor** | Pointer to **string** |  | [optional] 
**TrimColor** | Pointer to **string** |  | [optional] 
**IsHandling** | Pointer to **bool** |  | [optional] 
**CalibrationsLastModifiedAt** | Pointer to **string** |  | [optional] 
**GesturesOn** | Pointer to **bool** |  | [optional] 
**PortraitPairingOff** | Pointer to **string** |  | [optional] 
**LivePhotosOn** | Pointer to **bool** |  | [optional] 
**AutoProcessedPlaylistIds** | Pointer to **[]string** |  | [optional] 
**TimeZone** | Pointer to **string** |  | [optional] 
**WifiNetwork** | Pointer to **string** |  | [optional] 
**ColdBootAt** | Pointer to **string** |  | [optional] 
**IsCharityWaterFrame** | Pointer to **bool** |  | [optional] 
**NumAssets** | Pointer to **string** |  | [optional] 
**ThanksOn** | Pointer to **bool** |  | [optional] 
**FrameQueueUrl** | Pointer to **string** |  | [optional] 
**ClientQueueUrl** | Pointer to **string** |  | [optional] 
**ScheduledDisplaySleep** | Pointer to **bool** |  | [optional] 
**ScheduledDisplayOnAt** | Pointer to **string** |  | [optional] 
**ScheduledDisplayOffAt** | Pointer to **string** |  | [optional] 
**ForcedWifiState** | Pointer to **string** |  | [optional] 
**ForcedWifiRecipientEmail** | Pointer to **string** |  | [optional] 
**IsAnalogFrame** | Pointer to **bool** |  | [optional] 
**ControlType** | Pointer to **string** |  | [optional] 
**DisplayAspectRatio** | Pointer to **string** |  | [optional] 
**HasClaimableGift** | Pointer to **bool** |  | [optional] 
**GiftBillingHint** | Pointer to **string** |  | [optional] 
**Locale** | Pointer to **string** |  | [optional] 
**FrameType** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**RepresentativeAssetId** | Pointer to **string** |  | [optional] 
**SortMode** | Pointer to **string** |  | [optional] 
**EmailAddress** | Pointer to **string** |  | [optional] 
**Features** | Pointer to **[]string** |  | [optional] 
**Volume** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 
**Playlists** | Pointer to **[]string** |  | [optional] 
**LastFeedItem** | Pointer to [**FeedItem**](FeedItem.md) |  | [optional] 
**LastImpression** | Pointer to [**Impression**](Impression.md) |  | [optional] 
**RecentAssets** | Pointer to [**[]Asset**](Asset.md) |  | [optional] 
**Contributors** | Pointer to [**[]User**](User.md) |  | [optional] 
**FrameEnvironment** | Pointer to [**FrameEnvironment**](FrameEnvironment.md) |  | [optional] 
**CurrentPrintSet** | Pointer to **string** |  | [optional] 
**FirstPrintSet** | Pointer to **string** |  | [optional] 
**ChildAlbums** | Pointer to **[]string** |  | [optional] 
**SmartAdds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewFrame

`func NewFrame() *Frame`

NewFrame instantiates a new Frame object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameWithDefaults

`func NewFrameWithDefaults() *Frame`

NewFrameWithDefaults instantiates a new Frame object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Frame) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Frame) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Frame) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Frame) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Frame) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Frame) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Frame) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Frame) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserId

`func (o *Frame) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Frame) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Frame) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Frame) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetSoftwareVersion

`func (o *Frame) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *Frame) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *Frame) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *Frame) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### GetBuildVersion

`func (o *Frame) GetBuildVersion() string`

GetBuildVersion returns the BuildVersion field if non-nil, zero value otherwise.

### GetBuildVersionOk

`func (o *Frame) GetBuildVersionOk() (*string, bool)`

GetBuildVersionOk returns a tuple with the BuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildVersion

`func (o *Frame) SetBuildVersion(v string)`

SetBuildVersion sets BuildVersion field to given value.

### HasBuildVersion

`func (o *Frame) HasBuildVersion() bool`

HasBuildVersion returns a boolean if a field has been set.

### GetHwAndroidVersion

`func (o *Frame) GetHwAndroidVersion() string`

GetHwAndroidVersion returns the HwAndroidVersion field if non-nil, zero value otherwise.

### GetHwAndroidVersionOk

`func (o *Frame) GetHwAndroidVersionOk() (*string, bool)`

GetHwAndroidVersionOk returns a tuple with the HwAndroidVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwAndroidVersion

`func (o *Frame) SetHwAndroidVersion(v string)`

SetHwAndroidVersion sets HwAndroidVersion field to given value.

### HasHwAndroidVersion

`func (o *Frame) HasHwAndroidVersion() bool`

HasHwAndroidVersion returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Frame) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Frame) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Frame) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Frame) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Frame) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Frame) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Frame) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Frame) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetHandledAt

`func (o *Frame) GetHandledAt() string`

GetHandledAt returns the HandledAt field if non-nil, zero value otherwise.

### GetHandledAtOk

`func (o *Frame) GetHandledAtOk() (*string, bool)`

GetHandledAtOk returns a tuple with the HandledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandledAt

`func (o *Frame) SetHandledAt(v string)`

SetHandledAt sets HandledAt field to given value.

### HasHandledAt

`func (o *Frame) HasHandledAt() bool`

HasHandledAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Frame) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Frame) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Frame) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Frame) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetUpdatedAtOnClient

`func (o *Frame) GetUpdatedAtOnClient() string`

GetUpdatedAtOnClient returns the UpdatedAtOnClient field if non-nil, zero value otherwise.

### GetUpdatedAtOnClientOk

`func (o *Frame) GetUpdatedAtOnClientOk() (*string, bool)`

GetUpdatedAtOnClientOk returns a tuple with the UpdatedAtOnClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtOnClient

`func (o *Frame) SetUpdatedAtOnClient(v string)`

SetUpdatedAtOnClient sets UpdatedAtOnClient field to given value.

### HasUpdatedAtOnClient

`func (o *Frame) HasUpdatedAtOnClient() bool`

HasUpdatedAtOnClient returns a boolean if a field has been set.

### GetLastImpressionAt

`func (o *Frame) GetLastImpressionAt() string`

GetLastImpressionAt returns the LastImpressionAt field if non-nil, zero value otherwise.

### GetLastImpressionAtOk

`func (o *Frame) GetLastImpressionAtOk() (*string, bool)`

GetLastImpressionAtOk returns a tuple with the LastImpressionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastImpressionAt

`func (o *Frame) SetLastImpressionAt(v string)`

SetLastImpressionAt sets LastImpressionAt field to given value.

### HasLastImpressionAt

`func (o *Frame) HasLastImpressionAt() bool`

HasLastImpressionAt returns a boolean if a field has been set.

### GetOrientation

`func (o *Frame) GetOrientation() string`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *Frame) GetOrientationOk() (*string, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *Frame) SetOrientation(v string)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *Frame) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetAutoBrightness

`func (o *Frame) GetAutoBrightness() bool`

GetAutoBrightness returns the AutoBrightness field if non-nil, zero value otherwise.

### GetAutoBrightnessOk

`func (o *Frame) GetAutoBrightnessOk() (*bool, bool)`

GetAutoBrightnessOk returns a tuple with the AutoBrightness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBrightness

`func (o *Frame) SetAutoBrightness(v bool)`

SetAutoBrightness sets AutoBrightness field to given value.

### HasAutoBrightness

`func (o *Frame) HasAutoBrightness() bool`

HasAutoBrightness returns a boolean if a field has been set.

### GetMinBrightness

`func (o *Frame) GetMinBrightness() string`

GetMinBrightness returns the MinBrightness field if non-nil, zero value otherwise.

### GetMinBrightnessOk

`func (o *Frame) GetMinBrightnessOk() (*string, bool)`

GetMinBrightnessOk returns a tuple with the MinBrightness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBrightness

`func (o *Frame) SetMinBrightness(v string)`

SetMinBrightness sets MinBrightness field to given value.

### HasMinBrightness

`func (o *Frame) HasMinBrightness() bool`

HasMinBrightness returns a boolean if a field has been set.

### GetMaxBrightness

`func (o *Frame) GetMaxBrightness() string`

GetMaxBrightness returns the MaxBrightness field if non-nil, zero value otherwise.

### GetMaxBrightnessOk

`func (o *Frame) GetMaxBrightnessOk() (*string, bool)`

GetMaxBrightnessOk returns a tuple with the MaxBrightness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBrightness

`func (o *Frame) SetMaxBrightness(v string)`

SetMaxBrightness sets MaxBrightness field to given value.

### HasMaxBrightness

`func (o *Frame) HasMaxBrightness() bool`

HasMaxBrightness returns a boolean if a field has been set.

### GetBrightness

`func (o *Frame) GetBrightness() string`

GetBrightness returns the Brightness field if non-nil, zero value otherwise.

### GetBrightnessOk

`func (o *Frame) GetBrightnessOk() (*string, bool)`

GetBrightnessOk returns a tuple with the Brightness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrightness

`func (o *Frame) SetBrightness(v string)`

SetBrightness sets Brightness field to given value.

### HasBrightness

`func (o *Frame) HasBrightness() bool`

HasBrightness returns a boolean if a field has been set.

### GetSenseMotion

`func (o *Frame) GetSenseMotion() bool`

GetSenseMotion returns the SenseMotion field if non-nil, zero value otherwise.

### GetSenseMotionOk

`func (o *Frame) GetSenseMotionOk() (*bool, bool)`

GetSenseMotionOk returns a tuple with the SenseMotion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenseMotion

`func (o *Frame) SetSenseMotion(v bool)`

SetSenseMotion sets SenseMotion field to given value.

### HasSenseMotion

`func (o *Frame) HasSenseMotion() bool`

HasSenseMotion returns a boolean if a field has been set.

### GetDefaultSpeed

`func (o *Frame) GetDefaultSpeed() string`

GetDefaultSpeed returns the DefaultSpeed field if non-nil, zero value otherwise.

### GetDefaultSpeedOk

`func (o *Frame) GetDefaultSpeedOk() (*string, bool)`

GetDefaultSpeedOk returns a tuple with the DefaultSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSpeed

`func (o *Frame) SetDefaultSpeed(v string)`

SetDefaultSpeed sets DefaultSpeed field to given value.

### HasDefaultSpeed

`func (o *Frame) HasDefaultSpeed() bool`

HasDefaultSpeed returns a boolean if a field has been set.

### GetSlideshowInterval

`func (o *Frame) GetSlideshowInterval() string`

GetSlideshowInterval returns the SlideshowInterval field if non-nil, zero value otherwise.

### GetSlideshowIntervalOk

`func (o *Frame) GetSlideshowIntervalOk() (*string, bool)`

GetSlideshowIntervalOk returns a tuple with the SlideshowInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlideshowInterval

`func (o *Frame) SetSlideshowInterval(v string)`

SetSlideshowInterval sets SlideshowInterval field to given value.

### HasSlideshowInterval

`func (o *Frame) HasSlideshowInterval() bool`

HasSlideshowInterval returns a boolean if a field has been set.

### GetSlideshowAuto

`func (o *Frame) GetSlideshowAuto() bool`

GetSlideshowAuto returns the SlideshowAuto field if non-nil, zero value otherwise.

### GetSlideshowAutoOk

`func (o *Frame) GetSlideshowAutoOk() (*bool, bool)`

GetSlideshowAutoOk returns a tuple with the SlideshowAuto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlideshowAuto

`func (o *Frame) SetSlideshowAuto(v bool)`

SetSlideshowAuto sets SlideshowAuto field to given value.

### HasSlideshowAuto

`func (o *Frame) HasSlideshowAuto() bool`

HasSlideshowAuto returns a boolean if a field has been set.

### GetDigits

`func (o *Frame) GetDigits() string`

GetDigits returns the Digits field if non-nil, zero value otherwise.

### GetDigitsOk

`func (o *Frame) GetDigitsOk() (*string, bool)`

GetDigitsOk returns a tuple with the Digits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigits

`func (o *Frame) SetDigits(v string)`

SetDigits sets Digits field to given value.

### HasDigits

`func (o *Frame) HasDigits() bool`

HasDigits returns a boolean if a field has been set.

### GetContributorTokens

`func (o *Frame) GetContributorTokens() []ContributorTokens`

GetContributorTokens returns the ContributorTokens field if non-nil, zero value otherwise.

### GetContributorTokensOk

`func (o *Frame) GetContributorTokensOk() (*[]ContributorTokens, bool)`

GetContributorTokensOk returns a tuple with the ContributorTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributorTokens

`func (o *Frame) SetContributorTokens(v []ContributorTokens)`

SetContributorTokens sets ContributorTokens field to given value.

### HasContributorTokens

`func (o *Frame) HasContributorTokens() bool`

HasContributorTokens returns a boolean if a field has been set.

### GetHwSerial

`func (o *Frame) GetHwSerial() string`

GetHwSerial returns the HwSerial field if non-nil, zero value otherwise.

### GetHwSerialOk

`func (o *Frame) GetHwSerialOk() (*string, bool)`

GetHwSerialOk returns a tuple with the HwSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwSerial

`func (o *Frame) SetHwSerial(v string)`

SetHwSerial sets HwSerial field to given value.

### HasHwSerial

`func (o *Frame) HasHwSerial() bool`

HasHwSerial returns a boolean if a field has been set.

### GetMattingColor

`func (o *Frame) GetMattingColor() string`

GetMattingColor returns the MattingColor field if non-nil, zero value otherwise.

### GetMattingColorOk

`func (o *Frame) GetMattingColorOk() (*string, bool)`

GetMattingColorOk returns a tuple with the MattingColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMattingColor

`func (o *Frame) SetMattingColor(v string)`

SetMattingColor sets MattingColor field to given value.

### HasMattingColor

`func (o *Frame) HasMattingColor() bool`

HasMattingColor returns a boolean if a field has been set.

### GetTrimColor

`func (o *Frame) GetTrimColor() string`

GetTrimColor returns the TrimColor field if non-nil, zero value otherwise.

### GetTrimColorOk

`func (o *Frame) GetTrimColorOk() (*string, bool)`

GetTrimColorOk returns a tuple with the TrimColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrimColor

`func (o *Frame) SetTrimColor(v string)`

SetTrimColor sets TrimColor field to given value.

### HasTrimColor

`func (o *Frame) HasTrimColor() bool`

HasTrimColor returns a boolean if a field has been set.

### GetIsHandling

`func (o *Frame) GetIsHandling() bool`

GetIsHandling returns the IsHandling field if non-nil, zero value otherwise.

### GetIsHandlingOk

`func (o *Frame) GetIsHandlingOk() (*bool, bool)`

GetIsHandlingOk returns a tuple with the IsHandling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHandling

`func (o *Frame) SetIsHandling(v bool)`

SetIsHandling sets IsHandling field to given value.

### HasIsHandling

`func (o *Frame) HasIsHandling() bool`

HasIsHandling returns a boolean if a field has been set.

### GetCalibrationsLastModifiedAt

`func (o *Frame) GetCalibrationsLastModifiedAt() string`

GetCalibrationsLastModifiedAt returns the CalibrationsLastModifiedAt field if non-nil, zero value otherwise.

### GetCalibrationsLastModifiedAtOk

`func (o *Frame) GetCalibrationsLastModifiedAtOk() (*string, bool)`

GetCalibrationsLastModifiedAtOk returns a tuple with the CalibrationsLastModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalibrationsLastModifiedAt

`func (o *Frame) SetCalibrationsLastModifiedAt(v string)`

SetCalibrationsLastModifiedAt sets CalibrationsLastModifiedAt field to given value.

### HasCalibrationsLastModifiedAt

`func (o *Frame) HasCalibrationsLastModifiedAt() bool`

HasCalibrationsLastModifiedAt returns a boolean if a field has been set.

### GetGesturesOn

`func (o *Frame) GetGesturesOn() bool`

GetGesturesOn returns the GesturesOn field if non-nil, zero value otherwise.

### GetGesturesOnOk

`func (o *Frame) GetGesturesOnOk() (*bool, bool)`

GetGesturesOnOk returns a tuple with the GesturesOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGesturesOn

`func (o *Frame) SetGesturesOn(v bool)`

SetGesturesOn sets GesturesOn field to given value.

### HasGesturesOn

`func (o *Frame) HasGesturesOn() bool`

HasGesturesOn returns a boolean if a field has been set.

### GetPortraitPairingOff

`func (o *Frame) GetPortraitPairingOff() string`

GetPortraitPairingOff returns the PortraitPairingOff field if non-nil, zero value otherwise.

### GetPortraitPairingOffOk

`func (o *Frame) GetPortraitPairingOffOk() (*string, bool)`

GetPortraitPairingOffOk returns a tuple with the PortraitPairingOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortraitPairingOff

`func (o *Frame) SetPortraitPairingOff(v string)`

SetPortraitPairingOff sets PortraitPairingOff field to given value.

### HasPortraitPairingOff

`func (o *Frame) HasPortraitPairingOff() bool`

HasPortraitPairingOff returns a boolean if a field has been set.

### GetLivePhotosOn

`func (o *Frame) GetLivePhotosOn() bool`

GetLivePhotosOn returns the LivePhotosOn field if non-nil, zero value otherwise.

### GetLivePhotosOnOk

`func (o *Frame) GetLivePhotosOnOk() (*bool, bool)`

GetLivePhotosOnOk returns a tuple with the LivePhotosOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivePhotosOn

`func (o *Frame) SetLivePhotosOn(v bool)`

SetLivePhotosOn sets LivePhotosOn field to given value.

### HasLivePhotosOn

`func (o *Frame) HasLivePhotosOn() bool`

HasLivePhotosOn returns a boolean if a field has been set.

### GetAutoProcessedPlaylistIds

`func (o *Frame) GetAutoProcessedPlaylistIds() []string`

GetAutoProcessedPlaylistIds returns the AutoProcessedPlaylistIds field if non-nil, zero value otherwise.

### GetAutoProcessedPlaylistIdsOk

`func (o *Frame) GetAutoProcessedPlaylistIdsOk() (*[]string, bool)`

GetAutoProcessedPlaylistIdsOk returns a tuple with the AutoProcessedPlaylistIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoProcessedPlaylistIds

`func (o *Frame) SetAutoProcessedPlaylistIds(v []string)`

SetAutoProcessedPlaylistIds sets AutoProcessedPlaylistIds field to given value.

### HasAutoProcessedPlaylistIds

`func (o *Frame) HasAutoProcessedPlaylistIds() bool`

HasAutoProcessedPlaylistIds returns a boolean if a field has been set.

### GetTimeZone

`func (o *Frame) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *Frame) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *Frame) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *Frame) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetWifiNetwork

`func (o *Frame) GetWifiNetwork() string`

GetWifiNetwork returns the WifiNetwork field if non-nil, zero value otherwise.

### GetWifiNetworkOk

`func (o *Frame) GetWifiNetworkOk() (*string, bool)`

GetWifiNetworkOk returns a tuple with the WifiNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiNetwork

`func (o *Frame) SetWifiNetwork(v string)`

SetWifiNetwork sets WifiNetwork field to given value.

### HasWifiNetwork

`func (o *Frame) HasWifiNetwork() bool`

HasWifiNetwork returns a boolean if a field has been set.

### GetColdBootAt

`func (o *Frame) GetColdBootAt() string`

GetColdBootAt returns the ColdBootAt field if non-nil, zero value otherwise.

### GetColdBootAtOk

`func (o *Frame) GetColdBootAtOk() (*string, bool)`

GetColdBootAtOk returns a tuple with the ColdBootAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColdBootAt

`func (o *Frame) SetColdBootAt(v string)`

SetColdBootAt sets ColdBootAt field to given value.

### HasColdBootAt

`func (o *Frame) HasColdBootAt() bool`

HasColdBootAt returns a boolean if a field has been set.

### GetIsCharityWaterFrame

`func (o *Frame) GetIsCharityWaterFrame() bool`

GetIsCharityWaterFrame returns the IsCharityWaterFrame field if non-nil, zero value otherwise.

### GetIsCharityWaterFrameOk

`func (o *Frame) GetIsCharityWaterFrameOk() (*bool, bool)`

GetIsCharityWaterFrameOk returns a tuple with the IsCharityWaterFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCharityWaterFrame

`func (o *Frame) SetIsCharityWaterFrame(v bool)`

SetIsCharityWaterFrame sets IsCharityWaterFrame field to given value.

### HasIsCharityWaterFrame

`func (o *Frame) HasIsCharityWaterFrame() bool`

HasIsCharityWaterFrame returns a boolean if a field has been set.

### GetNumAssets

`func (o *Frame) GetNumAssets() string`

GetNumAssets returns the NumAssets field if non-nil, zero value otherwise.

### GetNumAssetsOk

`func (o *Frame) GetNumAssetsOk() (*string, bool)`

GetNumAssetsOk returns a tuple with the NumAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAssets

`func (o *Frame) SetNumAssets(v string)`

SetNumAssets sets NumAssets field to given value.

### HasNumAssets

`func (o *Frame) HasNumAssets() bool`

HasNumAssets returns a boolean if a field has been set.

### GetThanksOn

`func (o *Frame) GetThanksOn() bool`

GetThanksOn returns the ThanksOn field if non-nil, zero value otherwise.

### GetThanksOnOk

`func (o *Frame) GetThanksOnOk() (*bool, bool)`

GetThanksOnOk returns a tuple with the ThanksOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThanksOn

`func (o *Frame) SetThanksOn(v bool)`

SetThanksOn sets ThanksOn field to given value.

### HasThanksOn

`func (o *Frame) HasThanksOn() bool`

HasThanksOn returns a boolean if a field has been set.

### GetFrameQueueUrl

`func (o *Frame) GetFrameQueueUrl() string`

GetFrameQueueUrl returns the FrameQueueUrl field if non-nil, zero value otherwise.

### GetFrameQueueUrlOk

`func (o *Frame) GetFrameQueueUrlOk() (*string, bool)`

GetFrameQueueUrlOk returns a tuple with the FrameQueueUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameQueueUrl

`func (o *Frame) SetFrameQueueUrl(v string)`

SetFrameQueueUrl sets FrameQueueUrl field to given value.

### HasFrameQueueUrl

`func (o *Frame) HasFrameQueueUrl() bool`

HasFrameQueueUrl returns a boolean if a field has been set.

### GetClientQueueUrl

`func (o *Frame) GetClientQueueUrl() string`

GetClientQueueUrl returns the ClientQueueUrl field if non-nil, zero value otherwise.

### GetClientQueueUrlOk

`func (o *Frame) GetClientQueueUrlOk() (*string, bool)`

GetClientQueueUrlOk returns a tuple with the ClientQueueUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientQueueUrl

`func (o *Frame) SetClientQueueUrl(v string)`

SetClientQueueUrl sets ClientQueueUrl field to given value.

### HasClientQueueUrl

`func (o *Frame) HasClientQueueUrl() bool`

HasClientQueueUrl returns a boolean if a field has been set.

### GetScheduledDisplaySleep

`func (o *Frame) GetScheduledDisplaySleep() bool`

GetScheduledDisplaySleep returns the ScheduledDisplaySleep field if non-nil, zero value otherwise.

### GetScheduledDisplaySleepOk

`func (o *Frame) GetScheduledDisplaySleepOk() (*bool, bool)`

GetScheduledDisplaySleepOk returns a tuple with the ScheduledDisplaySleep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDisplaySleep

`func (o *Frame) SetScheduledDisplaySleep(v bool)`

SetScheduledDisplaySleep sets ScheduledDisplaySleep field to given value.

### HasScheduledDisplaySleep

`func (o *Frame) HasScheduledDisplaySleep() bool`

HasScheduledDisplaySleep returns a boolean if a field has been set.

### GetScheduledDisplayOnAt

`func (o *Frame) GetScheduledDisplayOnAt() string`

GetScheduledDisplayOnAt returns the ScheduledDisplayOnAt field if non-nil, zero value otherwise.

### GetScheduledDisplayOnAtOk

`func (o *Frame) GetScheduledDisplayOnAtOk() (*string, bool)`

GetScheduledDisplayOnAtOk returns a tuple with the ScheduledDisplayOnAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDisplayOnAt

`func (o *Frame) SetScheduledDisplayOnAt(v string)`

SetScheduledDisplayOnAt sets ScheduledDisplayOnAt field to given value.

### HasScheduledDisplayOnAt

`func (o *Frame) HasScheduledDisplayOnAt() bool`

HasScheduledDisplayOnAt returns a boolean if a field has been set.

### GetScheduledDisplayOffAt

`func (o *Frame) GetScheduledDisplayOffAt() string`

GetScheduledDisplayOffAt returns the ScheduledDisplayOffAt field if non-nil, zero value otherwise.

### GetScheduledDisplayOffAtOk

`func (o *Frame) GetScheduledDisplayOffAtOk() (*string, bool)`

GetScheduledDisplayOffAtOk returns a tuple with the ScheduledDisplayOffAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDisplayOffAt

`func (o *Frame) SetScheduledDisplayOffAt(v string)`

SetScheduledDisplayOffAt sets ScheduledDisplayOffAt field to given value.

### HasScheduledDisplayOffAt

`func (o *Frame) HasScheduledDisplayOffAt() bool`

HasScheduledDisplayOffAt returns a boolean if a field has been set.

### GetForcedWifiState

`func (o *Frame) GetForcedWifiState() string`

GetForcedWifiState returns the ForcedWifiState field if non-nil, zero value otherwise.

### GetForcedWifiStateOk

`func (o *Frame) GetForcedWifiStateOk() (*string, bool)`

GetForcedWifiStateOk returns a tuple with the ForcedWifiState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedWifiState

`func (o *Frame) SetForcedWifiState(v string)`

SetForcedWifiState sets ForcedWifiState field to given value.

### HasForcedWifiState

`func (o *Frame) HasForcedWifiState() bool`

HasForcedWifiState returns a boolean if a field has been set.

### GetForcedWifiRecipientEmail

`func (o *Frame) GetForcedWifiRecipientEmail() string`

GetForcedWifiRecipientEmail returns the ForcedWifiRecipientEmail field if non-nil, zero value otherwise.

### GetForcedWifiRecipientEmailOk

`func (o *Frame) GetForcedWifiRecipientEmailOk() (*string, bool)`

GetForcedWifiRecipientEmailOk returns a tuple with the ForcedWifiRecipientEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedWifiRecipientEmail

`func (o *Frame) SetForcedWifiRecipientEmail(v string)`

SetForcedWifiRecipientEmail sets ForcedWifiRecipientEmail field to given value.

### HasForcedWifiRecipientEmail

`func (o *Frame) HasForcedWifiRecipientEmail() bool`

HasForcedWifiRecipientEmail returns a boolean if a field has been set.

### GetIsAnalogFrame

`func (o *Frame) GetIsAnalogFrame() bool`

GetIsAnalogFrame returns the IsAnalogFrame field if non-nil, zero value otherwise.

### GetIsAnalogFrameOk

`func (o *Frame) GetIsAnalogFrameOk() (*bool, bool)`

GetIsAnalogFrameOk returns a tuple with the IsAnalogFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnalogFrame

`func (o *Frame) SetIsAnalogFrame(v bool)`

SetIsAnalogFrame sets IsAnalogFrame field to given value.

### HasIsAnalogFrame

`func (o *Frame) HasIsAnalogFrame() bool`

HasIsAnalogFrame returns a boolean if a field has been set.

### GetControlType

`func (o *Frame) GetControlType() string`

GetControlType returns the ControlType field if non-nil, zero value otherwise.

### GetControlTypeOk

`func (o *Frame) GetControlTypeOk() (*string, bool)`

GetControlTypeOk returns a tuple with the ControlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlType

`func (o *Frame) SetControlType(v string)`

SetControlType sets ControlType field to given value.

### HasControlType

`func (o *Frame) HasControlType() bool`

HasControlType returns a boolean if a field has been set.

### GetDisplayAspectRatio

`func (o *Frame) GetDisplayAspectRatio() string`

GetDisplayAspectRatio returns the DisplayAspectRatio field if non-nil, zero value otherwise.

### GetDisplayAspectRatioOk

`func (o *Frame) GetDisplayAspectRatioOk() (*string, bool)`

GetDisplayAspectRatioOk returns a tuple with the DisplayAspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayAspectRatio

`func (o *Frame) SetDisplayAspectRatio(v string)`

SetDisplayAspectRatio sets DisplayAspectRatio field to given value.

### HasDisplayAspectRatio

`func (o *Frame) HasDisplayAspectRatio() bool`

HasDisplayAspectRatio returns a boolean if a field has been set.

### GetHasClaimableGift

`func (o *Frame) GetHasClaimableGift() bool`

GetHasClaimableGift returns the HasClaimableGift field if non-nil, zero value otherwise.

### GetHasClaimableGiftOk

`func (o *Frame) GetHasClaimableGiftOk() (*bool, bool)`

GetHasClaimableGiftOk returns a tuple with the HasClaimableGift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasClaimableGift

`func (o *Frame) SetHasClaimableGift(v bool)`

SetHasClaimableGift sets HasClaimableGift field to given value.

### HasHasClaimableGift

`func (o *Frame) HasHasClaimableGift() bool`

HasHasClaimableGift returns a boolean if a field has been set.

### GetGiftBillingHint

`func (o *Frame) GetGiftBillingHint() string`

GetGiftBillingHint returns the GiftBillingHint field if non-nil, zero value otherwise.

### GetGiftBillingHintOk

`func (o *Frame) GetGiftBillingHintOk() (*string, bool)`

GetGiftBillingHintOk returns a tuple with the GiftBillingHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftBillingHint

`func (o *Frame) SetGiftBillingHint(v string)`

SetGiftBillingHint sets GiftBillingHint field to given value.

### HasGiftBillingHint

`func (o *Frame) HasGiftBillingHint() bool`

HasGiftBillingHint returns a boolean if a field has been set.

### GetLocale

`func (o *Frame) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *Frame) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *Frame) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *Frame) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetFrameType

`func (o *Frame) GetFrameType() string`

GetFrameType returns the FrameType field if non-nil, zero value otherwise.

### GetFrameTypeOk

`func (o *Frame) GetFrameTypeOk() (*string, bool)`

GetFrameTypeOk returns a tuple with the FrameType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameType

`func (o *Frame) SetFrameType(v string)`

SetFrameType sets FrameType field to given value.

### HasFrameType

`func (o *Frame) HasFrameType() bool`

HasFrameType returns a boolean if a field has been set.

### GetDescription

`func (o *Frame) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Frame) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Frame) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Frame) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRepresentativeAssetId

`func (o *Frame) GetRepresentativeAssetId() string`

GetRepresentativeAssetId returns the RepresentativeAssetId field if non-nil, zero value otherwise.

### GetRepresentativeAssetIdOk

`func (o *Frame) GetRepresentativeAssetIdOk() (*string, bool)`

GetRepresentativeAssetIdOk returns a tuple with the RepresentativeAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepresentativeAssetId

`func (o *Frame) SetRepresentativeAssetId(v string)`

SetRepresentativeAssetId sets RepresentativeAssetId field to given value.

### HasRepresentativeAssetId

`func (o *Frame) HasRepresentativeAssetId() bool`

HasRepresentativeAssetId returns a boolean if a field has been set.

### GetSortMode

`func (o *Frame) GetSortMode() string`

GetSortMode returns the SortMode field if non-nil, zero value otherwise.

### GetSortModeOk

`func (o *Frame) GetSortModeOk() (*string, bool)`

GetSortModeOk returns a tuple with the SortMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortMode

`func (o *Frame) SetSortMode(v string)`

SetSortMode sets SortMode field to given value.

### HasSortMode

`func (o *Frame) HasSortMode() bool`

HasSortMode returns a boolean if a field has been set.

### GetEmailAddress

`func (o *Frame) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *Frame) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *Frame) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *Frame) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetFeatures

`func (o *Frame) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *Frame) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *Frame) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *Frame) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetVolume

`func (o *Frame) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *Frame) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *Frame) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *Frame) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetUser

`func (o *Frame) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Frame) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Frame) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *Frame) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetPlaylists

`func (o *Frame) GetPlaylists() []string`

GetPlaylists returns the Playlists field if non-nil, zero value otherwise.

### GetPlaylistsOk

`func (o *Frame) GetPlaylistsOk() (*[]string, bool)`

GetPlaylistsOk returns a tuple with the Playlists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylists

`func (o *Frame) SetPlaylists(v []string)`

SetPlaylists sets Playlists field to given value.

### HasPlaylists

`func (o *Frame) HasPlaylists() bool`

HasPlaylists returns a boolean if a field has been set.

### GetLastFeedItem

`func (o *Frame) GetLastFeedItem() FeedItem`

GetLastFeedItem returns the LastFeedItem field if non-nil, zero value otherwise.

### GetLastFeedItemOk

`func (o *Frame) GetLastFeedItemOk() (*FeedItem, bool)`

GetLastFeedItemOk returns a tuple with the LastFeedItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFeedItem

`func (o *Frame) SetLastFeedItem(v FeedItem)`

SetLastFeedItem sets LastFeedItem field to given value.

### HasLastFeedItem

`func (o *Frame) HasLastFeedItem() bool`

HasLastFeedItem returns a boolean if a field has been set.

### GetLastImpression

`func (o *Frame) GetLastImpression() Impression`

GetLastImpression returns the LastImpression field if non-nil, zero value otherwise.

### GetLastImpressionOk

`func (o *Frame) GetLastImpressionOk() (*Impression, bool)`

GetLastImpressionOk returns a tuple with the LastImpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastImpression

`func (o *Frame) SetLastImpression(v Impression)`

SetLastImpression sets LastImpression field to given value.

### HasLastImpression

`func (o *Frame) HasLastImpression() bool`

HasLastImpression returns a boolean if a field has been set.

### GetRecentAssets

`func (o *Frame) GetRecentAssets() []Asset`

GetRecentAssets returns the RecentAssets field if non-nil, zero value otherwise.

### GetRecentAssetsOk

`func (o *Frame) GetRecentAssetsOk() (*[]Asset, bool)`

GetRecentAssetsOk returns a tuple with the RecentAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentAssets

`func (o *Frame) SetRecentAssets(v []Asset)`

SetRecentAssets sets RecentAssets field to given value.

### HasRecentAssets

`func (o *Frame) HasRecentAssets() bool`

HasRecentAssets returns a boolean if a field has been set.

### GetContributors

`func (o *Frame) GetContributors() []User`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *Frame) GetContributorsOk() (*[]User, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *Frame) SetContributors(v []User)`

SetContributors sets Contributors field to given value.

### HasContributors

`func (o *Frame) HasContributors() bool`

HasContributors returns a boolean if a field has been set.

### GetFrameEnvironment

`func (o *Frame) GetFrameEnvironment() FrameEnvironment`

GetFrameEnvironment returns the FrameEnvironment field if non-nil, zero value otherwise.

### GetFrameEnvironmentOk

`func (o *Frame) GetFrameEnvironmentOk() (*FrameEnvironment, bool)`

GetFrameEnvironmentOk returns a tuple with the FrameEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameEnvironment

`func (o *Frame) SetFrameEnvironment(v FrameEnvironment)`

SetFrameEnvironment sets FrameEnvironment field to given value.

### HasFrameEnvironment

`func (o *Frame) HasFrameEnvironment() bool`

HasFrameEnvironment returns a boolean if a field has been set.

### GetCurrentPrintSet

`func (o *Frame) GetCurrentPrintSet() string`

GetCurrentPrintSet returns the CurrentPrintSet field if non-nil, zero value otherwise.

### GetCurrentPrintSetOk

`func (o *Frame) GetCurrentPrintSetOk() (*string, bool)`

GetCurrentPrintSetOk returns a tuple with the CurrentPrintSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPrintSet

`func (o *Frame) SetCurrentPrintSet(v string)`

SetCurrentPrintSet sets CurrentPrintSet field to given value.

### HasCurrentPrintSet

`func (o *Frame) HasCurrentPrintSet() bool`

HasCurrentPrintSet returns a boolean if a field has been set.

### GetFirstPrintSet

`func (o *Frame) GetFirstPrintSet() string`

GetFirstPrintSet returns the FirstPrintSet field if non-nil, zero value otherwise.

### GetFirstPrintSetOk

`func (o *Frame) GetFirstPrintSetOk() (*string, bool)`

GetFirstPrintSetOk returns a tuple with the FirstPrintSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPrintSet

`func (o *Frame) SetFirstPrintSet(v string)`

SetFirstPrintSet sets FirstPrintSet field to given value.

### HasFirstPrintSet

`func (o *Frame) HasFirstPrintSet() bool`

HasFirstPrintSet returns a boolean if a field has been set.

### GetChildAlbums

`func (o *Frame) GetChildAlbums() []string`

GetChildAlbums returns the ChildAlbums field if non-nil, zero value otherwise.

### GetChildAlbumsOk

`func (o *Frame) GetChildAlbumsOk() (*[]string, bool)`

GetChildAlbumsOk returns a tuple with the ChildAlbums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildAlbums

`func (o *Frame) SetChildAlbums(v []string)`

SetChildAlbums sets ChildAlbums field to given value.

### HasChildAlbums

`func (o *Frame) HasChildAlbums() bool`

HasChildAlbums returns a boolean if a field has been set.

### GetSmartAdds

`func (o *Frame) GetSmartAdds() []string`

GetSmartAdds returns the SmartAdds field if non-nil, zero value otherwise.

### GetSmartAddsOk

`func (o *Frame) GetSmartAddsOk() (*[]string, bool)`

GetSmartAddsOk returns a tuple with the SmartAdds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartAdds

`func (o *Frame) SetSmartAdds(v []string)`

SetSmartAdds sets SmartAdds field to given value.

### HasSmartAdds

`func (o *Frame) HasSmartAdds() bool`

HasSmartAdds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


