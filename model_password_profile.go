/*
Libre Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v0.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
)

// PasswordProfile Password Profile associated with a user
type PasswordProfile struct {
	// If true the user is required to change their password upon the next login
	ForceChangePasswordNextSignIn *bool `json:"forceChangePasswordNextSignIn,omitempty"`
	// The user's password
	Password *string `json:"password,omitempty"`
}

// NewPasswordProfile instantiates a new PasswordProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordProfile() *PasswordProfile {
	this := PasswordProfile{}
	var forceChangePasswordNextSignIn bool = false
	this.ForceChangePasswordNextSignIn = &forceChangePasswordNextSignIn
	return &this
}

// NewPasswordProfileWithDefaults instantiates a new PasswordProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordProfileWithDefaults() *PasswordProfile {
	this := PasswordProfile{}
	var forceChangePasswordNextSignIn bool = false
	this.ForceChangePasswordNextSignIn = &forceChangePasswordNextSignIn
	return &this
}

// GetForceChangePasswordNextSignIn returns the ForceChangePasswordNextSignIn field value if set, zero value otherwise.
func (o *PasswordProfile) GetForceChangePasswordNextSignIn() bool {
	if o == nil || o.ForceChangePasswordNextSignIn == nil {
		var ret bool
		return ret
	}
	return *o.ForceChangePasswordNextSignIn
}

// GetForceChangePasswordNextSignInOk returns a tuple with the ForceChangePasswordNextSignIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordProfile) GetForceChangePasswordNextSignInOk() (*bool, bool) {
	if o == nil || o.ForceChangePasswordNextSignIn == nil {
		return nil, false
	}
	return o.ForceChangePasswordNextSignIn, true
}

// HasForceChangePasswordNextSignIn returns a boolean if a field has been set.
func (o *PasswordProfile) HasForceChangePasswordNextSignIn() bool {
	if o != nil && o.ForceChangePasswordNextSignIn != nil {
		return true
	}

	return false
}

// SetForceChangePasswordNextSignIn gets a reference to the given bool and assigns it to the ForceChangePasswordNextSignIn field.
func (o *PasswordProfile) SetForceChangePasswordNextSignIn(v bool) {
	o.ForceChangePasswordNextSignIn = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *PasswordProfile) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordProfile) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *PasswordProfile) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *PasswordProfile) SetPassword(v string) {
	o.Password = &v
}

func (o PasswordProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ForceChangePasswordNextSignIn != nil {
		toSerialize["forceChangePasswordNextSignIn"] = o.ForceChangePasswordNextSignIn
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullablePasswordProfile struct {
	value *PasswordProfile
	isSet bool
}

func (v NullablePasswordProfile) Get() *PasswordProfile {
	return v.value
}

func (v *NullablePasswordProfile) Set(val *PasswordProfile) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordProfile) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordProfile(val *PasswordProfile) *NullablePasswordProfile {
	return &NullablePasswordProfile{value: val, isSet: true}
}

func (v NullablePasswordProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


