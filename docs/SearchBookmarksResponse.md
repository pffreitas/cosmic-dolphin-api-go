# SearchBookmarksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bookmarks** | [**[]Bookmark**](Bookmark.md) |  | 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewSearchBookmarksResponse

`func NewSearchBookmarksResponse(bookmarks []Bookmark, ) *SearchBookmarksResponse`

NewSearchBookmarksResponse instantiates a new SearchBookmarksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchBookmarksResponseWithDefaults

`func NewSearchBookmarksResponseWithDefaults() *SearchBookmarksResponse`

NewSearchBookmarksResponseWithDefaults instantiates a new SearchBookmarksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBookmarks

`func (o *SearchBookmarksResponse) GetBookmarks() []Bookmark`

GetBookmarks returns the Bookmarks field if non-nil, zero value otherwise.

### GetBookmarksOk

`func (o *SearchBookmarksResponse) GetBookmarksOk() (*[]Bookmark, bool)`

GetBookmarksOk returns a tuple with the Bookmarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookmarks

`func (o *SearchBookmarksResponse) SetBookmarks(v []Bookmark)`

SetBookmarks sets Bookmarks field to given value.


### GetTotal

`func (o *SearchBookmarksResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SearchBookmarksResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SearchBookmarksResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SearchBookmarksResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


