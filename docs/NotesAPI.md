# \NotesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotesCreate**](NotesAPI.md#NotesCreate) | **Post** /notes | 
[**NotesFindAll**](NotesAPI.md#NotesFindAll) | **Get** /notes | 
[**NotesFindById**](NotesAPI.md#NotesFindById) | **Get** /notes/{id} | 



## NotesCreate

> Note NotesCreate(ctx).CreateNoteRequest(createNoteRequest).Execute()



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
	createNoteRequest := *openapiclient.NewCreateNoteRequest("Body_example", openapiclient.NoteType("fup")) // CreateNoteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotesAPI.NotesCreate(context.Background()).CreateNoteRequest(createNoteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotesAPI.NotesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotesCreate`: Note
	fmt.Fprintf(os.Stdout, "Response from `NotesAPI.NotesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNoteRequest** | [**CreateNoteRequest**](CreateNoteRequest.md) |  | 

### Return type

[**Note**](Note.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotesFindAll

> []Note NotesFindAll(ctx).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotesAPI.NotesFindAll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotesAPI.NotesFindAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotesFindAll`: []Note
	fmt.Fprintf(os.Stdout, "Response from `NotesAPI.NotesFindAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNotesFindAllRequest struct via the builder pattern


### Return type

[**[]Note**](Note.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotesFindById

> Note NotesFindById(ctx, id).Execute()



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotesAPI.NotesFindById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotesAPI.NotesFindById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotesFindById`: Note
	fmt.Fprintf(os.Stdout, "Response from `NotesAPI.NotesFindById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotesFindByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Note**](Note.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

