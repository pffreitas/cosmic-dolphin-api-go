/*
Cosmic Dolphin

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cosmicdolphinapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CosmicError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CosmicError{}

// CosmicError struct for CosmicError
type CosmicError struct {
	Code int32 `json:"code"`
	Message string `json:"message"`
}

type _CosmicError CosmicError

// NewCosmicError instantiates a new CosmicError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCosmicError(code int32, message string) *CosmicError {
	this := CosmicError{}
	this.Code = code
	this.Message = message
	return &this
}

// NewCosmicErrorWithDefaults instantiates a new CosmicError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCosmicErrorWithDefaults() *CosmicError {
	this := CosmicError{}
	return &this
}

// GetCode returns the Code field value
func (o *CosmicError) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *CosmicError) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *CosmicError) SetCode(v int32) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *CosmicError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *CosmicError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *CosmicError) SetMessage(v string) {
	o.Message = v
}

func (o CosmicError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CosmicError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *CosmicError) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"message",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCosmicError := _CosmicError{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCosmicError)

	if err != nil {
		return err
	}

	*o = CosmicError(varCosmicError)

	return err
}

type NullableCosmicError struct {
	value *CosmicError
	isSet bool
}

func (v NullableCosmicError) Get() *CosmicError {
	return v.value
}

func (v *NullableCosmicError) Set(val *CosmicError) {
	v.value = val
	v.isSet = true
}

func (v NullableCosmicError) IsSet() bool {
	return v.isSet
}

func (v *NullableCosmicError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCosmicError(val *CosmicError) *NullableCosmicError {
	return &NullableCosmicError{value: val, isSet: true}
}

func (v NullableCosmicError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCosmicError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


