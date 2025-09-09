# Resource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ResourceType**](ResourceType.md) |  | 
**OpenGraph** | Pointer to [**OpenGraph**](OpenGraph.md) |  | [optional] 

## Methods

### NewResource

`func NewResource(type_ ResourceType, ) *Resource`

NewResource instantiates a new Resource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceWithDefaults

`func NewResourceWithDefaults() *Resource`

NewResourceWithDefaults instantiates a new Resource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Resource) GetType() ResourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Resource) GetTypeOk() (*ResourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Resource) SetType(v ResourceType)`

SetType sets Type field to given value.


### GetOpenGraph

`func (o *Resource) GetOpenGraph() OpenGraph`

GetOpenGraph returns the OpenGraph field if non-nil, zero value otherwise.

### GetOpenGraphOk

`func (o *Resource) GetOpenGraphOk() (*OpenGraph, bool)`

GetOpenGraphOk returns a tuple with the OpenGraph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenGraph

`func (o *Resource) SetOpenGraph(v OpenGraph)`

SetOpenGraph sets OpenGraph field to given value.

### HasOpenGraph

`func (o *Resource) HasOpenGraph() bool`

HasOpenGraph returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


