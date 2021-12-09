/*
Libre Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v0.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package opengraph

import (
	"encoding/json"
)

// Hashes Hashes of the file's binary content, if available. Read-only.
type Hashes struct {
	// The CRC32 value of the file (if available). Read-only.
	Crc32Hash *string `json:"crc32Hash,omitempty"`
	// A proprietary hash of the file that can be used to determine if the contents of the file have changed (if available). Read-only.
	QuickXorHash *string `json:"quickXorHash,omitempty"`
	// SHA1 hash for the contents of the file (if available). Read-only.
	Sha1Hash *string `json:"sha1Hash,omitempty"`
	// SHA256 hash for the contents of the file (if available). Read-only.
	Sha256Hash *string `json:"sha256Hash,omitempty"`
}

// NewHashes instantiates a new Hashes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHashes() *Hashes {
	this := Hashes{}
	return &this
}

// NewHashesWithDefaults instantiates a new Hashes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHashesWithDefaults() *Hashes {
	this := Hashes{}
	return &this
}

// GetCrc32Hash returns the Crc32Hash field value if set, zero value otherwise.
func (o *Hashes) GetCrc32Hash() string {
	if o == nil || o.Crc32Hash == nil {
		var ret string
		return ret
	}
	return *o.Crc32Hash
}

// GetCrc32HashOk returns a tuple with the Crc32Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hashes) GetCrc32HashOk() (*string, bool) {
	if o == nil || o.Crc32Hash == nil {
		return nil, false
	}
	return o.Crc32Hash, true
}

// HasCrc32Hash returns a boolean if a field has been set.
func (o *Hashes) HasCrc32Hash() bool {
	if o != nil && o.Crc32Hash != nil {
		return true
	}

	return false
}

// SetCrc32Hash gets a reference to the given string and assigns it to the Crc32Hash field.
func (o *Hashes) SetCrc32Hash(v string) {
	o.Crc32Hash = &v
}

// GetQuickXorHash returns the QuickXorHash field value if set, zero value otherwise.
func (o *Hashes) GetQuickXorHash() string {
	if o == nil || o.QuickXorHash == nil {
		var ret string
		return ret
	}
	return *o.QuickXorHash
}

// GetQuickXorHashOk returns a tuple with the QuickXorHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hashes) GetQuickXorHashOk() (*string, bool) {
	if o == nil || o.QuickXorHash == nil {
		return nil, false
	}
	return o.QuickXorHash, true
}

// HasQuickXorHash returns a boolean if a field has been set.
func (o *Hashes) HasQuickXorHash() bool {
	if o != nil && o.QuickXorHash != nil {
		return true
	}

	return false
}

// SetQuickXorHash gets a reference to the given string and assigns it to the QuickXorHash field.
func (o *Hashes) SetQuickXorHash(v string) {
	o.QuickXorHash = &v
}

// GetSha1Hash returns the Sha1Hash field value if set, zero value otherwise.
func (o *Hashes) GetSha1Hash() string {
	if o == nil || o.Sha1Hash == nil {
		var ret string
		return ret
	}
	return *o.Sha1Hash
}

// GetSha1HashOk returns a tuple with the Sha1Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hashes) GetSha1HashOk() (*string, bool) {
	if o == nil || o.Sha1Hash == nil {
		return nil, false
	}
	return o.Sha1Hash, true
}

// HasSha1Hash returns a boolean if a field has been set.
func (o *Hashes) HasSha1Hash() bool {
	if o != nil && o.Sha1Hash != nil {
		return true
	}

	return false
}

// SetSha1Hash gets a reference to the given string and assigns it to the Sha1Hash field.
func (o *Hashes) SetSha1Hash(v string) {
	o.Sha1Hash = &v
}

// GetSha256Hash returns the Sha256Hash field value if set, zero value otherwise.
func (o *Hashes) GetSha256Hash() string {
	if o == nil || o.Sha256Hash == nil {
		var ret string
		return ret
	}
	return *o.Sha256Hash
}

// GetSha256HashOk returns a tuple with the Sha256Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hashes) GetSha256HashOk() (*string, bool) {
	if o == nil || o.Sha256Hash == nil {
		return nil, false
	}
	return o.Sha256Hash, true
}

// HasSha256Hash returns a boolean if a field has been set.
func (o *Hashes) HasSha256Hash() bool {
	if o != nil && o.Sha256Hash != nil {
		return true
	}

	return false
}

// SetSha256Hash gets a reference to the given string and assigns it to the Sha256Hash field.
func (o *Hashes) SetSha256Hash(v string) {
	o.Sha256Hash = &v
}

func (o Hashes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Crc32Hash != nil {
		toSerialize["crc32Hash"] = o.Crc32Hash
	}
	if o.QuickXorHash != nil {
		toSerialize["quickXorHash"] = o.QuickXorHash
	}
	if o.Sha1Hash != nil {
		toSerialize["sha1Hash"] = o.Sha1Hash
	}
	if o.Sha256Hash != nil {
		toSerialize["sha256Hash"] = o.Sha256Hash
	}
	return json.Marshal(toSerialize)
}

type NullableHashes struct {
	value *Hashes
	isSet bool
}

func (v NullableHashes) Get() *Hashes {
	return v.value
}

func (v *NullableHashes) Set(val *Hashes) {
	v.value = val
	v.isSet = true
}

func (v NullableHashes) IsSet() bool {
	return v.isSet
}

func (v *NullableHashes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHashes(val *Hashes) *NullableHashes {
	return &NullableHashes{value: val, isSet: true}
}

func (v NullableHashes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHashes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


