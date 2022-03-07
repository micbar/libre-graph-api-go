/*
Libre Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v0.12.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
)

// Deleted Information about the deleted state of the item. Read-only.
type Deleted struct {
	// Represents the state of the deleted item.
	State *string `json:"state,omitempty"`
}

// NewDeleted instantiates a new Deleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleted() *Deleted {
	this := Deleted{}
	return &this
}

// NewDeletedWithDefaults instantiates a new Deleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeletedWithDefaults() *Deleted {
	this := Deleted{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Deleted) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deleted) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Deleted) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Deleted) SetState(v string) {
	o.State = &v
}

func (o Deleted) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableDeleted struct {
	value *Deleted
	isSet bool
}

func (v NullableDeleted) Get() *Deleted {
	return v.value
}

func (v *NullableDeleted) Set(val *Deleted) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleted) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleted(val *Deleted) *NullableDeleted {
	return &NullableDeleted{value: val, isSet: true}
}

func (v NullableDeleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


