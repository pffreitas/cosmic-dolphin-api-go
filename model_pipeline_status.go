/*
Cosmic Dolphin

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cosmicdolphinapi

import (
	"encoding/json"
	"fmt"
)

// PipelineStatus the model 'PipelineStatus'
type PipelineStatus string

// List of PipelineStatus
const (
	PENDING PipelineStatus = "pending"
	RUNNING PipelineStatus = "running"
	COMPLETE PipelineStatus = "complete"
	FAILED PipelineStatus = "failed"
)

// All allowed values of PipelineStatus enum
var AllowedPipelineStatusEnumValues = []PipelineStatus{
	"pending",
	"running",
	"complete",
	"failed",
}

func (v *PipelineStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PipelineStatus(value)
	for _, existing := range AllowedPipelineStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PipelineStatus", value)
}

// NewPipelineStatusFromValue returns a pointer to a valid PipelineStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPipelineStatusFromValue(v string) (*PipelineStatus, error) {
	ev := PipelineStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PipelineStatus: valid values are %v", v, AllowedPipelineStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PipelineStatus) IsValid() bool {
	for _, existing := range AllowedPipelineStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PipelineStatus value
func (v PipelineStatus) Ptr() *PipelineStatus {
	return &v
}

type NullablePipelineStatus struct {
	value *PipelineStatus
	isSet bool
}

func (v NullablePipelineStatus) Get() *PipelineStatus {
	return v.value
}

func (v *NullablePipelineStatus) Set(val *PipelineStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineStatus(val *PipelineStatus) *NullablePipelineStatus {
	return &NullablePipelineStatus{value: val, isSet: true}
}

func (v NullablePipelineStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

