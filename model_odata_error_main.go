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

// OdataErrorMain struct for OdataErrorMain
type OdataErrorMain struct {
	Code    string             `json:"code"`
	Message string             `json:"message"`
	Target  *string            `json:"target,omitempty"`
	Details []OdataErrorDetail `json:"details,omitempty"`
	// The structure of this object is service-specific
	Innererror map[string]interface{} `json:"innererror,omitempty"`
}

// NewOdataErrorMain instantiates a new OdataErrorMain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOdataErrorMain(code string, message string) *OdataErrorMain {
	this := OdataErrorMain{}
	this.Code = code
	this.Message = message
	return &this
}

// NewOdataErrorMainWithDefaults instantiates a new OdataErrorMain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOdataErrorMainWithDefaults() *OdataErrorMain {
	this := OdataErrorMain{}
	return &this
}

// GetCode returns the Code field value
func (o *OdataErrorMain) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *OdataErrorMain) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *OdataErrorMain) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *OdataErrorMain) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *OdataErrorMain) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *OdataErrorMain) SetMessage(v string) {
	o.Message = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *OdataErrorMain) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OdataErrorMain) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *OdataErrorMain) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *OdataErrorMain) SetTarget(v string) {
	o.Target = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *OdataErrorMain) GetDetails() []OdataErrorDetail {
	if o == nil || o.Details == nil {
		var ret []OdataErrorDetail
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OdataErrorMain) GetDetailsOk() ([]OdataErrorDetail, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *OdataErrorMain) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []OdataErrorDetail and assigns it to the Details field.
func (o *OdataErrorMain) SetDetails(v []OdataErrorDetail) {
	o.Details = v
}

// GetInnererror returns the Innererror field value if set, zero value otherwise.
func (o *OdataErrorMain) GetInnererror() map[string]interface{} {
	if o == nil || o.Innererror == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Innererror
}

// GetInnererrorOk returns a tuple with the Innererror field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OdataErrorMain) GetInnererrorOk() (map[string]interface{}, bool) {
	if o == nil || o.Innererror == nil {
		return nil, false
	}
	return o.Innererror, true
}

// HasInnererror returns a boolean if a field has been set.
func (o *OdataErrorMain) HasInnererror() bool {
	if o != nil && o.Innererror != nil {
		return true
	}

	return false
}

// SetInnererror gets a reference to the given map[string]interface{} and assigns it to the Innererror field.
func (o *OdataErrorMain) SetInnererror(v map[string]interface{}) {
	o.Innererror = v
}

func (o OdataErrorMain) MarshalJSON() ([]byte, error) {
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
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Innererror != nil {
		toSerialize["innererror"] = o.Innererror
	}
	return json.Marshal(toSerialize)
}

type NullableOdataErrorMain struct {
	value *OdataErrorMain
	isSet bool
}

func (v NullableOdataErrorMain) Get() *OdataErrorMain {
	return v.value
}

func (v *NullableOdataErrorMain) Set(val *OdataErrorMain) {
	v.value = val
	v.isSet = true
}

func (v NullableOdataErrorMain) IsSet() bool {
	return v.isSet
}

func (v *NullableOdataErrorMain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOdataErrorMain(val *OdataErrorMain) *NullableOdataErrorMain {
	return &NullableOdataErrorMain{value: val, isSet: true}
}

func (v NullableOdataErrorMain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOdataErrorMain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
