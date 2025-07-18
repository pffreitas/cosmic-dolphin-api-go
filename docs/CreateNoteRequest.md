# CreateNoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **string** |  | 
**Type** | [**NoteType**](NoteType.md) |  | 

## Methods

### NewCreateNoteRequest

`func NewCreateNoteRequest(body string, type_ NoteType, ) *CreateNoteRequest`

NewCreateNoteRequest instantiates a new CreateNoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNoteRequestWithDefaults

`func NewCreateNoteRequestWithDefaults() *CreateNoteRequest`

NewCreateNoteRequestWithDefaults instantiates a new CreateNoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *CreateNoteRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CreateNoteRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CreateNoteRequest) SetBody(v string)`

SetBody sets Body field to given value.


### GetType

`func (o *CreateNoteRequest) GetType() NoteType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateNoteRequest) GetTypeOk() (*NoteType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateNoteRequest) SetType(v NoteType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


