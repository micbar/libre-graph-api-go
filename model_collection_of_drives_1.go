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

// CollectionOfDrives1 struct for CollectionOfDrives1
type CollectionOfDrives1 struct {
	Value []Drive `json:"value,omitempty"`
}

// NewCollectionOfDrives1 instantiates a new CollectionOfDrives1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfDrives1() *CollectionOfDrives1 {
	this := CollectionOfDrives1{}
	return &this
}

// NewCollectionOfDrives1WithDefaults instantiates a new CollectionOfDrives1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfDrives1WithDefaults() *CollectionOfDrives1 {
	this := CollectionOfDrives1{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfDrives1) GetValue() []Drive {
	if o == nil || o.Value == nil {
		var ret []Drive
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfDrives1) GetValueOk() ([]Drive, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfDrives1) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []Drive and assigns it to the Value field.
func (o *CollectionOfDrives1) SetValue(v []Drive) {
	o.Value = v
}

func (o CollectionOfDrives1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfDrives1 struct {
	value *CollectionOfDrives1
	isSet bool
}

func (v NullableCollectionOfDrives1) Get() *CollectionOfDrives1 {
	return v.value
}

func (v *NullableCollectionOfDrives1) Set(val *CollectionOfDrives1) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfDrives1) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfDrives1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfDrives1(val *CollectionOfDrives1) *NullableCollectionOfDrives1 {
	return &NullableCollectionOfDrives1{value: val, isSet: true}
}

func (v NullableCollectionOfDrives1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfDrives1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


