# \PipelinesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PipelinesFindByRefId**](PipelinesAPI.md#PipelinesFindByRefId) | **Get** /pipelines/{refId} | 



## PipelinesFindByRefId

> []Pipeline PipelinesFindByRefId(ctx, refId).Execute()



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
	refId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PipelinesAPI.PipelinesFindByRefId(context.Background(), refId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.PipelinesFindByRefId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PipelinesFindByRefId`: []Pipeline
	fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.PipelinesFindByRefId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**refId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPipelinesFindByRefIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Pipeline**](Pipeline.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

