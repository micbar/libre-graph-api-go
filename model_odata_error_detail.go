/*
Libre Graph API

Libre Graph is a free API for cloud collaboration inspired by the MS Graph API.

API version: v1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"encoding/json"
)

// OdataErrorDetail struct for OdataErrorDetail
type OdataErrorDetail struct {
	Code    string  `json:"code"`
	Message string  `json:"message"`
	Target  *string `json:"target,omitempty"`
}

// NewOdataErrorDetail instantiates a new OdataErrorDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOdataErrorDetail(code string, message string) *OdataErrorDetail {
	this := OdataErrorDetail{}
	this.Code = code
	this.Message = message
	return &this
}

// NewOdataErrorDetailWithDefaults instantiates a new OdataErrorDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOdataErrorDetailWithDefaults() *OdataErrorDetail {
	this := OdataErrorDetail{}
	return &this
}

// GetCode returns the Code field value
func (o *OdataErrorDetail) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *OdataErrorDetail) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *OdataErrorDetail) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *OdataErrorDetail) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *OdataErrorDetail) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *OdataErrorDetail) SetMessage(v string) {
	o.Message = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *OdataErrorDetail) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OdataErrorDetail) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *OdataErrorDetail) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *OdataErrorDetail) SetTarget(v string) {
	o.Target = &v
}

func (o OdataErrorDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}

type NullableOdataErrorDetail struct {
	value *OdataErrorDetail
	isSet bool
}

func (v NullableOdataErrorDetail) Get() *OdataErrorDetail {
	return v.value
}

func (v *NullableOdataErrorDetail) Set(val *OdataErrorDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableOdataErrorDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableOdataErrorDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOdataErrorDetail(val *OdataErrorDetail) *NullableOdataErrorDetail {
	return &NullableOdataErrorDetail{value: val, isSet: true}
}

func (v NullableOdataErrorDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOdataErrorDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
