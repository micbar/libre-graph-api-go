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

// CollectionOfUser struct for CollectionOfUser
type CollectionOfUser struct {
	Value         []User  `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfUser instantiates a new CollectionOfUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfUser() *CollectionOfUser {
	this := CollectionOfUser{}
	return &this
}

// NewCollectionOfUserWithDefaults instantiates a new CollectionOfUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfUserWithDefaults() *CollectionOfUser {
	this := CollectionOfUser{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfUser) GetValue() []User {
	if o == nil || o.Value == nil {
		var ret []User
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfUser) GetValueOk() ([]User, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfUser) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []User and assigns it to the Value field.
func (o *CollectionOfUser) SetValue(v []User) {
	o.Value = v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfUser) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfUser) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfUser) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfUser) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfUser struct {
	value *CollectionOfUser
	isSet bool
}

func (v NullableCollectionOfUser) Get() *CollectionOfUser {
	return v.value
}

func (v *NullableCollectionOfUser) Set(val *CollectionOfUser) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfUser) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfUser(val *CollectionOfUser) *NullableCollectionOfUser {
	return &NullableCollectionOfUser{value: val, isSet: true}
}

func (v NullableCollectionOfUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
