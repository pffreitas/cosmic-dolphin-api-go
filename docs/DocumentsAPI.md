# \DocumentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DocumentServiceInsertDocument**](DocumentsAPI.md#DocumentServiceInsertDocument) | **Post** /insert-document | 



## DocumentServiceInsertDocument

> DocumentServiceInsertDocument(ctx).InsertDocumentRequest(insertDocumentRequest).Execute()



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
	insertDocumentRequest := *openapiclient.NewInsertDocumentRequest("Type_example", "Source_example") // InsertDocumentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DocumentsAPI.DocumentServiceInsertDocument(context.Background()).InsertDocumentRequest(insertDocumentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.DocumentServiceInsertDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDocumentServiceInsertDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insertDocumentRequest** | [**InsertDocumentRequest**](InsertDocumentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

