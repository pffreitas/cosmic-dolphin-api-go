# Bookmark

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SourceUrl** | **string** |  | 
**CollectionId** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**IsArchived** | Pointer to **bool** |  | [optional] 
**IsFavorite** | Pointer to **bool** |  | [optional] 
**CosmicImages** | Pointer to [**[]BookmarkImage**](BookmarkImage.md) |  | [optional] 
**CosmicLinks** | Pointer to [**[]BookmarkLink**](BookmarkLink.md) |  | [optional] 
**CosmicSummary** | Pointer to **string** |  | [optional] 
**CosmicTags** | Pointer to **[]string** |  | [optional] 
**Metadata** | Pointer to [**BookmarkMetadata**](BookmarkMetadata.md) |  | [optional] 
**UserId** | **string** |  | 

## Methods

### NewBookmark

`func NewBookmark(id string, sourceUrl string, userId string, ) *Bookmark`

NewBookmark instantiates a new Bookmark object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookmarkWithDefaults

`func NewBookmarkWithDefaults() *Bookmark`

NewBookmarkWithDefaults instantiates a new Bookmark object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Bookmark) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Bookmark) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Bookmark) SetId(v string)`

SetId sets Id field to given value.


### GetSourceUrl

`func (o *Bookmark) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *Bookmark) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *Bookmark) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.


### GetCollectionId

`func (o *Bookmark) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *Bookmark) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *Bookmark) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *Bookmark) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### GetTitle

`func (o *Bookmark) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Bookmark) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Bookmark) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Bookmark) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetIsArchived

`func (o *Bookmark) GetIsArchived() bool`

GetIsArchived returns the IsArchived field if non-nil, zero value otherwise.

### GetIsArchivedOk

`func (o *Bookmark) GetIsArchivedOk() (*bool, bool)`

GetIsArchivedOk returns a tuple with the IsArchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArchived

`func (o *Bookmark) SetIsArchived(v bool)`

SetIsArchived sets IsArchived field to given value.

### HasIsArchived

`func (o *Bookmark) HasIsArchived() bool`

HasIsArchived returns a boolean if a field has been set.

### GetIsFavorite

`func (o *Bookmark) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *Bookmark) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *Bookmark) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *Bookmark) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### GetCosmicImages

`func (o *Bookmark) GetCosmicImages() []BookmarkImage`

GetCosmicImages returns the CosmicImages field if non-nil, zero value otherwise.

### GetCosmicImagesOk

`func (o *Bookmark) GetCosmicImagesOk() (*[]BookmarkImage, bool)`

GetCosmicImagesOk returns a tuple with the CosmicImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosmicImages

`func (o *Bookmark) SetCosmicImages(v []BookmarkImage)`

SetCosmicImages sets CosmicImages field to given value.

### HasCosmicImages

`func (o *Bookmark) HasCosmicImages() bool`

HasCosmicImages returns a boolean if a field has been set.

### GetCosmicLinks

`func (o *Bookmark) GetCosmicLinks() []BookmarkLink`

GetCosmicLinks returns the CosmicLinks field if non-nil, zero value otherwise.

### GetCosmicLinksOk

`func (o *Bookmark) GetCosmicLinksOk() (*[]BookmarkLink, bool)`

GetCosmicLinksOk returns a tuple with the CosmicLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosmicLinks

`func (o *Bookmark) SetCosmicLinks(v []BookmarkLink)`

SetCosmicLinks sets CosmicLinks field to given value.

### HasCosmicLinks

`func (o *Bookmark) HasCosmicLinks() bool`

HasCosmicLinks returns a boolean if a field has been set.

### GetCosmicSummary

`func (o *Bookmark) GetCosmicSummary() string`

GetCosmicSummary returns the CosmicSummary field if non-nil, zero value otherwise.

### GetCosmicSummaryOk

`func (o *Bookmark) GetCosmicSummaryOk() (*string, bool)`

GetCosmicSummaryOk returns a tuple with the CosmicSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosmicSummary

`func (o *Bookmark) SetCosmicSummary(v string)`

SetCosmicSummary sets CosmicSummary field to given value.

### HasCosmicSummary

`func (o *Bookmark) HasCosmicSummary() bool`

HasCosmicSummary returns a boolean if a field has been set.

### GetCosmicTags

`func (o *Bookmark) GetCosmicTags() []string`

GetCosmicTags returns the CosmicTags field if non-nil, zero value otherwise.

### GetCosmicTagsOk

`func (o *Bookmark) GetCosmicTagsOk() (*[]string, bool)`

GetCosmicTagsOk returns a tuple with the CosmicTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosmicTags

`func (o *Bookmark) SetCosmicTags(v []string)`

SetCosmicTags sets CosmicTags field to given value.

### HasCosmicTags

`func (o *Bookmark) HasCosmicTags() bool`

HasCosmicTags returns a boolean if a field has been set.

### GetMetadata

`func (o *Bookmark) GetMetadata() BookmarkMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Bookmark) GetMetadataOk() (*BookmarkMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Bookmark) SetMetadata(v BookmarkMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Bookmark) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetUserId

`func (o *Bookmark) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Bookmark) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Bookmark) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


