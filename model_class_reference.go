/*
Libre Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
)

// ClassReference struct for ClassReference
type ClassReference struct {
	OdataId *string `json:"@odata.id,omitempty"`
}

// NewClassReference instantiates a new ClassReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClassReference() *ClassReference {
	this := ClassReference{}
	return &this
}

// NewClassReferenceWithDefaults instantiates a new ClassReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClassReferenceWithDefaults() *ClassReference {
	this := ClassReference{}
	return &this
}

// GetOdataId returns the OdataId field value if set, zero value otherwise.
func (o *ClassReference) GetOdataId() string {
	if o == nil || o.OdataId == nil {
		var ret string
		return ret
	}
	return *o.OdataId
}

// GetOdataIdOk returns a tuple with the OdataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClassReference) GetOdataIdOk() (*string, bool) {
	if o == nil || o.OdataId == nil {
		return nil, false
	}
	return o.OdataId, true
}

// HasOdataId returns a boolean if a field has been set.
func (o *ClassReference) HasOdataId() bool {
	if o != nil && o.OdataId != nil {
		return true
	}

	return false
}

// SetOdataId gets a reference to the given string and assigns it to the OdataId field.
func (o *ClassReference) SetOdataId(v string) {
	o.OdataId = &v
}

func (o ClassReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OdataId != nil {
		toSerialize["@odata.id"] = o.OdataId
	}
	return json.Marshal(toSerialize)
}

type NullableClassReference struct {
	value *ClassReference
	isSet bool
}

func (v NullableClassReference) Get() *ClassReference {
	return v.value
}

func (v *NullableClassReference) Set(val *ClassReference) {
	v.value = val
	v.isSet = true
}

func (v NullableClassReference) IsSet() bool {
	return v.isSet
}

func (v *NullableClassReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClassReference(val *ClassReference) *NullableClassReference {
	return &NullableClassReference{value: val, isSet: true}
}

func (v NullableClassReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClassReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
