/*
Cosmic Dolphin

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cosmicdolphinapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the PipelineStage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PipelineStage{}

// PipelineStage struct for PipelineStage
type PipelineStage struct {
	Id *int64 `json:"id,omitempty"`
	Name string `json:"name"`
	Key string `json:"key"`
	Status PipelineStatus `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type _PipelineStage PipelineStage

// NewPipelineStage instantiates a new PipelineStage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineStage(name string, key string, status PipelineStatus, createdAt time.Time, updatedAt time.Time) *PipelineStage {
	this := PipelineStage{}
	this.Name = name
	this.Key = key
	this.Status = status
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewPipelineStageWithDefaults instantiates a new PipelineStage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineStageWithDefaults() *PipelineStage {
	this := PipelineStage{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PipelineStage) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineStage) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PipelineStage) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *PipelineStage) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *PipelineStage) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PipelineStage) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PipelineStage) SetName(v string) {
	o.Name = v
}

// GetKey returns the Key field value
func (o *PipelineStage) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *PipelineStage) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *PipelineStage) SetKey(v string) {
	o.Key = v
}

// GetStatus returns the Status field value
func (o *PipelineStage) GetStatus() PipelineStatus {
	if o == nil {
		var ret PipelineStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PipelineStage) GetStatusOk() (*PipelineStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PipelineStage) SetStatus(v PipelineStatus) {
	o.Status = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PipelineStage) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PipelineStage) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PipelineStage) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *PipelineStage) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PipelineStage) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *PipelineStage) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o PipelineStage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PipelineStage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["key"] = o.Key
	toSerialize["status"] = o.Status
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *PipelineStage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"key",
		"status",
		"createdAt",
		"updatedAt",
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

	varPipelineStage := _PipelineStage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPipelineStage)

	if err != nil {
		return err
	}

	*o = PipelineStage(varPipelineStage)

	return err
}

type NullablePipelineStage struct {
	value *PipelineStage
	isSet bool
}

func (v NullablePipelineStage) Get() *PipelineStage {
	return v.value
}

func (v *NullablePipelineStage) Set(val *PipelineStage) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineStage) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineStage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineStage(val *PipelineStage) *NullablePipelineStage {
	return &NullablePipelineStage{value: val, isSet: true}
}

func (v NullablePipelineStage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineStage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


