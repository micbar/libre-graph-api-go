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

// CollectionOfGroup struct for CollectionOfGroup
type CollectionOfGroup struct {
	Value         []Group `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfGroup instantiates a new CollectionOfGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfGroup() *CollectionOfGroup {
	this := CollectionOfGroup{}
	return &this
}

// NewCollectionOfGroupWithDefaults instantiates a new CollectionOfGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfGroupWithDefaults() *CollectionOfGroup {
	this := CollectionOfGroup{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfGroup) GetValue() []Group {
	if o == nil || o.Value == nil {
		var ret []Group
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfGroup) GetValueOk() ([]Group, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfGroup) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []Group and assigns it to the Value field.
func (o *CollectionOfGroup) SetValue(v []Group) {
	o.Value = v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfGroup) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfGroup) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfGroup) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfGroup) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfGroup struct {
	value *CollectionOfGroup
	isSet bool
}

func (v NullableCollectionOfGroup) Get() *CollectionOfGroup {
	return v.value
}

func (v *NullableCollectionOfGroup) Set(val *CollectionOfGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfGroup(val *CollectionOfGroup) *NullableCollectionOfGroup {
	return &NullableCollectionOfGroup{value: val, isSet: true}
}

func (v NullableCollectionOfGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
