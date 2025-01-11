# CosmicError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** |  | 
**Message** | **string** |  | 

## Methods

### NewCosmicError

`func NewCosmicError(code int32, message string, ) *CosmicError`

NewCosmicError instantiates a new CosmicError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCosmicErrorWithDefaults

`func NewCosmicErrorWithDefaults() *CosmicError`

NewCosmicErrorWithDefaults instantiates a new CosmicError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CosmicError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CosmicError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CosmicError) SetCode(v int32)`

SetCode sets Code field to given value.


### GetMessage

`func (o *CosmicError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CosmicError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CosmicError) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


