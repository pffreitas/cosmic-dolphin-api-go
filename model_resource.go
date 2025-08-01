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

// checks if the Resource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Resource{}

// Resource struct for Resource
type Resource struct {
	Type ResourceType `json:"type"`
	OpenGraph *OpenGraph `json:"openGraph,omitempty"`
}

type _Resource Resource

// NewResource instantiates a new Resource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResource(type_ ResourceType) *Resource {
	this := Resource{}
	this.Type = type_
	return &this
}

// NewResourceWithDefaults instantiates a new Resource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceWithDefaults() *Resource {
	this := Resource{}
	return &this
}

// GetType returns the Type field value
func (o *Resource) GetType() ResourceType {
	if o == nil {
		var ret ResourceType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Resource) GetTypeOk() (*ResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Resource) SetType(v ResourceType) {
	o.Type = v
}

// GetOpenGraph returns the OpenGraph field value if set, zero value otherwise.
func (o *Resource) GetOpenGraph() OpenGraph {
	if o == nil || IsNil(o.OpenGraph) {
		var ret OpenGraph
		return ret
	}
	return *o.OpenGraph
}

// GetOpenGraphOk returns a tuple with the OpenGraph field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetOpenGraphOk() (*OpenGraph, bool) {
	if o == nil || IsNil(o.OpenGraph) {
		return nil, false
	}
	return o.OpenGraph, true
}

// HasOpenGraph returns a boolean if a field has been set.
func (o *Resource) HasOpenGraph() bool {
	if o != nil && !IsNil(o.OpenGraph) {
		return true
	}

	return false
}

// SetOpenGraph gets a reference to the given OpenGraph and assigns it to the OpenGraph field.
func (o *Resource) SetOpenGraph(v OpenGraph) {
	o.OpenGraph = &v
}

func (o Resource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Resource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.OpenGraph) {
		toSerialize["openGraph"] = o.OpenGraph
	}
	return toSerialize, nil
}

func (o *Resource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varResource := _Resource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResource)

	if err != nil {
		return err
	}

	*o = Resource(varResource)

	return err
}

type NullableResource struct {
	value *Resource
	isSet bool
}

func (v NullableResource) Get() *Resource {
	return v.value
}

func (v *NullableResource) Set(val *Resource) {
	v.value = val
	v.isSet = true
}

func (v NullableResource) IsSet() bool {
	return v.isSet
}

func (v *NullableResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResource(val *Resource) *NullableResource {
	return &NullableResource{value: val, isSet: true}
}

func (v NullableResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


