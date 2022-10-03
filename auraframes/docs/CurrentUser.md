# CurrentUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminAccount** | Pointer to **string** |  | [optional] 
**AttributionId** | Pointer to **string** |  | [optional] 
**AttributionString** | Pointer to **string** |  | [optional] 
**AuthToken** | Pointer to **string** |  | [optional] 
**AutoUploadOff** | Pointer to **bool** |  | [optional] 
**AvatarFileName** | Pointer to **string** |  | [optional] 
**BillingInfo** | Pointer to **string** |  | [optional] 
**CharitySubscriptionsLaunched** | Pointer to **bool** |  | [optional] 
**ConfirmedEmail** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CurrentSourceId** | Pointer to **string** |  | [optional] 
**EligibleForAppReviewPrompt** | Pointer to **bool** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Features** | Pointer to **[]string** |  | [optional] 
**GooglePhotosDisabled** | Pointer to **string** |  | [optional] 
**HasAccessToNewGooglePhotos** | Pointer to **bool** |  | [optional] 
**HasFrame** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LatestAppVersion** | Pointer to **string** |  | [optional] 
**LivePhotosLaunched** | Pointer to **bool** |  | [optional] 
**Locale** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PrintSubscriptionsLaunched** | Pointer to **bool** |  | [optional] 
**RecurlyAccountCode** | Pointer to **string** |  | [optional] 
**ShortId** | Pointer to **string** |  | [optional] 
**ShowPushPrompt** | Pointer to **bool** |  | [optional] 
**SmartAlbumsOff** | Pointer to **bool** |  | [optional] 
**SmartSuggestionsOff** | Pointer to **bool** |  | [optional] 
**SubscriptionsLaunched** | Pointer to **bool** |  | [optional] 
**TestAccount** | Pointer to **string** |  | [optional] 
**ThanksLaunched** | Pointer to **bool** |  | [optional] 
**TooltipAddPhotosSeen** | Pointer to **bool** |  | [optional] 
**TooltipAddedPhotosSeen** | Pointer to **bool** |  | [optional] 
**TooltipGesturesSeen** | Pointer to **bool** |  | [optional] 
**TooltipInboxSeen** | Pointer to **bool** |  | [optional] 
**TooltipManageFramesSeen** | Pointer to **bool** |  | [optional] 
**TooltipSettingsSeen** | Pointer to **bool** |  | [optional] 
**UnconfirmedEmail** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**VerboseLoggingEnabled** | Pointer to **bool** |  | [optional] 
**WarnSmartAlbumsDeprecated** | Pointer to **bool** |  | [optional] 

## Methods

### NewCurrentUser

`func NewCurrentUser() *CurrentUser`

NewCurrentUser instantiates a new CurrentUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentUserWithDefaults

`func NewCurrentUserWithDefaults() *CurrentUser`

NewCurrentUserWithDefaults instantiates a new CurrentUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminAccount

`func (o *CurrentUser) GetAdminAccount() string`

GetAdminAccount returns the AdminAccount field if non-nil, zero value otherwise.

### GetAdminAccountOk

`func (o *CurrentUser) GetAdminAccountOk() (*string, bool)`

GetAdminAccountOk returns a tuple with the AdminAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminAccount

`func (o *CurrentUser) SetAdminAccount(v string)`

SetAdminAccount sets AdminAccount field to given value.

### HasAdminAccount

`func (o *CurrentUser) HasAdminAccount() bool`

HasAdminAccount returns a boolean if a field has been set.

### GetAttributionId

`func (o *CurrentUser) GetAttributionId() string`

GetAttributionId returns the AttributionId field if non-nil, zero value otherwise.

### GetAttributionIdOk

`func (o *CurrentUser) GetAttributionIdOk() (*string, bool)`

GetAttributionIdOk returns a tuple with the AttributionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributionId

`func (o *CurrentUser) SetAttributionId(v string)`

SetAttributionId sets AttributionId field to given value.

### HasAttributionId

`func (o *CurrentUser) HasAttributionId() bool`

HasAttributionId returns a boolean if a field has been set.

### GetAttributionString

