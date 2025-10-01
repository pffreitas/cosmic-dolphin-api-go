# BookmarkMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpenGraph** | Pointer to [**OpenGraphMetadata**](OpenGraphMetadata.md) |  | [optional] 
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


