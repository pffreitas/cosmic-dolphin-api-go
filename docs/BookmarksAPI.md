# \BookmarksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BookmarksCreate**](BookmarksAPI.md#BookmarksCreate) | **Post** /bookmarks | 
[**BookmarksList**](BookmarksAPI.md#BookmarksList) | **Get** /bookmarks | 



## BookmarksCreate

> CreateBookmarkResponse BookmarksCreate(ctx).CreateBookmarkRequest(createBookmarkRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pffreitas/cosmic-dolphin-api-go"
)

func main() {
	createBookmarkRequest := *openapiclient.NewCreateBookmarkRequest("SourceUrl_example") // CreateBookmarkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BookmarksAPI.BookmarksCreate(context.Background()).CreateBookmarkRequest(createBookmarkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BookmarksAPI.BookmarksCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BookmarksCreate`: CreateBookmarkResponse
	fmt.Fprintf(os.Stdout, "Response from `BookmarksAPI.BookmarksCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBookmarksCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBookmarkRequest** | [**CreateBookmarkRequest**](CreateBookmarkRequest.md) |  | 

### Return type

[**CreateBookmarkResponse**](CreateBookmarkResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BookmarksList

> GetBookmarksResponse BookmarksList(ctx).CollectionId(collectionId).Limit(limit).Offset(offset).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pffreitas/cosmic-dolphin-api-go"
)

func main() {
	collectionId := "collectionId_example" // string | Filter bookmarks by collection ID (optional)
	limit := int32(56) // int32 | Maximum number of bookmarks to return (optional) (default to 50)
	offset := int32(56) // int32 | Number of bookmarks to skip (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BookmarksAPI.BookmarksList(context.Background()).CollectionId(collectionId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BookmarksAPI.BookmarksList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BookmarksList`: GetBookmarksResponse
	fmt.Fprintf(os.Stdout, "Response from `BookmarksAPI.BookmarksList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBookmarksListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collectionId** | **string** | Filter bookmarks by collection ID | 
 **limit** | **int32** | Maximum number of bookmarks to return | [default to 50]
 **offset** | **int32** | Number of bookmarks to skip | [default to 0]

### Return type

[**GetBookmarksResponse**](GetBookmarksResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

