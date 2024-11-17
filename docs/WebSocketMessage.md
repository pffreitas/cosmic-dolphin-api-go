# WebSocketMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** |  | 
**Data** | [**WebSocketMessageData**](WebSocketMessageData.md) |  | 

## Methods

### NewWebSocketMessage

`func NewWebSocketMessage(event string, data WebSocketMessageData, ) *WebSocketMessage`

NewWebSocketMessage instantiates a new WebSocketMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebSocketMessageWithDefaults

`func NewWebSocketMessageWithDefaults() *WebSocketMessage`

NewWebSocketMessageWithDefaults instantiates a new WebSocketMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *WebSocketMessage) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebSocketMessage) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebSocketMessage) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetData

`func (o *WebSocketMessage) GetData() WebSocketMessageData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebSocketMessage) GetDataOk() (*WebSocketMessageData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebSocketMessage) SetData(v WebSocketMessageData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


