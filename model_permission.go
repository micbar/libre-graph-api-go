/*
Libre Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v0.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
)

// Permission The Permission resource provides information about a sharing permission granted for a DriveItem resource.
type Permission struct {
	GrantedTo []IdentitySet `json:"grantedTo,omitempty"`
	Roles []string `json:"roles,omitempty"`
}

// NewPermission instantiates a new Permission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermission() *Permission {
	this := Permission{}
	return &this
}

// NewPermissionWithDefaults instantiates a new Permission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionWithDefaults() *Permission {
	this := Permission{}
	return &this
}

// GetGrantedTo returns the GrantedTo field value if set, zero value otherwise.
func (o *Permission) GetGrantedTo() []IdentitySet {
	if o == nil || o.GrantedTo == nil {
		var ret []IdentitySet
		return ret
	}
	return o.GrantedTo
}

// GetGrantedToOk returns a tuple with the GrantedTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permission) GetGrantedToOk() ([]IdentitySet, bool) {
	if o == nil || o.GrantedTo == nil {
		return nil, false
	}
	return o.GrantedTo, true
}

// HasGrantedTo returns a boolean if a field has been set.
func (o *Permission) HasGrantedTo() bool {
	if o != nil && o.GrantedTo != nil {
		return true
	}

	return false
}

// SetGrantedTo gets a reference to the given []IdentitySet and assigns it to the GrantedTo field.
func (o *Permission) SetGrantedTo(v []IdentitySet) {
	o.GrantedTo = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *Permission) GetRoles() []string {
	if o == nil || o.Roles == nil {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Permission) GetRolesOk() ([]string, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *Permission) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *Permission) SetRoles(v []string) {
	o.Roles = v
}

func (o Permission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GrantedTo != nil {
		toSerialize["grantedTo"] = o.GrantedTo
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	return json.Marshal(toSerialize)
}

type NullablePermission struct {
	value *Permission
	isSet bool
}

func (v NullablePermission) Get() *Permission {
	return v.value
}

func (v *NullablePermission) Set(val *Permission) {
	v.value = val
	v.isSet = true
}

func (v NullablePermission) IsSet() bool {
	return v.isSet
}

func (v *NullablePermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermission(val *Permission) *NullablePermission {
	return &NullablePermission{value: val, isSet: true}
}

func (v NullablePermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


