# CreateBookmarkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceUrl** | **string** | The source URL to bookmark | 
**CollectionId** | Pointer to **string** | Optional collection ID to add the bookmark to | [optional] 

## Methods

### NewCreateBookmarkRequest

`func NewCreateBookmarkRequest(sourceUrl string, ) *CreateBookmarkRequest`

NewCreateBookmarkRequest instantiates a new CreateBookmarkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBookmarkRequestWithDefaults

`func NewCreateBookmarkRequestWithDefaults() *CreateBookmarkRequest`

NewCreateBookmarkRequestWithDefaults instantiates a new CreateBookmarkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceUrl

`func (o *CreateBookmarkRequest) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *CreateBookmarkRequest) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *CreateBookmarkRequest) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.


### GetCollectionId

`func (o *CreateBookmarkRequest) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *CreateBookmarkRequest) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *CreateBookmarkRequest) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *CreateBookmarkRequest) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


