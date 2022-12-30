/*
Libre Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
)

// TagAssignment struct for TagAssignment
type TagAssignment struct {
	ResourceId string `json:"resourceId"`
	Tags []string `json:"tags"`
}

// NewTagAssignment instantiates a new TagAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagAssignment(resourceId string, tags []string) *TagAssignment {
	this := TagAssignment{}
	this.ResourceId = resourceId
	this.Tags = tags
	return &this
}

// NewTagAssignmentWithDefaults instantiates a new TagAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagAssignmentWithDefaults() *TagAssignment {
	this := TagAssignment{}
	return &this
}

// GetResourceId returns the ResourceId field value
func (o *TagAssignment) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *TagAssignment) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *TagAssignment) SetResourceId(v string) {
	o.ResourceId = v
}

// GetTags returns the Tags field value
func (o *TagAssignment) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *TagAssignment) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *TagAssignment) SetTags(v []string) {
	o.Tags = v
}

func (o TagAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resourceId"] = o.ResourceId
	}
	if true {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableTagAssignment struct {
	value *TagAssignment
	isSet bool
}

func (v NullableTagAssignment) Get() *TagAssignment {
	return v.value
}

func (v *NullableTagAssignment) Set(val *TagAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableTagAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableTagAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagAssignment(val *TagAssignment) *NullableTagAssignment {
	return &NullableTagAssignment{value: val, isSet: true}
}

func (v NullableTagAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

