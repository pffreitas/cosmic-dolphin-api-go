# \PipelinesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PipelinesFindPipelinesByRefId**](PipelinesAPI.md#PipelinesFindPipelinesByRefId) | **Get** /pipelines/{refId} | 



## PipelinesFindPipelinesByRefId

> []Pipeline PipelinesFindPipelinesByRefId(ctx, refId).Execute()



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
	resp, r, err := apiClient.PipelinesAPI.PipelinesFindPipelinesByRefId(context.Background(), refId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.PipelinesFindPipelinesByRefId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PipelinesFindPipelinesByRefId`: []Pipeline
	fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.PipelinesFindPipelinesByRefId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**refId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPipelinesFindPipelinesByRefIdRequest struct via the builder pattern


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