`func (o *CurrentUser) GetAttributionString() string`

GetAttributionString returns the AttributionString field if non-nil, zero value otherwise.

### GetAttributionStringOk

`func (o *CurrentUser) GetAttributionStringOk() (*string, bool)`

GetAttributionStringOk returns a tuple with the AttributionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributionString

`func (o *CurrentUser) SetAttributionString(v string)`

SetAttributionString sets AttributionString field to given value.

### HasAttributionString

`func (o *CurrentUser) HasAttributionString() bool`

HasAttributionString returns a boolean if a field has been set.

### GetAuthToken

`func (o *CurrentUser) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *CurrentUser) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *CurrentUser) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *CurrentUser) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### GetAutoUploadOff

`func (o *CurrentUser) GetAutoUploadOff() bool`

GetAutoUploadOff returns the AutoUploadOff field if non-nil, zero value otherwise.

### GetAutoUploadOffOk

`func (o *CurrentUser) GetAutoUploadOffOk() (*bool, bool)`

GetAutoUploadOffOk returns a tuple with the AutoUploadOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUploadOff

`func (o *CurrentUser) SetAutoUploadOff(v bool)`

SetAutoUploadOff sets AutoUploadOff field to given value.

### HasAutoUploadOff

`func (o *CurrentUser) HasAutoUploadOff() bool`

HasAutoUploadOff returns a boolean if a field has been set.

### GetAvatarFileName

`func (o *CurrentUser) GetAvatarFileName() string`

GetAvatarFileName returns the AvatarFileName field if non-nil, zero value otherwise.

### GetAvatarFileNameOk

`func (o *CurrentUser) GetAvatarFileNameOk() (*string, bool)`

GetAvatarFileNameOk returns a tuple with the AvatarFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarFileName

`func (o *CurrentUser) SetAvatarFileName(v string)`

SetAvatarFileName sets AvatarFileName field to given value.

### HasAvatarFileName

`func (o *CurrentUser) HasAvatarFileName() bool`

HasAvatarFileName returns a boolean if a field has been set.

### GetBillingInfo

`func (o *CurrentUser) GetBillingInfo() string`

GetBillingInfo returns the BillingInfo field if non-nil, zero value otherwise.

### GetBillingInfoOk

`func (o *CurrentUser) GetBillingInfoOk() (*string, bool)`

GetBillingInfoOk returns a tuple with the BillingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInfo

`func (o *CurrentUser) SetBillingInfo(v string)`

SetBillingInfo sets BillingInfo field to given value.

### HasBillingInfo

`func (o *CurrentUser) HasBillingInfo() bool`

HasBillingInfo returns a boolean if a field has been set.

### GetCharitySubscriptionsLaunched

`func (o *CurrentUser) GetCharitySubscriptionsLaunched() bool`

GetCharitySubscriptionsLaunched returns the CharitySubscriptionsLaunched field if non-nil, zero value otherwise.

### GetCharitySubscriptionsLaunchedOk

`func (o *CurrentUser) GetCharitySubscriptionsLaunchedOk() (*bool, bool)`

GetCharitySubscriptionsLaunchedOk returns a tuple with the CharitySubscriptionsLaunched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharitySubscriptionsLaunched

`func (o *CurrentUser) SetCharitySubscriptionsLaunched(v bool)`

SetCharitySubscriptionsLaunched sets CharitySubscriptionsLaunched field to given value.

### HasCharitySubscriptionsLaunched

`func (o *CurrentUser) HasCharitySubscriptionsLaunched() bool`

HasCharitySubscriptionsLaunched returns a boolean if a field has been set.

### GetConfirmedEmail

`func (o *CurrentUser) GetConfirmedEmail() string`

GetConfirmedEmail returns the ConfirmedEmail field if non-nil, zero value otherwise.

### GetConfirmedEmailOk

`func (o *CurrentUser) GetConfirmedEmailOk() (*string, bool)`

GetConfirmedEmailOk returns a tuple with the ConfirmedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedEmail

`func (o *CurrentUser) SetConfirmedEmail(v string)`

SetConfirmedEmail sets ConfirmedEmail field to given value.

### HasConfirmedEmail

`func (o *CurrentUser) HasConfirmedEmail() bool`

HasConfirmedEmail returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CurrentUser) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CurrentUser) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CurrentUser) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CurrentUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrentSourceId

`func (o *CurrentUser) GetCurrentSourceId() string`

GetCurrentSourceId returns the CurrentSourceId field if non-nil, zero value otherwise.

### GetCurrentSourceIdOk

`func (o *CurrentUser) GetCurrentSourceIdOk() (*string, bool)`

GetCurrentSourceIdOk returns a tuple with the CurrentSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSourceId

`func (o *CurrentUser) SetCurrentSourceId(v string)`

SetCurrentSourceId sets CurrentSourceId field to given value.

### HasCurrentSourceId

`func (o *CurrentUser) HasCurrentSourceId() bool`

HasCurrentSourceId returns a boolean if a field has been set.

### GetEligibleForAppReviewPrompt

`func (o *CurrentUser) GetEligibleForAppReviewPrompt() bool`

GetEligibleForAppReviewPrompt returns the EligibleForAppReviewPrompt field if non-nil, zero value otherwise.

### GetEligibleForAppReviewPromptOk

`func (o *CurrentUser) GetEligibleForAppReviewPromptOk() (*bool, bool)`

GetEligibleForAppReviewPromptOk returns a tuple with the EligibleForAppReviewPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleForAppReviewPrompt

`func (o *CurrentUser) SetEligibleForAppReviewPrompt(v bool)`

SetEligibleForAppReviewPrompt sets EligibleForAppReviewPrompt field to given value.

### HasEligibleForAppReviewPrompt

`func (o *CurrentUser) HasEligibleForAppReviewPrompt() bool`

HasEligibleForAppReviewPrompt returns a boolean if a field has been set.

### GetEmail

`func (o *CurrentUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CurrentUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CurrentUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CurrentUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFeatures

`func (o *CurrentUser) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *CurrentUser) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *CurrentUser) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *CurrentUser) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetGooglePhotosDisabled

`func (o *CurrentUser) GetGooglePhotosDisabled() string`

GetGooglePhotosDisabled returns the GooglePhotosDisabled field if non-nil, zero value otherwise.

### GetGooglePhotosDisabledOk

`func (o *CurrentUser) GetGooglePhotosDisabledOk() (*string, bool)`

GetGooglePhotosDisabledOk returns a tuple with the GooglePhotosDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGooglePhotosDisabled

`func (o *CurrentUser) SetGooglePhotosDisabled(v string)`

SetGooglePhotosDisabled sets GooglePhotosDisabled field to given value.

### HasGooglePhotosDisabled

`func (o *CurrentUser) HasGooglePhotosDisabled() bool`

HasGooglePhotosDisabled returns a boolean if a field has been set.

### GetHasAccessToNewGooglePhotos

`func (o *CurrentUser) GetHasAccessToNewGooglePhotos() bool`

GetHasAccessToNewGooglePhotos returns the HasAccessToNewGooglePhotos field if non-nil, zero value otherwise.

### GetHasAccessToNewGooglePhotosOk

`func (o *CurrentUser) GetHasAccessToNewGooglePhotosOk() (*bool, bool)`

GetHasAccessToNewGooglePhotosOk returns a tuple with the HasAccessToNewGooglePhotos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccessToNewGooglePhotos

`func (o *CurrentUser) SetHasAccessToNewGooglePhotos(v bool)`

SetHasAccessToNewGooglePhotos sets HasAccessToNewGooglePhotos field to given value.

### HasHasAccessToNewGooglePhotos

`func (o *CurrentUser) HasHasAccessToNewGooglePhotos() bool`

HasHasAccessToNewGooglePhotos returns a boolean if a field has been set.

### GetHasFrame

`func (o *CurrentUser) GetHasFrame() bool`

GetHasFrame returns the HasFrame field if non-nil, zero value otherwise.

### GetHasFrameOk

`func (o *CurrentUser) GetHasFrameOk() (*bool, bool)`

GetHasFrameOk returns a tuple with the HasFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFrame

`func (o *CurrentUser) SetHasFrame(v bool)`

SetHasFrame sets HasFrame field to given value.

### HasHasFrame

`func (o *CurrentUser) HasHasFrame() bool`

HasHasFrame returns a boolean if a field has been set.

### GetId

`func (o *CurrentUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CurrentUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CurrentUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CurrentUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLatestAppVersion

`func (o *CurrentUser) GetLatestAppVersion() string`

GetLatestAppVersion returns the LatestAppVersion field if non-nil, zero value otherwise.

### GetLatestAppVersionOk

`func (o *CurrentUser) GetLatestAppVersionOk() (*string, bool)`

GetLatestAppVersionOk returns a tuple with the LatestAppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestAppVersion

`func (o *CurrentUser) SetLatestAppVersion(v string)`

SetLatestAppVersion sets LatestAppVersion field to given value.

### HasLatestAppVersion

`func (o *CurrentUser) HasLatestAppVersion() bool`

HasLatestAppVersion returns a boolean if a field has been set.

### GetLivePhotosLaunched

`func (o *CurrentUser) GetLivePhotosLaunched() bool`

GetLivePhotosLaunched returns the LivePhotosLaunched field if non-nil, zero value otherwise.

### GetLivePhotosLaunchedOk

`func (o *CurrentUser) GetLivePhotosLaunchedOk() (*bool, bool)`

GetLivePhotosLaunchedOk returns a tuple with the LivePhotosLaunched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivePhotosLaunched

`func (o *CurrentUser) SetLivePhotosLaunched(v bool)`

SetLivePhotosLaunched sets LivePhotosLaunched field to given value.

### HasLivePhotosLaunched

`func (o *CurrentUser) HasLivePhotosLaunched() bool`

HasLivePhotosLaunched returns a boolean if a field has been set.

### GetLocale

`func (o *CurrentUser) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *CurrentUser) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *CurrentUser) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *CurrentUser) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetName

`func (o *CurrentUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CurrentUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CurrentUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CurrentUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrintSubscriptionsLaunched

`func (o *CurrentUser) GetPrintSubscriptionsLaunched() bool`

GetPrintSubscriptionsLaunched returns the PrintSubscriptionsLaunched field if non-nil, zero value otherwise.

### GetPrintSubscriptionsLaunchedOk

`func (o *CurrentUser) GetPrintSubscriptionsLaunchedOk() (*bool, bool)`

GetPrintSubscriptionsLaunchedOk returns a tuple with the PrintSubscriptionsLaunched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintSubscriptionsLaunched

`func (o *CurrentUser) SetPrintSubscriptionsLaunched(v bool)`

SetPrintSubscriptionsLaunched sets PrintSubscriptionsLaunched field to given value.

### HasPrintSubscriptionsLaunched

`func (o *CurrentUser) HasPrintSubscriptionsLaunched() bool`

HasPrintSubscriptionsLaunched returns a boolean if a field has been set.

### GetRecurlyAccountCode

`func (o *CurrentUser) GetRecurlyAccountCode() string`

GetRecurlyAccountCode returns the RecurlyAccountCode field if non-nil, zero value otherwise.

### GetRecurlyAccountCodeOk

`func (o *CurrentUser) GetRecurlyAccountCodeOk() (*string, bool)`

GetRecurlyAccountCodeOk returns a tuple with the RecurlyAccountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurlyAccountCode

`func (o *CurrentUser) SetRecurlyAccountCode(v string)`

SetRecurlyAccountCode sets RecurlyAccountCode field to given value.

### HasRecurlyAccountCode

`func (o *CurrentUser) HasRecurlyAccountCode() bool`

HasRecurlyAccountCode returns a boolean if a field has been set.

### GetShortId

`func (o *CurrentUser) GetShortId() string`

GetShortId returns the ShortId field if non-nil, zero value otherwise.

### GetShortIdOk

`func (o *CurrentUser) GetShortIdOk() (*string, bool)`

GetShortIdOk returns a tuple with the ShortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortId

`func (o *CurrentUser) SetShortId(v string)`

SetShortId sets ShortId field to given value.

### HasShortId

`func (o *CurrentUser) HasShortId() bool`

HasShortId returns a boolean if a field has been set.

### GetShowPushPrompt

`func (o *CurrentUser) GetShowPushPrompt() bool`

GetShowPushPrompt returns the ShowPushPrompt field if non-nil, zero value otherwise.

### GetShowPushPromptOk

`func (o *CurrentUser) GetShowPushPromptOk() (*bool, bool)`

GetShowPushPromptOk returns a tuple with the ShowPushPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPushPrompt

`func (o *CurrentUser) SetShowPushPrompt(v bool)`

SetShowPushPrompt sets ShowPushPrompt field to given value.

### HasShowPushPrompt

`func (o *CurrentUser) HasShowPushPrompt() bool`

HasShowPushPrompt returns a boolean if a field has been set.

### GetSmartAlbumsOff

`func (o *CurrentUser) GetSmartAlbumsOff() bool`

GetSmartAlbumsOff returns the SmartAlbumsOff field if non-nil, zero value otherwise.

### GetSmartAlbumsOffOk

`func (o *CurrentUser) GetSmartAlbumsOffOk() (*bool, bool)`

GetSmartAlbumsOffOk returns a tuple with the SmartAlbumsOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartAlbumsOff

`func (o *CurrentUser) SetSmartAlbumsOff(v bool)`

SetSmartAlbumsOff sets SmartAlbumsOff field to given value.

### HasSmartAlbumsOff

`func (o *CurrentUser) HasSmartAlbumsOff() bool`

HasSmartAlbumsOff returns a boolean if a field has been set.

### GetSmartSuggestionsOff

`func (o *CurrentUser) GetSmartSuggestionsOff() bool`

GetSmartSuggestionsOff returns the SmartSuggestionsOff field if non-nil, zero value otherwise.

### GetSmartSuggestionsOffOk

`func (o *CurrentUser) GetSmartSuggestionsOffOk() (*bool, bool)`

GetSmartSuggestionsOffOk returns a tuple with the SmartSuggestionsOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartSuggestionsOff

`func (o *CurrentUser) SetSmartSuggestionsOff(v bool)`

SetSmartSuggestionsOff sets SmartSuggestionsOff field to given value.

### HasSmartSuggestionsOff

`func (o *CurrentUser) HasSmartSuggestionsOff() bool`

HasSmartSuggestionsOff returns a boolean if a field has been set.

### GetSubscriptionsLaunched

`func (o *CurrentUser) GetSubscriptionsLaunched() bool`

GetSubscriptionsLaunched returns the SubscriptionsLaunched field if non-nil, zero value otherwise.

### GetSubscriptionsLaunchedOk

`func (o *CurrentUser) GetSubscriptionsLaunchedOk() (*bool, bool)`

GetSubscriptionsLaunchedOk returns a tuple with the SubscriptionsLaunched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionsLaunched

`func (o *CurrentUser) SetSubscriptionsLaunched(v bool)`

SetSubscriptionsLaunched sets SubscriptionsLaunched field to given value.

### HasSubscriptionsLaunched

`func (o *CurrentUser) HasSubscriptionsLaunched() bool`

HasSubscriptionsLaunched returns a boolean if a field has been set.

### GetTestAccount

`func (o *CurrentUser) GetTestAccount() string`

GetTestAccount returns the TestAccount field if non-nil, zero value otherwise.

### GetTestAccountOk

`func (o *CurrentUser) GetTestAccountOk() (*string, bool)`

GetTestAccountOk returns a tuple with the TestAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestAccount

`func (o *CurrentUser) SetTestAccount(v string)`

SetTestAccount sets TestAccount field to given value.

### HasTestAccount

`func (o *CurrentUser) HasTestAccount() bool`

HasTestAccount returns a boolean if a field has been set.

### GetThanksLaunched

`func (o *CurrentUser) GetThanksLaunched() bool`

GetThanksLaunched returns the ThanksLaunched field if non-nil, zero value otherwise.

### GetThanksLaunchedOk

`func (o *CurrentUser) GetThanksLaunchedOk() (*bool, bool)`

GetThanksLaunchedOk returns a tuple with the ThanksLaunched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThanksLaunched

`func (o *CurrentUser) SetThanksLaunched(v bool)`

SetThanksLaunched sets ThanksLaunched field to given value.

### HasThanksLaunched

`func (o *CurrentUser) HasThanksLaunched() bool`

HasThanksLaunched returns a boolean if a field has been set.

### GetTooltipAddPhotosSeen

`func (o *CurrentUser) GetTooltipAddPhotosSeen() bool`

GetTooltipAddPhotosSeen returns the TooltipAddPhotosSeen field if non-nil, zero value otherwise.

### GetTooltipAddPhotosSeenOk

`func (o *CurrentUser) GetTooltipAddPhotosSeenOk() (*bool, bool)`

GetTooltipAddPhotosSeenOk returns a tuple with the TooltipAddPhotosSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTooltipAddPhotosSeen

`func (o *CurrentUser) SetTooltipAddPhotosSeen(v bool)`

SetTooltipAddPhotosSeen sets TooltipAddPhotosSeen field to given value.

### HasTooltipAddPhotosSeen

`func (o *CurrentUser) HasTooltipAddPhotosSeen() bool`

HasTooltipAddPhotosSeen returns a boolean if a field has been set.

### GetTooltipAddedPhotosSeen

`func (o *CurrentUser) GetTooltipAddedPhotosSeen() bool`

GetTooltipAddedPhotosSeen returns the TooltipAddedPhotosSeen field if non-nil, zero value otherwise.

### GetTooltipAddedPhotosSeenOk

`func (o *CurrentUser) GetTooltipAddedPhotosSeenOk() (*bool, bool)`

GetTooltipAddedPhotosSeenOk returns a tuple with the TooltipAddedPhotosSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTooltipAddedPhotosSeen

`func (o *CurrentUser) SetTooltipAddedPhotosSeen(v bool)`

SetTooltipAddedPhotosSeen sets TooltipAddedPhotosSeen field to given value.

### HasTooltipAddedPhotosSeen

`func (o *CurrentUser) HasTooltipAddedPhotosSeen() bool`

HasTooltipAddedPhotosSeen returns a boolean if a field has been set.

### GetTooltipGesturesSeen

`func (o *CurrentUser) GetTooltipGesturesSeen() bool`

GetTooltipGesturesSeen returns the TooltipGesturesSeen field if non-nil, zero value otherwise.

### GetTooltipGesturesSeenOk

`func (o *CurrentUser) GetTooltipGesturesSeenOk() (*bool, bool)`

GetTooltipGesturesSeenOk returns a tuple with the TooltipGesturesSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTooltipGesturesSeen

`func (o *CurrentUser) SetTooltipGesturesSeen(v bool)`

SetTooltipGesturesSeen sets TooltipGesturesSeen field to given value.

### HasTooltipGesturesSeen

`func (o *CurrentUser) HasTooltipGesturesSeen() bool`

HasTooltipGesturesSeen returns a boolean if a field has been set.

### GetTooltipInboxSeen

`func (o *CurrentUser) GetTooltipInboxSeen() bool`

GetTooltipInboxSeen returns the TooltipInboxSeen field if non-nil, zero value otherwise.

### GetTooltipInboxSeenOk

`func (o *CurrentUser) GetTooltipInboxSeenOk() (*bool, bool)`

GetTooltipInboxSeenOk returns a tuple with the TooltipInboxSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTooltipInboxSeen

`func (o *CurrentUser) SetTooltipInboxSeen(v bool)`

SetTooltipInboxSeen sets TooltipInboxSeen field to given value.

### HasTooltipInboxSeen

`func (o *CurrentUser) HasTooltipInboxSeen() bool`

HasTooltipInboxSeen returns a boolean if a field has been set.

### GetTooltipManageFramesSeen

`func (o *CurrentUser) GetTooltipManageFramesSeen() bool`

GetTooltipManageFramesSeen returns the TooltipManageFramesSeen field if non-nil, zero value otherwise.

### GetTooltipManageFramesSeenOk

`func (o *CurrentUser) GetTooltipManageFramesSeenOk() (*bool, bool)`

GetTooltipManageFramesSeenOk returns a tuple with the TooltipManageFramesSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTooltipManageFramesSeen

`func (o *CurrentUser) SetTooltipManageFramesSeen(v bool)`

SetTooltipManageFramesSeen sets TooltipManageFramesSeen field to given value.

### HasTooltipManageFramesSeen

`func (o *CurrentUser) HasTooltipManageFramesSeen() bool`

HasTooltipManageFramesSeen returns a boolean if a field has been set.

### GetTooltipSettingsSeen

`func (o *CurrentUser) GetTooltipSettingsSeen() bool`

GetTooltipSettingsSeen returns the TooltipSettingsSeen field if non-nil, zero value otherwise.

### GetTooltipSettingsSeenOk

`func (o *CurrentUser) GetTooltipSettingsSeenOk() (*bool, bool)`

GetTooltipSettingsSeenOk returns a tuple with the TooltipSettingsSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTooltipSettingsSeen

`func (o *CurrentUser) SetTooltipSettingsSeen(v bool)`

SetTooltipSettingsSeen sets TooltipSettingsSeen field to given value.

### HasTooltipSettingsSeen

`func (o *CurrentUser) HasTooltipSettingsSeen() bool`

HasTooltipSettingsSeen returns a boolean if a field has been set.

### GetUnconfirmedEmail

`func (o *CurrentUser) GetUnconfirmedEmail() string`

GetUnconfirmedEmail returns the UnconfirmedEmail field if non-nil, zero value otherwise.

### GetUnconfirmedEmailOk

`func (o *CurrentUser) GetUnconfirmedEmailOk() (*string, bool)`

GetUnconfirmedEmailOk returns a tuple with the UnconfirmedEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnconfirmedEmail

`func (o *CurrentUser) SetUnconfirmedEmail(v string)`

SetUnconfirmedEmail sets UnconfirmedEmail field to given value.

### HasUnconfirmedEmail

`func (o *CurrentUser) HasUnconfirmedEmail() bool`

HasUnconfirmedEmail returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CurrentUser) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CurrentUser) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CurrentUser) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CurrentUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVerboseLoggingEnabled

