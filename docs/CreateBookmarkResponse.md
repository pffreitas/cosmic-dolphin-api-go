# CreateBookmarkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bookmark** | [**Bookmark**](Bookmark.md) |  | 
**Message** | **string** |  | 

## Methods

### NewCreateBookmarkResponse

`func NewCreateBookmarkResponse(bookmark Bookmark, message string, ) *CreateBookmarkResponse`

NewCreateBookmarkResponse instantiates a new CreateBookmarkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBookmarkResponseWithDefaults

`func NewCreateBookmarkResponseWithDefaults() *CreateBookmarkResponse`

NewCreateBookmarkResponseWithDefaults instantiates a new CreateBookmarkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBookmark

`func (o *CreateBookmarkResponse) GetBookmark() Bookmark`

GetBookmark returns the Bookmark field if non-nil, zero value otherwise.

### GetBookmarkOk

`func (o *CreateBookmarkResponse) GetBookmarkOk() (*Bookmark, bool)`

GetBookmarkOk returns a tuple with the Bookmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookmark

`func (o *CreateBookmarkResponse) SetBookmark(v Bookmark)`

SetBookmark sets Bookmark field to given value.


### GetMessage

`func (o *CreateBookmarkResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateBookmarkResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateBookmarkResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


