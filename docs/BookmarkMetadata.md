# BookmarkMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpenGraph** | Pointer to [**OpenGraphMetadata**](OpenGraphMetadata.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Favicon** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**WordCount** | Pointer to **int32** |  | [optional] 
**ReadingTime** | Pointer to **int32** |  | [optional] 

## Methods

### NewBookmarkMetadata

`func NewBookmarkMetadata() *BookmarkMetadata`

NewBookmarkMetadata instantiates a new BookmarkMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookmarkMetadataWithDefaults

`func NewBookmarkMetadataWithDefaults() *BookmarkMetadata`

NewBookmarkMetadataWithDefaults instantiates a new BookmarkMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenGraph

`func (o *BookmarkMetadata) GetOpenGraph() OpenGraphMetadata`

GetOpenGraph returns the OpenGraph field if non-nil, zero value otherwise.

### GetOpenGraphOk

`func (o *BookmarkMetadata) GetOpenGraphOk() (*OpenGraphMetadata, bool)`

GetOpenGraphOk returns a tuple with the OpenGraph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenGraph

`func (o *BookmarkMetadata) SetOpenGraph(v OpenGraphMetadata)`

SetOpenGraph sets OpenGraph field to given value.

### HasOpenGraph

`func (o *BookmarkMetadata) HasOpenGraph() bool`

HasOpenGraph returns a boolean if a field has been set.

### GetTitle

`func (o *BookmarkMetadata) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BookmarkMetadata) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BookmarkMetadata) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BookmarkMetadata) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *BookmarkMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BookmarkMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BookmarkMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BookmarkMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFavicon

`func (o *BookmarkMetadata) GetFavicon() string`

GetFavicon returns the Favicon field if non-nil, zero value otherwise.

### GetFaviconOk

`func (o *BookmarkMetadata) GetFaviconOk() (*string, bool)`

GetFaviconOk returns a tuple with the Favicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavicon

`func (o *BookmarkMetadata) SetFavicon(v string)`

SetFavicon sets Favicon field to given value.

### HasFavicon

`func (o *BookmarkMetadata) HasFavicon() bool`

HasFavicon returns a boolean if a field has been set.

### GetContentType

`func (o *BookmarkMetadata) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *BookmarkMetadata) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *BookmarkMetadata) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *BookmarkMetadata) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetWordCount

`func (o *BookmarkMetadata) GetWordCount() int32`

GetWordCount returns the WordCount field if non-nil, zero value otherwise.

### GetWordCountOk

`func (o *BookmarkMetadata) GetWordCountOk() (*int32, bool)`

GetWordCountOk returns a tuple with the WordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWordCount

`func (o *BookmarkMetadata) SetWordCount(v int32)`

SetWordCount sets WordCount field to given value.

### HasWordCount

`func (o *BookmarkMetadata) HasWordCount() bool`

HasWordCount returns a boolean if a field has been set.

### GetReadingTime

`func (o *BookmarkMetadata) GetReadingTime() int32`

GetReadingTime returns the ReadingTime field if non-nil, zero value otherwise.

### GetReadingTimeOk

`func (o *BookmarkMetadata) GetReadingTimeOk() (*int32, bool)`

GetReadingTimeOk returns a tuple with the ReadingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingTime

`func (o *BookmarkMetadata) SetReadingTime(v int32)`

SetReadingTime sets ReadingTime field to given value.

### HasReadingTime

`func (o *BookmarkMetadata) HasReadingTime() bool`

HasReadingTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