`func (o *CurrentUser) GetVerboseLoggingEnabled() bool`

GetVerboseLoggingEnabled returns the VerboseLoggingEnabled field if non-nil, zero value otherwise.

### GetVerboseLoggingEnabledOk

`func (o *CurrentUser) GetVerboseLoggingEnabledOk() (*bool, bool)`

GetVerboseLoggingEnabledOk returns a tuple with the VerboseLoggingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerboseLoggingEnabled

`func (o *CurrentUser) SetVerboseLoggingEnabled(v bool)`

SetVerboseLoggingEnabled sets VerboseLoggingEnabled field to given value.

### HasVerboseLoggingEnabled

`func (o *CurrentUser) HasVerboseLoggingEnabled() bool`

HasVerboseLoggingEnabled returns a boolean if a field has been set.

### GetWarnSmartAlbumsDeprecated

`func (o *CurrentUser) GetWarnSmartAlbumsDeprecated() bool`

GetWarnSmartAlbumsDeprecated returns the WarnSmartAlbumsDeprecated field if non-nil, zero value otherwise.

### GetWarnSmartAlbumsDeprecatedOk

`func (o *CurrentUser) GetWarnSmartAlbumsDeprecatedOk() (*bool, bool)`

GetWarnSmartAlbumsDeprecatedOk returns a tuple with the WarnSmartAlbumsDeprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnSmartAlbumsDeprecated

`func (o *CurrentUser) SetWarnSmartAlbumsDeprecated(v bool)`

SetWarnSmartAlbumsDeprecated sets WarnSmartAlbumsDeprecated field to given value.

### HasWarnSmartAlbumsDeprecated

`func (o *CurrentUser) HasWarnSmartAlbumsDeprecated() bool`

HasWarnSmartAlbumsDeprecated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


